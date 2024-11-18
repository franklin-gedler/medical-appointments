package logger

import (
	"context"
	"fmt"
	"math"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type contextKey struct{}

var loggerKey = &contextKey{}

// LoggerConfig define si queremos salida en JSON o en formato personalizado.
type LoggerConfig struct {
	UseJSON bool
}

// InitializeLogger inicializa el logger basado en la configuración.
func InitializeLogger(config LoggerConfig) *zap.Logger {
	var encoder zapcore.Encoder

	if config.UseJSON {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = newCustomTextEncoder()
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	return zap.New(core, zap.AddCaller())
}

// InitializeLoggerContext inicializa el contexto con el logger.
func InitializeLoggerContext(config LoggerConfig) context.Context {
	logger := InitializeLogger(config)
	return context.WithValue(context.Background(), loggerKey, logger)
}

// FromContext extrae el logger del contexto. Si no existe, devuelve un logger predeterminado.
func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}

// customTextEncoder implementa un formato de salida personalizado para zapcore.
type customTextEncoder struct {
	zapcore.Encoder
}

// newCustomTextEncoder crea un encoder que usa customTextEncoder.
func newCustomTextEncoder() zapcore.Encoder {
	return &customTextEncoder{}
}

// bufferPool nos ayuda a administrar los buffers para evitar asignaciones constantes.
var bufferPool = buffer.NewPool()

func (enc *customTextEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// Obtiene un nuevo buffer de la pool
	line := bufferPool.Get()
	timestamp := entry.Time.Format("2006-01-02T15:04:05.000Z0700")
	level := entry.Level.CapitalString()
	caller := entry.Caller.TrimmedPath()
	line.AppendString("[" + timestamp + "][level: " + level + "][caller: " + caller + "][msg: " + entry.Message + "]")
	for _, field := range fields {
		//line.AppendString("[" + field.Key + ": " + field.String + "]")
		line.AppendString("[" + field.Key + ": " + fieldToString(field) + "]")
	}
	line.AppendString("\n")
	return line, nil
}

// Función de conversión de campos a cadenas
type fieldConverter func(field zapcore.Field) string

// Mapa de funciones de conversión
var fieldConverters = map[zapcore.FieldType]fieldConverter{
	zapcore.BoolType:  func(field zapcore.Field) string { return fmt.Sprintf("%v", field.Integer == 1) },
	zapcore.Int64Type: func(field zapcore.Field) string { return fmt.Sprintf("%d", field.Integer) },
	zapcore.Int32Type: func(field zapcore.Field) string { return fmt.Sprintf("%d", field.Integer) },
	zapcore.Int16Type: func(field zapcore.Field) string { return fmt.Sprintf("%d", field.Integer) },
	zapcore.Int8Type:  func(field zapcore.Field) string { return fmt.Sprintf("%d", field.Integer) },
	zapcore.Float64Type: func(field zapcore.Field) string {
		return fmt.Sprintf("%f", math.Float64frombits(uint64(field.Integer)))
	},
	zapcore.Float32Type: func(field zapcore.Field) string {
		return fmt.Sprintf("%f", math.Float32frombits(uint32(field.Integer)))
	},
	zapcore.StringType: func(field zapcore.Field) string { return field.String },
	zapcore.ErrorType:  func(field zapcore.Field) string { return field.Interface.(error).Error() },
	// Agrega más tipos según sea necesario
}

func fieldToString(field zapcore.Field) string {
	if converter, ok := fieldConverters[field.Type]; ok {
		return converter(field)
	}
	return fmt.Sprintf("%v", field.Interface)
}

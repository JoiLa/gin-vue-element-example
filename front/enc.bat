javascript-obfuscator ./dist
--output ./obfuscated
--compact true
--control-flow-flattening true
--control-flow-flattening-threshold 0.75
--dead-code-injection true














:: 高度混淆，低性能(官方文档)
javascript-obfuscator index.html -o index.html --compact true --control-flow-flattening true --control-flow-flattening-threshold 1 --dead-code-injection true --dead-code-injection-threshold 1 --debug-protection true --debug-protection-interval true --disable-console-output true --identifier-names-generator hexadecimal --log false --rename-globals false --rotate-string-array true --self-defending true --string-array true --string-array-encoding rc4 --string-array-threshold 1 --transform-object-keys true --unicode-escape-sequence false

::中等混淆，最佳性能(官方文档)
javascript-obfuscator ./dist -o ./dist-obfuscator
--compact true
--control-flow-flattening true
--control-flow-flattening-threshold 0.5
--dead-code-injection true
--dead-code-injection-threshold 0.6
--debug-protection false
--debug-protection-interval false
--disable-console-output true
--identifier-names-generator mangled
--log false
--rename-globals false
--rotate-string-array true
--self-defending true
--string-array true
--string-array-encoding rc4
--string-array-threshold 0.75
--transform-object-keys true
--unicode-escape-sequence false














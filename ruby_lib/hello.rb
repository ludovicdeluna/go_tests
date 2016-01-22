require 'fiddle'
require 'fiddle/import'

module Tester
  extend Fiddle::Importer

  # Below, for each DSL -> sample without Importer
  begin
    dlload 'my_lib.so'                      # @my_lib = Fiddle.dlopen("my_lib.so")
  rescue
    puts "Please, compile main.go before launch ruby script (since go1.5) :"
    puts "go build -buildmode=c-shared -o my_lib.so main.go"
    exit 1
  end

  extern '*void sayHello()'               # sayHello = Fiddle::Function.new(@my_lib["sayHello"], [], Fiddle::TYPE_VOIDP)
  extern 'void clearChar(char * pointer)' # clearChar = Fiddle::Function.new(@my_lib["clearChar"], [Fiddle::TYPE_VOIDP], Fiddle::TYPE_VOID)

  def self.hello
    result = ""
    clearChar(
      # return *C.Char pointer (Not GC by Ruby nor by Go runtime)
      sayHello.tap do |pointer|
        result = pointer.to_s
      end
      # Resulting pointer. It's memory is deallocated by C lib stdlib:free call
    )
    return result
  end
end

puts Tester.hello

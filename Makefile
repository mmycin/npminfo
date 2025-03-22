build:
	@cmd /c if not exist bin/ mkdir bin
	@nim c -d:release -o:npmdc main.nim
	@cmd /c move npmdc.exe bin
	@cmd /c echo Building Successful

run:
	@cmd /c cls
	@nim r --verbosity:0 main.nim
	@cmd /c pause
	
buildandrun: build
	@./bin/npmdc.exe

clean:
	@cmd /c echo Cleaning the project...
	@cmd /c if exist bin rmdir /s /q bin
	@cmd /c echo Cleaning Success
:: Please install upx first, https://github.com/upx/upx/releases
for %%* in (.) do set CurrDirName=%%~nx*
echo %CurrDirName%
for /f "delims=" %%i in ('dir /b /a-d /s "%CurrDirName%*"') do upx --best "%%i"

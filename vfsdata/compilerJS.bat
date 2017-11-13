::@echo off  
rem 正在搜索...  
rem 压缩文件  
for /f "delims=" %%i in ('dir /b /a-d /s "*.js"') do java -jar compiler.jar --js_output_file="%%i" "%%i"
rem 压缩完毕  
rem 输出完成
pause  
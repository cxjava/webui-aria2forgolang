::@echo off  
rem ��������...  
rem ѹ���ļ�  
for /f "delims=" %%i in ('dir /b /a-d /s "*.js"') do java -jar compiler.jar --js_output_file="%%i" "%%i"
rem ѹ�����  
rem ������
pause  
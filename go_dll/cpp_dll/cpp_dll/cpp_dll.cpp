// cpp_dll.cpp : ���� DLL Ӧ�ó���ĵ���������
//

#include "stdafx.h"
#include <stdio.h>

extern "C" __declspec(dllexport) int add(int a, int b)
{
	return (a + b);
}

extern "C" __declspec(dllexport) int sub(int a, int b)
{
	return (a - b);
}

// ����golangָ�봫��
extern "C" __declspec(dllexport) void * point(void *ctx){
	printf("ctx:%p\n", ctx);
	return ctx;
}


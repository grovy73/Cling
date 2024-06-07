#include <raylib.h>

// CLING GUIDE
// libs raylib

int main() {
	InitWindow(800, 600, "Balls");
	while (!WindowShouldClose()) {
		BeginDrawing();
			ClearBackground(RAYWHITE);
		EndDrawing();
	}
	CloseWindow();
}

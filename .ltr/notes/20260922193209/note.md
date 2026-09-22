# Ebiten basics

## Kako ebiten funkcionira?

Imamo tri glavne metode koje napravimo:

- `Update()`
- `Draw()`
- `Layout()`

Update se poziva 60 puta "ticks" po sekundi

Draw se poziva ovisno o refresh rate-u monitora, znaci meni 180 puta u sekundi odnosno on je FPS, dock je update TPS

Zasto?

Tako da logika uvijek vrti 60 puta u sekundi i da je konzistentno, dok frame rate / draw bude uskladen sa monitorom.

---

```
1. Layout (if needed)  —  window created or resized?
2. Update              —  run your game logic (60 times/sec by default)
3. Draw                —  render the current state to the screen
4. Swap buffers        —  show the new frame to the player
5. Repeat
```

---

U **Update** citamo input i mjenjamo game state.
U **Draw** samo citamo trenutni state i renderiramo. Ni slucajno mutiramo nista u **Draw**

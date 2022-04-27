package kata

/* not mine but a clever solution */

func Goals(g ...int) (r int) {
  for _, n := range g { r += n }
  return
}

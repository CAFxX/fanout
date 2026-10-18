package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 int, p1 uint16, p2 int32) int32 {
    var v7 int = (^(p0))
    v8 := []int{p0, v7, (v7 & 270)}
    var v9 uint8 = 0
    v7 -= 127
    v10 := GenS1{F0: ("\n\t" + ""), F1: ((!((823 != 32767) && (!((false || ("hello" > "")) || false)))) || (310 == 303))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 255
    var v3 int = -9223372036854775808
    v4 := []string{("foo bar" + "\n\t")}
    v5 := []int{}
    v6 := []string{"\n\t"}
    v3 -= (v3 & -2147483648)
    fmt.Println(v2, v1, v1)
    var v11 int16 = 344
    var v12 float32 = 3.14159
    fmt.Println(v11)
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    var v13 string = "hello"
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 int = 4294967295
        var v16 int16 = v11
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

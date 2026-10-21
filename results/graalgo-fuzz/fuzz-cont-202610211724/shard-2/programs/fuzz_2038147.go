package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 string }
func genFn0(p0 uint, p1 int, p2 string) int16 {
    fmt.Println(p1)
    var v6 uint8 = 127
    v7 := []int{p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 611
    var v2 int8 = 42
    var v3 int = -2147483648
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    v8 := GenS1{F0: 65535}
    var v9 uint64 = 631
    v1 -= (v1 & v1)
    fmt.Println(v9)
    var v10 float32 = -2.25
    var v11 bool = true
    v10 += (3.14159 * v10)
    v4[""] = (!true)
    switch v3 {
    case 0:
        fmt.Println(v3, v9, v9)
    case -2147483648:
        var v12 uint8 = 1
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        switch "hello" {
        case "a":
            for fi14 := 0; fi14 < 16; fi14++ {
                fmt.Println(v3)
                v15 := []int{(fi13 ^ fi14), (v3 >> ((18446744073709551615) & 63)), 127, (fi13 * (fi13 + fi13))}
                _ = fi14
                _ = v15
            }
        default:
            v2 = v2
        }
        _ = fi13
    }
    var v16 string = (("a" + "") + (("foo bar" + "hello") + "a"))
    var v17 int8 = (v2 >> ((4294967295) & 63))
    v9 = (v9 << ((4294967295) & 63))
    v18 := GenS2{F0: -852, F1: (^(v17)), F2: v16}
    if ((v16 > (v16 + ((v16 + "") + v16))) && ((((v1 - 32767) < 65535) && v11) && true)) {
        v5[65535] = v3
    } else {
        fmt.Println(v3, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

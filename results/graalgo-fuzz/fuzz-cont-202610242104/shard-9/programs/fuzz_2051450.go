package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int16; F1 uint16; F2 bool }
func genFn0(p0 uint32) int {
    v5 := GenS1{F0: "", F1: 0}
    v6 := GenS2{F0: (0.1 > 1.5), F1: -1}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float64 = -2.25
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 int64 = -426
    v2 := make(map[string]bool)
    v3 := []string{"a", "\n\t", "\n\t", "hello"}
    var v4 int16 = -32768
    v4 = 42
    var v9 uint16 = 987
    v4 -= (v4 >> ((65535) & 63))
    v10 := GenS1{F0: ("\n\t" + ""), F1: 2}
    var v11 int = 196
    var v12 int8 = 42
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint = 42
        v15 := GenS3{F0: (318 == 1), F1: v14, F2: (v11 - genFn0(uint32(v9)))}
        _ = fi13
        _ = v14
        _ = v15
    }
    v2[("foo bar" + ((("\n\t" + "\n\t") + "\n\t") + "hello"))] = (!true)
    v11 = v11
    v4 -= (-(v4))
    v4 += v4
    for fi16 := 0; fi16 < 20; fi16++ {
        if ((!((v4 ^ 1000) >= (v4 * v4))) && true) {
            var v17 int16 = (-(v4))
            v18 := GenS4{F0: ("foo bar" + "hello"), F1: v1}
            _ = v17
            _ = v18
        } else {
            if ((!(3.14159 > -2.25)) && false) {
                var v19 uint = 18446744073709551615
                _ = v19
            }
            v20 := v2["foo bar"]
            switch int(v12) {
            case 32768:
                _ = len(v3)
            case 32767:
                v21 := GenS5{F0: (v4 * v4), F1: v9, F2: (0 >= 1)}
                _ = v21
            }
            _ = v20
        }
        _ = fi16
    }
    fmt.Println(v1, v12)
    fmt.Println(v1, v4, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}

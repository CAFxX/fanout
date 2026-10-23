package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 int) int {
    var v5 bool = (true && true)
    p1 -= p1
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 uint8, p1 int8) int32 {
    if true {
        var v6 int32 = 1
        p0 += 6
        v6 += (v6 >> ((255) & 63))
        _ = v6
    }
    p1 += p1
    var v7 uint8 = 42
    var v8 bool = true
    v9 := GenS1{F0: 1, F1: (!(713 > 2147483647))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    var v3 string = ("\n\t" + (("a" + "") + "\n\t"))
    var v4 float64 = 0.1
    v3 = ""
    var v10 bool = false
    var v11 float64 = float64(1e10)
    for fi12 := 0; fi12 < 9; fi12++ {
        v4 = (v4 - ((float64(-2.25) / (float64(((1.5 - float32((0.0 - (v4 * 0.1)))) / (3.14159 + 0.0))) - 1.5)) - ((v11 - v4) + (0.1 / 3.14159))))
        _ = fi12
    }
    var v13 uint32 = 204
    v14 := GenS1{F0: 4294967295, F1: false}
    v3 = (((v3 + v3) + "") + (v3 + "hello"))
    switch v3 {
    case "foo bar":
        for fi15 := 0; fi15 < 20; fi15++ {
            v1 = (v1 | 631)
            _ = fi15
        }
    case "a":
        if v10 {
            v4 += ((float64((-2.25 - -2.25)) / 1e10) * 0.1)
            if ((127 > 255) && (127 <= genFn0(18446744073709551615, v13, 128))) {
                v16 := GenS2{F0: 65535, F1: 1}
                var v17 float64 = ((v4 + (float64(1e10) + v4)) / v11)
                _ = v16
                _ = v17
            }
        }
    case "\n\t":
        v10 = (!v10)
    }
    v18 := GenS2{F0: 0, F1: 145}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

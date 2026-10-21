package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 bool }
func genFn0() uint {
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 42
        if (!(42 <= 308)) {
            v8 -= v8
            v9 := []int{42, 2147483647}
            _ = v9
        }
        v10 := []int{(fi7 + fi7), (fi7 << ((255) & 63))}
        _ = fi7
        _ = v10
        _ = v8
    }
    _ = v6
    return 65535
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    var v3 uint64 = 461
    var v4 int16 = 42
    v5 := make(map[int]int)
    var v11 uint32 = 42
    if (109 == 95) {
        switch "a" {
        case "hello":
            if ((((0.1 / 0.1) + 0.0) * 1e10) >= 0.1) {
                switch -494 {
                case -2147483648:
                    var v12 int = int(v3)
                    _ = v12
                }
                fmt.Println(v3, v11, v3)
            }
        case "":
            var v13 uint64 = 406
            _ = v13
        case "a":
            v11 += (v11 ^ v11)
        default:
            v11 += (v11 >> ((42) & 63))
        }
    }
    var v14 bool = ((v4 >> ((18446744073709551615) & 63)) != -544)
    v15 := GenS1{F0: 42, F1: ""}
    v4 += (-(v4))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 uint = 1
    if v14 {
        v1 = append(v1, v11)
        for fi17 := 0; fi17 < 11; fi17++ {
            _ = len(v2)
            var v18 float64 = ((3.14159 * 1e10) / -2.25)
            _ = fi17
            _ = v18
        }
    }
    var v19 uint32 = (^(v11))
    v20 := GenS1{F0: -128, F1: (("a" + (("" + "\n\t") + ("hello" + "\n\t"))) + ("hello" + "foo bar"))}
    var v21 int64 = 42
    fmt.Println(v3)
    var v22 string = "a"
    v23 := GenS2{F0: 9223372036854775807, F1: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}

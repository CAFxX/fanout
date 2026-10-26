package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }

func main() {
    var v1 int = 42
    var v2 float64 = 0.0
    var v3 int64 = 0
    var v4 uint32 = 2147483647
    var v5 int8 = 44
    v5 = 21
    v6 := []int{}
    if (true || ((744 <= 690) || false)) {
        fmt.Println(v1)
        var v7 int64 = (v3 % ((0) | 1))
        var v8 int32 = 1
        _ = v7
        _ = v8
    } else {
        var v9 string = ("" + "a")
        _ = v9
    }
    if true {
        v3 -= (v3 * 1)
        switch (-(v1)) {
        case 32768:
            v1 = 42
        case 657:
            if ((!(1 < 597)) || true) {
                v2 += v2
                if false {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 float32 = 0.1
                        var v12 int32 = -120
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                } else {
                    v2 = 0.0
                    if (541 > 255) {
                        v13 := GenS1{F0: (-(v5)), F1: true, F2: 65535}
                        v1 = (v1 + 2147483647)
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 4; fi14++ {
                        var v15 uint64 = 2
                        _ = fi14
                        _ = v15
                    }
                }
            } else {
                var v16 uint64 = 18446744073709551615
                v2 -= v2
                switch "foo bar" {
                case "hello":
                    v5 = (v5 << ((37) & 63))
                case "a":
                    v5 = -128
                }
                _ = v16
            }
        case 4294967295:
            v5 -= (v5 + int8(v4))
        }
    }
    v4 += 1
    if len(v6) > 0 {
        v6[0] = v1
    }
    fmt.Println(v2, v1)
    var v17 int16 = 1
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

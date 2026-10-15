package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 int }
func genFn0() bool {
    if (1 < 401) {
        switch 32767 {
        case 32768:
            var v7 uint8 = 255
            _ = v7
        default:
            var v8 string = ("hello" + "hello")
            _ = v8
        }
        switch 127 {
        case -63:
            var v9 uint64 = 557
            _ = v9
        }
        var v10 bool = false
        _ = v10
    } else {
        v11 := GenS1{F0: 32767, F1: -9223372036854775808, F2: 396}
        _ = v11
    }
    var v12 int64 = 0
    for fi13 := 0; fi13 < 10; fi13++ {
        switch "" {
        case "a":
            v12 = (v12 >> ((42) & 63))
        default:
            for fi14 := 0; fi14 < 8; fi14++ {
                switch fi13 {
                case 4294967295:
                    v12 = v12
                }
                _ = fi14
            }
        }
        _ = fi13
    }
    _ = v12
    return (false || false)
}
func genFn1(p0 bool) int64 {
    for fi15 := 0; fi15 < 8; fi15++ {
        p0 = p0
        if (9223372036854775807 <= 42) {
            var v16 uint8 = 0
            _ = v16
        }
        _ = fi15
    }
    v17 := []int{}
    _ = p0
    _ = v17
    return 0
}

func main() {
    v1 := []string{"foo bar", ("\n\t" + "hello"), "\n\t", "a"}
    v2 := make(map[uint32]bool)
    var v3 string = "foo bar"
    var v4 float64 = (3.14159 / float64((((1e10 / -2.25) / 1.0) - float32(0.0))))
    var v5 uint32 = 2147483648
    var v6 int16 = -28
    var v18 string = v3
    if (30 <= 582) {
        _ = len(v2)
    }
    fmt.Println(v6, v6)
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int64 = int64(v5)
        v21 := GenS2{F0: (((v18 + v18) + v3) + "hello"), F1: -1, F2: fi19}
        fmt.Println(v3, v6)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 string = "hello"
    delete(v2, (^(v5)))
    var v23 float32 = float32(v4)
    switch -220 {
    case -874:
        fmt.Println(v5, v18, v18)
    case 9223372036854775807:
        var v24 int = 2147483647
        _ = v24
    case -1:
        v25 := GenS1{F0: 32767, F1: 719, F2: v5}
        _ = v25
    default:
        switch -206 {
        case -1:
            v26 := v2[(v5 >> ((42) & 63))]
            _ = v26
        default:
            v18 = ("" + v3)
        }
    }
    var v27 float32 = v23
    var v28 uint8 = 255
    fmt.Println(v4, v5, v6)
    for fi29 := 0; fi29 < 18; fi29++ {
        v2[(v5 >> ((255) & 63))] = ((v18 + v18) < v18)
        var v30 int16 = v6
        if false {
            v28 = (v28 % ((v28) | 1))
            v22 = v18
        } else {
            for fi31 := 0; fi31 < 1; fi31++ {
                switch (fi29 & (^(fi31))) {
                case 9223372036854775807:
                    var v32 uint16 = 65535
                    _ = v32
                }
                _ = fi31
            }
            v4 += (v4 - 0.0)
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

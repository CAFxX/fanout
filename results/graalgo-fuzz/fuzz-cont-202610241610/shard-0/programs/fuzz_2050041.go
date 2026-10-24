package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
func genFn0(p0 uint16, p1 int64) uint32 {
    var v6 string = ("\n\t" + "")
    fmt.Println(p0, p1, p1)
    p1 = -9223372036854775808
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            if (2 == 547) {
                var v8 uint32 = 1
                v9 := []int{(fi7 * (^(fi7)))}
                _ = v8
                _ = v9
            }
        } else {
            switch fi7 {
            case 59:
                if ((0.1 * 1.5) > 0.0) {
                    fmt.Println(v6, p0)
                    p1 = p1
                }
            default:
                for fi10 := 0; fi10 < 2; fi10++ {
                    p0 = 405
                    _ = fi10
                }
            }
            var v11 float64 = (0.1 * 1e10)
            _ = v11
        }
        _ = fi7
    }
    v12 := GenS1{F0: 255, F1: 42, F2: 2147483647}
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return 2
}
func genFn1() string {
    var v13 uint = 255
    fmt.Println(v13)
    switch 482 {
    case -457:
        fmt.Println(v13)
    }
    _ = v13
    return "foo bar"
}

func main() {
    v1 := []string{"hello", "foo bar"}
    var v2 int16 = -235
    v3 := []string{"hello", "hello", "hello"}
    var v4 uint = 245
    var v5 uint16 = 255
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint64 = 4294967295
        _ = fi14
        _ = v15
    }
    var v16 float64 = float64((1.5 * 0.0))
    var v17 uint64 = 0
    v4 -= (v4 >> ((v4) & 63))
    v18 := GenS1{F0: 128, F1: -712, F2: -2147483648}
    for fi19 := 0; fi19 < 7; fi19++ {
        fmt.Println(v17, v2)
        switch "" {
        case "hello":
            var v20 float64 = 3.14159
            _ = v20
        default:
            _ = len(v1)
        }
        _ = fi19
    }
    var v21 uint16 = (v5 << ((42) & 63))
    var v22 uint8 = 151
    v23 := GenS1{F0: v22, F1: 170, F2: 200}
    v2 = (v2 >> (((v4 + 4294967295)) & 63))
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int64 = 42
        v22 = (v22 + v22)
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}

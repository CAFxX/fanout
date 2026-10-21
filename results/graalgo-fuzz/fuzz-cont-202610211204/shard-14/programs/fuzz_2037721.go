package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint8 }
func genFn0(p0 int, p1 uint8) bool {
    v5 := GenS1{F0: p1, F1: 0}
    var v6 uint8 = 181
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (1 > 1)
}
func genFn1(p0 int32, p1 uint32) uint {
    switch -248 {
    case 4294967295:
        var v7 float64 = (0.1 * 1.5)
        _ = v7
    }
    switch 127 {
    case 16:
        v8 := GenS1{F0: 245, F1: 1}
        _ = v8
    case 1:
        switch "" {
        case "hello":
            var v9 int16 = -806
            _ = v9
        default:
            v10 := GenS1{F0: 42, F1: 0}
            _ = v10
        }
    case 42:
        var v11 int32 = p0
        _ = v11
    default:
        for fi12 := 0; fi12 < 14; fi12++ {
            if false {
                p0 = 42
            }
            var v13 string = ("foo bar" + ((("\n\t" + "hello") + "") + "hello"))
            for fi14 := 0; fi14 < 18; fi14++ {
                if ((-100 != -14) || false) {
                    p0 = p0
                } else {
                    v15 := []int{(fi14 << ((4294967295) & 63)), fi12}
                    v16 := GenS1{F0: 26, F1: -1}
                    _ = v15
                    _ = v16
                }
                v17 := []int{fi12, 42, 32767}
                _ = fi14
                _ = v17
            }
            _ = fi12
            _ = v13
        }
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(p1, p1, p0)
        var v19 uint16 = 658
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    return 490
}

func main() {
    var v1 uint8 = 99
    v2 := make(map[string]int)
    var v3 float32 = float32(-2.25)
    var v4 bool = (-957 == -1)
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 bool = v4
        _ = fi20
        _ = v21
    }
    delete(v2, "")
    switch ((((("a" + "a") + ("foo bar" + "a")) + "a") + ("foo bar" + "hello")) + (("" + "a") + (("" + "a") + "\n\t"))) {
    case "hello":
        v22 := GenS2{F0: -1, F1: v4, F2: (^(v1))}
        _ = v22
    default:
        v23 := GenS2{F0: -58, F1: v4, F2: v1}
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 uint16 = 42
            _ = fi25
            _ = v26
        }
        var v27 int8 = 42
        _ = fi24
        _ = v27
    }
    v4 = (("a" + ("\n\t" + "\n\t")) != "hello")
    var v28 uint32 = 4294967295
    var v29 int32 = 392
    fmt.Println(v3)
    for fi30 := 0; fi30 < 7; fi30++ {
        v3 = (v3 * v3)
        _ = fi30
    }
    v31 := GenS2{F0: 0, F1: true, F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}

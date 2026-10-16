package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0(p0 uint16, p1 int8, p2 uint8) int64 {
    p0 = 1
    var v5 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 42
    var v3 uint = 255
    var v4 uint16 = 0
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float64 = 3.14159
        if false {
            var v8 int = fi6
            switch fi6 {
            case -626:
                v3 = v3
            case 32768:
                if (true || true) {
                    var v9 uint32 = 835
                    v9 = v9
                    _ = v9
                } else {
                    v4 = (^(v4))
                    var v10 float64 = float64(0.0)
                    fmt.Println(v3)
                    _ = v10
                }
            default:
                var v11 float64 = (v7 * -2.25)
                _ = v11
            }
            _ = v8
        } else {
            fmt.Println(v2, v7)
            v7 -= v7
        }
        _ = fi6
        _ = v7
    }
    v12 := GenS1{F0: uint32(v4), F1: 32767, F2: 255}
    var v13 uint64 = 2
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 string = (("\n\t" + ("foo bar" + ("foo bar" + (("\n\t" + "") + (("hello" + "a") + ("\n\t" + ("foo bar" + ("\n\t" + "")))))))) + ((("a" + "a") + ("foo bar" + "foo bar")) + "a"))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 4294967295, F1: v4, F2: (^(v4))}
    v17 := []int{0, 1, -2147483648, 127}
    v4 += v4
    v18 := v1[2147483648]
    var v19 float32 = 3.14159
    switch "a" {
    case "":
        v4 = v4
    default:
        var v20 uint16 = 42
        _ = v20
    }
    var v21 uint32 = 0
    var v22 uint8 = 42
    delete(v1, v21)
    var v23 uint64 = (^(v13))
    for fi24 := 0; fi24 < 19; fi24++ {
        _ = len(v17)
        var v25 float64 = 0.1
        _ = fi24
        _ = v25
    }
    var v26 uint64 = v2
    var v27 string = ("" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}

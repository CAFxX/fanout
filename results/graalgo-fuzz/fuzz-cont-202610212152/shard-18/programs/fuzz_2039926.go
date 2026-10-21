package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 string, p2 bool) bool {
    p2 = false
    v5 := GenS1{F0: 100, F1: -1, F2: (false && (!(3.14159 >= (((float32(1.5) / 1e10) / (1e10 * (float32(0.1) / 1e10))) + (1.5 * 3.14159)))))}
    p0 -= p0
    v6 := []int{32767, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 4294967295
    v3 := make(map[uint32]string)
    var v4 uint32 = 0
    switch "\n\t" {
    case "foo bar":
        v2 -= (v2 % ((v2) | 1))
    }
    v7 := GenS1{F0: -128, F1: 2147483647, F2: false}
    var v8 float64 = v1
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            switch fi9 {
            case 42:
                v10 := GenS1{F0: 127, F1: 42, F2: true}
                _ = v10
            case 127:
                var v11 int16 = 1
                _ = v11
            case 4294967295:
                fmt.Println(v2, fi9, v4)
            }
            _ = fi9
        }
    }
    if (!(-1 < 1)) {
        delete(v3, (v4 % (((^(v4))) | 1)))
        v4 -= 373
    }
    v12 := GenS2{F0: 128}
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            v2 = 0
            _ = fi14
        }
        v1 += float64(3.14159)
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS2{F0: 127}
            _ = fi15
            _ = v16
        }
        _ = fi13
    }
    if false {
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint = 255
            v19 := []int{(fi17 - (fi17 - (fi17 + fi17))), (fi17 & fi17)}
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS1{F0: -128, F1: 42, F2: true}
                switch ("foo bar" + "a") {
                case "":
                    var v22 int32 = 239
                    _ = v22
                case "\n\t":
                    if (107 >= uint8(v18)) {
                        delete(v3, 65535)
                    }
                case "foo bar":
                    fmt.Println(v8, v18, v1)
                }
                fmt.Println(fi20, v4, fi20)
                _ = fi20
                _ = v21
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        var v23 uint32 = v4
        _ = v23
    } else {
        v24 := GenS1{F0: -37, F1: 842, F2: ("\n\t" >= "")}
        _ = v24
    }
    var v25 int = 2147483647
    if (float32((((float64(0.0) / v1) / v1) + (0.0 + v1))) <= (-2.25 / (-2.25 - float32(((3.14159 - 1.5) * v8))))) {
        var v26 bool = (3.14159 <= ((0.1 + 0.1) / 1.5))
        _ = v26
    } else {
        v27 := v3[305]
        var v28 uint32 = v4
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint16 = 0
            v28 += (v4 + v4)
            _ = fi29
            _ = v30
        }
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}

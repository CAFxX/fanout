package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint) int8 {
    var v5 uint64 = 2
    v5 -= v5
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6, -2147483648}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: 63}
    _ = p0
    _ = v5
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 255
    var v2 bool = false
    var v3 int16 = -698
    v4 := make(map[int]bool)
    switch 775 {
    case 1:
        v2 = v2
    case 602:
        switch ("\n\t" + "foo bar") {
        case "a":
            for fi9 := 0; fi9 < 15; fi9++ {
                v2 = ((2 <= 42) && (true && ((fi9 % ((fi9) | 1)) != (^(fi9)))))
                v10 := []int{}
                _ = fi9
                _ = v10
            }
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                v3 = int16(fi11)
                v3 += (v3 & (v3 % ((v3) | 1)))
                _ = fi11
            }
        }
    }
    switch 128 {
    case 127:
        var v12 int32 = 1
        _ = v12
    case -1:
        v13 := []int{}
        _ = v13
    default:
        var v14 bool = true
        _ = v14
    }
    delete(v4, 128)
    v15 := []int{-1}
    for fi16 := 0; fi16 < 11; fi16++ {
        delete(v4, -2147483648)
        _ = fi16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        fmt.Println(v1)
        v1 = 42
        v3 = (v3 >> ((1) & 63))
        _ = fi17
    }
    switch 0 {
    case 128:
        var v18 string = (("hello" + ("foo bar" + "foo bar")) + "foo bar")
        _ = v18
    case 9223372036854775807:
        if v2 {
            switch 0 {
            case 128:
                var v19 float64 = 1.5
                _ = v19
            case -494:
                var v20 uint8 = (v1 >> ((4294967295) & 63))
                _ = v20
            default:
                v1 = 127
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 bool = false
                _ = fi21
                _ = v22
            }
        } else {
            v3 += v3
            v1 -= v1
            for fi23 := 0; fi23 < 11; fi23++ {
                fmt.Println(v1, fi23)
                if len(v15) > 0 {
                    v15[0] = fi23
                }
                fmt.Println(v1)
                _ = fi23
            }
        }
    default:
        fmt.Println(v1, v3, v2)
    }
    if (((((0.1 / (1.5 - (0.1 - float64(1.5)))) * float64(0.1)) - (0.1 - (1e10 * (-2.25 + 1.5)))) + ((((0.0 - 0.0) + float64(float32(0.0))) + 1.5) / (float64(1e10) + -2.25))) != 1.5) {
        fmt.Println(v1)
    } else {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := GenS1{F0: -1, F1: (^(v1))}
                var v27 uint64 = 867
                v2 = (((-377 <= 1) || true) || (int(v3) == (fi25 + fi25)))
                _ = fi25
                _ = v26
                _ = v27
            }
            v1 = v1
            v28 := GenS2{F0: 1}
            _ = fi24
            _ = v28
        }
    }
    fmt.Println(v2, v3, v2)
    var v29 bool = (!v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}

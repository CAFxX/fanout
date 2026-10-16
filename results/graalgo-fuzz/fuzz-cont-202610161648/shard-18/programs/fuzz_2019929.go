package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0() uint8 {
    v5 := GenS1{F0: 718, F1: 869}
    var v6 int64 = -314
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 string, p1 int8) uint {
    if true {
        p1 = (p1 << ((312) & 63))
    }
    var v7 uint8 = 42
    _ = p0
    _ = p1
    _ = v7
    return 565
}

func main() {
    var v1 int32 = -392
    var v2 string = ""
    var v3 int = 223
    v4 := []uint32{411, 2, 2147483648, 4294967295}
    fmt.Println(v3, v2)
    var v8 int64 = 1
    v8 += (v8 * v8)
    v4 = append(v4, 42)
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 float32 = 1.5
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: (v1 + (-(v1))), F1: 286}
        v4 = append(v4, 391)
        _ = fi9
        _ = v12
    }
    v4 = append(v4, 0)
    v2 = "\n\t"
    v4 = append(v4, 2147483648)
    switch 0 {
    case 32768:
        v8 = (v8 >> ((0) & 63))
    case -879:
        switch "foo bar" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        case "foo bar":
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: 2147483647, F1: 1}
                var v15 uint32 = 181
                _ = fi13
                _ = v14
                _ = v15
            }
        case "":
            if (!(!false)) {
                for fi16 := 0; fi16 < 2; fi16++ {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v8 += v8
                        v4 = append(v4, 353)
                        _ = fi17
                    }
                    _ = fi16
                }
            }
        default:
            if ((998 == 42) || false) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v1, v2)
                for fi18 := 0; fi18 < 5; fi18++ {
                    fmt.Println(v8)
                    _ = fi18
                }
            }
        }
    case 32767:
        v8 = (v8 | v8)
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int16 = 0
        v20 -= v20
        if (v20 >= v20) {
            switch (v2 + v2) {
            case "foo bar":
                for fi21 := 0; fi21 < 8; fi21++ {
                    v20 += v20
                    v1 -= (v1 << ((4294967295) & 63))
                    _ = fi21
                }
            case "\n\t":
                v3 = fi19
            }
            var v22 bool = false
            var v23 float64 = (0.1 / (0.1 * 1e10))
            _ = v22
            _ = v23
        } else {
            v2 = ""
        }
        _ = fi19
        _ = v20
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        switch 4294967295 {
        case 128:
            v1 -= 0
        case 42:
            var v25 bool = false
            _ = v25
        case 32768:
            var v26 int = 4294967295
            _ = v26
        default:
            v3 += (v3 & (-(fi24)))
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) int8 {
    p1 = p1
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = (3.14159 * (3.14159 + float32(((1.5 - float64(1.5)) / (1e10 * 1.5)))))
        _ = fi7
        _ = v8
    }
    p1 += (^(p1))
    if ((0 <= 32767) || false) {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int = fi9
            v11 := GenS1{F0: 130, F1: 16, F2: 9223372036854775807}
            p0 = p0
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        switch 1 {
        case 9223372036854775807:
            v12 := []int{}
            _ = v12
        case -1:
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p1)
            case "a":
                v13 := GenS2{F0: 42, F1: (("a" + (("" + (("" + ("foo bar" + "a")) + "foo bar")) + "\n\t")) + ("a" + "foo bar"))}
                _ = v13
            default:
                if true {
                    fmt.Println(p0)
                    v14 := []int{128}
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 float64 = float64((0.0 * float32((0.1 * 0.0))))
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
            }
        }
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint8 = 165
    var v2 float32 = 1e10
    var v3 float64 = -2.25
    v4 := []int{-94, 127}
    var v5 bool = (v3 > (v3 + v3))
    v6 := make(map[int]uint32)
    v2 = float32(v3)
    v17 := GenS2{F0: 42, F1: "\n\t"}
    v18 := GenS3{F0: ("a" + "a"), F1: (v1 * v1)}
    var v19 int16 = 42
    var v20 uint8 = 0
    var v21 int = -2147483648
    if ((-1 <= 9223372036854775807) || true) {
        if true {
            var v22 uint64 = 1
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                for fi23 := 0; fi23 < 5; fi23++ {
                    fmt.Println(v1)
                    var v24 int8 = int8(v19)
                    _ = fi23
                    _ = v24
                }
            }
            _ = v22
        }
        v20 -= 128
        v20 += 120
    } else {
        v5 = (463 != 65535)
        var v25 string = "foo bar"
        _ = v25
    }
    fmt.Println(v1, v5, v2)
    v3 = v3
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 bool = v5
        _ = len(v4)
        if len(v4) > 0 {
            v4[0] = fi26
        }
        _ = fi26
        _ = v27
    }
    v20 = (v20 << ((1) & 63))
    v28 := GenS4{F0: genFn0(458, 216)}
    fmt.Println(v21)
    if (!((!false) || (0 >= 4294967295))) {
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 int8 = -128
            _ = fi29
            _ = v30
        }
        switch v21 {
        case 39:
            v31 := GenS2{F0: 0, F1: ("hello" + "\n\t")}
            _ = v31
        case 42:
            if (v5 && ((v19 & (v19 - v19)) >= 32767)) {
                var v32 uint32 = 42
                v33 := GenS1{F0: 116, F1: uint16(v20), F2: (v21 % (((v21 * 290)) | 1))}
                v34 := GenS1{F0: 292, F1: 1, F2: v21}
                _ = v32
                _ = v33
                _ = v34
            }
        default:
            v4 = append(v4, 127)
        }
        var v35 uint32 = uint32(v20)
        _ = v35
    } else {
        var v36 int = v21
        v21 -= (v21 * (-(v21)))
        _ = v36
    }
    v2 += -2.25
    v19 = 32767
    v37 := v6[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

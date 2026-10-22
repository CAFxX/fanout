package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }
type GenS2 struct { F0 uint32; F1 bool; F2 uint }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 uint8; F2 int16 }
type GenS8 struct { F0 int8; F1 int8 }
func genFn0(p0 bool, p1 uint8, p2 uint64) string {
    var v6 int16 = 32767
    fmt.Println(p1, p1, p2)
    v7 := GenS1{F0: 106, F1: (933 >= uint32(p2)), F2: ("hello" + "a")}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{"\n\t"}
    v3 := make(map[string]uint32)
    var v4 int16 = 1
    v5 := []string{("foo bar" + ""), "foo bar", ""}
    v9 := GenS2{F0: 156, F1: (184 >= 0), F2: 841}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v10 uint8 = 255
    v10 = v10
    var v11 string = ("" + "\n\t")
    v12 := GenS1{F0: 55, F1: true, F2: "hello"}
    v13 := GenS3{F0: 1, F1: (v10 ^ v10)}
    fmt.Println(v11)
    v14 := GenS2{F0: 711, F1: false, F2: 387}
    if true {
        var v15 int = -1
        v16 := GenS2{F0: 456, F1: ((float64(float32(-2.25)) + 0.1) < ((1e10 - ((1.5 / 0.1) / 0.1)) - 1.5)), F2: 2}
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                switch 79 {
                case 9223372036854775807:
                    v4 += (v4 % ((1000) | 1))
                case 127:
                    v4 = (v4 / (((-(v4))) | 1))
                    fmt.Println(v10, fi17)
                    fmt.Println(fi18)
                case -2147483648:
                    v19 := GenS4{F0: 0, F1: 1}
                    _ = v19
                default:
                    v20 := GenS2{F0: 0, F1: (((fi18 & -66) != (fi18 << ((4294967295) & 63))) || (v1 <= (((v1 / float32((0.0 / (float64(v1) - (1e10 + ((1.5 + (float64(float32((float64(v1) + (0.1 - 1e10)))) / (1e10 + float64(v1)))) * (float64(float32(float64(v1))) - 1e10))))))) - ((((v1 - v1) * v1) - (((v1 + 3.14159) + float32(-2.25)) / (1.5 * (v1 / v1)))) * v1)) - v1))), F2: 4294967295}
                    _ = v20
                }
                v4 += 1000
                _ = fi18
            }
            _ = fi17
        }
        v1 += 1e10
    }
    if (0 > -1) {
        v21 := GenS3{F0: -1, F1: (v10 / ((1) | 1))}
        var v22 int = 42
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS5{F0: 42, F1: 42}
            v4 = (v4 >> ((684) & 63))
            if true {
                fmt.Println(v4)
            } else {
                if (1 >= 18446744073709551615) {
                    v25 := GenS1{F0: 865, F1: false, F2: v11}
                    _ = v25
                }
                var v26 int = 0
                for fi27 := 0; fi27 < 5; fi27++ {
                    v28 := GenS6{F0: 1}
                    delete(v3, v11)
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = v21
        _ = v22
    }
    v29 := GenS7{F0: 42, F1: v10, F2: (v4 | v4)}
    var v30 uint64 = 748
    delete(v3, "a")
    fmt.Println(v30)
    switch "" {
    case "":
        if ((v10 > 2) || (4294967295 >= 2)) {
            var v31 uint32 = 65535
            for fi32 := 0; fi32 < 8; fi32++ {
                fmt.Println(v10)
                var v33 uint8 = (v10 / ((225) | 1))
                _ = fi32
                _ = v33
            }
            var v34 float32 = (1e10 * v1)
            _ = v31
            _ = v34
        }
    case "foo bar":
        v4 -= v4
    case "\n\t":
        v35 := GenS8{F0: 127, F1: 1}
        _ = v35
    }
    v3["foo bar"] = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}

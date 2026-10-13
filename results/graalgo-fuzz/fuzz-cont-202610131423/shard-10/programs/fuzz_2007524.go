package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 uint32 }
func genFn0() uint8 {
    var v5 float32 = 0.1
    switch 972 {
    case 32767:
        v5 -= 1e10
    case 815:
        v5 += 0.0
    default:
        v6 := GenS1{F0: ("\n\t" + ("a" + "hello"))}
        _ = v6
    }
    _ = v5
    return 128
}
func genFn1(p0 int32, p1 int8) uint32 {
    fmt.Println(p0, p0)
    p1 = 42
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 156
    v2 := make(map[uint32]string)
    var v3 uint8 = v1
    var v4 uint = 42
    var v7 int = 32768
    if true {
        var v8 string = "hello"
        v9 := GenS2{F0: 1, F1: false, F2: (v7 % ((v7) | 1))}
        var v10 int8 = 1
        _ = v10
        _ = v8
        _ = v9
    }
    switch "" {
    case "hello":
        if true {
            var v11 int16 = 32767
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := GenS2{F0: 696, F1: true, F2: (fi12 | fi12)}
                var v14 uint32 = 4294967295
                fmt.Println(v7)
                _ = fi12
                _ = v13
                _ = v14
            }
            if (2147483647 != 2147483647) {
                var v15 float64 = ((1e10 - (-2.25 * (1.5 - 0.1))) - (1e10 + (((((float64((0.1 * -2.25)) * (0.1 + -2.25)) / 0.1) / -2.25) - ((-2.25 - ((0.1 + float64(1.5)) - 0.0)) / 3.14159)) * -2.25)))
                var v16 float32 = 1.5
                _ = v15
                _ = v16
            } else {
                v1 = v1
                var v17 int = 9223372036854775807
                _ = v17
            }
            _ = v11
        } else {
            var v18 uint = (v4 % (((v4 * v4)) | 1))
            var v19 int32 = 0
            _ = v18
            _ = v19
        }
    case "\n\t":
        var v20 uint8 = genFn0()
        _ = v20
    }
    fmt.Println(v4)
    v7 = v7
    v21 := GenS1{F0: ""}
    if (!false) {
        var v22 float64 = 1e10
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v4 -= 0
        _ = fi23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        if (false && true) {
            var v25 float32 = 3.14159
            _ = v25
        } else {
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 int32 = 564
                _ = fi26
                _ = v27
            }
            var v28 int64 = int64(v4)
            _ = v28
        }
        _ = fi24
    }
    v29 := v2[2147483648]
    v30 := GenS1{F0: ("foo bar" + "")}
    if false {
        var v31 float64 = float64(0.0)
        for fi32 := 0; fi32 < 11; fi32++ {
            for fi33 := 0; fi33 < 10; fi33++ {
                fmt.Println(v3)
                v2[4294967295] = v29
                v2[0] = "foo bar"
                _ = fi33
            }
            for fi34 := 0; fi34 < 17; fi34++ {
                fmt.Println(v4, v31, fi32)
                v35 := []int{}
                switch fi32 {
                case -796:
                    if true {
                        v36 := GenS1{F0: v29}
                        var v37 bool = false
                        _ = v36
                        _ = v37
                    }
                case 564:
                    v38 := v2[genFn1(581, -128)]
                    _ = v38
                default:
                    v39 := GenS1{F0: (((("a" + "hello") + (v29 + (v29 + v29))) + (v29 + "")) + ((v29 + v29) + v29))}
                    _ = v39
                }
                _ = fi34
                _ = v35
            }
            _ = fi32
        }
        v40 := GenS3{F0: 126}
        _ = v31
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}

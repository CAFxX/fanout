package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 string; F2 uint8 }
type GenS5 struct { F0 int32; F1 int64 }
func genFn0(p0 uint32, p1 uint32, p2 int) uint16 {
    if true {
        p2 += (p2 / (((p2 - 32768)) | 1))
        var v5 int = (p2 ^ 4294967295)
        _ = v5
    }
    var v6 int64 = 0
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 int = -2147483648
    v2 := []uint32{2, 65535, 2147483647}
    v3 := make(map[uint32]uint32)
    var v4 uint = 2
    v4 -= (^(v4))
    for fi7 := 0; fi7 < 3; fi7++ {
        if true {
            var v8 uint32 = 4294967295
            _ = v8
        } else {
            var v9 int32 = -1
            var v10 int32 = v9
            _ = v10
            _ = v9
        }
        var v11 float32 = (((1.5 / ((float32(3.14159) / -2.25) * 3.14159)) + (1.5 / (1.5 - (1e10 + 1e10)))) + float32((3.14159 + 3.14159)))
        if false {
            var v12 string = "a"
            _ = v12
        }
        _ = fi7
        _ = v11
    }
    switch 4294967295 {
    case 128:
        for fi13 := 0; fi13 < 15; fi13++ {
            v3[4294967295] = 2
            v14 := GenS1{F0: (1 >= 277)}
            _ = fi13
            _ = v14
        }
    case 1:
        var v15 uint8 = 2
        _ = v15
    }
    var v16 uint64 = 1
    v1 -= -2147483648
    v17 := GenS1{F0: ("foo bar" <= "a")}
    var v18 string = ("\n\t" + ("" + "foo bar"))
    v19 := GenS2{F0: 9223372036854775807}
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint = 4294967295
        _ = fi20
        _ = v21
    }
    if (!(0 != 127)) {
        if (-15 < 0) {
            if ((true && true) || (42 > -23)) {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int = fi22
                    v24 := GenS3{F0: 2147483647}
                    v25 := GenS4{F0: v16, F1: v18, F2: 141}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            } else {
                switch (v18 + v18) {
                case "foo bar":
                    v18 = v18
                default:
                    if ((!(false && (-1 != 32767))) && false) {
                        fmt.Println(v4, v18, v16)
                        fmt.Println(v18, v4)
                        v26 := GenS3{F0: 2147483647}
                        _ = v26
                    }
                }
                v27 := GenS5{F0: -946, F1: -9223372036854775808}
                v3[2147483648] = 906
                _ = v27
            }
        }
        v28 := GenS2{F0: 1}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

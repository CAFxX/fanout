package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint32, p1 int, p2 int64) int {
    v6 := GenS1{F0: ("a" > ("a" + ("hello" + "hello"))), F1: -128, F2: 1}
    p1 -= (p1 >> ((283) & 63))
    for fi7 := 0; fi7 < 15; fi7++ {
        if ((p0 / (((p0 | p0)) | 1)) < p0) {
            p2 += p2
            switch (fi7 / (((fi7 - fi7)) | 1)) {
            case 405:
                for fi8 := 0; fi8 < 18; fi8++ {
                    fmt.Println(fi8, p0, p2)
                    _ = fi8
                }
            case 9223372036854775807:
                var v9 float32 = 0.0
                _ = v9
            default:
                switch (p1 + (^(fi7))) {
                case -9223372036854775808:
                    p0 = (p0 / ((4294967295) | 1))
                case -412:
                    v10 := GenS1{F0: true, F1: -121, F2: 2}
                    _ = v10
                case 128:
                    p0 = 2147483647
                default:
                    fmt.Println(p2, p0)
                }
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1(p0 int8, p1 int32, p2 int64) bool {
    p1 += (p1 / ((p1) | 1))
    var v11 float32 = (float32((0.0 + 1.5)) - 1.5)
    var v12 int64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return false
}

func main() {
    var v1 string = ("" + "foo bar")
    v2 := make(map[uint32]int)
    var v3 uint32 = 42
    var v4 int = 32768
    v5 := []uint32{v3, v3, (v3 / ((v3) | 1))}
    var v13 int16 = -32768
    var v14 uint32 = 913
    switch (((("" + "\n\t") + v1) + v1) + v1) {
    case "\n\t":
        if true {
            var v15 uint16 = 208
            v16 := v2[(v3 ^ (v3 >> ((0) & 63)))]
            _ = v15
            _ = v16
        } else {
            if len(v5) > 0 {
                v5[0] = (v14 + v3)
            }
        }
    case "":
        v17 := v2[(v3 - v14)]
        _ = v17
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v18 int16 = v13
    var v19 uint16 = 255
    var v20 uint16 = 713
    var v21 int64 = -703
    delete(v2, v14)
    v22 := v2[v14]
    v23 := GenS2{F0: 1, F1: (-(v18)), F2: (v21 ^ v21)}
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 bool = false
        v26 := GenS3{F0: ((^(v21)) == 1), F1: 2}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v19, v3, v3)
    fmt.Println(v3, v20)
    v13 = (^(v13))
    var v27 uint16 = (v20 << ((42) & 63))
    var v28 bool = (false || false)
    fmt.Println(v28, v27)
    v29 := v2[(v14 / ((v3) | 1))]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

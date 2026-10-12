package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }

func main() {
    var v1 string = ("\n\t" + "hello")
    v2 := []string{("a" + "foo bar"), v1, ((((((((("" + v1) + v1) + ((v1 + "") + "")) + v1) + v1) + v1) + ((v1 + v1) + v1)) + v1) + v1), (v1 + (((v1 + (v1 + "hello")) + "\n\t") + (("" + v1) + v1)))}
    v3 := []uint32{4294967295}
    v4 := GenS1{F0: 401, F1: 42, F2: 946}
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int8 = -3
        if len(v3) > 0 {
            v3[0] = 856
        }
        v1 = v1
        _ = fi5
        _ = v6
    }
    if false {
        v1 = "\n\t"
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v1 = ""
        _ = fi7
    }
    var v8 uint = 18446744073709551615
    if (2 < 51) {
        if ((0 != -1) && true) {
            fmt.Println(v1, v8)
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(v8, fi9, v8)
                _ = fi9
            }
        } else {
            if (-1 < -804) {
                var v10 float64 = 3.14159
                var v11 uint8 = 128
                v3 = append(v3, uint32(v8))
                _ = v10
                _ = v11
            } else {
                v2 = append(v2, v1)
                v1 = (v1 + "hello")
                var v12 int64 = 9223372036854775807
                _ = v12
            }
        }
    } else {
        v8 += v8
        var v13 float32 = 0.0
        _ = v13
    }
    _ = len(v3)
    if (((255 == 1) && (true && (42 < 4294967295))) || true) {
        _ = len(v2)
        v14 := GenS1{F0: v8, F1: v8, F2: 505}
        v15 := GenS1{F0: v8, F1: 1, F2: 42}
        _ = v14
        _ = v15
    } else {
        var v16 int16 = 1000
        v8 -= v8
        v17 := GenS1{F0: (v8 << ((4294967295) & 63)), F1: (v8 ^ v8), F2: 0}
        _ = v16
        _ = v17
    }
    if false {
        var v18 float32 = -2.25
        _ = len(v2)
        _ = v18
    }
    v8 = (v8 >> (((v8 / ((18446744073709551615) | 1))) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v1)
    v8 = 1
    v1 = (v1 + ((v1 + v1) + v1))
    var v19 int64 = -53
    _ = len(v2)
    fmt.Println(v1, v19)
    for fi20 := 0; fi20 < 12; fi20++ {
        v1 = (("a" + "\n\t") + v1)
        for fi21 := 0; fi21 < 6; fi21++ {
            if (!(int16(v19) == -741)) {
                v22 := GenS1{F0: 18446744073709551615, F1: (v8 ^ v8), F2: 108}
                switch ((v1 + v1) + v1) {
                case "hello":
                    var v23 uint = v8
                    _ = v23
                case "":
                    v19 += (v19 | 322)
                    v8 = v8
                    var v24 float32 = 0.1
                    _ = v24
                case "a":
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                default:
                    var v25 int32 = 269
                    _ = v25
                }
                var v26 float32 = 0.1
                _ = v22
                _ = v26
            } else {
                if len(v3) > 0 {
                    v3[0] = 0
                }
            }
            _ = fi21
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

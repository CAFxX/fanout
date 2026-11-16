package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
func genFn0() uint8 {
    if true {
        v7 := []int{-350}
        v7 = append(v7, 42)
        var v8 float32 = 1e10
        _ = v7
        _ = v8
    }
    var v9 uint32 = 4294967295
    var v10 int64 = 42
    _ = v10
    _ = v9
    return 2
}
func genFn1(p0 int, p1 uint8, p2 uint) int64 {
    var v11 string = ""
    var v12 int8 = 1
    fmt.Println(v12, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -203
}

func main() {
    var v1 int8 = -128
    var v2 bool = true
    var v3 uint16 = 82
    var v4 bool = (!(!(!v2)))
    var v5 uint16 = (v3 + 1)
    var v6 int = 127
    var v13 uint8 = 42
    var v14 int = (v6 ^ v6)
    v15 := []int{9223372036854775807}
    if (0.0 != (3.14159 * (-2.25 + 0.1))) {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int64 = int64(v1)
            _ = fi16
            _ = v17
        }
        if len(v15) > 0 {
            v15[0] = v14
        }
        var v18 float64 = (1e10 - ((float64(float32(0.0)) - (-2.25 / float64(3.14159))) - -2.25))
        _ = v18
    }
    v1 = (v1 ^ (v1 * (v1 % ((v1) | 1))))
    if v4 {
        v5 = 502
        var v19 int16 = 603
        _ = v19
    } else {
        if (1 <= 4294967295) {
            v20 := GenS1{F0: 2147483647}
            if (v13 == (v13 / (((v13 * (v13 / (((v13 + v13)) | 1)))) | 1))) {
                v15 = append(v15, v6)
                var v21 bool = v2
                _ = v21
            } else {
                var v22 float32 = 0.1
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                for fi24 := 0; fi24 < 19; fi24++ {
                    fmt.Println(v14, v5, v4)
                    fmt.Println(v3, v13)
                    _ = fi24
                }
                switch ("hello" + "") {
                case "\n\t":
                    var v25 int32 = -583
                    _ = v25
                case "a":
                    v2 = (v5 < v3)
                case "foo bar":
                    switch "" {
                    case "hello":
                        fmt.Println(fi23, v2)
                    }
                default:
                    var v26 int = (v14 - 32768)
                    _ = v26
                }
                _ = fi23
            }
            _ = v20
        }
        fmt.Println(v13, v3, v13)
        var v27 int64 = int64(v14)
        _ = v27
    }
    var v28 uint8 = (v13 << ((4294967295) & 63))
    var v29 int16 = 32767
    fmt.Println(v6, v4, v28)
    switch (v14 >> ((0) & 63)) {
    case 32768:
        switch (v14 * 822) {
        case 32768:
            v28 += v28
        }
    }
    v30 := GenS1{F0: 837}
    if len(v15) > 0 {
        v15[0] = (^(v14))
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        fmt.Println(v2, fi31, v13)
        v29 = v29
        _ = fi31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(v13, fi32)
        _ = fi32
    }
    v29 = (v29 | v29)
    v33 := GenS2{F0: 107, F1: 2147483647, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

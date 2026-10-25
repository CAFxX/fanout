package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint8 }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    v7 := []int{-1, 42, 1, 4294967295}
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint = 42
    v2 := []uint32{271, 2, 42}
    v3 := []int{127, 32768, -47}
    var v4 uint8 = 223
    var v5 bool = (true || (!(0 == 42)))
    var v6 string = "hello"
    var v8 int64 = -1
    switch (v6 + v6) {
    case "foo bar":
        var v9 float32 = (0.0 * -2.25)
        _ = v9
    case "":
        var v10 uint = (v1 + (v1 | (v1 >> ((genFn0(uint32(v4))) & 63))))
        _ = v10
    case "hello":
        var v11 string = "\n\t"
        _ = v11
    }
    var v12 int = 127
    var v13 int64 = (-(v8))
    v14 := GenS1{F0: 1000, F1: 255, F2: 1}
    if v5 {
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!v5) {
                v16 := GenS1{F0: int16(v4), F1: (v4 | (v4 % (((v4 << ((v1) & 63))) | 1))), F2: 1}
                if v5 {
                    fmt.Println(v6, v6)
                    fmt.Println(v8, v6, v1)
                    fmt.Println(v6, v4)
                } else {
                    if len(v3) > 0 {
                        v3[0] = (^(fi15))
                    }
                    v17 := GenS1{F0: -253, F1: v4, F2: -77}
                    _ = v17
                }
                _ = v16
            } else {
                v18 := GenS1{F0: 42, F1: v4, F2: 42}
                _ = v18
            }
            if v5 {
                var v19 uint8 = 255
                switch "" {
                case "foo bar":
                    v5 = true
                case "a":
                    v20 := GenS1{F0: 810, F1: 1, F2: -128}
                    _ = v20
                default:
                    var v21 int = v12
                    _ = v21
                }
                v22 := GenS1{F0: -1, F1: (v19 ^ 247), F2: -128}
                _ = v19
                _ = v22
            }
            switch (v6 + (v6 + "foo bar")) {
            case "\n\t":
                v6 = v6
            case "a":
                if (!v5) {
                    if (!false) {
                        _ = len(v3)
                        fmt.Println(v4, fi15, fi15)
                        v12 -= fi15
                    } else {
                        v5 = false
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if (int(v4) < (fi15 << (((^(v1))) & 63))) {
                        var v23 int32 = 608
                        v24 := GenS2{F0: 899, F1: 1, F2: 1}
                        v5 = ((!(1.5 > 0.1)) && v5)
                        _ = v23
                        _ = v24
                    }
                } else {
                    fmt.Println(v5, v1)
                }
            default:
                for fi25 := 0; fi25 < 8; fi25++ {
                    v2 = append(v2, 1)
                    v26 := GenS1{F0: 32767, F1: 1, F2: 1}
                    switch -47 {
                    case -2147483648:
                        fmt.Println(fi15, v5, fi25)
                    case 127:
                        var v27 string = "\n\t"
                        _ = v27
                    case -1:
                        switch (fi25 + fi15) {
                        case -9223372036854775808:
                            var v28 bool = true
                            _ = v28
                        case 9223372036854775807:
                            v29 := GenS2{F0: 0, F1: 1, F2: 270}
                            _ = v29
                        default:
                            var v30 uint64 = 4294967295
                            _ = v30
                        }
                    }
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi15
        }
        v31 := GenS2{F0: 0, F1: 814, F2: 82}
        _ = v31
    } else {
        if v5 {
            v4 = (^(v4))
            v32 := GenS3{F0: v4, F1: 0, F2: v4}
            _ = v32
        }
        var v33 uint32 = 4294967295
        switch (v12 / (((^(v12))) | 1)) {
        case -629:
            var v34 uint = (v1 / (((v1 - (v1 | v1))) | 1))
            _ = v34
        }
        _ = v33
    }
    v35 := GenS3{F0: 252, F1: 2, F2: 128}
    if len(v3) > 0 {
        v3[0] = v12
    }
    var v36 int32 = 0
    v1 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

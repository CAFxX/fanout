package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 uint8; F2 uint }
type GenS7 struct { F0 uint32; F1 int }
func genFn0() string {
    var v5 int = 9223372036854775807
    v6 := GenS1{F0: 4294967295}
    var v7 uint32 = 4294967295
    _ = v5
    _ = v6
    _ = v7
    return (("a" + "hello") + "hello")
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{4294967295, 42, 2, 939}
    var v8 uint64 = v3
    var v9 int32 = -2147483648
    v3 = v8
    v10 := GenS2{F0: 2147483647, F1: (1e10 != 0.1)}
    for fi11 := 0; fi11 < 11; fi11++ {
        v8 -= (v3 % (((v3 + v3)) | 1))
        _ = fi11
    }
    var v12 uint64 = (v8 >> ((0) & 63))
    switch 32767 {
    case 42:
        for fi13 := 0; fi13 < 4; fi13++ {
            if ((1.5 >= -2.25) && false) {
                var v14 uint64 = (v12 << ((2) & 63))
                v14 += v12
                var v15 int64 = -9223372036854775808
                _ = v14
                _ = v15
            } else {
                _ = len(v1)
                if true {
                    v16 := v1[2]
                    v17 := GenS3{F0: 42, F1: ("\n\t" + ("" + ("hello" + ("foo bar" + "hello")))), F2: uint(v12)}
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS3{F0: 0, F1: "hello", F2: 4294967295}
                    v19 := GenS4{F0: 0, F1: 42, F2: ""}
                    _ = v18
                    _ = v19
                }
                var v20 int8 = int8(v9)
                _ = v20
            }
            if ("\n\t" != "a") {
                switch fi13 {
                case -1:
                    v9 += (-(v9))
                default:
                    var v21 float64 = (3.14159 - ((0.0 + 1e10) - 0.1))
                    _ = v21
                }
            } else {
                v22 := GenS5{F0: 2147483647}
                var v23 uint16 = 870
                for fi24 := 0; fi24 < 1; fi24++ {
                    v8 -= (v8 & (v8 + 0))
                    var v25 bool = v2
                    var v26 int8 = 127
                    v8 += (v8 << ((4294967295) & 63))
                    v8 = v12
                    switch ("hello" + "hello") {
                    case "":
                        v3 = (^(v8))
                    case "a":
                        if (0 < 1) {
                            var v27 int64 = -689
                            var v28 uint16 = (v23 ^ (v23 | (v23 & v23)))
                            var v29 uint16 = 617
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    default:
                        fmt.Println(v23, fi24)
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = v22
                _ = v23
            }
            _ = fi13
        }
    case 127:
        if ((v9 % (((v9 & v9)) | 1)) <= v9) {
            var v30 int = 693
            _ = len(v4)
            if (2 == 201) {
                var v31 uint8 = 2
                v32 := GenS6{F0: int(v12), F1: v31, F2: 0}
                _ = v31
                _ = v32
            }
            _ = v30
        } else {
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := v1[2147483647]
                _ = fi33
                _ = v34
            }
            v2 = (!v2)
            var v35 int = 0
            _ = v35
        }
    default:
        var v36 uint64 = v8
        _ = v36
    }
    switch -9223372036854775808 {
    case 127:
        v37 := GenS7{F0: 2, F1: 336}
        _ = v37
    case 274:
        v4 = append(v4, 65535)
    default:
        v3 = 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

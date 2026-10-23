package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint64 }
func genFn0() uint8 {
    switch "a" {
    case "hello":
        var v4 int8 = 100
        _ = v4
    default:
        var v5 int32 = -1
        _ = v5
    }
    var v6 uint32 = 226
    v6 += (v6 % (((v6 + v6)) | 1))
    _ = v6
    return 151
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = -1
    var v3 bool = true
    var v7 uint32 = 2147483647
    var v8 int32 = 0
    v9 := GenS1{F0: 42}
    v1 += (v1 << ((2) & 63))
    v10 := []int{}
    v11 := GenS1{F0: 17}
    if (!(0 < 0)) {
        for fi12 := 0; fi12 < 8; fi12++ {
            if v3 {
                var v13 string = ("a" + ("" + "hello"))
                v2 = (v2 & v2)
                _ = v13
            } else {
                var v14 uint16 = 1
                _ = v14
            }
            v1 -= (v1 - (v1 - v1))
            _ = fi12
        }
        if v3 {
            var v15 int64 = (v1 | v1)
            v7 += v7
            fmt.Println(v2, v8)
            _ = v15
        }
    } else {
        _ = len(v10)
        v1 -= v1
        if v3 {
            if true {
                var v16 float32 = (((float32(0.1) / (1.5 / (((1.5 + 1.5) + (1.5 + (1.5 - -2.25))) - 3.14159))) - 1e10) / (0.0 - -2.25))
                v7 = (v7 - (^(v7)))
                _ = v16
            } else {
                v8 -= -2147483648
                var v17 int16 = (v2 << ((612) & 63))
                var v18 float32 = 0.1
                _ = v17
                _ = v18
            }
            if v3 {
                if len(v10) > 0 {
                    v10[0] = 693
                }
            } else {
                v19 := GenS2{F0: v7, F1: (v8 & v8)}
                fmt.Println(v2, v3, v2)
                var v20 int64 = v1
                _ = v19
                _ = v20
            }
            v2 -= (v2 ^ (v2 >> ((65535) & 63)))
        } else {
            switch -718 {
            case -414:
                v2 -= -1
            case -860:
                fmt.Println(v8)
            case 42:
                for fi21 := 0; fi21 < 5; fi21++ {
                    if v3 {
                        v1 = v1
                    }
                    _ = fi21
                }
            default:
                for fi22 := 0; fi22 < 4; fi22++ {
                    v1 = v1
                    v8 += (v8 >> ((18446744073709551615) & 63))
                    _ = fi22
                }
            }
            if ((true && ((!(genFn0() < 1)) && v3)) || v3) {
                switch "\n\t" {
                case "hello":
                    v23 := GenS3{F0: 364}
                    _ = v23
                }
                var v24 string = (("" + "a") + "foo bar")
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 int64 = -78
        v8 = (v8 + v8)
        var v27 float64 = 3.14159
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint64 = 18446744073709551615
    v7 = v7
    var v29 uint64 = (^(v28))
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 int8 = 127
        for fi32 := 0; fi32 < 5; fi32++ {
            fmt.Println(v8, v31, v31)
            fmt.Println(v28)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 uint64; F1 string }
func genFn0() uint64 {
    var v7 uint = 893
    v8 := GenS1{F0: 377, F1: -2147483648, F2: v7}
    v9 := []int{-9223372036854775808, -429, 128, 32768}
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []int{-9223372036854775808, 127, 9223372036854775807, 987}
    v2 := []int{-9223372036854775808, 9223372036854775807, 2147483647, -1}
    var v3 int = -994
    var v4 uint64 = 0
    v5 := []string{}
    v6 := make(map[uint32]uint32)
    var v10 uint8 = 151
    var v11 int16 = 386
    v11 = 1
    v12 := v6[415]
    v3 += v3
    if ((((true || (-1 >= -711)) || ("hello" < ("hello" + ("foo bar" + "\n\t")))) && (!(!true))) || (!(((!false) || false) && true))) {
        v13 := GenS1{F0: 2, F1: -1, F2: 4294967295}
        _ = v13
    } else {
        fmt.Println(v10, v10, v10)
        v12 -= (v12 % ((uint32(v10)) | 1))
        fmt.Println(v4, v3)
    }
    v14 := GenS1{F0: 42, F1: 0, F2: 2}
    switch v3 {
    case 959:
        for fi15 := 0; fi15 < 17; fi15++ {
            if (v3 < (v3 * fi15)) {
                v3 = fi15
            } else {
                fmt.Println(v12)
            }
            _ = fi15
        }
    }
    var v16 uint = 1
    switch (v3 & 598) {
    case 32767:
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int = v3
            _ = fi17
            _ = v18
        }
    case 9223372036854775807:
        if (-67 < 127) {
            v12 -= v12
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int16 = (v11 << ((v16) & 63))
                _ = fi19
                _ = v20
            }
            var v21 uint16 = 1
            _ = v21
        }
    case 29:
        fmt.Println(v12, v16)
    }
    fmt.Println(v4, v10, v3)
    for fi22 := 0; fi22 < 12; fi22++ {
        fmt.Println(v12, fi22, v12)
        _ = fi22
    }
    if false {
        var v23 string = "a"
        _ = v23
    } else {
        var v24 float32 = 0.0
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS2{F0: 0, F1: ""}
                fmt.Println(v16, v4, v11)
                _ = fi26
                _ = v27
            }
            fmt.Println(v24, v3, v11)
            for fi28 := 0; fi28 < 5; fi28++ {
                v1 = append(v1, fi28)
                if true {
                    var v29 int8 = 42
                    v29 = (v29 + v29)
                    fmt.Println(v10, v4, v10)
                    _ = v29
                }
                _ = fi28
            }
            _ = fi25
        }
        if (-933 != -1) {
            var v30 uint32 = 4294967295
            _ = v30
        } else {
            if (!(((v11 % ((v11) | 1)) > (v11 + (v11 ^ v11))) || true)) {
                var v31 uint64 = v4
                v4 += (v4 / ((v31) | 1))
                _ = v31
            }
            var v32 int = v3
            _ = v32
        }
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

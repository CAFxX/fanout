package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 int64; F1 int; F2 uint8 }
type GenS4 struct { F0 bool; F1 string }
func genFn0(p0 uint64) uint {
    v6 := []int{-366, 32767, -9223372036854775808}
    p0 = (p0 % (((^(p0))) | 1))
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]bool)
    var v3 uint = 255
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v7 := GenS1{F0: "a"}
    for fi8 := 0; fi8 < 3; fi8++ {
        if true {
            v3 = genFn0(4294967295)
            var v9 int8 = 1
            _ = v9
        }
        _ = fi8
    }
    v10 := GenS2{F0: 1, F1: 149, F2: "foo bar"}
    v11 := GenS3{F0: 42, F1: -374, F2: 1}
    var v12 bool = (2 < 127)
    if v12 {
        v12 = (42 >= 32767)
        v12 = true
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := []int{fi13}
        for fi15 := 0; fi15 < 18; fi15++ {
            v12 = true
            v14 = append(v14, fi15)
            var v16 uint = v3
            _ = fi15
            _ = v16
        }
        v12 = v12
        _ = fi13
        _ = v14
    }
    if (v12 || v12) {
        v3 += (v3 / ((v3) | 1))
        var v17 string = ("" + "a")
        if v12 {
            v18 := GenS1{F0: v17}
            v17 = v17
            _ = v18
        } else {
            v17 = v17
        }
        _ = v17
    }
    if (((^(v1)) == v1) && v12) {
        if (255 == 2) {
            delete(v5, 2147483648)
            v19 := GenS3{F0: -90, F1: 127, F2: 0}
            _ = v19
        } else {
            v20 := []int{}
            var v21 int64 = -99
            _ = v20
            _ = v21
        }
    } else {
        var v22 int64 = 9223372036854775807
        switch -1 {
        case -1:
            for fi23 := 0; fi23 < 7; fi23++ {
                for fi24 := 0; fi24 < 17; fi24++ {
                    v25 := v5[42]
                    switch (fi23 << (((v3 & (v3 << ((4294967295) & 63)))) & 63)) {
                    case -9223372036854775808:
                        if v12 {
                            var v26 uint16 = 205
                            v26 -= (v26 & v26)
                            var v27 int32 = 100
                            _ = v26
                            _ = v27
                        }
                    case 4294967295:
                        fmt.Println(v22, v22, v1)
                    case 128:
                        fmt.Println(fi23, v22, fi24)
                    default:
                        v3 = (v3 + (v3 << ((v3) & 63)))
                    }
                    v4[330] = (fi23 - fi24)
                    _ = fi24
                    _ = v25
                }
                if true {
                    var v28 uint16 = 65535
                    var v29 uint8 = 127
                    _ = len(v2)
                    _ = v28
                    _ = v29
                }
                _ = fi23
            }
        default:
            v1 += (v1 - v1)
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            if false {
                v31 := []int{fi30, (fi30 % (((-(fi30))) | 1))}
                v3 = v3
                _ = v31
            }
            v32 := GenS4{F0: v12, F1: ("hello" + "hello")}
            var v33 int32 = v1
            _ = fi30
            _ = v32
            _ = v33
        }
        _ = v22
    }
    v12 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}

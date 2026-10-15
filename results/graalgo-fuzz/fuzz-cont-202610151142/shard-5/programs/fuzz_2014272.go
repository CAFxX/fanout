package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int = 4294967295
    var v2 int = 32767
    var v3 uint = 2
    v4 := make(map[uint32]uint32)
    v3 = (^(v3))
    switch v2 {
    case 110:
        v5 := []int{v1}
        _ = v5
    case -9223372036854775808:
        if true {
            for fi6 := 0; fi6 < 16; fi6++ {
                fmt.Println(v1)
                v1 = v2
                _ = fi6
            }
        } else {
            var v7 float64 = (0.0 - 0.0)
            _ = v7
        }
    default:
        _ = len(v4)
    }
    v8 := []int{(v2 + (v2 + (v2 & v2))), (v1 << ((42) & 63))}
    var v9 uint32 = 236
    fmt.Println(v1, v9)
    var v10 uint32 = 65535
    v11 := GenS1{F0: 24}
    if false {
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                v9 -= (v9 >> (((v3 - v3)) & 63))
                _ = fi13
            }
            _ = fi12
        }
        var v14 int = (v1 & (v1 / ((v1) | 1)))
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                v2 += fi16
                v9 = 637
                v17 := GenS2{F0: 558, F1: false, F2: 2}
                _ = fi16
                _ = v17
            }
            fmt.Println(v10)
            _ = fi15
        }
        _ = v14
    }
    v18 := GenS2{F0: -134, F1: (32767 > 1), F2: 42}
    if true {
        var v19 string = "hello"
        for fi20 := 0; fi20 < 10; fi20++ {
            switch v19 {
            case "":
                if false {
                    v19 = "foo bar"
                }
            case "a":
                for fi21 := 0; fi21 < 20; fi21++ {
                    if (2147483647 >= 2147483647) {
                        if len(v8) > 0 {
                            v8[0] = fi21
                        }
                    } else {
                        fmt.Println(v2, fi21, fi20)
                        v1 = 4294967295
                    }
                    _ = fi21
                }
            }
            _ = fi20
        }
        v22 := GenS3{F0: 793, F1: -893}
        _ = v19
        _ = v22
    } else {
        v2 = (v2 ^ 128)
        for fi23 := 0; fi23 < 19; fi23++ {
            v2 += fi23
            _ = len(v8)
            var v24 float32 = 1.5
            _ = fi23
            _ = v24
        }
    }
    v3 -= v3
    if true {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v25 := GenS1{F0: -9223372036854775808}
        for fi26 := 0; fi26 < 5; fi26++ {
            if (int8(v1) < -30) {
                for fi27 := 0; fi27 < 3; fi27++ {
                    fmt.Println(fi26, v10, v1)
                    v10 = v10
                    var v28 int16 = 32767
                    _ = fi27
                    _ = v28
                }
                switch (^(v2)) {
                case 128:
                    v29 := GenS1{F0: 321}
                    _ = v29
                case 550:
                    var v30 int16 = -32768
                    _ = v30
                case 127:
                    delete(v4, v10)
                default:
                    if false {
                        v10 += v10
                    }
                }
            }
            if (0 < 18446744073709551615) {
                v9 += v9
                _ = len(v8)
            } else {
                v31 := GenS4{F0: 0, F1: 477}
                _ = v31
            }
            _ = fi26
        }
        _ = v25
    } else {
        var v32 int16 = 42
        var v33 float32 = 0.1
        _ = v32
        _ = v33
    }
    fmt.Println(v3, v3)
    var v34 int64 = -1
    v35 := GenS3{F0: 0, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v8
    _ = v9
}

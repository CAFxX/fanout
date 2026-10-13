package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int32 = -1
    var v2 uint8 = 128
    var v3 float32 = ((1e10 - -2.25) + (float32(-2.25) + 1e10))
    var v4 uint64 = 999
    v5 := []int{}
    var v6 int8 = 0
    v7 := GenS1{F0: false, F1: -225, F2: v2}
    if len(v5) > 0 {
        v5[0] = 128
    }
    var v8 int32 = (^(v1))
    switch 891 {
    case 0:
        if (-128 > 23) {
            if len(v5) > 0 {
                v5[0] = -1
            }
            v9 := GenS1{F0: true, F1: -1, F2: 128}
            var v10 float32 = v3
            _ = v10
            _ = v9
        }
    }
    switch "hello" {
    case "":
        fmt.Println(v2, v6, v3)
    }
    var v11 uint64 = uint64(v2)
    for fi12 := 0; fi12 < 1; fi12++ {
        v2 = v2
        switch (fi12 ^ -684) {
        case -2147483648:
            if (42 == 4294967295) {
                var v13 uint16 = 255
                v14 := GenS2{F0: 18446744073709551615}
                for fi15 := 0; fi15 < 17; fi15++ {
                    v13 = v13
                    v16 := GenS1{F0: (false || (42 != 0)), F1: 0, F2: (v2 >> ((1) & 63))}
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            } else {
                if (false && (!true)) {
                    v5 = append(v5, (fi12 & -971))
                    var v17 string = ("hello" + "hello")
                    _ = v17
                }
                if (fi12 >= fi12) {
                    switch (fi12 * fi12) {
                    case -908:
                        v11 -= v4
                    case 32768:
                        v2 = (v2 & (^(v2)))
                    }
                } else {
                    v3 -= float32(1.5)
                }
                _ = len(v5)
            }
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v1 += (v8 / (((v1 / ((42) | 1))) | 1))
        _ = fi12
    }
    var v18 string = ""
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 uint32 = 498
        _ = fi19
        _ = v20
    }
    v2 += 88
    v1 = (-(v8))
    if ((v1 | (^(v1))) > (v8 & (v8 % ((0) | 1)))) {
        for fi21 := 0; fi21 < 7; fi21++ {
            if false {
                for fi22 := 0; fi22 < 2; fi22++ {
                    _ = len(v5)
                    fmt.Println(v1)
                    v23 := GenS2{F0: 2}
                    _ = fi22
                    _ = v23
                }
                var v24 int8 = (v6 ^ 100)
                _ = len(v5)
                _ = v24
            } else {
                if false {
                    v11 += 0
                    v25 := GenS2{F0: 18446744073709551615}
                    _ = v25
                }
                if len(v5) > 0 {
                    v5[0] = (fi21 / (((fi21 ^ fi21)) | 1))
                }
            }
            _ = fi21
        }
        v26 := GenS2{F0: 2}
        _ = v26
    } else {
        v27 := GenS1{F0: ((v2 << ((65535) & 63)) >= uint8(v11)), F1: 317, F2: (v2 >> ((606) & 63))}
        switch 32768 {
        case 4294967295:
            v2 = (v2 * 31)
        case 525:
            v28 := GenS1{F0: false, F1: -1, F2: (^(v2))}
            _ = v28
        case 468:
            if false {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v4 = (v11 & 1)
                    if (((v3 * (-2.25 * v3)) / v3) < (3.14159 + 0.1)) {
                        var v30 uint64 = (^(v4))
                        _ = v30
                    }
                    var v31 float32 = v3
                    _ = fi29
                    _ = v31
                }
                var v32 bool = true
                _ = v32
            }
        }
        switch v18 {
        case "a":
            v33 := GenS3{F0: v8}
            _ = v33
        default:
            v11 = (v11 % (((^(v4))) | 1))
        }
        _ = v27
    }
    fmt.Println(v2)
    v8 = v1
    if len(v5) > 0 {
        v5[0] = int(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

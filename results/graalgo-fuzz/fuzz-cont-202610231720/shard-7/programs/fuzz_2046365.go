package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]string)
    var v3 float32 = -2.25
    var v4 uint32 = 661
    v5 := []int{128}
    var v6 int16 = -299
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 2, F1: 47, F2: (^(v6))}
            v1 = (v1 | 42)
            if true {
                v3 += v3
                var v10 int32 = -1
                fmt.Println(fi7)
                var v11 uint64 = 0
                _ = v10
                _ = v11
            } else {
                var v12 uint = 65535
                v12 -= (v12 << (((v12 | (^(v12)))) & 63))
                var v13 float32 = 1.5
                _ = v12
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v4)
            fmt.Println(fi14, fi7)
            if false {
                v4 -= v4
                v3 = 1.5
            }
            _ = fi14
        }
        v15 := GenS2{F0: (v6 ^ (v6 + v6)), F1: 42}
        _ = fi7
        _ = v15
    }
    v6 += v6
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            if (255 >= 127) {
                v5 = append(v5, fi16)
                v3 -= v3
                if true {
                    v1 = v1
                    v18 := GenS3{F0: 159, F1: 42, F2: -2147483648}
                    _ = v18
                }
            }
            _ = fi17
        }
        _ = fi16
    }
    var v19 int64 = v1
    v20 := GenS1{F0: v4, F1: 1, F2: v6}
    v21 := GenS2{F0: v6, F1: 255}
    v4 = 2147483648
    var v22 int = -113
    for fi23 := 0; fi23 < 5; fi23++ {
        if true {
            switch ("foo bar" + "a") {
            case "a":
                v6 = v6
            case "":
                var v24 bool = true
                _ = v24
            default:
                v19 -= (v19 ^ (v1 ^ v1))
            }
            if (!true) {
                v19 += v19
            }
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                if true {
                    v1 += -9223372036854775808
                    var v26 float32 = (0.0 / (0.1 - v3))
                    _ = v26
                } else {
                    v27 := GenS2{F0: v6, F1: 1}
                    v28 := v2[fi23]
                    if len(v5) > 0 {
                        v5[0] = 127
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int = 32768
                var v31 int64 = (v1 % (((v19 - v19)) | 1))
                var v32 uint8 = 42
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 string = ""
                v35 := GenS1{F0: 95, F1: -1, F2: (v6 / ((643) | 1))}
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

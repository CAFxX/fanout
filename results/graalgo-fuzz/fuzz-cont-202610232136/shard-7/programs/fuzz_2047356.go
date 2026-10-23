package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 string; F2 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int16; F2 uint8 }
func genFn0(p0 uint, p1 int16, p2 int64) uint16 {
    var v6 uint8 = 215
    fmt.Println(v6)
    var v7 bool = (18446744073709551615 == 261)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 282
}

func main() {
    var v1 uint32 = 0
    var v2 float32 = (((float32(float64(((0.1 * (-2.25 / 1.0)) - 1.5))) / (float32(0.0) - 3.14159)) - 3.14159) * 1e10)
    var v3 int = 42
    var v4 uint8 = 0
    v5 := []int{}
    var v8 int16 = 32767
    if false {
        var v9 uint = 255
        _ = v9
    }
    v10 := GenS1{F0: 1}
    if (false && (!(127 >= 0))) {
        v11 := GenS1{F0: 2147483647}
        var v12 bool = false
        fmt.Println(v12, v8, v1)
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: 400, F1: (v4 >> ((410) & 63)), F2: 500}
        v1 -= (v1 >> ((0) & 63))
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                if ((65535 == 151) || (2147483647 == -1)) {
                    var v16 bool = (127 <= 53)
                    fmt.Println(v2)
                    _ = v16
                } else {
                    v5 = append(v5, -494)
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    fmt.Println(fi15, v4, v3)
                    var v17 float32 = float32(0.1)
                    _ = v17
                }
                v1 = 2147483648
                var v18 bool = true
                _ = fi15
                _ = v18
            }
            if false {
                v4 -= (v4 << ((255) & 63))
            } else {
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 bool = (true || false)
                    _ = fi19
                    _ = v20
                }
                var v21 uint = 42
                _ = v21
            }
            fmt.Println(v1, fi14)
            _ = fi14
        }
        _ = v13
    }
    var v22 uint16 = 1
    switch (v3 & v3) {
    case 2147483647:
        var v23 bool = true
        _ = v23
    default:
        v24 := GenS3{F0: 0, F1: v3, F2: (v1 ^ (v1 + (v1 / (((v1 - 703)) | 1))))}
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            v2 += float32(1.5)
            fmt.Println(v8)
            v4 += (v4 + uint8(v1))
            _ = fi26
        }
        _ = fi25
    }
    v27 := GenS4{F0: 42}
    var v28 uint8 = uint8(v1)
    for fi29 := 0; fi29 < 1; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            v3 = (fi30 * 127)
            v31 := GenS5{F0: -1, F1: ("hello" + (("\n\t" + "") + "hello")), F2: 960}
            fmt.Println(v28)
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    var v32 int = (v3 / ((-9223372036854775808) | 1))
    fmt.Println(v8, v22)
    if (523 <= -9223372036854775808) {
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := GenS3{F0: 0, F1: v3, F2: v1}
            _ = fi33
            _ = v34
        }
        var v35 uint = 255
        switch ((("" + "foo bar") + "hello") + "\n\t") {
        case "\n\t":
            fmt.Println(v3)
        case "hello":
            if (v1 != v1) {
                v28 = (^(v28))
                fmt.Println(v3, v8, v8)
            }
        }
        _ = v35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v37 := GenS6{F0: v8}
        _ = fi36
        _ = v37
    }
    var v38 uint32 = (v1 % (((^(v1))) | 1))
    v39 := GenS7{F0: -9223372036854775808, F1: (v8 << ((42) & 63)), F2: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v8
}

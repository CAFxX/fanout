package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int; F2 uint }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint8; F1 int32; F2 int }
func genFn0() int16 {
    var v4 uint16 = 0
    var v5 uint32 = 65535
    v5 += v5
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 uint, p1 uint64, p2 uint16) uint32 {
    p1 += p1
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = (float64(0.0) / 3.14159)
        v7 += ((v7 + 1.5) * v7)
        var v8 string = "hello"
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (true || (true || ((p2 ^ p2) <= (p2 - p2)))), F1: int(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}

func main() {
    var v1 int = 128
    var v2 int64 = 9223372036854775807
    var v3 string = "\n\t"
    var v10 int = -1
    switch v10 {
    case 127:
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS2{F0: 18446744073709551615, F1: (v2 - (v2 * (v2 >> ((364) & 63)))), F2: 2}
            _ = fi11
            _ = v12
        }
    case 552:
        if false {
            v1 = 32768
            var v13 float32 = 0.0
            var v14 int32 = -237
            _ = v13
            _ = v14
        } else {
            v15 := GenS3{F0: v2}
            _ = v15
        }
    default:
        v3 = v3
    }
    var v16 int = -274
    fmt.Println(v10, v2)
    fmt.Println(v2, v3)
    var v17 uint32 = 2
    v18 := GenS4{F0: 0, F1: -1, F2: 1}
    if ((-2.25 > float64(3.14159)) && (3.14159 != -2.25)) {
        v19 := GenS5{F0: -846, F1: v17, F2: 431}
        var v20 uint8 = 128
        for fi21 := 0; fi21 < 15; fi21++ {
            v2 = (v2 >> ((358) & 63))
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    fmt.Println(v10)
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 float32 = 3.14159
        if (4294967295 == 4294967295) {
            v16 = 2147483647
            var v24 bool = (42 >= 255)
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    v25 := GenS6{F0: (v17 >> ((42) & 63)), F1: 128}
    if ((v17 >> ((1) & 63)) >= (^(v17))) {
        v26 := GenS7{F0: 2, F1: -1, F2: 32767}
        _ = v26
    }
    fmt.Println(v10, v10, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}

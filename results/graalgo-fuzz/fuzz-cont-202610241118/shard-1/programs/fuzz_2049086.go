package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint8) int16 {
    p0 = p0
    fmt.Println(p0, p1)
    p0 = p0
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 = true
        fmt.Println(p1, p1, p0)
        p1 = p1
        _ = fi5
    }
    var v6 int = -1
    _ = p0
    _ = p1
    _ = v6
    return int16(p1)
}
func genFn1() int {
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint = 2
            v8 = v8
            _ = fi7
            _ = v8
        }
        var v9 int = 32768
        _ = v9
    }
    v10 := GenS1{F0: 337, F1: (4294967295 > 866)}
    var v11 float32 = -2.25
    var v12 uint64 = 622
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 uint32 = 771
    var v2 uint8 = 127
    var v3 int8 = -45
    var v4 uint8 = (v2 ^ (v2 & v2))
    v3 += v3
    var v13 uint = 0
    if (v4 == (v2 - uint8(v1))) {
        v4 = (v4 - v4)
        v14 := GenS2{F0: 1, F1: -424}
        _ = v14
    }
    var v15 uint8 = v4
    var v16 int64 = 9223372036854775807
    v15 += (v2 * (v15 >> (((v13 ^ v13)) & 63)))
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint32 = (v1 & 1)
        v19 := GenS3{F0: 0, F1: 18446744073709551615}
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS4{F0: ("a" + "\n\t")}
            _ = fi20
            _ = v21
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v1, v15)
    var v22 int8 = v3
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS5{F0: 2}
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}

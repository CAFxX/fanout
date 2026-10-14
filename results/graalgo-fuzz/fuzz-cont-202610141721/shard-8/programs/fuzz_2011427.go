package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
func genFn0(p0 bool, p1 uint32) string {
    v6 := GenS1{F0: -2147483648, F1: 428, F2: 32767}
    p0 = (p0 && ((p0 && (!true)) || (36 > 128)))
    p1 += p1
    var v7 int64 = 42
    v8 := GenS2{F0: (v7 >> ((674) & 63)), F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 int32) uint8 {
    if ((255 <= 0) || (!false)) {
        var v9 int16 = 0
        var v10 int16 = (^(v9))
        _ = v10
        _ = v9
    }
    var v11 float32 = (3.14159 - 0.1)
    _ = p0
    _ = v11
    return 208
}

func main() {
    var v1 string = (((("hello" + ("hello" + "")) + "foo bar") + "foo bar") + ("a" + (("foo bar" + "\n\t") + "\n\t")))
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 float64 = 3.14159
    var v5 uint64 = 1
    v1 = v1
    switch -599 {
    case -2147483648:
        var v12 uint32 = 2
        _ = v12
    }
    v4 = 1.5
    v1 = v1
    v13 := GenS3{F0: 2147483648, F1: -409, F2: 18446744073709551615}
    var v14 int8 = -128
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 bool = false
        var v17 uint8 = 190
        v4 -= v4
        _ = fi15
        _ = v16
        _ = v17
    }
    v14 -= (v14 & -104)
    if true {
        delete(v3, (v1 + v1))
        v14 = -1
    }
    v1 = ("hello" + (v1 + "a"))
    fmt.Println(v1, v1)
    var v18 float64 = ((v4 + v4) / v4)
    v18 -= (v18 / 1e10)
    v5 += (v5 << ((42) & 63))
    v18 -= -2.25
    var v19 uint = 84
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

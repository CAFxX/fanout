package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS4 struct { F0 string; F1 uint32; F2 uint64 }
func genFn0() uint16 {
    if (-323 > -32768) {
        var v6 bool = ((true || (-207 <= 406)) && (4294967295 == 980))
        var v7 bool = (!v6)
        _ = v6
        _ = v7
    } else {
        var v8 int = -750
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int16 = -137
            var v11 int32 = -658
            v11 -= v11
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    if (!true) {
        v12 := GenS1{F0: 42}
        var v13 int64 = 0
        v13 = 9223372036854775807
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 32767}
    _ = v14
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 0
    var v3 int16 = (v2 % (((v2 << ((28) & 63))) | 1))
    var v4 float32 = 0.1
    var v5 int32 = -2147483648
    if true {
        v3 += (v3 << ((4294967295) & 63))
    }
    var v15 string = ("\n\t" + ("\n\t" + "a"))
    v16 := []int{2147483647, 0, -556}
    if false {
        switch v15 {
        case "\n\t":
            v17 := GenS1{F0: 0}
            _ = v17
        case "hello":
            switch v15 {
            case "foo bar":
                v15 = (((v15 + v15) + v15) + v15)
            }
        case "foo bar":
            fmt.Println(v3, v4, v5)
        default:
            var v18 int8 = 127
            _ = v18
        }
        v4 = 1e10
    } else {
        if len(v16) > 0 {
            v16[0] = -9223372036854775808
        }
        switch 128 {
        case 312:
            var v19 uint32 = 42
            _ = v19
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                v4 -= ((v4 + 0.0) - v4)
                _ = fi20
            }
        }
    }
    fmt.Println(v15, v5, v2)
    v2 -= v2
    v15 = "\n\t"
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint16 = 2
        fmt.Println(v3, v4)
        var v23 uint = 4294967295
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 285, F1: 834, F2: genFn0()}
    v15 = v15
    for fi25 := 0; fi25 < 17; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v5, fi26)
            var v27 int8 = 0
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    if true {
        v2 = 32767
    }
    var v28 float32 = 1e10
    v29 := GenS3{F0: 4294967295, F1: 1, F2: 2}
    v30 := GenS4{F0: v15, F1: 1, F2: 116}
    var v31 int = -604
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

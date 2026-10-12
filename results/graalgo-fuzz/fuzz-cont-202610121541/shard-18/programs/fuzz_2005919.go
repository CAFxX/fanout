package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint8; F2 int64 }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint) string {
    v7 := GenS1{F0: 127, F1: 0, F2: 18446744073709551615}
    fmt.Println(p0)
    var v8 int = -752
    for fi9 := 0; fi9 < 16; fi9++ {
        switch v8 {
        case 350:
            v10 := GenS1{F0: 42, F1: 45, F2: 478}
            _ = v10
        case 1:
            v11 := []int{(fi9 & (fi9 ^ -9223372036854775808)), fi9, (-(v8)), (v8 | v8)}
            _ = v11
        case 541:
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 int16 = 1
                var v14 float32 = (0.1 * 1.5)
                p0 = 42
                var v15 int32 = 42
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        default:
            p0 += 65535
        }
        var v16 int8 = -128
        var v17 bool = true
        _ = fi9
        _ = v16
        _ = v17
    }
    if (0 <= 952) {
        v8 -= v8
    } else {
        var v18 float32 = (3.14159 / 1e10)
        p0 = (p0 & 42)
        _ = v18
    }
    _ = p0
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 % (((v1 + v1)) | 1))
    var v3 int64 = -1
    v4 := make(map[string]bool)
    var v5 uint32 = 0
    var v6 int64 = 363
    var v19 uint = 610
    var v20 int16 = 1000
    v19 = (v19 + v19)
    v4[("foo bar" + "\n\t")] = (1000 != v20)
    if true {
        v1 -= -128
        _ = len(v4)
    }
    v21 := GenS2{F0: (v20 ^ (v20 - v20))}
    if false {
        var v22 string = ""
        fmt.Println(v6)
        _ = v22
    }
    if true {
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                if true {
                    v25 := []int{fi24}
                    _ = v25
                }
                _ = fi24
            }
            v26 := GenS3{F0: v20, F1: 223, F2: (v3 >> (((v19 / ((uint(v5)) | 1))) & 63))}
            v27 := GenS3{F0: -1, F1: 76, F2: 9223372036854775807}
            _ = fi23
            _ = v26
            _ = v27
        }
    } else {
        v2 -= v2
    }
    if true {
        fmt.Println(v1)
        v28 := GenS1{F0: 0, F1: 42, F2: 0}
        var v29 uint = v19
        _ = v28
        _ = v29
    } else {
        fmt.Println(v5, v19, v1)
        if true {
            fmt.Println(v1)
        }
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        switch (("hello" + "a") + (("hello" + "") + "\n\t")) {
        case "a":
            var v31 int = int(v5)
            _ = v31
        case "":
            v20 -= v20
        case "foo bar":
            var v32 uint32 = v5
            _ = v32
        default:
            v33 := GenS3{F0: (v20 + v20), F1: 63, F2: 9223372036854775807}
            _ = v33
        }
        v34 := GenS4{F0: -4, F1: 42}
        _ = fi30
        _ = v34
    }
    var v35 int64 = (v6 ^ (v3 + int64(v2)))
    if (2 == 1) {
        if (false || ((true && false) || false)) {
            v5 = (v5 | v5)
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            v4[("" + "hello")] = true
            for fi37 := 0; fi37 < 19; fi37++ {
                fmt.Println(v35, fi36, v19)
                v19 = (v19 - v19)
                fmt.Println(v1, v19, fi37)
                _ = fi37
            }
            var v38 uint32 = 2
            _ = fi36
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}

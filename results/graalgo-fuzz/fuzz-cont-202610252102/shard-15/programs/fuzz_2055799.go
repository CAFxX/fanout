package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16) int8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int64 = -1
        p0 += (p0 % ((1) | 1))
        _ = fi7
        _ = v8
    }
    p0 -= (p0 | p0)
    _ = p0
    return -128
}
func genFn1(p0 uint8, p1 int8, p2 int) string {
    var v9 int32 = -250
    p1 = genFn0(32767)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 128
    var v3 bool = true
    v4 := make(map[uint32]int)
    var v5 int8 = -111
    var v6 string = "hello"
    _ = len(v1)
    v10 := GenS1{F0: v5, F1: 2, F2: v6}
    var v11 bool = (1 > 42)
    var v12 uint16 = 1
    var v13 int64 = 864
    var v14 int64 = 42
    v5 = int8(v13)
    v13 = v14
    for fi15 := 0; fi15 < 8; fi15++ {
        if (32767 > 863) {
            fmt.Println(v2)
        } else {
            if ((!v3) && v3) {
                var v16 int8 = genFn0((v12 * v12))
                _ = v16
            }
            var v17 int16 = 0
            _ = v17
        }
        v18 := []int{fi15, -2147483648, v2, (v2 & fi15)}
        _ = fi15
        _ = v18
    }
    if ((v5 / (((v5 << ((65535) & 63))) | 1)) <= v5) {
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v14)
            _ = fi19
        }
        v13 = 42
        var v20 int16 = -452
        _ = v20
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v13 -= v13
            var v22 uint32 = 2147483648
            _ = fi21
            _ = v22
        }
        var v23 int8 = (v5 + v5)
        for fi24 := 0; fi24 < 14; fi24++ {
            v6 = (v6 + v6)
            for fi25 := 0; fi25 < 18; fi25++ {
                v11 = (2147483648 != 393)
                var v26 uint16 = (v12 ^ v12)
                fmt.Println(v26, v13, v14)
                _ = fi25
                _ = v26
            }
            v27 := GenS1{F0: int8(v12), F1: (v12 - (v12 << ((255) & 63))), F2: ""}
            _ = fi24
            _ = v27
        }
        _ = v23
    }
    if (!true) {
        v13 = (v13 % (((v13 & 121)) | 1))
        switch (v2 / ((698) | 1)) {
        case 4294967295:
            fmt.Println(v2)
        case -2147483648:
            v14 = v14
        case 0:
            v28 := GenS2{F0: 1, F1: -1, F2: -1}
            _ = v28
        default:
            fmt.Println(v3, v11)
        }
        var v29 int = v2
        _ = v29
    }
    v30 := GenS3{F0: (v2 + (v2 | (^(v2))))}
    var v31 int = -806
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 int16 = 1
        var v34 float64 = ((float64(0.0) * 1e10) + 0.1)
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

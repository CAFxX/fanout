package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    var v2 int = 42
    var v3 uint32 = 0
    for fi4 := 0; fi4 < 8; fi4++ {
        if false {
            var v5 string = "a"
            _ = v5
        } else {
            var v6 int64 = -9223372036854775808
            _ = v6
        }
        var v7 string = ("hello" + ("a" + ("hello" + "foo bar")))
        if ((v3 - (^(v3))) == 1) {
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(v7, fi8)
                _ = fi8
            }
            if (!false) {
                v7 = "hello"
                v1 = append(v1, (v7 + v7))
            } else {
                v7 = v7
                v3 = v3
                if false {
                    var v9 float32 = 0.1
                    _ = v9
                } else {
                    var v10 uint = 18446744073709551615
                    v11 := GenS1{F0: v2, F1: int32(v3)}
                    _ = v10
                    _ = v11
                }
            }
        }
        _ = fi4
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v3)
    if (!((48 == 1) || false)) {
        v3 -= (v3 + v3)
        v12 := GenS1{F0: v2, F1: -480}
        _ = v12
    } else {
        var v13 float64 = (1e10 / 1.0)
        v13 -= ((v13 - 0.0) + v13)
        _ = v13
    }
    v2 = (v2 - 291)
    var v14 uint32 = 122
    var v15 int16 = 742
    for fi16 := 0; fi16 < 8; fi16++ {
        if (false || false) {
            v17 := GenS1{F0: (-(v2)), F1: 0}
            var v18 uint8 = 143
            var v19 int8 = 82
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(v15, v15)
        var v21 float32 = 1.5
        v15 += v15
        _ = fi20
        _ = v21
    }
    v1 = append(v1, "")
    v22 := GenS1{F0: v2, F1: -741}
    var v23 int32 = 42
    v3 = (^(v3))
    var v24 uint16 = 65535
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 uint32 = v14
        _ = fi25
        _ = v26
    }
    switch (v2 & 1) {
    case 32768:
        if true {
            v27 := GenS2{F0: 0, F1: (v14 & 844), F2: (v23 & (v23 * -76))}
            _ = v27
        } else {
            var v28 float64 = 0.0
            _ = v28
        }
    case -911:
        v29 := GenS2{F0: 1, F1: (v3 * v3), F2: (v23 % (((v23 * v23)) | 1))}
        _ = v29
    case 42:
        fmt.Println(v2, v3)
    }
    v24 += (v24 ^ 0)
    var v30 uint32 = (v14 * (v3 ^ (v3 << ((2) & 63))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 bool; F1 int32; F2 int64 }
func genFn0(p0 int32, p1 int32) uint64 {
    p1 -= -2147483648
    fmt.Println(p1)
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 uint16 = 472
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int = (fi6 + fi4)
            if true {
                if true {
                    v5 -= v5
                }
            } else {
                v8 := GenS1{F0: v5, F1: 1, F2: 0}
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        p1 = int32(v5)
        _ = fi4
        _ = v5
    }
    p0 += p1
    if (!(33 < -1)) {
        if false {
            p1 += (p1 * p1)
        }
        var v9 float32 = -2.25
        _ = v9
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int32 = 1
    var v2 int32 = 1
    var v3 string = "a"
    switch 32767 {
    case 2147483647:
        var v10 int8 = -1
        _ = v10
    }
    v3 = "\n\t"
    v11 := []int{-9223372036854775808}
    v1 = v1
    if (-32768 > 1000) {
        v2 += 318
        fmt.Println(v2)
    } else {
        var v12 int = 32767
        _ = v12
    }
    var v13 uint32 = 0
    v14 := GenS1{F0: 1, F1: 4294967295, F2: genFn0((-(v1)), v2)}
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            if false {
                switch (-(fi15)) {
                case 127:
                    var v16 int = fi15
                    _ = v16
                case 9223372036854775807:
                    v2 = -212
                default:
                    fmt.Println(v3, fi15)
                }
                v17 := GenS2{F0: v2, F1: v3}
                _ = v17
            }
            if true {
                v13 = (v13 | (v13 & v13))
                var v18 int = -9223372036854775808
                fmt.Println(fi15)
                _ = v18
            }
        }
        _ = fi15
    }
    var v19 uint = 349
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 float32 = 3.14159
        v22 := GenS2{F0: (v2 % (((v2 % ((242) | 1))) | 1)), F1: "hello"}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int32 = v1
    v24 := GenS3{F0: (1.5 <= -2.25), F1: 354, F2: 863}
    v25 := GenS3{F0: (65535 >= v19), F1: 2147483647, F2: 1}
    v2 = v2
    if false {
        v3 = v3
        var v26 uint32 = 2147483648
        v2 = (v2 * (v1 / ((-1) | 1)))
        _ = v26
    }
    var v27 bool = (false || true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}

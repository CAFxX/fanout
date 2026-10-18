package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint32 }

func main() {
    var v1 int8 = 21
    var v2 int = -103
    var v3 int8 = (v1 - v1)
    var v4 uint64 = 302
    if (18446744073709551615 > 1) {
        if (!true) {
            fmt.Println(v2)
            v1 = v3
            v5 := GenS1{F0: (v4 >= (v4 >> ((176) & 63))), F1: v4, F2: 32767}
            _ = v5
        } else {
            if false {
                var v6 uint = 42
                v2 -= v2
                var v7 uint32 = 120
                _ = v6
                _ = v7
            }
            var v8 bool = true
            _ = v8
        }
    } else {
        v1 = (v1 | v1)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{fi9, (^(fi9)), fi9}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if len(v10) > 0 {
            v10[0] = (fi9 >> ((2) & 63))
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: (v2 * (^(v2))), F1: (v2 - (-(v2))), F2: 42}
    fmt.Println(v1)
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: -9223372036854775808, F1: (v2 / ((42) | 1)), F2: v4}
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int = (v2 / ((v2) | 1))
    v16 := GenS3{F0: -2147483648}
    switch (v15 / ((268) | 1)) {
    case 128:
        v17 := []int{(v15 % ((128) | 1))}
        _ = v17
    case 32768:
        v18 := GenS4{F0: 255}
        _ = v18
    }
    if false {
        switch v2 {
        case 42:
            var v19 uint64 = (v4 ^ v4)
            _ = v19
        case 0:
            v20 := GenS4{F0: 0}
            _ = v20
        case -2147483648:
            v21 := GenS5{F0: (0.1 < -2.25), F1: 1, F2: 4294967295}
            _ = v21
        }
        v22 := GenS4{F0: 82}
        for fi23 := 0; fi23 < 20; fi23++ {
            switch v15 {
            case 0:
                fmt.Println(v1, v3, v2)
            case -2147483648:
                v15 = (-(v2))
            }
            v24 := []int{(-(fi23))}
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v3 = 1
    var v25 int8 = (v1 ^ (v3 & -128))
    v26 := []int{(v15 | v2)}
    var v27 uint = 933
    _ = len(v26)
    v15 = int(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}

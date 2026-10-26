package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16; F2 int8 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 uint16, p1 uint16, p2 uint8) uint64 {
    p2 -= (p2 | (p2 | 0))
    var v4 uint8 = p2
    v5 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int, p1 uint) uint64 {
    if false {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := []int{int(p1)}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v8 uint16 = 2
            _ = fi6
            _ = v7
            _ = v8
        }
        if false {
            v9 := []int{(p0 | (p0 / ((p0) | 1))), 2147483647, p0, 2147483647}
            _ = v9
        } else {
            p0 = (p0 % ((-77) | 1))
            var v10 int16 = 21
            _ = v10
        }
        v11 := []int{p0}
        _ = v11
    } else {
        if false {
            v12 := GenS1{F0: (p0 % ((p0) | 1)), F1: (^(p0)), F2: -32768}
            _ = v12
        } else {
            var v13 float32 = 1e10
            v14 := GenS1{F0: (p0 % (((-(p0))) | 1)), F1: -158, F2: 667}
            _ = v13
            _ = v14
        }
        v15 := []int{}
        _ = v15
    }
    v16 := GenS1{F0: (p0 ^ 128), F1: -9223372036854775808, F2: 32767}
    switch (^(p0)) {
    case 32768:
        switch "" {
        case "a":
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS2{F0: 4294967295}
                _ = fi17
                _ = v18
            }
        }
    case 128:
        v19 := GenS2{F0: 407}
        _ = v19
    case 965:
        if (-1 == 1) {
            var v20 int32 = -517
            if (0.1 >= float32(1e10)) {
                v21 := []int{4294967295, -9223372036854775808, p0, p0}
                if len(v21) > 0 {
                    v21[0] = 42
                }
                _ = v21
            }
            _ = v20
        }
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v16
    return uint64(p1)
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = float32((1.5 + 0.0))
    var v3 uint = 255
    v3 = (v3 * (^(v3)))
    if false {
        var v22 int32 = 1
        v23 := GenS1{F0: -9223372036854775808, F1: 1, F2: 1}
        _ = v22
        _ = v23
    }
    switch -392 {
    case 32767:
        if (!true) {
            var v24 uint32 = 50
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v25 := v1[826]
                _ = v25
            }
            var v26 uint16 = 870
            _ = v24
            _ = v26
        } else {
            v27 := GenS1{F0: 1, F1: 0, F2: 1}
            v28 := GenS1{F0: 522, F1: 32768, F2: 1}
            _ = v27
            _ = v28
        }
    case 599:
        v29 := []int{}
        _ = v29
    default:
        delete(v1, 63)
    }
    var v30 uint = (v3 ^ v3)
    var v31 uint16 = 255
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := []int{fi32, -9223372036854775808}
        _ = len(v33)
        v34 := GenS3{F0: 32767, F1: genFn1((fi32 << ((v30) & 63)), (v3 ^ v3))}
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: -9223372036854775808, F1: int16(v31), F2: 100}
    fmt.Println(v30, v3)
    v30 = v30
    v36 := GenS5{F0: v3, F1: 255, F2: v3}
    for fi37 := 0; fi37 < 15; fi37++ {
        v30 = (v30 | (v3 * 42))
        fmt.Println(v3)
        _ = fi37
    }
    fmt.Println(v2, v2, v2)
    v38 := GenS2{F0: 42}
    for fi39 := 0; fi39 < 17; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            v2 -= 1e10
            var v41 uint64 = 750
            var v42 uint32 = 65535
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v38
}

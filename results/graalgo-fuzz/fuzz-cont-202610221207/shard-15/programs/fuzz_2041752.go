package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS5 struct { F0 int8; F1 int; F2 int64 }
func genFn0(p0 int8, p1 int) int {
    for fi6 := 0; fi6 < 12; fi6++ {
        p1 = (-(fi6))
        _ = fi6
    }
    var v7 int = p1
    _ = p0
    _ = p1
    _ = v7
    return (v7 << ((0) & 63))
}
func genFn1(p0 uint64, p1 uint32) uint64 {
    for fi8 := 0; fi8 < 20; fi8++ {
        switch (fi8 >> ((1) & 63)) {
        case -915:
            p1 -= 2147483647
        case 4294967295:
            var v9 float64 = 3.14159
            _ = v9
        default:
            var v10 uint32 = p1
            _ = v10
        }
        _ = fi8
    }
    var v11 float64 = (-2.25 / 3.14159)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int64 = -1
        if false {
            p1 = 2147483647
            v14 := []int{}
            _ = v14
        } else {
            v11 = 0.0
            p1 = (p1 | (p1 << ((uint(p1)) & 63)))
        }
        _ = fi12
        _ = v13
    }
    v15 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v15
    return uint64(p1)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 128
    v3 := []int{}
    var v4 float32 = -2.25
    var v5 float32 = v4
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint16 = 666
                v19 := GenS2{F0: 42, F1: -1, F2: -106}
                v4 = float32(0.1)
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: 99, F1: 0, F2: -45}
            _ = fi16
            _ = v20
        }
    } else {
        var v21 uint32 = 2147483648
        switch 32768 {
        case 0:
            var v22 int64 = int64(v2)
            _ = v22
        case 2147483647:
            fmt.Println(v5)
        case 128:
            if true {
                v23 := GenS3{F0: 885}
                _ = v23
            } else {
                v2 = v2
                var v24 uint16 = 672
                _ = v24
            }
        default:
            delete(v1, 1)
        }
        _ = v21
    }
    switch 32768 {
    case -263:
        switch "foo bar" {
        case "hello":
            var v25 int = genFn0(-99, -427)
            _ = v25
        default:
            v3 = append(v3, 2147483647)
        }
    case 861:
        switch 127 {
        case 2147483647:
            v5 += float32(0.1)
        case 613:
            v26 := GenS4{F0: v2, F1: 56, F2: 926}
            _ = v26
        case -177:
            if (2 != 18446744073709551615) {
                var v27 bool = (0.1 <= 3.14159)
                var v28 int8 = -45
                switch ("" + ("\n\t" + "a")) {
                case "foo bar":
                    var v29 uint64 = 0
                    _ = v29
                }
                _ = v27
                _ = v28
            } else {
                var v30 uint32 = 2
                _ = v30
            }
        default:
            _ = len(v3)
        }
    }
    var v31 int32 = 928
    if true {
        delete(v1, 405)
        if false {
            var v32 uint = 681
            var v33 int16 = 314
            _ = v32
            _ = v33
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v1[-91] = "\n\t"
        v34 := GenS5{F0: 16, F1: -83, F2: -9223372036854775808}
        _ = v34
    }
    v35 := GenS1{F0: 42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v36 string = "foo bar"
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 int32 = v31
        if false {
            var v39 uint = 1
            _ = v39
        }
        _ = len(v3)
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}

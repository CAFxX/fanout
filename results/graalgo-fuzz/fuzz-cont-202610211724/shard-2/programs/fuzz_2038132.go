package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint32; F2 string }
func genFn0(p0 string, p1 uint64) uint64 {
    p0 = p0
    var v4 uint32 = 2
    var v5 float64 = (1e10 * 0.1)
    p1 = (p1 / (((^(p1))) | 1))
    switch 128 {
    case -459:
        var v6 int16 = 0
        _ = v6
    case 128:
        if true {
            var v7 float32 = (-2.25 * (0.1 + 1e10))
            v7 = v7
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p1 * p1)
}
func genFn1(p0 uint16, p1 int32, p2 uint) uint64 {
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8, p1, p0)
        var v9 int = fi8
        if false {
            var v10 uint = 0
            _ = v10
        } else {
            if false {
                var v11 int64 = -487
                _ = v11
            } else {
                p2 = (p2 << ((p2) & 63))
            }
            switch v9 {
            case 233:
                p0 += 42
            default:
                if true {
                    v12 := []int{fi8, (v9 + fi8), (v9 << ((18446744073709551615) & 63)), v9}
                    _ = v12
                }
            }
            fmt.Println(p0)
        }
        _ = fi8
        _ = v9
    }
    p2 = (p2 & p2)
    for fi13 := 0; fi13 < 19; fi13++ {
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            var v14 int = fi13
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 uint32 = 42
                var v17 uint8 = 0
                v18 := GenS1{F0: (^(fi15)), F1: uint(v17)}
                if false {
                    switch (fi13 % ((fi13) | 1)) {
                    case -857:
                        switch (fi13 << ((uint(p0)) & 63)) {
                        case 127:
                            var v19 uint32 = v16
                            _ = v19
                        case -2147483648:
                            v16 += v16
                        }
                    default:
                        v16 = (v16 | v16)
                    }
                    v20 := GenS1{F0: fi13, F1: (p2 | (p2 ^ (p2 + (^(p2)))))}
                    v21 := GenS2{F0: -128, F1: uint64(p0)}
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi13, p2)
                fmt.Println(v16, fi15, p0)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        case "a":
            v22 := GenS1{F0: 256, F1: p2}
            _ = v22
        }
        v23 := GenS2{F0: 1, F1: 42}
        var v24 int16 = 86
        _ = fi13
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 int8 = 1
    v25 := GenS2{F0: (-(v3)), F1: 18446744073709551615}
    for fi26 := 0; fi26 < 14; fi26++ {
        v1 += -214
        _ = fi26
    }
    var v27 uint = 42
    v1 -= (-(v1))
    v28 := GenS1{F0: 32767, F1: (v27 - v27)}
    v1 -= v1
    v29 := GenS3{F0: 42}
    var v30 int = -2147483648
    v31 := GenS4{F0: (true && true)}
    if false {
        var v32 float32 = ((-2.25 / (1.5 / (3.14159 / 0.1))) - 0.0)
        _ = v32
    } else {
        fmt.Println(v27)
        if (!false) {
            v33 := v2[386]
            v34 := GenS1{F0: (v30 | (v30 % (((v30 ^ (v30 >> (((v27 ^ (^(v27)))) & 63)))) | 1))), F1: (v27 ^ uint(v33))}
            _ = v33
            _ = v34
        }
        switch -1 {
        case 4294967295:
            var v35 string = ""
            _ = v35
        case 333:
            v27 += (v27 / ((v27) | 1))
        case 868:
            v36 := GenS3{F0: -1}
            _ = v36
        default:
            var v37 uint = (^(v27))
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        v39 := GenS2{F0: v3, F1: 0}
        v40 := []int{(fi38 - (v30 << ((v27) & 63)))}
        var v41 uint64 = 4294967295
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    switch "a" {
    case "foo bar":
        v42 := []int{9223372036854775807}
        _ = v42
    case "\n\t":
        v3 -= (v3 ^ v3)
    case "a":
        v1 = v1
    }
    v43 := GenS5{F0: 32767, F1: 42, F2: "a"}
    var v44 int64 = v1
    v30 -= (v30 % ((-335) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v43
    _ = v44
}

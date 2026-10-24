package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 string; F1 uint64; F2 bool }
type GenS5 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 uint, p1 int32) uint {
    p1 -= -677
    v4 := []int{0, -326, -9223372036854775808, 990}
    v5 := GenS1{F0: 0, F1: 2}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 int, p1 uint64, p2 string) bool {
    switch ((p2 + "hello") + "") {
    case "":
        if true {
            var v6 int64 = 345
            var v7 int = (p0 & p0)
            _ = v6
            _ = v7
        } else {
            switch p2 {
            case "":
                v8 := GenS2{F0: 65535}
                _ = v8
            case "hello":
                p0 = (p0 | 127)
            default:
                var v9 uint8 = 252
                _ = v9
            }
        }
    }
    var v10 uint32 = 2147483647
    var v11 uint64 = 1
    if false {
        var v12 uint8 = 171
        p0 = (^(p0))
        _ = v12
    } else {
        var v13 float64 = float64(float32((-2.25 - (1.5 * float64(3.14159)))))
        v11 = (v11 / ((p1) | 1))
        var v14 int = (p0 * p0)
        _ = v13
        _ = v14
    }
    v15 := []int{(p0 ^ (p0 >> ((2) & 63))), (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v15
    return true
}

func main() {
    var v1 int32 = -968
    var v2 uint64 = 0
    var v3 float64 = 0.0
    v16 := GenS2{F0: 4294967295}
    fmt.Println(v2, v2, v2)
    var v17 int8 = 42
    if (((("\n\t" + (("a" + "\n\t") + "a")) + "foo bar") + "a") == "\n\t") {
        v17 = (v17 % ((v17) | 1))
    } else {
        if false {
            for fi18 := 0; fi18 < 6; fi18++ {
                v1 = -2147483648
                if false {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 int64 = -1
                        var v21 uint8 = 152
                        v22 := []int{-306, fi19, (fi19 << ((2) & 63)), (fi18 - fi19)}
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS3{F0: 138, F1: 0, F2: 32768}
                    var v24 string = ("" + ("\n\t" + "hello"))
                    _ = v23
                    _ = v24
                }
                _ = fi18
            }
            var v25 float64 = 0.1
            _ = v25
        }
    }
    v1 += v1
    for fi26 := 0; fi26 < 3; fi26++ {
        v17 += v17
        _ = fi26
    }
    v27 := GenS4{F0: ("foo bar" + ""), F1: v2, F2: false}
    for fi28 := 0; fi28 < 9; fi28++ {
        switch (fi28 / ((fi28) | 1)) {
        case 127:
            switch (fi28 ^ fi28) {
            case 2147483647:
                v17 += -1
            case 273:
                var v29 int16 = 1
                _ = v29
            case 32767:
                fmt.Println(v1)
            }
        case 305:
            v30 := []int{}
            _ = v30
        }
        _ = fi28
    }
    var v31 uint64 = 497
    v32 := GenS5{F0: int(v17), F1: 946, F2: 266}
    if false {
        v2 -= 524
        var v33 int = 4294967295
        v2 = (v31 >> ((genFn0(531, v1)) & 63))
        _ = v33
    }
    fmt.Println(v1)
    switch int(v1) {
    case 9223372036854775807:
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 float64 = v3
            _ = fi34
            _ = v35
        }
    case -9223372036854775808:
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 int8 = 127
            _ = fi36
            _ = v37
        }
    case 2147483647:
        v2 = v2
    default:
        var v38 float64 = (0.0 / 1e10)
        _ = v38
    }
    switch 457 {
    case 32767:
        if false {
            v39 := []int{9223372036854775807, 128, 2147483647}
            _ = v39
        } else {
            v1 = (v1 - (^(v1)))
            v3 += -2.25
        }
    case 2147483647:
        v40 := []int{488, 894, 1, -344}
        _ = v40
    default:
        var v41 uint = 42
        _ = v41
    }
    v42 := GenS2{F0: 18446744073709551615}
    v43 := GenS1{F0: (v2 & (v31 & v31)), F1: (v2 & (v31 | (^(v31))))}
    v3 = v3
    v31 = (v2 | (v2 | (v31 >> ((uint(v2)) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v42
    _ = v43
}

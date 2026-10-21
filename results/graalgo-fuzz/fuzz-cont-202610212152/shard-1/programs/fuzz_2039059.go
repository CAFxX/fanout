package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
type GenS3 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint8, p1 int64) int {
    p1 -= (p1 ^ 1)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{}
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int64) int64 {
    p0 -= (p0 - (p0 - (p0 >> ((65535) & 63))))
    v6 := []int{-923, 32767, 0}
    if true {
        p0 = p0
        p0 = (p0 >> ((18446744073709551615) & 63))
    } else {
        var v7 int = 901
        _ = v7
    }
    p0 = (p0 + (p0 + (^(p0))))
    if true {
        p0 += 1
    }
    _ = p0
    _ = v6
    return (p0 & p0)
}

func main() {
    var v1 float32 = (0.1 * (1e10 / 1.5))
    var v2 uint8 = 255
    v3 := make(map[uint32]string)
    v8 := GenS1{F0: 0, F1: 42, F2: 210}
    var v9 int = 4294967295
    var v10 int = v9
    if true {
        v3[65535] = ("foo bar" + "\n\t")
        v11 := []int{v10}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint16 = 0
            v14 := GenS1{F0: 1, F1: 4294967295, F2: -1}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: 544, F1: 4294967295, F2: 42}
        switch ("\n\t" + "a") {
        case "foo bar":
            for fi16 := 0; fi16 < 18; fi16++ {
                switch 9223372036854775807 {
                case 378:
                    v17 := GenS2{F0: 1, F1: 18446744073709551615, F2: 2}
                    _ = v17
                }
                _ = fi16
            }
        case "\n\t":
            for fi18 := 0; fi18 < 12; fi18++ {
                v10 += genFn0(127, genFn1(1))
                for fi19 := 0; fi19 < 1; fi19++ {
                    v10 = (fi18 + v10)
                    _ = fi19
                }
                _ = fi18
            }
        case "a":
            if false {
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 int64 = -9223372036854775808
                    var v22 int8 = -128
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v9 += -107
                v9 = 340
            } else {
                fmt.Println(v2, v2, v1)
            }
        }
        _ = v15
    }
    v23 := GenS3{F0: 4294967295, F1: genFn1(9223372036854775807)}
    v2 -= 120
    switch "a" {
    case "\n\t":
        var v24 uint32 = 2147483648
        _ = v24
    default:
        var v25 int16 = 1000
        _ = v25
    }
    switch "" {
    case "\n\t":
        var v26 int = v9
        _ = v26
    case "foo bar":
        var v27 int64 = -1
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 bool = false
        v29 = false
        var v30 uint8 = v2
        _ = fi28
        _ = v29
        _ = v30
    }
    switch v9 {
    case 32767:
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS1{F0: 42, F1: 443, F2: -156}
            _ = fi31
            _ = v32
        }
    case -9223372036854775808:
        var v33 uint16 = 1
        _ = v33
    case 128:
        v3[65535] = "\n\t"
    default:
        v34 := GenS2{F0: 2147483647, F1: 4294967295, F2: 868}
        _ = v34
    }
    var v35 uint16 = 0
    var v36 uint16 = 503
    v35 = v36
    for fi37 := 0; fi37 < 12; fi37++ {
        fmt.Println(v10, v2, fi37)
        fmt.Println(v9, v10)
        var v38 float64 = (0.1 + 0.0)
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v8
    _ = v9
}

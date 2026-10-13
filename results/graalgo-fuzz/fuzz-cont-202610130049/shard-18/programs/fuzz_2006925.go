package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 bool, p2 int8) int32 {
    fmt.Println(p2, p1)
    fmt.Println(p1, p1)
    p1 = false
    switch "\n\t" {
    case "a":
        v4 := []int{}
        _ = v4
    case "foo bar":
        var v5 uint64 = 1
        _ = v5
    case "\n\t":
        p2 -= (p2 ^ (p2 >> ((42) & 63)))
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int16, p1 int32) bool {
    v6 := []int{127}
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v7, fi8)
        var v9 float64 = 0.0
        v6 = append(v6, (fi8 & fi8))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -89, F1: 1, F2: 2}
    v11 := GenS1{F0: 24, F1: 0, F2: 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    _ = v7
    return (!true)
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 4294967295
    var v3 float64 = (1e10 / 1e10)
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            v13 := []int{(^(fi12)), fi12, -2147483648, fi12}
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            if false {
                fmt.Println(v3, v3)
                v2 = (v2 / ((v2) | 1))
            } else {
                v2 -= 255
                fmt.Println(fi12, v3, v2)
            }
            v15 := GenS2{F0: -823, F1: (v2 << (((v2 + (v2 << ((v2) & 63)))) & 63)), F2: 395}
            _ = fi14
            _ = v15
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if false {
                v18 := GenS3{F0: "foo bar"}
                fmt.Println(fi17, v1)
                _ = v18
            } else {
                fmt.Println(fi17)
                v2 += v2
                v19 := []int{}
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if true {
                    v20 := GenS3{F0: ""}
                    _ = v20
                }
                _ = v19
            }
            v21 := []int{(fi17 | (fi16 + fi17))}
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    fmt.Println(v1)
    for fi22 := 0; fi22 < 18; fi22++ {
        v1 = v1
        v23 := []int{fi22, fi22}
        _ = fi22
        _ = v23
    }
    var v24 bool = (!false)
    var v25 uint8 = 255
    for fi26 := 0; fi26 < 12; fi26++ {
        v25 -= v25
        v27 := GenS2{F0: 9223372036854775807, F1: 2, F2: 2}
        var v28 uint = v2
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint32 = 920
        _ = fi29
        _ = v30
    }
    v2 = v2
    var v31 int16 = 1
    fmt.Println(v2)
    v32 := GenS3{F0: (("foo bar" + "") + "a")}
    fmt.Println(v24, v31)
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := GenS2{F0: 42, F1: uint(v1), F2: 1}
        v35 := GenS1{F0: -128, F1: 42, F2: 18446744073709551615}
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v36
}

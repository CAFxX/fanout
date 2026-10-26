package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) uint16 {
    p0 = 0
    v4 := GenS1{F0: 32767, F1: "a", F2: 1}
    var v5 float64 = (-2.25 + -2.25)
    fmt.Println(v5)
    var v6 uint16 = 842
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (^(v6))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    v3 := make(map[int]string)
    var v7 float32 = 0.0
    if ((((0.1 + (3.14159 * 1e10)) * 0.1) * 0.0) < 1e10) {
        v7 = 0.1
        v7 -= ((v7 / 1.0) * (((v7 / (v7 - ((1.5 * v7) - (v7 / v7)))) * float32(1e10)) * (v7 + v7)))
    } else {
        v7 = float32(float64((float32(0.0) / (((0.0 * v7) * v7) + v7))))
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v7 = float32((3.14159 - 0.0))
        for fi9 := 0; fi9 < 11; fi9++ {
            if (!(480 == 2147483648)) {
                fmt.Println(v7)
                fmt.Println(v2, fi8, fi8)
                v10 := v1[4294967295]
                var v11 uint32 = 733
                _ = v10
                _ = v11
            }
            var v12 int16 = 42
            if false {
                v7 -= (v7 - v7)
            }
            _ = fi9
            _ = v12
        }
        v13 := GenS1{F0: 65535, F1: "\n\t", F2: 1}
        _ = fi8
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 float64 = 1.5
        _ = fi14
        _ = v15
    }
    v7 = v7
    v16 := GenS1{F0: 0, F1: "hello", F2: -334}
    v17 := GenS2{F0: -1, F1: -139}
    switch "foo bar" {
    case "foo bar":
        var v18 uint16 = 42
        _ = v18
    default:
        var v19 uint = 4294967295
        _ = v19
    }
    switch ("\n\t" + (("" + "hello") + "a")) {
    case "\n\t":
        for fi20 := 0; fi20 < 4; fi20++ {
            switch "foo bar" {
            case "foo bar":
                v2 = v2
            case "\n\t":
                var v21 uint = 356
                _ = v21
            }
            _ = fi20
        }
    default:
        fmt.Println(v2, v2)
    }
    v7 += v7
    var v22 uint16 = 65535
    if ((!(215 != 13)) && (100 <= 127)) {
        if v2 {
            v1[217] = (true || ((("" + "") + "foo bar") > "foo bar"))
            switch ("foo bar" + "a") {
            case "":
                v22 = genFn0(42)
            case "hello":
                if (1 > 0) {
                    v7 += (v7 - 1.5)
                }
            }
            v23 := []int{-9223372036854775808, -499}
            _ = v23
        }
        if false {
            switch -2147483648 {
            case 662:
                switch "a" {
                case "":
                    v2 = (v2 && v2)
                }
            }
            v24 := GenS3{F0: 2147483647}
            _ = v24
        }
    } else {
        if true {
            v25 := v3[42]
            _ = v25
        } else {
            if (65535 <= 746) {
                fmt.Println(v2, v7)
            } else {
                v26 := GenS3{F0: 1}
                v27 := []int{-266, 9223372036854775807}
                _ = v26
                _ = v27
            }
            v28 := []int{-9223372036854775808}
            _ = v28
        }
    }
    var v29 int8 = 42
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := []int{}
        _ = fi30
        _ = v31
    }
    v22 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v7
}

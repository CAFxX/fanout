package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16, p1 int8) int8 {
    switch 9223372036854775807 {
    case -9223372036854775808:
        v7 := []int{-736, 888, 128, -932}
        _ = v7
    }
    p0 = (^(p0))
    if false {
        p1 += p1
        switch "foo bar" {
        case "hello":
            p0 = 153
        case "foo bar":
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: false, F1: "a"}
                v10 := []int{(fi8 & 128), (fi8 % ((int(p0)) | 1))}
                if true {
                    p1 = p1
                } else {
                    fmt.Println(p0, fi8)
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    }
    p0 += (p0 % ((-648) | 1))
    if false {
        for fi11 := 0; fi11 < 7; fi11++ {
            p1 -= p1
            var v12 uint32 = 4294967295
            _ = fi11
            _ = v12
        }
    } else {
        fmt.Println(p0, p1)
        if false {
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(fi13, p1)
                var v14 int16 = p0
                fmt.Println(v14)
                _ = fi13
                _ = v14
            }
            if (42 >= 128) {
                v15 := []int{}
                fmt.Println(p0)
                var v16 int16 = p0
                _ = v15
                _ = v16
            }
        }
    }
    _ = p0
    _ = p1
    return (p1 / ((p1) | 1))
}
func genFn1(p0 uint, p1 int32, p2 uint16) int64 {
    var v17 uint64 = 0
    if false {
        v18 := []int{128, 150, -835}
        fmt.Println(p1, v17)
        _ = v18
    }
    var v19 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v19
    return int64(p1)
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    var v3 int = -2147483648
    var v4 uint64 = 42
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    v20 := GenS2{F0: "\n\t", F1: true, F2: 133}
    fmt.Println(v3, v4, v3)
    v21 := GenS1{F0: (uint16(v4) >= 754), F1: "hello"}
    if (0 < 45) {
        v22 := v5[v3]
        _ = v22
    } else {
        fmt.Println(v3, v2)
        var v23 int32 = int32(v4)
        _ = v23
    }
    var v24 int32 = -2147483648
    fmt.Println(v24)
    fmt.Println(v24, v3)
    v2 -= (v2 | (^(v2)))
    switch int(v2) {
    case -1:
        if (!(626 < 42)) {
            v25 := []int{2147483647}
            v4 -= 249
            _ = v25
        } else {
            var v26 int = 0
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 uint = 18446744073709551615
                _ = fi27
                _ = v28
            }
            switch "a" {
            case "":
                var v29 bool = (true || false)
                _ = v29
            case "hello":
                for fi30 := 0; fi30 < 9; fi30++ {
                    fmt.Println(v3, v24)
                    v31 := GenS1{F0: false, F1: "foo bar"}
                    _ = fi30
                    _ = v31
                }
            case "\n\t":
                v4 -= v4
            default:
                v3 += (^(v3))
            }
            _ = v26
        }
    case 32767:
        v32 := GenS3{F0: 370}
        _ = v32
    }
    var v33 float64 = (float64(float32(3.14159)) * -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 string, p1 bool, p2 uint8) int64 {
    switch p0 {
    case "a":
        var v7 int16 = 665
        _ = v7
    default:
        for fi8 := 0; fi8 < 10; fi8++ {
            p1 = true
            fmt.Println(p1, fi8)
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if (1e10 < 0.1) {
            v10 := []int{(fi9 % (((fi9 << ((255) & 63))) | 1)), 288, fi9, fi9}
            _ = v10
        } else {
            var v11 int32 = 2147483647
            _ = v11
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            if ((p0 >= (p0 + p0)) && p1) {
                v13 := []int{}
                var v14 int64 = 42
                v15 := GenS1{F0: 2147483647}
                v16 := GenS2{F0: 186}
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            switch "hello" {
            case "":
                var v17 int8 = 127
                _ = v17
            }
            v18 := []int{fi9, -933}
            _ = fi12
            _ = v18
        }
        p1 = ((int64(p2) >= -1) && (!(850 < 4294967295)))
        _ = fi9
    }
    v19 := []int{127, 42, 9223372036854775807, 128}
    v20 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 int32) int8 {
    fmt.Println(p1, p1)
    p1 -= p1
    p0 = 0
    if false {
        fmt.Println(p1, p0)
        v21 := []int{-1, 4294967295}
        _ = v21
    } else {
        if true {
            p1 -= p1
            v22 := []int{-2147483648, 363}
            v23 := GenS3{F0: p0}
            _ = v22
            _ = v23
        }
        var v24 uint8 = 0
        if (0 != 1) {
            switch "a" {
            case "":
                v24 -= (v24 / (((^(v24))) | 1))
            case "hello":
                v25 := []int{127, 32768}
                _ = v25
            }
        }
        _ = v24
    }
    p0 += 1
    _ = p0
    _ = p1
    return 100
}

func main() {
    var v1 string = "hello"
    var v2 int32 = 0
    var v3 string = v1
    var v4 float32 = (float32(1e10) / 1.5)
    var v5 uint8 = 255
    v6 := make(map[string]bool)
    var v26 int32 = -698
    if true {
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := GenS4{F0: 20}
                _ = fi28
                _ = v29
            }
            fmt.Println(fi27, v1, v5)
            _ = fi27
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            v26 = v2
            _ = fi30
        }
        var v31 int8 = genFn1(4294967295, v2)
        _ = v31
    } else {
        v32 := GenS5{F0: 1, F1: 0, F2: true}
        v33 := GenS6{F0: -2147483648}
        v34 := GenS7{F0: 0}
        _ = v32
        _ = v33
        _ = v34
    }
    v35 := GenS8{F0: 42}
    switch v3 {
    case "a":
        v36 := []int{}
        _ = v36
    default:
        var v37 uint = 18446744073709551615
        _ = v37
    }
    var v38 int16 = -405
    for fi39 := 0; fi39 < 10; fi39++ {
        fmt.Println(v3)
        if true {
            var v40 int64 = 595
            _ = v40
        }
        var v41 int64 = -9223372036854775808
        _ = fi39
        _ = v41
    }
    delete(v6, (v1 + v1))
    v26 = v2
    if false {
        v3 = ((v3 + (v3 + (v1 + v3))) + (v3 + v3))
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        if false {
            for fi43 := 0; fi43 < 1; fi43++ {
                v44 := []int{(fi43 | 32767), (fi42 << ((106) & 63)), (fi43 - (fi43 ^ (fi43 ^ fi42)))}
                v26 = (-(v2))
                v45 := GenS9{F0: 2}
                _ = fi43
                _ = v44
                _ = v45
            }
            if (4294967295 <= 18446744073709551615) {
                v2 = 814
            } else {
                fmt.Println(v4, v2)
                fmt.Println(v3, v4, v2)
            }
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

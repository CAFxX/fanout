package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint; F1 uint8; F2 int }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0(p0 int64, p1 int) uint64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 -= (-(p0))
        _ = fi5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int16 = 419
        v7 = -534
        for fi8 := 0; fi8 < 1; fi8++ {
            switch "a" {
            case "foo bar":
                var v9 uint32 = 4294967295
                _ = v9
            case "":
                p1 += -392
            }
            p0 = 156
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int, p1 int, p2 uint32) int64 {
    var v10 int64 = -1
    v11 := GenS1{F0: true}
    p0 -= 5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return v10
}

func main() {
    var v1 int = 32768
    var v2 int16 = -1
    var v3 uint = 2
    var v4 float64 = 3.14159
    var v12 int32 = 50
    v13 := GenS2{F0: true, F1: (v1 & -2147483648)}
    if true {
        var v14 uint32 = 65535
        var v15 float64 = v4
        _ = v14
        _ = v15
    }
    var v16 int64 = 1
    var v17 float32 = ((float32(v4) / 3.14159) / -2.25)
    var v18 uint16 = 365
    var v19 float64 = (v4 + float64(v17))
    v19 += 1e10
    v3 += 2
    if false {
        v1 = 1
    }
    v20 := GenS1{F0: (false && true)}
    v21 := []int{v1, 86, (v1 >> ((1) & 63))}
    var v22 uint64 = 42
    if (1 > 0) {
        fmt.Println(v3, v17, v18)
        v23 := GenS3{F0: true, F1: ((("foo bar" + "a") + "\n\t") + "hello")}
        var v24 int64 = (^(v16))
        _ = v23
        _ = v24
    } else {
        switch "" {
        case "hello":
            fmt.Println(v12, v3, v2)
        case "":
            if true {
                fmt.Println(v22, v17)
            } else {
                var v25 int8 = -68
                switch v1 {
                case 2147483647:
                    v26 := GenS2{F0: ((^(v3)) > v3), F1: v1}
                    _ = v26
                }
                _ = v25
            }
        }
        switch (v1 * v1) {
        case 128:
            switch "hello" {
            case "foo bar":
                v27 := GenS3{F0: true, F1: "a"}
                _ = v27
            case "\n\t":
                switch "a" {
                case "a":
                    var v28 int16 = int16(v1)
                    _ = v28
                case "foo bar":
                    if true {
                        v29 := GenS4{F0: (v3 - 0), F1: 142, F2: int(v2)}
                        if false {
                            v12 -= (v12 + -683)
                            var v30 int32 = 1
                            _ = v30
                        }
                        _ = v29
                    }
                }
            case "hello":
                if (-1 > -21) {
                    switch "\n\t" {
                    case "\n\t":
                        v31 := GenS5{F0: 42, F1: v3}
                        _ = v31
                    case "hello":
                        var v32 int16 = (v2 << (((v3 << ((v3) & 63))) & 63))
                        _ = v32
                    default:
                        var v33 uint8 = 42
                        _ = v33
                    }
                } else {
                    if false {
                        if true {
                            var v34 int = v1
                            _ = v34
                        }
                    } else {
                        v35 := GenS5{F0: (^(v18)), F1: uint(v18)}
                        var v36 int32 = v12
                        var v37 float32 = v17
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    if true {
                        v19 -= v4
                    }
                    var v38 uint = 408
                    _ = v38
                }
            }
        case 42:
            switch "foo bar" {
            case "foo bar":
                v16 = (v16 >> ((42) & 63))
            case "a":
                fmt.Println(v4, v3, v17)
            default:
                fmt.Println(v1)
            }
        case 0:
            v39 := GenS5{F0: v18, F1: v3}
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}

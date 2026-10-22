package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0() bool {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint32 = 0
        _ = fi5
        _ = v6
    }
    switch "foo bar" {
    case "foo bar":
        var v7 float32 = (1.5 * 1.5)
        _ = v7
    case "":
        if ("hello" <= "\n\t") {
            var v8 string = "\n\t"
            _ = v8
        }
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            var v9 float64 = 0.0
            _ = v9
        default:
            v10 := []int{}
            _ = v10
        }
    default:
        v11 := []int{-1, 811, 32768}
        _ = v11
    }
    var v12 uint16 = 986
    fmt.Println(v12)
    v12 -= 785
    _ = v12
    return false
}
func genFn1(p0 int64, p1 string, p2 int) int32 {
    var v13 float64 = 1.5
    v14 := []int{}
    if len(v14) > 0 {
        v14[0] = 4294967295
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        switch (fi15 >> ((500) & 63)) {
        case -743:
            switch p2 {
            case 476:
                v16 := GenS1{F0: 1}
                _ = v16
            case -9223372036854775808:
                p1 = "a"
            case 578:
                if ((42 >= 230) && true) {
                    if true {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            p2 -= -1
                            _ = fi17
                        }
                        fmt.Println(fi15)
                        var v18 uint8 = 0
                        _ = v18
                    }
                }
            default:
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 string = ""
                    v21 := GenS1{F0: 4294967295}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        case 974:
            var v22 uint64 = 0
            _ = v22
        default:
            p0 += 1
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 954
}

func main() {
    var v1 float32 = (1e10 + 1.5)
    v2 := make(map[int]int)
    var v3 int64 = 0
    v4 := []uint32{2147483647, 0}
    switch "foo bar" {
    case "\n\t":
        v23 := GenS1{F0: 1}
        _ = v23
    default:
        fmt.Println(v3)
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        switch 128 {
        case 128:
            fmt.Println(fi24)
        }
        var v25 bool = false
        var v26 int8 = -99
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: 1, F1: 4294967295}
    var v28 int32 = 0
    for fi29 := 0; fi29 < 19; fi29++ {
        v1 += 1.5
        _ = fi29
    }
    fmt.Println(v1)
    if false {
        var v30 float32 = v1
        switch "hello" {
        case "":
            fmt.Println(v3)
        }
        _ = v30
    } else {
        if ((v28 * -1) <= genFn1(v3, "hello", 1)) {
            v31 := GenS2{F0: 0, F1: 114}
            var v32 uint = 42
            if false {
                if false {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        fmt.Println(v1)
                        _ = fi33
                    }
                }
            }
            _ = v31
            _ = v32
        }
        var v34 string = ""
        for fi35 := 0; fi35 < 8; fi35++ {
            v1 = (1e10 + float32(-2.25))
            var v36 string = v34
            for fi37 := 0; fi37 < 1; fi37++ {
                v36 = v36
                _ = fi37
            }
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    var v38 int8 = 1
    v3 = (v3 / ((v3) | 1))
    v28 -= (v28 & (v28 % (((v28 % ((1) | 1))) | 1)))
    fmt.Println(v28, v38, v3)
    v1 -= v1
    var v39 uint8 = 42
    v39 = (v39 + v39)
    v1 -= v1
    v39 += 203
    if true {
        var v40 string = "foo bar"
        _ = v40
    } else {
        fmt.Println(v28, v3)
        if len(v4) > 0 {
            v4[0] = 615
        }
    }
    v41 := GenS1{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
}

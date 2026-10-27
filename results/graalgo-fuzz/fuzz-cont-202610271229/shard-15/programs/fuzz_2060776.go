package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
func genFn0() uint64 {
    var v5 int16 = 939
    if false {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint16 = 0
            v7 += 871
            _ = fi6
            _ = v7
        }
        var v8 int8 = -1
        _ = v8
    }
    v9 := GenS1{F0: true, F1: 255}
    for fi10 := 0; fi10 < 17; fi10++ {
        if (false && false) {
            var v11 string = "\n\t"
            v12 := GenS1{F0: true, F1: 1}
            fmt.Println(v11)
            _ = v11
            _ = v12
        }
        if true {
            if false {
                var v13 int8 = -128
                var v14 float32 = 0.0
                _ = v13
                _ = v14
            }
            var v15 bool = (true && (!(275 == 4294967295)))
            _ = v15
        } else {
            v16 := GenS1{F0: true, F1: 42}
            _ = v16
        }
        _ = fi10
    }
    var v17 string = "foo bar"
    _ = v17
    _ = v5
    _ = v9
    return 4294967295
}
func genFn1(p0 bool, p1 string) uint8 {
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(p0)
            _ = fi19
        }
        v20 := GenS2{F0: 12, F1: -128}
        fmt.Println(fi18, fi18)
        _ = fi18
        _ = v20
    }
    p1 = p1
    p0 = true
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 float64 = 1e10
        switch fi21 {
        case 941:
            switch "\n\t" {
            case "a":
                p1 = "hello"
            case "\n\t":
                var v23 uint64 = 365
                _ = v23
            case "hello":
                if (-673 > 0) {
                    var v24 float32 = -2.25
                    _ = v24
                }
            }
        default:
            v22 = (-2.25 - float64(1e10))
        }
        if p0 {
            if p0 {
                if (true || p0) {
                    var v25 uint8 = 1
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v22 = v22
                        _ = fi26
                    }
                }
                v27 := []int{fi21, fi21, 2147483647}
                var v28 uint64 = 4294967295
                _ = v27
                _ = v28
            } else {
                var v29 float32 = ((-2.25 - (-2.25 / 1.0)) / 1e10)
                for fi30 := 0; fi30 < 5; fi30++ {
                    v29 -= ((((float32((float64((v29 + 0.1)) / v22)) * (v29 - float32(v22))) - 0.1) * 3.14159) * ((v29 - v29) * 0.1))
                    switch fi30 {
                    case 0:
                        v31 := []int{(fi30 % ((fi30) | 1)), fi30, fi21}
                        _ = v31
                    case -9223372036854775808:
                        p1 = "hello"
                    }
                    fmt.Println(v22, p1, fi21)
                    var v32 int32 = 0
                    _ = fi30
                    _ = v32
                }
                v29 -= 1e10
                _ = v29
            }
            switch "\n\t" {
            case "foo bar":
                for fi33 := 0; fi33 < 11; fi33++ {
                    v34 := GenS1{F0: (!((p1 + ("a" + p1)) >= p1)), F1: 42}
                    var v35 float64 = v22
                    v35 -= v35
                    v36 := GenS2{F0: 87, F1: 15}
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            case "":
                var v37 int16 = 1000
                _ = v37
            default:
                v22 += (1e10 + (0.0 + (((v22 - (1.5 - v22)) / ((v22 - (v22 - v22)) + v22)) + v22)))
            }
        } else {
            for fi38 := 0; fi38 < 10; fi38++ {
                if (true || (p0 && false)) {
                    v39 := []int{}
                    fmt.Println(p1, fi38)
                    _ = v39
                }
                v40 := []int{fi38, fi38}
                var v41 uint8 = 0
                _ = fi38
                _ = v40
                _ = v41
            }
            v22 -= float64(0.0)
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS3{F0: p0}
                _ = fi42
                _ = v43
            }
        }
        _ = fi21
        _ = v22
    }
    p1 = (p1 + (p1 + p1))
    _ = p0
    _ = p1
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 281
    v3 := make(map[int]bool)
    var v4 string = ("" + ("\n\t" + ("hello" + "hello")))
    var v44 float32 = -2.25
    v45 := GenS4{F0: -2147483648}
    var v46 uint32 = 0
    if ((!(2 != 240)) || false) {
        v47 := []int{-2147483648, 381, 1, 9223372036854775807}
        var v48 uint8 = uint8(v46)
        v49 := GenS2{F0: uint16(v48), F1: 0}
        _ = v47
        _ = v48
        _ = v49
    }
    v44 = v44
    var v50 float32 = v44
    switch v4 {
    case "\n\t":
        v46 = (v46 << ((518) & 63))
    }
    fmt.Println(v44, v2, v4)
    var v51 int16 = 607
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v50
    _ = v51
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string; F1 int64; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 string }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float64 = 1.5
        _ = fi6
        _ = v7
    }
    return 246
}
func genFn1(p0 uint64, p1 uint64) int8 {
    switch -2147483648 {
    case 4294967295:
        if (!true) {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 bool = (!((0.0 - 3.14159) == (1.5 / 1.0)))
                var v10 bool = ((v9 || true) || true)
                var v11 float32 = ((float32(((1.5 / float64(1e10)) + 3.14159)) / ((1.5 / 0.1) * -2.25)) + 3.14159)
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
    case 32767:
        var v12 uint = 255
        _ = v12
    case -2147483648:
        p1 += (p0 | p0)
    }
    if (!((true && true) && false)) {
        if false {
            switch 2147483647 {
            case 767:
                v13 := GenS1{F0: -111}
                _ = v13
            case -906:
                p1 = p1
            default:
                var v14 uint32 = 2
                _ = v14
            }
            var v15 int = 32768
            _ = v15
        } else {
            fmt.Println(p0, p1)
            p0 -= 18446744073709551615
            v16 := []int{}
            _ = v16
        }
        var v17 int8 = 100
        _ = v17
    }
    fmt.Println(p1, p0)
    if false {
        if true {
            v18 := []int{-193, 4294967295, 32768, 0}
            _ = v18
        }
        var v19 bool = true
        if ((^(p1)) <= p1) {
            v20 := GenS2{F0: 42, F1: -1}
            var v21 uint = 65535
            if false {
                v22 := GenS1{F0: -1}
                var v23 uint = v21
                v24 := []int{1, 793}
                _ = v22
                _ = v23
                _ = v24
            }
            _ = v20
            _ = v21
        }
        _ = v19
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 uint64 = 42
    v4 := make(map[string]int)
    v3 = (v3 % ((v3) | 1))
    switch 658 {
    case 42:
        if true {
            v3 -= v3
            fmt.Println(v3)
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(fi25)
                _ = fi25
            }
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        if false {
            var v27 uint = 2
            switch "\n\t" {
            case "foo bar":
                var v28 uint32 = 805
                _ = v28
            case "":
                if ((0.1 / -2.25) > (1e10 / 0.1)) {
                    if ((1 != 0) && false) {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 uint64 = (v3 | (v3 << (((^(v27))) & 63)))
                            _ = fi29
                            _ = v30
                        }
                    }
                } else {
                    v3 += (v3 % (((v3 ^ v3)) | 1))
                    v31 := GenS3{F0: (fi26 + -692), F1: "foo bar"}
                    _ = v31
                }
            default:
                switch "foo bar" {
                case "\n\t":
                    var v32 string = ("hello" + "\n\t")
                    _ = v32
                case "a":
                    v33 := []int{791}
                    _ = v33
                }
            }
            _ = v27
        } else {
            var v34 float64 = 1.5
            v3 = (v3 ^ (^(v3)))
            v3 += v3
            _ = v34
        }
        _ = len(v4)
        _ = fi26
    }
    v35 := GenS4{F0: "hello", F1: -9223372036854775808, F2: 566}
    v36 := GenS2{F0: 65535, F1: -2147483648}
    for fi37 := 0; fi37 < 16; fi37++ {
        switch "hello" {
        case "foo bar":
            v38 := GenS1{F0: 88}
            _ = v38
        case "a":
            var v39 float32 = ((1.5 / 1.5) * 1e10)
            _ = v39
        case "hello":
            v3 = v3
        default:
            v40 := GenS5{F0: ((("" + "foo bar") + "a") + "foo bar")}
            _ = v40
        }
        if ((1e10 >= 0.1) && (-1 > -1)) {
            var v41 int64 = -10
            _ = v41
        }
        v3 += v3
        _ = fi37
    }
    fmt.Println(v3)
    v42 := []int{}
    var v43 int16 = -1
    v3 = v3
    if false {
        fmt.Println(v3, v43)
    } else {
        fmt.Println(v3, v3)
        v44 := GenS4{F0: "a", F1: -9223372036854775808, F2: 0}
        _ = v44
    }
    _ = len(v42)
    v43 -= 42
    var v45 bool = false
    v46 := GenS6{F0: 32767, F1: "foo bar"}
    v47 := GenS4{F0: ("" + "a"), F1: 0, F2: 65535}
    v43 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v47
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64; F2 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int16, p1 string, p2 uint32) uint16 {
    switch p1 {
    case "a":
        fmt.Println(p0, p1)
    case "foo bar":
        switch 128 {
        case 42:
            v5 := []int{341}
            _ = v5
        case 32767:
            if false {
                p2 -= (p2 & p2)
                switch p1 {
                case "":
                    var v6 string = ("a" + (p1 + "hello"))
                    _ = v6
                }
                if ((9223372036854775807 != -95) && false) {
                    var v7 int32 = -656
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        var v9 int8 = 127
                        var v10 int64 = -9223372036854775808
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    switch 42 {
                    case 42:
                        p0 = p0
                    case 681:
                        p1 = ((p1 + (p1 + (p1 + ""))) + "")
                    case 726:
                        var v11 int16 = -32768
                        _ = v11
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(p0, p0, p2)
                        }
                    }
                }
            }
        case -817:
            var v12 float32 = (3.14159 * 1e10)
            _ = v12
        default:
            var v13 float64 = (1.5 - 1.5)
            _ = v13
        }
    case "":
        v14 := GenS1{F0: (!(2 != 42))}
        _ = v14
    default:
        fmt.Println(p2, p0, p1)
    }
    v15 := GenS2{F0: 18446744073709551615, F1: 153, F2: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 633
}

func main() {
    var v1 uint8 = 115
    v2 := make(map[uint32]int)
    var v3 uint32 = 65535
    v4 := make(map[string]string)
    for fi16 := 0; fi16 < 9; fi16++ {
        if (!(uint32(v1) < v3)) {
            var v17 int64 = 0
            v18 := GenS1{F0: true}
            fmt.Println(v17, v17)
            _ = v17
            _ = v18
        }
        if ((v1 % ((v1) | 1)) < 1) {
            var v19 uint = 255
            _ = v19
        }
        if (65535 > 142) {
            v20 := []int{(fi16 + fi16), fi16, 128, -440}
            v1 += (^(v1))
            _ = v20
        } else {
            fmt.Println(v3, v1, v1)
        }
        _ = fi16
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := []int{(fi21 << ((226) & 63)), (fi21 >> ((65535) & 63)), fi22, -2147483648}
            v24 := GenS2{F0: 212, F1: 1, F2: "a"}
            v3 += uint32(v1)
            _ = fi22
            _ = v23
            _ = v24
        }
        switch fi21 {
        case 182:
            v1 -= v1
        }
        _ = fi21
    }
    switch ("" + "\n\t") {
    case "foo bar":
        v25 := []int{0, 4294967295}
        _ = v25
    default:
        v26 := GenS3{F0: 32767, F1: 2147483647}
        _ = v26
    }
    v1 -= 248
    v27 := []int{934, 0, -179, 32767}
    v28 := GenS1{F0: true}
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int32 = 666
        var v31 int64 = 955
        for fi32 := 0; fi32 < 10; fi32++ {
            v31 = 0
            _ = fi32
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    v33 := GenS3{F0: 1000, F1: -1}
    v34 := GenS4{F0: 185, F1: 1}
    var v35 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0(p0 int64, p1 int16) uint32 {
    var v5 uint8 = 128
    v6 := GenS1{F0: v5, F1: (^(v5))}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16, p1 int32) int64 {
    switch "foo bar" {
    case "hello":
        if (!true) {
            v7 := GenS1{F0: 127, F1: 128}
            var v8 uint16 = (p0 % (((^(p0))) | 1))
            _ = v7
            _ = v8
        }
    }
    if ((false && (-9223372036854775808 != 997)) && false) {
        if true {
            var v9 uint32 = genFn0(0, -1)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int8 = -128
                var v12 uint16 = p0
                _ = fi10
                _ = v11
                _ = v12
            }
            if false {
                v13 := []int{-294, -1, 127}
                _ = v13
            }
        }
        v14 := GenS2{F0: "foo bar", F1: 42}
        p1 += (p1 % (((p1 >> ((79) & 63))) | 1))
        _ = v14
    } else {
        if (true && false) {
            if (42 == 255) {
                var v15 uint32 = 980
                _ = v15
            }
        }
        var v16 int16 = 32767
        _ = v16
    }
    switch 32767 {
    case -365:
        var v17 bool = (!(-2.25 == 3.14159))
        _ = v17
    default:
        v18 := GenS1{F0: 128, F1: 128}
        _ = v18
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.0
    var v3 float64 = (0.1 + 1.5)
    v4 := make(map[int]int)
    delete(v4, 898)
    fmt.Println(v2)
    for fi19 := 0; fi19 < 1; fi19++ {
        v2 += v2
        var v20 uint = 8
        v21 := []int{fi19, 32767}
        _ = fi19
        _ = v20
        _ = v21
    }
    v2 -= (0.1 + v2)
    if true {
        var v22 int8 = 42
        if true {
            v2 -= (0.1 + 1.5)
            var v23 string = "hello"
            _ = v23
        } else {
            switch ("" + "a") {
            case "hello":
                v24 := GenS2{F0: "\n\t", F1: 128}
                _ = v24
            case "":
                var v25 int8 = (^(v22))
                _ = v25
            }
            switch -622 {
            case 258:
                if (32767 > 2) {
                    v3 = 0.0
                    if true {
                        v26 := GenS3{F0: "a", F1: 65535, F2: 42}
                        var v27 uint16 = 42
                        _ = v26
                        _ = v27
                    }
                    var v28 int16 = 1
                    _ = v28
                } else {
                    v22 -= (v22 - v22)
                    if ((-2.25 + 1e10) != 0.1) {
                        v29 := GenS3{F0: "foo bar", F1: genFn0(-998, -1), F2: 0}
                        v3 = v3
                        v30 := GenS3{F0: "a", F1: 65535, F2: 398}
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS1{F0: 223, F1: 2}
                        fmt.Println(v3, v22, v22)
                        _ = v31
                    }
                }
            case 0:
                for fi32 := 0; fi32 < 6; fi32++ {
                    v3 -= ((float64((-2.25 / (v2 / v2))) - (v3 / 1.5)) - 1.5)
                    _ = fi32
                }
            case 32767:
                v33 := []int{0, -1, -115, 0}
                _ = v33
            default:
                fmt.Println(v22, v3, v22)
            }
            var v34 string = "foo bar"
            _ = v34
        }
        fmt.Println(v2)
        _ = v22
    } else {
        if (42 == 687) {
            _ = len(v4)
            var v35 string = "hello"
            for fi36 := 0; fi36 < 19; fi36++ {
                v35 = ((((v35 + v35) + v35) + (v35 + v35)) + "hello")
                v35 = (((v35 + "foo bar") + v35) + "\n\t")
                _ = fi36
            }
            _ = v35
        }
        switch "hello" {
        case "":
            var v37 int32 = -537
            _ = v37
        case "\n\t":
            _ = len(v4)
        default:
            var v38 int16 = -32768
            _ = v38
        }
        v2 = (v2 + 3.14159)
    }
    v2 = (v2 - v2)
    v39 := GenS3{F0: ("hello" + "foo bar"), F1: genFn0(1, -32768), F2: 2147483647}
    v40 := []int{4294967295, 42, 9223372036854775807, 319}
    var v41 uint8 = 128
    for fi42 := 0; fi42 < 7; fi42++ {
        v43 := v1["foo bar"]
        _ = fi42
        _ = v43
    }
    v4[692] = -9223372036854775808
    fmt.Println(v2, v2)
    _ = len(v4)
    var v44 uint32 = 149
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
}

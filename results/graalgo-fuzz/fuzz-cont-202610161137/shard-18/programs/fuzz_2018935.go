package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint16) int16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            if true {
                v7 := []int{fi6, fi6}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 4; fi8++ {
                    fmt.Println(p0, fi8)
                    var v9 int64 = -9223372036854775808
                    _ = fi8
                    _ = v9
                }
                fmt.Println(p0, p0)
                p0 += p0
            }
        } else {
            switch "foo bar" {
            case "a":
                v10 := []int{(fi6 + fi6)}
                _ = v10
            }
        }
        p0 -= p0
        _ = fi6
    }
    v11 := []int{-9223372036854775808, -1}
    var v12 float32 = (-2.25 * 0.0)
    for fi13 := 0; fi13 < 7; fi13++ {
        p0 = 32767
        switch "" {
        case "a":
            v14 := GenS1{F0: 94, F1: 0}
            _ = v14
        case "hello":
            v15 := GenS2{F0: 401, F1: fi13}
            _ = v15
        }
        _ = fi13
    }
    switch "hello" {
    case "hello":
        fmt.Println(v12, v12)
    case "":
        v16 := GenS3{F0: "", F1: 9223372036854775807}
        _ = v16
    default:
        for fi17 := 0; fi17 < 4; fi17++ {
            v12 += (3.14159 + 0.1)
            fmt.Println(v12, fi17, fi17)
            _ = fi17
        }
    }
    _ = p0
    _ = v11
    _ = v12
    return -1
}
func genFn1(p0 int8) string {
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS4{F0: (((((("" + "foo bar") + "hello") + "a") + "a") + "hello") + "\n\t")}
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: -34, F1: 955}
    var v21 int64 = -1
    var v22 int32 = -92
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    return "hello"
}

func main() {
    var v1 uint = 0
    var v2 uint = (^(v1))
    v3 := []int{2147483647, 846, -1}
    var v4 uint16 = 355
    var v5 int32 = 726
    v1 -= 901
    var v23 bool = true
    v5 = (v5 % ((-2147483648) | 1))
    var v24 uint16 = 491
    v3 = append(v3, 127)
    var v25 float64 = 1e10
    for fi26 := 0; fi26 < 6; fi26++ {
        if (!true) {
            if ("foo bar" >= ("" + "\n\t")) {
                if ("\n\t" < "foo bar") {
                    var v27 float32 = 0.1
                    _ = v27
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 int8 = 0
                    var v30 bool = ("a" < "hello")
                    v2 = 261
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v24 = v24
            } else {
                for fi31 := 0; fi31 < 14; fi31++ {
                    fmt.Println(fi31)
                    v23 = v23
                    v32 := GenS4{F0: "foo bar"}
                    _ = fi31
                    _ = v32
                }
                for fi33 := 0; fi33 < 17; fi33++ {
                    v5 = 2147483647
                    var v34 uint8 = 2
                    _ = fi33
                    _ = v34
                }
            }
            if v23 {
                var v35 float64 = v25
                v25 += ((1e10 - ((v25 / 1.0) - (0.1 * (v25 / float64(-2.25))))) / v25)
                v36 := GenS1{F0: 127, F1: -1}
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 4; fi37++ {
                    fmt.Println(v2, fi26, v25)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi37
                }
                _ = len(v3)
                v38 := GenS1{F0: -28, F1: v5}
                _ = v38
            }
            v23 = (!true)
        }
        v23 = v23
        _ = fi26
    }
    v1 = 42
    fmt.Println(v23, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

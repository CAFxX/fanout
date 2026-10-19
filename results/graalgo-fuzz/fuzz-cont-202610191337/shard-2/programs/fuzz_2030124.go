package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
type GenS3 struct { F0 string; F1 uint }
func genFn0() int16 {
    if (("hello" >= (("a" + (("" + "hello") + "\n\t")) + "a")) && true) {
        var v4 int32 = -1
        _ = v4
    } else {
        var v5 float64 = float64(0.0)
        _ = v5
    }
    if (!false) {
        var v6 uint = 2
        _ = v6
    } else {
        var v7 int8 = -1
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint8 = 0
            _ = fi8
            _ = v9
        }
        v10 := []int{-1, 100, 32768}
        _ = v10
        _ = v7
    }
    var v11 bool = (!true)
    v11 = (424 > 418)
    _ = v11
    return 253
}

func main() {
    var v1 uint16 = 19
    var v2 float32 = 0.1
    var v3 float32 = float32(3.14159)
    v12 := GenS1{F0: 18, F1: 1000, F2: -501}
    for fi13 := 0; fi13 < 14; fi13++ {
        switch "hello" {
        case "foo bar":
            switch "a" {
            case "":
                if false {
                    v3 += float32(0.0)
                }
            case "a":
                v2 = (0.0 * ((v2 + v3) / float32(0.1)))
            case "foo bar":
                v14 := GenS2{F0: -2147483648, F1: (-57 < 0), F2: 1000}
                _ = v14
            default:
                if false {
                    v15 := []int{(fi13 - fi13), -391, (-(fi13)), fi13}
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 uint16 = v1
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v1 = v1
                        fmt.Println(fi13, v1, v3)
                        v3 = (v3 + 1.5)
                        _ = fi18
                    }
                    fmt.Println(v1, v2, v1)
                }
            }
        }
        if false {
            v3 = v3
            var v19 int64 = 114
            if false {
                var v20 string = ""
                _ = v20
            } else {
                v21 := []int{fi13, fi13, fi13}
                var v22 int32 = 42
                _ = len(v21)
                _ = v21
                _ = v22
            }
            _ = v19
        } else {
            v23 := GenS3{F0: "", F1: 1}
            v2 = ((0.0 * float32((float64(v3) - 0.0))) - (0.1 - (1e10 * (float32(3.14159) - (v3 + v2)))))
            var v24 float32 = v2
            _ = v23
            _ = v24
        }
        v3 = (1e10 + 1e10)
        _ = fi13
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        if true {
            if false {
                v2 = v3
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 float64 = (float64((v3 - v2)) / 1.0)
                    switch ("" + ("" + "a")) {
                    case "foo bar":
                        v1 += 2
                    default:
                        var v28 int16 = -1
                        _ = v28
                    }
                    v27 += 1.5
                    _ = fi26
                    _ = v27
                }
            } else {
                var v29 int8 = 100
                _ = v29
            }
            v30 := []int{fi25, (fi25 >> ((255) & 63)), (fi25 % ((2147483647) | 1))}
            _ = v30
        } else {
            switch "a" {
            case "foo bar":
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 int32 = 2147483647
                    var v33 int16 = 1
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            case "\n\t":
                var v34 string = ""
                _ = v34
            }
        }
        _ = fi25
    }
    v35 := []int{490}
    _ = len(v35)
    if (42 < 127) {
        if len(v35) > 0 {
            v35[0] = -1
        }
    }
    if len(v35) > 0 {
        v35[0] = 435
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 float32 = float32((-2.25 * 0.1))
        v2 = (((v3 - 1e10) / float32(-2.25)) / v3)
        v38 := GenS2{F0: 0, F1: (false || true), F2: -1}
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 string = ("" + "a")
    if true {
        if (4294967295 != 255) {
            var v40 string = v39
            _ = v40
        }
        v35 = append(v35, 708)
        var v41 int16 = -1
        _ = v41
    }
    v39 = "hello"
    v3 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v39
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint64, p1 bool) int32 {
    var v7 string = "a"
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := []int{(fi8 % (((fi8 >> ((255) & 63))) | 1)), fi8, 4294967295}
        var v10 string = v7
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 42}
    var v12 int8 = 42
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v7
    return -1
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = (((v1 - (((0.0 + 1.5) * (v1 * 1.5)) + -2.25)) * v1) / (float32(0.1) + (1e10 * ((v1 * v1) + v1))))
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    v5 := []int{-975}
    var v6 uint8 = 0
    _ = len(v5)
    v14 := GenS1{F0: 2}
    var v15 bool = (-1 != 0)
    for fi16 := 0; fi16 < 14; fi16++ {
        switch "foo bar" {
        case "hello":
            if (0 > fi16) {
                switch ((("\n\t" + "") + "\n\t") + ("foo bar" + "hello")) {
                case "foo bar":
                    v6 = (v6 >> ((150) & 63))
                }
            }
        case "":
            var v17 bool = (308 > 255)
            _ = v17
        default:
            v2 += v2
        }
        _ = fi16
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (720 < 42) {
        v5 = append(v5, 1)
        var v18 int32 = 0
        if ((^(v6)) > v6) {
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int32 = (v18 + v18)
                _ = fi19
                _ = v20
            }
            v6 = v6
        } else {
            v1 = 3.14159
            v6 -= (v6 % (((^(v6))) | 1))
            var v21 float64 = (1.5 / 1e10)
            _ = v21
        }
        _ = v18
    } else {
        var v22 int = 9223372036854775807
        v6 -= (v6 >> ((65535) & 63))
        _ = v22
    }
    if (v15 || v15) {
        switch 1 {
        case -110:
            if v15 {
                switch "foo bar" {
                case "":
                    v1 += v2
                case "a":
                    v23 := GenS2{F0: -2147483648}
                    _ = v23
                case "foo bar":
                    var v24 int = 4294967295
                    _ = v24
                default:
                    if v15 {
                        var v25 int8 = -128
                        var v26 string = "a"
                        _ = v25
                        _ = v26
                    }
                }
                var v27 float64 = (3.14159 - 1e10)
                _ = v27
            } else {
                switch ((("\n\t" + ((("\n\t" + ("foo bar" + ("" + "foo bar"))) + ("a" + (("" + (("\n\t" + "\n\t") + "a")) + (("\n\t" + "\n\t") + ("\n\t" + "a"))))) + ("\n\t" + ""))) + "hello") + "") {
                case "a":
                    v4 += v4
                case "hello":
                    if (-772 < 1000) {
                        v3[638] = 2147483647
                        var v28 uint = 254
                        _ = v28
                    }
                case "foo bar":
                    v29 := GenS1{F0: 0}
                    _ = v29
                default:
                    delete(v3, 4294967295)
                }
                var v30 int = -357
                _ = v30
            }
        default:
            _ = len(v5)
        }
        v4 = v4
        var v31 uint16 = 0
        _ = v31
    } else {
        v1 = v2
    }
    v1 = (float32(1e10) / 0.1)
    v6 -= v6
    for fi32 := 0; fi32 < 1; fi32++ {
        fmt.Println(v6)
        switch (fi32 / ((fi32) | 1)) {
        case 42:
            v2 = 1e10
        case -9223372036854775808:
            var v33 int8 = 42
            _ = v33
        case 179:
            v2 = (float32(3.14159) - ((-2.25 * 1e10) / -2.25))
        default:
            var v34 float64 = (1.5 * (-2.25 + (0.0 - 3.14159)))
            _ = v34
        }
        v35 := GenS3{F0: 65535, F1: fi32, F2: v6}
        _ = fi32
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        if v15 {
            v2 -= 1.5
            v4 = v4
            if (-21 > v4) {
                for fi37 := 0; fi37 < 17; fi37++ {
                    v1 -= v1
                    if v15 {
                        switch -2147483648 {
                        case 42:
                            var v38 uint32 = 42
                            _ = v38
                        case 128:
                            var v39 uint = 255
                            _ = v39
                        case 2147483647:
                            v40 := GenS4{F0: 404, F1: -1}
                            _ = v40
                        default:
                            v2 += (((1e10 * (v1 + -2.25)) / v1) * v1)
                        }
                    }
                    fmt.Println(v15, fi37)
                    _ = fi37
                }
                v2 = float32(float64(v1))
            }
        }
        fmt.Println(v15)
        _ = fi36
    }
    v2 -= v1
    _ = len(v3)
    switch -252 {
    case 32767:
        v1 = v2
    case 42:
        var v41 int32 = -895
        _ = v41
    case -2147483648:
        fmt.Println(v4, v2)
    }
    switch 895 {
    case 128:
        fmt.Println(v4, v4, v6)
    }
    v42 := GenS5{F0: (v6 - (v6 - v6))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 string; F1 uint8 }
func genFn0(p0 uint, p1 string, p2 uint) uint64 {
    var v7 int32 = 344
    switch 648 {
    case 4294967295:
        if false {
            switch p1 {
            case "foo bar":
                var v8 float64 = 1.5
                _ = v8
            case "hello":
                fmt.Println(v7, p0, v7)
            case "a":
                var v9 uint64 = 954
                _ = v9
            default:
                p0 = p2
            }
            v10 := []int{0, -2147483648}
            _ = v10
        } else {
            if ((!(!((v7 ^ (v7 % (((v7 >> (((p2 / (((p0 >> ((p2) & 63))) | 1))) & 63))) | 1))) >= v7))) || (-132 < 314)) {
                v11 := GenS1{F0: 791}
                var v12 int16 = 125
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: 1000}
                    _ = fi13
                    _ = v14
                }
                p0 = (p0 >> ((255) & 63))
            }
            var v15 int8 = 118
            _ = v15
        }
    case -961:
        if false {
            if (1000 >= 462) {
                switch (p1 + "hello") {
                case "hello":
                    if false {
                        fmt.Println(p1)
                    } else {
                        v16 := GenS2{F0: 42, F1: (^(p0))}
                        var v17 float32 = 0.0
                        v18 := []int{127, 127}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                case "\n\t":
                    v7 = (v7 >> (((^(p2))) & 63))
                default:
                    var v19 string = ((((p1 + ("foo bar" + p1)) + "\n\t") + p1) + p1)
                    _ = v19
                }
                for fi20 := 0; fi20 < 9; fi20++ {
                    p0 -= (p2 - (p2 & p0))
                    _ = fi20
                }
            }
        } else {
            if (!true) {
                p1 = (p1 + p1)
                p1 = (p1 + ((p1 + "\n\t") + "foo bar"))
                v21 := []int{9223372036854775807}
                _ = v21
            } else {
                var v22 int64 = 0
                fmt.Println(v7)
                _ = v22
            }
            var v23 int64 = 9223372036854775807
            var v24 int = -9223372036854775808
            _ = v23
            _ = v24
        }
    default:
        fmt.Println(p0)
    }
    p0 = 745
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 157
}
func genFn1(p0 uint) int {
    v25 := []int{-2147483648, 9223372036854775807}
    if true {
        _ = len(v25)
        _ = len(v25)
    } else {
        fmt.Println(p0)
        for fi26 := 0; fi26 < 1; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                fmt.Println(fi27, fi26, fi26)
                _ = fi27
            }
            p0 += (p0 % ((p0) | 1))
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi26
        }
    }
    _ = p0
    _ = v25
    return 269
}

func main() {
    var v1 int8 = 1
    var v2 bool = true
    v3 := []string{"\n\t", ((("hello" + "foo bar") + ("\n\t" + "")) + "")}
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    var v6 float64 = 1e10
    var v28 uint8 = 2
    switch "" {
    case "hello":
        v29 := GenS2{F0: genFn1(760), F1: 346}
        _ = v29
    case "a":
        var v30 int16 = 1000
        _ = v30
    case "\n\t":
        var v31 int64 = -9223372036854775808
        _ = v31
    }
    delete(v5, 2)
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 uint16 = 599
        var v34 bool = (0 <= 1)
        _ = len(v3)
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v2, v6, v28)
    switch -376 {
    case 42:
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS3{F0: (("" + "hello") + "hello"), F1: 128}
            for fi37 := 0; fi37 < 6; fi37++ {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                var v38 string = (((("\n\t" + "") + ("\n\t" + "hello")) + "") + ("\n\t" + "hello"))
                if true {
                    if v2 {
                        fmt.Println(fi37, v2, v38)
                        var v39 int64 = -1
                        var v40 float64 = v6
                        _ = v39
                        _ = v40
                    }
                } else {
                    switch (v38 + v38) {
                    case "\n\t":
                        fmt.Println(fi35)
                    case "hello":
                        _ = len(v3)
                    default:
                        if false {
                            var v41 float64 = v6
                            _ = v41
                        }
                    }
                    fmt.Println(fi35, v38)
                    if (157 <= 65535) {
                        v1 -= v1
                        v42 := GenS3{F0: "foo bar", F1: 143}
                        _ = v42
                    }
                }
                _ = fi37
                _ = v38
            }
            v1 -= (^(v1))
            _ = fi35
            _ = v36
        }
    case -709:
        var v43 bool = v2
        _ = v43
    }
    fmt.Println(v28)
    v4[1] = "foo bar"
    delete(v5, 671)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

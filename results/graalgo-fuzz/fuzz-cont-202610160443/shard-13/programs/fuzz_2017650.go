package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
func genFn0() uint16 {
    v5 := []int{1, 32768}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = 2147483648
        if ((-9223372036854775808 > 567) && true) {
            fmt.Println(v7, v7)
        } else {
            if len(v5) > 0 {
                v5[0] = 594
            }
            v8 := GenS1{F0: 32768, F1: v7, F2: 32768}
            if (false && true) {
                v9 := GenS1{F0: fi6, F1: v7, F2: -14}
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 uint32 = v7
                    var v12 bool = ((!false) && (30 < 1))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 bool = (-128 == 98)
                _ = v13
                _ = v9
            }
            _ = v8
        }
        if true {
            switch fi6 {
            case 2147483647:
                if true {
                    if false {
                        fmt.Println(fi6)
                    }
                }
            default:
                switch (-(fi6)) {
                case 2147483647:
                    var v14 uint32 = (v7 ^ v7)
                    _ = v14
                case 9223372036854775807:
                    v15 := GenS2{F0: -128, F1: -13, F2: 1}
                    _ = v15
                case -860:
                    var v16 float32 = 1.5
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if len(v5) > 0 {
                            v5[0] = fi17
                        }
                        _ = fi17
                    }
                }
            }
        }
        _ = fi6
        _ = v7
    }
    var v18 int32 = -1
    var v19 uint64 = 42
    _ = v18
    _ = v19
    _ = v5
    return 329
}
func genFn1(p0 uint64) int32 {
    p0 = p0
    if ((0.1 == 0.1) && (true && (!("\n\t" != "a")))) {
        v20 := []int{}
        _ = v20
    }
    var v21 uint32 = 0
    v21 = 4294967295
    _ = p0
    _ = v21
    return 805
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar", (("\n\t" + "foo bar") + "a"), "hello", ""}
    var v3 uint16 = 255
    var v4 uint = 18446744073709551615
    v22 := GenS3{F0: 475, F1: "foo bar", F2: "a"}
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "a" {
        case "":
            _ = len(v2)
        case "a":
            for fi24 := 0; fi24 < 7; fi24++ {
                fmt.Println(fi24, fi23)
                v2 = append(v2, "\n\t")
                _ = fi24
            }
        default:
            v4 += v4
        }
        var v25 uint8 = 66
        var v26 uint = (v4 / ((0) | 1))
        _ = fi23
        _ = v25
        _ = v26
    }
    v27 := GenS4{F0: 42}
    if ((-747 < 9223372036854775807) || ((("foo bar" + "a") + "a") != ("" + ""))) {
        v3 -= genFn0()
        for fi28 := 0; fi28 < 2; fi28++ {
            switch (^(fi28)) {
            case 32768:
                v3 += 2
            case 2147483647:
                v2 = append(v2, "foo bar")
            case 127:
                _ = len(v2)
            default:
                var v29 bool = false
                _ = v29
            }
            v30 := GenS5{F0: true}
            v3 = genFn0()
            _ = fi28
            _ = v30
        }
    } else {
        v1[0] = true
    }
    fmt.Println(v3)
    switch "a" {
    case "":
        for fi31 := 0; fi31 < 7; fi31++ {
            v4 = v4
            var v32 int8 = 42
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 float32 = 0.1
                switch "\n\t" {
                case "a":
                    if (v32 > v32) {
                        fmt.Println(v3, fi31)
                        var v35 bool = true
                        if (!(("a" != ("\n\t" + "\n\t")) && false)) {
                            if v35 {
                                v4 = (v4 & 0)
                            }
                            v36 := GenS5{F0: (2147483647 >= 698)}
                            switch "\n\t" {
                            case "foo bar":
                                var v37 float32 = float32(1e10)
                                _ = v37
                            case "\n\t":
                                v34 = (v34 + (v34 + 0.1))
                            case "hello":
                                fmt.Println(fi33, fi33, v32)
                            default:
                                var v38 uint = 65535
                                _ = v38
                            }
                            _ = v36
                        }
                        var v39 string = "hello"
                        _ = v35
                        _ = v39
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    default:
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 bool = true
            v2 = append(v2, "\n\t")
            _ = fi40
            _ = v41
        }
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v42 := v1[4294967295]
    fmt.Println(v4)
    fmt.Println(v4, v3)
    v4 -= 18446744073709551615
    var v43 int = 9223372036854775807
    switch "foo bar" {
    case "":
        fmt.Println(v3)
    case "foo bar":
        _ = len(v2)
    case "hello":
        switch (-(v43)) {
        case 2147483647:
            for fi44 := 0; fi44 < 11; fi44++ {
                if (!(!true)) {
                    var v45 int32 = genFn1(249)
                    for fi46 := 0; fi46 < 5; fi46++ {
                        v43 = (fi44 ^ -9223372036854775808)
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        switch int(v45) {
                        case 4294967295:
                            delete(v1, 0)
                        }
                        fmt.Println(v3, v42)
                        _ = fi46
                    }
                    _ = v45
                } else {
                    v47 := GenS5{F0: v42}
                    _ = v47
                }
                _ = fi44
            }
        case -1:
            fmt.Println(v3)
        case 4294967295:
            if false {
                for fi48 := 0; fi48 < 1; fi48++ {
                    fmt.Println(v43, v42)
                    _ = fi48
                }
                v3 += v3
                v42 = v42
            }
        }
    default:
        fmt.Println(v42)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 string }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint8, p1 int16, p2 uint16) int32 {
    switch -1 {
    case 128:
        if false {
            p1 -= p1
            var v7 float32 = 1.5
            v8 := GenS1{F0: 42, F1: 4294967295}
            _ = v7
            _ = v8
        } else {
            v9 := []int{32767, 2147483647}
            p0 = (p0 ^ (^(p0)))
            var v10 int8 = 91
            _ = v10
            _ = v9
        }
    case -499:
        v11 := GenS2{F0: 42, F1: (p1 * 32767), F2: 2}
        _ = v11
    default:
        v12 := []int{-1, 2147483647, 127}
        _ = v12
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 170
}
func genFn1(p0 int8, p1 bool) uint64 {
    if (p1 || (1e10 <= 0.1)) {
        v13 := []int{1, -2147483648, 9223372036854775807, -2147483648}
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int32 = 0
        if p1 {
            var v16 uint64 = 1
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        switch -2147483648 {
        case 32767:
            p0 -= 121
        case -9223372036854775808:
            var v18 int32 = -672
            _ = v18
        case 634:
            p0 -= p0
        }
        v19 := []int{int(p0), fi17, fi17, fi17}
        _ = fi17
        _ = v19
    }
    if (!p1) {
        switch (("\n\t" + (("" + "a") + "hello")) + "foo bar") {
        case "foo bar":
            for fi20 := 0; fi20 < 13; fi20++ {
                p0 -= -1
                p0 = p0
                _ = fi20
            }
        case "a":
            p0 = p0
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float64 = -2.25
    v2 := []string{""}
    var v3 int32 = 2147483647
    var v4 float64 = 1e10
    var v5 string = "\n\t"
    v6 := make(map[uint32]int)
    var v21 uint32 = 2147483648
    if len(v2) > 0 {
        v2[0] = (v5 + "hello")
    }
    if false {
        var v22 int32 = genFn0(127, -32768, 700)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 3; fi23++ {
            switch fi23 {
            case -1:
                if false {
                    v24 := v6[1]
                    v21 -= v21
                    _ = v24
                }
            case 32768:
                for fi25 := 0; fi25 < 17; fi25++ {
                    v5 = ((v5 + (v5 + "\n\t")) + (v5 + (v5 + v5)))
                    _ = fi25
                }
            case 286:
                fmt.Println(fi23)
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                switch "" {
                case "":
                    var v27 uint16 = 65535
                    _ = v27
                case "hello":
                    var v28 int16 = 1000
                    _ = v28
                default:
                    var v29 float32 = 1.5
                    _ = v29
                }
                if (42 >= 0) {
                    switch "\n\t" {
                    case "hello":
                        switch (fi26 ^ fi26) {
                        case -1:
                            if false {
                                v21 = v21
                                v5 = (((v5 + v5) + ("hello" + ((v5 + v5) + "hello"))) + v5)
                                var v30 uint32 = 0
                                _ = v30
                            }
                        case 42:
                            fmt.Println(fi26, v21)
                        default:
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        }
                        switch fi23 {
                        case -9223372036854775808:
                            switch "\n\t" {
                            case "foo bar":
                                if (0 < 839) {
                                    fmt.Println(fi26, fi23)
                                } else {
                                    v31 := v6[v21]
                                    v32 := GenS3{F0: 686, F1: -1}
                                    v33 := GenS4{F0: 128}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            case "a":
                                fmt.Println(fi23)
                            }
                        case 1:
                            var v34 float64 = ((float64(3.14159) + float64(float32(1e10))) - v1)
                            _ = v34
                        case 2147483647:
                            delete(v6, v21)
                        }
                    case "":
                        v35 := GenS4{F0: 128}
                        _ = v35
                    default:
                        v1 -= v4
                    }
                    fmt.Println(v21, v1, v5)
                    if false {
                        fmt.Println(fi23, v4, v5)
                        v21 = (^(v21))
                    }
                }
                if ((((0 >= 65535) || ("a" != "\n\t")) && true) && true) {
                    var v36 int64 = 9223372036854775807
                    _ = v36
                }
                _ = fi26
            }
            _ = fi23
        }
        fmt.Println(v5)
    }
    var v37 bool = false
    fmt.Println(v21)
    var v38 uint16 = 255
    fmt.Println(v4, v4, v38)
    _ = len(v6)
    switch v5 {
    case "foo bar":
        for fi39 := 0; fi39 < 12; fi39++ {
            v40 := GenS5{F0: (v21 / ((2) | 1)), F1: v21, F2: (((v5 + v5) + v5) + "")}
            _ = fi39
            _ = v40
        }
    case "a":
        v38 = 65535
    }
    if (((v38 / ((42) | 1)) == v38) && v37) {
        fmt.Println(v5, v38, v4)
        v41 := GenS6{F0: -32768}
        _ = v41
    }
    var v42 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}

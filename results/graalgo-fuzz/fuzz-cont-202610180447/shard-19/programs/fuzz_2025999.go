package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint64; F2 bool }
func genFn0(p0 int8) int32 {
    if false {
        var v5 int8 = -124
        var v6 float64 = float64(1.5)
        v7 := GenS1{F0: (!false)}
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 bool = true
        _ = v8
    }
    v9 := []int{195, -9223372036854775808, 127, -1}
    var v10 int32 = 1
    _ = p0
    _ = v10
    _ = v9
    return (v10 + (v10 * v10))
}
func genFn1() uint16 {
    switch "\n\t" {
    case "":
        v11 := []int{}
        _ = v11
    case "hello":
        var v12 float32 = 0.1
        _ = v12
    case "foo bar":
        var v13 uint32 = 589
        _ = v13
    }
    var v14 string = "a"
    switch 9223372036854775807 {
    case -625:
        v15 := GenS1{F0: false}
        _ = v15
    case 720:
        var v16 int = -2147483648
        _ = v16
    case 245:
        v17 := []int{520, 128, -2147483648, -307}
        _ = v17
    }
    fmt.Println(v14)
    _ = v14
    return 42
}

func main() {
    v1 := []uint32{285}
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := []int{-1}
    v18 := GenS2{F0: ("foo bar" + "")}
    var v19 uint16 = 650
    if ("a" < "") {
        v20 := GenS2{F0: (("foo bar" + "") + "")}
        v21 := GenS3{F0: v19, F1: 33}
        v19 = 448
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int64 = 42
        var v24 uint8 = 1
        switch "foo bar" {
        case "\n\t":
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v23, v19, v23)
                _ = fi25
            }
        default:
            for fi26 := 0; fi26 < 7; fi26++ {
                v4 = append(v4, 9223372036854775807)
                var v27 float32 = 3.14159
                var v28 bool = (!(-53 > 1))
                v27 -= (v27 - v27)
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    switch "" {
    case "hello":
        var v29 int = 4294967295
        _ = v29
    case "a":
        var v30 uint32 = 98
        _ = v30
    case "\n\t":
        v19 = v19
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 int32 = -1
        v33 := GenS2{F0: ("a" + "foo bar")}
        _ = fi31
        _ = v32
        _ = v33
    }
    v19 -= v19
    v19 = genFn1()
    v19 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((true || ((0 == 1) && (!true))) && false) {
        v1 = append(v1, 1)
        v19 -= 0
    } else {
        v19 += 954
    }
    v19 = 2
    switch 0 {
    case 4294967295:
        v19 -= (v19 - (v19 - 206))
    case 32768:
        v34 := GenS4{F0: 0, F1: 4294967295, F2: true}
        _ = v34
    default:
        delete(v2, 4294967295)
    }
    if false {
        for fi35 := 0; fi35 < 12; fi35++ {
            if false {
                var v36 uint8 = 42
                if (float64((float32((0.0 / (1e10 / 1e10))) + float32(-2.25))) > (-2.25 - 1e10)) {
                    var v37 uint8 = 34
                    _ = v37
                }
                _ = v36
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(fi35, fi35)
            }
            _ = fi35
        }
        var v38 uint = 0
        _ = v38
    }
    if false {
        var v39 int8 = 116
        _ = v39
    } else {
        var v40 uint16 = v19
        v3[""] = -9223372036854775808
        var v41 bool = false
        _ = v40
        _ = v41
    }
    fmt.Println(v19)
    if ((((0.0 + (0.0 * 1e10)) * 0.1) + (float32(float64(0.1)) * 1.5)) >= (1e10 - ((1e10 * 1e10) + 0.0))) {
        for fi42 := 0; fi42 < 7; fi42++ {
            fmt.Println(fi42)
            switch ("a" + "") {
            case "hello":
                var v43 uint32 = 0
                _ = v43
            default:
                switch "foo bar" {
                case "hello":
                    v19 += v19
                case "":
                    switch (fi42 ^ fi42) {
                    case 1:
                        switch (("hello" + "foo bar") + "a") {
                        case "hello":
                            delete(v2, 2147483648)
                        case "\n\t":
                            fmt.Println(fi42, v19)
                        case "":
                            var v44 string = "a"
                            _ = v44
                        }
                    case 128:
                        v19 = v19
                    case 42:
                        delete(v3, "")
                    default:
                        for fi45 := 0; fi45 < 7; fi45++ {
                            if (245 >= -9223372036854775808) {
                                var v46 uint8 = 42
                                var v47 uint = 737
                                v48 := GenS1{F0: true}
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            _ = fi45
                        }
                    }
                case "foo bar":
                    var v49 float64 = (-2.25 / 1.0)
                    _ = v49
                }
            }
            _ = fi42
        }
        if true {
            v19 = (v19 * (v19 << ((42) & 63)))
            var v50 float32 = ((((0.0 + 1e10) / (1.5 + 0.1)) * float32(((1.5 / -2.25) + (float64(1e10) + -2.25)))) / -2.25)
            _ = v50
        } else {
            v19 = (v19 | 138)
        }
    }
    v51 := GenS3{F0: (v19 ^ (v19 - (v19 % (((^(v19))) | 1)))), F1: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v51
}

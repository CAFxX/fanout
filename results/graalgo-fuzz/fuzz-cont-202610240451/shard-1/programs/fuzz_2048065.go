package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool) uint64 {
    v7 := []int{-170, 2147483647, 2147483647, 32768}
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 = append(v7, 9223372036854775807)
        _ = fi8
    }
    v7 = append(v7, 4294967295)
    v7 = append(v7, 4294967295)
    v9 := GenS1{F0: ("foo bar" + ""), F1: 1}
    _ = p0
    _ = v7
    _ = v9
    return 0
}

func main() {
    v1 := []int{42, 127}
    v2 := []uint32{2}
    v3 := []uint32{2147483648, 1, 65535, 65535}
    v4 := []uint32{465, 2147483648}
    var v5 float64 = ((0.0 / (1.5 / 1.0)) / 1e10)
    v6 := make(map[int]int)
    if true {
        var v10 float64 = 3.14159
        var v11 uint8 = 42
        v12 := GenS1{F0: "hello", F1: 42}
        _ = v10
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: ("" + ("foo bar" + ((("" + ("" + (("a" + "foo bar") + "foo bar"))) + "\n\t") + ""))), F1: 1}
        fmt.Println(v5)
        v14 := GenS1{F0: ("" + ((("a" + "\n\t") + ("hello" + "\n\t")) + "\n\t")), F1: 611}
        _ = v13
        _ = v14
    }
    v5 -= (v5 + float64(((-2.25 + (1e10 + 3.14159)) / 1.5)))
    if ((((((false && true) && ((!false) && (42 <= 2147483647))) || true) && (127 > 127)) || ((1.5 * 0.1) <= v5)) || false) {
        switch 42 {
        case 319:
            fmt.Println(v5)
        case 699:
            var v15 int64 = -9223372036854775808
            _ = v15
        case -14:
            v16 := v6[186]
            _ = v16
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(fi17, v5)
            v18 := v6[184]
            v19 := GenS1{F0: ((((("" + "foo bar") + "\n\t") + "hello") + "hello") + ""), F1: 0}
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS1{F0: "\n\t", F1: 384}
        _ = v20
    }
    fmt.Println(v5)
    if true {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 string = ("hello" + "hello")
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: ("hello" + (("hello" + "") + "foo bar")), F1: 0, F2: 692}
        _ = v23
    } else {
        var v24 bool = false
        _ = v24
    }
    v5 -= (1e10 - v5)
    v5 -= v5
    var v25 int32 = 1
    switch 32767 {
    case -993:
        var v26 int = 127
        _ = v26
    }
    if (0 == 480) {
        v27 := v6[1]
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        v6[v27] = (^(v27))
        _ = v27
    } else {
        v25 = v25
        if ((true && true) || false) {
            if false {
                fmt.Println(v5, v5)
                v25 -= v25
            }
        }
    }
    if false {
        var v28 string = "foo bar"
        switch -2147483648 {
        case 32767:
            v28 = ("foo bar" + ((v28 + "a") + v28))
        case 32768:
            if true {
                var v29 uint16 = 255
                v29 = (^(v29))
                for fi30 := 0; fi30 < 11; fi30++ {
                    v28 = (("" + ((v28 + "") + "hello")) + (v28 + ((v28 + v28) + "")))
                    _ = fi30
                }
                _ = v29
            } else {
                for fi31 := 0; fi31 < 14; fi31++ {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 float64 = v5
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v34 := GenS3{F0: genFn0(true)}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v36 := v6[fi31]
                        _ = fi35
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 3; fi37++ {
                        v38 := GenS1{F0: v28, F1: 562}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v25 -= v25
                        _ = fi37
                        _ = v38
                    }
                    _ = fi31
                }
            }
        case -288:
            var v39 int32 = -1
            _ = v39
        default:
            if (false || true) {
                if true {
                    switch v28 {
                    case "foo bar":
                        v28 = (("a" + "hello") + (v28 + ""))
                    case "a":
                        v40 := GenS2{F0: "", F1: 1, F2: 20}
                        _ = v40
                    }
                    v28 = "hello"
                    var v41 uint64 = 2
                    _ = v41
                }
                v42 := v6[-373]
                _ = v42
            }
        }
        _ = v28
    }
    fmt.Println(v5, v25)
    fmt.Println(v25, v5)
    for fi43 := 0; fi43 < 12; fi43++ {
        v5 -= v5
        _ = fi43
    }
    var v44 uint8 = 127
    if (116 < 1) {
        switch "" {
        case "hello":
            var v45 int32 = (v25 << ((65535) & 63))
            _ = v45
        default:
            fmt.Println(v44)
        }
        fmt.Println(v5)
        switch 9223372036854775807 {
        case 32768:
            v46 := GenS2{F0: "foo bar", F1: 1000, F2: 869}
            _ = v46
        case 957:
            v5 -= v5
        default:
            v44 += (v44 + v44)
        }
    }
    var v47 uint64 = 134
    v5 += float64(1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 bool }
func genFn0() uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 uint32 = 65535
        if true {
            v6 := []int{(fi4 / ((fi4) | 1)), 32768}
            var v7 bool = false
            for fi8 := 0; fi8 < 14; fi8++ {
                if (127 == 4294967295) {
                    var v9 int64 = 61
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v9
                } else {
                    switch -2147483648 {
                    case 32768:
                        fmt.Println(fi4, v5)
                    case 4294967295:
                        var v10 int64 = 0
                        _ = v10
                    case 928:
                        if v7 {
                            var v11 uint32 = (v5 % ((v5) | 1))
                            var v12 int64 = 0
                            _ = v11
                            _ = v12
                        } else {
                            v13 := GenS1{F0: -32768, F1: -36, F2: -128}
                            _ = v13
                        }
                    }
                }
                _ = fi8
            }
            _ = v6
            _ = v7
        } else {
            if (("hello" == (("foo bar" + "\n\t") + ("" + "hello"))) && (42 > 2)) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    switch 527 {
                    case 32767:
                        var v15 float64 = 1.5
                        _ = v15
                    }
                    switch fi14 {
                    case 128:
                        var v16 uint8 = 42
                        _ = v16
                    case -2147483648:
                        var v17 bool = ((^(v5)) >= 42)
                        _ = v17
                    case 127:
                        v5 = (v5 * 2147483647)
                    }
                    v5 = v5
                    _ = fi14
                }
            }
            if (32767 <= 466) {
                if (433 > -11) {
                    fmt.Println(v5)
                    var v18 int8 = 127
                    v19 := GenS2{F0: fi4, F1: 474}
                    _ = v18
                    _ = v19
                } else {
                    v5 -= v5
                    for fi20 := 0; fi20 < 16; fi20++ {
                        var v21 int32 = 991
                        v22 := []int{fi20, (fi4 % (((^(fi20))) | 1)), (fi4 | 498)}
                        v5 = v5
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := []int{(fi4 + -260)}
                    _ = v23
                }
                switch ("" + ("hello" + (("a" + (("" + "\n\t") + "hello")) + ("a" + "a")))) {
                case "foo bar":
                    var v24 int16 = -430
                    _ = v24
                case "hello":
                    v25 := []int{}
                    _ = v25
                }
            }
            v5 += 2
        }
        var v26 int16 = -1
        _ = fi4
        _ = v26
        _ = v5
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS3{F0: 168, F1: (true || (0.1 > (float64(float32((3.14159 * 3.14159))) - 3.14159)))}
        _ = fi27
        _ = v28
    }
    return 920
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{}
    var v3 float32 = (((-2.25 - 1.5) / 0.1) * (0.0 * 0.1))
    if len(v2) > 0 {
        v2[0] = 42
    }
    fmt.Println(v3)
    var v29 uint32 = 956
    v29 = (v29 | 248)
    if false {
        var v30 int32 = 230
        _ = v30
    }
    var v31 bool = (!true)
    if false {
        switch "a" {
        case "\n\t":
            if v31 {
                v32 := GenS4{F0: 708, F1: 128, F2: ((float32(float64(1e10)) != ((v3 + v3) - float32(-2.25))) && false)}
                fmt.Println(v29)
                _ = v32
            }
        case "a":
            v33 := GenS3{F0: 1, F1: (32767 > 9223372036854775807)}
            _ = v33
        }
    } else {
        var v34 int = 0
        var v35 uint = 832
        var v36 int16 = v1
        _ = v34
        _ = v35
        _ = v36
    }
    var v37 float64 = 1e10
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v31, v1)
    case "":
        switch ("" + "foo bar") {
        case "":
            fmt.Println(v29)
        case "foo bar":
            v2 = append(v2, (^(v29)))
        }
    case "a":
        if v31 {
            switch -639 {
            case 992:
                _ = len(v2)
            case 127:
                v38 := GenS4{F0: 65535, F1: 2, F2: v31}
                _ = v38
            }
            v37 += (((v37 - v37) / float64(1.5)) - ((0.1 * (v37 / ((3.14159 / 1e10) + float64(3.14159)))) / (3.14159 / v37)))
        }
    default:
        switch (("a" + "hello") + "hello") {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        case "":
            if v31 {
                v1 -= v1
                var v39 uint16 = 42
                _ = v39
            } else {
                var v40 uint8 = 128
                _ = v40
            }
        }
    }
    v37 += v37
    v2 = append(v2, (^(v29)))
    _ = len(v2)
    switch -671 {
    case 884:
        v3 = 0.1
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        if (genFn0() != 255) {
            var v42 uint16 = 42
            _ = v42
        } else {
            fmt.Println(v29)
        }
        var v43 int32 = 42
        switch fi41 {
        case -1:
            var v44 uint8 = 1
            _ = v44
        case -9223372036854775808:
            v1 -= v1
        case 128:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi41
        _ = v43
    }
    var v45 int32 = -695
    v29 = (v29 | v29)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v37
    _ = v45
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0() int32 {
    if false {
        var v6 int = 32767
        _ = v6
    } else {
        var v7 bool = (141 != 42)
        var v8 int16 = -32768
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int8 = 100
        var v11 float32 = 0.1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint8 = 2
    v12 = 1
    _ = v12
    return 0
}
func genFn1() bool {
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint32 = 4294967295
        _ = fi13
        _ = v14
    }
    var v15 int = 9223372036854775807
    var v16 int = 1
    if true {
        v16 -= v16
    }
    v17 := []int{}
    _ = v15
    _ = v16
    _ = v17
    return (true || (false && false))
}

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 float32 = 3.14159
    var v5 float32 = v4
    v5 -= (v5 - v4)
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 string = "foo bar"
        if (219 >= 174) {
            v1 = append(v1, v19)
        } else {
            switch v19 {
            case "a":
                v20 := GenS1{F0: true, F1: 0, F2: 65535}
                _ = v20
            case "\n\t":
                v5 = v5
            case "":
                for fi21 := 0; fi21 < 20; fi21++ {
                    v5 += v5
                    switch v19 {
                    case "foo bar":
                        v4 -= v4
                    default:
                        v22 := GenS1{F0: true, F1: 114, F2: 255}
                        _ = v22
                    }
                    _ = fi21
                }
            }
            var v23 int32 = 2147483647
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    v24 := GenS1{F0: true, F1: 127, F2: 1}
    var v25 uint = 356
    _ = len(v3)
    v25 += (v25 ^ 1)
    v5 += (v5 + (float32(1.5) + -2.25))
    switch -9223372036854775808 {
    case 32768:
        for fi26 := 0; fi26 < 15; fi26++ {
            v1 = append(v1, "a")
            delete(v3, 128)
            _ = fi26
        }
    }
    switch "a" {
    case "a":
        for fi27 := 0; fi27 < 14; fi27++ {
            var v28 int32 = 1
            if (!(!true)) {
                var v29 string = "\n\t"
                v4 += (float32(1.5) - (v5 + float32(-2.25)))
                _ = v29
            } else {
                if false {
                    if false {
                        var v30 float64 = 0.0
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v30
                    } else {
                        var v31 uint64 = 307
                        fmt.Println(fi27, v28, v4)
                        switch int(v25) {
                        case 127:
                            _ = len(v2)
                        case -2147483648:
                            var v32 int = 9223372036854775807
                            _ = v32
                        default:
                            switch fi27 {
                            case -431:
                                fmt.Println(v28)
                            case -9223372036854775808:
                                var v33 uint8 = 230
                                _ = v33
                            }
                        }
                        _ = v31
                    }
                    switch "a" {
                    case "":
                        var v34 float64 = 1.5
                        _ = v34
                    case "foo bar":
                        var v35 int32 = -367
                        _ = v35
                    case "\n\t":
                        var v36 int16 = -1
                        _ = v36
                    default:
                        var v37 float32 = (1e10 / (0.0 + (v5 - v5)))
                        _ = v37
                    }
                } else {
                    switch ((("foo bar" + "foo bar") + ("hello" + "foo bar")) + "a") {
                    case "hello":
                        v5 -= 0.1
                    case "\n\t":
                        v38 := GenS2{F0: v25}
                        _ = v38
                    case "foo bar":
                        for fi39 := 0; fi39 < 16; fi39++ {
                            v5 -= (v4 * v5)
                            v2["foo bar"] = 128
                            if true {
                                var v40 uint8 = 1
                                if (42 < 0) {
                                    v5 = float32((-2.25 + 1e10))
                                    var v41 int8 = 100
                                    _ = v41
                                }
                                v2[(("\n\t" + "\n\t") + ("foo bar" + ("\n\t" + ("foo bar" + ""))))] = (fi27 / (((fi39 * (fi39 << ((uint(v40)) & 63)))) | 1))
                                _ = v40
                            }
                            _ = fi39
                        }
                    }
                    for fi42 := 0; fi42 < 12; fi42++ {
                        _ = len(v3)
                        var v43 int32 = v28
                        _ = fi42
                        _ = v43
                    }
                }
            }
            _ = fi27
            _ = v28
        }
    case "\n\t":
        for fi44 := 0; fi44 < 2; fi44++ {
            switch fi44 {
            case 1:
                for fi45 := 0; fi45 < 2; fi45++ {
                    v4 = (0.1 + 1e10)
                    v4 = (v5 / (1e10 + (float32((((-2.25 * 0.0) - ((0.1 * (3.14159 / 1e10)) - -2.25)) / float64(1e10))) / (v4 * 3.14159))))
                    _ = fi45
                }
            }
            _ = fi44
        }
    default:
        for fi46 := 0; fi46 < 16; fi46++ {
            switch (fi46 + fi46) {
            case -2147483648:
                var v47 int64 = -9223372036854775808
                _ = v47
            case -1:
                if false {
                    v25 = v25
                    var v48 int64 = 660
                    if (!true) {
                        v5 -= v5
                        if false {
                            v1 = append(v1, "hello")
                            var v49 int64 = v48
                            _ = v49
                        }
                    } else {
                        v50 := GenS1{F0: ((v25 << (((v25 | v25)) & 63)) != 890), F1: 0, F2: 1}
                        v48 += v48
                        _ = v50
                    }
                    _ = v48
                }
            }
            switch ("foo bar" + ("hello" + "a")) {
            case "hello":
                if true {
                    if true {
                        v5 = (v5 * v5)
                        v25 -= (^(v25))
                        v51 := GenS3{F0: 127, F1: 127, F2: 0}
                        _ = v51
                    }
                    if false {
                        var v52 uint = 1
                        fmt.Println(v25, v52)
                        _ = v52
                    } else {
                        v53 := GenS3{F0: 100, F1: -1, F2: 2147483647}
                        var v54 float32 = 3.14159
                        _ = v53
                        _ = v54
                    }
                }
            case "foo bar":
                var v55 int = (fi46 >> ((42) & 63))
                _ = v55
            case "\n\t":
                var v56 float32 = 3.14159
                _ = v56
            default:
                var v57 uint32 = 0
                _ = v57
            }
            _ = fi46
        }
    }
    v25 += 65535
    for fi58 := 0; fi58 < 14; fi58++ {
        switch "hello" {
        case "a":
            var v59 uint8 = 42
            _ = v59
        case "foo bar":
            v25 = 0
        }
        v3[(fi58 >> ((181) & 63))] = false
        for fi60 := 0; fi60 < 2; fi60++ {
            if true {
                delete(v3, fi60)
                var v61 int64 = -1
                _ = v61
            }
            delete(v2, "hello")
            var v62 float64 = 1.5
            _ = fi60
            _ = v62
        }
        _ = fi58
    }
    v3[0] = (false && true)
    var v63 uint64 = 936
    var v64 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }
func genFn0(p0 string) bool {
    v7 := GenS1{F0: 1, F1: -141}
    if true {
        fmt.Println(p0)
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 = p0
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint32 = 42
        var v11 uint16 = 255
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint64 = 42
            p0 = p0
            _ = fi13
            _ = v14
        }
        v15 := []int{}
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint8 = 0
            var v18 uint = 0
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi12
        _ = v15
    }
    v19 := []int{-9223372036854775808, 4294967295, -832, 934}
    _ = p0
    _ = v19
    _ = v7
    return (32768 == 127)
}
func genFn1() int32 {
    var v20 int16 = 169
    fmt.Println(v20)
    var v21 float64 = 1.5
    _ = v20
    _ = v21
    return 0
}

func main() {
    v1 := []string{"a"}
    var v2 float64 = 3.14159
    var v3 int8 = -87
    var v4 int = 128
    v5 := make(map[int]int)
    var v6 uint16 = 22
    delete(v5, (v4 - 0))
    var v22 float32 = ((0.1 + 1e10) + 0.1)
    var v23 uint32 = 0
    var v24 bool = false
    v25 := v5[v4]
    fmt.Println(v4)
    if v24 {
        if ((v23 * (v23 >> ((2) & 63))) < (v23 >> ((335) & 63))) {
            v5[(v25 >> ((4294967295) & 63))] = (v25 & v4)
            if (v25 == v4) {
                var v26 string = "a"
                v26 = "\n\t"
                v3 -= -104
                _ = v26
            } else {
                for fi27 := 0; fi27 < 16; fi27++ {
                    if len(v1) > 0 {
                        v1[0] = ("foo bar" + (((("foo bar" + "\n\t") + ("\n\t" + (("hello" + "\n\t") + "hello"))) + "foo bar") + "\n\t"))
                    }
                    _ = fi27
                }
                v22 += (0.1 - (3.14159 - ((0.0 * v22) + 0.1)))
                v28 := GenS2{F0: (v6 >> ((2) & 63)), F1: ((("\n\t" + (((("a" + ("hello" + "\n\t")) + "\n\t") + ("" + ("hello" + ""))) + "\n\t")) + "\n\t") + "hello"), F2: int64(v6)}
                _ = v28
            }
            v25 += 985
        }
        var v29 int16 = 42
        switch (("a" + ("foo bar" + "")) + ("a" + (("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) + ((("\n\t" + "foo bar") + "") + ("\n\t" + ("hello" + "")))))) {
        case "foo bar":
            var v30 uint32 = v23
            _ = v30
        }
        _ = v29
    } else {
        switch "\n\t" {
        case "\n\t":
            var v31 int16 = 0
            _ = v31
        case "a":
            if v24 {
                if ("a" == "\n\t") {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 float32 = v22
                        delete(v5, (v4 & -290))
                        var v34 uint32 = (v23 + 65535)
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch (-(v25)) {
                    case 2147483647:
                        if len(v1) > 0 {
                            v1[0] = ("\n\t" + "hello")
                        }
                    case 127:
                        for fi35 := 0; fi35 < 16; fi35++ {
                            for fi36 := 0; fi36 < 9; fi36++ {
                                if (v4 <= -9223372036854775808) {
                                    v37 := GenS2{F0: 65535, F1: "", F2: -1}
                                    _ = v37
                                }
                                if v24 {
                                    v6 += 1000
                                } else {
                                    v24 = (v24 || ((v6 - (v6 % ((v6) | 1))) >= (v6 * v6)))
                                    if (!false) {
                                        var v38 float32 = v22
                                        _ = v38
                                    }
                                    v39 := GenS3{F0: ("foo bar" + ("foo bar" + "\n\t"))}
                                    _ = v39
                                }
                                v40 := GenS3{F0: "\n\t"}
                                _ = fi36
                                _ = v40
                            }
                            for fi41 := 0; fi41 < 19; fi41++ {
                                var v42 float32 = 3.14159
                                fmt.Println(v42)
                                var v43 float64 = (v2 / 1.0)
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS1{F0: v25, F1: -366}
                            _ = fi35
                            _ = v44
                        }
                    case -9223372036854775808:
                        _ = len(v5)
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                delete(v5, (v25 ^ v25))
            } else {
                var v45 int8 = (v3 / (((^(v3))) | 1))
                switch ((("foo bar" + "hello") + "a") + "foo bar") {
                case "hello":
                    v22 = float32(v2)
                case "":
                    switch ("hello" + "a") {
                    case "hello":
                        if (1e10 >= float32(v2)) {
                            switch "\n\t" {
                            case "":
                                if v24 {
                                    for fi46 := 0; fi46 < 4; fi46++ {
                                        var v47 bool = (481 > 190)
                                        v25 -= (fi46 ^ -774)
                                        v48 := GenS4{F0: 1}
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                    v49 := GenS5{F0: uint8(v23)}
                                    _ = v49
                                } else {
                                    for fi50 := 0; fi50 < 11; fi50++ {
                                        _ = len(v5)
                                        var v51 uint64 = 18446744073709551615
                                        _ = fi50
                                        _ = v51
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        var v52 string = "\n\t"
                                        _ = v52
                                    case "hello":
                                        var v53 bool = v24
                                        _ = v53
                                    }
                                }
                            case "a":
                                var v54 uint8 = 0
                                _ = v54
                            default:
                                for fi55 := 0; fi55 < 7; fi55++ {
                                    v56 := v5[2147483647]
                                    var v57 float64 = v2
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                            }
                            if v24 {
                                var v58 int8 = (v45 + v3)
                                _ = len(v1)
                                _ = v58
                            }
                        }
                    default:
                        if ("hello" != "a") {
                            if true {
                                var v59 int64 = 1
                                _ = v59
                            }
                            for fi60 := 0; fi60 < 3; fi60++ {
                                v61 := GenS6{F0: 1, F1: 65535, F2: 1}
                                _ = fi60
                                _ = v61
                            }
                            switch (v4 + 9223372036854775807) {
                            case 127:
                                _ = len(v1)
                            default:
                                var v62 int32 = -151
                                _ = v62
                            }
                        }
                    }
                }
                if (42 <= 42) {
                    v63 := GenS4{F0: 2}
                    if ((-804 != -1) && (184 >= 216)) {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        v22 = (v22 - (0.1 * 0.0))
                    }
                    var v64 uint8 = 127
                    _ = v63
                    _ = v64
                }
                _ = v45
            }
        }
    }
    for fi65 := 0; fi65 < 16; fi65++ {
        v66 := GenS4{F0: 18446744073709551615}
        var v67 bool = (!(!(v2 != (1.5 - ((v2 - ((0.0 / (float64((((float32(1e10) + v22) - v22) + (v22 / (v22 - v22)))) / (v2 * (v2 - v2)))) + (-2.25 - float64(1.5)))) - v2)))))
        fmt.Println(v22, v23, v24)
        _ = fi65
        _ = v66
        _ = v67
    }
    var v68 int = 704
    v22 = (0.1 / 0.1)
    v69 := GenS2{F0: v6, F1: ("" + ""), F2: 46}
    var v70 int32 = 673
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 bool }
type GenS2 struct { F0 int16; F1 int16; F2 string }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    switch "hello" {
    case "":
        for fi4 := 0; fi4 < 19; fi4++ {
            for fi5 := 0; fi5 < 12; fi5++ {
                var v6 int32 = -1
                v7 := []int{1, fi5, p0}
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 bool = (false && true)
            _ = fi4
            _ = v8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := []int{}
            var v11 uint64 = 0
            var v12 bool = (false && (!true))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            p0 = (fi13 >> ((2) & 63))
            _ = fi13
        }
    }
    _ = p0
    return 677
}
func genFn1(p0 uint32) int {
    fmt.Println(p0)
    if false {
        v14 := GenS1{F0: false, F1: (("\n\t" + ("foo bar" + ("foo bar" + ("" + ("hello" + (("foo bar" + ("foo bar" + "foo bar")) + "foo bar")))))) + ("\n\t" + "")), F2: true}
        v15 := GenS1{F0: true, F1: "hello", F2: ((!(9223372036854775807 == -9223372036854775808)) || true)}
        switch ("" + "hello") {
        case "a":
            if false {
                if true {
                    v16 := []int{32767}
                    switch 42 {
                    case 564:
                        for fi17 := 0; fi17 < 19; fi17++ {
                            if (true && ("foo bar" < "a")) {
                                var v18 int = fi17
                                if (265 >= -492) {
                                    v18 += fi17
                                }
                                v18 = (fi17 | 52)
                                _ = v18
                            }
                            v19 := GenS1{F0: (0 > int16(p0)), F1: ("hello" + "foo bar"), F2: false}
                            if len(v16) > 0 {
                                v16[0] = (fi17 % ((fi17) | 1))
                            }
                            _ = fi17
                            _ = v19
                        }
                    default:
                        var v20 float32 = (0.1 - 1e10)
                        _ = v20
                    }
                    var v21 uint = uint(p0)
                    _ = v16
                    _ = v21
                }
                v22 := []int{-394}
                for fi23 := 0; fi23 < 14; fi23++ {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        fmt.Println(fi24)
                        _ = fi24
                    }
                    var v25 bool = (2 != 0)
                    _ = fi23
                    _ = v25
                }
                _ = v22
            } else {
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := []int{(fi26 * fi26), fi26, fi26}
                    var v28 int8 = -85
                    for fi29 := 0; fi29 < 10; fi29++ {
                        v28 = -44
                        p0 -= genFn0(fi29)
                        _ = fi29
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        case "\n\t":
            fmt.Println(p0)
        }
        _ = v14
        _ = v15
    } else {
        if (true || false) {
            switch "\n\t" {
            case "\n\t":
                if true {
                    var v30 int = -421
                    _ = v30
                }
            case "foo bar":
                var v31 float32 = (1e10 - -2.25)
                _ = v31
            case "a":
                var v32 uint64 = 2
                _ = v32
            default:
                if (0 == (p0 << ((701) & 63))) {
                    if (!((((255 >= 2) && true) || (!("hello" > ""))) || false)) {
                        v33 := GenS2{F0: 1, F1: 816, F2: ""}
                        _ = v33
                    } else {
                        v34 := []int{2147483647}
                        switch "\n\t" {
                        case "\n\t":
                            var v35 int64 = -1
                            _ = v35
                        case "hello":
                            if (42 >= 127) {
                                v34 = append(v34, 2147483647)
                            } else {
                                if len(v34) > 0 {
                                    v34[0] = 32768
                                }
                                if true {
                                    var v36 int32 = -1
                                    var v37 int32 = -2147483648
                                    _ = v36
                                    _ = v37
                                }
                            }
                        }
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                        _ = v34
                    }
                    if (-114 > -128) {
                        for fi38 := 0; fi38 < 5; fi38++ {
                            v39 := GenS2{F0: 1000, F1: -646, F2: "\n\t"}
                            _ = fi38
                            _ = v39
                        }
                        v40 := GenS3{F0: 184, F1: 292}
                        _ = v40
                    } else {
                        var v41 uint64 = 601
                        _ = v41
                    }
                }
            }
        } else {
            p0 = (p0 << ((158) & 63))
            var v42 int64 = 42
            _ = v42
        }
        v43 := []int{2147483647}
        _ = v43
    }
    if (!(true || true)) {
        var v44 float32 = (3.14159 * (((1.5 / 1.0) - (float32((1.5 + 3.14159)) / 1.5)) - (0.1 - 0.0)))
        v44 = 0.0
        _ = v44
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        p0 += (p0 ^ (p0 & (p0 | p0)))
        for fi46 := 0; fi46 < 16; fi46++ {
            p0 += p0
            var v47 uint = 1
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    _ = p0
    return int(p0)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{127, 128}
    var v3 uint64 = 2
    v1[genFn0(genFn1(2147483647))] = 0
    v48 := GenS1{F0: (!false), F1: "hello", F2: false}
    if false {
        v3 = 1
        if (!false) {
            v49 := GenS4{F0: 622}
            _ = v49
        } else {
            delete(v1, 36)
            var v50 uint8 = 0
            v51 := GenS1{F0: true, F1: ("" + (("a" + (((("" + ((("" + "") + ("" + "\n\t")) + "a")) + (("hello" + "foo bar") + "a")) + "a") + "\n\t")) + "a")), F2: true}
            _ = v50
            _ = v51
        }
        fmt.Println(v3)
    } else {
        for fi52 := 0; fi52 < 1; fi52++ {
            _ = len(v1)
            var v53 bool = (v3 == (v3 & 1))
            _ = fi52
            _ = v53
        }
    }
    switch ("foo bar" + "hello") {
    case "":
        v3 -= v3
    case "a":
        if (true && true) {
            var v54 int8 = 127
            delete(v1, 1)
            _ = v54
        }
    default:
        for fi55 := 0; fi55 < 5; fi55++ {
            v3 += v3
            _ = fi55
        }
    }
    var v56 float32 = 0.0
    v57 := GenS2{F0: 125, F1: -32768, F2: ("\n\t" + "foo bar")}
    v58 := GenS2{F0: 42, F1: 0, F2: "foo bar"}
    v3 = (v3 << ((67) & 63))
    v3 = v3
    fmt.Println(v56, v56)
    var v59 int = 32768
    v59 = (v59 ^ v59)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}

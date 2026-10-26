package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int64; F2 bool }
func genFn0(p0 bool, p1 int64, p2 uint8) bool {
    p2 -= (p2 | p2)
    if (!((p0 || (3.14159 <= float64(float32(-2.25)))) && p0)) {
        p2 += p2
        var v5 int8 = 100
        _ = v5
    }
    switch 128 {
    case 1:
        v6 := GenS1{F0: 2, F1: ("hello" != "a"), F2: 255}
        _ = v6
    }
    var v7 int64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}
func genFn1(p0 bool, p1 uint) bool {
    switch 1 {
    case 9223372036854775807:
        v8 := []int{}
        _ = v8
    case 127:
        if (p0 && p0) {
            var v9 uint8 = 127
            _ = v9
        } else {
            if (-128 > -9) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    p1 += p1
                    for fi11 := 0; fi11 < 19; fi11++ {
                        switch (("foo bar" + "hello") + (((("foo bar" + "a") + "") + "foo bar") + ("hello" + ("hello" + ((("hello" + ("hello" + ("" + ""))) + (("\n\t" + "a") + ("foo bar" + "foo bar"))) + ""))))) {
                        case "":
                            p0 = true
                        case "hello":
                            var v12 uint32 = uint32(p1)
                            _ = v12
                        case "\n\t":
                            if (622 >= 315) {
                                var v13 string = "hello"
                                _ = v13
                            } else {
                                v14 := GenS2{F0: 1, F1: 0, F2: 150}
                                _ = v14
                            }
                        default:
                            v15 := GenS3{F0: 2}
                            _ = v15
                        }
                        p0 = (false && (1 > 0))
                        _ = fi11
                    }
                    _ = fi10
                }
                var v16 uint8 = 42
                _ = v16
            }
            if false {
                v17 := []int{-1, 32768}
                fmt.Println(p0)
                _ = v17
            } else {
                var v18 float32 = (1e10 * 1.5)
                _ = v18
            }
            if (!p0) {
                v19 := GenS4{F0: 116, F1: -361, F2: 32767}
                _ = v19
            } else {
                p0 = (-9223372036854775808 != 0)
                if (0 != 42) {
                    if (!((float32((1e10 + 0.1)) - -2.25) > 0.1)) {
                        switch 128 {
                        case -9223372036854775808:
                            var v20 float32 = float32(0.1)
                            _ = v20
                        }
                    } else {
                        var v21 int = 0
                        v22 := []int{(v21 & 127), v21, 32768, -288}
                        v22 = append(v22, int(p1))
                        _ = v21
                        _ = v22
                    }
                    var v23 uint64 = 1
                    fmt.Println(p1)
                    _ = v23
                }
                if (!false) {
                    fmt.Println(p0, p1)
                    if false {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            p0 = p0
                            _ = fi24
                        }
                    } else {
                        v25 := GenS5{F0: 2147483648}
                        if p0 {
                            v26 := []int{-810, 32767, -833, 32767}
                            var v27 uint8 = 127
                            _ = v26
                            _ = v27
                        }
                        _ = v25
                    }
                    var v28 uint8 = 0
                    _ = v28
                }
            }
        }
    default:
        switch 4294967295 {
        case 0:
            v29 := []int{-136, -979, 0}
            _ = v29
        case -854:
            v30 := []int{221, -1}
            _ = v30
        case 42:
            if (!(1 <= -208)) {
                p0 = false
                v31 := GenS2{F0: 0, F1: 0, F2: 2}
                var v32 int8 = -128
                _ = v31
                _ = v32
            }
        }
    }
    p1 = p1
    v33 := GenS6{F0: 42, F1: 9223372036854775807, F2: p0}
    var v34 int8 = 127
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    return (42 < 1)
}

func main() {
    v1 := []string{(("hello" + ("foo bar" + "")) + "a"), "\n\t", ("hello" + ("foo bar" + ""))}
    var v2 int64 = -894
    var v3 float64 = ((0.0 - 1e10) - 0.0)
    v4 := make(map[string]string)
    switch -2147483648 {
    case 128:
        if false {
            switch 1 {
            case 251:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 9223372036854775807:
                var v35 float64 = v3
                _ = v35
            case 32768:
                var v36 uint64 = 1
                _ = v36
            default:
                delete(v4, "a")
            }
            var v37 uint32 = 1
            _ = v37
        }
    default:
        var v38 int8 = 42
        _ = v38
    }
    v4["foo bar"] = ((("" + "\n\t") + ("hello" + "hello")) + ("a" + ("" + "\n\t")))
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 uint = 0
        for fi41 := 0; fi41 < 9; fi41++ {
            fmt.Println(v2)
            v1 = append(v1, (("foo bar" + "hello") + "foo bar"))
            v40 -= (v40 | v40)
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    v42 := GenS1{F0: 2, F1: false, F2: 255}
    fmt.Println(v3, v2)
    v43 := GenS5{F0: 2147483648}
    var v44 int8 = -1
    for fi45 := 0; fi45 < 3; fi45++ {
        v44 -= v44
        _ = fi45
    }
    fmt.Println(v2, v3)
    if (738 < 198) {
        var v46 float64 = 0.0
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
}

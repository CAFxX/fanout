package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 bool; F2 string }
type GenS8 struct { F0 int64; F1 int }
func genFn0(p0 string, p1 string, p2 int) int64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{(p2 | (fi5 >> ((65535) & 63))), (fi5 + (p2 ^ 42)), (fi5 | (fi5 << ((0) & 63))), fi5}
        v6 = append(v6, fi5)
        p2 -= 0
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int32, p1 bool, p2 uint) uint32 {
    switch ("\n\t" + (("foo bar" + (((("hello" + "hello") + "foo bar") + "foo bar") + "hello")) + ("" + ("\n\t" + "a")))) {
    case "a":
        var v8 uint16 = 999
        _ = v8
    case "":
        var v9 int8 = -128
        _ = v9
    default:
        if p1 {
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 uint = p2
                    var v13 int16 = 32767
                    if (!false) {
                        if (p1 && false) {
                            var v14 uint = 4294967295
                            v15 := GenS1{F0: 9223372036854775807}
                            _ = v14
                            _ = v15
                        } else {
                            if p1 {
                                switch (fi11 + fi11) {
                                case -1:
                                    if (v13 <= 1) {
                                        var v16 float64 = 0.0
                                        if true {
                                            v17 := []int{(^(fi11)), 4294967295, (-(fi11)), 685}
                                            var v18 uint64 = 42
                                            _ = v17
                                            _ = v18
                                        }
                                        var v19 uint = p2
                                        _ = v16
                                        _ = v19
                                    } else {
                                        v12 -= 0
                                        p1 = (-9223372036854775808 != genFn0("", "a", fi11))
                                    }
                                case 9223372036854775807:
                                    p1 = p1
                                default:
                                    v20 := GenS1{F0: genFn0("hello", "\n\t", fi11)}
                                    _ = v20
                                }
                            }
                            var v21 float32 = 1e10
                            v22 := GenS1{F0: -296}
                            _ = v21
                            _ = v22
                        }
                        fmt.Println(v13)
                    } else {
                        fmt.Println(p1, v13)
                        var v23 uint8 = 128
                        _ = v23
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if (!p1) {
                    if p1 {
                        v24 := GenS2{F0: 1}
                        var v25 float64 = ((1e10 * 1e10) - 1e10)
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(p0, p0)
                    var v26 uint32 = 788
                    _ = v26
                }
                _ = fi10
            }
            v27 := GenS3{F0: -258, F1: 0}
            _ = v27
        } else {
            fmt.Println(p2, p2)
        }
    }
    fmt.Println(p0)
    if p1 {
        for fi28 := 0; fi28 < 1; fi28++ {
            if (-9223372036854775808 >= 75) {
                p1 = true
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := GenS4{F0: 57}
                    var v31 bool = ((p0 & 1) > p0)
                    v32 := []int{-949, (fi28 << (((p2 & (p2 + 255))) & 63))}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
            v33 := GenS5{F0: 1, F1: 4294967295}
            _ = fi28
            _ = v33
        }
        v34 := GenS6{F0: p2}
        v35 := []int{127, int(p0), -1000, 0}
        _ = v34
        _ = v35
    } else {
        var v36 int64 = -1
        _ = v36
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 98
}

func main() {
    var v1 int = 42
    var v2 int16 = -1
    var v3 uint16 = 120
    var v4 uint8 = 212
    v3 = uint16(v4)
    var v37 int8 = 42
    var v38 int32 = 2147483647
    var v39 int16 = int16(v37)
    v40 := GenS5{F0: v38, F1: 425}
    if false {
        fmt.Println(v1, v3)
        v39 = (^(v2))
    } else {
        v41 := GenS5{F0: (v38 / ((v38) | 1)), F1: 0}
        switch ("foo bar" + ((("\n\t" + "") + ("\n\t" + "foo bar")) + ("\n\t" + ("hello" + "a")))) {
        case "a":
            v42 := []int{v1, v1, (v1 * int(v4))}
            _ = v42
        case "":
            v39 = v39
        }
        v43 := []int{}
        _ = v41
        _ = v43
    }
    switch ((("\n\t" + (((("a" + "a") + (("\n\t" + "\n\t") + "a")) + "\n\t") + "a")) + ("foo bar" + ("" + "\n\t"))) + "") {
    case "":
        if false {
            v44 := []int{4294967295, (v1 & v1)}
            var v45 string = ((("foo bar" + ("hello" + ("" + "foo bar"))) + "") + ("" + ""))
            _ = v44
            _ = v45
        } else {
            v46 := []int{}
            _ = v46
        }
    case "hello":
        v1 = (v1 + v1)
    default:
        v47 := GenS7{F0: 0, F1: true, F2: ("foo bar" + ("" + "a"))}
        _ = v47
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        var v49 bool = false
        for fi50 := 0; fi50 < 2; fi50++ {
            v51 := []int{}
            v51 = append(v51, fi48)
            v3 = (^(v3))
            _ = fi50
            _ = v51
        }
        _ = fi48
        _ = v49
    }
    var v52 uint16 = v3
    v53 := []int{1}
    switch v1 {
    case 42:
        if false {
            for fi54 := 0; fi54 < 15; fi54++ {
                for fi55 := 0; fi55 < 19; fi55++ {
                    fmt.Println(fi55)
                    _ = fi55
                }
                for fi56 := 0; fi56 < 11; fi56++ {
                    v38 = (v38 % (((v38 >> ((255) & 63))) | 1))
                    _ = len(v53)
                    var v57 int16 = (v2 | v2)
                    v58 := GenS8{F0: genFn0("a", "a", (fi54 ^ fi54)), F1: fi56}
                    v53 = append(v53, (fi56 + (^(fi54))))
                    v4 = 111
                    var v59 int32 = v38
                    _ = fi56
                    _ = v57
                    _ = v58
                    _ = v59
                }
                _ = fi54
            }
            switch "a" {
            case "hello":
                v39 += (v39 | (v39 >> ((4294967295) & 63)))
            default:
                fmt.Println(v3, v2, v2)
            }
            var v60 float32 = 0.0
            _ = v60
        } else {
            var v61 uint = 65535
            _ = v61
        }
    case 0:
        for fi62 := 0; fi62 < 1; fi62++ {
            var v63 float32 = 3.14159
            _ = fi62
            _ = v63
        }
    case -9223372036854775808:
        if len(v53) >= 2 {
            v53 = v53[1:len(v53)]
        }
    }
    fmt.Println(v39, v38)
    fmt.Println(v39)
    var v64 int16 = v39
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v52
    _ = v53
    _ = v64
}

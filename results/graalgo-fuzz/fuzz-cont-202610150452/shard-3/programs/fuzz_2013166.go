package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0(p0 uint, p1 int8) int {
    p0 = (p0 - p0)
    var v7 uint = (^(p0))
    v8 := GenS1{F0: 749}
    var v9 int64 = -952
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 262
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (-1 > 1)
    var v3 uint64 = 42
    v4 := make(map[int]string)
    var v5 string = "\n\t"
    v6 := []string{"", (v5 + v5)}
    fmt.Println(v2, v3)
    for fi10 := 0; fi10 < 8; fi10++ {
        v5 = v5
        _ = fi10
    }
    v5 = ("a" + v5)
    v2 = (0 <= -32768)
    for fi11 := 0; fi11 < 19; fi11++ {
        if v2 {
            v5 = v5
        } else {
            if (423 >= -217) {
                switch v5 {
                case "foo bar":
                    if false {
                        _ = len(v6)
                    }
                case "\n\t":
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 int32 = -2147483648
                        _ = len(v6)
                        var v14 float64 = (0.1 - float64(-2.25))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "":
                    v5 = "a"
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    switch (fi11 / ((127) | 1)) {
                    case -2147483648:
                        fmt.Println(v5, v2)
                    case 32768:
                        delete(v4, fi11)
                    }
                    if false {
                        if ((v5 + v5) != v5) {
                            fmt.Println(v5)
                            var v16 string = ((v5 + "") + "foo bar")
                            v5 = ("a" + (v5 + (v5 + v16)))
                            v16 = v16
                            _ = v16
                        } else {
                            var v17 float64 = 1.5
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v17
                        }
                        var v18 int64 = 616
                        _ = v18
                    }
                    var v19 uint8 = 0
                    _ = fi15
                    _ = v19
                }
            }
            if (0 != fi11) {
                if false {
                    if v2 {
                        fmt.Println(v5)
                    }
                    var v20 bool = (!true)
                    _ = v20
                }
            } else {
                var v21 string = v5
                v5 = v5
                v22 := GenS2{F0: fi11, F1: -9223372036854775808}
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                _ = len(v1)
                var v24 float64 = ((0.0 / float64(-2.25)) / 3.14159)
                v3 += v3
                var v25 int64 = -462
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        var v26 int32 = 42
        v3 += (v3 + 1)
        _ = fi11
        _ = v26
    }
    v3 += (v3 + v3)
    var v27 uint = 65535
    var v28 int32 = 2147483647
    if (42 > -754) {
        v29 := GenS2{F0: genFn0(0, -5), F1: 128}
        if false {
            if ((v28 / (((^(v28))) | 1)) <= (v28 ^ (v28 >> ((v27) & 63)))) {
                var v30 uint64 = (v3 | v3)
                _ = v30
            }
            var v31 uint32 = 65535
            v32 := GenS1{F0: uint64(v31)}
            _ = v31
            _ = v32
        } else {
            if ((1.5 / ((1e10 / -2.25) - (0.0 / -2.25))) > (1.5 + -2.25)) {
                v6 = append(v6, "")
                var v33 int = 42
                v34 := GenS3{F0: 481, F1: v28}
                _ = v33
                _ = v34
            }
            v35 := GenS2{F0: -535, F1: 32768}
            _ = v35
        }
        _ = v29
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        v5 = ""
        fmt.Println(v27, v28)
        _ = fi36
    }
    if true {
        v37 := GenS4{F0: 1, F1: 1}
        _ = v37
    } else {
        switch -826 {
        case 128:
            if false {
                v3 = (^(v3))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case 949:
            v2 = true
        case 32767:
            if true {
                if (v5 > (("a" + (("\n\t" + (("hello" + "hello") + v5)) + v5)) + v5)) {
                    if v2 {
                        v2 = (false || false)
                        var v38 int = 32768
                        if (!v2) {
                            for fi39 := 0; fi39 < 19; fi39++ {
                                v3 = v3
                                _ = fi39
                            }
                            switch genFn0((^(v27)), 42) {
                            case 32768:
                                if true {
                                    for fi40 := 0; fi40 < 11; fi40++ {
                                        v41 := GenS5{F0: 895, F1: (!v2), F2: 42}
                                        _ = fi40
                                        _ = v41
                                    }
                                    var v42 string = ((("foo bar" + v5) + v5) + v5)
                                    fmt.Println(v2, v27, v2)
                                    _ = v42
                                } else {
                                    v28 += (v28 % (((-(v28))) | 1))
                                }
                            }
                            for fi43 := 0; fi43 < 18; fi43++ {
                                v38 += (v38 | (fi43 | (v38 ^ genFn0(v27, 1))))
                                v2 = v2
                                _ = fi43
                            }
                        } else {
                            var v44 uint = v27
                            _ = v44
                        }
                        _ = v38
                    }
                    fmt.Println(v3, v3)
                } else {
                    var v45 uint64 = 900
                    _ = v45
                }
                fmt.Println(v5, v3)
            }
        default:
            v27 -= v27
        }
        v27 -= 404
        v46 := GenS2{F0: 127, F1: 1}
        _ = v46
    }
    switch v5 {
    case "\n\t":
        switch -847 {
        case -545:
            var v47 int16 = 1
            _ = v47
        default:
            v48 := GenS3{F0: 42, F1: -2147483648}
            _ = v48
        }
    default:
        var v49 uint8 = 255
        _ = v49
    }
    if (1.5 <= ((1.5 + ((1.5 + 0.0) - 1e10)) / float32(1.5))) {
        fmt.Println(v28, v2)
        if (32767 >= -1) {
            fmt.Println(v28)
            v28 = v28
        }
        if true {
            v4[genFn0(65535, -31)] = "hello"
        } else {
            var v50 int32 = v28
            if (!(543 < -104)) {
                var v51 bool = true
                _ = v51
            }
            var v52 uint64 = 4294967295
            _ = v50
            _ = v52
        }
    }
    if (322 > uint16(v27)) {
        var v53 float32 = 0.1
        _ = v53
    }
    fmt.Println(v3, v5)
    delete(v1, 908)
    v28 = 2147483647
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

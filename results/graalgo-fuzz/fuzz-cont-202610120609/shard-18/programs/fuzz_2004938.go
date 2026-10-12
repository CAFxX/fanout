package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 int) string {
    v4 := []int{}
    var v5 float32 = 1.5
    _ = p0
    _ = v4
    _ = v5
    return "hello"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    var v3 uint64 = 54
    fmt.Println(v3, v2)
    v6 := []int{1, -710}
    v7 := GenS1{F0: -2147483648}
    switch "\n\t" {
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "\n\t":
        var v8 string = ""
        _ = v8
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            v2 += (v2 & (v2 >> ((65535) & 63)))
            v3 = (v3 % ((v3) | 1))
            _ = fi9
        }
    }
    var v10 int32 = 0
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            v6 = append(v6, 0)
            _ = fi11
        }
        fmt.Println(v2)
    }
    v12 := GenS2{F0: 1000}
    var v13 string = ""
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v3)
        for fi15 := 0; fi15 < 3; fi15++ {
            v10 = v10
            _ = fi15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        switch (fi16 / ((128) | 1)) {
        case 128:
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := v1[65535]
                v3 = 908
                _ = fi17
                _ = v18
            }
        case -1:
            v13 = (("foo bar" + ("a" + ("a" + v13))) + "foo bar")
        default:
            if false {
                if ("foo bar" < v13) {
                    fmt.Println(v3)
                    v13 = ("" + (((v13 + v13) + v13) + ("\n\t" + (v13 + (((v13 + v13) + ("" + v13)) + ((v13 + ("" + v13)) + v13))))))
                    v6 = append(v6, fi16)
                }
                var v19 float32 = 3.14159
                _ = v19
            }
        }
        _ = fi16
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v13 = "hello"
        _ = fi20
    }
    fmt.Println(v2, v13, v2)
    var v21 string = "foo bar"
    if ((-2.25 == -2.25) || true) {
        fmt.Println(v2, v3)
        for fi22 := 0; fi22 < 7; fi22++ {
            v1[1] = true
            _ = fi22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            v2 -= (v2 & 100)
            _ = fi23
        }
    }
    var v24 float32 = 3.14159
    for fi25 := 0; fi25 < 1; fi25++ {
        switch v13 {
        case "\n\t":
            v26 := GenS3{F0: v3}
            _ = v26
        case "foo bar":
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS4{F0: int64(v2), F1: fi27, F2: v2}
                if true {
                    if false {
                        v24 -= (float32(((0.1 + 0.1) - 1e10)) - v24)
                    } else {
                        var v29 uint16 = 42
                        v29 = (v29 + v29)
                        var v30 float32 = -2.25
                        _ = v29
                        _ = v30
                    }
                    switch (fi25 % ((-317) | 1)) {
                    case 2147483647:
                        if false {
                            fmt.Println(v13)
                            fmt.Println(v10, v3)
                        } else {
                            v31 := GenS5{F0: 42}
                            var v32 uint16 = 110
                            _ = v31
                            _ = v32
                        }
                    case 193:
                        fmt.Println(fi27)
                    case -9223372036854775808:
                        v21 = v21
                    default:
                        v21 = "a"
                    }
                } else {
                    v21 = "\n\t"
                    v2 += v2
                    var v33 int64 = 851
                    if true {
                        v10 -= (v10 | -1)
                        v34 := GenS6{F0: (v33 ^ (v33 ^ v33)), F1: (0 <= 728), F2: ((v21 + (v13 + (v21 + (("\n\t" + (((v21 + (("foo bar" + v13) + v13)) + (v13 + v21)) + "")) + (v13 + (v21 + ("\n\t" + v13))))))) + (("foo bar" + v13) + ("\n\t" + "")))}
                        fmt.Println(fi27, v2, v21)
                        switch fi27 {
                        case -150:
                            var v35 uint8 = 127
                            _ = v35
                        case -187:
                            var v36 float32 = float32(float64(3.14159))
                            _ = v36
                        case 522:
                            fmt.Println(v24, fi27, fi27)
                        default:
                            if true {
                                var v37 float32 = 0.0
                                v38 := GenS4{F0: v33, F1: -1, F2: v2}
                                _ = v37
                                _ = v38
                            }
                        }
                        _ = v34
                    }
                    var v39 bool = (v10 > (v10 + v10))
                    if (v39 && v39) {
                        v6 = append(v6, (-(fi25)))
                        delete(v1, 0)
                        var v40 int16 = 32767
                        _ = v40
                    }
                    _ = v33
                    _ = v39
                }
                v41 := GenS2{F0: 906}
                _ = fi27
                _ = v28
                _ = v41
            }
        }
        _ = fi25
    }
    switch 9223372036854775807 {
    case 418:
        v3 = 0
    case -498:
        v13 = "foo bar"
    case -9223372036854775808:
        var v42 bool = false
        _ = v42
    default:
        v10 += v10
    }
    fmt.Println(v3, v13, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v6
    _ = v7
}

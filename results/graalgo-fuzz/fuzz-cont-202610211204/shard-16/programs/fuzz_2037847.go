package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint32; F1 int16; F2 string }
type GenS5 struct { F0 bool; F1 int32; F2 bool }
type GenS6 struct { F0 bool; F1 int64; F2 uint }
type GenS7 struct { F0 int8; F1 int }
func genFn0(p0 bool) uint {
    v4 := []int{-665, 127, -2147483648, -2147483648}
    if len(v4) > 0 {
        v4[0] = 32767
    }
    var v5 int8 = 127
    _ = len(v4)
    var v6 int16 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 424
}
func genFn1() uint8 {
    v7 := GenS1{F0: "foo bar", F1: 2}
    var v8 float64 = 0.0
    _ = v7
    _ = v8
    return 49
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 127
    var v3 int64 = 9223372036854775807
    for fi9 := 0; fi9 < 9; fi9++ {
        switch (fi9 * fi9) {
        case 9223372036854775807:
            v10 := GenS1{F0: "a", F1: 0}
            _ = v10
        case 133:
            fmt.Println(v2, v3, fi9)
        case 2147483647:
            if (18446744073709551615 <= 185) {
                v11 := []int{fi9, (-(fi9)), (fi9 % ((fi9) | 1))}
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: ((("" + "") + "\n\t") + ("\n\t" + "a")), F1: 2}
            if (2 == uint(v2)) {
                if false {
                    v1["foo bar"] = (("hello" + "a") + "foo bar")
                    v14 := GenS2{F0: 42, F1: 18446744073709551615, F2: v2}
                    if true {
                        if false {
                            fmt.Println(fi9)
                            if (true || (false || (!true))) {
                                v2 = v2
                            }
                            switch fi9 {
                            case 711:
                                if true {
                                    v2 = genFn1()
                                }
                            case 9223372036854775807:
                                v15 := GenS3{F0: 42, F1: 1}
                                _ = v15
                            case -1:
                                v2 = v2
                            default:
                                switch (("hello" + "a") + ("foo bar" + "foo bar")) {
                                case "":
                                    v16 := GenS4{F0: 825, F1: 1000, F2: "\n\t"}
                                    _ = v16
                                }
                            }
                        }
                    } else {
                        v3 = (v3 - (v3 ^ v3))
                    }
                    _ = v14
                } else {
                    if true {
                        switch "a" {
                        case "":
                            v2 += v2
                        default:
                            v17 := []int{fi12, 32767}
                            var v18 float32 = ((1e10 / 1.5) * float32(-2.25))
                            var v19 int64 = v3
                            var v20 float32 = float32((-2.25 + 0.1))
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    } else {
                        fmt.Println(fi12, v3)
                        v1[""] = ("foo bar" + (("" + (("a" + "") + (("foo bar" + "foo bar") + "hello"))) + "\n\t"))
                    }
                    switch -459 {
                    case -2147483648:
                        v21 := GenS5{F0: (-9 != -2147483648), F1: 42, F2: true}
                        _ = v21
                    case 2147483647:
                        switch 4294967295 {
                        case 32768:
                            fmt.Println(v3)
                        case 156:
                            fmt.Println(fi9)
                        case 127:
                            v3 -= (v3 + v3)
                        default:
                            if true {
                                v2 -= genFn1()
                                fmt.Println(fi12, fi9)
                            } else {
                                fmt.Println(fi12, fi9)
                                fmt.Println(v2, v2)
                            }
                            fmt.Println(v2, fi9, fi9)
                        }
                        v2 = v2
                    case 0:
                        v2 = 0
                    default:
                        if false {
                            v3 += -1
                            v22 := GenS6{F0: (1e10 < 1.5), F1: v3, F2: 315}
                            _ = v22
                        }
                    }
                    v23 := GenS1{F0: "a", F1: 65535}
                    v3 = v3
                    if (1.5 == 0.0) {
                        fmt.Println(v2)
                        var v24 string = (("" + "a") + ("\n\t" + (("foo bar" + "") + ("\n\t" + "a"))))
                        v25 := GenS7{F0: 42, F1: fi9}
                        _ = v24
                        _ = v25
                    }
                    _ = v23
                }
                var v26 uint16 = 32767
                _ = v26
            } else {
                fmt.Println(fi9)
                v27 := GenS5{F0: false, F1: -1, F2: false}
                _ = v27
            }
            v28 := GenS1{F0: "foo bar", F1: 0}
            _ = fi12
            _ = v13
            _ = v28
        }
        v29 := GenS6{F0: (935 > 679), F1: (v3 ^ v3), F2: 65535}
        _ = fi9
        _ = v29
    }
    var v30 uint = 18446744073709551615
    v30 += (^(v30))
    var v31 uint64 = 237
    if false {
        v32 := v1[("foo bar" + "foo bar")]
        var v33 uint16 = 0
        v30 = v30
        _ = v32
        _ = v33
    }
    v34 := GenS7{F0: 127, F1: 32768}
    v35 := GenS2{F0: v30, F1: 4294967295, F2: (v2 << ((v30) & 63))}
    var v36 float64 = ((0.1 * (3.14159 * 0.0)) - 3.14159)
    var v37 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}

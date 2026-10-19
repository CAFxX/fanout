package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
type GenS6 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint8, p1 int64, p2 bool) string {
    v4 := GenS1{F0: p2}
    fmt.Println(p1, p2, p0)
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5, p0)
        p0 = p0
        _ = fi5
    }
    fmt.Println(p1)
    switch 42 {
    case 9223372036854775807:
        for fi6 := 0; fi6 < 5; fi6++ {
            p2 = p2
            _ = fi6
        }
    case -126:
        p2 = true
    case 32767:
        p1 -= (p1 | (p1 - -1))
    default:
        for fi7 := 0; fi7 < 8; fi7++ {
            p1 += p1
            var v8 int8 = 0
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}

func main() {
    var v1 float32 = (3.14159 - -2.25)
    var v2 float32 = v1
    var v3 int32 = 932
    var v9 uint16 = 0
    fmt.Println(v3)
    fmt.Println(v1, v1, v9)
    fmt.Println(v9, v3)
    v10 := []int{0}
    _ = len(v10)
    var v11 uint64 = 1
    if true {
        var v12 string = "a"
        _ = v12
    }
    var v13 float32 = v1
    var v14 int16 = 621
    if ((795 != 1) && (v14 <= (v14 / (((v14 | (-(v14)))) | 1)))) {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS2{F0: -1, F1: -110}
            _ = fi15
            _ = v16
        }
        switch ("foo bar" + "") {
        case "":
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int = (fi17 ^ fi18)
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        case "a":
            v2 -= 0.0
        case "\n\t":
            v9 = v9
        default:
            for fi20 := 0; fi20 < 10; fi20++ {
                if len(v10) > 0 {
                    v10[0] = (-(fi20))
                }
                fmt.Println(v14, v13)
                for fi21 := 0; fi21 < 17; fi21++ {
                    v22 := GenS3{F0: "foo bar"}
                    switch "hello" {
                    case "":
                        v10 = append(v10, 1)
                    case "foo bar":
                        if false {
                            if true {
                                v10 = append(v10, (fi21 >> ((865) & 63)))
                                var v23 string = ""
                                fmt.Println(v14, v2)
                                _ = v23
                            } else {
                                switch (fi21 & 127) {
                                case 435:
                                    v3 -= (v3 / (((v3 * v3)) | 1))
                                case 127:
                                    v24 := GenS2{F0: -1, F1: int8(v9)}
                                    _ = v24
                                default:
                                    v10 = append(v10, fi21)
                                }
                            }
                            v11 = (v11 * v11)
                        } else {
                            v13 -= float32(0.0)
                            var v25 int32 = 1
                            if false {
                                v26 := GenS3{F0: ""}
                                v11 -= (v11 << ((65535) & 63))
                                v14 += v14
                                _ = v26
                            } else {
                                v1 = v13
                                var v27 uint = 23
                                _ = v27
                            }
                            _ = v25
                        }
                        v28 := GenS4{F0: v14, F1: 70}
                        var v29 int8 = 1
                        _ = v28
                        _ = v29
                    case "hello":
                        v9 -= v9
                    default:
                        v30 := GenS5{F0: 979, F1: "", F2: (v14 - -206)}
                        _ = v30
                    }
                    v3 = (v3 >> ((uint(v11)) & 63))
                    v9 = v9
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        }
    } else {
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 int32 = v3
            var v33 int = 2147483647
            var v34 int64 = 1
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS6{F0: 102, F1: 191}
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}

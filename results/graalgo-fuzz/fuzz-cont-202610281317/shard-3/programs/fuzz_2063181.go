package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    var v1 uint16 = 2
    var v2 int8 = -128
    v3 := make(map[string]string)
    var v4 int32 = 1
    var v5 string = "hello"
    var v6 uint8 = 255
    var v7 int8 = (v2 << ((0) & 63))
    v8 := GenS1{F0: "a", F1: uint32(v1)}
    if false {
        v4 = v4
    } else {
        if ((v6 * 42) < (v6 ^ uint8(v1))) {
            if false {
                v7 += v2
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 int = (fi9 * 0)
                    v11 := v3[((v5 + (((("hello" + v5) + ((v5 + "foo bar") + v5)) + v5) + v5)) + v5)]
                    v5 = (v5 + "foo bar")
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                switch "a" {
                case "foo bar":
                    if (-109 >= (v7 % ((3) | 1))) {
                        v7 = v7
                        v12 := v3[v5]
                        switch v5 {
                        case "a":
                            fmt.Println(v5)
                        case "\n\t":
                            v5 = "foo bar"
                        default:
                            v13 := []int{429}
                            _ = v13
                        }
                        _ = v12
                    }
                default:
                    fmt.Println(v5)
                }
            } else {
                if true {
                    v14 := []int{367, 32768, 368}
                    v4 -= (v4 + v4)
                    switch v5 {
                    case "foo bar":
                        switch -154 {
                        case 1:
                            var v15 int = -816
                            _ = v15
                        case -9223372036854775808:
                            fmt.Println(v7, v2)
                        }
                    case "\n\t":
                        for fi16 := 0; fi16 < 14; fi16++ {
                            if (((-2.25 - ((float32((3.14159 * float64(1.5))) + float32(3.14159)) / (float32(1.5) * -2.25))) / (1e10 - (0.0 / 1.0))) == 1e10) {
                                v4 = 0
                            }
                            _ = len(v3)
                            v17 := v3[v5]
                            _ = fi16
                            _ = v17
                        }
                    default:
                        if (1000 <= 1) {
                            var v18 string = (v5 + v5)
                            _ = len(v3)
                            _ = v18
                        }
                    }
                    _ = v14
                } else {
                    v4 += 604
                    switch 2147483647 {
                    case 935:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 int8 = (^(v7))
                            v7 = -111
                            if true {
                                if true {
                                    fmt.Println(v2)
                                }
                                if true {
                                    if true {
                                        _ = len(v3)
                                        for fi21 := 0; fi21 < 6; fi21++ {
                                            v4 -= -2147483648
                                            fmt.Println(v20)
                                            _ = fi21
                                        }
                                    } else {
                                        v2 -= (v2 >> ((2) & 63))
                                        for fi22 := 0; fi22 < 18; fi22++ {
                                            v2 = v2
                                            var v23 uint16 = (v1 << ((uint(v6)) & 63))
                                            _ = fi22
                                            _ = v23
                                        }
                                        v1 += (v1 >> ((165) & 63))
                                    }
                                }
                            }
                            _ = fi19
                            _ = v20
                        }
                    case 4294967295:
                        if true {
                            fmt.Println(v1, v5)
                        } else {
                            var v24 bool = (42 != 2147483648)
                            _ = v24
                        }
                    case -772:
                        v5 = "hello"
                    }
                }
                if false {
                    var v25 int8 = 1
                    _ = v25
                }
            }
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                v2 = v2
                v5 = v5
                _ = fi26
            }
            var v27 int32 = -858
            v28 := []int{-9223372036854775808}
            _ = v27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int32 = v4
        for fi31 := 0; fi31 < 16; fi31++ {
            v7 = (v2 / ((v7) | 1))
            _ = fi31
        }
        v32 := GenS2{F0: (v4 * 42)}
        _ = fi29
        _ = v30
        _ = v32
    }
    v33 := GenS2{F0: v4}
    for fi34 := 0; fi34 < 4; fi34++ {
        fmt.Println(v4, v6, v4)
        v1 = v1
        v35 := GenS1{F0: ("hello" + "a"), F1: 1}
        _ = fi34
        _ = v35
    }
    v2 += v2
    v36 := GenS3{F0: (v4 * v4), F1: (-400 == 32768)}
    var v37 int = 42
    v38 := GenS1{F0: "a", F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

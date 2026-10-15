package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
type GenS3 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 int8, p1 uint) int8 {
    v5 := GenS1{F0: 0, F1: ("\n\t" + "hello")}
    if (2 <= 1) {
        var v6 bool = true
        p1 = 0
        switch "hello" {
        case "":
            if false {
                v7 := []int{-1, 354, 127, -97}
                _ = v7
            }
        case "hello":
            v6 = v6
        default:
            fmt.Println(p0, p1, p0)
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return (^(p0))
}
func genFn1() uint32 {
    switch "" {
    case "a":
        var v8 uint32 = 0
        _ = v8
    default:
        var v9 string = ("foo bar" + "")
        _ = v9
    }
    v10 := GenS1{F0: 65535, F1: ("a" + "")}
    _ = v10
    return 307
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 uint64 = 0
    var v4 float64 = ((float64(0.0) + 3.14159) / 1.0)
    fmt.Println(v3, v4)
    switch ("\n\t" + "") {
    case "hello":
        v3 = v3
    case "\n\t":
        var v11 uint32 = 4294967295
        _ = v11
    case "foo bar":
        v12 := v2[4294967295]
        _ = v12
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 127 {
    case -510:
        switch "foo bar" {
        case "a":
            _ = len(v1)
        case "hello":
            var v13 uint32 = 65535
            _ = v13
        }
    case -72:
        var v14 uint8 = 0
        _ = v14
    case -9223372036854775808:
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 float32 = (1e10 / (1e10 + (float32(float64((1.5 - (3.14159 / 0.1)))) - -2.25)))
            v4 += float64((((0.0 + v16) - v16) / v16))
            _ = fi15
            _ = v16
        }
    default:
        if (32767 == 709) {
            v17 := GenS1{F0: 462, F1: "a"}
            v4 += 1e10
            switch "" {
            case "hello":
                fmt.Println(v3)
            case "":
                if true {
                    var v18 bool = (42 <= -1)
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = v18
                } else {
                    fmt.Println(v3, v4)
                }
            case "a":
                fmt.Println(v4)
            default:
                var v19 int16 = -1
                _ = v19
            }
            _ = v17
        } else {
            for fi20 := 0; fi20 < 20; fi20++ {
                switch fi20 {
                case 78:
                    v21 := GenS2{F0: -696, F1: ("hello" + "foo bar"), F2: 65535}
                    _ = v21
                case -460:
                    switch (^(fi20)) {
                    case 687:
                        v4 = ((v4 - v4) + v4)
                    case -302:
                        if (genFn1() > 774) {
                            var v22 bool = false
                            v4 -= 3.14159
                            if v22 {
                                v23 := GenS2{F0: 1, F1: ("\n\t" + "\n\t"), F2: 379}
                                _ = v23
                            }
                            _ = v22
                        }
                    case -954:
                        _ = len(v1)
                    }
                case 128:
                    var v24 bool = false
                    _ = v24
                }
                _ = fi20
            }
            var v25 int = -908
            fmt.Println(v25)
            _ = v25
        }
    }
    var v26 uint32 = 65535
    var v27 int8 = 1
    switch -666 {
    case 2147483647:
        v28 := GenS2{F0: 540, F1: "a", F2: (^(v26))}
        _ = v28
    case 32767:
        var v29 int8 = (v27 - genFn0((v27 + -114), 792))
        _ = v29
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    v30 := GenS1{F0: 2, F1: ("foo bar" + "foo bar")}
    v4 -= ((v4 / (v4 / v4)) - (v4 - v4))
    v31 := GenS3{F0: -9223372036854775808, F1: 2147483647, F2: (-(v27))}
    var v32 string = "a"
    fmt.Println(v32)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}

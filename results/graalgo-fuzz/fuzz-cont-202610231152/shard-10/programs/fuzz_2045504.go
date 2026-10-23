package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 string; F1 int32; F2 int64 }
func genFn0(p0 uint32) uint {
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 -= (p0 ^ (p0 << ((255) & 63)))
        _ = fi7
    }
    p0 = (p0 >> ((42) & 63))
    p0 += p0
    _ = p0
    return 42
}
func genFn1() uint64 {
    var v8 uint64 = 2
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(v8)
    default:
        var v9 int32 = 345
        _ = v9
    }
    v8 = v8
    if (!false) {
        var v10 string = "hello"
        _ = v10
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int8 = 127
            var v13 int8 = 42
            switch 128 {
            case 32768:
                v12 += v13
            case -1:
                v14 := GenS1{F0: 42, F1: 1, F2: 9223372036854775807}
                _ = v14
            case -2147483648:
                switch "a" {
                case "a":
                    var v15 float32 = 0.0
                    _ = v15
                default:
                    var v16 uint8 = 2
                    _ = v16
                }
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := []int{}
            if false {
                if false {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        if false {
                            v8 += v8
                        }
                        if len(v18) > 0 {
                            v18[0] = (^(fi17))
                        }
                        var v20 uint16 = 335
                        _ = len(v18)
                        v8 = (^(v8))
                        _ = fi19
                        _ = v20
                    }
                    _ = len(v18)
                }
                v8 -= v8
            }
            var v21 uint16 = 1
            _ = fi17
            _ = v18
            _ = v21
        }
    }
    _ = v8
    return v8
}

func main() {
    var v1 uint8 = 0
    v2 := []string{"\n\t", "foo bar", "hello", "hello"}
    v3 := []string{}
    v4 := make(map[string]int)
    var v5 float64 = 0.0
    var v6 uint16 = 2
    v22 := GenS1{F0: -128, F1: 2, F2: int(v6)}
    if false {
        fmt.Println(v1, v1)
    }
    v1 -= (^(v1))
    if (2 > 42) {
        if true {
            var v23 int32 = -917
            v24 := GenS1{F0: -1, F1: 649, F2: 32768}
            for fi25 := 0; fi25 < 10; fi25++ {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi25
            }
            _ = v23
            _ = v24
        }
    } else {
        v6 = 65535
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    v1 = (v1 * 29)
    var v26 float64 = v5
    fmt.Println(v26, v5, v6)
    v1 = (v1 << ((579) & 63))
    v27 := GenS2{F0: 925, F1: (("a" + "hello") + "\n\t"), F2: "\n\t"}
    if true {
        v28 := GenS3{F0: 255, F1: 18446744073709551615, F2: 0}
        _ = v28
    } else {
        v6 = v6
        var v29 float64 = v5
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v3 = append(v3, ("" + "foo bar"))
        switch "a" {
        case "hello":
            v5 += (3.14159 + (1e10 / v26))
        case "":
            v26 = v5
        case "a":
            v31 := GenS4{F0: true, F1: -128}
            _ = v31
        default:
            fmt.Println(v6, v6)
        }
        if (-60 < -1) {
            v32 := GenS5{F0: ("\n\t" + ("" + "foo bar")), F1: 42, F2: 9223372036854775807}
            _ = v32
        } else {
            v6 = v6
        }
        _ = fi30
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        if len(v3) > 0 {
            v3[0] = (("\n\t" + ("" + ("foo bar" + "a"))) + "")
        }
        v26 -= v5
        _ = fi33
    }
    _ = len(v3)
    v26 += 1.5
    var v34 uint8 = (v1 + 101)
    _ = len(v4)
    switch 32768 {
    case 1:
        v5 -= 1.5
    case 470:
        _ = len(v4)
    default:
        for fi35 := 0; fi35 < 10; fi35++ {
            switch "a" {
            case "a":
                fmt.Println(fi35, v1, v6)
            default:
                v36 := GenS1{F0: -1, F1: 0, F2: (-(fi35))}
                _ = v36
            }
            var v37 float32 = 1.5
            _ = fi35
            _ = v37
        }
    }
    switch "a" {
    case "a":
        for fi38 := 0; fi38 < 6; fi38++ {
            v39 := GenS2{F0: -32768, F1: ((("hello" + "a") + "") + ("\n\t" + "\n\t")), F2: "hello"}
            v6 = v6
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

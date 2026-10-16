package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 string; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint }
func genFn0(p0 string, p1 string) string {
    v4 := GenS1{F0: 581, F1: 295, F2: 925}
    v5 := []int{4294967295}
    fmt.Println(p0, p0)
    switch 32768 {
    case 2147483647:
        v6 := GenS2{F0: (0 != 294)}
        _ = v6
    case 32768:
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                if (496 == 0) {
                    p0 = "a"
                    switch fi7 {
                    case -157:
                        for fi8 := 0; fi8 < 7; fi8++ {
                            var v9 int16 = 32767
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi8
                            _ = v9
                        }
                    case 267:
                        var v10 bool = (false || true)
                        _ = v10
                    }
                } else {
                    v11 := GenS3{F0: 255, F1: "foo bar", F2: (1 <= 33)}
                    _ = v11
                }
                var v12 string = ""
                _ = fi7
                _ = v12
            }
            p1 = p0
        }
    }
    if (2 > 168) {
        v13 := GenS4{F0: 255}
        switch ("a" + ("foo bar" + p1)) {
        case "":
            var v14 bool = (42 >= 563)
            _ = v14
        case "\n\t":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v13
    } else {
        var v15 float32 = -2.25
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = -2.25
    var v3 int16 = 1000
    _ = len(v1)
    if (!(0 >= 162)) {
        fmt.Println(v2, v3)
        v2 += v2
    }
    v3 -= v3
    v2 += -2.25
    v3 = -32768
    v16 := v1[("hello" + "hello")]
    v3 -= (v3 / (((v3 >> ((400) & 63))) | 1))
    v17 := GenS5{F0: "\n\t", F1: 65535, F2: -128}
    v1["a"] = v16
    var v18 uint8 = 0
    for fi19 := 0; fi19 < 18; fi19++ {
        switch ("a" + v16) {
        case "a":
            v20 := []int{fi19, -1, fi19, 127}
            _ = v20
        case "foo bar":
            v21 := GenS3{F0: 785, F1: (v16 + "foo bar"), F2: false}
            _ = v21
        case "":
            for fi22 := 0; fi22 < 15; fi22++ {
                switch -208 {
                case 42:
                    switch ((v16 + "") + v16) {
                    case "hello":
                        v16 = v16
                        v23 := GenS4{F0: 2}
                        if true {
                            var v24 uint64 = 18446744073709551615
                            v25 := GenS5{F0: (v16 + (("foo bar" + "hello") + "\n\t")), F1: 42, F2: -118}
                            _ = v24
                            _ = v25
                        } else {
                            v26 := []int{fi22, fi22, fi19}
                            var v27 int64 = -1
                            var v28 int8 = 127
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = v23
                    case "foo bar":
                        if (v3 > v3) {
                            v2 -= (float32(1e10) - v2)
                        } else {
                            var v29 bool = false
                            v16 = (v16 + (((((("hello" + "foo bar") + "hello") + v16) + "a") + v16) + "a"))
                            _ = v29
                        }
                    case "":
                        var v30 int32 = -2147483648
                        _ = v30
                    default:
                        fmt.Println(fi19, v3, v2)
                    }
                }
                v18 -= (v18 ^ (v18 / (((v18 - (v18 ^ (v18 % ((127) | 1))))) | 1)))
                var v31 uint16 = 32767
                var v32 uint32 = 2147483648
                _ = fi22
                _ = v31
                _ = v32
            }
        default:
            v18 += (v18 * 1)
        }
        v1["a"] = ""
        _ = fi19
    }
    v3 -= (v3 * (-(v3)))
    _ = len(v1)
    if (v3 >= 82) {
        var v33 int8 = 42
        if (!((((v3 - (v3 % ((42) | 1))) >= (v3 ^ (v3 ^ (v3 << ((42) & 63))))) || false) || false)) {
            switch 32768 {
            case 32767:
                v34 := GenS2{F0: (((true || (uint(v18) < 4294967295)) || false) && (871 == 65535))}
                _ = v34
            case -1:
                var v35 uint8 = (v18 | v18)
                _ = v35
            }
            fmt.Println(v33)
            v33 -= 0
        }
        _ = v33
    } else {
        delete(v1, v16)
        var v36 bool = (362 <= 249)
        var v37 uint = 255
        _ = v36
        _ = v37
    }
    v38 := GenS6{F0: 1}
    fmt.Println(v3, v18, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
}

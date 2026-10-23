package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint64) string {
    var v6 string = "\n\t"
    v7 := []int{-766, 32767}
    p0 -= p0
    p0 -= p0
    _ = p0
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 int8 = 100
    var v2 float64 = (1e10 + 3.14159)
    var v3 uint32 = 620
    v4 := []string{("\n\t" + "\n\t"), ("" + "hello")}
    v5 := make(map[string]bool)
    if (826 < 2147483647) {
        if true {
            if false {
                if false {
                    for fi8 := 0; fi8 < 13; fi8++ {
                        v5["a"] = false
                        _ = len(v5)
                        fmt.Println(v1, v1, v2)
                        _ = fi8
                    }
                }
                v9 := GenS1{F0: 127, F1: 34}
                _ = v9
            } else {
                if false {
                    fmt.Println(v1, v1, v1)
                    v5["\n\t"] = false
                }
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint8 = 0
                    if (false || (!true)) {
                        if false {
                            for fi12 := 0; fi12 < 17; fi12++ {
                                v1 += v1
                                _ = fi12
                            }
                        }
                        if (false || true) {
                            if ((v1 >> ((2) & 63)) == (v1 - 0)) {
                                v13 := v5[""]
                                switch "a" {
                                case "hello":
                                    v13 = (2 != uint64(v11))
                                default:
                                    _ = len(v4)
                                }
                                _ = v13
                            } else {
                                _ = len(v4)
                                v14 := GenS1{F0: (v1 ^ v1), F1: 44}
                                switch fi10 {
                                case 4294967295:
                                    for fi15 := 0; fi15 < 16; fi15++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi15
                                    }
                                case -2147483648:
                                    v4 = append(v4, (("" + ("a" + "a")) + ("foo bar" + "hello")))
                                default:
                                    v3 += v3
                                }
                                _ = v14
                            }
                            if (42 < 18446744073709551615) {
                                v16 := GenS1{F0: v1, F1: -722}
                                _ = v16
                            }
                        }
                        var v17 uint = 249
                        _ = v17
                    }
                    v5[""] = (-9223372036854775808 <= 0)
                    _ = fi10
                    _ = v11
                }
                v18 := GenS1{F0: v1, F1: -1}
                _ = v18
            }
        } else {
            var v19 string = ("" + "foo bar")
            _ = v19
        }
        v3 += (^(v3))
    } else {
        fmt.Println(v3, v2)
    }
    switch 442 {
    case 4294967295:
        v20 := GenS2{F0: 1, F1: 0}
        _ = v20
    case -2147483648:
        v21 := GenS3{F0: 0, F1: -1}
        _ = v21
    case 127:
        _ = len(v4)
    default:
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int16 = 1000
            v5[(("a" + "hello") + "a")] = (((v3 >> ((255) & 63)) >= (v3 - (v3 - 2147483647))) || ((false && (0.1 >= -2.25)) && false))
            v2 = 0.1
            _ = fi22
            _ = v23
        }
    }
    v1 += (^(v1))
    v1 = v1
    if (1 <= 1) {
        fmt.Println(v2, v1)
    }
    v2 += ((float64((3.14159 + float32(-2.25))) - v2) / float64(1e10))
    if (!(true && (true && ("foo bar" > "foo bar")))) {
        v1 -= v1
        if (203 == 2147483647) {
            v2 -= v2
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int32 = -1
        _ = fi24
        _ = v25
    }
    v5["hello"] = false
    v26 := GenS2{F0: 499, F1: -2147483648}
    v27 := v5["foo bar"]
    var v28 string = "hello"
    var v29 int16 = -32768
    for fi30 := 0; fi30 < 3; fi30++ {
        switch "hello" {
        case "hello":
            var v31 int32 = -444
            _ = v31
        case "":
            v29 += v29
        }
        v1 = v1
        _ = fi30
    }
    fmt.Println(v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

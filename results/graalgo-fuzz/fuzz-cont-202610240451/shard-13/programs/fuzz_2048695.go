package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int8 }

func main() {
    var v1 bool = true
    var v2 int8 = -1
    var v3 uint32 = 737
    var v4 int64 = 9223372036854775807
    v5 := []int{128, 535}
    v1 = v1
    v4 += v4
    switch "" {
    case "\n\t":
        var v6 int64 = (v4 / (((v4 % ((522) | 1))) | 1))
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 4294967295, F1: 32768}
            if (v2 == v2) {
                v2 = (v2 * (v2 << ((255) & 63)))
            } else {
                if v1 {
                    v9 := GenS1{F0: 2, F1: fi7}
                    var v10 uint16 = 255
                    switch "hello" {
                    case "hello":
                        for fi11 := 0; fi11 < 16; fi11++ {
                            fmt.Println(fi7)
                            if v1 {
                                if v1 {
                                    fmt.Println(fi11, v2)
                                }
                            } else {
                                v12 := GenS2{F0: (v2 - (^(v2)))}
                                if true {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if v1 {
                                        v13 := GenS1{F0: 462, F1: 9223372036854775807}
                                        v14 := GenS2{F0: -1}
                                        v10 = 42
                                        if len(v5) > 0 {
                                            v5[0] = 368
                                        }
                                        fmt.Println(v10, v10)
                                        switch "" {
                                        case "a":
                                            var v15 uint16 = v10
                                            _ = v15
                                        case "":
                                            if (612 >= 1) {
                                                v5 = append(v5, int(v2))
                                                var v16 bool = v1
                                                _ = v16
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = fi11
                                                }
                                            }
                                        case "foo bar":
                                            v3 = (v3 & v3)
                                        }
                                        _ = v13
                                        _ = v14
                                    }
                                    var v17 int32 = 193
                                    _ = v17
                                }
                                fmt.Println(v2, v2, v1)
                                _ = v12
                            }
                            var v18 uint = 2
                            _ = fi11
                            _ = v18
                        }
                    }
                    _ = v10
                    _ = v9
                }
            }
            _ = fi7
            _ = v8
        }
    }
    var v19 uint64 = 2
    var v20 int16 = int16(v2)
    var v21 float64 = float64(1.5)
    var v22 uint8 = 127
    v23 := GenS3{F0: (((false && v1) || (("\n\t" + "foo bar") == "a")) || v1), F1: v2}
    v1 = false
    if len(v5) > 0 {
        v5[0] = 306
    }
    v3 -= (v3 / ((v3) | 1))
    switch ("a" + "hello") {
    case "":
        switch 9223372036854775807 {
        case -684:
            var v24 float64 = v21
            _ = v24
        case 4294967295:
            var v25 uint = 600
            _ = v25
        case 75:
            for fi26 := 0; fi26 < 1; fi26++ {
                if len(v5) > 0 {
                    v5[0] = fi26
                }
                var v27 string = ("a" + ("\n\t" + ""))
                if ((!((-(v2)) > 1)) && ((v20 * v20) <= v20)) {
                    _ = len(v5)
                    v5 = append(v5, fi26)
                }
                _ = fi26
                _ = v27
            }
        }
    case "a":
        v1 = (!v1)
    }
    v28 := GenS3{F0: ((!(v2 > 42)) || ((1e10 + -2.25) < float32((v21 * float64((0.0 - float32(float64(1e10)))))))), F1: (v2 % (((v2 ^ (v2 | v2))) | 1))}
    var v29 string = (("" + "foo bar") + "")
    switch 32767 {
    case 4294967295:
        v30 := GenS1{F0: (v19 & 4294967295), F1: 2147483647}
        _ = v30
    case 2147483647:
        var v31 string = ("foo bar" + v29)
        _ = v31
    case -846:
        v20 = (v20 | v20)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

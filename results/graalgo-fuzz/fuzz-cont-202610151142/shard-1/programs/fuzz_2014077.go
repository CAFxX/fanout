package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int = 673
    var v2 uint = 65535
    var v3 uint64 = 363
    v4 := make(map[int]int)
    var v5 float32 = (1.5 - -2.25)
    for fi6 := 0; fi6 < 18; fi6++ {
        v1 = fi6
        _ = fi6
    }
    if (65535 <= 962) {
        v2 -= v2
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v1 = fi7
        if true {
            if true {
                if (("foo bar" + "hello") == "a") {
                    v3 = (v3 * (v3 & 4294967295))
                    v8 := v4[(-(v1))]
                    for fi9 := 0; fi9 < 10; fi9++ {
                        var v10 uint = 255
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                } else {
                    v2 += (^(v2))
                    var v11 int = 127
                    _ = v11
                }
            }
        } else {
            fmt.Println(v5, v2, v2)
            v12 := GenS1{F0: int32(v2), F1: 42}
            if (3.14159 >= v5) {
                v13 := GenS1{F0: 0, F1: 206}
                v2 = v2
                _ = v13
            }
            _ = v12
        }
        _ = fi7
    }
    var v14 int8 = 100
    v4[v1] = 0
    v3 = 4294967295
    var v15 float64 = 0.1
    v16 := GenS2{F0: 32768, F1: -17}
    var v17 uint8 = 146
    switch 2147483647 {
    case -640:
        fmt.Println(v1)
    case 960:
        v5 += v5
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(v15, v5)
        v1 = -2147483648
        _ = fi18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(v14, v1, v14)
        for fi20 := 0; fi20 < 17; fi20++ {
            if ((false || true) || false) {
                v21 := GenS1{F0: 1, F1: -1}
                if (!true) {
                    v3 -= 18446744073709551615
                    v22 := GenS3{F0: -1}
                    var v23 int16 = 32767
                    _ = v22
                    _ = v23
                } else {
                    v4[(fi19 | v1)] = (fi20 % (((v1 | (v1 ^ fi20))) | 1))
                    if (!false) {
                        _ = len(v4)
                        if ("" >= "hello") {
                            var v24 bool = false
                            v25 := []int{fi19, (fi19 >> ((v2) & 63)), 4294967295}
                            v2 = v2
                            _ = v24
                            _ = v25
                        } else {
                            if true {
                                switch (v1 & fi19) {
                                case -226:
                                    v26 := []int{fi20}
                                    _ = v26
                                }
                                switch "hello" {
                                case "":
                                    v27 := GenS4{F0: v17}
                                    _ = v27
                                case "a":
                                    v28 := []int{}
                                    _ = v28
                                }
                                v3 += v3
                                v29 := GenS4{F0: v17}
                                _ = v29
                            }
                            switch "" {
                            case "":
                                v30 := []int{(fi19 * v1)}
                                v3 -= v3
                                v2 = v2
                                _ = v30
                            }
                            if (false || false) {
                                v1 -= (fi20 * 32768)
                            } else {
                                if (((v14 - (v14 * int8(v17))) != v14) || (!false)) {
                                    fmt.Println(v3, fi19, v3)
                                    v31 := GenS3{F0: -9223372036854775808}
                                    v32 := []int{(-(fi19)), v1, (fi19 >> ((uint(v17)) & 63)), fi19}
                                    _ = v31
                                    _ = v32
                                }
                                v15 -= float64((0.1 - v5))
                                fmt.Println(fi20, v14)
                            }
                        }
                        v33 := []int{}
                        switch -9223372036854775808 {
                        case -502:
                            _ = len(v4)
                        case 0:
                            v3 = (v3 + 18446744073709551615)
                        }
                        _ = v33
                    } else {
                        v34 := GenS3{F0: 282}
                        _ = v34
                    }
                    var v35 int = -462
                    _ = v35
                }
                _ = v21
            } else {
                v36 := v4[(fi20 << ((4294967295) & 63))]
                _ = v36
            }
            v5 += (v5 - 0.1)
            _ = fi20
        }
        _ = fi19
    }
    switch 4294967295 {
    case -1:
        v5 = v5
    default:
        v17 = v17
    }
    v2 -= v2
    v37 := []int{int(v3)}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}

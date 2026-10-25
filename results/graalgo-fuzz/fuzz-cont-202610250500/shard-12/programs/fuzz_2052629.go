package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint32; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int64; F2 uint64 }

func main() {
    var v1 float32 = 1e10
    var v2 float32 = float32((float64(v1) + 3.14159))
    v3 := []int{}
    var v4 int64 = 83
    v5 := []int{-1}
    if ("hello" != "foo bar") {
        fmt.Println(v4)
    } else {
        v6 := GenS1{F0: true, F1: 1, F2: 0}
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint16 = 0
            v8 = v8
            if (32767 >= (v8 & 890)) {
                fmt.Println(v4)
                if true {
                    var v9 uint16 = v8
                    var v10 int8 = -128
                    _ = len(v5)
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: (0 >= 849), F1: 2, F2: 119}
                    _ = fi11
                    _ = v12
                }
            } else {
                switch "a" {
                case "hello":
                    v13 := GenS2{F0: 2147483647}
                    _ = v13
                case "foo bar":
                    var v14 uint16 = 255
                    _ = v14
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 uint32 = 65535
                    switch "\n\t" {
                    case "hello":
                        if ((true && true) || ((false && (-1 <= 51)) || (31 == 1))) {
                            v17 := GenS3{F0: fi7, F1: 2, F2: false}
                            _ = v17
                        } else {
                            switch 579 {
                            case 32767:
                                switch fi15 {
                                case 2147483647:
                                    var v18 bool = false
                                    _ = v18
                                }
                            case 1:
                                v1 = v2
                            default:
                                if false {
                                    v16 -= (v16 >> ((191) & 63))
                                }
                            }
                        }
                    default:
                        if true {
                            v2 -= (v2 * (v1 - v1))
                        }
                    }
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi7
            _ = v8
        }
        var v19 int32 = -1
        _ = v19
        _ = v6
    }
    if true {
        var v20 int = 0
        _ = v20
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            v5 = append(v5, fi21)
            for fi22 := 0; fi22 < 4; fi22++ {
                if ("\n\t" < ("" + "\n\t")) {
                    var v23 int = 4294967295
                    if true {
                        if true {
                            fmt.Println(v2, fi22, v23)
                        } else {
                            switch "hello" {
                            case "hello":
                                fmt.Println(v2)
                            }
                        }
                    } else {
                        fmt.Println(fi21, v4, fi21)
                        v1 -= ((v2 / v2) + float32((0.1 * 3.14159)))
                        v2 += float32((-2.25 + 1.5))
                        v23 -= (fi22 + (fi21 | (fi22 - v23)))
                    }
                    switch (-(v23)) {
                    case 0:
                        v24 := GenS4{F0: 4294967295}
                        _ = v24
                    case 32767:
                        v3 = append(v3, v23)
                        fmt.Println(v4, v2, v1)
                        v1 += (float32(1.5) + v1)
                        switch (v23 >> ((2) & 63)) {
                        case 988:
                            _ = len(v5)
                        }
                    }
                    v3 = append(v3, (-(fi22)))
                    _ = v23
                } else {
                    v5 = append(v5, fi22)
                    if len(v5) > 0 {
                        v5[0] = (fi21 | fi21)
                    }
                    v2 = v2
                }
                if true {
                    v2 = 1.5
                    v25 := GenS1{F0: (true && (32767 >= 32767)), F1: 2, F2: 127}
                    _ = v25
                } else {
                    v2 = (0.0 - v1)
                    fmt.Println(fi21, fi22)
                    var v26 float64 = (float64(((v1 + v1) + v1)) / float64(v1))
                    _ = v26
                }
                v4 += v4
                v4 -= v4
                switch fi21 {
                case -1:
                    var v27 int8 = -1
                    _ = v27
                default:
                    v3 = append(v3, 931)
                }
                _ = fi22
            }
            v28 := GenS5{F0: ("a" + "hello"), F1: 0, F2: 4294967295}
            _ = fi21
            _ = v28
        }
    }
    if (42 > 4294967295) {
        v29 := GenS5{F0: "\n\t", F1: (-(v4)), F2: 0}
        v30 := GenS5{F0: "foo bar", F1: v4, F2: 42}
        _ = v29
        _ = v30
    }
    var v31 uint8 = 116
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 -= -2.25
    fmt.Println(v4, v1, v1)
    var v32 int16 = 765
    v3 = append(v3, 4294967295)
    var v33 string = "hello"
    var v34 float64 = 1e10
    v2 -= (float32((v34 / 3.14159)) * v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}

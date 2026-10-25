package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 bool, p1 uint32, p2 int) int64 {
    p0 = p0
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 uint64 = 4294967295
        _ = fi4
        _ = v5
    }
    if p0 {
        switch ("hello" + ("\n\t" + "hello")) {
        case "hello":
            v6 := GenS1{F0: 42, F1: "\n\t"}
            _ = v6
        }
    }
    p0 = (!((0.1 * ((0.0 * 3.14159) + 3.14159)) <= 1.5))
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v7 int32 = 42
    var v8 int8 = -128
    fmt.Println(v8, v7)
    var v9 float32 = float32(0.0)
    fmt.Println(v8)
    switch "\n\t" {
    case "\n\t":
        v10 := []int{}
        _ = v10
    case "a":
        switch "" {
        case "":
            if false {
                switch "hello" {
                case "hello":
                    v8 -= v8
                default:
                    v11 := []int{}
                    _ = v11
                }
                v12 := []int{32767}
                var v13 int32 = (v7 - -1)
                _ = v12
                _ = v13
            }
        }
    case "hello":
        var v14 int = 128
        _ = v14
    }
    var v15 float32 = v9
    v9 = 3.14159
    v16 := GenS2{F0: "foo bar", F1: -862, F2: 0}
    v17 := []int{}
    v8 -= (v8 / ((1) | 1))
    fmt.Println(v9)
    var v18 uint32 = 251
    v19 := GenS2{F0: ("" + ("foo bar" + "")), F1: 539, F2: 42}
    v20 := GenS3{F0: 42, F1: 451, F2: 49}
    v7 -= -496
    switch 128 {
    case 960:
        if false {
            switch "\n\t" {
            case "":
                switch (("" + (("" + "hello") + "")) + ("\n\t" + ((("hello" + "foo bar") + "\n\t") + ""))) {
                case "":
                    switch "a" {
                    case "hello":
                        if false {
                            var v21 int8 = (v8 << ((401) & 63))
                            var v22 int64 = 42
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v15)
                        }
                    case "a":
                        if ((((v18 ^ 1) > 644) || (0 < -285)) || (-142 >= 4294967295)) {
                            v23 := GenS2{F0: (("\n\t" + "\n\t") + "a"), F1: -1, F2: 22}
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                v24 := GenS4{F0: (v7 >> ((2) & 63))}
                                _ = v24
                            case "hello":
                                if ((v7 & 611) <= 2147483647) {
                                    if false {
                                        var v25 uint32 = v18
                                        _ = v25
                                    }
                                } else {
                                    switch 32767 {
                                    case 32767:
                                        v8 = 127
                                    case -593:
                                        v7 = v7
                                    case 0:
                                        v9 = v9
                                    default:
                                        if true {
                                            fmt.Println(v7, v8)
                                            var v26 bool = true
                                            var v27 uint16 = 255
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                    v1[-9223372036854775808] = v18
                                    v17 = append(v17, 4294967295)
                                }
                            case "":
                                switch -509 {
                                case -675:
                                    var v28 string = "a"
                                    _ = v28
                                }
                            default:
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    v18 += 838
                                    _ = fi29
                                }
                            }
                            if true {
                                v30 := GenS5{F0: -784, F1: 545, F2: 255}
                                _ = v30
                            }
                            _ = v23
                        } else {
                            fmt.Println(v7, v18, v9)
                        }
                    default:
                        var v31 int8 = 100
                        _ = v31
                    }
                }
            case "foo bar":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case "a":
                v32 := GenS4{F0: v7}
                _ = v32
            }
        } else {
            var v33 uint16 = 42
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 uint16 = uint16(v18)
                _ = fi34
                _ = v35
            }
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}

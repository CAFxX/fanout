package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 string) uint8 {
    var v7 int = 42
    p0 = (p0 + "hello")
    var v8 uint64 = 0
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 1
    v3 := make(map[int]bool)
    var v4 int32 = -1
    var v5 uint64 = 4294967295
    v6 := make(map[string]bool)
    v9 := []int{628}
    var v10 uint = 18446744073709551615
    v10 += v10
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint32 = 2147483647
        _ = fi11
        _ = v12
    }
    v4 -= v4
    if true {
        switch "a" {
        case "foo bar":
            var v13 bool = false
            _ = v13
        }
    }
    var v14 int32 = v4
    if true {
        _ = len(v9)
    }
    var v15 float64 = 0.1
    v16 := GenS1{F0: -2, F1: (v10 >= (v10 & 390))}
    v10 = v10
    switch ("foo bar" + (("a" + ((("" + (("a" + ("\n\t" + "hello")) + "foo bar")) + "") + "foo bar")) + "foo bar")) {
    case "foo bar":
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint8 = (^(v2))
            if (!(("a" + "\n\t") <= "foo bar")) {
                switch ("\n\t" + ((("hello" + "") + "\n\t") + "foo bar")) {
                case "hello":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 uint = 500
                        switch (fi19 ^ 0) {
                        case 32767:
                            fmt.Println(v5, v10)
                        default:
                            v21 := GenS1{F0: -1, F1: (!false)}
                            _ = v21
                        }
                        var v22 float64 = 3.14159
                        _ = fi19
                        _ = v20
                        _ = v22
                    }
                }
                fmt.Println(v10)
            }
            var v23 uint64 = (^(v5))
            _ = fi17
            _ = v18
            _ = v23
        }
    case "":
        _ = len(v9)
    case "\n\t":
        switch (((("foo bar" + "\n\t") + "") + ("\n\t" + "foo bar")) + "") {
        case "a":
            v24 := GenS1{F0: 6, F1: false}
            _ = v24
        case "":
            switch 42 {
            case -577:
                v2 += 1
            case -1:
                _ = len(v9)
            case -123:
                if (-613 == 1) {
                    fmt.Println(v15, v4)
                    var v25 string = "\n\t"
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v25 = "\n\t"
                        v27 := GenS2{F0: v5}
                        _ = len(v9)
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    if (32768 >= -342) {
                        var v28 int8 = 1
                        _ = v28
                    } else {
                        var v29 int32 = 534
                        v30 := GenS2{F0: (^(v5))}
                        for fi31 := 0; fi31 < 14; fi31++ {
                            v5 += v5
                            switch -2147483648 {
                            case 1:
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    v4 = v14
                                    _ = fi32
                                }
                            case -711:
                                v29 -= v4
                            case 9223372036854775807:
                                for fi33 := 0; fi33 < 4; fi33++ {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = fi33
                                }
                            default:
                                v34 := GenS1{F0: -128, F1: false}
                                _ = v34
                            }
                            _ = fi31
                        }
                        _ = v29
                        _ = v30
                    }
                    for fi35 := 0; fi35 < 12; fi35++ {
                        if (genFn0("a") != v2) {
                            if true {
                                if true {
                                    var v36 int64 = 1
                                    _ = v36
                                } else {
                                    v15 = v15
                                    if false {
                                        fmt.Println(v14, v10, v14)
                                        var v37 int16 = 32767
                                        for fi38 := 0; fi38 < 15; fi38++ {
                                            v5 = (v5 / ((v5) | 1))
                                            _ = fi38
                                        }
                                        _ = v37
                                    } else {
                                        if ("hello" <= "") {
                                            var v39 int32 = 0
                                            _ = v39
                                        }
                                    }
                                }
                            } else {
                                v9 = append(v9, (fi35 + 9223372036854775807))
                            }
                            for fi40 := 0; fi40 < 9; fi40++ {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                v14 += (^(v14))
                                _ = fi40
                            }
                            v41 := GenS3{F0: (v14 & (v4 - 2147483647))}
                            _ = v41
                        }
                        _ = fi35
                    }
                    v14 = (v4 / ((2147483647) | 1))
                }
            }
        case "\n\t":
            v42 := GenS4{F0: 9223372036854775807, F1: v2}
            _ = v42
        default:
            if len(v9) > 0 {
                v9[0] = 42
            }
        }
    }
    fmt.Println(v15, v4)
    var v43 int16 = -887
    v2 = genFn0(("\n\t" + "hello"))
    v15 += ((float64((3.14159 - 3.14159)) + v15) - ((v15 * (1.5 + v15)) * 1.5))
    fmt.Println(v15, v5)
    var v44 uint = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}

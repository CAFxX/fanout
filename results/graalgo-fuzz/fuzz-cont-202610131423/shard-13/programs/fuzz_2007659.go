package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 string; F1 uint32; F2 bool }
type GenS5 struct { F0 string; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 int }
func genFn0(p0 uint64) bool {
    v7 := []int{931, -9223372036854775808}
    p0 -= (p0 / (((p0 + 315)) | 1))
    _ = p0
    _ = v7
    return false
}
func genFn1(p0 string) uint {
    fmt.Println(p0)
    if true {
        switch p0 {
        case "foo bar":
            v8 := GenS1{F0: 1}
            _ = v8
        default:
            v9 := GenS2{F0: 122}
            _ = v9
        }
    }
    _ = p0
    return 530
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (-2.25 + 3.14159)
    var v3 int = -58
    var v4 uint32 = 1
    v5 := make(map[int]int)
    v6 := []string{"hello", "hello", "", (((("hello" + "\n\t") + ("hello" + "hello")) + ("hello" + ("" + ""))) + ("a" + ""))}
    fmt.Println(v2, v3, v4)
    v10 := GenS3{F0: -1, F1: 255}
    var v11 bool = false
    switch "a" {
    case "foo bar":
        var v12 uint32 = 865
        _ = v12
    default:
        fmt.Println(v2, v11, v3)
    }
    switch "hello" {
    case "hello":
        delete(v5, (v3 & 32767))
    case "a":
        switch ("a" + "") {
        case "a":
            switch ("a" + (((("" + ("\n\t" + "")) + "") + "") + "foo bar")) {
            case "foo bar":
                fmt.Println(v4, v4)
            case "":
                v4 = v4
            case "a":
                if true {
                    if (!((607 >= 2147483648) && ("\n\t" < "a"))) {
                        var v13 bool = (((-1 >= -1) || v11) && v11)
                        switch "foo bar" {
                        case "\n\t":
                            var v14 int32 = 2147483647
                            _ = v14
                        case "hello":
                            v15 := GenS4{F0: "hello", F1: v4, F2: v11}
                            _ = v15
                        default:
                            v1[42] = 65535
                        }
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                        _ = v13
                    } else {
                        v11 = (!(true || (54 <= 100)))
                    }
                } else {
                    v16 := v5[v3]
                    for fi17 := 0; fi17 < 13; fi17++ {
                        fmt.Println(v11, v2, v2)
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v4 -= v4
                            fmt.Println(fi17)
                            _ = fi18
                        }
                        _ = fi17
                    }
                    switch (("\n\t" + "hello") + ("\n\t" + ((("a" + "\n\t") + "hello") + ""))) {
                    case "\n\t":
                        var v19 int8 = 1
                        _ = v19
                    case "a":
                        switch 1 {
                        case -703:
                            v3 = (^(v3))
                        default:
                            if true {
                                v20 := GenS5{F0: "hello", F1: 42, F2: (v4 * (v4 - (^(v4))))}
                                _ = v20
                            }
                        }
                    case "foo bar":
                        v21 := GenS3{F0: 100, F1: genFn1(("a" + "foo bar"))}
                        _ = v21
                    }
                    _ = v16
                }
            default:
                v22 := GenS2{F0: 80}
                _ = v22
            }
        }
    }
    fmt.Println(v4, v11, v2)
    var v23 float64 = float64(1e10)
    v24 := GenS5{F0: "hello", F1: 0, F2: 42}
    v25 := GenS4{F0: "", F1: (v4 / ((206) | 1)), F2: v11}
    switch "" {
    case "":
        v1[42] = v4
    case "a":
        v2 -= v23
    case "foo bar":
        if true {
            if v11 {
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 int = v3
                    var v28 bool = true
                    if true {
                        if (!(fi26 > 32767)) {
                            var v29 int8 = 42
                            _ = v29
                        }
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v30 := GenS6{F0: (v3 >> ((genFn1("a")) & 63)), F1: 4294967295, F2: (^(v4))}
                _ = len(v6)
                _ = v30
            } else {
                for fi31 := 0; fi31 < 6; fi31++ {
                    _ = len(v6)
                    switch "hello" {
                    case "a":
                        v32 := GenS7{F0: 709}
                        _ = v32
                    }
                    _ = fi31
                }
                switch 1 {
                case -933:
                    var v33 float64 = v23
                    _ = v33
                default:
                    if (!(0 >= -1)) {
                        var v34 int = v3
                        for fi35 := 0; fi35 < 16; fi35++ {
                            v6 = append(v6, ("a" + "a"))
                            v36 := GenS2{F0: 255}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(v4)
                        _ = v34
                    }
                }
            }
            var v37 uint = 65535
            for fi38 := 0; fi38 < 3; fi38++ {
                v39 := GenS8{F0: v4, F1: v3}
                if (289 == -9223372036854775808) {
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        var v40 float32 = 1e10
                        _ = v40
                    default:
                        v4 = (v4 & (v4 << ((genFn1("")) & 63)))
                    }
                    var v41 float32 = 1.5
                    v37 = (v37 >> ((v37) & 63))
                    _ = v41
                } else {
                    v37 = 1
                    var v42 int = 42
                    var v43 int32 = 843
                    _ = v42
                    _ = v43
                }
                _ = fi38
                _ = v39
            }
            _ = v37
        } else {
            v44 := v5[v3]
            _ = v44
        }
    }
    fmt.Println(v11, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

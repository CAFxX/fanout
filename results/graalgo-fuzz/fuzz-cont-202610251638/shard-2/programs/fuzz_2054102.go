package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 string) uint64 {
    v5 := []int{47, 4294967295, -2147483648}
    switch p0 {
    case "\n\t":
        var v6 uint32 = 2
        _ = v6
    }
    p0 = "foo bar"
    var v7 string = p0
    fmt.Println(p0, v7)
    _ = p0
    _ = v5
    _ = v7
    return 2
}

func main() {
    var v1 uint = 255
    v2 := []int{-9223372036854775808, 32768}
    var v3 float32 = 0.0
    v4 := []int{}
    fmt.Println(v3, v1)
    switch "" {
    case "a":
        v1 += v1
    }
    if ((127 > 1) && false) {
        v8 := GenS1{F0: 42}
        v1 += (v1 * v1)
        if false {
            var v9 float64 = 3.14159
            var v10 uint32 = 0
            _ = v10
            _ = v9
        } else {
            if len(v4) > 0 {
                v4[0] = 128
            }
            if (0 == 285) {
                var v11 uint32 = 0
                if (127 != 111) {
                    v11 = (v11 * v11)
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v3 -= 3.14159
                        _ = fi12
                    }
                } else {
                    _ = len(v2)
                }
                v11 -= uint32(v1)
                _ = v11
            }
            var v13 int64 = -701
            _ = v13
        }
        _ = v8
    } else {
        if true {
            var v14 uint32 = 42
            var v15 int64 = -1
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint8 = 2
                v18 := GenS1{F0: 2}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = len(v4)
            switch "\n\t" {
            case "a":
                var v19 string = "\n\t"
                _ = v19
            case "":
                if (!false) {
                    if false {
                        var v20 bool = false
                        v1 = (v1 * v1)
                        var v21 float32 = 3.14159
                        _ = v20
                        _ = v21
                    }
                    _ = len(v4)
                }
            case "foo bar":
                if true {
                    switch ("a" + "") {
                    case "":
                        v4 = append(v4, -879)
                    case "hello":
                        v1 = v1
                    }
                    if false {
                        v3 = 1.5
                        v22 := GenS2{F0: 9223372036854775807, F1: 4294967295}
                        var v23 int8 = 100
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint32 = 504
                        switch (("hello" + ("hello" + "foo bar")) + "foo bar") {
                        case "foo bar":
                            switch "" {
                            case "a":
                                v24 = 700
                            case "\n\t":
                                if false {
                                    v1 -= (v1 >> ((v1) & 63))
                                    v25 := GenS1{F0: 0}
                                    switch 517 {
                                    case 452:
                                        for fi26 := 0; fi26 < 2; fi26++ {
                                            v27 := GenS1{F0: 0}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v3, v3)
                                            _ = fi26
                                            _ = v27
                                        }
                                    case 2147483647:
                                        var v28 string = "foo bar"
                                        _ = v28
                                    default:
                                        v1 += (v1 >> ((v1) & 63))
                                    }
                                    _ = v25
                                }
                            }
                        case "hello":
                            for fi29 := 0; fi29 < 15; fi29++ {
                                var v30 string = "hello"
                                v3 -= v3
                                _ = fi29
                                _ = v30
                            }
                        default:
                            var v31 int = -628
                            _ = v31
                        }
                        switch "" {
                        case "a":
                            if true {
                                v2 = append(v2, 835)
                            } else {
                                v24 += (v24 >> ((v1) & 63))
                                v3 -= float32(0.0)
                            }
                        case "foo bar":
                            fmt.Println(v1)
                        }
                        _ = v24
                    }
                }
            default:
                if (!true) {
                    v32 := GenS2{F0: 0, F1: 1}
                    v3 += 1e10
                    _ = v32
                } else {
                    v33 := GenS1{F0: 54}
                    if (false || ((!(0 > 22)) || (1e10 < 0.1))) {
                        fmt.Println(v1, v1)
                        var v34 int32 = 958
                        _ = v34
                    } else {
                        v3 = (v3 / -2.25)
                        v35 := GenS2{F0: 1, F1: 2147483647}
                        _ = v35
                    }
                    v3 += v3
                    _ = v33
                }
            }
        }
        v3 = (1e10 + v3)
    }
    v36 := GenS1{F0: 2}
    var v37 int = 32767
    var v38 bool = false
    if (!(("a" + ("" + ("a" + "foo bar"))) == "")) {
        var v39 float64 = (1e10 / 1e10)
        fmt.Println(v39)
        _ = v39
    }
    if len(v2) > 0 {
        v2[0] = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}

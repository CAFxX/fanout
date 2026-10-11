package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 int8, p1 int64) uint64 {
    fmt.Println(p1)
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(p0, p1)
    case 42:
        var v4 uint = 18446744073709551615
        _ = v4
    case 32768:
        switch 0 {
        case 1:
            if true {
                switch -226 {
                case -417:
                    switch 19 {
                    case 4294967295:
                        p1 = (^(p1))
                    case 127:
                        v5 := []int{}
                        _ = v5
                    default:
                        p0 = 127
                    }
                default:
                    var v6 uint8 = 1
                    _ = v6
                }
                switch 0 {
                case -711:
                    for fi7 := 0; fi7 < 6; fi7++ {
                        fmt.Println(fi7, p0, fi7)
                        var v8 int8 = p0
                        _ = fi7
                        _ = v8
                    }
                }
            } else {
                switch ("a" + ("hello" + "hello")) {
                case "a":
                    var v9 float32 = 1e10
                    _ = v9
                case "foo bar":
                    v10 := []int{-9223372036854775808, 32768, int(p0)}
                    _ = v10
                default:
                    var v11 float64 = 0.0
                    _ = v11
                }
                v12 := []int{-9223372036854775808, -2147483648}
                var v13 uint8 = 249
                _ = v12
                _ = v13
            }
        case 945:
            var v14 uint8 = 0
            _ = v14
        }
    default:
        var v15 float32 = 1e10
        _ = v15
    }
    if (!(1 < 1000)) {
        p0 = (p0 & p0)
    } else {
        p0 = p0
    }
    var v16 int8 = p0
    _ = p0
    _ = p1
    _ = v16
    return 2
}

func main() {
    var v1 int = -904
    v2 := []uint32{}
    v3 := []string{"\n\t", "a", "\n\t", "a"}
    var v17 uint16 = 40
    v1 += (-(v1))
    var v18 uint8 = 104
    var v19 uint32 = 1
    if true {
        for fi20 := 0; fi20 < 16; fi20++ {
            switch "a" {
            case "hello":
                for fi21 := 0; fi21 < 18; fi21++ {
                    if ("foo bar" <= "") {
                        fmt.Println(v17)
                        var v22 uint32 = (v19 >> ((53) & 63))
                        _ = v22
                    } else {
                        var v23 uint64 = 2
                        if ("foo bar" == "foo bar") {
                            v24 := GenS1{F0: -292, F1: 4294967295, F2: (v23 | 0)}
                            _ = v24
                        } else {
                            if true {
                                v25 := GenS1{F0: 1, F1: (v19 * v19), F2: 42}
                                var v26 int8 = 42
                                switch v1 {
                                case -360:
                                    v17 = 65535
                                default:
                                    v1 = fi21
                                }
                                _ = len(v3)
                                v27 := GenS2{F0: (-(fi20))}
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                v1 = (fi21 - v1)
                                fmt.Println(fi20, fi21, v19)
                            }
                        }
                        _ = v23
                    }
                    v28 := GenS1{F0: 558, F1: v19, F2: 0}
                    _ = fi21
                    _ = v28
                }
            case "foo bar":
                _ = len(v3)
            }
            var v29 uint16 = 42
            _ = fi20
            _ = v29
        }
        v17 = v17
    }
    var v30 uint8 = v18
    _ = len(v3)
    fmt.Println(v17)
    v1 = 9223372036854775807
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            v1 = (^(fi31))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = ""
            }
            v17 = v17
            _ = fi32
        }
        _ = fi31
    }
    var v33 string = ""
    fmt.Println(v30, v33)
    fmt.Println(v1)
    v3 = append(v3, "a")
    var v34 int8 = 117
    fmt.Println(v17)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}

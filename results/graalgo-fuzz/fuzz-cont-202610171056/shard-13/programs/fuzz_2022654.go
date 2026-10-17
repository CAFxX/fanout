package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint, p1 bool, p2 int32) uint16 {
    if p1 {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 uint16 = 0
            _ = fi4
            _ = v5
        }
        if p1 {
            if p1 {
                v6 := GenS1{F0: -9223372036854775808, F1: 0, F2: 2147483647}
                v7 := GenS1{F0: -1, F1: p0, F2: 2147483647}
                switch -1 {
                case 32768:
                    for fi8 := 0; fi8 < 20; fi8++ {
                        v9 := GenS1{F0: 42, F1: 0, F2: 994}
                        _ = fi8
                        _ = v9
                    }
                case 0:
                    if (1 < 65535) {
                        for fi10 := 0; fi10 < 6; fi10++ {
                            if p1 {
                                switch "\n\t" {
                                case "foo bar":
                                    v11 := GenS1{F0: 9223372036854775807, F1: 1, F2: 1}
                                    _ = v11
                                case "a":
                                    for fi12 := 0; fi12 < 16; fi12++ {
                                        p2 += p2
                                        _ = fi12
                                    }
                                }
                                if (!p1) {
                                    v13 := []int{128, 633, 959}
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    fmt.Println(fi10, fi10, p1)
                                    _ = v13
                                }
                            }
                            fmt.Println(p1)
                            var v14 string = "a"
                            _ = fi10
                            _ = v14
                        }
                        var v15 string = ""
                        _ = v15
                    }
                default:
                    var v16 float64 = 0.1
                    _ = v16
                }
                _ = v6
                _ = v7
            }
            var v17 float64 = 3.14159
            _ = v17
        }
        v18 := []int{128, -725, -9223372036854775808, 745}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int64 = 0
            v21 := []int{(fi19 % (((fi19 << (((^(p0))) & 63))) | 1)), fi19}
            v20 += v20
            _ = fi19
            _ = v20
            _ = v21
        }
        fmt.Println(p2, p0, p1)
        switch 42 {
        case 1:
            v22 := []int{-9223372036854775808, 1, -646}
            _ = v22
        case 42:
            for fi23 := 0; fi23 < 16; fi23++ {
                p2 += 42
                _ = fi23
            }
        default:
            if (p0 >= (p0 % ((18446744073709551615) | 1))) {
                if p1 {
                    fmt.Println(p1, p0)
                    var v24 uint8 = 208
                    _ = v24
                } else {
                    switch 0 {
                    case -1:
                        p1 = ("foo bar" == ("foo bar" + "\n\t"))
                    case 9223372036854775807:
                        var v25 string = ("a" + "")
                        _ = v25
                    default:
                        p1 = p1
                    }
                    fmt.Println(p2, p1)
                }
            }
        }
    }
    var v26 int32 = (p2 ^ (^(p2)))
    switch ("\n\t" + "") {
    case "a":
        v26 += v26
    case "hello":
        for fi27 := 0; fi27 < 6; fi27++ {
            p2 = p2
            _ = fi27
        }
    }
    v26 = 42
    v28 := GenS2{F0: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v28
    return 2
}
func genFn1() int64 {
    var v29 int = 337
    v30 := GenS3{F0: (v29 | (v29 >> ((0) & 63))), F1: 406}
    v29 = (v29 - (v29 >> ((42) & 63)))
    v29 += v29
    _ = v29
    _ = v30
    return 380
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = (("foo bar" + "\n\t") + "a")
    v3 := make(map[string]bool)
    v3[v2] = false
    for fi31 := 0; fi31 < 7; fi31++ {
        v2 = v2
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 int16 = -281
            v33 = v33
            _ = fi32
            _ = v33
        }
        var v34 int64 = 1
        _ = fi31
        _ = v34
    }
    fmt.Println(v2)
    _ = len(v3)
    switch v2 {
    case "":
        if true {
            v2 = "hello"
            var v35 int16 = -32768
            _ = v35
        }
    case "foo bar":
        v2 = v2
    }
    var v36 float32 = 1.5
    fmt.Println(v2)
    v37 := []int{2147483647, -374}
    switch 32767 {
    case 1:
        _ = len(v37)
    case 42:
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
    }
    var v38 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}

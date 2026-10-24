package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint64) int {
    p0 = (p0 * 18446744073709551615)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int8 = 0
        _ = fi5
        _ = v6
    }
    _ = p0
    return 388
}

func main() {
    var v1 string = "a"
    var v2 float64 = 3.14159
    var v3 string = ("\n\t" + v1)
    v4 := make(map[int]uint32)
    v7 := GenS1{F0: 32767, F1: -32768, F2: 2147483648}
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch v3 {
            case "foo bar":
                fmt.Println(v2, v1)
            }
            _ = fi8
        }
        switch v3 {
        case "":
            switch v1 {
            case "\n\t":
                var v9 int = 32768
                _ = v9
            case "hello":
                v10 := GenS2{F0: "a", F1: 1}
                _ = v10
            }
        case "a":
            fmt.Println(v1, v1)
        }
    } else {
        var v11 uint32 = 42
        switch 32768 {
        case -1:
            v12 := GenS2{F0: v1, F1: -1}
            _ = v12
        case 128:
            v13 := v4[364]
            _ = v13
        }
        v3 = "foo bar"
        _ = v11
    }
    v4[-112] = 0
    for fi14 := 0; fi14 < 14; fi14++ {
        v2 = 1.5
        v15 := []int{fi14, genFn0(18446744073709551615), 127, (fi14 << ((858) & 63))}
        _ = fi14
        _ = v15
    }
    v16 := []int{2147483647}
    if false {
        for fi17 := 0; fi17 < 16; fi17++ {
            v1 = "foo bar"
            var v18 uint = 4294967295
            if (true && ((1.5 * 3.14159) == 0.1)) {
                var v19 float64 = v2
                v1 = v1
                fmt.Println(v2, v19)
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 bool = true
                v1 = ((v3 + v3) + v1)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    } else {
        fmt.Println(v3, v1)
    }
    v1 = "foo bar"
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: v1, F1: -1}
        if true {
            var v25 string = (v1 + ("a" + ""))
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 uint = 1
                _ = fi26
                _ = v27
            }
            v16 = append(v16, fi23)
            _ = v25
        } else {
            var v28 int64 = 9223372036854775807
            _ = v28
        }
        _ = fi23
        _ = v24
    }
    switch v3 {
    case "foo bar":
        fmt.Println(v2, v1, v3)
    case "":
        if false {
            v29 := v4[32767]
            if (!true) {
                if true {
                    fmt.Println(v2, v3)
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 int = 128
                        if false {
                            for fi32 := 0; fi32 < 3; fi32++ {
                                v33 := GenS2{F0: "foo bar", F1: -128}
                                v29 = (v29 - 4294967295)
                                switch 32767 {
                                case 107:
                                    var v34 int = genFn0(0)
                                    _ = v34
                                case -9223372036854775808:
                                    v29 = (^(v29))
                                case 1:
                                    if false {
                                        _ = len(v16)
                                        v29 = v29
                                        var v35 bool = true
                                        _ = v35
                                    }
                                }
                                _ = fi32
                                _ = v33
                            }
                            var v36 int8 = 42
                            _ = v36
                        }
                        _ = fi30
                        _ = v31
                    }
                    v29 = (v29 & 2147483647)
                }
                for fi37 := 0; fi37 < 18; fi37++ {
                    if (false || false) {
                        v1 = (v3 + v3)
                    }
                    var v38 uint64 = 709
                    v29 -= 2147483647
                    _ = fi37
                    _ = v38
                }
                v39 := GenS2{F0: v3, F1: -1}
                _ = v39
            }
            fmt.Println(v29)
            _ = v29
        } else {
            if (100 == -20) {
                var v40 float64 = ((1.5 / float64((0.1 / 1.0))) * 0.1)
                v41 := GenS3{F0: 1, F1: 1}
                var v42 bool = (false || false)
                _ = v40
                _ = v41
                _ = v42
            } else {
                v43 := v4[-923]
                v1 = v3
                v44 := GenS4{F0: 42, F1: 305}
                _ = v43
                _ = v44
            }
        }
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

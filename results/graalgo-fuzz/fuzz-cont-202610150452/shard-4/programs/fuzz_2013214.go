package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32) uint64 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += p0
        var v8 uint32 = 2147483648
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 251
    _ = p0
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 uint32, p2 uint32) int8 {
    var v10 int64 = 352
    v11 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -128
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -1
    v3 := []string{"foo bar"}
    v4 := make(map[string]int)
    var v5 int8 = 0
    var v6 string = "hello"
    switch -367 {
    case -9223372036854775808:
        var v12 bool = true
        _ = v12
    case 1:
        var v13 int = 127
        _ = v13
    default:
        _ = len(v4)
    }
    fmt.Println(v2, v5)
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS1{F0: -950, F1: -425, F2: 1}
            var v17 float32 = -2.25
            switch -275 {
            case 42:
                fmt.Println(v5, v17)
            case 9223372036854775807:
                _ = len(v3)
            case -9223372036854775808:
                if len(v3) > 0 {
                    v3[0] = v6
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 1000, F1: 42, F2: (v5 * (v5 | -128))}
        _ = fi14
        _ = v18
    }
    if (true || (65535 >= 1)) {
        v6 = v6
        if (-744 > -798) {
            fmt.Println(v5, v2)
        }
        if true {
            v19 := GenS2{F0: 235}
            v6 = "foo bar"
            var v20 int64 = 1
            _ = v19
            _ = v20
        } else {
            var v21 int32 = 2147483647
            v4[v6] = 9223372036854775807
            _ = v21
        }
    }
    _ = len(v4)
    var v22 int32 = 2147483647
    if true {
        if true {
            _ = len(v1)
            v22 = v22
            if true {
                for fi23 := 0; fi23 < 8; fi23++ {
                    switch v6 {
                    case "foo bar":
                        v24 := GenS1{F0: 1000, F1: -1, F2: 0}
                        _ = v24
                    case "a":
                        v5 = (v5 & v5)
                    case "":
                        if false {
                            v4["foo bar"] = fi23
                        } else {
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v22 = (v22 & v22)
                                delete(v4, ("" + ("foo bar" + v6)))
                                v6 = v6
                                _ = fi25
                            }
                        }
                    }
                    var v26 int = fi23
                    _ = fi23
                    _ = v26
                }
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v27 int = 9223372036854775807
                _ = v27
            }
        }
    } else {
        v5 = (v5 * v5)
        fmt.Println(v2)
        v6 = "foo bar"
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS1{F0: 0, F1: (^(v22)), F2: genFn1(42, 521, 2147483647)}
        _ = fi28
        _ = v29
    }
    var v30 float64 = 3.14159
    var v31 int64 = (v2 ^ (v2 % ((v2) | 1)))
    var v32 bool = false
    switch -787 {
    case -2147483648:
        var v33 int = 127
        _ = v33
    case -259:
        var v34 uint64 = 470
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v22 += v22
        v32 = true
        delete(v1, v6)
        _ = fi35
    }
    if (v32 || (-1 > v5)) {
        v36 := GenS1{F0: 674, F1: v22, F2: v5}
        var v37 uint = 812
        for fi38 := 0; fi38 < 5; fi38++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v6 = ""
            _ = fi38
        }
        _ = v36
        _ = v37
    } else {
        _ = len(v1)
        switch 4294967295 {
        case 32767:
            switch -2147483648 {
            case 42:
                var v39 int16 = -32768
                _ = v39
            default:
                if v32 {
                    v6 = (v6 + "foo bar")
                } else {
                    _ = len(v3)
                }
            }
        }
        fmt.Println(v30, v2, v6)
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        switch fi40 {
        case 42:
            v6 = "\n\t"
        case 127:
            v30 = v30
        }
        _ = fi40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

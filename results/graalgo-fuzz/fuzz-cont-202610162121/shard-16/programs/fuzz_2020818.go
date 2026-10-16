package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int8 = 1
    var v2 int8 = v1
    var v3 uint32 = 4294967295
    v4 := []string{}
    v5 := make(map[string]string)
    _ = len(v5)
    fmt.Println(v2)
    v1 = (-(v2))
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch ("foo bar" + "") {
            case "\n\t":
                v1 += 1
            case "foo bar":
                fmt.Println(v3, v2, v1)
            }
            v8 := GenS1{F0: 2}
            v2 -= (v1 ^ (^(v1)))
            v2 = 100
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v1 = 1
    for fi9 := 0; fi9 < 8; fi9++ {
        v2 -= v1
        v10 := GenS1{F0: 255}
        var v11 uint = 255
        _ = fi9
        _ = v10
        _ = v11
    }
    v4 = append(v4, "")
    v12 := GenS2{F0: 701, F1: 961, F2: 42}
    var v13 int16 = 0
    switch "\n\t" {
    case "a":
        v14 := v5["a"]
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 9; fi15++ {
            if false {
                _ = len(v4)
                if len(v4) > 0 {
                    v4[0] = (("hello" + (("a" + "foo bar") + "a")) + "a")
                }
            }
            _ = fi15
        }
    case "foo bar":
        switch "foo bar" {
        case "a":
            v13 += v13
        default:
            if len(v4) > 0 {
                v4[0] = ("hello" + ("\n\t" + "\n\t"))
            }
        }
    default:
        v1 = -128
    }
    switch ("" + ("a" + "foo bar")) {
    case "":
        switch (("\n\t" + "hello") + "") {
        case "a":
            var v16 uint8 = 0
            _ = v16
        case "":
            if ((v1 * v1) >= -1) {
                v17 := GenS2{F0: 4294967295, F1: 32767, F2: uint64(v3)}
                _ = v17
            } else {
                switch 710 {
                case 967:
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        delete(v5, "foo bar")
                        var v18 uint = 42
                        fmt.Println(v13)
                        _ = v18
                    }
                default:
                    v2 -= v1
                }
                switch 4294967295 {
                case 127:
                    if ((-2.25 != -2.25) && (-1 == v2)) {
                        v1 -= v1
                        fmt.Println(v3, v3)
                        var v19 uint = 1
                        _ = v19
                    } else {
                        v1 = (v1 + v2)
                        v3 = 2147483648
                        fmt.Println(v1)
                    }
                default:
                    if true {
                        fmt.Println(v13)
                        v13 += v13
                        v1 += (v1 | v1)
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ("hello" + ("hello" + "hello"))
                        }
                        v20 := GenS2{F0: -747, F1: 465, F2: 2}
                        _ = v20
                    }
                }
            }
        default:
            _ = len(v4)
        }
    case "\n\t":
        v2 = v1
    case "foo bar":
        var v21 string = "\n\t"
        _ = v21
    default:
        var v22 int16 = 32767
        _ = v22
    }
    switch -18 {
    case -9223372036854775808:
        var v23 int64 = 471
        _ = v23
    }
    v24 := GenS3{F0: v13, F1: 127}
    switch 333 {
    case 0:
        if true {
            var v25 int32 = 68
            _ = v25
        }
    default:
        v26 := GenS4{F0: -581}
        _ = v26
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        fmt.Println(v13, fi27)
        fmt.Println(v2, fi27)
        var v28 float64 = 3.14159
        _ = fi27
        _ = v28
    }
    var v29 uint = 0
    var v30 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 string = (((("foo bar" + ("" + ("foo bar" + (("" + "hello") + "hello")))) + ("foo bar" + (("foo bar" + "") + ""))) + ("foo bar" + "hello")) + "foo bar")
    v2 := make(map[string]string)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]uint32)
    var v5 bool = true
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            var v8 uint16 = 42
            v9 := GenS1{F0: 128, F1: 65535}
            _ = v8
            _ = v9
        }
        if ((v3 ^ v3) != v3) {
            if (!v5) {
                var v10 int64 = -9223372036854775808
                _ = v10
            }
            for fi11 := 0; fi11 < 13; fi11++ {
                switch (fi7 / (((fi7 / ((fi7) | 1))) | 1)) {
                case -576:
                    v12 := []int{}
                    _ = v12
                default:
                    var v13 int16 = 1
                    _ = v13
                }
                v5 = v5
                _ = fi11
            }
            var v14 int8 = 100
            _ = v14
        } else {
            fmt.Println(v1, v5, v5)
            switch v1 {
            case "":
                if v5 {
                    var v15 uint8 = 42
                    v16 := []int{(fi7 + fi7), (fi7 >> ((18446744073709551615) & 63))}
                    _ = v15
                    _ = v16
                } else {
                    v5 = v5
                }
            case "a":
                var v17 int32 = 0
                _ = v17
            }
        }
        _ = fi7
    }
    v5 = (true && false)
    var v18 bool = v5
    fmt.Println(v5, v1)
    v19 := []int{42}
    v6[1] = 0
    v19 = append(v19, 4294967295)
    v1 = (v1 + v1)
    switch 925 {
    case 310:
        if v18 {
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 15; fi21++ {
                    switch fi21 {
                    case 32768:
                        var v22 uint8 = 1
                        _ = v22
                    case -9223372036854775808:
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    case -2147483648:
                        v6[0] = 2147483647
                    }
                    if v18 {
                        var v24 int = fi21
                        v5 = (0 > 0)
                        _ = v24
                    } else {
                        delete(v4, 128)
                        var v25 uint8 = 174
                        delete(v2, (("\n\t" + "foo bar") + ("a" + "foo bar")))
                        var v26 int = (fi21 + (fi20 >> ((80) & 63)))
                        v27 := GenS1{F0: (^(v25)), F1: 255}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = len(v2)
                    if (-1 != 42) {
                        v1 = ("" + "hello")
                        var v28 float64 = 3.14159
                        var v29 float32 = (0.1 - 0.0)
                        _ = v28
                        _ = v29
                    }
                    _ = fi21
                }
                _ = fi20
            }
            v5 = true
        } else {
            v1 = v1
        }
    case 0:
        var v30 bool = v5
        _ = v30
    }
    var v31 uint = 593
    if false {
        switch v1 {
        case "\n\t":
            var v32 int8 = 10
            _ = v32
        case "hello":
            for fi33 := 0; fi33 < 15; fi33++ {
                for fi34 := 0; fi34 < 10; fi34++ {
                    if len(v19) > 0 {
                        v19[0] = (-(fi34))
                    }
                    switch v1 {
                    case "":
                        delete(v4, 32768)
                    default:
                        var v35 int32 = -1
                        _ = v35
                    }
                    _ = fi34
                }
                fmt.Println(v31, v3)
                _ = fi33
            }
        default:
            var v36 uint = (v31 - v31)
            _ = v36
        }
        switch "a" {
        case "a":
            fmt.Println(v18, v3, v5)
        case "\n\t":
            delete(v4, -1)
        case "":
            var v37 int32 = -819
            _ = v37
        default:
            v31 -= v31
        }
        if len(v19) > 0 {
            v19[0] = 0
        }
    }
    if len(v19) > 0 {
        v19[0] = -1
    }
    v1 = "foo bar"
    switch -90 {
    case 511:
        for fi38 := 0; fi38 < 3; fi38++ {
            for fi39 := 0; fi39 < 1; fi39++ {
                v5 = ((((float64((float32(1e10) * 1e10)) + 0.0) - (1e10 * -2.25)) != 3.14159) && true)
                if len(v19) > 0 {
                    v19[0] = fi39
                }
                v3 = 177
                _ = fi39
            }
            _ = fi38
        }
    case 908:
        fmt.Println(v31, v1, v5)
    }
    var v40 uint16 = uint16(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

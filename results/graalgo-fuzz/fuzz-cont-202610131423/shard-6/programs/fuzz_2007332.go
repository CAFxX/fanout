package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint16) int32 {
    if false {
        p0 -= 32767
        switch -1 {
        case 636:
            v4 := []int{-1, 127}
            _ = v4
        case 4294967295:
            switch ((("" + "foo bar") + ((((("a" + "hello") + ("a" + "hello")) + "foo bar") + "a") + ("\n\t" + "a"))) + ("foo bar" + "a")) {
            case "hello":
                v5 := []int{9223372036854775807, 32768, 128}
                _ = v5
            default:
                var v6 uint8 = 88
                _ = v6
            }
        default:
            if (!(866 >= 2147483647)) {
                var v7 int32 = -778
                _ = v7
            } else {
                if (((!false) || false) || false) {
                    v8 := []int{}
                    _ = v8
                } else {
                    var v9 float32 = (-2.25 * 0.1)
                    var v10 uint8 = 37
                    _ = v10
                    _ = v9
                }
                if (!false) {
                    var v11 uint8 = 174
                    _ = v11
                }
                switch "a" {
                case "a":
                    switch 32767 {
                    case 0:
                        var v12 uint8 = 255
                        _ = v12
                    case 2147483647:
                        var v13 int16 = 4
                        _ = v13
                    default:
                        var v14 int16 = -32768
                        _ = v14
                    }
                case "":
                    var v15 int16 = 293
                    _ = v15
                }
            }
        }
        fmt.Println(p0)
    } else {
        p0 = (p0 / ((p0) | 1))
        p0 = (p0 & (^(p0)))
    }
    fmt.Println(p0)
    v16 := GenS1{F0: 255}
    for fi17 := 0; fi17 < 6; fi17++ {
        switch "foo bar" {
        case "":
            v18 := []int{(fi17 | fi17), 191, (fi17 & fi17)}
            _ = v18
        case "\n\t":
            var v19 bool = (!(!false))
            _ = v19
        default:
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int = fi17
                _ = fi20
                _ = v21
            }
        }
        var v22 float64 = 3.14159
        _ = fi17
        _ = v22
    }
    _ = p0
    _ = v16
    return -2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    switch ("a" + ((("" + ((("hello" + "") + "foo bar") + "hello")) + "\n\t") + ("\n\t" + "hello"))) {
    case "foo bar":
        switch ("" + "a") {
        case "a":
            fmt.Println(v3)
        case "foo bar":
            _ = len(v1)
        case "":
            fmt.Println(v3)
        }
    case "\n\t":
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 string = "a"
            if false {
                var v25 int32 = 304
                _ = v25
            } else {
                switch "foo bar" {
                case "hello":
                    v24 = v24
                }
                var v26 uint32 = 4294967295
                v27 := []int{fi23}
                _ = v26
                _ = v27
            }
            v3 = (v3 | 2)
            _ = fi23
            _ = v24
        }
    default:
        var v28 uint16 = 32767
        _ = v28
    }
    if ((0.1 == 0.1) && (0.0 == (-2.25 / -2.25))) {
        for fi29 := 0; fi29 < 10; fi29++ {
            v3 += 18446744073709551615
            _ = fi29
        }
    }
    _ = len(v1)
    fmt.Println(v3)
    v30 := GenS1{F0: 2}
    v31 := GenS1{F0: 18446744073709551615}
    v3 = 425
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            v1[407] = 65535
            var v34 uint64 = 2
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    v35 := GenS2{F0: 617, F1: 370, F2: 42}
    if false {
        for fi36 := 0; fi36 < 7; fi36++ {
            v3 = v3
            v37 := []int{fi36}
            _ = fi36
            _ = v37
        }
        v3 = 42
        for fi38 := 0; fi38 < 14; fi38++ {
            v3 += (v3 / (((^(v3))) | 1))
            _ = fi38
        }
    }
    v3 = v3
    v3 += v3
    if true {
        v3 += (^(v3))
        for fi39 := 0; fi39 < 11; fi39++ {
            if false {
                v3 = v3
            }
            _ = fi39
        }
    } else {
        v40 := GenS3{F0: genFn0(65535), F1: 2147483647}
        var v41 uint = 2
        fmt.Println(v3, v3)
        _ = v40
        _ = v41
    }
    var v42 int32 = 0
    v43 := GenS3{F0: 692, F1: 629}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v42
    _ = v43
}

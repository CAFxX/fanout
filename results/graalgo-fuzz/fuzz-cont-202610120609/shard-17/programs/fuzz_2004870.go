package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 int64 }

func main() {
    v1 := []uint32{2147483648, 265}
    v2 := make(map[int]string)
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 255
    v5 := []uint32{42, 1, 42}
    v3 = (v3 & v3)
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(fi6, v3, v3)
        _ = fi6
    }
    if ("\n\t" <= ((("\n\t" + "foo bar") + (((((("foo bar" + "hello") + "foo bar") + (((("a" + "") + ("foo bar" + "\n\t")) + "hello") + "a")) + "\n\t") + "") + ("hello" + "foo bar"))) + "")) {
        _ = len(v5)
    }
    var v7 int16 = int16(v3)
    var v8 int8 = 43
    var v9 int64 = -471
    switch 9223372036854775807 {
    case 128:
        for fi10 := 0; fi10 < 17; fi10++ {
            v3 = v3
            v11 := GenS1{F0: -65}
            _ = fi10
            _ = v11
        }
    case 42:
        var v12 int = 32767
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 2; fi13++ {
            v7 = -32768
            var v14 int64 = (-(v9))
            var v15 float32 = float32(3.14159)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := GenS2{F0: 336, F1: (v4 + v4), F2: true}
    switch ("a" + (("hello" + "a") + (("\n\t" + "") + ("" + ("\n\t" + ("a" + ("hello" + ""))))))) {
    case "a":
        for fi17 := 0; fi17 < 7; fi17++ {
            if true {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS3{F0: 0}
                    switch fi18 {
                    case 32767:
                        if false {
                            var v20 float32 = 3.14159
                            var v21 int64 = 9223372036854775807
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v7, v3)
                            v8 = v8
                            v8 += v8
                        }
                    case 2147483647:
                        fmt.Println(v7, fi18)
                    default:
                        var v22 uint64 = v3
                        _ = v22
                    }
                    v23 := GenS4{F0: 443, F1: 1, F2: int64(v3)}
                    _ = fi18
                    _ = v19
                    _ = v23
                }
                v24 := GenS3{F0: 993}
                _ = v24
            }
            if true {
                var v25 float64 = (3.14159 - -2.25)
                _ = v25
            }
            _ = fi17
        }
    case "foo bar":
        switch "hello" {
        case "foo bar":
            if (!(!true)) {
                v2[434] = "hello"
                fmt.Println(v7, v8, v4)
            }
        case "\n\t":
            if (!true) {
                for fi26 := 0; fi26 < 10; fi26++ {
                    switch "foo bar" {
                    case "a":
                        v2[(fi26 + 660)] = "hello"
                    default:
                        v27 := GenS3{F0: 255}
                        _ = v27
                    }
                    v28 := v2[(fi26 | fi26)]
                    var v29 string = v28
                    _ = fi26
                    _ = v28
                    _ = v29
                }
                if (((v8 == v8) || true) || ((v3 == v3) || true)) {
                    var v30 string = "\n\t"
                    _ = len(v1)
                    _ = v30
                }
                switch 923 {
                case 9223372036854775807:
                    switch "foo bar" {
                    case "\n\t":
                        v31 := GenS2{F0: 32767, F1: v4, F2: (!false)}
                        _ = v31
                    }
                case 2147483647:
                    var v32 int = 42
                    _ = v32
                case -926:
                    var v33 string = "\n\t"
                    _ = v33
                default:
                    v34 := GenS3{F0: uint(v4)}
                    _ = v34
                }
            } else {
                fmt.Println(v7, v9, v9)
                v8 -= -57
            }
        case "":
            v8 += v8
        }
    case "\n\t":
        var v35 uint8 = 2
        _ = v35
    default:
        for fi36 := 0; fi36 < 17; fi36++ {
            v9 = (v9 | 445)
            var v37 int8 = v8
            v9 -= v9
            _ = fi36
            _ = v37
        }
    }
    if false {
        var v38 float64 = float64(0.0)
        switch "hello" {
        case "\n\t":
            for fi39 := 0; fi39 < 5; fi39++ {
                v7 = v7
                _ = fi39
            }
        default:
            var v40 bool = false
            _ = v40
        }
        v41 := GenS4{F0: 0, F1: 1, F2: (v9 + v9)}
        _ = v38
        _ = v41
    }
    v7 -= v7
    var v42 uint16 = v4
    if (true || (234 != -680)) {
        v43 := GenS3{F0: 592}
        var v44 int8 = v8
        v45 := GenS3{F0: 65535}
        _ = v43
        _ = v44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}

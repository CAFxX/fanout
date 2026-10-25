package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    v1 := []uint32{42}
    v2 := make(map[int]string)
    var v3 string = "foo bar"
    if (true || (!true)) {
        for fi4 := 0; fi4 < 7; fi4++ {
            v3 = ("" + v3)
            _ = fi4
        }
        if (1 >= 100) {
            var v5 uint = 1
            var v6 uint16 = 598
            _ = v5
            _ = v6
        }
        fmt.Println(v3)
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            var v8 uint8 = 200
            v9 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v8
            _ = v9
        } else {
            var v10 uint16 = 2
            delete(v2, (fi7 >> ((42) & 63)))
            _ = v10
        }
        if (false || ((v3 + (v3 + v3)) != (v3 + "foo bar"))) {
            var v11 int64 = 0
            var v12 float32 = 3.14159
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(fi13, v3)
        _ = fi13
    }
    if (660 >= 2) {
        v3 = v3
        v2[127] = (v3 + "a")
    }
    v14 := GenS1{F0: -1, F1: 255}
    if true {
        v3 = v3
        fmt.Println(v3)
        var v15 bool = false
        _ = v15
    }
    switch "hello" {
    case "hello":
        v16 := GenS1{F0: -1, F1: 42}
        _ = v16
    case "foo bar":
        var v17 int = 423
        _ = v17
    default:
        _ = len(v1)
    }
    if false {
        v2[9223372036854775807] = v3
        var v18 uint8 = 0
        if (1 <= 548) {
            var v19 uint64 = 0
            var v20 float64 = 1.5
            _ = v19
            _ = v20
        }
        _ = v18
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v3 = ("" + "hello")
        var v22 int = fi21
        _ = fi21
        _ = v22
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = v3
    if (42 != 18446744073709551615) {
        var v23 uint32 = 1
        var v24 string = v3
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int8 = -1
            _ = fi25
            _ = v26
        }
        _ = v23
        _ = v24
    }
    switch 42 {
    case 128:
        if (true || (false || false)) {
            if ("" != v3) {
                v3 = ("a" + v3)
                if true {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v28 := GenS1{F0: 680, F1: 939}
                        for fi29 := 0; fi29 < 18; fi29++ {
                            var v30 int8 = -1
                            v31 := GenS1{F0: 9223372036854775807, F1: 42}
                            if false {
                                if (false && (!((-631 >= 1000) && ((true && true) || (2 > 65535))))) {
                                    v30 = v30
                                    v30 += (v30 % (((v30 | (v30 << ((409) & 63)))) | 1))
                                }
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        var v32 int16 = 1
                        _ = fi27
                        _ = v28
                        _ = v32
                    }
                    v3 = v3
                    var v33 float64 = 1.5
                    _ = v33
                }
            }
        } else {
            if false {
                switch ((("" + "") + "foo bar") + (("a" + v3) + ("" + ("hello" + "")))) {
                case "\n\t":
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v3 = (v3 + v3)
                        _ = fi34
                    }
                case "hello":
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v36 := GenS2{F0: 42}
                        _ = fi35
                        _ = v36
                    }
                }
                var v37 uint32 = 65535
                v38 := GenS2{F0: 0}
                _ = v37
                _ = v38
            } else {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v3)
                case "hello":
                    var v39 uint32 = 0
                    _ = v39
                default:
                    fmt.Println(v3)
                }
            }
            v3 = v3
            var v40 int8 = 0
            _ = v40
        }
    }
    v2[1] = ((v3 + v3) + "")
    switch ("" + (("\n\t" + ("\n\t" + (((((v3 + v3) + (("foo bar" + (v3 + "foo bar")) + "")) + "") + v3) + (v3 + "foo bar")))) + ("hello" + (v3 + ("foo bar" + ""))))) {
    case "a":
        for fi41 := 0; fi41 < 1; fi41++ {
            v3 = v3
            _ = fi41
        }
    }
    switch 127 {
    case -406:
        v42 := GenS3{F0: 0, F1: 2147483648}
        _ = v42
    case -2147483648:
        if false {
            if (18446744073709551615 < 255) {
                if false {
                    v3 = ((v3 + v3) + (v3 + ("\n\t" + "")))
                }
                if false {
                    switch v3 {
                    case "\n\t":
                        fmt.Println(v3)
                    case "a":
                        var v43 int32 = 42
                        _ = v43
                    case "hello":
                        if true {
                            v3 = v3
                        }
                    default:
                        var v44 string = v3
                        _ = v44
                    }
                    switch 506 {
                    case 4294967295:
                        var v45 int16 = -32768
                        _ = v45
                    }
                }
            } else {
                v3 = (v3 + "foo bar")
                v3 = "a"
            }
            if (-1 <= -2147483648) {
                v3 = (((v3 + v3) + v3) + v3)
                fmt.Println(v3)
            } else {
                if len(v1) > 0 {
                    v1[0] = 166
                }
                switch -1 {
                case 4294967295:
                    var v46 float64 = float64(0.1)
                    _ = v46
                case -290:
                    fmt.Println(v3)
                }
            }
        }
    default:
        v2[-886] = (v3 + ((("" + ("\n\t" + ((v3 + "\n\t") + v3))) + v3) + ("foo bar" + (((v3 + v3) + "") + v3))))
    }
    var v47 int64 = 1
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v47
}

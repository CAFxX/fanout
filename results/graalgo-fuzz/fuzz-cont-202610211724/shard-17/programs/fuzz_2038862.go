package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int8 }
func genFn0(p0 uint, p1 int8, p2 int64) uint64 {
    p1 = -1
    if false {
        if false {
            p2 = p2
            var v6 uint64 = 1
            var v7 float64 = (0.1 - (float64(((1.5 / 1.0) * -2.25)) * 1e10))
            _ = v6
            _ = v7
        }
    } else {
        if (false && false) {
            if ((!(!(((!true) || true) || true))) || false) {
                p2 += 30
                if (127 >= 646) {
                    var v8 int64 = (p2 + p2)
                    fmt.Println(v8, p1, p0)
                    _ = v8
                } else {
                    fmt.Println(p2, p0, p0)
                    p1 = p1
                }
            }
        }
        fmt.Println(p0)
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{(fi9 * (fi9 ^ (fi9 + 128))), fi9, 128}
        _ = fi9
        _ = v10
    }
    var v11 bool = (255 >= 127)
    var v12 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 826
}
func genFn1(p0 uint16, p1 bool) uint8 {
    var v13 int32 = 42
    v14 := []int{375}
    var v15 float32 = 3.14159
    for fi16 := 0; fi16 < 19; fi16++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            switch "hello" {
            case "hello":
                v18 := GenS1{F0: ""}
                _ = v18
            }
            var v19 uint = 390
            _ = fi17
            _ = v19
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 208
}

func main() {
    var v1 bool = (2 > 2)
    var v2 bool = (100 < 1)
    v3 := []string{(("" + "a") + "foo bar")}
    var v4 uint = 0
    var v5 float32 = 0.1
    var v20 uint8 = 2
    switch ("a" + ("a" + "foo bar")) {
    case "foo bar":
        fmt.Println(v4, v5)
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            v5 += (3.14159 * (v5 * v5))
        case "":
            if (v4 == (^(v4))) {
                v2 = false
            }
        case "a":
            v3 = append(v3, ("" + ""))
        default:
            switch "foo bar" {
            case "":
                for fi21 := 0; fi21 < 10; fi21++ {
                    v22 := GenS2{F0: -1, F1: 1}
                    fmt.Println(fi21, v2, fi21)
                    var v23 float32 = (float32(-2.25) + float32(0.1))
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case "a":
                v24 := GenS2{F0: 1, F1: 59}
                _ = v24
            default:
                for fi25 := 0; fi25 < 13; fi25++ {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    _ = fi25
                }
            }
        }
    case "hello":
        fmt.Println(v2, v4)
    default:
        if v2 {
            if ("" == ("" + ("foo bar" + ("a" + (("" + "foo bar") + (("hello" + "") + "a")))))) {
                if false {
                    if len(v3) > 0 {
                        v3[0] = ("a" + "")
                    }
                }
                var v26 uint32 = 0
                var v27 uint8 = 0
                _ = v26
                _ = v27
            } else {
                v1 = ((826 == (v4 / (((v4 & 817)) | 1))) && v1)
            }
            switch "\n\t" {
            case "\n\t":
                var v28 int16 = -1
                _ = v28
            default:
                fmt.Println(v5, v4, v4)
            }
            fmt.Println(v20)
        }
    }
    var v29 uint16 = 0
    if false {
        var v30 uint8 = (v20 | v20)
        v31 := GenS3{F0: (^(v29)), F1: uint16(v4)}
        if len(v3) > 0 {
            v3[0] = ("a" + "")
        }
        _ = v30
        _ = v31
    }
    if (100 <= 1) {
        switch ("\n\t" + "a") {
        case "\n\t":
            if (!v2) {
                v32 := GenS4{F0: "foo bar"}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 5; fi33++ {
                    for fi34 := 0; fi34 < 18; fi34++ {
                        var v35 int32 = 42
                        v36 := GenS5{F0: (v1 && true), F1: 89}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = fi33
                }
                var v37 uint32 = 2147483647
                switch ("" + ("a" + "foo bar")) {
                case "\n\t":
                    var v38 float32 = v5
                    _ = v38
                case "":
                    v39 := GenS3{F0: v29, F1: uint16(v4)}
                    _ = v39
                case "hello":
                    var v40 bool = false
                    _ = v40
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v37
            }
        case "a":
            v3 = append(v3, "")
        default:
            var v41 uint16 = 255
            _ = v41
        }
        fmt.Println(v4, v2, v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v42 int8 = 42
    v4 = (^(v4))
    switch "foo bar" {
    case "":
        fmt.Println(v1, v20, v5)
    case "a":
        v43 := GenS5{F0: ((^(v20)) >= 128), F1: v42}
        _ = v43
    }
    v42 = v42
    var v44 uint = 1
    v3 = append(v3, "a")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v42 = int8(v20)
    var v45 int = 2147483647
    var v46 uint64 = 196
    var v47 uint8 = genFn1((v29 ^ (v29 << ((559) & 63))), (0 != (v46 * v46)))
    var v48 uint64 = (v46 + 589)
    for fi49 := 0; fi49 < 7; fi49++ {
        v29 -= 464
        v3 = append(v3, "hello")
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}

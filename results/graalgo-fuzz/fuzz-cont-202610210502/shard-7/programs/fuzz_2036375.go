package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0(p0 int8) int64 {
    fmt.Println(p0)
    p0 += (p0 | -70)
    fmt.Println(p0)
    _ = p0
    return -9223372036854775808
}

func main() {
    v1 := []int{128}
    v2 := []int{}
    v3 := []string{(("foo bar" + (("a" + ("a" + ("" + (("hello" + ((("hello" + ("hello" + ("" + ""))) + ((("a" + ("" + "\n\t")) + "foo bar") + (("" + "a") + ("hello" + "hello")))) + "\n\t")) + "hello")))) + (((("a" + "\n\t") + (("a" + "foo bar") + "hello")) + (("foo bar" + "\n\t") + "foo bar")) + ""))) + ("" + "a")), ("a" + ("hello" + ((("a" + ("foo bar" + ("\n\t" + ("a" + "\n\t")))) + "hello") + ("foo bar" + "a"))))}
    var v4 string = ""
    v5 := []uint32{4294967295, 596, 42, 2147483647}
    var v6 int16 = 1
    v4 = "hello"
    var v7 uint64 = 524
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint = 2
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int32 = 343
            var v13 float64 = 3.14159
            v12 = (^(v12))
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: (v4 + "\n\t"), F1: 65535}
        _ = fi10
        _ = v14
    }
    var v15 int8 = 42
    var v16 uint8 = 2
    if (true && (!false)) {
        fmt.Println(v16, v16, v4)
    }
    v17 := GenS2{F0: (v6 ^ (-(v6))), F1: 2147483647}
    switch -962 {
    case 4294967295:
        switch ("hello" + "a") {
        case "\n\t":
            switch 42 {
            case -639:
                v1 = append(v1, 4294967295)
            default:
                switch 127 {
                case -2147483648:
                    v16 = (^(v16))
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        case "foo bar":
            v4 = ("" + v4)
        case "hello":
            if (2 == 42) {
                v3 = append(v3, "a")
                for fi18 := 0; fi18 < 5; fi18++ {
                    v6 = v6
                    var v19 bool = true
                    _ = fi18
                    _ = v19
                }
                if true {
                    v20 := GenS1{F0: v4, F1: 65535}
                    _ = v20
                }
            } else {
                fmt.Println(v16, v4, v4)
            }
        default:
            var v21 bool = false
            _ = v21
        }
    case 722:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 1:
        for fi22 := 0; fi22 < 15; fi22++ {
            v6 = (v6 & (v6 - v6))
            _ = fi22
        }
    }
    if false {
        for fi23 := 0; fi23 < 8; fi23++ {
            fmt.Println(v6)
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        if true {
            var v25 int64 = genFn0((v15 * (v15 ^ 0)))
            _ = v25
        }
        fmt.Println(v6)
        fmt.Println(fi24, v16)
        _ = fi24
    }
    switch 32767 {
    case 512:
        var v26 float32 = -2.25
        _ = v26
    default:
        fmt.Println(v7, v16, v15)
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + v4)
    }
    var v27 int64 = genFn0(v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

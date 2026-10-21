package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 bool) int8 {
    var v6 float32 = 3.14159
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 17; fi8++ {
        if (!p1) {
            v7 = v7
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []string{"hello", (("a" + ("hello" + "hello")) + (("hello" + "a") + "a")), "a"}
    v2 := make(map[uint32]uint32)
    var v3 float32 = -2.25
    v4 := make(map[string]int)
    var v5 uint64 = 132
    if (false || false) {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            fmt.Println(v3, v5)
        case "\n\t":
            v9 := GenS1{F0: 2147483647, F1: -558}
            _ = v9
        case "":
            if (74 < 42) {
                var v10 uint64 = (v5 / ((v5) | 1))
                if (true || false) {
                    switch 9223372036854775807 {
                    case 32767:
                        delete(v4, "hello")
                    case 395:
                        var v11 float64 = 3.14159
                        _ = v11
                    case 127:
                        _ = len(v2)
                    }
                    var v12 string = ("" + ("a" + "a"))
                    var v13 int = 128
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        default:
            switch "hello" {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int64 = 9223372036854775807
                        _ = fi14
                        _ = v15
                    }
                }
            default:
                v4["hello"] = 698
            }
        }
        fmt.Println(v5, v3)
    } else {
        fmt.Println(v5)
        v1 = append(v1, "")
        v16 := GenS2{F0: ((("a" + "") + (("foo bar" + "a") + "a")) + ("\n\t" + ("a" + "a")))}
        _ = v16
    }
    v5 = 1
    for fi17 := 0; fi17 < 20; fi17++ {
        v3 -= v3
        _ = fi17
    }
    v5 = 616
    var v18 uint64 = (v5 | v5)
    var v19 string = "a"
    var v20 float32 = (v3 + v3)
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int = (fi21 / (((^(fi21))) | 1))
        _ = len(v2)
        _ = fi21
        _ = v22
    }
    v23 := v4["\n\t"]
    for fi24 := 0; fi24 < 20; fi24++ {
        if false {
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(v5, v20)
                _ = fi25
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                v19 = (("a" + v19) + ("foo bar" + "foo bar"))
                _ = fi26
            }
        } else {
            v27 := GenS2{F0: v19}
            fmt.Println(v18, v3)
            _ = v27
        }
        _ = fi24
    }
    if false {
        if ((0.1 * v3) == (1e10 * float32(0.1))) {
            fmt.Println(v20)
        }
        fmt.Println(v19, v19)
        v2[2147483647] = 2147483647
    } else {
        for fi28 := 0; fi28 < 14; fi28++ {
            v19 = v19
            for fi29 := 0; fi29 < 12; fi29++ {
                fmt.Println(fi28)
                var v30 float32 = (3.14159 + -2.25)
                v30 = 1.5
                _ = fi29
                _ = v30
            }
            _ = len(v1)
            _ = fi28
        }
        v20 += (float32(1e10) - v3)
    }
    v4[(((v19 + ((("\n\t" + (v19 + "\n\t")) + (v19 + ((v19 + v19) + "foo bar"))) + v19)) + (("hello" + "") + (v19 + "a"))) + v19)] = v23
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(v5, v20, v3)
        if (!true) {
            _ = len(v1)
        } else {
            if (!true) {
                for fi32 := 0; fi32 < 1; fi32++ {
                    v23 = 127
                    v33 := GenS1{F0: 398, F1: 42}
                    _ = fi32
                    _ = v33
                }
                var v34 int8 = 82
                var v35 float64 = 3.14159
                _ = v34
                _ = v35
            } else {
                var v36 bool = true
                _ = v36
            }
            v19 = "\n\t"
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}

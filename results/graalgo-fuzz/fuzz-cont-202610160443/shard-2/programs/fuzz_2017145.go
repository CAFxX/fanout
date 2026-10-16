package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (((3.14159 + float32(3.14159)) - (1e10 - 1e10)) * float32(3.14159))
    v3 := []string{"foo bar", ((("" + "hello") + "hello") + "")}
    var v4 bool = false
    var v5 float64 = 1.5
    if ("" != "a") {
        v6 := GenS1{F0: false, F1: 42, F2: 1}
        delete(v1, ("hello" + ((("a" + "\n\t") + "foo bar") + "")))
        _ = v6
    }
    v7 := GenS2{F0: (v2 == v2), F1: "\n\t"}
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            v5 -= v5
            var v9 int64 = -1
            _ = v9
        } else {
            _ = len(v3)
            v1[""] = 2147483647
        }
        _ = fi8
    }
    v10 := GenS2{F0: (1.5 >= ((0.1 - (v5 * (v5 + 0.1))) + 0.0)), F1: "\n\t"}
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v2)
        v4 = (false && v4)
        if ((138 == 0) && true) {
            v2 = -2.25
            var v12 uint16 = 895
            _ = v12
        } else {
            var v13 uint32 = 2147483648
            if false {
                v5 += v5
                v4 = false
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        _ = len(v1)
        var v15 bool = (255 > 255)
        _ = fi14
        _ = v15
    }
    if (((!(true && v4)) || (!false)) || true) {
        fmt.Println(v5, v5, v5)
    }
    v5 -= v5
    v16 := GenS2{F0: ((true && (0 == 805)) || (128 >= 127)), F1: "a"}
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v17 := GenS1{F0: ((v4 && v4) || v4), F1: 990, F2: -159}
    fmt.Println(v2)
    if ((!(false && (-1 <= 362))) || true) {
        v18 := GenS3{F0: 1, F1: "a", F2: 42}
        switch "" {
        case "a":
            v2 -= v2
        case "foo bar":
            for fi19 := 0; fi19 < 6; fi19++ {
                if v4 {
                    delete(v1, "\n\t")
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 int8 = 42
                        _ = fi20
                        _ = v21
                    }
                    if v4 {
                        var v22 int64 = -1
                        _ = v22
                    }
                }
                fmt.Println(v2, fi19, v2)
                _ = fi19
            }
        case "":
            var v23 int = 127
            _ = v23
        }
        _ = v18
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        delete(v1, ("hello" + ((("hello" + "\n\t") + "") + "")))
        if v4 {
            v5 = (0.0 * 3.14159)
            v4 = true
        }
        switch "hello" {
        case "\n\t":
            var v25 int32 = 0
            _ = v25
        case "hello":
            v4 = (v4 && false)
        default:
            var v26 string = "hello"
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS1{F0: false, F1: 65535, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}

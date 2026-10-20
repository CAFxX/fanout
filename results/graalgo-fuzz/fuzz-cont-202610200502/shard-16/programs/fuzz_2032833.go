package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    v1 := []string{}
    var v2 int8 = 124
    var v3 uint = 4294967295
    v4 := make(map[string]string)
    v5 := []int{}
    if false {
        if false {
            switch ("hello" + "") {
            case "\n\t":
                var v6 int64 = 42
                _ = v6
            case "foo bar":
                var v7 int8 = 0
                _ = v7
            case "a":
                if ((255 >= 769) && (95 == 255)) {
                    var v8 uint16 = 1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v8
                }
            }
            for fi9 := 0; fi9 < 14; fi9++ {
                v3 -= (v3 << (((^(v3))) & 63))
                _ = fi9
            }
            v2 = (v2 * -78)
        } else {
            var v10 float32 = ((1e10 - (1e10 + ((-2.25 - 0.0) - ((0.1 + float32((((((1.5 * 0.0) / float64(0.1)) + 0.0) / 1.0) + (((3.14159 * -2.25) / -2.25) / 1.0)))) * 1.5)))) * -2.25)
            var v11 bool = (65535 < 1)
            var v12 int64 = 1
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 float32 = float32(((((3.14159 * float64(0.0)) - (1.5 - (float64(float32(1e10)) - (((1e10 + 1.5) / 1.5) + 0.1)))) + 1.5) * (((-2.25 * ((0.0 + float64(0.0)) - 0.0)) - (-2.25 - 3.14159)) + (1e10 * 3.14159))))
        if true {
            var v14 int16 = 1000
            _ = v14
        }
        _ = v13
    } else {
        _ = len(v5)
        for fi15 := 0; fi15 < 2; fi15++ {
            v2 += (v2 & v2)
            v16 := GenS1{F0: 2, F1: ("a" + (("a" + ("a" + "hello")) + "\n\t"))}
            _ = fi15
            _ = v16
        }
        v2 -= (v2 << ((v3) & 63))
    }
    fmt.Println(v3, v2)
    v17 := GenS1{F0: 42, F1: "hello"}
    v1 = append(v1, "hello")
    v2 -= 127
    v4["foo bar"] = ""
    v4["foo bar"] = ("\n\t" + ("foo bar" + ((("\n\t" + "") + "foo bar") + "\n\t")))
    for fi18 := 0; fi18 < 1; fi18++ {
        v3 = v3
        v2 -= v2
        _ = fi18
    }
    if true {
        var v19 float64 = 0.1
        switch ("foo bar" + "a") {
        case "hello":
            var v20 int64 = 9223372036854775807
            _ = v20
        case "":
            _ = len(v4)
        }
        switch 914 {
        case -1:
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        case 284:
            v3 = 255
        case 104:
            fmt.Println(v3)
        default:
            v19 = v19
        }
        _ = v19
    } else {
        v2 += v2
        v3 += v3
    }
    v21 := GenS1{F0: 873, F1: (("foo bar" + ("a" + "\n\t")) + "\n\t")}
    if true {
        if (!(42 > 1)) {
            var v22 int8 = (v2 & v2)
            _ = v22
        } else {
            v3 -= (^(v3))
            if true {
                if false {
                    v23 := GenS1{F0: 705, F1: "\n\t"}
                    _ = v23
                }
            } else {
                var v24 float64 = 1.5
                fmt.Println(v24)
                v2 = v2
                _ = v24
            }
        }
        var v25 uint = (v3 / ((v3) | 1))
        _ = v25
    } else {
        v26 := GenS1{F0: 18446744073709551615, F1: "hello"}
        _ = v26
    }
    fmt.Println(v2)
    v2 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

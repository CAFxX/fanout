package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int = -9223372036854775808
    var v3 int64 = -522
    v4 := []string{"\n\t", ("hello" + "\n\t"), ("foo bar" + "\n\t"), "a"}
    var v5 uint32 = 812
    if ((false || (((0.1 - (1e10 + 1.5)) <= 1.5) && false)) && true) {
        var v6 float32 = 3.14159
        var v7 string = (("hello" + "") + "a")
        v6 = (((((float32(0.1) * (1.5 / (-2.25 - float32(((float64((v6 * 1.5)) - (1e10 / 1.0)) / float64(-2.25)))))) / 1.0) + v6) / -2.25) - v6)
        _ = v6
        _ = v7
    } else {
        if (true || false) {
            if (int(v5) <= v2) {
                v2 += (v2 | 4294967295)
            } else {
                switch "" {
                case "a":
                    v5 -= (v5 ^ (v5 / (((v5 / (((v5 & (v5 ^ (v5 % ((2147483647) | 1))))) | 1))) | 1)))
                }
                fmt.Println(v5, v2)
                v2 = v2
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                if true {
                    v9 := GenS1{F0: 136, F1: v3}
                    var v10 int16 = 32767
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
            fmt.Println(v3, v3)
        }
        var v11 uint16 = 347
        switch v2 {
        case -1:
            var v12 string = "a"
            _ = v12
        case -286:
            v13 := v1["foo bar"]
            _ = v13
        case -9223372036854775808:
            _ = len(v4)
        }
        _ = v11
    }
    delete(v1, "a")
    var v14 int16 = 0
    _ = len(v4)
    var v15 float32 = (3.14159 + 0.1)
    for fi16 := 0; fi16 < 12; fi16++ {
        _ = len(v1)
        for fi17 := 0; fi17 < 4; fi17++ {
            switch ("a" + "\n\t") {
            case "hello":
                switch "hello" {
                case "a":
                    v15 = 0.1
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    if false {
                        fmt.Println(v3, v5)
                    }
                }
                v3 = v3
                v15 -= (1.5 - v15)
            }
            var v18 bool = false
            _ = fi17
            _ = v18
        }
        fmt.Println(v2)
        _ = fi16
    }
    var v19 string = ""
    fmt.Println(v19, v2, v5)
    fmt.Println(v14)
    var v20 int8 = 100
    var v21 uint = 602
    switch v2 {
    case 9223372036854775807:
        fmt.Println(v15, v21)
    default:
        switch v19 {
        case "foo bar":
            var v22 int = (v2 & v2)
            _ = v22
        case "a":
            fmt.Println(v21, v21, v20)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

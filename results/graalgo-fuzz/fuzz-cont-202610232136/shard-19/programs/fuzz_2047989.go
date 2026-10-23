package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }
type GenS2 struct { F0 string }

func main() {
    var v1 int32 = 0
    v2 := make(map[string]string)
    v3 := []string{"\n\t"}
    var v4 uint = 18446744073709551615
    var v5 int = 32767
    v6 := GenS1{F0: 4294967295, F1: 128, F2: ("\n\t" + ("foo bar" + "hello"))}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint16 = 148
        var v9 float32 = (0.1 + 0.1)
        switch v5 {
        case 0:
            v10 := GenS2{F0: "\n\t"}
            _ = v10
        default:
            v4 += v4
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v1, v4)
    fmt.Println(v1, v5)
    v2["hello"] = ("foo bar" + ("foo bar" + (("hello" + "foo bar") + ("foo bar" + (("foo bar" + (("hello" + "") + "")) + "foo bar")))))
    switch ("\n\t" + "a") {
    case "foo bar":
        var v11 int8 = 100
        _ = v11
    case "a":
        if (628 > 32767) {
            fmt.Println(v5, v5, v4)
            v4 -= (v4 >> ((494) & 63))
        }
    default:
        switch 128 {
        case 32767:
            if false {
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint64 = 2
                    v14 := GenS2{F0: "a"}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            } else {
                var v15 float64 = 0.0
                fmt.Println(v4)
                _ = v15
            }
        case -70:
            var v16 float32 = float32(((0.0 - float64(-2.25)) * 3.14159))
            _ = v16
        case 127:
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint64 = 0
                v1 += (v1 * (v1 / ((v1) | 1)))
                var v19 string = "a"
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            v5 = (fi20 | (fi20 % ((v5) | 1)))
            switch v5 {
            case -34:
                switch "foo bar" {
                case "":
                    v1 += (v1 & (v1 ^ (v1 ^ v1)))
                case "a":
                    v4 -= (v4 | (v4 >> ((v4) & 63)))
                case "foo bar":
                    fmt.Println(v1)
                default:
                    var v21 uint16 = 65535
                    _ = v21
                }
            case -329:
                for fi22 := 0; fi22 < 20; fi22++ {
                    _ = len(v2)
                    if ((-2.25 < 0.1) && ((false || true) || (!false))) {
                        fmt.Println(v5)
                        var v23 uint8 = 255
                        fmt.Println(v5, v4, v5)
                        _ = v23
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v4 = (v4 * (v4 >> ((499) & 63)))
                        if true {
                            v3 = append(v3, "foo bar")
                        }
                        var v24 int = (-(v5))
                        _ = v24
                    }
                    var v25 uint = 64
                    _ = fi22
                    _ = v25
                }
            case 447:
                v1 = 129
            }
        }
        _ = fi20
    }
    v26 := GenS2{F0: (("hello" + "") + "hello")}
    v1 -= v1
    v27 := GenS2{F0: "foo bar"}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

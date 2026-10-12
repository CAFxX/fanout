package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    var v1 int16 = 42
    var v2 uint8 = 127
    v3 := []string{"hello", "hello"}
    var v4 int64 = -1
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: 2147483647}
    fmt.Println(v4)
    v7 := GenS2{F0: ((1e10 >= 1e10) || (((("\n\t" + "a") + "") != (("hello" + ("foo bar" + ("foo bar" + "hello"))) + "\n\t")) || false)), F1: 895}
    v4 += v4
    var v8 uint64 = 42
    v9 := GenS2{F0: false, F1: 4294967295}
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int16 = v1
        fmt.Println(fi10, v11)
        _ = fi10
        _ = v11
    }
    switch 0 {
    case 9223372036854775807:
        v8 -= v8
    case -2147483648:
        fmt.Println(v1, v2, v2)
    default:
        var v12 uint8 = v2
        _ = v12
    }
    v2 = v2
    var v13 float64 = float64((1e10 / float32(1.5)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "foo bar" {
    case "":
        switch 294 {
        case 127:
            for fi14 := 0; fi14 < 12; fi14++ {
                if true {
                    switch int(v8) {
                    case -2147483648:
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                    case 2147483647:
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v8 = 154
                            v15 := GenS1{F0: -26}
                            _ = v15
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v16 := GenS1{F0: -2147483648}
                        _ = v16
                    case "":
                        v13 = (v13 + (v13 + v13))
                    case "a":
                        switch ("hello" + "\n\t") {
                        case "a":
                            fmt.Println(v13)
                        case "hello":
                            v3 = append(v3, "a")
                        case "foo bar":
                            fmt.Println(v2, fi14, v8)
                        }
                    }
                    switch "hello" {
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = ("a" + ("a" + ("hello" + (("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + ""))))
                        }
                    default:
                        fmt.Println(v13)
                    }
                }
                _ = fi14
            }
        case 111:
            v3 = append(v3, "\n\t")
        }
    case "\n\t":
        v17 := GenS1{F0: 2147483647}
        _ = v17
    default:
        var v18 uint = 787
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

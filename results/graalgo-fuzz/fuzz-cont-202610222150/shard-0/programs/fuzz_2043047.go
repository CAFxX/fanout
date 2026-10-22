package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{42}
    var v2 bool = false
    v3 := make(map[string]uint32)
    v4 := []int{42, 32767, 127}
    switch "\n\t" {
    case "":
        if v2 {
            delete(v3, (("" + "\n\t") + (((((("hello" + "foo bar") + "") + ((((("foo bar" + (("a" + "") + "hello")) + "") + (("foo bar" + ("foo bar" + "\n\t")) + "hello")) + "") + "")) + "") + "a") + ("a" + ""))))
            var v5 uint32 = 65535
            _ = v5
        } else {
            switch "" {
            case "":
                v6 := GenS1{F0: -1, F1: 42}
                _ = v6
            default:
                if true {
                    for fi7 := 0; fi7 < 8; fi7++ {
                        fmt.Println(fi7, fi7)
                        v8 := GenS1{F0: -317, F1: 2}
                        v9 := GenS1{F0: -9, F1: 2}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    v2 = (3.14159 > 0.1)
                }
            }
        }
    }
    _ = len(v1)
    for fi10 := 0; fi10 < 2; fi10++ {
        v2 = true
        if len(v1) > 0 {
            v1[0] = 127
        }
        _ = fi10
    }
    v11 := GenS1{F0: 0, F1: 65535}
    if (!(!v2)) {
        v4 = append(v4, 32768)
        var v12 int32 = 36
        _ = v12
    }
    var v13 uint = 65535
    v13 -= v13
    if v2 {
        fmt.Println(v13, v13)
    } else {
        if (!((true && true) || false)) {
            var v14 uint16 = 65535
            switch 2147483647 {
            case -9223372036854775808:
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := v3["a"]
                    _ = fi15
                    _ = v16
                }
            default:
                v13 = v13
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v14
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch -549 {
            case -1:
                fmt.Println(v13, v13)
            case -951:
                v3["hello"] = 923
            }
            switch ("a" + ("\n\t" + "\n\t")) {
            case "":
                v2 = (!(-32768 == -196))
            case "foo bar":
                var v17 int64 = -1
                _ = v17
            case "hello":
                fmt.Println(v2, v13)
            }
        }
        if v2 {
            v13 = 2
        } else {
            fmt.Println(v13)
            switch "foo bar" {
            case "":
                v18 := GenS2{F0: 128}
                _ = v18
            case "\n\t":
                var v19 bool = true
                _ = v19
            case "a":
                if true {
                    fmt.Println(v13, v2)
                    if v2 {
                        v20 := GenS1{F0: 2147483647, F1: uint32(v13)}
                        _ = v20
                    } else {
                        var v21 uint8 = 42
                        _ = v21
                    }
                    fmt.Println(v2)
                }
            default:
                fmt.Println(v2)
            }
        }
    }
    v2 = true
    var v22 float64 = (1e10 * -2.25)
    v2 = (((v22 - v22) + 1.5) > v22)
    var v23 int = -298
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}

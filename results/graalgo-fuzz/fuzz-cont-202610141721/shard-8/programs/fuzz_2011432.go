package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }

func main() {
    v1 := []uint32{42, 698, 2147483647, 130}
    v2 := []string{(("foo bar" + (("hello" + (("hello" + ("\n\t" + ("a" + "hello"))) + (("foo bar" + "hello") + ("" + "\n\t")))) + "a")) + "\n\t")}
    var v3 int = -455
    var v4 int = v3
    if false {
        if (true || (!(true || (v3 != v3)))) {
            if true {
                var v5 int = v3
                _ = v5
            } else {
                v6 := GenS1{F0: false, F1: 0, F2: "\n\t"}
                v3 = -2147483648
                _ = v6
            }
            for fi7 := 0; fi7 < 5; fi7++ {
                v3 = 1
                if (!(!true)) {
                    var v8 int16 = -32768
                    _ = v8
                }
                _ = fi7
            }
            v4 = (v3 ^ v3)
        }
    }
    fmt.Println(v4, v4)
    if (1 != 4294967295) {
        _ = len(v2)
    } else {
        var v9 bool = false
        switch 2147483647 {
        case 2147483647:
            var v10 string = ""
            _ = v10
        default:
            v3 = (v4 | (v4 + v4))
        }
        _ = v9
    }
    v4 = 9223372036854775807
    var v11 uint16 = 42
    v4 += (v4 / ((361) | 1))
    var v12 bool = false
    v11 = (v11 % ((v11) | 1))
    if false {
        v13 := GenS1{F0: v12, F1: 4294967295, F2: "a"}
        _ = v13
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "foo bar":
        v14 := GenS1{F0: (-32768 <= 42), F1: 1, F2: ("hello" + ("" + "hello"))}
        _ = v14
    case "":
        var v15 int64 = -9223372036854775808
        _ = v15
    case "hello":
        var v16 float64 = 1.5
        _ = v16
    default:
        var v17 uint8 = 115
        _ = v17
    }
    var v18 uint32 = 65535
    var v19 float64 = (3.14159 / (float64(1e10) - 0.1))
    if v12 {
        fmt.Println(v19, v12, v3)
        var v20 bool = (-736 < -278)
        v18 -= 304
        _ = v20
    }
    v19 -= -2.25
    switch (v4 ^ (v4 % ((303) | 1))) {
    case -1:
        v21 := GenS2{F0: 1, F1: 42, F2: 255}
        _ = v21
    default:
        if (834 >= -20) {
            var v22 uint = 18446744073709551615
            for fi23 := 0; fi23 < 15; fi23++ {
                if v12 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    _ = len(v1)
                }
                if ((^(v11)) > v11) {
                    v22 = v22
                    v22 -= 0
                }
                _ = fi23
            }
            var v24 int32 = 1
            _ = v22
            _ = v24
        } else {
            v4 = v4
        }
    }
    v25 := GenS1{F0: (((!(!(!(-128 >= -128)))) || v12) || true), F1: uint(v18), F2: ("hello" + "\n\t")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}

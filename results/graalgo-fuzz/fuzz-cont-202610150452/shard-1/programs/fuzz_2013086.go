package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint) int64 {
    fmt.Println(p1, p0, p0)
    var v4 bool = (0 >= 255)
    p0 -= (^(p0))
    switch ("\n\t" + "\n\t") {
    case "hello":
        var v5 bool = (((81 >= -1) || v4) || (!(4294967295 < uint32(p1))))
        _ = v5
    case "":
        switch 14 {
        case -2147483648:
            switch -2147483648 {
            case 0:
                var v6 int = 128
                _ = v6
            default:
                switch "\n\t" {
                case "hello":
                    var v7 uint16 = 42
                    _ = v7
                default:
                    v8 := GenS1{F0: "\n\t", F1: (("\n\t" + "\n\t") + ("\n\t" + (("" + "\n\t") + ((("\n\t" + ("a" + "")) + ("a" + ("foo bar" + "hello"))) + "hello")))), F2: int16(p1)}
                    _ = v8
                }
            }
        case 1:
            var v9 uint16 = p1
            _ = v9
        case 32768:
            p0 += 996
        default:
            var v10 float32 = (0.1 + 3.14159)
            _ = v10
        }
    case "a":
        if ((!(v4 || true)) || (!(!("" > ((("" + ("hello" + "a")) + ("" + "\n\t")) + ((("foo bar" + (("foo bar" + "") + "hello")) + ("a" + "\n\t")) + "hello")))))) {
            fmt.Println(p2, v4)
        }
    default:
        p0 = 614
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1() string {
    var v11 bool = (false || false)
    var v12 int16 = -32768
    var v13 int8 = -29
    var v14 int32 = 2147483647
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return ("foo bar" + "a")
}

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 42
    v3 := []string{"\n\t"}
    v2 = -432
    fmt.Println(v2, v2)
    v1 -= v1
    v15 := GenS1{F0: "\n\t", F1: "a", F2: 0}
    v2 = -195
    v1 = v1
    v2 = (^(v2))
    for fi16 := 0; fi16 < 11; fi16++ {
        switch "\n\t" {
        case "":
            v17 := GenS1{F0: "hello", F1: ("\n\t" + (((((((("foo bar" + ("" + (((("foo bar" + "hello") + ((("hello" + "foo bar") + "foo bar") + "")) + "hello") + ("hello" + "")))) + "\n\t") + "\n\t") + "a") + ("" + "\n\t")) + "\n\t") + "\n\t") + ("\n\t" + "hello"))), F2: -1}
            _ = v17
        }
        var v18 uint16 = 42
        _ = fi16
        _ = v18
    }
    v19 := GenS1{F0: "\n\t", F1: (("" + "hello") + (("hello" + "foo bar") + "a")), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint = 283
    v2 := []int{32767}
    var v3 uint = (v1 & 65535)
    var v4 uint32 = 2147483647
    v5 := make(map[string]string)
    v6 := []int{}
    switch "a" {
    case "":
        _ = len(v2)
    case "hello":
        v1 = v1
    case "\n\t":
        if false {
            var v7 uint = 42
            v8 := GenS1{F0: 2}
            _ = v7
            _ = v8
        } else {
            var v9 int64 = 9223372036854775807
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := v5[(("\n\t" + "a") + ("foo bar" + "hello"))]
                v3 += (v3 - 295)
                _ = fi10
                _ = v11
            }
            _ = v9
        }
    }
    fmt.Println(v1, v1, v4)
    var v12 int16 = -1
    fmt.Println(v12, v12, v1)
    var v13 int8 = 68
    if (false || (true || false)) {
        switch "foo bar" {
        case "\n\t":
            v13 += (v13 >> ((v1) & 63))
        case "foo bar":
            var v14 int = 42
            _ = v14
        default:
            fmt.Println(v12)
        }
        var v15 uint16 = 65535
        if false {
            switch ("hello" + "hello") {
            case "a":
                switch 313 {
                case -925:
                    fmt.Println(v15)
                case -9223372036854775808:
                    v12 = (v12 & 437)
                case -1:
                    var v16 float32 = (((-2.25 - (-2.25 - 3.14159)) + (0.1 + 0.0)) + (0.1 / (0.1 - 0.0)))
                    _ = v16
                }
            }
            if false {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v12 -= -295
                    _ = fi17
                }
                var v18 uint64 = 2
                v5["hello"] = ((((("foo bar" + "") + ("\n\t" + ("hello" + (("hello" + (("" + "") + "a")) + "foo bar")))) + (("" + "") + "a")) + "\n\t") + "")
                _ = v18
            }
        }
        _ = v15
    } else {
        var v19 int64 = -635
        _ = v19
    }
    var v20 int8 = -128
    v1 = (^(v1))
    var v21 uint64 = uint64(v4)
    var v22 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

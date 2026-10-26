package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint8 }

func main() {
    v1 := []string{("a" + (("" + (((((("\n\t" + ("\n\t" + "hello")) + ("a" + ("foo bar" + "foo bar"))) + ((("hello" + "foo bar") + "hello") + ((("\n\t" + "a") + ("" + "\n\t")) + "a"))) + "") + ("hello" + ("\n\t" + ("hello" + "\n\t")))) + ((("foo bar" + "") + "a") + "foo bar"))) + "")), ("hello" + "a"), ""}
    var v2 uint16 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    switch "foo bar" {
    case "\n\t":
        if (1 > -128) {
            v2 -= v2
            switch "a" {
            case "\n\t":
                switch 354 {
                case 32768:
                    v2 = 65535
                case -357:
                    v2 = (v2 >> ((1) & 63))
                }
            case "":
                delete(v3, "")
            default:
                v2 += v2
            }
        }
    case "":
        fmt.Println(v2)
    }
    v5 := GenS1{F0: 108, F1: 423}
    v6 := GenS2{F0: 4294967295, F1: 42, F2: 255}
    switch 1 {
    case -471:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 32768:
        for fi7 := 0; fi7 < 12; fi7++ {
            if false {
                var v8 uint32 = 580
                var v9 int64 = int64(fi7)
                _ = v8
                _ = v9
            }
            _ = fi7
        }
    default:
        v2 -= v2
    }
    var v10 float32 = 1e10
    _ = len(v4)
    var v11 float32 = 0.0
    v12 := v4[49]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

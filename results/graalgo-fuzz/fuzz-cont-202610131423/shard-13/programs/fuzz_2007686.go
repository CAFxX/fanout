package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []string{"", ("a" + "a"), "\n\t", (("foo bar" + "foo bar") + "a")}
    var v2 uint16 = 32767
    v3 := []string{"\n\t", ""}
    v4 := []string{(("\n\t" + "foo bar") + ("foo bar" + ("foo bar" + "foo bar"))), ("" + (("hello" + ("foo bar" + "a")) + "")), "\n\t"}
    v5 := make(map[string]string)
    v2 = (v2 >> ((42) & 63))
    var v6 float64 = float64(1.5)
    switch "a" {
    case "hello":
        switch "a" {
        case "hello":
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 int64 = 9223372036854775807
                var v9 uint16 = v2
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 int16 = -32768
                    v12 := GenS1{F0: (fi10 - fi10), F1: ("a" + (("\n\t" + "\n\t") + ""))}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    default:
        v13 := GenS1{F0: 2147483647, F1: "\n\t"}
        _ = v13
    }
    var v14 int32 = 459
    fmt.Println(v6, v14, v14)
    v15 := GenS2{F0: 0, F1: (v2 % (((v2 & (^(v2)))) | 1))}
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v14)
        v17 := GenS1{F0: (fi16 + (-(fi16))), F1: "\n\t"}
        v6 += -2.25
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            delete(v5, "")
            if true {
                v14 -= v14
                var v20 int8 = 0
                _ = v20
            }
            var v21 uint = 4294967295
            _ = fi19
            _ = v21
        }
        var v22 uint8 = 255
        _ = fi18
        _ = v22
    }
    var v23 uint64 = 86
    _ = len(v5)
    var v24 int16 = 1
    var v25 int16 = v24
    v25 = (v24 - int16(v2))
    v24 = (^(v25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

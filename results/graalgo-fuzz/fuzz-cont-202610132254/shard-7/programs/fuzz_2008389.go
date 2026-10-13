package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 uint8, p1 uint16, p2 string) int16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint32 = 4294967295
        var v7 int64 = 9223372036854775807
        var v8 int = fi5
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 255
    var v10 int8 = 127
    v11 := GenS1{F0: p2, F1: -2147483648}
    v12 := GenS1{F0: "a", F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 1
    v2 := []string{("foo bar" + ((("" + "a") + (("hello" + ((("\n\t" + "") + ("\n\t" + "\n\t")) + "foo bar")) + "hello")) + "\n\t")), "hello", "a"}
    var v3 int8 = (v1 | -1)
    v4 := make(map[string]bool)
    v3 = v1
    v13 := GenS1{F0: "foo bar", F1: 432}
    var v14 int32 = 0
    delete(v4, "")
    var v15 uint32 = 2
    if (65535 == 905) {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int8 = (^(v1))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi16
            _ = v17
        }
        switch "" {
        case "foo bar":
            switch ("\n\t" + "foo bar") {
            case "hello":
                fmt.Println(v14)
            case "a":
                v18 := GenS1{F0: "", F1: 454}
                _ = v18
            case "\n\t":
                fmt.Println(v14)
            default:
                var v19 int64 = -1
                _ = v19
            }
        case "hello":
            fmt.Println(v14)
        case "\n\t":
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int8 = 126
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = fi20
                _ = v21
            }
        }
        fmt.Println(v15)
    }
    var v22 bool = (v3 > v1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}

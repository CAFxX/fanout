package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
func genFn0(p0 int32) uint8 {
    var v4 int8 = 79
    fmt.Println(v4, v4)
    var v5 float32 = 0.1
    _ = p0
    _ = v4
    _ = v5
    return 199
}

func main() {
    var v1 uint = 672
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    if false {
        fmt.Println(v1)
        var v6 int8 = 127
        if true {
            fmt.Println(v6)
            v6 = (v6 & 100)
            v7 := GenS1{F0: (^(v1)), F1: 1000, F2: ("foo bar" + ("" + ("\n\t" + ("" + ""))))}
            _ = v7
        }
        _ = v6
    } else {
        v1 += 882
        switch ("foo bar" + "a") {
        case "hello":
            v2[-87] = true
        }
        var v8 uint8 = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v1, fi9)
        _ = fi9
    }
    if true {
        _ = len(v3)
        switch "\n\t" {
        case "\n\t":
            switch -9223372036854775808 {
            case -134:
                v1 = v1
            case 32768:
                delete(v3, -9223372036854775808)
            case 549:
                var v10 bool = false
                _ = v10
            }
        case "foo bar":
            var v11 int16 = -32768
            _ = v11
        }
        v1 = 18446744073709551615
    }
    v12 := GenS1{F0: (v1 % ((65535) | 1)), F1: 0, F2: ("\n\t" + "hello")}
    v1 = v1
    var v13 int64 = 42
    switch (("" + "a") + ("a" + ("\n\t" + "hello"))) {
    case "foo bar":
        v2[-9223372036854775808] = true
    default:
        v13 -= 838
    }
    var v14 int16 = 1
    _ = len(v2)
    v15 := GenS1{F0: v1, F1: v14, F2: "hello"}
    var v16 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}

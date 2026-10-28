package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint64 = 488
    var v2 bool = (!true)
    v3 := make(map[string]int)
    var v4 int32 = 0
    delete(v3, "hello")
    switch "hello" {
    case "":
        var v5 int16 = -1
        _ = v5
    default:
        fmt.Println(v1, v1)
    }
    v2 = true
    if (("\n\t" + "a") < "") {
        var v6 int32 = v4
        v4 = 555
        _ = v6
    }
    var v7 uint = 188
    switch -989 {
    case 127:
        if v2 {
            v8 := []int{128}
            _ = v8
        }
    case -9223372036854775808:
        var v9 int8 = 82
        _ = v9
    }
    var v10 float64 = 0.0
    v11 := []int{}
    for fi12 := 0; fi12 < 11; fi12++ {
        switch (("" + ("\n\t" + ("foo bar" + ("foo bar" + "")))) + "\n\t") {
        case "hello":
            var v13 uint32 = 2147483648
            _ = v13
        case "a":
            v2 = (v7 != (v7 * 65535))
        case "foo bar":
            var v14 uint8 = 1
            _ = v14
        }
        _ = fi12
    }
    v15 := GenS1{F0: (v4 | 2147483647), F1: (^(v4))}
    var v16 string = "hello"
    var v17 int64 = int64(v4)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    switch 32767 {
    case 9223372036854775807:
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 bool = v2
            switch v16 {
            case "\n\t":
                fmt.Println(v16, v2)
            default:
                var v20 int16 = 32767
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

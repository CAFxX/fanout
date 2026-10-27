package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 uint32 = 42
    if v2 {
        delete(v1, "")
    } else {
        v4 := GenS1{F0: -2147483648}
        _ = v4
    }
    v3 -= 2147483648
    fmt.Println(v3)
    if v2 {
        switch ("hello" + "") {
        case "hello":
            var v5 string = ("foo bar" + "hello")
            _ = v5
        case "foo bar":
            var v6 int = 128
            _ = v6
        }
        var v7 uint32 = (v3 & 2147483647)
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint32 = (v3 / ((v3) | 1))
            var v10 uint64 = 18446744073709551615
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        var v11 float64 = 3.14159
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if (-1 == 42) {
            var v13 uint8 = 1
            var v14 bool = (1 != 1)
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v2 = v2
        v3 = v3
        _ = fi15
    }
    if ("hello" >= "a") {
        fmt.Println(v3, v3)
    }
    v3 = v3
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 float64 = 1.5
        if v2 {
            v3 = 4294967295
            switch fi16 {
            case -9223372036854775808:
                var v18 int32 = 2147483647
                _ = v18
            }
        }
        var v19 bool = (210 == 18446744073709551615)
        _ = fi16
        _ = v17
        _ = v19
    }
    var v20 uint64 = 911
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}

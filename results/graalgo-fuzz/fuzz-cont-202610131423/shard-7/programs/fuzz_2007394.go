package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = 1
    var v2 string = (("a" + "foo bar") + "hello")
    var v3 uint16 = 2
    v4 := []string{"a", "", "a", v2}
    var v5 int8 = 42
    v6 := make(map[uint32]string)
    fmt.Println(v2, v3, v2)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint64 = 42
        var v9 uint64 = v8
        v10 := GenS1{F0: v2}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (1 == 1) {
        switch ("hello" + ("" + "hello")) {
        case "foo bar":
            v11 := GenS1{F0: "\n\t"}
            _ = v11
        }
        var v12 uint16 = v3
        _ = v12
    }
    v3 += 255
    var v13 int = 32768
    v3 -= v3
    switch v2 {
    case "hello":
        switch v2 {
        case "":
            var v14 float32 = (0.0 * 1.5)
            _ = v14
        case "\n\t":
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint = 651
                var v17 int8 = 1
                _ = fi15
                _ = v16
                _ = v17
            }
        case "foo bar":
            _ = len(v6)
        default:
            var v18 float32 = -2.25
            _ = v18
        }
    }
    v6[2147483647] = ""
    v2 = "a"
    v19 := GenS1{F0: ""}
    fmt.Println(v1, v13, v5)
    switch v13 {
    case -9223372036854775808:
        switch (v13 - (^(v13))) {
        case 32768:
            switch (v2 + (v2 + (v2 + "\n\t"))) {
            case "a":
                _ = len(v6)
            default:
                fmt.Println(v2, v5, v13)
            }
        case -1:
            v2 = "foo bar"
        case 2147483647:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    var v20 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

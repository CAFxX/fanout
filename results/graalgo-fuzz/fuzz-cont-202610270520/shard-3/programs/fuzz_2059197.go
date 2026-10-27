package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 int16 = 736
    v3 := make(map[uint32]int)
    var v4 float64 = ((float64(3.14159) * (1.5 - -2.25)) * 1e10)
    v5 := []string{v1, "hello", v1, "a"}
    var v6 float64 = v4
    v7 := GenS1{F0: 46}
    if false {
        if ((!false) || (42 < -32768)) {
            v6 = v6
            var v8 int64 = -324
            _ = v8
        }
    }
    var v9 uint = 4294967295
    v9 = v9
    var v10 uint32 = uint32(v9)
    fmt.Println(v9)
    if false {
        switch (v1 + v1) {
        case "\n\t":
            fmt.Println(v6, v2)
        }
        var v11 uint32 = 564
        var v12 bool = true
        _ = v11
        _ = v12
    }
    switch "a" {
    case "hello":
        var v13 int64 = 1
        _ = v13
    case "foo bar":
        var v14 string = (("a" + "") + "hello")
        _ = v14
    }
    switch 4294967295 {
    case 32767:
        var v15 uint64 = 42
        _ = v15
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = v4
    var v16 uint64 = uint64(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

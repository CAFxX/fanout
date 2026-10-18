package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 42
    v3 := []string{"foo bar", "", (("hello" + "") + "hello")}
    v4 := make(map[int]uint32)
    var v5 int32 = (^(v2))
    if (0 <= 9223372036854775807) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        v6 := v1[""]
        _ = v6
    } else {
        var v7 uint16 = 1
        _ = v7
    }
    var v8 int32 = v2
    var v9 int64 = 9223372036854775807
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            _ = fi11
        }
        switch "foo bar" {
        case "hello":
            if len(v3) > 0 {
                v3[0] = ("a" + ("" + ("hello" + "a")))
            }
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            _ = len(v3)
            if len(v3) > 0 {
                v3[0] = ""
            }
            v9 = (v9 ^ v9)
            _ = fi12
        }
        _ = fi10
    }
    switch "foo bar" {
    case "":
        var v13 uint64 = 2
        _ = v13
    case "a":
        var v14 int = 0
        _ = v14
    default:
        if false {
            v2 -= v2
        }
    }
    fmt.Println(v5, v5, v8)
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint8 = 127
        _ = fi15
        _ = v16
    }
    var v17 bool = true
    v2 = (v2 | -149)
    v2 = 1
    var v18 uint32 = 42
    v9 = v9
    v19 := GenS1{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

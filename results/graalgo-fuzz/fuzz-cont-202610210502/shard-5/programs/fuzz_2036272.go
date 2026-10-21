package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }

func main() {
    var v1 uint = 2
    v2 := []string{(("hello" + "hello") + ((("\n\t" + "\n\t") + ("a" + "a")) + ""))}
    var v3 int8 = -1
    v4 := make(map[uint32]string)
    var v5 uint64 = 4294967295
    v6 := GenS1{F0: 32767, F1: "", F2: v3}
    switch 290 {
    case 127:
        var v7 int16 = 1000
        _ = v7
    case 42:
        _ = len(v4)
    case 2147483647:
        switch "hello" {
        case "":
            var v8 int32 = 0
            _ = v8
        case "\n\t":
            var v9 int64 = 1
            _ = v9
        default:
            v4[2147483648] = "a"
        }
    }
    v10 := GenS1{F0: 32767, F1: "a", F2: v3}
    var v11 uint32 = 65535
    if ("hello" < "a") {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if false {
            fmt.Println(v11, v3)
        }
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v13 int32 = -916
        var v14 int8 = v3
        _ = fi12
        _ = v13
        _ = v14
    }
    if true {
        if false {
            var v15 uint16 = 267
            _ = v15
        }
    } else {
        var v16 uint32 = 2147483647
        switch 2147483647 {
        case 42:
            if true {
                v17 := v4[v11]
                var v18 int32 = 0
                v5 -= v5
                _ = v17
                _ = v18
            }
        default:
            var v19 int64 = 42
            _ = v19
        }
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = v16
    }
    v11 = 65535
    var v20 int8 = 127
    fmt.Println(v3, v20, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

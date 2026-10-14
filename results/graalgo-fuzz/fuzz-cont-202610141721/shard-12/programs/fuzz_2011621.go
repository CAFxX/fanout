package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []string{("a" + "a"), ("\n\t" + ("foo bar" + "hello")), (("\n\t" + "foo bar") + "\n\t"), "hello"}
    var v2 uint16 = 65535
    var v3 uint16 = (v2 + 265)
    v4 := make(map[int]uint32)
    var v5 string = "a"
    var v6 uint = 4294967295
    var v7 int16 = 1000
    var v8 uint8 = 255
    v9 := GenS1{F0: -788}
    if false {
        fmt.Println(v3)
    } else {
        var v10 int8 = 13
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v2 -= (^(v3))
        var v12 int64 = 9223372036854775807
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v7 += v7
        _ = fi13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = v5
    v14 := GenS2{F0: 4294967295, F1: 42}
    for fi15 := 0; fi15 < 16; fi15++ {
        v8 -= v8
        v6 = (^(v6))
        _ = fi15
    }
    fmt.Println(v7, v6)
    v4[9223372036854775807] = uint32(v6)
    var v16 int64 = 629
    v1 = append(v1, (((((v5 + v5) + v5) + ("hello" + v5)) + "a") + "a"))
    v3 += (v3 & (^(v2)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

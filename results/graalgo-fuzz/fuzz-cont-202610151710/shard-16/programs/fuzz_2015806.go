package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 uint8 = 128
    var v4 uint64 = 2
    var v5 uint64 = v4
    var v6 uint64 = (^(v4))
    v6 = v5
    v4 -= (v6 >> ((844) & 63))
    v7 := []int{0}
    if false {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 uint16 = 0
            _ = fi8
            _ = v9
        }
    }
    if true {
        if true {
            v4 += (v5 >> ((4294967295) & 63))
        }
    } else {
        v10 := GenS1{F0: ("hello" + ("a" + (((("foo bar" + ("hello" + "foo bar")) + (((("foo bar" + (("hello" + "a") + "hello")) + "\n\t") + "") + "a")) + "a") + ("foo bar" + "a")))), F1: 1, F2: 4294967295}
        var v11 int32 = -764
        v6 -= 18446744073709551615
        _ = v10
        _ = v11
    }
    _ = len(v1)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int = fi12
        _ = fi12
        _ = v13
    }
    switch (("a" + "\n\t") + "foo bar") {
    case "hello":
        var v14 int32 = 42
        _ = v14
    case "foo bar":
        fmt.Println(v5, v3, v4)
    case "":
        var v15 int16 = 510
        _ = v15
    }
    v16 := GenS1{F0: "\n\t", F1: 339, F2: 2147483648}
    if ((0 < -393) && false) {
        switch "a" {
        case "hello":
            fmt.Println(v5, v6)
        }
    } else {
        var v17 uint16 = 0
        _ = v17
    }
    v6 = v5
    if true {
        var v18 int32 = -105
        v6 = (^(v6))
        _ = v18
    } else {
        fmt.Println(v3, v4, v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

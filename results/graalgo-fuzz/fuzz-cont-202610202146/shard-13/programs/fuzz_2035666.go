package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 string = ("hello" + ("hello" + ("foo bar" + ("foo bar" + ("" + "foo bar")))))
    v4 := []int{42}
    var v5 string = ("\n\t" + v3)
    _ = len(v4)
    if false {
        v6 := GenS1{F0: 127, F1: -1}
        _ = v6
    } else {
        switch v3 {
        case "a":
            v7 := GenS1{F0: 128, F1: 4294967295}
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: 1, F1: 2147483647}
            _ = v8
        case "hello":
            fmt.Println(v3, v5)
        }
        fmt.Println(v3, v3)
        var v9 int = 0
        _ = v9
    }
    v10 := v1[((((v3 + (("foo bar" + "") + "a")) + v3) + ("\n\t" + ("foo bar" + "hello"))) + v3)]
    var v11 uint32 = 2147483647
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int32 = 527
        fmt.Println(v3, v13, v10)
        _ = fi12
        _ = v13
    }
    switch -9223372036854775808 {
    case -1:
        if false {
            v5 = ((v10 + "hello") + "a")
        } else {
            switch (v5 + "a") {
            case "":
                fmt.Println(v3)
            case "hello":
                _ = len(v2)
            case "foo bar":
                fmt.Println(v11, v11, v11)
            }
        }
    case 347:
        v3 = ((v5 + ("foo bar" + ("" + v3))) + "a")
    case 32767:
        var v14 uint32 = v11
        _ = v14
    default:
        var v15 uint32 = (v11 << ((372) & 63))
        _ = v15
    }
    fmt.Println(v11)
    if (!true) {
        v3 = ""
    } else {
        var v16 uint16 = 255
        _ = v16
    }
    v17 := GenS1{F0: 128, F1: 791}
    v5 = v5
    fmt.Println(v5, v11, v10)
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    v5 = ("" + v3)
    v5 = (v10 + v10)
    if true {
        var v19 uint = 65535
        v20 := GenS2{F0: -1, F1: -855}
        var v21 bool = (0 != 42)
        _ = v19
        _ = v20
        _ = v21
    }
    if (-32768 > 1) {
        var v22 uint64 = 4294967295
        v3 = (("a" + v3) + (v3 + v10))
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

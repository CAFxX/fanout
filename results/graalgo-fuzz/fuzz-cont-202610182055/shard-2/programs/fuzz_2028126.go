package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int }

func main() {
    var v1 uint8 = 0
    var v2 float32 = 3.14159
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    v2 = v2
    var v7 string = ""
    var v8 uint8 = v1
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v8 = v1
        var v12 int16 = 42
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(fi11, v1)
            fmt.Println(v1)
            var v14 uint = 4294967295
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    switch int(v1) {
    case -1:
        v7 = (("a" + ((("foo bar" + (("hello" + v7) + v7)) + v7) + v7)) + "hello")
    case 627:
        fmt.Println(v1, v1, v1)
    }
    if (1 > 255) {
        switch 874 {
        case 128:
            switch (v7 + v7) {
            case "foo bar":
                v6["foo bar"] = 332
            case "a":
                var v15 float32 = v2
                _ = v15
            }
        case 2147483647:
            switch -2147483648 {
            case 549:
                v8 = v8
            }
        case -1:
            if true {
                fmt.Println(v8, v8, v8)
                v1 = (v8 % ((v1) | 1))
            }
        }
    } else {
        switch -9223372036854775808 {
        case -170:
            v1 = v1
        case -1:
            v16 := []int{216, -875}
            _ = v16
        default:
            fmt.Println(v1, v1)
        }
        v17 := v5[0]
        _ = v17
    }
    v8 = v1
    fmt.Println(v1)
    _ = len(v5)
    v18 := GenS2{F0: 18446744073709551615, F1: 728}
    v2 = v2
    var v19 uint64 = 258
    var v20 int32 = int32(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

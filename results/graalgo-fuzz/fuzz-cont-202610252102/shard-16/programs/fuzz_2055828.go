package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 1000
    var v3 int = 9223372036854775807
    var v4 uint16 = 255
    fmt.Println(v4)
    fmt.Println(v3, v4, v2)
    v5 := GenS1{F0: 170, F1: v4, F2: -3}
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            switch "foo bar" {
            case "hello":
                var v7 int32 = 42
                _ = v7
            case "\n\t":
                var v8 uint8 = 127
                _ = v8
            default:
                v4 = v4
            }
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := []int{(v3 ^ v3), (v3 % ((fi6) | 1)), fi9}
                var v11 int32 = -155
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(v3, v4)
            _ = fi6
        }
        var v12 int64 = -105
        delete(v1, 2147483648)
        _ = v12
    } else {
        var v13 bool = (false && (("\n\t" + "foo bar") >= ""))
        fmt.Println(v13)
        _ = v13
    }
    v2 = v2
    var v14 int8 = -1
    var v15 uint = 1
    if false {
        var v16 uint64 = 240
        v15 = (v15 - (^(v15)))
        if (false || false) {
            switch v3 {
            case -2147483648:
                v17 := GenS1{F0: v2, F1: (v4 - (^(v4))), F2: 49}
                _ = v17
            }
        }
        _ = v16
    } else {
        fmt.Println(v2, v3)
        var v18 string = ("foo bar" + (("a" + ((("foo bar" + "foo bar") + "a") + (("hello" + "a") + ("\n\t" + "hello")))) + "\n\t"))
        _ = v18
    }
    fmt.Println(v4, v15)
    for fi19 := 0; fi19 < 10; fi19++ {
        v15 -= 42
        _ = fi19
    }
    var v20 bool = true
    v21 := GenS1{F0: (v2 * 42), F1: 485, F2: 0}
    switch 9223372036854775807 {
    case -463:
        v22 := v1[2]
        _ = v22
    case 9223372036854775807:
        var v23 int8 = 71
        _ = v23
    case -1:
        v24 := GenS1{F0: 1000, F1: 32767, F2: 196}
        _ = v24
    default:
        _ = len(v1)
    }
    fmt.Println(v4, v15, v3)
    v20 = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 213
    v2 := make(map[int]int)
    var v3 int8 = -128
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: 42}
    fmt.Println(v1, v1, v1)
    fmt.Println(v3, v1)
    v6 := []int{}
    if (!(0 >= 32767)) {
        if (!(true || false)) {
            v3 = (v3 << ((2) & 63))
            var v7 bool = (("" >= ("hello" + ("a" + "hello"))) && true)
            fmt.Println(v3, v3, v4)
            _ = v7
        } else {
            if false {
                switch ("a" + "hello") {
                case "\n\t":
                    switch 304 {
                    case 9223372036854775807:
                        var v8 int = 2147483647
                        _ = v8
                    case 905:
                        _ = len(v6)
                    }
                case "foo bar":
                    v9 := GenS2{F0: v1, F1: (v1 / (((v1 + (v1 >> ((0) & 63)))) | 1)), F2: 2}
                    _ = v9
                }
            } else {
                var v10 string = "hello"
                fmt.Println(v4, v3)
                v4 = v4
                _ = v10
            }
            v11 := GenS3{F0: v1}
            _ = v11
        }
    }
    v12 := v2[583]
    _ = len(v2)
    v12 += (v12 - 128)
    switch ("\n\t" + "") {
    case "a":
        var v13 int = v12
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 uint = 0
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        delete(v2, (v12 / ((fi16) | 1)))
        fmt.Println(v4, fi16)
        _ = fi16
    }
    switch 42 {
    case 93:
        var v17 uint16 = 32767
        _ = v17
    case -9223372036854775808:
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(fi18, v4)
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS2{F0: v1, F1: (^(v1)), F2: 97}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    default:
        var v21 bool = true
        _ = v21
    }
    if false {
        var v22 uint = 4294967295
        _ = v22
    } else {
        var v23 int64 = -40
        switch v12 {
        case -1:
            fmt.Println(v1)
        }
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

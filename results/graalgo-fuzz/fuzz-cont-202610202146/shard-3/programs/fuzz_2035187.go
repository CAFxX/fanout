package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 int16 = -1
    var v2 bool = true
    v3 := []string{"", "\n\t", "a"}
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 0, F1: -1}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        switch "" {
        case "":
            v3 = append(v3, "a")
        case "hello":
            var v7 uint64 = 2
            _ = v7
        default:
            var v8 uint16 = 479
            _ = v8
        }
        v9 := GenS2{F0: ("\n\t" + "")}
        if (!("" > "hello")) {
            v1 += v1
        }
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint64 = 359
            _ = fi11
            _ = v12
        }
        fmt.Println(v1, v1)
        if true {
            v2 = v2
            var v13 int8 = 1
            _ = v13
        } else {
            var v14 uint64 = 937
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS1{F0: 4294967295, F1: -67}
                var v17 int64 = -9223372036854775808
                v17 -= v17
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        _ = fi10
    }
    v1 += -181
    var v18 int8 = 54
    v18 += v18
    v1 = v1
    for fi19 := 0; fi19 < 4; fi19++ {
        if v2 {
            v2 = v2
        }
        _ = fi19
    }
    if false {
        var v20 string = "foo bar"
        v2 = v2
        for fi21 := 0; fi21 < 9; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                if (!(42 < 32767)) {
                    var v23 int16 = (v1 ^ v1)
                    _ = v23
                }
                _ = fi22
            }
            _ = fi21
        }
        _ = v20
    }
    var v24 string = "hello"
    v1 = v1
    delete(v4, -2147483648)
    var v25 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

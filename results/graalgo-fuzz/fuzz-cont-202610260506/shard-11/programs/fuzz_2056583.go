package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    v1 := []uint32{}
    v2 := []string{}
    v3 := []uint32{65535, 4294967295}
    var v4 string = ("\n\t" + "\n\t")
    var v5 int64 = 42
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        if len(v2) > 0 {
            v2[0] = v4
        }
        _ = fi9
    }
    if false {
        fmt.Println(v8, v5, v5)
        var v10 int64 = v5
        _ = len(v1)
        _ = v10
    } else {
        var v11 uint = 0
        fmt.Println(v8)
        _ = v11
    }
    var v12 string = ""
    var v13 uint = 65535
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            if len(v1) > 0 {
                v1[0] = (v8 << (((v13 + v13)) & 63))
            }
            var v16 float64 = 0.1
            var v17 uint16 = 42
            v8 = (v8 & v8)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        switch ("hello" + (("hello" + "") + "hello")) {
        case "hello":
            var v18 int8 = 1
            _ = v18
        case "":
            v19 := GenS1{F0: 1, F1: false}
            _ = v19
        }
        _ = fi14
    }
    var v20 uint32 = v8
    v21 := GenS1{F0: 2, F1: (((^(v13)) == (v13 * (v13 * 551))) || false)}
    var v22 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

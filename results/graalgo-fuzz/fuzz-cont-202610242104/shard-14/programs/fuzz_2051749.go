package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint64 = 0
    var v2 float32 = 1.5
    v3 := make(map[string]uint32)
    var v4 int8 = 100
    fmt.Println(v1, v2, v2)
    v5 := []int{}
    for fi6 := 0; fi6 < 4; fi6++ {
        v4 = (^(v4))
        switch fi6 {
        case -538:
            v7 := GenS1{F0: 934, F1: false, F2: 49}
            _ = v7
        case 32767:
            v4 = (^(v4))
        case 42:
            v2 = ((1e10 * v2) / (1.5 * v2))
        }
        _ = fi6
    }
    switch -2147483648 {
    case 32767:
        fmt.Println(v4, v4)
    case 42:
        if false {
            delete(v3, ("hello" + (("a" + ((("" + "") + "foo bar") + "foo bar")) + "foo bar")))
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS2{F0: -123, F1: 2, F2: false}
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 int8 = (v4 * 0)
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 255, F1: true, F2: 42}
                _ = fi8
                _ = v12
                _ = v9
            }
            if true {
                var v13 uint = 22
                v14 := v3[(((("\n\t" + "hello") + "\n\t") + ("foo bar" + "")) + "\n\t")]
                _ = v13
                _ = v14
            } else {
                v1 += v1
                v4 += v4
            }
        } else {
            v4 -= (v4 << ((1) & 63))
        }
    }
    v5 = append(v5, 4294967295)
    if (false && false) {
        var v15 string = ("foo bar" + ("foo bar" + ((("\n\t" + "a") + "\n\t") + "\n\t")))
        var v16 uint32 = 42
        if (0 > 42) {
            v17 := GenS1{F0: 0, F1: (false && false), F2: v4}
            _ = v17
        }
        _ = v15
        _ = v16
    }
    if (1 < 1) {
        delete(v3, (("a" + ("" + ("\n\t" + "\n\t"))) + (("hello" + "foo bar") + "hello")))
    }
    var v18 uint32 = 2147483648
    v19 := GenS3{F0: 615, F1: v1, F2: 912}
    switch 32768 {
    case 445:
        v20 := GenS4{F0: ""}
        _ = v20
    default:
        var v21 int8 = -107
        _ = v21
    }
    if (1 < 9223372036854775807) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v22 uint32 = 1
    var v23 int = 128
    var v24 uint = 42
    switch ("" + ("hello" + "hello")) {
    case "":
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v24)
            var v26 int32 = 2147483647
            _ = fi25
            _ = v26
        }
    case "a":
        fmt.Println(v23)
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

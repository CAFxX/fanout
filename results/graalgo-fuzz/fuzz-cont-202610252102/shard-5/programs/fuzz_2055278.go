package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0() int {
    var v4 uint64 = 1
    v4 = v4
    v4 = v4
    _ = v4
    return 128
}

func main() {
    v1 := []int{}
    var v2 string = "\n\t"
    v3 := make(map[uint32]string)
    if true {
        if (-353 != 42) {
            switch 219 {
            case 9223372036854775807:
                v5 := GenS1{F0: -1, F1: -1}
                _ = v5
            default:
                v2 = v2
            }
        } else {
            var v6 int16 = -1
            if (0 >= 78) {
                fmt.Println(v2)
                var v7 int8 = 127
                v6 -= 0
                _ = v7
            }
            var v8 uint32 = 0
            _ = v6
            _ = v8
        }
        var v9 uint32 = 4294967295
        switch v2 {
        case "a":
            switch -1 {
            case -394:
                var v10 float32 = -2.25
                _ = v10
            }
        case "foo bar":
            v2 = v2
        default:
            v11 := GenS1{F0: -269, F1: 0}
            _ = v11
        }
        _ = v9
    } else {
        v2 = v2
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v2 = (("" + "\n\t") + "hello")
            v2 = v2
            _ = fi13
        }
        v14 := GenS2{F0: 255, F1: 1}
        var v15 uint32 = 65535
        _ = fi12
        _ = v14
        _ = v15
    }
    v2 = "hello"
    v16 := v3[2147483647]
    var v17 int = -326
    v16 = ("foo bar" + (v2 + ((v16 + v2) + v2)))
    var v18 uint64 = 42
    if true {
        v19 := GenS3{F0: 57}
        if false {
            for fi20 := 0; fi20 < 8; fi20++ {
                if false {
                    v21 := GenS2{F0: 32767, F1: 1}
                    _ = v21
                } else {
                    v22 := GenS4{F0: v18, F1: true}
                    switch v2 {
                    case "a":
                        v16 = (v16 + (v2 + ((v2 + "hello") + "a")))
                    case "hello":
                        var v23 bool = false
                        _ = v23
                    }
                    _ = v22
                }
                _ = fi20
            }
        }
        delete(v3, 42)
        _ = v19
    }
    v24 := v3[2147483647]
    fmt.Println(v16)
    fmt.Println(v17)
    var v25 float32 = 0.0
    var v26 float32 = float32(-2.25)
    switch "" {
    case "":
        var v27 uint64 = (v18 + 42)
        _ = v27
    case "\n\t":
        var v28 int64 = 893
        _ = v28
    case "foo bar":
        var v29 bool = false
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}

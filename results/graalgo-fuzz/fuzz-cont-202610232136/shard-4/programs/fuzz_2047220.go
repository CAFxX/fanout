package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]bool)
    var v3 int = -1
    switch "a" {
    case "":
        v3 = (v3 | (v3 + v3))
    case "foo bar":
        if (1e10 != 0.0) {
            if (!true) {
                switch "\n\t" {
                case "\n\t":
                    var v4 string = (("foo bar" + "\n\t") + "\n\t")
                    _ = v4
                case "":
                    if true {
                        _ = len(v2)
                        v1 = v1
                    }
                case "foo bar":
                    delete(v2, "hello")
                }
                v3 += (v3 * v3)
                var v5 int64 = -1
                _ = v5
            }
            v3 -= (-(v3))
            v1 = (v1 & v1)
        } else {
            var v6 uint8 = 128
            v2[("\n\t" + "")] = (0 != 42)
            v1 = v1
            _ = v6
        }
    case "hello":
        v7 := []int{v3, v3}
        _ = v7
    default:
        if (!(true && (56 < 127))) {
            v8 := []int{v3, -289, v3}
            for fi9 := 0; fi9 < 13; fi9++ {
                for fi10 := 0; fi10 < 2; fi10++ {
                    fmt.Println(fi9, fi9, v1)
                    _ = fi10
                }
                _ = fi9
            }
            _ = v8
        }
    }
    _ = len(v2)
    v3 = v3
    switch "hello" {
    case "hello":
        var v11 int16 = -1
        _ = v11
    case "":
        fmt.Println(v3)
    case "\n\t":
        var v12 uint64 = 42
        _ = v12
    default:
        for fi13 := 0; fi13 < 13; fi13++ {
            v3 = v3
            fmt.Println(fi13)
            _ = fi13
        }
    }
    v14 := []int{v3}
    var v15 uint64 = 354
    v15 = (^(v15))
    fmt.Println(v3, v3)
    var v16 uint8 = uint8(v15)
    if (false || (false || (((-2.25 * 1.5) + ((0.1 - (-2.25 / -2.25)) / 1.0)) <= 1.5))) {
        var v17 int32 = 79
        fmt.Println(v1)
        _ = v17
    }
    v18 := GenS1{F0: false}
    fmt.Println(v1)
    v3 -= 1
    v19 := v2[""]
    if (!(65535 <= 42)) {
        v20 := GenS2{F0: int32(v3)}
        v19 = false
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 float32 = -2.25
            switch "a" {
            case "foo bar":
                v16 += v16
            case "a":
                var v23 string = "foo bar"
                _ = v23
            case "":
                fmt.Println(v3, v15)
            default:
                v24 := GenS1{F0: false}
                _ = v24
            }
            var v25 float32 = (float32(0.1) + v22)
            _ = fi21
            _ = v22
            _ = v25
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}

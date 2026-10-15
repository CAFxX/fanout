package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        if (-2147483648 >= 1) {
            v7 := GenS1{F0: 2147483647, F1: ""}
            p0 += 0
            _ = v7
        }
        _ = fi6
    }
    if true {
        if (255 < 255) {
            switch "foo bar" {
            case "foo bar":
                v8 := GenS1{F0: -9223372036854775808, F1: ("foo bar" + "a")}
                _ = v8
            case "\n\t":
                var v9 string = (("\n\t" + "\n\t") + "foo bar")
                _ = v9
            case "a":
                if (901 < 944) {
                    fmt.Println(p0)
                }
            }
            p0 = 2
        }
        p0 = 4294967295
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: fi10, F1: ""}
            _ = fi10
            _ = v11
        }
    }
    if false {
        p0 = p0
        p0 -= (^(p0))
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS2{F0: 127, F1: 2}
            var v14 int32 = -1
            if false {
                v15 := []int{-9223372036854775808, (fi12 - fi12), (fi12 & fi12)}
                _ = v15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 42
    var v3 int8 = 1
    var v4 bool = true
    v5 := []string{(((("hello" + ("foo bar" + "\n\t")) + "") + "foo bar") + ""), "\n\t"}
    var v16 int = 9223372036854775807
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS2{F0: 128, F1: 65535}
        _ = fi17
        _ = v18
    }
    var v19 uint32 = (v2 << ((2) & 63))
    var v20 uint8 = 136
    v19 += (v2 >> ((4294967295) & 63))
    v21 := GenS3{F0: 42, F1: 2, F2: (v20 % (((v20 * v20)) | 1))}
    v22 := GenS4{F0: false}
    v4 = ((v20 >> ((65535) & 63)) <= v20)
    switch "hello" {
    case "a":
        fmt.Println(v20, v19)
    case "hello":
        v23 := v1[v19]
        _ = v23
    default:
        v2 += v19
    }
    var v24 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

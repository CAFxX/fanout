package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 bool }
func genFn0(p0 int16) uint8 {
    v7 := GenS1{F0: 1, F1: 32767}
    if (0.0 > 0.1) {
        var v8 uint = 65535
        switch "hello" {
        case "foo bar":
            v9 := GenS2{F0: 27}
            _ = v9
        case "a":
            p0 -= p0
        }
        _ = v8
    } else {
        switch "\n\t" {
        case "\n\t":
            p0 += 42
        case "a":
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := []int{}
                _ = fi10
                _ = v11
            }
        default:
            p0 = 744
        }
    }
    var v12 uint32 = 0
    _ = p0
    _ = v12
    _ = v7
    return 255
}
func genFn1(p0 bool) bool {
    v13 := []int{127, 0, 127, 32767}
    p0 = p0
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v14 := GenS3{F0: 1}
    _ = p0
    _ = v13
    _ = v14
    return (1 >= 0)
}

func main() {
    var v1 uint8 = 2
    var v2 int16 = 1
    v3 := make(map[string]uint32)
    v4 := []string{("" + ((("hello" + "foo bar") + "hello") + "a")), ("a" + ""), "\n\t", "\n\t"}
    var v5 uint64 = 2
    var v6 string = "foo bar"
    fmt.Println(v1, v6)
    v15 := GenS4{F0: 4294967295, F1: 816, F2: (((v2 | v2) >= 0) || true)}
    v2 += (v2 / ((v2) | 1))
    fmt.Println(v6, v1)
    v16 := GenS1{F0: genFn0((v2 << ((65535) & 63))), F1: 298}
    v6 = "\n\t"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint32 = 42
    var v18 uint32 = 0
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS4{F0: v18, F1: 2, F2: false}
        _ = fi19
        _ = v20
    }
    v4 = append(v4, (v6 + "hello"))
    fmt.Println(v6, v6, v1)
    _ = len(v4)
    fmt.Println(v1, v18)
    var v21 uint64 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

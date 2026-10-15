package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint8 = 56
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 uint16 = 32767
    v2[2147483647] = (("foo bar" + "hello") + ("\n\t" + (("foo bar" + "foo bar") + "")))
    var v6 uint32 = 1
    v1 += uint8(v6)
    if false {
        switch -1 {
        case -984:
            var v7 uint = 42
            _ = v7
        case 1:
            if false {
                var v8 int8 = int8(v6)
                v9 := []int{-1, 42, 2147483647, -2147483648}
                if true {
                    v1 = 1
                }
                _ = v8
                _ = v9
            }
        }
    } else {
        if true {
            var v10 float32 = (3.14159 + 1e10)
            _ = v10
        }
    }
    fmt.Println(v5, v5)
    fmt.Println(v6, v5, v1)
    delete(v3, 2147483648)
    if (!true) {
        v2[v6] = "foo bar"
        v2[(v6 | (v6 ^ (^(v6))))] = "\n\t"
    }
    v11 := GenS1{F0: "\n\t"}
    v1 -= v1
    var v12 int32 = 202
    v13 := GenS1{F0: ("hello" + ("hello" + ""))}
    switch int(v5) {
    case 128:
        v14 := []int{2147483647, int(v6)}
        _ = v14
    case 127:
        var v15 int32 = (v12 ^ v12)
        _ = v15
    case -803:
        v6 = v6
    default:
        v16 := GenS1{F0: "\n\t"}
        _ = v16
    }
    v17 := GenS2{F0: (v6 >> ((4294967295) & 63))}
    v18 := GenS3{F0: int8(v5), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

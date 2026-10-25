package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint8 = 0
    var v2 int32 = -2147483648
    var v3 int = 9223372036854775807
    var v4 bool = false
    var v5 int16 = 1000
    fmt.Println(v5)
    var v6 string = ""
    fmt.Println(v5)
    fmt.Println(v4)
    v4 = v4
    v7 := []int{v3}
    var v8 int32 = (v2 / ((v2) | 1))
    v9 := GenS1{F0: v3}
    v10 := GenS1{F0: (v3 * 128)}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 = v2
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                if len(v7) > 0 {
                    v7[0] = fi12
                }
            }
            if v4 {
                var v13 string = "hello"
                var v14 int32 = v8
                var v15 uint16 = 65535
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        switch "\n\t" {
        case "hello":
            if v4 {
                v7 = append(v7, v3)
                fmt.Println(v3, v2, v3)
            }
        case "":
            var v16 uint8 = (^(v1))
            _ = v16
        case "\n\t":
            var v17 uint64 = 822
            _ = v17
        default:
            var v18 string = (v6 + (v6 + v6))
            _ = v18
        }
        _ = fi11
    }
    fmt.Println(v4, v4, v3)
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(v3)
        v5 = 150
        _ = fi19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        if ((-(v3)) < (-(fi20))) {
            if (v5 <= v5) {
                if len(v7) > 0 {
                    v7[0] = v3
                }
                v5 += v5
                if true {
                    _ = len(v7)
                }
            }
            var v21 int64 = -1
            _ = v21
        }
        switch "" {
        case "":
            fmt.Println(v2)
        case "foo bar":
            v22 := GenS1{F0: (fi20 % ((fi20) | 1))}
            _ = v22
        default:
            v23 := GenS2{F0: 1, F1: v5}
            _ = v23
        }
        v1 = (^(v1))
        _ = fi20
    }
    v24 := GenS3{F0: 238}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

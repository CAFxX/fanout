package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int32, p1 int8, p2 uint16) bool {
    v6 := GenS1{F0: 0}
    v7 := GenS2{F0: 4294967295}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint = 255
    var v2 int = 32767
    v3 := []int{v2, (v2 & 127), v2, 32768}
    var v4 uint = (v1 + 42)
    var v5 int8 = 0
    var v9 bool = (999 != 32767)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v9)
    v3 = append(v3, 32767)
    v2 += (^(v2))
    v10 := GenS1{F0: 2}
    v11 := GenS3{F0: 2147483647, F1: v2}
    for fi12 := 0; fi12 < 12; fi12++ {
        v2 -= fi12
        _ = fi12
    }
    var v13 bool = true
    var v14 uint = 4294967295
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint = 279
        if len(v3) > 0 {
            v3[0] = fi15
        }
        v3 = append(v3, (v2 << (((v14 & 2)) & 63)))
        _ = fi15
        _ = v16
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v14, v4, v4)
    case "foo bar":
        v5 += (v5 ^ (v5 ^ (-(v5))))
    case "a":
        v4 += (v4 + (v14 ^ 438))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

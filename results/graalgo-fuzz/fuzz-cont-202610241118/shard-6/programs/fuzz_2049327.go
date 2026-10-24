package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int8; F2 string }
func genFn0(p0 uint64, p1 string, p2 int32) int {
    v6 := []int{}
    v7 := GenS1{F0: 38, F1: p1}
    if len(v6) > 0 {
        v6[0] = -1
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(p0, fi8, p1)
        _ = fi8
    }
    v9 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return int(p2)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 4294967295
    var v3 int32 = 1
    var v4 uint32 = 2147483647
    var v5 int32 = (v3 % (((-(v3))) | 1))
    for fi10 := 0; fi10 < 19; fi10++ {
        switch "a" {
        case "hello":
            v3 += v5
        }
        v5 = (^(v3))
        v11 := GenS1{F0: 9223372036854775807, F1: "a"}
        _ = fi10
        _ = v11
    }
    fmt.Println(v4, v4, v1)
    v12 := GenS3{F0: 196, F1: 0, F2: ("" + ("hello" + ((("a" + "") + (("foo bar" + "a") + "a")) + (("hello" + "") + ("a" + "foo bar")))))}
    switch "" {
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := []int{int(v2), -246, 32768}
                var v16 string = "hello"
                v1 = 2147483647
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    case "hello":
        v3 += v5
    }
    fmt.Println(v1, v3, v1)
    var v17 uint16 = 255
    var v18 float32 = float32(-2.25)
    v19 := GenS2{F0: (v17 + 65535)}
    v20 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 float32 = 3.14159
    v4 := make(map[int]bool)
    var v5 uint = 200
    for fi6 := 0; fi6 < 10; fi6++ {
        v5 = (v5 ^ v5)
        var v7 uint8 = 108
        _ = fi6
        _ = v7
    }
    v1 = v1
    var v8 int = 128
    var v9 float32 = v3
    v5 -= 415
    if (!false) {
        var v10 bool = (2 == uint32(v5))
        v11 := GenS1{F0: 9223372036854775807, F1: ((v1 + v1) + (("" + ((((v1 + v1) + v1) + (v1 + v1)) + (v1 + v1))) + (v1 + (((v1 + (v1 + "foo bar")) + "") + "foo bar")))), F2: 791}
        _ = v10
        _ = v11
    }
    v3 = v3
    if false {
        v12 := []int{}
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 string = v1
            v15 := GenS1{F0: 42, F1: v14, F2: 0}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    v16 := []int{v8, 2147483647, (^(v8))}
    var v17 string = (v1 + "hello")
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := GenS1{F0: 0, F1: "hello", F2: 18446744073709551615}
        var v20 int32 = -220
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

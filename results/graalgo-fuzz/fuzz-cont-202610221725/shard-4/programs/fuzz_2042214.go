package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int8 }
func genFn0(p0 uint32) uint8 {
    var v6 string = ("" + ("\n\t" + ""))
    fmt.Println(v6)
    v6 = v6
    var v7 uint64 = 18446744073709551615
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1() uint64 {
    v8 := []int{-2147483648}
    v9 := GenS1{F0: 1}
    var v10 uint64 = 1
    _ = v10
    _ = v8
    _ = v9
    return v10
}

func main() {
    v1 := []string{"", ("a" + ("" + "")), (("hello" + "hello") + "a")}
    var v2 uint32 = 155
    var v3 int32 = 42
    var v4 uint32 = v2
    var v5 bool = (!false)
    v11 := GenS2{F0: (^(v4)), F1: 42}
    fmt.Println(v3, v5)
    var v12 int8 = 42
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v5 = (v5 && v5)
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v4, v2, v12)
        fmt.Println(v12, v2)
        if ((v3 % ((0) | 1)) > v3) {
            _ = len(v1)
            var v14 int64 = -552
            v15 := GenS3{F0: -620, F1: (v14 << ((42) & 63)), F2: v14}
            _ = v14
            _ = v15
        } else {
            var v16 int32 = v3
            _ = v16
        }
        _ = fi13
    }
    v17 := GenS1{F0: 1000}
    fmt.Println(v5, v4)
    fmt.Println(v12, v4, v2)
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS4{F0: 0, F1: (v4 % (((v4 - (v2 - (^(v2))))) | 1)), F2: 100}
        v5 = (int16(v3) > 0)
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

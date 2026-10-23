package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16) bool {
    p0 -= (p0 & p0)
    var v6 int16 = -32768
    _ = p0
    _ = v6
    return (!((v6 == (v6 * 32767)) || (!(!false))))
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{65535}
    var v3 uint64 = v1
    v4 := []int{9223372036854775807, 32767}
    var v5 bool = (((("\n\t" + ("\n\t" + "a")) + "hello") + ("" + "\n\t")) == "foo bar")
    v1 += 1
    var v7 uint32 = 1
    var v8 uint = 18446744073709551615
    v9 := GenS1{F0: 2147483647, F1: 74}
    v10 := GenS1{F0: 2147483647, F1: 39}
    v11 := GenS2{F0: 0}
    v7 = uint32(v1)
    _ = len(v4)
    var v12 float32 = (0.1 - 1.5)
    fmt.Println(v8, v1, v5)
    var v13 bool = false
    if false {
        v14 := GenS1{F0: 4294967295, F1: -1}
        var v15 bool = true
        var v16 int16 = 271
        _ = v14
        _ = v15
        _ = v16
    }
    if (!false) {
        if v13 {
            v3 = v1
            _ = len(v4)
            fmt.Println(v3, v13)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v17 float32 = ((((v12 / float32((1e10 - 3.14159))) * float32(0.1)) - 3.14159) / v12)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v17
        }
        v13 = (47 > uint8(v7))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}

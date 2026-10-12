package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int16; F1 uint; F2 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 int8 = 42
    var v4 uint32 = 0
    v5 := []int{2147483647, -847, 32767}
    v6 := GenS1{F0: -32768, F1: 223}
    var v7 int16 = -1
    var v8 int32 = 0
    v7 -= 1
    var v9 uint8 = 255
    if true {
        fmt.Println(v7, v7)
        v9 = (^(v9))
    } else {
        var v10 string = ("foo bar" + "foo bar")
        if ((65535 != v4) && true) {
            fmt.Println(v4)
        }
        _ = v10
    }
    v9 -= 255
    v4 = (v4 ^ v4)
    for fi11 := 0; fi11 < 13; fi11++ {
        if false {
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
        }
        v12 := v2[""]
        _ = fi11
        _ = v12
    }
    v13 := v2["foo bar"]
    v7 = -1
    fmt.Println(v4, v3, v4)
    v14 := GenS2{F0: v7, F1: 0, F2: v7}
    v15 := GenS3{F0: (-(v7)), F1: 1, F2: 2}
    delete(v1, "foo bar")
    if false {
        v7 += (v7 ^ -32768)
        var v16 bool = ((42 >= v9) && false)
        var v17 int32 = v8
        _ = v16
        _ = v17
    } else {
        var v18 int64 = 1
        _ = v18
    }
    v19 := GenS4{F0: 353}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

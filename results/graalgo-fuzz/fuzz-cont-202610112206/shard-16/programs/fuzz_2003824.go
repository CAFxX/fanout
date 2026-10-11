package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int; F1 int; F2 uint8 }
func genFn0() uint64 {
    v5 := []int{32768}
    var v6 float64 = 3.14159
    v6 = -2.25
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 305, F1: 0, F2: -247}
        _ = fi7
        _ = v8
    }
    if (2147483647 > 42) {
        var v9 uint32 = 0
        fmt.Println(v9)
        _ = v9
    } else {
        var v10 int8 = 100
        _ = v10
    }
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []string{("" + ("\n\t" + "a")), ("foo bar" + ("\n\t" + "foo bar")), (("" + ("hello" + "hello")) + "foo bar")}
    var v2 int8 = 1
    var v3 uint = 65535
    var v4 int8 = (v2 * v2)
    var v11 int32 = -2147483648
    v12 := GenS2{F0: ""}
    v4 = v4
    v4 = 29
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int64 = -9223372036854775808
        _ = fi13
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v4
    fmt.Println(v2)
    fmt.Println(v3, v4, v2)
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS3{F0: 42, F1: 2147483647, F2: 0}
        v17 := GenS3{F0: v11, F1: 245, F2: 1}
        v18 := GenS4{F0: -484, F1: fi15, F2: 39}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 float64 = 1e10
        fmt.Println(v3, v11)
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}

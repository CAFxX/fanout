package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 uint8 }
func genFn0() int64 {
    var v6 int16 = -737
    var v7 uint64 = 1
    v6 += (v6 << ((316) & 63))
    _ = v6
    _ = v7
    return 419
}
func genFn1(p0 int64, p1 uint64) string {
    fmt.Println(p1, p1)
    if (2 < 526) {
        var v8 int = -24
        fmt.Println(p1)
        _ = v8
    } else {
        var v9 int8 = -1
        v9 = int8(p1)
        v10 := GenS1{F0: 104}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    v1 := []int{-371, -9223372036854775808, 32768, -2147483648}
    v2 := []int{4294967295, 326, -59}
    var v3 uint8 = 87
    var v4 int16 = 0
    var v5 uint32 = 2
    var v11 uint8 = (v3 - 128)
    switch ("foo bar" + "\n\t") {
    case "a":
        _ = len(v1)
    case "foo bar":
        v12 := GenS1{F0: v3}
        _ = v12
    }
    var v13 int64 = -495
    var v14 uint64 = 42
    if false {
        v5 -= v5
    }
    v15 := GenS2{F0: 579, F1: 18446744073709551615, F2: ("a" + ("\n\t" + ("hello" + "foo bar")))}
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS2{F0: 42, F1: 42, F2: "\n\t"}
        _ = fi16
        _ = v17
    }
    var v18 float32 = float32(float64(1e10))
    v19 := GenS3{F0: (v3 | (v3 * (v11 ^ 152)))}
    fmt.Println(v11)
    fmt.Println(v3)
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS3{F0: v3}
        _ = fi20
        _ = v21
    }
    fmt.Println(v14, v4, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int) int32 {
    fmt.Println(p0)
    p0 += p0
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 float32 = 1.5
        _ = fi6
        _ = v7
    }
    if false {
        p0 = (p0 * p0)
        v8 := []int{}
        _ = v8
    }
    _ = p0
    return 1
}
func genFn1(p0 int16, p1 uint) int {
    switch "foo bar" {
    case "a":
        fmt.Println(p0, p1)
    case "hello":
        var v9 int16 = p0
        _ = v9
    }
    var v10 uint32 = 0
    for fi11 := 0; fi11 < 12; fi11++ {
        p0 += p0
        p0 = (p0 % (((p0 * (p0 << ((435) & 63)))) | 1))
        p0 = p0
        _ = fi11
    }
    if false {
        v12 := GenS1{F0: 18446744073709551615, F1: p1, F2: ("a" + "")}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 32768
}

func main() {
    var v1 bool = (!false)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 int16 = -1
    var v5 float32 = 1.5
    delete(v3, 9223372036854775807)
    v4 = (v4 * v4)
    var v13 int = 1
    v14 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: "hello"}
    var v15 int = 0
    delete(v2, (v15 - -2147483648))
    if (!(4294967295 > 2147483647)) {
        v5 -= -2.25
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v5, v1, v1)
        fmt.Println(v5, v4)
        v17 := GenS2{F0: 2}
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 172
    if false {
        v19 := GenS3{F0: 4294967295}
        _ = v19
    } else {
        v20 := GenS1{F0: 690, F1: 4294967295, F2: ""}
        var v21 int64 = 9223372036854775807
        fmt.Println(v4, v13)
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0(p0 int8) uint16 {
    v7 := GenS1{F0: 2147483648}
    p0 += (p0 ^ 1)
    _ = p0
    _ = v7
    return 0
}
func genFn1() int8 {
    var v8 int32 = -1
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint64 = 2
        v11 := GenS2{F0: "foo bar"}
        fmt.Println(fi9, v10, fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: ("a" + "foo bar")}
    fmt.Println(v8)
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint64 = 0
        _ = fi13
        _ = v14
    }
    _ = v12
    _ = v8
    return 27
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]int)
    var v3 uint16 = v1
    var v4 uint16 = (^(v1))
    var v5 uint = 2
    v6 := make(map[uint32]int)
    var v15 int = 9223372036854775807
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int64 = 0
        fmt.Println(v1, v17)
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 425}
    if true {
        var v19 float64 = ((-2.25 / 3.14159) - 0.1)
        v20 := GenS1{F0: 0}
        _ = v19
        _ = v20
    }
    if false {
        v15 = int(v5)
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := []int{(fi21 ^ (v15 << ((v5) & 63))), (v15 ^ 2147483647), fi21}
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 string = (("foo bar" + "") + ("foo bar" + "\n\t"))
                v25 := GenS3{F0: v5}
                fmt.Println(v4, fi23)
                var v26 bool = true
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi21
            _ = v22
        }
    } else {
        if true {
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(v3)
                fmt.Println(v3)
                _ = fi27
            }
        }
        v28 := GenS2{F0: ("foo bar" + "foo bar")}
        v29 := GenS1{F0: 42}
        _ = v28
        _ = v29
    }
    _ = len(v6)
    delete(v6, 42)
    var v30 int16 = -32768
    var v31 int = (v15 * 1)
    v15 = v15
    for fi32 := 0; fi32 < 14; fi32++ {
        _ = len(v2)
        _ = fi32
    }
    v33 := GenS4{F0: 108, F1: (v4 / ((v3) | 1)), F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

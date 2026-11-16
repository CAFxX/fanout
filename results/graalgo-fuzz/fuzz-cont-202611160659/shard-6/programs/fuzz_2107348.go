package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 42
    v2 := make(map[int]string)
    var v3 int64 = 847
    v4 := make(map[uint32]string)
    var v5 bool = ("a" > ("a" + "hello"))
    if v5 {
        v3 -= v3
        if true {
            v6 := GenS1{F0: -128, F1: -899}
            v7 := GenS2{F0: -114}
            _ = v6
            _ = v7
        }
        var v8 int64 = v3
        _ = v8
    }
    v9 := []int{(^(v1)), (v1 * -1), v1}
    if true {
        var v10 bool = v5
        for fi11 := 0; fi11 < 6; fi11++ {
            v3 += v3
            var v12 int16 = -32768
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        var v13 int = (v1 ^ 42)
        _ = v13
    }
    v14 := GenS3{F0: -1}
    fmt.Println(v5, v5)
    var v15 bool = v5
    fmt.Println(v5, v5, v15)
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 string = ("a" + "foo bar")
        _ = fi16
        _ = v17
    }
    switch 983 {
    case 4294967295:
        v5 = false
    case 0:
        var v18 float32 = (float32((0.1 + 0.0)) * 1e10)
        _ = v18
    case 2147483647:
        for fi19 := 0; fi19 < 1; fi19++ {
            if false {
                v3 -= -60
                v20 := GenS3{F0: -12}
                _ = len(v4)
                _ = v20
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                v4[13] = ((("a" + (("" + "hello") + ("\n\t" + "\n\t"))) + "") + "a")
                fmt.Println(fi21, v5, v1)
                _ = fi21
            }
            _ = fi19
        }
    default:
        v15 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0() int {
    var v7 string = "a"
    if true {
        if true {
            v8 := GenS1{F0: 2}
            _ = v8
        } else {
            v7 = ""
        }
        v7 = "foo bar"
        fmt.Println(v7)
    } else {
        v9 := []int{}
        _ = v9
    }
    if false {
        fmt.Println(v7)
        v10 := []int{127, 32767}
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(v7)
            switch "\n\t" {
            case "foo bar":
                var v12 int64 = -9223372036854775808
                _ = v12
            }
            _ = fi11
        }
        _ = v10
    } else {
        var v13 string = v7
        _ = v13
    }
    v14 := GenS2{F0: 0, F1: false, F2: 9223372036854775807}
    var v15 bool = true
    _ = v14
    _ = v15
    _ = v7
    return 127
}

func main() {
    var v1 int8 = 116
    v2 := []uint32{2147483648, 0, 53}
    var v3 uint32 = 834
    var v4 float32 = -2.25
    var v5 uint8 = 127
    var v6 int64 = -9223372036854775808
    fmt.Println(v1)
    var v16 int8 = v1
    v17 := GenS2{F0: v3, F1: false, F2: (-(v6))}
    for fi18 := 0; fi18 < 14; fi18++ {
        fmt.Println(v4, v4, v6)
        if true {
            fmt.Println(v6, v16, v6)
            v16 -= v16
        }
        _ = fi18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v2 = append(v2, 1)
        _ = fi19
    }
    var v20 int8 = v1
    v21 := GenS3{F0: -19, F1: v3}
    var v22 uint64 = 976
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

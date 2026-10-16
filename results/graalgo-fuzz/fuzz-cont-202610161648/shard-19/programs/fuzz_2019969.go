package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint) int {
    var v6 string = "foo bar"
    v7 := GenS1{F0: p1}
    v6 = v6
    if true {
        switch v6 {
        case "foo bar":
            v8 := []int{1, 1}
            _ = v8
        }
        var v9 string = v6
        v10 := GenS2{F0: true, F1: 32767, F2: int16(p0)}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -53
}

func main() {
    v1 := []int{32767}
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v4 uint = 449
    v5 := []uint32{2, 0, 469, 765}
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    switch -450 {
    case 42:
        var v12 bool = false
        _ = v12
    case -1:
        fmt.Println(v2, v2)
    case 1:
        var v13 uint8 = 21
        _ = v13
    }
    v14 := v3[v2]
    v2 = "a"
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v2, v14)
        v4 = v4
        _ = fi15
    }
    v2 = (("foo bar" + ("a" + (v2 + ""))) + v2)
    v4 = 2
    v16 := v3[v2]
    v17 := GenS3{F0: -1}
    if false {
        v2 = "hello"
        if false {
            v4 += v4
            var v18 uint16 = 442
            _ = v18
        }
    }
    v19 := GenS3{F0: 0}
    var v20 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

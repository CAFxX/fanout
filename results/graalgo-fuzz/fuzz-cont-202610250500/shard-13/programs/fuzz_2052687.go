package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool, p1 bool) uint16 {
    switch -217 {
    case -2147483648:
        switch -521 {
        case -855:
            p1 = p0
        default:
            var v4 uint64 = 0
            _ = v4
        }
    default:
        if p1 {
            var v5 uint32 = 2
            var v6 int8 = 90
            p0 = p1
            _ = v5
            _ = v6
        } else {
            v7 := GenS1{F0: 1}
            p1 = p0
            _ = v7
        }
    }
    p1 = p1
    p1 = p1
    _ = p0
    _ = p1
    return 65
}
func genFn1() uint16 {
    var v8 int16 = -560
    var v9 string = ("a" + ("hello" + ("\n\t" + "foo bar")))
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 128
    var v3 uint64 = 735
    var v10 int64 = 42
    for fi11 := 0; fi11 < 13; fi11++ {
        if (263 == uint32(v2)) {
            v3 -= v3
        }
        var v12 uint32 = 861
        var v13 float32 = 0.1
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int64 = v10
    v10 -= (v10 | v14)
    if true {
        var v15 int16 = 1
        v16 := GenS1{F0: v2}
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: (^(v2))}
        _ = v17
    }
    v18 := []int{639}
    v10 = (v10 % ((251) | 1))
    if len(v18) > 0 {
        v18[0] = 42
    }
    v18 = append(v18, -798)
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            if true {
                var v21 int32 = 42
                _ = v21
            }
            var v22 int16 = -34
            v10 += v14
            _ = fi20
            _ = v22
        }
        _ = fi19
    }
    if false {
        v10 += -1
    } else {
        v14 += v10
    }
    switch (("a" + "\n\t") + "hello") {
    case "hello":
        v23 := GenS2{F0: 1}
        _ = v23
    case "":
        fmt.Println(v2, v1)
    }
    v2 += (v2 - (v2 | v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}

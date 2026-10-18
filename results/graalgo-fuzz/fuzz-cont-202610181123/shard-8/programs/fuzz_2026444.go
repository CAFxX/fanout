package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 string }
func genFn0(p0 bool, p1 int) uint32 {
    var v4 int8 = 42
    if true {
        fmt.Println(v4, v4, p1)
        p0 = p0
        v5 := []int{(p1 - (p1 % ((p1) | 1))), p1, -279}
        _ = v5
    } else {
        p0 = (1 != 298)
        p0 = true
    }
    if p0 {
        var v6 int16 = 0
        p0 = p0
        _ = v6
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 int32, p1 int64, p2 uint64) uint32 {
    v7 := GenS1{F0: (p0 % (((p0 / ((int32(p1)) | 1))) | 1))}
    if (!(true || ("" >= "\n\t"))) {
        v8 := []int{-1}
        _ = v8
    }
    v9 := []int{1, 4294967295, -365, 9223372036854775807}
    v10 := GenS1{F0: (p0 << ((0) & 63))}
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v9
    return genFn0((true && false), -1)
}

func main() {
    v1 := []string{"a", ((("" + ("hello" + ("foo bar" + "\n\t"))) + ("foo bar" + "hello")) + "foo bar")}
    v2 := make(map[int]bool)
    v3 := []int{128}
    var v11 string = ("\n\t" + ((("a" + "") + "") + "\n\t"))
    switch v11 {
    case "foo bar":
        v11 = "a"
    case "hello":
        for fi12 := 0; fi12 < 9; fi12++ {
            _ = len(v3)
            v11 = v11
            v11 = ""
            _ = fi12
        }
    case "a":
        var v13 int = 0
        _ = v13
    }
    v2[1] = ((42 <= -2147483648) || (102 != 522))
    switch 2147483647 {
    case 42:
        v14 := GenS2{F0: 598}
        _ = v14
    }
    v15 := GenS3{F0: -1, F1: v11}
    if len(v1) > 0 {
        v1[0] = v11
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint = 1
        _ = fi16
        _ = v17
    }
    v11 = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}

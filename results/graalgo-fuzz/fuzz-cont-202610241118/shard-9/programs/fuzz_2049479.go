package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int, p1 int) bool {
    switch -420 {
    case 1:
        for fi5 := 0; fi5 < 3; fi5++ {
            fmt.Println(p0)
            _ = fi5
        }
    case -2147483648:
        fmt.Println(p0, p1)
    case 628:
        if false {
            p1 = p0
        } else {
            p0 = (p1 * 1)
            v6 := []int{p0, 4294967295, (p1 & (p0 & 4294967295)), (p1 - p1)}
            _ = v6
        }
    }
    v7 := GenS1{F0: (p1 % ((p1) | 1)), F1: (("\n\t" + "foo bar") + "\n\t"), F2: 1}
    _ = p0
    _ = p1
    _ = v7
    return (-128 < -97)
}

func main() {
    var v1 string = ("\n\t" + "\n\t")
    v2 := []int{-290, 4294967295, 508}
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v1 = (("hello" + "a") + (((((v1 + "foo bar") + ("foo bar" + (v1 + v1))) + v1) + v1) + "foo bar"))
    var v8 uint8 = 42
    var v9 uint = 18446744073709551615
    fmt.Println(v9, v9)
    var v10 int32 = -1
    v11 := GenS2{F0: 0}
    v8 += v8
    _ = len(v4)
    v12 := GenS3{F0: 144}
    if true {
        switch v1 {
        case "a":
            var v13 int64 = 948
            _ = v13
        default:
            v9 = v9
        }
    }
    _ = len(v4)
    var v14 uint64 = 577
    var v15 bool = false
    if v15 {
        fmt.Println(v1, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

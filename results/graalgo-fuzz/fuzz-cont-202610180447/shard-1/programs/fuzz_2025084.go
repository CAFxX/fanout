package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int16, p1 uint64, p2 string) int8 {
    v6 := []int{32767, -142, 128, 42}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 22
}
func genFn1(p0 int16) uint32 {
    v7 := GenS1{F0: -288}
    var v8 uint16 = 32767
    _ = p0
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 65535
    v3 := make(map[string]bool)
    v4 := []int{}
    v5 := make(map[int]bool)
    var v9 uint = v2
    var v10 string = "a"
    var v11 float64 = 3.14159
    v12 := GenS2{F0: -89}
    if true {
        if false {
            var v13 int32 = 1
            if false {
                _ = len(v4)
                v10 = "foo bar"
            } else {
                var v14 uint16 = 595
                _ = len(v4)
                _ = v14
            }
            switch ("" + v10) {
            case "foo bar":
                v13 = v13
            default:
                v11 += (v11 - (0.0 / v11))
            }
            _ = v13
        } else {
            v10 = "\n\t"
            v15 := GenS3{F0: "hello", F1: 65535, F2: 0}
            _ = v15
        }
        _ = len(v4)
    }
    v5[-670] = ("foo bar" != "foo bar")
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint64 = 42
        if true {
            _ = len(v1)
        }
        var v18 uint16 = 854
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v10, v10, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint, p1 uint32) int32 {
    var v5 bool = true
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 128
    v3 := make(map[int]bool)
    v4 := []uint32{4294967295, 4294967295}
    var v7 bool = true
    v2 += v2
    if v7 {
        if (-1 > -421) {
            for fi8 := 0; fi8 < 3; fi8++ {
                v7 = v7
                v9 := GenS1{F0: fi8}
                _ = fi8
                _ = v9
            }
        }
    }
    v2 = v2
    fmt.Println(v2, v7)
    if v7 {
        for fi10 := 0; fi10 < 18; fi10++ {
            v1[128] = "hello"
            var v11 uint = 170
            if (1 == 31) {
                var v12 string = ""
                v13 := GenS2{F0: 0, F1: ((("a" + v12) + "a") + (((v12 + (v12 + v12)) + v12) + "hello"))}
                v2 = (v2 >> (((v11 & v11)) & 63))
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = len(v3)
    }
    fmt.Println(v7)
    var v14 int = 2147483647
    v7 = v7
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

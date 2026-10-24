package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint64 }
func genFn0(p0 uint, p1 uint64, p2 string) int8 {
    fmt.Println(p1)
    if (((42 > -1) && true) && (!((false && (!((0.0 + 1e10) < (((1.5 * 1.5) - 0.0) + (3.14159 + float32(((-2.25 + (0.1 - (float64((-2.25 - (1e10 - 3.14159))) / 1.5))) - 3.14159))))))) && true))) {
        switch 32768 {
        case -329:
            v4 := []int{-623}
            _ = v4
        case 42:
            fmt.Println(p0, p1)
        default:
            p0 = 65535
        }
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 uint16) string {
    fmt.Println(p0)
    var v5 int8 = 42
    v6 := GenS1{F0: 405}
    _ = p0
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := []string{}
    v2 := []string{"hello", "foo bar"}
    var v3 float64 = (1e10 * 0.0)
    v7 := GenS1{F0: 4294967295}
    var v8 int32 = -2147483648
    v8 += (-(v8))
    v8 += (v8 & 0)
    v8 += (v8 % ((v8) | 1))
    v3 += 1.5
    var v9 int = 128
    v8 += 1
    fmt.Println(v3, v9, v8)
    v10 := GenS2{F0: ("hello" >= "a"), F1: 2147483647, F2: 637}
    fmt.Println(v8)
    v9 = int(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}

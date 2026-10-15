package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0() int {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 float32 = 0.0
        _ = fi5
        _ = v6
    }
    var v7 float64 = (1e10 - 0.1)
    _ = v7
    return 245
}
func genFn1(p0 int16, p1 int, p2 uint) int {
    p1 = (p1 + int(p2))
    v8 := GenS1{F0: (p1 >> (((p2 << ((p2) & 63))) & 63))}
    if (18446744073709551615 != 2) {
        switch p1 {
        case 0:
            v9 := []int{}
            _ = v9
        }
        v10 := []int{(p1 + 208), p1, (p1 % ((p1) | 1))}
        p0 = p0
        _ = v10
    }
    v11 := []int{int(p0), 333, (p1 + (p1 ^ (p1 << ((332) & 63))))}
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return (p1 % (((^(p1))) | 1))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 42
    v3 := make(map[int]int)
    var v4 uint32 = 65535
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    v4 = v4
    var v12 int64 = 1
    fmt.Println(v4)
    v13 := GenS2{F0: true}
    switch ("a" + "") {
    case "a":
        var v14 uint32 = v4
        _ = v14
    case "":
        v15 := GenS2{F0: ("\n\t" < "\n\t")}
        _ = v15
    case "\n\t":
        var v16 uint = 762
        _ = v16
    }
    v17 := []int{}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}

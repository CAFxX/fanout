package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 uint, p1 bool) string {
    var v5 uint8 = 0
    var v6 bool = (-579 < -243)
    switch "" {
    case "":
        v7 := []int{1, 278, 128}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 int32, p1 int8) int64 {
    if true {
        p1 += (p1 * p1)
        v8 := []int{42}
        if (!("" <= ("a" + "hello"))) {
            p0 = p0
            fmt.Println(p0)
        } else {
            p0 += 42
        }
        _ = v8
    }
    v9 := GenS1{F0: 357}
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 uint8 = 127
    v4 := make(map[int]string)
    v3 = 1
    var v10 uint32 = 4294967295
    v1["\n\t"] = false
    fmt.Println(v10, v10)
    if (64 >= 0) {
        v11 := GenS1{F0: 65535}
        var v12 bool = true
        if (-1 <= genFn1(2147483647, 127)) {
            var v13 bool = ((42 <= 42) && (0.0 >= 1.5))
            var v14 int8 = 127
            _ = v13
            _ = v14
        }
        _ = v11
        _ = v12
    } else {
        fmt.Println(v3, v3)
    }
    var v15 int64 = 111
    if ((!(951 >= 42)) && false) {
        var v16 int = 32768
        fmt.Println(v15, v10, v10)
        v15 = genFn1(0, 0)
        _ = v16
    } else {
        var v17 int = 19
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: (v3 + 86), F2: 0}
    var v19 uint64 = 1
    v3 += 202
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

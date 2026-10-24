package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64, p1 uint32) uint64 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float64 = 3.14159
        fmt.Println(fi4, v5)
        var v6 uint = 42
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-2147483648, -9223372036854775808, -9}
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 15
    var v3 bool = ("a" < "")
    var v8 uint32 = 2
    v9 := GenS1{F0: 255}
    var v10 int8 = v2
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int64 = 740
        _ = fi11
        _ = v12
    }
    if (v8 < 0) {
        var v13 int8 = (v10 & -1)
        var v14 uint64 = 1
        _ = v13
        _ = v14
    } else {
        v15 := GenS2{F0: 0, F1: 750}
        v3 = false
        var v16 uint16 = 2
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 2, F1: 0}
    switch "hello" {
    case "a":
        var v18 bool = v3
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v3 = v3
        var v20 int64 = -694
        fmt.Println(v3)
        _ = fi19
        _ = v20
    }
    v2 = (v2 + (v10 + v10))
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(v8)
        _ = fi21
    }
    if (-1 == v2) {
        var v22 uint8 = 2
        _ = v22
    } else {
        v2 += -19
        var v23 int16 = -1
        v24 := GenS2{F0: 4294967295, F1: 18446744073709551615}
        _ = v23
        _ = v24
    }
    v3 = false
    fmt.Println(v2)
    for fi25 := 0; fi25 < 3; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int32 = -1
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    if false {
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 int8 = v10
            _ = fi28
            _ = v29
        }
        if true {
            var v30 uint32 = (v8 / (((v8 << ((255) & 63))) | 1))
            var v31 float32 = -2.25
            _ = v30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        v33 := v1["a"]
        fmt.Println(v3, v10)
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := []int{(^(fi34)), 1}
        v36 := GenS3{F0: 42}
        _ = len(v1)
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        v3 = (true || v3)
        fmt.Println(fi37, v10)
        switch "hello" {
        case "":
            v38 := GenS3{F0: 65535}
            _ = v38
        case "a":
            var v39 bool = (!false)
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}

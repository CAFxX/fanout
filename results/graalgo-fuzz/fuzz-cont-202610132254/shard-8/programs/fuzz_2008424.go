package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 uint, p1 uint8, p2 string) int64 {
    v5 := GenS1{F0: "hello", F1: 4294967295}
    fmt.Println(p2, p0, p1)
    v6 := []int{32767, -1, -507, -723}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 int16) uint8 {
    var v7 int8 = 42
    if true {
        var v8 string = "hello"
        var v9 uint16 = (p0 & (p0 >> ((42) & 63)))
        _ = v8
        _ = v9
    } else {
        if (1 >= 0) {
            fmt.Println(p0, v7, p1)
        } else {
            v10 := GenS1{F0: "hello", F1: 0}
            _ = v10
        }
    }
    var v11 uint16 = p0
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = "\n\t"
        switch (v13 + v13) {
        case "hello":
            if ((fi12 == (fi12 + (fi12 * 2147483647))) && false) {
                v7 = (v7 | v7)
            } else {
                fmt.Println(p0, v13, v11)
            }
        case "a":
            p1 += 0
        case "":
            p1 = p1
        default:
            fmt.Println(v13)
        }
        v14 := GenS1{F0: ((v13 + v13) + (v13 + (("hello" + "a") + v13))), F1: 227}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return uint8(p0)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 uint8 = 0
    var v4 uint16 = 1
    if ((!false) || false) {
        fmt.Println(v4)
        for fi15 := 0; fi15 < 2; fi15++ {
            v4 += 2
            _ = fi15
        }
    }
    fmt.Println(v3, v3)
    v4 = (v4 << ((432) & 63))
    var v16 string = "\n\t"
    if true {
        v17 := GenS1{F0: v16, F1: 65535}
        _ = v17
    } else {
        var v18 int64 = 42
        v18 = v18
        v19 := GenS1{F0: v16, F1: 42}
        _ = v18
        _ = v19
    }
    if false {
        switch 2147483647 {
        case -456:
            var v20 uint = 4294967295
            _ = v20
        case 32767:
            var v21 int32 = -1
            _ = v21
        }
    }
    v22 := []int{-867, 4294967295, -819, 32768}
    if ((62 > -1) && (1000 >= 60)) {
        var v23 uint8 = (v3 >> ((327) & 63))
        _ = v23
    }
    v24 := v2[""]
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}

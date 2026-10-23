package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int) int32 {
    for fi4 := 0; fi4 < 3; fi4++ {
        if true {
            p0 -= (p0 & fi4)
        } else {
            v5 := GenS1{F0: -98, F1: "\n\t"}
            _ = v5
        }
        v6 := GenS2{F0: 18446744073709551615}
        v7 := GenS3{F0: "a"}
        _ = fi4
        _ = v6
        _ = v7
    }
    if true {
        v8 := []int{(^(p0)), (p0 & 127), 0}
        _ = v8
    } else {
        var v9 int16 = -478
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 bool = (255 <= 1)
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 string = ""
    _ = p0
    _ = v12
    return 447
}
func genFn1(p0 int64, p1 uint16) uint {
    v13 := GenS1{F0: 42, F1: ("foo bar" + "")}
    if (p1 == (p1 << ((876) & 63))) {
        var v14 int16 = 1
        v14 -= 0
        _ = v14
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            p0 = (-(p0))
            var v16 uint = 819
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            p1 = p1
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return 65535
}

func main() {
    var v1 int32 = 42
    var v2 int32 = 317
    v3 := make(map[uint32]uint32)
    v18 := GenS1{F0: -98, F1: ""}
    for fi19 := 0; fi19 < 9; fi19++ {
        if true {
            _ = len(v3)
        }
        _ = fi19
    }
    fmt.Println(v1)
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS4{F0: -682, F1: true}
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 0, F1: "foo bar"}
    var v23 uint32 = 65535
    switch "" {
    case "foo bar":
        var v24 uint16 = 32767
        _ = v24
    case "a":
        for fi25 := 0; fi25 < 4; fi25++ {
            fmt.Println(v1, v2, v23)
            _ = fi25
        }
    }
    var v26 uint16 = 255
    fmt.Println(v23, v2, v2)
    v27 := []int{-412, -723, -173, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
}

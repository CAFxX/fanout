package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    v5 := GenS1{F0: 128, F1: 943, F2: 4294967295}
    switch "a" {
    case "":
        var v6 int8 = -128
        _ = v6
    default:
        v7 := GenS2{F0: 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(fi8, fi8)
        var v9 uint64 = 1
        v9 = v9
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    return -710
}
func genFn1(p0 int, p1 bool) uint16 {
    v10 := []int{(^(p0)), p0, p0}
    var v11 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint64 = 2
    v4 := make(map[string]int)
    _ = len(v4)
    v12 := GenS3{F0: 244, F1: 0}
    for fi13 := 0; fi13 < 17; fi13++ {
        if v1 {
            v14 := []int{(fi13 * (fi13 ^ fi13)), fi13}
            v2[("hello" + "\n\t")] = (fi13 / (((fi13 * fi13)) | 1))
            v1 = v1
            _ = v14
        } else {
            v15 := v4["hello"]
            v16 := []int{fi13, (v15 << ((268) & 63))}
            _ = v15
            _ = v16
        }
        var v17 bool = (true && (-88 != -313))
        switch ("\n\t" + (("a" + ((("" + (("hello" + "foo bar") + "hello")) + "a") + "\n\t")) + "")) {
        case "\n\t":
            v2["a"] = fi13
        case "foo bar":
            var v18 int32 = -1
            _ = v18
        }
        _ = fi13
        _ = v17
    }
    var v19 uint64 = v3
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS4{F0: 490, F1: 1}
        _ = fi20
        _ = v21
    }
    var v22 bool = (v1 || false)
    v23 := []int{}
    v1 = (!v1)
    for fi24 := 0; fi24 < 4; fi24++ {
        v1 = (!(846 >= -839))
        _ = fi24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 int8 = -1
        v1 = true
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 float32 = ((float32(1e10) - 0.1) - 1e10)
        for fi29 := 0; fi29 < 10; fi29++ {
            if (!v1) {
                var v30 uint = 0
                _ = v30
            } else {
                var v31 int16 = 1
                var v32 uint16 = 925
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        v22 = v22
        _ = fi27
        _ = v28
    }
    switch "hello" {
    case "":
        var v33 float32 = 1.5
        _ = v33
    case "hello":
        switch 75 {
        case 492:
            var v34 int16 = -32768
            _ = v34
        case 1:
            fmt.Println(v3)
        case 421:
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 uint = 42
                _ = fi35
                _ = v36
            }
        }
    default:
        switch 32767 {
        case 42:
            v3 -= v3
        }
    }
    var v37 bool = v22
    var v38 int8 = 1
    v39 := v2["foo bar"]
    v1 = v1
    v22 = v1
    var v40 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}

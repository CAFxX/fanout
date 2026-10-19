package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
func genFn0(p0 bool, p1 bool, p2 int16) uint8 {
    if (32767 > 128) {
        var v5 int16 = p2
        _ = v5
    }
    var v6 uint = 112
    p2 = (p2 & (p2 * (p2 * (p2 >> (((^(v6))) & 63)))))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 154
}
func genFn1(p0 int8) bool {
    var v7 string = "hello"
    var v8 float64 = (-2.25 - (((3.14159 / 1.0) + 3.14159) / -2.25))
    if false {
        v9 := []int{468, -9223372036854775808, 685}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(fi10, v7)
            _ = fi10
        }
        v11 := []int{4294967295, -1, 32767}
        v12 := GenS1{F0: true, F1: 806}
        _ = v11
        _ = v12
    }
    switch (v7 + "") {
    case "a":
        v13 := GenS1{F0: true, F1: 422}
        _ = v13
    case "hello":
        if (int64(p0) <= 42) {
            var v14 uint = 4294967295
            switch -161 {
            case 0:
                v14 -= v14
            }
            v7 = ""
            _ = v14
        } else {
            v15 := []int{227, -9223372036854775808}
            fmt.Println(p0)
            _ = v15
        }
    case "foo bar":
        v16 := []int{}
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 int32 = -2147483648
    v4 := []int{}
    v3 += v3
    v17 := v1[79]
    for fi18 := 0; fi18 < 1; fi18++ {
        switch "" {
        case "a":
            _ = len(v4)
        case "":
            v19 := GenS1{F0: (1000 < 32767), F1: fi18}
            _ = v19
        }
        v20 := GenS2{F0: 128, F1: 1}
        _ = fi18
        _ = v20
    }
    v17 -= (^(v17))
    if false {
        fmt.Println(v17)
    }
    v21 := GenS1{F0: true, F1: -242}
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 float32 = 1.5
        _ = fi22
        _ = v23
    }
    v3 = v3
    var v24 int32 = -1
    for fi25 := 0; fi25 < 20; fi25++ {
        delete(v2, (fi25 * fi25))
        var v26 uint64 = 18446744073709551615
        _ = fi25
        _ = v26
    }
    v24 = -2147483648
    for fi27 := 0; fi27 < 14; fi27++ {
        if true {
            v3 -= -1
            v3 = (-(v24))
        }
        if false {
            v3 -= -1
            v3 = v24
        }
        _ = fi27
    }
    var v28 uint64 = 721
    v24 -= v3
    v29 := GenS3{F0: v24, F1: 2}
    if false {
        fmt.Println(v3, v28)
    }
    var v30 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}

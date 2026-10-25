package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint; F2 uint16 }
func genFn0() int8 {
    var v7 int64 = 946
    v8 := []int{}
    for fi9 := 0; fi9 < 19; fi9++ {
        if true {
            fmt.Println(v7, fi9)
        } else {
            v7 += v7
            if len(v8) > 0 {
                v8[0] = fi9
            }
            var v10 uint8 = 127
            _ = v10
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 14
}

func main() {
    v1 := []int{128, 86, 127}
    var v2 uint = 0
    var v3 int64 = 42
    var v4 string = "foo bar"
    var v5 bool = (32767 < 130)
    v6 := []int{}
    var v11 uint8 = 2
    if true {
        var v12 uint8 = v11
        v13 := GenS1{F0: -1, F1: -1, F2: (v11 | v11)}
        for fi14 := 0; fi14 < 2; fi14++ {
            switch (^(fi14)) {
            case 32768:
                fmt.Println(v11, v12, v11)
            case 128:
                v15 := GenS1{F0: 637, F1: -1, F2: (v12 << (((v2 * (v2 % ((v2) | 1)))) & 63))}
                _ = v15
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    v16 := GenS2{F0: (v4 + "\n\t")}
    switch (v4 + (v4 + v4)) {
    case "foo bar":
        var v17 string = ""
        _ = v17
    case "hello":
        v11 -= (^(v11))
    case "\n\t":
        var v18 uint64 = 515
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := GenS2{F0: "foo bar"}
        fmt.Println(fi19)
        _ = fi19
        _ = v20
    }
    fmt.Println(v2, v2)
    fmt.Println(v5, v5)
    v1 = append(v1, -193)
    for fi21 := 0; fi21 < 16; fi21++ {
        v6 = append(v6, fi21)
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 string = (v4 + ("\n\t" + v4))
            if len(v1) > 0 {
                v1[0] = (fi22 / (((-(fi22))) | 1))
            }
            var v24 uint = 833
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = len(v1)
        _ = fi21
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int64 = -1
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v2 = 0
        fmt.Println(fi27, v11)
        _ = fi27
    }
    switch (v4 + v4) {
    case "foo bar":
        _ = len(v1)
    case "":
        if len(v1) > 0 {
            v1[0] = -556
        }
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(v11, v3)
        v29 := GenS3{F0: fi28, F1: v2, F2: uint16(v2)}
        _ = fi28
        _ = v29
    }
    v3 += 190
    var v30 string = v4
    v5 = (true && ((0 == -1) && v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

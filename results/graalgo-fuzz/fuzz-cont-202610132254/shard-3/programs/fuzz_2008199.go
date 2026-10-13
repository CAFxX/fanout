package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }

func main() {
    v1 := []string{}
    var v2 string = (("a" + ("" + "a")) + "a")
    v3 := []int{}
    var v4 uint64 = 0
    v5 := make(map[uint32]int)
    var v6 int8 = 0
    switch (("a" + ("" + v2)) + v2) {
    case "\n\t":
        _ = len(v5)
    case "a":
        var v7 uint8 = 128
        _ = v7
    default:
        v4 = (v4 >> ((42) & 63))
    }
    fmt.Println(v2, v6)
    delete(v5, 2147483647)
    for fi8 := 0; fi8 < 13; fi8++ {
        if false {
            for fi9 := 0; fi9 < 10; fi9++ {
                v5[486] = (-(fi8))
                v10 := v5[0]
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: (v4 & (v4 ^ v4))}
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: 643}
                fmt.Println(v2)
                v2 = v2
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            v2 = v2
            v2 = v2
        }
        fmt.Println(fi8, v6)
        _ = fi8
    }
    var v14 int16 = -1
    var v15 uint32 = 2147483648
    var v16 uint64 = 4294967295
    v4 = v4
    fmt.Println(v2)
    for fi17 := 0; fi17 < 15; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v18 int32 = -586
        v19 := v5[v15]
        _ = fi17
        _ = v18
        _ = v19
    }
    if ((3.14159 * 0.0) <= 3.14159) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 2147483647 {
        case -706:
            switch -2147483648 {
            case -28:
                fmt.Println(v14)
            }
        case 78:
            v20 := GenS2{F0: v4, F1: -356, F2: 0}
            _ = v20
        }
    } else {
        v2 = v2
        var v21 uint16 = 42
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 uint64 = (v4 ^ 185)
        var v24 uint16 = 598
        if (uint64(v15) >= (v4 % (((v23 ^ (v23 | v16))) | 1))) {
            v25 := v5[v15]
            v24 -= v24
            _ = v25
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    switch (("" + v2) + (v2 + v2)) {
    case "hello":
        var v26 uint16 = 2
        _ = v26
    case "":
        v2 = "\n\t"
    case "\n\t":
        var v27 float32 = ((1e10 / 1.5) - 1.5)
        _ = v27
    }
    var v28 uint64 = (v16 & (v4 << ((343) & 63)))
    v29 := GenS1{F0: v4}
    var v30 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

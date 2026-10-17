package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int64 = 0
    var v2 int16 = 1
    var v3 uint64 = 320
    v3 = v3
    switch 1 {
    case 42:
        for fi4 := 0; fi4 < 18; fi4++ {
            v5 := GenS1{F0: -83, F1: (v3 >> ((2) & 63))}
            v2 -= (v2 & v2)
            _ = fi4
            _ = v5
        }
    default:
        if true {
            fmt.Println(v3, v1)
        } else {
            v6 := GenS1{F0: 54, F1: (^(v3))}
            fmt.Println(v1, v1, v3)
            _ = v6
        }
    }
    fmt.Println(v2)
    fmt.Println(v2, v2, v3)
    fmt.Println(v1, v1, v2)
    if (1 != 0) {
        v7 := []int{}
        var v8 uint32 = 2147483648
        _ = v7
        _ = v8
    } else {
        v2 -= v2
        v2 = 600
        var v9 uint64 = 983
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{1, fi10, (^(fi10))}
        _ = fi10
        _ = v11
    }
    v2 = -32768
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint64 = v3
        _ = fi12
        _ = v13
    }
    v3 = 2
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := []int{fi14, fi14, 892}
        if len(v15) > 0 {
            v15[0] = (fi14 / ((fi14) | 1))
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v1)
    for fi16 := 0; fi16 < 4; fi16++ {
        if false {
            var v17 int64 = (^(v1))
            var v18 uint64 = v3
            _ = v17
            _ = v18
        } else {
            var v19 int = -65
            var v20 uint8 = 42
            var v21 int8 = -128
            _ = v19
            _ = v20
            _ = v21
        }
        v2 -= -1
        v22 := GenS1{F0: -19, F1: v3}
        _ = fi16
        _ = v22
    }
    if ((true && true) && (!false)) {
        v23 := []int{}
        for fi24 := 0; fi24 < 6; fi24++ {
            _ = len(v23)
            v2 = v2
            _ = len(v23)
            _ = fi24
        }
        _ = v23
    }
    v3 = v3
    switch "" {
    case "hello":
        v2 = v2
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := []int{}
        for fi27 := 0; fi27 < 13; fi27++ {
            switch "" {
            case "":
                v2 += 1
            default:
                if true {
                    var v28 int32 = 1
                    _ = v28
                } else {
                    v2 += 1
                    var v29 int64 = (v1 + v1)
                    var v30 uint16 = 0
                    _ = v29
                    _ = v30
                }
            }
            _ = fi27
        }
        fmt.Println(fi25, v3)
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}

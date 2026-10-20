package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0() int32 {
    var v4 uint16 = 156
    v5 := []int{127, 138, 32768}
    for fi6 := 0; fi6 < 8; fi6++ {
        v4 = v4
        v4 += (v4 - v4)
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: 94, F1: false}
            var v9 uint32 = 65535
            if true {
                v4 = (v4 >> ((0) & 63))
                v5 = append(v5, (fi6 * 128))
                v10 := GenS2{F0: 227, F1: 42, F2: 4294967295}
                _ = v10
            } else {
                v11 := GenS2{F0: 0, F1: 255, F2: 0}
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    switch "foo bar" {
    case "hello":
        if true {
            v4 = v4
            v12 := GenS2{F0: 1, F1: 750, F2: 1}
            _ = v12
        }
    case "a":
        v5 = append(v5, 149)
    case "foo bar":
        v4 += v4
    }
    var v13 int16 = 1000
    _ = v13
    _ = v4
    _ = v5
    return -583
}

func main() {
    v1 := []uint32{42, 388}
    v2 := []uint32{693, 853, 1, 0}
    v3 := make(map[uint32]string)
    var v14 uint64 = 0
    _ = len(v3)
    v14 = v14
    for fi15 := 0; fi15 < 4; fi15++ {
        v14 -= v14
        _ = fi15
    }
    v16 := GenS2{F0: 60, F1: 111, F2: v14}
    v14 -= 194
    fmt.Println(v14)
    v17 := GenS1{F0: 65535, F1: false}
    var v18 int16 = -370
    if false {
        var v19 uint32 = 0
        if len(v1) > 0 {
            v1[0] = uint32(v14)
        }
        fmt.Println(v18, v18)
        _ = v19
    } else {
        v20 := GenS1{F0: 65535, F1: true}
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        switch "hello" {
        case "foo bar":
            switch 969 {
            case -489:
                v18 += v18
            }
        case "a":
            var v22 uint32 = uint32(v14)
            _ = v22
        case "":
            v23 := v3[2147483647]
            _ = v23
        }
        var v24 int = (fi21 % ((fi21) | 1))
        switch fi21 {
        case 1:
            var v25 uint64 = 1
            _ = v25
        case -544:
            v24 += fi21
        }
        _ = fi21
        _ = v24
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    switch "foo bar" {
    case "hello":
        if true {
            v14 = (v14 ^ 846)
            switch 907 {
            case 2147483647:
                v26 := GenS1{F0: 0, F1: true}
                _ = v26
            case 128:
                var v27 int = 32767
                _ = v27
            default:
                var v28 float64 = 1.5
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 uint32 = 116
                v31 := GenS1{F0: v30, F1: (!true)}
                v14 -= v14
                _ = fi29
                _ = v30
                _ = v31
            }
            switch -2147483648 {
            case -1:
                v14 -= v14
            default:
                switch (("hello" + "foo bar") + "\n\t") {
                case "foo bar":
                    var v32 uint16 = 2
                    _ = v32
                case "hello":
                    var v33 uint16 = 65535
                    _ = v33
                }
            }
            if true {
                v2 = append(v2, 782)
                v34 := GenS3{F0: (true && true), F1: 2147483647}
                _ = v34
            }
        }
    case "a":
        var v35 int = 32767
        _ = v35
    case "":
        v18 = v18
    default:
        v18 += (v18 & v18)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 int64 }
func genFn0() uint32 {
    var v7 int64 = 42
    switch "\n\t" {
    case "a":
        var v8 int8 = 52
        _ = v8
    case "hello":
        var v9 int = 4294967295
        _ = v9
    case "":
        for fi10 := 0; fi10 < 12; fi10++ {
            switch "foo bar" {
            case "":
                v11 := GenS1{F0: (^(v7)), F1: (fi10 >> ((255) & 63))}
                _ = v11
            case "\n\t":
                if false {
                    switch (fi10 >> ((0) & 63)) {
                    case 0:
                        var v12 int32 = 42
                        _ = v12
                    case -48:
                        v7 += 808
                    }
                    v13 := []int{4294967295, fi10, (fi10 & (fi10 & 9223372036854775807)), fi10}
                    v7 -= v7
                    _ = v13
                } else {
                    if (true && (v7 < v7)) {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            fmt.Println(fi10)
                            _ = fi14
                        }
                        v7 = (-(v7))
                    }
                }
            default:
                var v15 bool = true
                _ = v15
            }
            var v16 int16 = 1
            if true {
                v17 := GenS2{F0: 127, F1: ("a" + ""), F2: 723}
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 bool = (true || ("a" >= (("" + "foo bar") + "\n\t")))
                    fmt.Println(fi18, fi18, fi18)
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                var v20 float32 = float32(1e10)
                if false {
                    var v21 uint32 = 365
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if true {
                            var v23 uint8 = 189
                            var v24 uint64 = 673
                            _ = v23
                            _ = v24
                        }
                        var v25 uint16 = 151
                        v21 = (v21 ^ (^(v21)))
                        v26 := GenS3{F0: (v21 << ((210) & 63))}
                        _ = fi22
                        _ = v25
                        _ = v26
                    }
                    v27 := []int{(fi10 * 42), (fi10 - 2147483647)}
                    _ = v21
                    _ = v27
                }
                _ = v20
            }
            _ = fi10
            _ = v16
        }
    default:
        v28 := []int{42}
        _ = v28
    }
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    v6 := make(map[int]uint32)
    var v29 uint64 = 2
    switch 32767 {
    case -2147483648:
        fmt.Println(v2, v29)
    }
    v30 := GenS4{F0: 119, F1: 122, F2: 0}
    var v31 string = "hello"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v32 := GenS5{F0: -1}
    if (2147483648 > 65535) {
        v33 := GenS5{F0: 9223372036854775807}
        var v34 uint8 = 1
        v29 = (v29 + 87)
        _ = v33
        _ = v34
    }
    delete(v4, genFn0())
    v2 = (^(v2))
    fmt.Println(v31, v29)
    var v35 int32 = 1
    v29 -= v29
    v36 := GenS6{F0: 42, F1: v2}
    for fi37 := 0; fi37 < 10; fi37++ {
        switch -25 {
        case 32767:
            var v38 float32 = (1e10 * -2.25)
            _ = v38
        default:
            fmt.Println(v35, v2, v35)
        }
        _ = fi37
    }
    var v39 int16 = -1
    var v40 int16 = (^(v39))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

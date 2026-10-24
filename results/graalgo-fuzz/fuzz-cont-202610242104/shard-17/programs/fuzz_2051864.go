package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 int16) uint16 {
    fmt.Println(p0)
    p0 = p0
    v4 := GenS1{F0: 100, F1: 1, F2: 597}
    var v5 uint16 = 0
    var v6 uint64 = 2
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v5
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 string = (("foo bar" + "a") + "")
    var v7 int64 = 42
    if false {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: 42, F1: 62, F2: 42}
            _ = fi8
            _ = v9
        }
        fmt.Println(v3)
        var v10 int32 = 42
        _ = v10
    }
    var v11 float32 = 1e10
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            if true {
                v7 = v7
                v3 = "\n\t"
            }
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        if true {
            delete(v2, v3)
        }
        _ = fi12
    }
    var v15 uint16 = 2
    if true {
        var v16 uint16 = (v15 / ((genFn0(42)) | 1))
        var v17 float32 = v11
        fmt.Println(v16, v7, v16)
        _ = v16
        _ = v17
    } else {
        if false {
            v18 := GenS2{F0: 1, F1: 1000, F2: (("a" + v3) < "a")}
            var v19 int8 = 1
            fmt.Println(v7, v11, v15)
            _ = v18
            _ = v19
        }
        v11 = v11
        v20 := GenS3{F0: 363, F1: (v15 & (v15 & genFn0(618)))}
        _ = v20
    }
    v15 -= (v15 >> ((345) & 63))
    v21 := GenS4{F0: 751, F1: v15, F2: 172}
    if true {
        for fi22 := 0; fi22 < 20; fi22++ {
            v3 = (v3 + (("hello" + ((v3 + ("\n\t" + v3)) + "")) + "a"))
            if (!true) {
                var v23 int64 = 42
                v24 := GenS5{F0: (fi22 ^ (fi22 + (fi22 >> ((42) & 63)))), F1: 18446744073709551615, F2: 954}
                _ = v23
                _ = v24
            }
            if (42 >= -1) {
                fmt.Println(v7)
                switch (^(fi22)) {
                case 890:
                    var v25 int = (-(fi22))
                    _ = v25
                case 42:
                    var v26 float32 = v11
                    _ = v26
                case 610:
                    var v27 string = v3
                    _ = v27
                default:
                    fmt.Println(v15, v11)
                }
            } else {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v29 := GenS4{F0: 978, F1: (v15 % ((2) | 1)), F2: 1}
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi22
        }
        v30 := GenS3{F0: 18446744073709551615, F1: v15}
        _ = v30
    }
    fmt.Println(v15)
    v31 := v2[("" + v3)]
    fmt.Println(v7, v31)
    switch -2147483648 {
    case 32768:
        var v32 int16 = 634
        _ = v32
    default:
        switch "\n\t" {
        case "a":
            v2[(v3 + "\n\t")] = ""
        case "hello":
            v33 := []int{-124, -979}
            _ = v33
        }
    }
    v34 := GenS6{F0: 2, F1: 42, F2: 2}
    v35 := []int{127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    var v3 uint8 = v1
    var v4 int64 = 1
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 1, F1: "hello", F2: 1}
        if true {
            v8 := []int{9223372036854775807, fi6}
            v9 := GenS1{F0: 671, F1: "foo bar", F2: -128}
            var v10 float32 = 3.14159
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := GenS1{F0: 589, F1: ((("a" + "foo bar") + "a") + ""), F2: 100}
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    if (!false) {
        switch 1 {
        case 32767:
            v1 = v1
        case -257:
            fmt.Println(v4)
        default:
            if (!true) {
                switch ("a" + "") {
                case "foo bar":
                    v12 := GenS2{F0: 0, F1: 4294967295, F2: -611}
                    _ = v12
                }
                if (true || (v4 == 9223372036854775807)) {
                    v3 -= (v3 * (v3 & (v1 ^ v1)))
                } else {
                    fmt.Println(v3, v3, v1)
                    var v13 uint8 = v1
                    _ = v13
                }
                var v14 string = "foo bar"
                _ = v14
            }
        }
    }
    _ = len(v5)
    fmt.Println(v1)
    v1 -= (v3 * 44)
    v15 := []int{}
    v1 = (v1 - 141)
    var v16 bool = true
    v17 := GenS1{F0: 1, F1: "foo bar", F2: 100}
    fmt.Println(v1, v3)
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 int16 = -79
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v1 -= (v1 << ((349) & 63))
        _ = fi20
    }
    fmt.Println(v4, v16)
    if v16 {
        v21 := GenS2{F0: 0, F1: 413, F2: 4294967295}
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint = 2
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v24 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

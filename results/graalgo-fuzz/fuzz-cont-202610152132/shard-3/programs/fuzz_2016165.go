package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }
type GenS3 struct { F0 string; F1 bool }
func genFn0() string {
    v4 := []int{736, 42, 9223372036854775807, 0}
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            if ((1 <= 2147483647) || true) {
                if true {
                    var v7 int64 = -362
                    v8 := GenS1{F0: 127, F1: 0, F2: 42}
                    var v9 bool = ((v7 * (v7 % (((v7 % (((v7 << ((42) & 63))) | 1))) | 1))) > 42)
                    _ = v7
                    _ = v8
                    _ = v9
                }
            } else {
                switch (fi6 | (fi5 ^ 2147483647)) {
                case 128:
                    var v10 int16 = -32768
                    _ = v10
                case 789:
                    v4 = append(v4, -2147483648)
                default:
                    _ = len(v4)
                }
                var v11 int16 = 0
                switch 270 {
                case 0:
                    var v12 int32 = -1
                    _ = v12
                case 675:
                    var v13 int8 = 105
                    _ = v13
                }
                _ = v11
            }
            v14 := GenS1{F0: 1, F1: 100, F2: 690}
            _ = fi6
            _ = v14
        }
        _ = fi5
    }
    var v15 uint = 18446744073709551615
    for fi16 := 0; fi16 < 8; fi16++ {
        v15 -= v15
        _ = fi16
    }
    _ = v15
    _ = v4
    return ""
}
func genFn1(p0 int16, p1 int16) int32 {
    var v17 uint = 1
    var v18 uint8 = 128
    var v19 uint32 = 0
    v20 := GenS2{F0: -684, F1: 112, F2: 42}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 uint = 43
    var v2 int8 = 27
    var v3 uint8 = 1
    var v21 uint = 42
    switch 2147483647 {
    case -9223372036854775808:
        fmt.Println(v3, v1, v3)
    case 1:
        var v22 uint16 = 235
        _ = v22
    case 42:
        v23 := []int{}
        _ = v23
    default:
        var v24 uint64 = 18446744073709551615
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := []int{(fi26 & 4294967295), (fi25 * -722), fi26, (fi26 - (fi26 | fi26))}
            _ = fi26
            _ = v27
        }
        v1 = (v21 % ((248) | 1))
        _ = fi25
    }
    v28 := GenS3{F0: ((("" + ("\n\t" + "")) + "a") + "foo bar"), F1: true}
    var v29 uint16 = 802
    v30 := []int{}
    v31 := GenS3{F0: ("a" + "hello"), F1: (42 < v2)}
    fmt.Println(v1, v1)
    switch 9223372036854775807 {
    case -399:
        var v32 string = ("hello" + ("hello" + "hello"))
        _ = v32
    case 32767:
        v33 := GenS1{F0: v3, F1: v2, F2: 2}
        _ = v33
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int64 {
    var v5 float64 = 1.5
    if (4294967295 <= 2) {
        fmt.Println(v5)
        if (!(-787 != -1)) {
            v6 := GenS1{F0: 47}
            switch 986 {
            case 4294967295:
                v7 := GenS1{F0: 4294967295}
                _ = v7
            default:
                var v8 uint32 = 65535
                _ = v8
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint32 = 42
                v10 -= (v10 ^ (^(v10)))
                switch "" {
                case "hello":
                    fmt.Println(fi9, v5)
                case "\n\t":
                    v11 := GenS1{F0: 4294967295}
                    _ = v11
                case "foo bar":
                    v12 := []int{(fi9 ^ fi9)}
                    _ = v12
                default:
                    v13 := GenS1{F0: 18446744073709551615}
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            _ = v6
        }
        switch "" {
        case "a":
            v14 := []int{9223372036854775807}
            _ = v14
        case "hello":
            if false {
                fmt.Println(v5)
            }
        case "":
            v5 = v5
        }
    }
    var v15 float32 = 0.1
    v16 := []int{}
    v17 := GenS1{F0: 638}
    _ = v15
    _ = v16
    _ = v17
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 int, p2 int8) int64 {
    var v18 float64 = 0.0
    var v19 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return int64(p1)
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 int16 = 32767
    var v3 uint64 = 1
    var v4 float32 = (((0.0 / 1.0) * 1e10) + 1.5)
    v4 = v4
    v3 = (v3 << ((4294967295) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v20 uint = 1
    var v21 string = ""
    fmt.Println(v2, v3, v3)
    fmt.Println(v21, v3)
    v22 := GenS1{F0: (v3 % ((v3) | 1))}
    v1 = append(v1, 42)
    v23 := GenS1{F0: v3}
    v2 = (^(v2))
    var v24 uint32 = 1
    v25 := GenS1{F0: v3}
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 float64 = 0.0
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS1{F0: (v3 | (^(v3)))}
        fmt.Println(v21, v4, v20)
        _ = fi28
        _ = v29
    }
    v24 -= (^(v24))
    v30 := GenS1{F0: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}

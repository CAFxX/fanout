package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    p0 -= (p0 - (p0 - 149))
    var v6 int32 = 1
    var v7 string = "foo bar"
    _ = p0
    _ = v6
    _ = v7
    return 65535
}
func genFn1() uint {
    var v8 int64 = -337
    if ((-2.25 - 0.1) != 1e10) {
        fmt.Println(v8)
        v9 := []int{}
        _ = v9
    }
    switch 9223372036854775807 {
    case 32767:
        for fi10 := 0; fi10 < 15; fi10++ {
            switch fi10 {
            case 128:
                var v11 uint16 = 0
                _ = v11
            default:
                for fi12 := 0; fi12 < 14; fi12++ {
                    if true {
                        var v13 uint8 = 131
                        v13 = v13
                        v8 += v8
                        _ = v13
                    }
                    _ = fi12
                }
            }
            v14 := []int{fi10}
            for fi15 := 0; fi15 < 2; fi15++ {
                switch (fi10 - fi15) {
                case -802:
                    v14 = append(v14, (fi10 | fi15))
                case 4294967295:
                    v8 = (v8 * v8)
                case 2147483647:
                    switch "hello" {
                    case "a":
                        var v16 int32 = 2147483647
                        _ = v16
                    }
                default:
                    v17 := GenS1{F0: fi10, F1: -2147483648, F2: 1}
                    _ = v17
                }
                _ = fi15
            }
            _ = fi10
            _ = v14
        }
    default:
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 int16 = 1000
            switch fi18 {
            case -1:
                var v20 int32 = 42
                _ = v20
            case -274:
                v21 := []int{(^(fi18))}
                _ = v21
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                fmt.Println(v19, fi22)
                v19 -= 1000
                _ = fi22
            }
            _ = fi18
            _ = v19
        }
    }
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = ((((1e10 * 0.0) + float32(-2.25)) / float32((0.0 + 3.14159))) * 1.5)
    v3 := []string{""}
    var v4 float64 = (3.14159 - 0.0)
    v5 := make(map[string]string)
    var v23 uint = genFn1()
    var v24 uint32 = (^(v1))
    switch 1 {
    case -9223372036854775808:
        var v25 float64 = v4
        _ = v25
    }
    var v26 uint64 = 986
    switch "\n\t" {
    case "hello":
        var v27 uint16 = 0
        _ = v27
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v28 := GenS1{F0: -1, F1: 966, F2: uint(v1)}
    v23 = (v23 >> (((^(v23))) & 63))
    fmt.Println(v23, v4, v2)
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 int8 = 101
        _ = fi29
        _ = v30
    }
    switch ("" + "\n\t") {
    case "\n\t":
        fmt.Println(v1, v26, v24)
    default:
        v31 := GenS1{F0: 9223372036854775807, F1: -132, F2: v23}
        _ = v31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        if true {
            v24 = 599
            v5[""] = "foo bar"
        } else {
            v33 := GenS1{F0: 673, F1: 42, F2: v23}
            _ = v33
        }
        _ = fi32
    }
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}

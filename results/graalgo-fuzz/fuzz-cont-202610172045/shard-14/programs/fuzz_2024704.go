package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint; F2 int }
func genFn0(p0 int64) int64 {
    for fi5 := 0; fi5 < 1; fi5++ {
        if (fi5 > -678) {
            v6 := GenS1{F0: 65535, F1: 375, F2: 213}
            v7 := []int{}
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    p0 += (p0 % ((9223372036854775807) | 1))
    _ = p0
    return 887
}

func main() {
    v1 := []int{42}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v8 := GenS1{F0: 9, F1: 326, F2: 18446744073709551615}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = 2
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 969, F1: 843, F2: 2}
    v1 = append(v1, -110)
    for fi12 := 0; fi12 < 5; fi12++ {
        switch ("" + "hello") {
        case "foo bar":
            _ = len(v1)
        case "":
            var v13 int16 = 633
            _ = v13
        }
        var v14 uint16 = 65535
        _ = fi12
        _ = v14
    }
    var v15 int8 = 127
    if (42 > (-(v15))) {
        if false {
            if (true && (255 >= 0)) {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v3[216] = 1
                    var v16 bool = false
                    _ = v16
                }
                if (100 >= -39) {
                    var v17 uint32 = 89
                    _ = v17
                }
                v2 -= v2
            } else {
                var v18 int64 = (v2 * v2)
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint64 = 1
                v21 := GenS2{F0: uint(v20), F1: 573, F2: 813}
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            v4[65535] = ("" + "a")
        }
        var v22 uint16 = 1
        switch 2147483647 {
        case 638:
            fmt.Println(v22)
        default:
            v23 := GenS3{F0: (v15 | v15), F1: 906, F2: 802}
            _ = v23
        }
        _ = v22
    }
    v24 := v3[427]
    var v25 int16 = 0
    for fi26 := 0; fi26 < 2; fi26++ {
        delete(v3, v24)
        v25 += (v25 | (v25 | v25))
        _ = fi26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 bool = (1 <= 907)
        var v29 int8 = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := v4[65535]
    fmt.Println(v2, v24)
    v25 = (v25 % (((^(v25))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint) int8 {
    var v4 uint16 = 988
    v5 := []int{}
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint8 = 207
    var v2 uint64 = 42
    v3 := []int{32767, -1}
    v2 += 317
    v2 = v2
    fmt.Println(v2)
    v6 := GenS1{F0: -167, F1: 2}
    var v7 uint64 = v2
    for fi8 := 0; fi8 < 12; fi8++ {
        if len(v3) > 0 {
            v3[0] = fi8
        }
        _ = len(v3)
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int64 = -1
            if (true && true) {
                v1 = (v1 + v1)
                if true {
                    var v11 int8 = 100
                    v3 = append(v3, (^(fi8)))
                    _ = v11
                } else {
                    var v12 uint64 = v2
                    var v13 uint16 = 0
                    var v14 uint = 903
                    _ = v12
                    _ = v13
                    _ = v14
                }
                fmt.Println(v1, fi9, v7)
                fmt.Println(fi8, v7)
                fmt.Println(fi9, fi8)
            } else {
                v15 := GenS1{F0: 816, F1: (^(v2))}
                if len(v3) > 0 {
                    v3[0] = 369
                }
                _ = v15
            }
            v2 += (v2 ^ v7)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    switch "foo bar" {
    case "a":
        if len(v3) > 0 {
            v3[0] = -837
        }
    default:
        v16 := GenS1{F0: 599, F1: v2}
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v2, v7)
        _ = fi17
    }
    v1 += (v1 & 1)
    if (!(false || ((^(v7)) < (v7 * (^(v2)))))) {
        if len(v3) > 0 {
            v3[0] = 155
        }
        var v18 float32 = 0.1
        _ = v18
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint16 = 601
            _ = fi19
            _ = v20
        }
        _ = len(v3)
        var v21 int32 = -2147483648
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS1{F0: -174, F1: v7}
        _ = fi22
        _ = v23
    }
    var v24 uint64 = v7
    for fi25 := 0; fi25 < 10; fi25++ {
        v26 := GenS2{F0: ("\n\t" + ""), F1: 1, F2: 1}
        if false {
            if true {
                for fi27 := 0; fi27 < 11; fi27++ {
                    v1 += (v1 << ((18446744073709551615) & 63))
                    _ = fi27
                }
                v28 := GenS3{F0: -463}
                fmt.Println(v24, v24)
                _ = v28
            } else {
                var v29 uint8 = (^(v1))
                _ = v29
            }
            var v30 uint = 65535
            _ = v30
        }
        if true {
            var v31 string = "foo bar"
            switch (v31 + v31) {
            case "a":
                v2 = (v2 | v7)
            }
            var v32 float64 = (1.5 + 0.1)
            _ = v31
            _ = v32
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v6
    _ = v7
}

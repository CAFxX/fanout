package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS5 struct { F0 int16; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int32) bool {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 634, F1: 517, F2: 944}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 = p0
        var v7 uint = 42
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return (true && true)
}
func genFn1(p0 uint64, p1 uint8) uint8 {
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(fi8, fi8)
            v9 := []int{2147483647}
            _ = fi8
            _ = v9
        }
    }
    switch ("a" + "") {
    case "a":
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(p1, p1)
            var v11 uint32 = 663
            v12 := GenS2{F0: 1, F1: 130, F2: 9223372036854775807}
            _ = fi10
            _ = v11
            _ = v12
        }
    case "foo bar":
        var v13 int16 = 1
        _ = v13
    case "hello":
        p1 = p1
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 uint64 = 1
    v14 := GenS3{F0: 9223372036854775807, F1: 2147483647, F2: 348}
    for fi15 := 0; fi15 < 4; fi15++ {
        v3 = (v3 << ((443) & 63))
        switch "foo bar" {
        case "foo bar":
            v16 := GenS2{F0: -128, F1: 0, F2: 1}
            _ = v16
        case "":
            v3 -= (^(v3))
        }
        _ = fi15
    }
    var v17 string = "hello"
    v18 := []int{42, 32768, 1}
    switch "" {
    case "a":
        v3 -= (^(v3))
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v3 -= v3
        _ = fi19
    }
    v20 := GenS3{F0: 1, F1: 4294967295, F2: 2}
    if true {
        switch (v17 + v17) {
        case "foo bar":
            var v21 float32 = 1.5
            _ = v21
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS4{F0: 42, F1: 309, F2: -949}
            v24 := GenS2{F0: 127, F1: 1, F2: -1}
            for fi25 := 0; fi25 < 16; fi25++ {
                v1[0] = uint32(v3)
                switch (fi25 * 128) {
                case 9223372036854775807:
                    v17 = "\n\t"
                case 32768:
                    v17 = v17
                }
                if false {
                    switch (^(fi22)) {
                    case 4294967295:
                        if true {
                            var v26 int16 = 1000
                            fmt.Println(fi25, v26, fi25)
                            v27 := GenS5{F0: v26, F1: v3, F2: (v26 * -1)}
                            _ = v26
                            _ = v27
                        }
                    case -9223372036854775808:
                        v17 = v17
                    case -564:
                        if len(v18) > 0 {
                            v18[0] = 0
                        }
                    }
                    var v28 int8 = 127
                    v18 = append(v18, -1)
                    _ = v28
                } else {
                    switch "\n\t" {
                    case "":
                        v29 := GenS2{F0: 100, F1: 9223372036854775807, F2: -553}
                        _ = v29
                    case "a":
                        switch 32768 {
                        case 32767:
                            if false {
                                var v30 uint64 = v3
                                fmt.Println(fi22)
                                _ = v30
                            }
                        case 127:
                            var v31 int16 = 0
                            _ = v31
                        }
                    default:
                        switch v17 {
                        case "foo bar":
                            v32 := GenS6{F0: 158}
                            _ = v32
                        case "":
                            if len(v18) > 0 {
                                v18[0] = (fi25 - fi22)
                            }
                        }
                    }
                }
                if len(v18) > 0 {
                    v18[0] = int(v3)
                }
                switch (^(fi25)) {
                case -1:
                    fmt.Println(v17, fi22)
                default:
                    v3 = 161
                }
                _ = fi25
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}

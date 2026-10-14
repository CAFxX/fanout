package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 int16; F2 int }
func genFn0() int8 {
    var v7 uint32 = 348
    v8 := []int{0, -338}
    _ = v7
    _ = v8
    return 100
}
func genFn1(p0 uint64) int32 {
    v9 := GenS1{F0: 42, F1: p0, F2: 100}
    v10 := []int{}
    switch "foo bar" {
    case "a":
        var v11 string = "hello"
        _ = v11
    case "":
        if (!(127 >= 32768)) {
            v12 := GenS2{F0: 42, F1: false}
            v13 := GenS3{F0: 42}
            _ = v12
            _ = v13
        }
    default:
        p0 = 1
    }
    fmt.Println(p0)
    if true {
        p0 += p0
        fmt.Println(p0)
    }
    _ = p0
    _ = v10
    _ = v9
    return -365
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483648}
    var v3 uint8 = 165
    v4 := []uint32{2147483648}
    var v5 uint16 = 1
    var v6 int8 = 1
    v14 := GenS2{F0: 0, F1: (uint(v5) == 65535)}
    switch 32767 {
    case 907:
        v15 := GenS3{F0: 42}
        _ = v15
    }
    v16 := GenS4{F0: 42}
    v3 = v3
    if ((-1 < 0) && (!false)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v17 := GenS5{F0: 977}
        fmt.Println(v3)
        var v18 uint16 = 0
        _ = v17
        _ = v18
    }
    switch 4294967295 {
    case 127:
        v4 = append(v4, 65535)
    case 1:
        v19 := GenS6{F0: (9223372036854775807 != 42)}
        _ = v19
    case 42:
        var v20 uint = 500
        _ = v20
    default:
        if false {
            switch ("foo bar" + "a") {
            case "":
                switch -802 {
                case -464:
                    v3 -= (v3 | (v3 - 249))
                case 127:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            var v23 uint64 = 683
                            _ = fi22
                            _ = v23
                        }
                        _ = fi21
                    }
                default:
                    if false {
                        fmt.Println(v5, v5)
                        var v24 int64 = 42
                        _ = v24
                    }
                }
            }
            fmt.Println(v6, v3)
        } else {
            switch "foo bar" {
            case "foo bar":
                v5 = v5
            case "hello":
                v5 = v5
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                v4 = append(v4, 2)
                _ = fi25
            }
            v26 := v1["\n\t"]
            _ = v26
        }
    }
    var v27 uint16 = v5
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 uint64 = 42
        v30 := GenS7{F0: -9223372036854775808, F1: -334, F2: fi28}
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint64 = (v29 << ((1) & 63))
            v29 = v32
            _ = len(v4)
            _ = fi31
            _ = v32
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    var v33 int16 = 1
    fmt.Println(v27, v6, v33)
    if false {
        v6 = v6
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 int64 = 1
            _ = fi34
            _ = v35
        }
        v36 := GenS5{F0: 0}
        _ = len(v1)
        _ = v36
    }
    var v37 uint8 = 0
    if ((^(v33)) > (v33 * v33)) {
        v5 = (v5 * v5)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

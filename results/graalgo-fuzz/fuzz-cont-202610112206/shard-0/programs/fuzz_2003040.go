package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint, p2 bool) bool {
    p1 += p1
    p1 -= 255
    if (-885 < 0) {
        var v5 uint8 = 255
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float32 = 1.5
            _ = fi6
            _ = v7
        }
        if (p0 >= (p0 ^ int8(v5))) {
            p2 = true
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: 32767, F1: 42, F2: 42}
                if false {
                    var v10 uint64 = 42
                    _ = v10
                } else {
                    if false {
                        p1 = 0
                    }
                    v11 := GenS2{F0: (fi8 / (((-(fi8))) | 1))}
                    p0 = -90
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            v12 := []int{127, -9223372036854775808}
            _ = v12
        }
        _ = v5
    }
    if p2 {
        var v13 uint = (p1 * p1)
        p0 = p0
        _ = v13
    } else {
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 4294967295
    v3 := make(map[uint32]string)
    var v4 uint = v2
    v14 := GenS1{F0: 2, F1: -917, F2: 42}
    var v15 int = -783
    v16 := []int{4294967295}
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS3{F0: (-1 <= fi17), F1: "foo bar"}
        v4 = (v4 >> ((469) & 63))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := GenS2{F0: (-(v15))}
            if false {
                var v22 bool = true
                var v23 int8 = 82
                var v24 string = ("foo bar" + "\n\t")
                _ = v22
                _ = v23
                _ = v24
            } else {
                fmt.Println(fi20)
                v25 := GenS4{F0: -2147483648}
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        switch "foo bar" {
        case "":
            fmt.Println(v15, fi19)
        case "hello":
            var v26 uint64 = 1
            _ = v26
        case "a":
            var v27 bool = true
            _ = v27
        default:
            v15 = fi19
        }
        v15 += 128
        _ = fi19
    }
    if true {
        var v28 uint = v2
        v28 -= 914
        _ = v28
    } else {
        v4 = (^(v4))
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 int8 = 42
            var v31 string = "hello"
            fmt.Println(v4, v15, v31)
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v33 := GenS4{F0: 1}
        if (1000 >= 42) {
            fmt.Println(fi32, fi32, v2)
        }
        var v34 float32 = ((0.0 - 0.1) / (-2.25 + -2.25))
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int64 = -9223372036854775808
        var v37 int32 = 578
        v2 -= (v2 - (v4 ^ (v4 >> ((v4) & 63))))
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        v39 := GenS5{F0: 42}
        _ = fi38
        _ = v39
    }
    var v40 uint64 = 1
    fmt.Println(v1, v40, v40)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}

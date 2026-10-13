package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32, p1 uint64) uint32 {
    v4 := GenS1{F0: 60, F1: 9223372036854775807}
    var v5 string = (("a" + "a") + ("foo bar" + ("a" + "hello")))
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            p0 = (p0 << ((4294967295) & 63))
            fmt.Println(v5)
            var v7 string = ""
            _ = v7
        } else {
            v8 := []int{-809}
            v9 := GenS2{F0: "a"}
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    for fi10 := 0; fi10 < 9; fi10++ {
        switch "" {
        case "foo bar":
            v11 := GenS3{F0: 1, F1: 4294967295}
            _ = v11
        case "a":
            var v12 int64 = 1
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: 42, F1: (fi10 % ((801) | 1))}
            _ = fi13
            _ = v14
        }
        _ = fi10
    }
    var v15 int64 = 147
    for fi16 := 0; fi16 < 17; fi16++ {
        v15 += (v15 / (((v15 | v15)) | 1))
        v17 := GenS2{F0: "hello"}
        v15 += (v15 % ((v15) | 1))
        _ = fi16
        _ = v17
    }
    v15 = v15
    v18 := []int{664, 9223372036854775807}
    if (!("a" >= (("foo bar" + "foo bar") + ("\n\t" + "")))) {
        v19 := GenS2{F0: ("" + "\n\t")}
        v15 = v15
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS2{F0: "hello"}
                var v23 int64 = int64(fi21)
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 uint16 = 187
            _ = fi20
            _ = v24
        }
        _ = v19
    }
    var v25 int16 = 32767
    var v26 uint32 = 65535
    v26 += genFn0(int32(v15), 969)
    v27 := GenS4{F0: ("\n\t" + "hello"), F1: (!true)}
    if true {
        switch 828 {
        case 592:
            if false {
                fmt.Println(v15, v15, v15)
            } else {
                v28 := GenS4{F0: "a", F1: (408 < 2)}
                _ = v28
            }
        case 32768:
            v29 := GenS3{F0: 1, F1: 18446744073709551615}
            _ = v29
        case 329:
            fmt.Println(v15, v15)
        default:
            v26 = v26
        }
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    v15 += v15
    switch 1 {
    case -9223372036854775808:
        v1[(v26 ^ 2147483648)] = v26
    case 32768:
        v18 = append(v18, -9223372036854775808)
    case 870:
        switch 523 {
        case 1:
            v30 := GenS4{F0: (("hello" + (((("foo bar" + ((("" + "\n\t") + "") + ("" + "\n\t"))) + ((((((("hello" + "\n\t") + ("" + "")) + "foo bar") + "a") + "a") + ("a" + "a")) + ("foo bar" + "foo bar"))) + ("\n\t" + ("a" + "foo bar"))) + "")) + "foo bar"), F1: false}
            _ = v30
        case 127:
            v31 := GenS5{F0: 32767, F1: 1, F2: -2147483648}
            _ = v31
        case 9223372036854775807:
            var v32 uint32 = 0
            _ = v32
        default:
            v33 := GenS6{F0: 1}
            _ = v33
        }
    default:
        if (!false) {
            var v34 string = "\n\t"
            _ = v34
        } else {
            var v35 int16 = (v25 & -4)
            var v36 uint64 = 4294967295
            v26 = v26
            _ = v35
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}

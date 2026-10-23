package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 int64, p1 int16) uint32 {
    var v5 int64 = -1
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (("hello" + "foo bar") + ("hello" + ("hello" + "a"))) {
        case "hello":
            var v7 uint64 = 18446744073709551615
            _ = v7
        case "a":
            v8 := []int{511}
            _ = v8
        default:
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint = 2
                v11 := GenS1{F0: 255}
                v12 := []int{fi6}
                p0 = (-(p0))
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            p0 = (v5 ^ (v5 ^ p0))
            _ = fi13
        }
        _ = fi6
    }
    var v14 float64 = 0.1
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return 779
}
func genFn1() uint32 {
    var v15 bool = (127 < -116)
    v16 := []int{854}
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v15 = v15
            var v19 uint32 = 25
            var v20 uint16 = 1
            v21 := GenS2{F0: 408, F1: 4294967295}
            var v22 int64 = -1
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi17
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        if len(v16) > 0 {
            v16[0] = 127
        }
        v15 = ("a" <= ("" + ("a" + ("foo bar" + "\n\t"))))
        _ = fi23
    }
    _ = v15
    _ = v16
    return genFn0(9223372036854775807, -1)
}

func main() {
    var v1 int32 = 1
    var v2 uint = 65535
    var v3 uint16 = 551
    var v4 int64 = -1
    v24 := []int{945}
    var v25 string = ("\n\t" + (("a" + "") + "\n\t"))
    v26 := GenS3{F0: 42, F1: v2}
    var v27 int16 = 42
    _ = len(v24)
    var v28 int16 = 0
    var v29 int64 = (v4 << ((v2) & 63))
    if (v25 >= v25) {
        fmt.Println(v2)
        fmt.Println(v25, v4)
    } else {
        if len(v24) > 0 {
            v24[0] = 32768
        }
        for fi30 := 0; fi30 < 11; fi30++ {
            if len(v24) > 0 {
                v24[0] = -2147483648
            }
            _ = fi30
        }
    }
    var v31 uint32 = 378
    var v32 int16 = (v28 % ((759) | 1))
    if true {
        var v33 int = -1
        for fi34 := 0; fi34 < 17; fi34++ {
            if (false || (!false)) {
                v28 = (v27 / ((0) | 1))
                fmt.Println(v2, v3, v28)
                var v35 uint16 = 255
                _ = v35
            }
            v36 := GenS1{F0: 42}
            _ = fi34
            _ = v36
        }
        var v37 uint64 = 1
        _ = v33
        _ = v37
    }
    if false {
        v38 := GenS2{F0: 18446744073709551615, F1: v2}
        if len(v24) > 0 {
            v24[0] = 749
        }
        var v39 int64 = 1
        _ = v38
        _ = v39
    }
    _ = len(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}

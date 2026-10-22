package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0() uint64 {
    var v7 uint = 42
    var v8 uint32 = uint32(v7)
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint32 {
    v9 := []int{}
    var v10 string = "hello"
    _ = len(v9)
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 bool = (236 >= 128)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int32 = 42
        v14 = (-(v14))
        _ = fi13
        _ = v14
    }
    _ = v10
    _ = v9
    return 957
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]string)
    v3 := []uint32{4294967295, 65535}
    v4 := make(map[uint32]string)
    var v5 uint64 = 573
    var v6 int32 = -655
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint = 18446744073709551615
        var v17 int64 = 42
        _ = fi15
        _ = v16
        _ = v17
    }
    if ("foo bar" >= "foo bar") {
        var v18 int32 = (v6 % ((1) | 1))
        v5 -= genFn0()
        var v19 int8 = 100
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: 100, F1: 10}
        var v21 uint16 = (v1 * v1)
        _ = v20
        _ = v21
    }
    fmt.Println(v1, v6)
    v5 -= (v5 ^ 283)
    if true {
        v1 = 32767
    } else {
        var v22 uint32 = 1
        _ = v22
    }
    var v23 int8 = 0
    for fi24 := 0; fi24 < 9; fi24++ {
        fmt.Println(v5, v23)
        for fi25 := 0; fi25 < 2; fi25++ {
            if false {
                var v26 uint32 = 2
                _ = v26
            } else {
                var v27 int64 = 9223372036854775807
                var v28 uint64 = (^(v5))
                var v29 int = int(v27)
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi25
        }
        _ = fi24
    }
    v5 += (v5 ^ v5)
    var v30 int64 = int64(v23)
    v5 = 1
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v31 uint64 = 42
        _ = v31
    default:
        v6 = (v6 >> ((0) & 63))
    }
    fmt.Println(v5, v30)
    for fi32 := 0; fi32 < 4; fi32++ {
        v5 -= 42
        fmt.Println(v30)
        if (((^(v1)) != (^(v1))) || true) {
            v2[(fi32 - (^(fi32)))] = ("" + "a")
            var v33 int64 = (v30 % (((v30 + v30)) | 1))
            if (2 > 2) {
                for fi34 := 0; fi34 < 5; fi34++ {
                    v35 := GenS1{F0: v23, F1: 162}
                    var v36 int8 = (v23 % (((-(v23))) | 1))
                    var v37 int = (fi32 + 1)
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            _ = v33
        }
        _ = fi32
    }
    v23 = v23
    for fi38 := 0; fi38 < 18; fi38++ {
        v5 = (v5 % ((96) | 1))
        v39 := v4[2147483648]
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        fmt.Println(v23, fi40, v5)
        v41 := GenS1{F0: v23, F1: 2}
        v42 := GenS1{F0: v23, F1: 2}
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint }
func genFn0() int64 {
    var v4 int16 = -32768
    var v5 bool = ((true || true) || false)
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    v6 := GenS1{F0: ""}
    var v7 uint32 = 4294967295
    if false {
        var v8 int32 = 837
        var v9 uint = 2
        v10 := GenS2{F0: 1, F1: 158}
        _ = v10
        _ = v8
        _ = v9
    } else {
        var v11 float64 = (((1.5 + 1e10) - 1e10) / (float64(1.5) - 1e10))
        _ = v11
    }
    if true {
        var v12 uint64 = 4294967295
        _ = v12
    }
    var v13 uint32 = 42
    v13 -= (v13 >> ((1) & 63))
    if true {
        var v14 float32 = (1.5 / float32(1e10))
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                v7 -= (^(v7))
                var v17 int32 = -1
                var v18 uint = 1
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                switch (("foo bar" + "\n\t") + "\n\t") {
                case "hello":
                    v20 := GenS3{F0: 2, F1: 32767}
                    _ = v20
                }
                var v21 uint = uint(v7)
                _ = fi19
                _ = v21
            }
            v14 = 0.1
            _ = fi15
        }
        if true {
            var v22 bool = (1 < 0)
            v23 := GenS3{F0: 0, F1: 1}
            _ = v22
            _ = v23
        }
        _ = v14
    } else {
        v1 = v1
        var v24 string = ("foo bar" + "\n\t")
        v1 = v1
        _ = v24
    }
    if ((v7 >> ((uint(v7)) & 63)) <= 909) {
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 uint32 = 643
            v26 += v26
            v27 := GenS3{F0: 2, F1: 65535}
            _ = fi25
            _ = v26
            _ = v27
        }
        v28 := GenS1{F0: ""}
        switch "a" {
        case "":
            var v29 uint8 = 255
            _ = v29
        case "hello":
            for fi30 := 0; fi30 < 18; fi30++ {
                fmt.Println(v7, v1)
                switch (("" + ("" + "hello")) + "a") {
                case "a":
                    v1 += v1
                case "\n\t":
                    var v31 uint64 = 1
                    _ = v31
                case "hello":
                    for fi32 := 0; fi32 < 13; fi32++ {
                        if false {
                            v33 := GenS4{F0: -1, F1: (113 >= 65535), F2: 18446744073709551615}
                            var v34 float32 = -2.25
                            v1 = (^(v1))
                            _ = v33
                            _ = v34
                        } else {
                            v35 := GenS4{F0: genFn0(), F1: (213 >= 127), F2: 1}
                            _ = v35
                        }
                        _ = fi32
                    }
                default:
                    var v36 int32 = int32(v13)
                    _ = v36
                }
                var v37 int16 = 369
                _ = fi30
                _ = v37
            }
        }
        _ = v28
    }
    var v38 int64 = genFn0()
    var v39 int64 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v6
    _ = v7
}

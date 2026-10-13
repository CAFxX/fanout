package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint8; F1 bool; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 uint = 4294967295
    var v4 int32 = 1
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int64 = 1
        fmt.Println(v4)
        if true {
            _ = len(v2)
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 int16 = 0
                v6 += int64(v8)
                if false {
                    var v9 int32 = 2147483647
                    if false {
                        var v10 bool = (!true)
                        v11 := GenS1{F0: 706}
                        _ = v10
                        _ = v11
                    } else {
                        v3 = (^(v3))
                    }
                    _ = v9
                } else {
                    var v12 uint16 = 973
                    var v13 float32 = 0.1
                    var v14 int8 = 127
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi7
                _ = v8
            }
            v15 := GenS1{F0: 0}
            _ = v15
        }
        _ = fi5
        _ = v6
    }
    v16 := []int{-9223372036854775808, 501}
    var v17 int16 = 1
    v17 = (v17 ^ (-(v17)))
    v18 := GenS1{F0: int16(v3)}
    v19 := GenS2{F0: 0}
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint64 = 2
            fmt.Println(fi21, v3, v17)
            fmt.Println(v4, v4, v3)
            fmt.Println(fi21, fi21, v22)
            _ = fi21
            _ = v22
        }
        if (!true) {
            v4 = v4
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        } else {
            v23 := GenS3{F0: ("" + ""), F1: (32767 > 12)}
            var v24 int32 = (v4 - 476)
            delete(v1, "foo bar")
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            if (102 < 1) {
                var v26 int16 = (v17 >> ((255) & 63))
                fmt.Println(v4, v17, fi25)
                _ = v26
            } else {
                v4 = v4
                v27 := GenS2{F0: 1}
                _ = len(v2)
                _ = v27
            }
            _ = fi25
        }
        _ = fi20
    }
    v28 := GenS4{F0: 2, F1: false, F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}

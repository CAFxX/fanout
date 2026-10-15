package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]uint32)
    var v3 int64 = 1
    v4 := make(map[string]int)
    if (false && true) {
        v5 := []int{471, 459, -422}
        var v6 float32 = 1.5
        var v7 bool = true
        _ = v5
        _ = v6
        _ = v7
    } else {
        if (false || (-1 > 57)) {
            var v8 string = "hello"
            _ = v8
        } else {
            var v9 string = "a"
            switch ("foo bar" + "") {
            case "a":
                v10 := GenS1{F0: v9, F1: 663, F2: ""}
                _ = v10
            case "":
                v3 = 42
            }
            fmt.Println(v1, v9, v1)
            _ = v9
        }
    }
    v11 := GenS2{F0: "", F1: v3}
    if (-953 > 42) {
        if (v1 <= 32767) {
            v3 -= v3
            if false {
                var v12 int32 = -877
                _ = v12
            } else {
                var v13 int = 128
                var v14 int8 = 127
                _ = v13
                _ = v14
            }
            var v15 float32 = -2.25
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 12; fi16++ {
            if (198 != 4294967295) {
                v17 := GenS3{F0: 100, F1: 255}
                v3 = (v3 - 1)
                _ = v17
            } else {
                v18 := GenS4{F0: v3, F1: 4294967295, F2: 2147483647}
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 uint64 = 739
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            _ = fi16
        }
    }
    v3 = (v3 >> ((2) & 63))
    v3 -= (v3 << ((643) & 63))
    v21 := GenS5{F0: false}
    v22 := GenS6{F0: 65535}
    var v23 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}

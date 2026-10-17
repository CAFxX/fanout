package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int32; F1 int; F2 uint8 }
func genFn0(p0 uint8, p1 uint64, p2 uint) uint8 {
    p1 -= p1
    switch -927 {
    case 792:
        v5 := GenS1{F0: 42}
        _ = v5
    case -1:
        p0 = 111
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 | (p0 ^ p0))
}
func genFn1(p0 uint8) int16 {
    var v6 float64 = -2.25
    if false {
        if false {
            p0 = p0
        }
        fmt.Println(v6, p0)
    }
    var v7 string = ("hello" + ("hello" + ""))
    _ = p0
    _ = v6
    _ = v7
    return -947
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 int64 = -978
    var v4 uint64 = 1
    var v8 int32 = int32(v3)
    if (!true) {
        v9 := []int{}
        _ = v9
    }
    v3 = 565
    var v10 string = ""
    if ((!(1 < 2147483648)) || true) {
        var v11 int8 = -128
        _ = v11
    }
    v3 -= (v3 ^ 0)
    v12 := []int{0, 810, 32768}
    switch (("a" + v10) + "\n\t") {
    case "":
        v13 := v2[32767]
        _ = v13
    }
    v8 = v8
    var v14 int32 = (-(v8))
    v15 := GenS2{F0: int32(v4), F1: 32768}
    _ = len(v2)
    fmt.Println(v8, v14)
    if false {
        switch -1 {
        case 32767:
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int32 = v8
                _ = fi16
                _ = v17
            }
        case 127:
            v8 += (^(v14))
        case 562:
            fmt.Println(v3, v10, v10)
        default:
            var v18 float64 = 1.5
            _ = v18
        }
        v19 := GenS3{F0: (v8 >> ((250) & 63))}
        _ = v19
    } else {
        v4 = (v4 >> ((0) & 63))
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int8 = -94
            v22 := GenS4{F0: 2}
            var v23 uint = 966
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        delete(v1, 2147483648)
    }
    if (0 >= 42) {
        v4 = 0
        if ((^(v3)) <= 989) {
            var v24 uint32 = 0
            switch -998 {
            case 32767:
                v25 := GenS5{F0: (!(!(!(((1 == -32768) || false) || true))))}
                _ = v25
            case 2147483647:
                v26 := GenS6{F0: "hello", F1: genFn1(228)}
                _ = v26
            default:
                _ = len(v1)
            }
            v3 = 1
            _ = v24
        } else {
            v10 = v10
            for fi27 := 0; fi27 < 13; fi27++ {
                for fi28 := 0; fi28 < 14; fi28++ {
                    v14 -= (v14 ^ -1)
                    _ = len(v1)
                    _ = fi28
                }
                v29 := v1[2]
                for fi30 := 0; fi30 < 6; fi30++ {
                    fmt.Println(v3)
                    _ = fi30
                }
                _ = fi27
                _ = v29
            }
            if (3.14159 >= ((0.1 - 3.14159) - 0.1)) {
                for fi31 := 0; fi31 < 14; fi31++ {
                    v10 = (("a" + (v10 + v10)) + v10)
                    _ = fi31
                }
            }
        }
        v8 = (v8 ^ (v14 / ((v8) | 1)))
    } else {
        if false {
            switch 42 {
            case 9223372036854775807:
                var v32 float32 = float32(1.5)
                _ = v32
            case 879:
                v12 = append(v12, 32768)
            default:
                v4 = 2
            }
            v33 := GenS3{F0: (v8 & v8)}
            switch (v10 + v10) {
            case "a":
                v3 = v3
            case "hello":
                v34 := GenS7{F0: (v8 + (v8 >> ((561) & 63))), F1: 4294967295, F2: 42}
                _ = v34
            }
            _ = v33
        } else {
            v8 = v8
            v12 = append(v12, 4294967295)
            v3 = v3
        }
        var v35 uint = 2
        v35 = (v35 & 4294967295)
        _ = v35
    }
    var v36 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v8
}

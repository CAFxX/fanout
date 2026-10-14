package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int16; F1 string; F2 uint }
func genFn0(p0 uint8) int32 {
    v7 := GenS1{F0: 65535, F1: 2147483647, F2: 1}
    v8 := GenS1{F0: 4294967295, F1: 2, F2: 4294967295}
    var v9 uint16 = 42
    if false {
        switch "hello" {
        case "a":
            switch ("\n\t" + "") {
            case "\n\t":
                var v10 float32 = -2.25
                _ = v10
            }
        default:
            v11 := GenS2{F0: v9, F1: 860}
            _ = v11
        }
        var v12 uint32 = 65535
        _ = v12
    } else {
        switch "hello" {
        case "foo bar":
            switch -2147483648 {
            case -183:
                v13 := GenS3{F0: 42, F1: 42, F2: 0}
                _ = v13
            case -210:
                v14 := []int{}
                _ = v14
            }
        case "a":
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS4{F0: (-(fi15)), F1: 1}
                p0 = p0
                _ = fi15
                _ = v16
            }
        }
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 bool, p1 int16, p2 int32) int16 {
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS4{F0: (^(fi17)), F1: 42}
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 1
    var v20 float64 = (float64(((0.0 / (1e10 / 0.1)) * float32(0.1))) - (0.1 - (-2.25 / (3.14159 + float64(((3.14159 - float32((-2.25 - (1e10 - 0.1)))) + 1.5))))))
    v21 := GenS3{F0: (^(v19)), F1: 0, F2: p2}
    p0 = (true && p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return p1
}

func main() {
    v1 := []int{-539}
    var v2 int16 = 1
    var v3 uint8 = 42
    v4 := make(map[string]uint32)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[uint32]string)
    v3 = 0
    var v22 float32 = -2.25
    var v23 int32 = 1
    v22 = v22
    var v24 uint16 = 2
    for fi25 := 0; fi25 < 8; fi25++ {
        v23 -= 0
        if true {
            v2 -= (v2 ^ (v2 >> ((18446744073709551615) & 63)))
        }
        var v26 string = "a"
        _ = fi25
        _ = v26
    }
    v27 := GenS5{F0: -152, F1: "", F2: 42}
    for fi28 := 0; fi28 < 11; fi28++ {
        v23 += int32(v24)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
        _ = fi28
    }
    var v29 int16 = (^(v2))
    v23 -= (v23 * (v23 + genFn0((v3 / ((2) | 1)))))
    if (v22 < ((v22 / (-2.25 - (((float32(0.1) * (1e10 + v22)) + 3.14159) * v22))) * v22)) {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int64 = int64(v3)
            var v32 int32 = (v23 >> ((2) & 63))
            _ = fi30
            _ = v31
            _ = v32
        }
    } else {
        fmt.Println(v23)
        v33 := GenS1{F0: 261, F1: 65535, F2: v5}
        _ = v33
    }
    var v34 uint64 = v5
    var v35 int = 819
    if false {
        var v36 string = "hello"
        fmt.Println(v3, v2, v5)
        _ = v36
    } else {
        v34 = (v5 & (v34 - (^(v34))))
    }
    fmt.Println(v35, v2, v24)
    if (-109 >= -84) {
        var v37 uint = 42
        v6[2] = "foo bar"
        var v38 uint32 = uint32(v37)
        _ = v37
        _ = v38
    }
    var v39 float32 = float32((-2.25 + 1e10))
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 int = (v35 | (v35 / ((42) | 1)))
        _ = fi40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

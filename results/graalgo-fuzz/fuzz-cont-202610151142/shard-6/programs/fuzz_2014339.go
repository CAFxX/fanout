package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 string) int16 {
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: 230, F1: 128, F2: 42}
            v8 := []int{}
            var v9 int32 = 42
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        p0 = "a"
        var v10 int = (fi5 ^ fi5)
        _ = fi5
        _ = v10
    }
    switch p0 {
    case "\n\t":
        var v11 float64 = 0.1
        _ = v11
    case "":
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: 518, F1: 42, F2: 611}
            if false {
                v14 := GenS2{F0: 11, F1: -323, F2: -1}
                var v15 int16 = 1000
                var v16 float64 = 3.14159
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]string)
    var v3 uint16 = 213
    v4 := make(map[int]uint32)
    switch 9223372036854775807 {
    case -553:
        _ = len(v4)
    case 32768:
        v17 := []int{209}
        _ = v17
    default:
        v18 := GenS3{F0: 42, F1: 998, F2: -57}
        _ = v18
    }
    v3 += 2
    if true {
        v19 := []int{}
        v19 = append(v19, 0)
        v3 += (v3 % ((32767) | 1))
        _ = v19
    }
    var v20 float64 = -2.25
    v21 := GenS1{F0: v1, F1: 1, F2: -2147483648}
    fmt.Println(v1, v1)
    for fi22 := 0; fi22 < 17; fi22++ {
        if ((fi22 * 0) == fi22) {
            var v23 uint8 = 255
            v24 := []int{}
            v1 = 742
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v3 = 32767
        _ = fi25
    }
    switch "a" {
    case "\n\t":
        v4[-940] = 2147483648
    case "hello":
        var v26 int64 = 281
        _ = v26
    case "":
        var v27 uint32 = 2147483648
        _ = v27
    default:
        v3 = (v3 ^ (v3 ^ 1))
    }
    v20 = (((((v20 / v20) * 1e10) + (v20 - 1.5)) / v20) / v20)
    fmt.Println(v1, v3, v3)
    fmt.Println(v3, v20, v3)
    var v28 bool = ((true || true) && (false && false))
    delete(v2, ("\n\t" + "a"))
    switch "foo bar" {
    case "a":
        switch 1 {
        case 1:
            var v29 uint = 1
            _ = v29
        }
    default:
        for fi30 := 0; fi30 < 2; fi30++ {
            for fi31 := 0; fi31 < 8; fi31++ {
                v28 = (v28 && (4294967295 < 2147483648))
                fmt.Println(v3, v20, fi30)
                _ = fi31
            }
            if v28 {
                var v32 uint8 = 42
                v33 := GenS2{F0: 13, F1: int32(fi30), F2: 0}
                _ = v32
                _ = v33
            } else {
                _ = len(v2)
                for fi34 := 0; fi34 < 18; fi34++ {
                    v20 = (v20 / (1e10 + v20))
                    fmt.Println(v20, v28)
                    _ = fi34
                }
            }
            _ = fi30
        }
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        fmt.Println(v3, v20)
        var v36 int64 = 0
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
}

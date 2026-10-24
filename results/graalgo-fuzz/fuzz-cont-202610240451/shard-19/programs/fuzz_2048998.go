package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 string, p1 int8) uint {
    p0 = ""
    var v5 int = 1
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 string = "a"
    if (genFn0((v4 + v4), -1) <= 42) {
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: "hello", F1: -1, F2: 861}
            v8 := GenS2{F0: 4294967295, F1: 2, F2: "foo bar"}
            delete(v2, fi6)
            _ = fi6
            _ = v7
            _ = v8
        }
        if (((0.0 + float64((3.14159 * ((((float32(1.5) / 1.0) - ((-2.25 / 1.0) + (1e10 + float32(0.1)))) + -2.25) + 1.5)))) == float64(0.0)) && true) {
            v4 = ("hello" + v4)
            v4 = v4
        }
        _ = len(v2)
    } else {
        if (2 == 834) {
            var v9 int32 = -312
            v10 := GenS1{F0: v4, F1: -766, F2: 638}
            v11 := v1[-9223372036854775808]
            _ = v10
            _ = v11
            _ = v9
        }
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        delete(v3, "\n\t")
        _ = fi12
    }
    v4 = ((((v4 + "a") + "\n\t") + "hello") + v4)
    v13 := GenS3{F0: 926, F1: -804, F2: 640}
    if false {
        v1[714] = 4294967295
        v4 = ((v4 + "a") + v4)
        v1[32767] = -1
    }
    var v14 uint16 = 42
    v15 := []int{825, 128, 32768}
    var v16 uint64 = 1
    var v17 int64 = -191
    var v18 int8 = 83
    v19 := GenS3{F0: 1000, F1: int16(v18), F2: 966}
    var v20 uint8 = 40
    fmt.Println(v16, v18)
    v21 := v3[("hello" + "")]
    if ((true && v21) && (75 == -1)) {
        fmt.Println(v14, v4, v17)
    } else {
        var v22 float64 = -2.25
        var v23 bool = (!v21)
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"", "a", "foo bar"}
    v3 := []int{1, 32767, 1, 4294967295}
    v4 := []uint32{799}
    v5 := []int{127, -1, -307, -858}
    var v6 uint32 = 2147483648
    var v7 int = 32767
    if (3.14159 == (0.0 / ((0.0 + 1.5) - 1e10))) {
        v8 := v1[v6]
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int16 = 806
            var v11 uint64 = 48
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 42, F1: ("\n\t" + "foo bar"), F2: v8}
        _ = v12
        _ = v8
    } else {
        v13 := GenS1{F0: 42, F1: "a", F2: (-(v7))}
        var v14 uint16 = 42
        _ = v13
        _ = v14
    }
    v7 += (v7 - (^(v7)))
    var v15 int16 = 32767
    if false {
        v16 := GenS2{F0: -2147483648, F1: 495, F2: 1}
        if true {
            fmt.Println(v6, v15)
            v1[(v6 % (((^(v6))) | 1))] = v7
            if (!true) {
                var v17 uint = 255
                var v18 int32 = -1
                var v19 bool = true
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 uint32 = (^(v6))
                _ = v20
            }
        } else {
            fmt.Println(v7)
            v7 = 9223372036854775807
            fmt.Println(v15, v7, v6)
        }
        if len(v2) > 0 {
            v2[0] = ("a" + "")
        }
        _ = v16
    }
    var v21 float32 = (1.5 + (((0.1 + float32(3.14159)) * float32(3.14159)) + ((((3.14159 / 1.5) * -2.25) + (-2.25 / 1.0)) / 1e10)))
    var v22 string = (("hello" + "foo bar") + (("foo bar" + "") + "a"))
    var v23 float32 = ((float32(0.0) - ((1.5 * (((0.0 / v21) + 1.5) + v21)) * v21)) / (v21 + v21))
    v15 = -1
    v24 := GenS3{F0: 1, F1: 42, F2: 42}
    if len(v5) > 0 {
        v5[0] = (v7 / (((v7 * v7)) | 1))
    }
    var v25 uint32 = 2
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint32 = (v25 * v6)
        var v28 int64 = -1
        v6 -= (v6 | (v25 + 1))
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS3{F0: 9223372036854775807, F1: 42, F2: 352}
    delete(v1, (v6 >> ((4294967295) & 63)))
    v6 = v25
    for fi30 := 0; fi30 < 13; fi30++ {
        v22 = v22
        switch (v22 + (v22 + ("hello" + v22))) {
        case "foo bar":
            v31 := v1[v6]
            _ = v31
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

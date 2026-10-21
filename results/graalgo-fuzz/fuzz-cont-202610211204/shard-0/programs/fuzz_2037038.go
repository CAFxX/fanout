package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 bool }
func genFn0() int32 {
    switch -893 {
    case 32767:
        if true {
            v4 := GenS1{F0: 100}
            var v5 float32 = 0.1
            if (32767 > 0) {
                var v6 uint32 = 42
                v6 -= (^(v6))
                v5 = 1e10
                _ = v6
            }
            _ = v4
            _ = v5
        }
    case 32768:
        var v7 float64 = (-2.25 * (1e10 + 0.0))
        _ = v7
    case 4294967295:
        v8 := []int{9223372036854775807, 32768}
        _ = v8
    }
    var v9 int32 = -742
    v10 := GenS2{F0: "foo bar", F1: -867, F2: 1}
    _ = v10
    _ = v9
    return (v9 >> ((42) & 63))
}

func main() {
    var v1 uint16 = 1
    var v2 float64 = float64(((0.0 + ((float32(-2.25) * 0.0) - 1.5)) - 0.0))
    var v3 int64 = 0
    fmt.Println(v3, v1)
    var v11 uint64 = 1
    v11 -= v11
    v12 := []int{-1, 9223372036854775807, 9223372036854775807}
    if true {
        if len(v12) > 0 {
            v12[0] = 127
        }
        v13 := GenS3{F0: false}
        var v14 string = ("a" + "\n\t")
        _ = v13
        _ = v14
    }
    v1 = (^(v1))
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch (("" + "hello") + "\n\t") {
            case "\n\t":
                v12 = append(v12, (-(fi15)))
            case "":
                var v17 uint8 = 255
                _ = v17
            case "foo bar":
                var v18 int64 = v3
                _ = v18
            }
            var v19 int = fi16
            _ = fi16
            _ = v19
        }
        v1 -= v1
        _ = fi15
    }
    var v20 uint32 = 0
    _ = len(v12)
    v3 = 648
    var v21 uint16 = (v1 * v1)
    v22 := GenS2{F0: "foo bar", F1: 841, F2: -128}
    var v23 string = "\n\t"
    var v24 uint16 = v21
    v20 += v20
    var v25 int32 = -1
    if (!((false || (!(0 <= 127))) && true)) {
        fmt.Println(v24)
    } else {
        var v26 int32 = -1
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 float32 = 0.0
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}

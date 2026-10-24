package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 string = ("" + "a")
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := []string{v1, "foo bar", (v1 + "a")}
    var v5 uint64 = 42
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 489, F1: (v5 ^ 1), F2: v5}
    var v8 int64 = 0
    var v9 bool = (true && (!true))
    fmt.Println(v9, v8)
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: 1, F1: 42}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: (v5 >> ((255) & 63)), F2: v5}
    v13 := GenS3{F0: 128}
    fmt.Println(v1, v8, v8)
    var v14 uint32 = 818
    for fi15 := 0; fi15 < 16; fi15++ {
        v9 = ((fi15 << ((18446744073709551615) & 63)) > fi15)
        _ = fi15
    }
    var v16 bool = ((v9 || v9) && v9)
    var v17 int8 = 68
    if ((v8 | (^(v8))) > (v8 >> ((0) & 63))) {
        if false {
            if true {
                var v18 uint = 42
                var v19 bool = true
                _ = v18
                _ = v19
            }
        } else {
            v16 = v9
            v5 = v5
            for fi20 := 0; fi20 < 4; fi20++ {
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 int8 = v17
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 int64 = (v8 % (((^(v8))) | 1))
                    _ = fi23
                    _ = v24
                }
                _ = fi20
            }
        }
        v25 := v2[-634]
        _ = v25
    } else {
        fmt.Println(v8, v14)
        var v26 int8 = v17
        _ = v26
    }
    var v27 float64 = float64(-2.25)
    if (!v9) {
        switch -391 {
        case 42:
            v28 := GenS2{F0: (v14 / ((65535) | 1)), F1: 32767}
            _ = v28
        case 127:
            var v29 int64 = (v8 % ((0) | 1))
            _ = v29
        }
        var v30 int16 = -876
        v31 := GenS1{F0: (v8 * (v8 * v8)), F1: 215, F2: v5}
        _ = v30
        _ = v31
    } else {
        if (2 != 255) {
            var v32 uint16 = 32767
            v16 = true
            if true {
                fmt.Println(v27)
            }
            _ = v32
        } else {
            var v33 float32 = float32(0.0)
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := GenS4{F0: 128, F1: 51}
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            switch 128 {
            case -1:
                v5 = v5
            case 128:
                fmt.Println(v14, v8, v17)
                v16 = true
            }
            v9 = v9
            _ = fi37
        }
        if (!(true && (-1 == -918))) {
            v5 = v5
            var v38 int64 = 0
            v38 += (v8 * 9223372036854775807)
            _ = v38
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

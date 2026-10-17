package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 bool }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]bool)
    var v3 string = ""
    var v4 uint = 377
    v5 := make(map[string]uint32)
    var v6 float32 = 1.5
    if false {
        v5[v3] = 0
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: (fi7 >> (((v4 & 42)) & 63)), F1: 495, F2: -1}
            v9 := GenS1{F0: fi7, F1: 718, F2: 1}
            switch ((((v3 + "foo bar") + ("" + v3)) + v3) + v3) {
            case "foo bar":
                v3 = (("foo bar" + "foo bar") + (v3 + ((v3 + "foo bar") + v3)))
            case "":
                for fi10 := 0; fi10 < 9; fi10++ {
                    v4 += (^(v4))
                    v11 := GenS1{F0: (fi10 >> (((v4 * 255)) & 63)), F1: 320, F2: -1}
                    v4 += (v4 | v4)
                    _ = fi10
                    _ = v11
                }
            default:
                fmt.Println(v4, v6, v4)
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    delete(v5, v3)
    var v12 int = 32768
    for fi13 := 0; fi13 < 12; fi13++ {
        switch v12 {
        case 127:
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int64 = 42
                _ = fi14
                _ = v15
            }
        default:
            if (float64(((v6 / v6) - (float32(1e10) - v6))) <= (1e10 + 1.5)) {
                var v16 float64 = (((float64((1e10 * float32(1.5))) * 0.1) * 0.0) - (1e10 / 1.5))
                fmt.Println(v3, fi13, v1)
                _ = v16
            } else {
                delete(v2, ((v3 + "foo bar") + ("foo bar" + ("" + "a"))))
            }
        }
        if false {
            var v17 float32 = 1.5
            v17 = (0.0 / (v6 - (v17 + (0.0 / v6))))
            fmt.Println(v3)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                if (1 == 42) {
                    fmt.Println(fi13, v4)
                } else {
                    v19 := GenS1{F0: v12, F1: 256, F2: 9223372036854775807}
                    v3 = "hello"
                    var v20 uint64 = 376
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
        }
        _ = fi13
    }
    var v21 uint8 = 173
    var v22 string = v3
    v23 := v2[(v3 + v22)]
    switch (v22 + v3) {
    case "":
        for fi24 := 0; fi24 < 1; fi24++ {
            fmt.Println(fi24, v22)
            _ = fi24
        }
    case "hello":
        fmt.Println(v1, v3)
    }
    v25 := []int{1, v12}
    fmt.Println(v1, v1, v21)
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 float64 = 0.0
        fmt.Println(v12)
        _ = fi26
        _ = v27
    }
    fmt.Println(v6, v1)
    v28 := GenS2{F0: false, F1: 2147483647, F2: v23}
    v29 := GenS2{F0: v23, F1: 2147483647, F2: ((0 < 42) && true)}
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    v4 = (v4 * 42)
    if false {
        v25 = append(v25, (v12 - (v12 - (v12 * (^(v12))))))
    } else {
        var v30 int64 = 9223372036854775807
        switch 32767 {
        case 824:
            var v31 int8 = 127
            _ = v31
        case 0:
            var v32 int16 = -507
            _ = v32
        case 4294967295:
            v5[(v22 + v22)] = 2
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

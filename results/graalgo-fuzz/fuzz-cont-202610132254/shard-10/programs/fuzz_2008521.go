package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int64) uint8 {
    if false {
        p0 += 51
        switch "" {
        case "\n\t":
            v4 := GenS1{F0: 1}
            _ = v4
        case "":
            if (2147483647 >= -2147483648) {
                v5 := GenS1{F0: 32767}
                var v6 uint16 = 1
                _ = v5
                _ = v6
            }
        }
        var v7 bool = (-18 <= 0)
        _ = v7
    } else {
        if false {
            var v8 uint32 = 2
            v9 := []int{128, 152, 4294967295}
            v10 := GenS2{F0: false}
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := GenS1{F0: -1}
        _ = v11
    }
    var v12 int64 = (p0 / ((p0) | 1))
    var v13 float64 = float64(0.0)
    v13 -= v13
    v12 += v12
    _ = p0
    _ = v12
    _ = v13
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{("\n\t" + (("foo bar" + "\n\t") + (("a" + "foo bar") + "a"))), "\n\t"}
    var v3 uint16 = 65535
    fmt.Println(v3)
    v3 = v3
    v3 -= v3
    v3 = 98
    v3 = v3
    var v14 int32 = 1
    if (!true) {
        v14 -= v14
        if ("hello" >= (("\n\t" + "foo bar") + "a")) {
            fmt.Println(v14)
            for fi15 := 0; fi15 < 7; fi15++ {
                switch "" {
                case "a":
                    var v16 string = (("" + "hello") + "a")
                    _ = v16
                default:
                    var v17 bool = (527 != 42)
                    _ = v17
                }
                if (float32(1.5) <= 1.5) {
                    var v18 int = 677
                    v19 := GenS3{F0: (939 != 1)}
                    var v20 uint = 65
                    _ = v18
                    _ = v19
                    _ = v20
                }
                fmt.Println(fi15, v14, fi15)
                _ = fi15
            }
            switch "foo bar" {
            case "hello":
                v21 := GenS1{F0: -641}
                _ = v21
            case "a":
                var v22 bool = (!(4294967295 >= 42))
                _ = v22
            }
        } else {
            if (v3 >= 255) {
                switch 9223372036854775807 {
                case 32768:
                    fmt.Println(v14, v3)
                }
                v23 := GenS1{F0: 42}
                _ = v23
            }
        }
        v24 := GenS2{F0: (!(0 >= -949))}
        _ = v24
    }
    var v25 uint8 = 127
    var v26 float64 = 1e10
    for fi27 := 0; fi27 < 8; fi27++ {
        if false {
            v14 = v14
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 uint8 = v25
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    if false {
        if (512 <= 135) {
            var v30 int32 = v14
            _ = v30
        } else {
            v31 := GenS4{F0: v25, F1: (v3 ^ v3)}
            v26 -= v26
            _ = v31
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v32 float32 = -2.25
    v32 += v32
    v33 := GenS1{F0: 0}
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            if (0 >= 255) {
                _ = len(v1)
                fmt.Println(fi35, fi35, v3)
                _ = len(v1)
                var v36 int32 = (v14 | v14)
                _ = v36
            } else {
                var v37 uint16 = (^(v3))
                var v38 float32 = ((0.1 * 1.5) - 0.1)
                _ = v37
                _ = v38
            }
            _ = fi35
        }
        v25 = (^(v25))
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 uint32, p1 int32) int8 {
    if (128 > 42) {
        switch "a" {
        case "foo bar":
            for fi5 := 0; fi5 < 2; fi5++ {
                var v6 int = (fi5 / ((-1) | 1))
                _ = fi5
                _ = v6
            }
        case "\n\t":
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 uint32 = 648
                p1 = p1
                _ = fi7
                _ = v8
            }
        }
        p0 = p0
    } else {
        fmt.Println(p1, p0)
        switch 4294967295 {
        case 347:
            var v9 uint8 = 255
            _ = v9
        case 32767:
            var v10 int8 = 100
            _ = v10
        }
        var v11 int16 = 0
        _ = v11
    }
    var v12 int64 = -609
    var v13 int32 = p1
    if (float32(float64(float32(((0.1 / 1.0) / float64((0.0 - float32(((0.1 - 0.0) / (0.1 / 0.1))))))))) <= 0.1) {
        fmt.Println(p0, p1)
        var v14 uint64 = 2
        _ = v14
    } else {
        v15 := GenS1{F0: true, F1: 659}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -1
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]bool)
    var v3 uint = 651
    var v4 string = "hello"
    if (!(((0.0 + 0.0) + 1e10) > 3.14159)) {
        v16 := GenS1{F0: false, F1: 42}
        _ = v16
    } else {
        v17 := []int{2147483647, 32768}
        _ = v17
    }
    fmt.Println(v4, v4, v1)
    if (false || false) {
        var v18 int32 = 42
        _ = v18
    }
    if (-2.25 > (((3.14159 / 1.5) / 3.14159) + 1.5)) {
        var v19 float64 = 3.14159
        v3 += 255
        _ = v19
    } else {
        v20 := GenS2{F0: 0}
        for fi21 := 0; fi21 < 7; fi21++ {
            fmt.Println(v4, v1, fi21)
            delete(v2, 1)
            _ = fi21
        }
        var v22 uint32 = (^(v1))
        _ = v20
        _ = v22
    }
    if true {
        v23 := GenS3{F0: 32767, F1: ((v4 + v4) >= v4)}
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS3{F0: 169, F1: (4294967295 > 4294967295)}
            delete(v2, (fi24 * (fi24 & 32768)))
            _ = fi24
            _ = v25
        }
        v26 := []int{289, 664, 879}
        _ = v23
        _ = v26
    }
    v1 = v1
    var v27 float32 = 0.1
    var v28 uint64 = 4294967295
    switch v4 {
    case "foo bar":
        var v29 int64 = -9223372036854775808
        _ = v29
    }
    fmt.Println(v1)
    v28 += v28
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}

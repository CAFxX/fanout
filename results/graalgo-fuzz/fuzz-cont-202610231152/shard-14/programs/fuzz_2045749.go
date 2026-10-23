package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int16, p1 int16) int32 {
    p0 = p0
    p1 += (p1 - p0)
    v5 := GenS1{F0: 514, F1: "hello"}
    p1 += -32768
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 int8) bool {
    var v6 float32 = 1.5
    var v7 int8 = (-(p1))
    var v8 int8 = p1
    v9 := GenS2{F0: 65535, F1: 32768}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 float32 = (3.14159 + 3.14159)
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 string = "hello"
        var v12 uint = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        if false {
            var v14 string = "hello"
            for fi15 := 0; fi15 < 9; fi15++ {
                fmt.Println(fi13)
                var v16 float32 = v4
                v17 := GenS3{F0: 0, F1: 877}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(fi18, v4)
            v19 := GenS4{F0: 0}
            fmt.Println(v4)
            v20 := GenS5{F0: 42}
            v21 := []int{331, 9223372036854775807, (fi13 | (fi13 % (((fi18 + fi13)) | 1))), (fi18 / ((fi18) | 1))}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi13
    }
    v4 += (((v4 * v4) / (v4 + (v4 / (v4 - float32(0.1))))) - 3.14159)
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(fi23)
            _ = len(v3)
            _ = fi23
        }
        _ = fi22
    }
    v24 := GenS4{F0: genFn0(1, 1000)}
    for fi25 := 0; fi25 < 14; fi25++ {
        v4 -= (v4 * 1e10)
        switch "a" {
        case "hello":
            var v26 uint = 455
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS2{F0: 255, F1: fi25}
            _ = fi27
            _ = v28
        }
        _ = fi25
    }
    if (!true) {
        fmt.Println(v4)
        _ = len(v2)
    } else {
        v29 := GenS6{F0: 127, F1: 925}
        var v30 uint16 = 42
        var v31 int = -1
        _ = v29
        _ = v30
        _ = v31
    }
    v32 := []int{127, -2147483648, 4294967295}
    for fi33 := 0; fi33 < 13; fi33++ {
        v4 += v4
        _ = fi33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 uint16 = 654
        _ = fi34
        _ = v35
    }
    var v36 int64 = -9223372036854775808
    if (!(1 > 975)) {
        fmt.Println(v36)
    } else {
        v36 -= (v36 * v36)
    }
    if true {
        fmt.Println(v36, v36)
        v4 = 1e10
        if true {
            fmt.Println(v36, v36)
            if ((128 > 127) || (v4 <= v4)) {
                v4 = v4
            }
        }
    }
    fmt.Println(v36, v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v36
    _ = v4
}

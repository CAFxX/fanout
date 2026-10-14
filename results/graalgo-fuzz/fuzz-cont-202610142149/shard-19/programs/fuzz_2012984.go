package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 string; F1 uint; F2 bool }
type GenS5 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0(p0 int64) int8 {
    var v6 int = -2147483648
    v7 := []int{}
    v8 := GenS1{F0: 278}
    switch (^(v6)) {
    case 1:
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 float32 = (-2.25 + ((((1.5 + 3.14159) / 3.14159) / float32(1.5)) + 1.5))
            _ = fi9
            _ = v10
        }
    }
    switch (v6 << ((2) & 63)) {
    case 0:
        v6 -= v6
    case 32767:
        if false {
            v6 = (v6 - (v6 >> ((97) & 63)))
            fmt.Println(p0, v6)
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -76
}

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := []int{309, 128}
    var v5 string = "\n\t"
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS2{F0: 855}
        _ = fi11
        _ = v12
    }
    if (439 >= -1) {
        var v13 uint = 2
        _ = v13
    } else {
        v14 := GenS2{F0: 332}
        v15 := GenS3{F0: 4294967295, F1: 42}
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v4 = append(v4, (fi16 / ((-91) | 1)))
        _ = fi16
    }
    var v17 uint = (v1 << (((v1 << ((v1) & 63))) & 63))
    v18 := GenS4{F0: ((v5 + "hello") + (v5 + "foo bar")), F1: 475, F2: true}
    for fi19 := 0; fi19 < 9; fi19++ {
        v5 = ("a" + "foo bar")
        for fi20 := 0; fi20 < 14; fi20++ {
            switch "hello" {
            case "foo bar":
                fmt.Println(v17, v5)
            case "a":
                var v21 float32 = float32((3.14159 * 1e10))
                _ = v21
            case "hello":
                delete(v2, 65535)
            }
            _ = fi20
        }
        _ = fi19
    }
    v17 += 4294967295
    var v22 uint64 = 18446744073709551615
    for fi23 := 0; fi23 < 19; fi23++ {
        if true {
            if (uint(v22) != 255) {
                v5 = v5
                v24 := GenS1{F0: 255}
                switch v5 {
                case "a":
                    fmt.Println(fi23, v1, fi23)
                case "":
                    v1 += (v1 >> ((v1) & 63))
                default:
                    var v25 uint = v17
                    _ = v25
                }
                _ = v24
            }
            fmt.Println(v17)
            fmt.Println(fi23, v1)
        } else {
            v5 = v5
        }
        fmt.Println(v17, v22)
        _ = fi23
    }
    _ = len(v4)
    v2[4294967295] = true
    var v26 string = v5
    if (324 >= 0) {
        fmt.Println(v1, v17, v5)
    } else {
        v1 -= 255
        v27 := GenS5{F0: -31, F1: int8(v1), F2: 0}
        var v28 int16 = -32768
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int }

func main() {
    var v1 float32 = 1.5
    var v2 int = 840
    v3 := make(map[uint32]bool)
    v4 := []int{}
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            var v7 int16 = 42
            if ((true && false) && (227 == 2)) {
                var v8 int8 = 123
                v1 = v1
                for fi9 := 0; fi9 < 6; fi9++ {
                    switch 9223372036854775807 {
                    case -1:
                        var v10 uint64 = 0
                        _ = v10
                    case 317:
                        _ = len(v4)
                    case -528:
                        v8 -= 111
                    }
                    _ = fi9
                }
                _ = v8
            } else {
                v7 += v7
                v2 += (^(v2))
                var v11 float32 = v1
                _ = v11
            }
            _ = v7
        }
        switch "a" {
        case "a":
            var v12 string = ("a" + ("a" + "hello"))
            _ = v12
        }
        v2 += fi6
        _ = fi6
    }
    var v13 int32 = 1
    switch v2 {
    case 504:
        v1 = (v1 + (1e10 / v1))
    case 373:
        switch "foo bar" {
        case "\n\t":
            v1 += (((3.14159 + float32(float64(float32(0.1)))) * ((float32(0.0) + 3.14159) / 3.14159)) * 1.5)
        case "":
            v2 += (v2 % ((v2) | 1))
        }
    }
    v14 := GenS1{F0: 868}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        v15 := GenS1{F0: 38}
        _ = v15
    case "\n\t":
        var v16 uint32 = 2147483648
        _ = v16
    case "hello":
        var v17 int64 = 0
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 11; fi18++ {
            v2 += (-(v2))
            var v19 uint16 = 65535
            v19 = (v19 - (v19 / ((884) | 1)))
            _ = fi18
            _ = v19
        }
    } else {
        fmt.Println(v13, v2)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = v1
    fmt.Println(v2, v2)
    switch "foo bar" {
    case "":
        fmt.Println(v2, v2, v13)
    case "hello":
        fmt.Println(v2, v13)
    case "\n\t":
        for fi20 := 0; fi20 < 8; fi20++ {
            if (!false) {
                var v21 uint64 = 0
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 float32 = 0.0
                    v1 += float32(1e10)
                    v24 := GenS2{F0: v13}
                    delete(v5, 65535)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v21
            } else {
                if len(v4) > 0 {
                    v4[0] = (fi20 - v2)
                }
                v2 += (v2 % (((fi20 + fi20)) | 1))
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 string = ("hello" + "foo bar")
                    var v27 uint8 = 50
                    var v28 int8 = 1
                    v29 := GenS3{F0: v28, F1: (fi25 + 2147483647)}
                    v27 = v27
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi20
        }
    }
    v13 += v13
    v1 -= (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

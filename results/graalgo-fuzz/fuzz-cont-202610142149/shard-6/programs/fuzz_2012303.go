package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
func genFn0() int32 {
    var v6 bool = (693 == 2)
    v7 := GenS1{F0: 127, F1: 251, F2: -1}
    var v8 bool = v6
    for fi9 := 0; fi9 < 2; fi9++ {
        v8 = false
        _ = fi9
    }
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int = 128
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint = 4294967295
    v5 := make(map[string]uint32)
    v10 := []int{415}
    var v11 uint8 = 42
    v11 = (v11 / (((v11 & uint8(v4))) | 1))
    v12 := GenS2{F0: 1}
    if true {
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS3{F0: v4, F1: 452, F2: v4}
                _ = fi13
                _ = v14
            }
        }
        switch 42 {
        case 629:
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 float32 = -2.25
                v17 := v3[(fi15 - -402)]
                v10 = append(v10, v1)
                _ = fi15
                _ = v16
                _ = v17
            }
        case 32767:
            var v18 int32 = 42
            _ = v18
        case -2147483648:
            var v19 int64 = -107
            _ = v19
        }
    }
    v1 = (v1 & v1)
    switch v1 {
    case -2147483648:
        var v20 float32 = 3.14159
        _ = v20
    case -101:
        var v21 float64 = (float64(1.5) / 1.5)
        _ = v21
    }
    switch "hello" {
    case "":
        delete(v5, ("\n\t" + ((("foo bar" + ("\n\t" + ("" + ("foo bar" + "hello")))) + "foo bar") + ("a" + "foo bar"))))
    case "hello":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 int64 = -1
            v5["foo bar"] = 65535
            v11 -= v11
            _ = fi22
            _ = v23
        }
    case "foo bar":
        if false {
            v1 -= 42
        }
    }
    if len(v10) > 0 {
        v10[0] = (v1 & 0)
    }
    if (127 <= (v1 >> (((v4 / ((818) | 1))) & 63))) {
        fmt.Println(v4)
    } else {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

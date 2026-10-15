package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
func genFn0() int8 {
    var v4 int = 32767
    var v5 string = ((("foo bar" + (("\n\t" + ("hello" + "")) + "")) + "a") + "foo bar")
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint) int8 {
    p0 += 1
    fmt.Println(p0)
    _ = p0
    return 7
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 bool = (114 >= 1)
    v6 := []int{9223372036854775807}
    var v7 int32 = -2147483648
    v7 += (^(v7))
    var v8 uint8 = 65
    var v9 int8 = 118
    var v10 float32 = (1.5 * 1e10)
    switch "a" {
    case "hello":
        switch "a" {
        case "hello":
            if false {
                var v11 int16 = -559
                v10 += ((v10 * (v10 / 1.0)) * (0.1 / ((v10 - (3.14159 + 1.5)) * ((v10 / (float32((3.14159 / (0.1 / 1.0))) / float32((0.1 / (float64((v10 + (v10 / 3.14159))) + (0.1 - -2.25)))))) * (1.5 / float32(float64(1.5)))))))
                v8 = v8
                _ = v11
            }
        case "foo bar":
            fmt.Println(v8, v9, v7)
        }
    case "\n\t":
        var v12 string = "\n\t"
        _ = v12
    case "":
        var v13 float32 = v10
        _ = v13
    default:
        var v14 int64 = 0
        _ = v14
    }
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v15 int64 = 9223372036854775807
        _ = v15
    } else {
        if (!v3) {
            v16 := GenS1{F0: 128, F1: 42}
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 float32 = float32(1e10)
                fmt.Println(v7, fi17, fi17)
                v3 = (0 > 731)
                _ = fi17
                _ = v18
            }
            fmt.Println(v3)
            _ = v16
        } else {
            for fi19 := 0; fi19 < 11; fi19++ {
                v20 := GenS2{F0: ""}
                var v21 uint64 = 2
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    var v22 float32 = float32(float64(((1e10 / 1.0) + (v10 / 0.1))))
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(fi23, v8)
        for fi24 := 0; fi24 < 11; fi24++ {
            v3 = v3
            _ = fi24
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS3{F0: v7, F1: 42, F2: true}
            _ = fi25
            _ = v26
        }
        _ = fi23
    }
    var v27 bool = (!((87 == (v9 - (^(v9)))) && v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

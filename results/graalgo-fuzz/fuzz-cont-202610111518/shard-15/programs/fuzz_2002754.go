package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 float64 = (float64(((float32(float64(1.5)) - -2.25) / (float32(float64(0.0)) + (1e10 + 1.5)))) * 1e10)
    v2 := []int{-775, 2147483647}
    v3 := make(map[uint32]bool)
    var v4 uint64 = 0
    var v5 bool = (!(((((-2.25 + (float64(1.5) - v1)) * ((v1 - 0.0) - 3.14159)) + v1) * 0.0) >= v1))
    v6 := make(map[int]bool)
    v7 := v6[0]
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 4294967295}
        switch "foo bar" {
        case "hello":
            var v10 uint32 = 898
            _ = v10
        case "":
            switch ("" + "hello") {
            case "a":
                var v11 uint32 = 1
                _ = v11
            case "":
                switch 0 {
                case 9223372036854775807:
                    if v5 {
                        fmt.Println(fi8)
                    } else {
                        fmt.Println(fi8)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case -61:
                    v12 := GenS2{F0: 0, F1: ("foo bar" + ("foo bar" + "\n\t")), F2: 0}
                    _ = v12
                case 32767:
                    fmt.Println(v7, v7, fi8)
                default:
                    var v13 int64 = 893
                    _ = v13
                }
            case "hello":
                switch -9223372036854775808 {
                case 4294967295:
                    v14 := GenS1{F0: (^(v4))}
                    _ = v14
                }
            default:
                var v15 int64 = 42
                _ = v15
            }
        }
        v7 = (42 != 1)
        _ = fi8
        _ = v9
    }
    if (!v5) {
        v4 = 18446744073709551615
        v7 = true
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint16 = 42
            v1 = v1
            _ = len(v6)
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        switch (("foo bar" + ("foo bar" + "a")) + "a") {
        case "":
            v7 = v5
        }
        _ = fi18
    }
    v7 = v7
    v7 = true
    v19 := GenS2{F0: 255, F1: "", F2: uint16(v4)}
    v7 = (-945 >= 42)
    var v20 int32 = -1
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 string = "hello"
        switch v22 {
        case "hello":
            for fi23 := 0; fi23 < 13; fi23++ {
                v2 = append(v2, 0)
                var v24 uint16 = 2
                _ = fi23
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v4 = v4
            v26 := GenS3{F0: 127}
            _ = fi25
            _ = v26
        }
        _ = fi21
        _ = v22
    }
    var v27 float32 = float32(((v1 + 1.5) - -2.25))
    for fi28 := 0; fi28 < 5; fi28++ {
        switch ("" + (("" + "hello") + ("" + "a"))) {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v29 float64 = float64(((((((v27 * 3.14159) / (0.0 / v27)) + v27) * v27) - v27) / 1e10))
            _ = v29
        }
        v1 += v1
        v5 = ((1e10 - v1) != (0.0 - 0.1))
        _ = fi28
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int16 = 1
            _ = fi31
            _ = v32
        }
        v5 = (!true)
        _ = len(v6)
        _ = fi30
    }
    var v33 int32 = -278
    var v34 uint16 = 0
    v35 := GenS3{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

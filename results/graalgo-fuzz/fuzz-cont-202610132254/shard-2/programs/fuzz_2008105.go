package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    v2 := []int{2147483647, -638}
    var v3 int = 576
    var v4 bool = (2 <= 656)
    v5 := []uint32{65535, 305, 2}
    var v6 bool = (false && false)
    v3 -= (v3 - (v3 / ((v3) | 1)))
    v3 -= (v3 / ((v3) | 1))
    var v7 uint = 255
    v4 = v6
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    var v8 int64 = -74
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint64 = 4294967295
        fmt.Println(v8)
        if (v7 >= (v7 / (((v7 ^ 18446744073709551615)) | 1))) {
            v4 = ("foo bar" == "hello")
            if v4 {
                for fi11 := 0; fi11 < 4; fi11++ {
                    _ = len(v5)
                    switch "hello" {
                    case "a":
                        if false {
                            var v12 uint = v7
                            var v13 float64 = float64(0.1)
                            var v14 int16 = 140
                            v15 := GenS1{F0: 723}
                            var v16 uint64 = v10
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    case "hello":
                        switch fi9 {
                        case -547:
                            var v17 uint8 = 9
                            _ = v17
                        case 4294967295:
                            var v18 string = "\n\t"
                            _ = v18
                        default:
                            var v19 int16 = -118
                            _ = v19
                        }
                    case "foo bar":
                        var v20 uint16 = 255
                        _ = v20
                    }
                    var v21 int32 = 2147483647
                    _ = fi11
                    _ = v21
                }
                fmt.Println(v4)
                for fi22 := 0; fi22 < 18; fi22++ {
                    v10 = v10
                    _ = fi22
                }
            } else {
                var v23 int = 42
                v5 = append(v5, 903)
                _ = v23
            }
            var v24 float32 = (0.1 - ((3.14159 + (float32(float64(3.14159)) * (0.0 * 3.14159))) - 1.5))
            _ = v24
        }
        _ = fi9
        _ = v10
    }
    v25 := GenS2{F0: v8}
    fmt.Println(v6, v8)
    if v6 {
        fmt.Println(v7)
        switch (v3 - -1) {
        case -44:
            switch ("" + ("" + "")) {
            case "hello":
                fmt.Println(v7, v4)
            default:
                var v26 int = v3
                _ = v26
            }
        case 831:
            if (v4 && v4) {
                _ = len(v1)
                _ = len(v5)
            }
        case 4294967295:
            for fi27 := 0; fi27 < 19; fi27++ {
                for fi28 := 0; fi28 < 5; fi28++ {
                    var v29 uint64 = 1
                    v6 = v6
                    var v30 uint32 = 1
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        }
        v7 = (^(v7))
    }
    v6 = v6
    v4 = v6
    fmt.Println(v3, v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

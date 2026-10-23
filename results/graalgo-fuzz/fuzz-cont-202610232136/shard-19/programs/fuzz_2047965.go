package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 uint32 = uint32(v2)
    if ("foo bar" >= ("a" + ("a" + ""))) {
        if false {
            v2 -= 0
            for fi5 := 0; fi5 < 17; fi5++ {
                switch "hello" {
                case "foo bar":
                    for fi6 := 0; fi6 < 14; fi6++ {
                        _ = len(v1)
                        fmt.Println(fi6)
                        _ = fi6
                    }
                case "\n\t":
                    v7 := []int{fi5}
                    _ = v7
                default:
                    v8 := []int{}
                    _ = v8
                }
                _ = fi5
            }
        }
    } else {
        var v9 uint8 = 255
        _ = v9
    }
    v10 := []int{}
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v2, v4, fi11)
        _ = fi11
    }
    v4 = (v4 + v4)
    if true {
        var v12 uint8 = 7
        v13 := v3[v4]
        v13 += 4294967295
        _ = v12
        _ = v13
    }
    fmt.Println(v4)
    _ = len(v10)
    if true {
        switch "a" {
        case "\n\t":
            v3[v4] = v4
        default:
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 uint32 = (^(v4))
                var v16 uint64 = 124
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    } else {
        fmt.Println(v4, v2)
        v17 := v1[v4]
        _ = v17
    }
    var v18 float64 = (-2.25 + 0.1)
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS1{F0: 0, F1: 503, F2: -833}
        var v21 string = ""
        v22 := GenS1{F0: -712, F1: -1, F2: 1}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int = -1
    if (!false) {
        var v24 float32 = 3.14159
        v2 = (v2 % ((2) | 1))
        var v25 string = ("\n\t" + "hello")
        _ = v24
        _ = v25
    } else {
        v26 := GenS2{F0: 2}
        for fi27 := 0; fi27 < 10; fi27++ {
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int8 = 90
                v23 -= 0
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        v30 := GenS3{F0: int8(v4), F1: "foo bar"}
        _ = v26
        _ = v30
    }
    v31 := GenS4{F0: 420, F1: 1}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v18 -= ((((3.14159 - v18) + 0.0) / ((v18 - ((float64(((1.5 / 1e10) * (((((float32(-2.25) / float32((v18 * ((float64(1.5) - (v18 / v18)) + ((float64((1.5 / 0.1)) - v18) / 3.14159))))) * 0.1) + 1e10) - ((-2.25 - 1e10) - 3.14159)) / (1e10 - (1.5 + 0.1))))) - v18) - v18)) * v18)) - float64(-2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
}

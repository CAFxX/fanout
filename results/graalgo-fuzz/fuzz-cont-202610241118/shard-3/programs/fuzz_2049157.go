package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 string; F2 uint }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0() uint16 {
    if false {
        var v6 uint = 1
        _ = v6
    } else {
        var v7 string = "hello"
        var v8 float64 = ((0.0 - (0.1 * -2.25)) * 1.5)
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1}
    var v10 uint16 = 0
    v10 -= v10
    var v11 int32 = 0
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 bool = false
    var v2 int = -33
    v3 := make(map[uint32]uint32)
    var v4 float32 = 0.1
    var v5 int8 = 1
    v2 += v2
    var v12 int8 = 0
    switch (v2 * v2) {
    case 1:
        switch -565 {
        case 242:
            v13 := GenS1{F0: 0}
            _ = v13
        case 0:
            if true {
                v5 = v12
            }
        case 42:
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := []int{v2}
                if len(v15) > 0 {
                    v15[0] = (fi14 & fi14)
                }
                _ = fi14
                _ = v15
            }
        default:
            delete(v3, 4294967295)
        }
    case 42:
        var v16 int32 = 1
        _ = v16
    case 9223372036854775807:
        v3[4294967295] = 42
    }
    v5 = (v5 % (((v12 * v5)) | 1))
    v12 -= v5
    var v17 bool = (("a" <= "\n\t") || v1)
    var v18 int = -137
    v19 := v3[617]
    v20 := []int{(v18 ^ v18), v18, v18}
    fmt.Println(v4, v19)
    if true {
        if v1 {
            var v21 int = -2147483648
            _ = v21
        } else {
            switch ("foo bar" + ("hello" + "a")) {
            case "\n\t":
                if len(v20) > 0 {
                    v20[0] = (^(v18))
                }
            default:
                fmt.Println(v1, v12, v19)
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                v5 += (v12 + v5)
                v23 := v3[(v19 & 623)]
                for fi24 := 0; fi24 < 13; fi24++ {
                    v25 := GenS1{F0: 1}
                    var v26 float32 = v4
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
        }
        v17 = (!v1)
        switch "" {
        case "hello":
            v27 := GenS2{F0: "foo bar", F1: "", F2: 146}
            _ = v27
        case "a":
            v28 := GenS1{F0: 128}
            _ = v28
        case "\n\t":
            if (329 != -1) {
                switch "foo bar" {
                case "hello":
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v19 = (v19 * 275)
                        v4 += float32(1.5)
                        _ = fi29
                    }
                case "\n\t":
                    var v30 int8 = (v5 >> ((2) & 63))
                    _ = v30
                case "a":
                    v31 := GenS1{F0: 1}
                    _ = v31
                default:
                    v1 = false
                }
                v32 := GenS3{F0: 0, F1: -45}
                _ = v32
            } else {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                switch ("hello" + "") {
                case "a":
                    fmt.Println(v5)
                case "\n\t":
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 float64 = float64(v4)
                        v5 -= (v12 * 100)
                        _ = fi33
                        _ = v34
                    }
                case "hello":
                    var v35 uint16 = 637
                    _ = v35
                default:
                    var v36 int16 = -1
                    _ = v36
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

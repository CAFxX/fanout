package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    var v3 int16 = 0
    v4 := []int{-362, 738, -292, -590}
    v5 := make(map[int]string)
    var v6 float32 = (((0.1 * (0.0 / 3.14159)) - (0.0 - ((0.0 * ((0.1 + -2.25) * 1e10)) / 1e10))) - 3.14159)
    v7 := GenS1{F0: 1, F1: -1, F2: v3}
    var v8 int8 = 1
    var v9 int16 = 1000
    var v10 int64 = 42
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v6)
        v3 = (v3 << ((42) & 63))
        _ = len(v2)
        _ = fi11
    }
    v2 = append(v2, -773)
    var v12 uint16 = 32767
    var v13 string = ("hello" + "")
    var v14 float32 = (v6 * (float32(float64((v6 / v6))) / 0.1))
    fmt.Println(v13, v6, v6)
    switch 32767 {
    case 1:
        switch "foo bar" {
        case "":
            switch -267 {
            case -939:
                switch (v13 + v13) {
                case "":
                    var v15 int = 128
                    _ = v15
                case "a":
                    _ = len(v2)
                case "foo bar":
                    v16 := GenS1{F0: -9223372036854775808, F1: 42, F2: (v9 ^ int16(v10))}
                    _ = v16
                default:
                    var v17 int = 128
                    _ = v17
                }
            case -9223372036854775808:
                v14 = v14
            case 891:
                var v18 uint64 = 266
                _ = v18
            }
        case "foo bar":
            var v19 float32 = v14
            _ = v19
        default:
            v20 := GenS1{F0: 663, F1: 2147483647, F2: (^(v3))}
            _ = v20
        }
    case 4294967295:
        for fi21 := 0; fi21 < 10; fi21++ {
            switch 34 {
            case -1:
                fmt.Println(v3, v10, v9)
            case 340:
                v1[(fi21 + fi21)] = (fi21 % (((fi21 ^ fi21)) | 1))
            case 2147483647:
                for fi22 := 0; fi22 < 9; fi22++ {
                    v3 = (v3 / ((v3) | 1))
                    switch -1 {
                    case -229:
                        fmt.Println(v8, v14, v14)
                    default:
                        if false {
                            v9 = 0
                            var v23 bool = (!true)
                            v13 = ""
                            _ = v23
                        } else {
                            v24 := v5[(fi22 * (-(fi22)))]
                            switch (^(fi22)) {
                            case 330:
                                fmt.Println(v8, fi22)
                            case 152:
                                fmt.Println(v14, v10)
                            default:
                                _ = len(v2)
                                var v25 uint16 = v12
                                _ = v25
                            }
                            v6 = v14
                            _ = v24
                        }
                    }
                    _ = fi22
                }
            }
            if false {
                v26 := GenS1{F0: fi21, F1: -426, F2: (v9 * (v9 | 222))}
                v2 = append(v2, -275)
                _ = v26
            }
            var v27 float32 = ((float32(0.0) - (((1.5 - (0.0 - 0.0)) / (0.1 + (float32(-2.25) / 0.1))) * -2.25)) + v6)
            _ = fi21
            _ = v27
        }
    }
    var v28 uint8 = 26
    v28 = (v28 % ((219) | 1))
    if false {
        var v29 float32 = v6
        _ = v29
    } else {
        if (((uint32(v28) > 303) && (!true)) || (65535 != 4294967295)) {
            var v30 int32 = -1
            _ = v30
        } else {
            if ((!(!false)) && false) {
                var v31 int = 127
                fmt.Println(v6)
                _ = v31
            }
        }
        var v32 uint8 = 39
        _ = v32
    }
    switch ((("\n\t" + (v13 + (v13 + v13))) + v13) + ((v13 + ((("foo bar" + v13) + "a") + v13)) + (v13 + (v13 + ("\n\t" + v13))))) {
    case "hello":
        v10 -= 9223372036854775807
    case "a":
        v33 := GenS1{F0: -640, F1: 2147483647, F2: (v3 ^ (v3 ^ v3))}
        _ = v33
    default:
        v8 += v8
    }
    v9 = v3
    var v34 int32 = -561
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

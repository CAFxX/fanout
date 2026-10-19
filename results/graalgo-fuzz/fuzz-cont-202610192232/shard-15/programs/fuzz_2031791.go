package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
type GenS4 struct { F0 uint; F1 uint32 }

func main() {
    var v1 uint64 = 0
    var v2 int = 4294967295
    var v3 uint8 = 0
    v4 := []int{(v2 << ((uint(v1)) & 63)), v2}
    for fi5 := 0; fi5 < 5; fi5++ {
        v2 -= (v2 / ((9223372036854775807) | 1))
        _ = fi5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = -74
        fmt.Println(v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
        _ = v7
    }
    var v8 int = (v2 + v2)
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 float64 = float64((float32(((3.14159 + 0.1) + (1.5 - (0.1 + (-2.25 + -2.25))))) + float32(((float64(1.5) + 0.1) + ((3.14159 / (1e10 + (((1e10 * (((((float64(((0.0 / 1.5) - 1e10)) + (0.0 / 3.14159)) - 1e10) + 1.5) - 1.5) - (-2.25 / (0.0 - (-2.25 * 1e10))))) / 3.14159) * 1.5))) - 0.0)))))
        _ = fi10
        _ = v11
    }
    switch v2 {
    case 4294967295:
        if ((!false) && (((^(v2)) != v2) && false)) {
            v9 = v9
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v2, v1, v3)
                _ = fi12
            }
            v1 += (v1 ^ (v1 ^ (v1 & v1)))
        } else {
            for fi13 := 0; fi13 < 13; fi13++ {
                if ((0.0 / float64(1.5)) == 0.1) {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        switch ("hello" + "a") {
                        case "foo bar":
                            var v15 uint8 = v3
                            _ = v15
                        default:
                            var v16 uint = 4294967295
                            _ = v16
                        }
                        if (1 >= 1) {
                            if false {
                                v2 = v8
                                switch (-(fi14)) {
                                case 32768:
                                    v4 = append(v4, (v8 << ((1) & 63)))
                                }
                                v4 = append(v4, (v8 << ((434) & 63)))
                            } else {
                                var v17 string = ""
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v17
                            }
                        } else {
                            v2 = fi13
                            var v18 int8 = 42
                            _ = v18
                        }
                        _ = fi14
                    }
                    v19 := GenS1{F0: 262, F1: true}
                    _ = v19
                } else {
                    var v20 int16 = 1
                    fmt.Println(v8)
                    _ = v20
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    fmt.Println(v9)
                    v4 = append(v4, -1)
                    _ = fi21
                }
                _ = fi13
            }
            v1 = (v1 / ((v1) | 1))
        }
    case 32767:
        v8 -= (v2 % (((v2 * v8)) | 1))
    case 9223372036854775807:
        v1 = (v1 * v1)
    default:
        if len(v4) > 0 {
            v4[0] = v2
        }
    }
    var v22 uint = 255
    fmt.Println(v2)
    for fi23 := 0; fi23 < 16; fi23++ {
        v9 = v9
        v24 := GenS2{F0: 1000}
        v25 := GenS3{F0: 1, F1: "a", F2: 659}
        _ = fi23
        _ = v24
        _ = v25
    }
    if ((1e10 * float32(float64(1.5))) > 3.14159) {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int32 = 801
            v2 = (fi26 | (v8 << ((uint(v3)) & 63)))
            _ = fi26
            _ = v27
        }
        if true {
            v28 := GenS3{F0: 2147483647, F1: v9, F2: -50}
            switch ("a" + "a") {
            case "":
                var v29 float32 = 1.5
                _ = v29
            case "foo bar":
                v30 := GenS4{F0: 42, F1: 0}
                _ = v30
            default:
                switch (v8 ^ (v8 * (v8 >> ((v22) & 63)))) {
                case 249:
                    var v31 int = v8
                    _ = v31
                default:
                    var v32 int16 = 1000
                    _ = v32
                }
            }
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

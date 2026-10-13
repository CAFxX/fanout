package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint) bool {
    fmt.Println(p0)
    var v4 float64 = -2.25
    switch ("" + ("a" + "a")) {
    case "\n\t":
        for fi5 := 0; fi5 < 20; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                switch "\n\t" {
                case "foo bar":
                    if true {
                        v4 -= v4
                        if false {
                            var v7 int8 = 0
                            fmt.Println(v7, fi5)
                            v8 := []int{fi5, (fi6 / (((fi6 / ((fi6) | 1))) | 1)), 32767}
                            _ = v7
                            _ = v8
                        } else {
                            var v9 int8 = 80
                            if (!false) {
                                p0 = p0
                                v10 := []int{(fi6 ^ (fi6 + fi6)), (fi5 % (((fi5 - -421)) | 1)), (fi6 + (fi6 >> (((p0 ^ p0)) & 63))), -9223372036854775808}
                                _ = v10
                            } else {
                                v11 := GenS1{F0: 308, F1: 1}
                                _ = v11
                            }
                            _ = v9
                        }
                        var v12 uint16 = 171
                        _ = v12
                    }
                default:
                    v13 := GenS1{F0: -9223372036854775808, F1: -1}
                    _ = v13
                }
                var v14 uint = 905
                fmt.Println(v14, v4, p0)
                v15 := GenS2{F0: 255, F1: (42 != 2), F2: 127}
                _ = fi6
                _ = v14
                _ = v15
            }
            _ = fi5
        }
    case "":
        v4 = float64(1e10)
    default:
        var v16 int = 332
        _ = v16
    }
    if (0 >= 1) {
        if (!(v4 < (v4 + 1.5))) {
            switch 0 {
            case -1:
                v17 := []int{int(p0), 32767, -9223372036854775808}
                _ = v17
            }
            v4 = 0.0
        }
    }
    _ = p0
    _ = v4
    return (-9223372036854775808 >= 548)
}

func main() {
    var v1 int = 127
    v2 := make(map[string]bool)
    var v3 uint64 = 846
    var v18 float64 = float64((3.14159 - 3.14159))
    fmt.Println(v1)
    if (304 < 221) {
        var v19 int16 = 32767
        _ = v19
    }
    v3 += v3
    v18 = 0.0
    if true {
        fmt.Println(v3)
        var v20 float64 = v18
        var v21 float32 = float32(float64(0.1))
        _ = v20
        _ = v21
    }
    var v22 int64 = 9223372036854775807
    var v23 int64 = (v22 % ((v22) | 1))
    if true {
        if true {
            v3 += v3
            var v24 int16 = 42
            _ = v24
        }
        fmt.Println(v22, v22, v18)
        v18 -= (v18 + float64((0.1 - ((0.0 + 3.14159) + 0.1))))
    } else {
        v25 := GenS3{F0: v23}
        v26 := GenS2{F0: 0, F1: (0 != 1), F2: (^(v1))}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        switch ("a" + "hello") {
        case "":
            switch v1 {
            case -9223372036854775808:
                if (32767 < -32768) {
                    var v28 uint32 = 4294967295
                    var v29 int32 = -701
                    _ = v28
                    _ = v29
                } else {
                    v2["a"] = ((!(3.14159 < v18)) && false)
                }
            case -314:
                if true {
                    var v30 float32 = (((1e10 + 0.0) * (3.14159 - float32(v18))) + (1.5 - 0.0))
                    _ = v30
                } else {
                    v31 := []int{0, 128, fi27}
                    _ = v31
                }
            default:
                var v32 bool = (!true)
                _ = v32
            }
        }
        v1 = -9223372036854775808
        _ = fi27
    }
    delete(v2, "hello")
    v33 := GenS3{F0: v23}
    v34 := GenS4{F0: v23}
    var v35 uint32 = 387
    var v36 int64 = 42
    var v37 string = "\n\t"
    v35 += v35
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}

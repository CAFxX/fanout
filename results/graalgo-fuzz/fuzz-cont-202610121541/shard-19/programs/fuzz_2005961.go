package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 string = "hello"
    var v4 bool = (false && (((1.5 + ((1.5 / -2.25) + 0.1)) * (1.5 + (3.14159 + 3.14159))) == 0.0))
    var v5 float32 = (0.1 + (3.14159 - 1.5))
    var v6 float64 = (-2.25 * float64(v5))
    v5 += 3.14159
    if (42 == 2147483647) {
        var v7 float32 = v5
        var v8 string = ((("\n\t" + "a") + v3) + "foo bar")
        v3 = v8
        _ = v7
        _ = v8
    }
    if (-752 == 0) {
        var v9 uint8 = 127
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: 1}
    var v11 string = v3
    if false {
        var v12 int64 = 42
        v3 = "hello"
        _ = v12
    } else {
        switch v11 {
        case "hello":
            v13 := GenS2{F0: 2, F1: 127, F2: -1}
            _ = v13
        }
        var v14 uint16 = 503
        _ = v14
    }
    switch v3 {
    case "hello":
        switch 943 {
        case 1:
            var v15 uint = 902
            _ = v15
        default:
            var v16 int32 = -80
            _ = v16
        }
    }
    if v4 {
        v17 := []int{-419, 9223372036854775807, 32767, 2147483647}
        _ = v17
    } else {
        v6 += 1e10
    }
    v4 = true
    for fi18 := 0; fi18 < 4; fi18++ {
        v3 = "hello"
        v19 := GenS2{F0: 32767, F1: 128, F2: 42}
        _ = fi18
        _ = v19
    }
    fmt.Println(v5, v3)
    if v4 {
        if (-2.25 <= v5) {
            v20 := []int{241}
            fmt.Println(v6, v3)
            _ = v20
        } else {
            var v21 uint64 = 0
            fmt.Println(v11, v3)
            _ = v21
        }
    }
    var v22 int32 = 2147483647
    v4 = ((1 <= 255) || (255 >= 255))
    v5 -= (v5 + v5)
    fmt.Println(v22, v22)
    for fi23 := 0; fi23 < 11; fi23++ {
        if (136 != 225) {
            v24 := GenS2{F0: 291, F1: fi23, F2: 0}
            fmt.Println(v22, v6, v11)
            _ = v24
        } else {
            v22 = (v22 & (v22 | -442))
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            switch "hello" {
            case "\n\t":
                if v4 {
                    if true {
                        var v26 int16 = -32768
                        _ = v26
                    } else {
                        v5 -= v5
                        var v27 uint16 = 65535
                        var v28 uint = 42
                        _ = v27
                        _ = v28
                    }
                    switch fi25 {
                    case 747:
                        var v29 uint32 = 784
                        _ = v29
                    case 127:
                        if false {
                            v4 = true
                        } else {
                            v30 := GenS2{F0: 255, F1: 0, F2: -641}
                            _ = v30
                        }
                    default:
                        if (v22 != (v22 % ((v22) | 1))) {
                            if v4 {
                                var v31 int16 = -810
                                var v32 int = int(v31)
                                v33 := GenS3{F0: -71}
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                var v34 string = v3
                                var v35 string = v34
                                v11 = ""
                                _ = v34
                                _ = v35
                            }
                            v1["a"] = (fi25 * -201)
                            v36 := GenS3{F0: 42}
                            _ = v36
                        }
                    }
                    fmt.Println(v5, v4)
                }
            case "a":
                v4 = v4
            default:
                v11 = (v11 + (v11 + ("hello" + "hello")))
            }
            _ = fi25
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            v3 = v11
            _ = fi37
        }
        _ = fi23
    }
    var v38 float32 = (0.0 + (float32(((v6 + float64(3.14159)) + 0.0)) / v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

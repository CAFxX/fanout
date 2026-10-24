package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 int16) int64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p2, p1, fi5)
        v6 := GenS1{F0: "hello", F1: 555}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: ("" + "foo bar"), F1: 42}
    switch "a" {
    case "foo bar":
        var v8 float64 = -2.25
        _ = v8
    case "":
        v9 := []int{}
        _ = v9
    }
    var v10 int = 0
    var v11 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 1
}

func main() {
    var v1 int64 = -578
    var v2 bool = false
    v3 := make(map[int]uint32)
    var v4 int64 = v1
    v12 := GenS2{F0: 521, F1: 0, F2: 2}
    switch "" {
    case "hello":
        v2 = (128 < 162)
    case "a":
        if (50 <= 162) {
            if (v2 || (725 == 255)) {
                v4 = (v1 >> ((65535) & 63))
                var v13 uint64 = 0
                for fi14 := 0; fi14 < 4; fi14++ {
                    if v2 {
                        v15 := GenS1{F0: "\n\t", F1: fi14}
                        v4 = (v4 | (v4 / ((genFn0((v13 + (^(v13))), (v13 << ((65535) & 63)), -1)) | 1)))
                        _ = v15
                    } else {
                        var v16 int16 = -904
                        _ = v16
                    }
                    fmt.Println(v4, v4, v4)
                    v17 := []int{}
                    _ = fi14
                    _ = v17
                }
                _ = v13
            } else {
                v3[-9223372036854775808] = 2
                for fi18 := 0; fi18 < 16; fi18++ {
                    fmt.Println(v4)
                    if ((0.0 + 1.5) == (1.5 * (((3.14159 - ((-2.25 - 1.5) + 0.0)) * float32((float64((3.14159 * (-2.25 / (1.5 - 3.14159)))) * (((-2.25 - (0.0 * 3.14159)) * 1e10) * 1e10)))) * (0.0 - -2.25)))) {
                        for fi19 := 0; fi19 < 2; fi19++ {
                            switch fi19 {
                            case 326:
                                v2 = true
                            case 32768:
                                if true {
                                    var v20 float64 = (3.14159 * float64(0.1))
                                    _ = v20
                                }
                            case -550:
                                var v21 int64 = (v4 | v4)
                                _ = v21
                            default:
                                fmt.Println(fi19)
                            }
                            _ = fi19
                        }
                        var v22 uint = 65535
                        if (!(true || (v2 || v2))) {
                            var v23 int8 = -127
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 2; fi24++ {
                                fmt.Println(v4, v4)
                                v22 = 1
                                _ = fi24
                            }
                        }
                        _ = v22
                    }
                    _ = fi18
                }
                var v25 bool = (0 >= 255)
                _ = v25
            }
        }
    default:
        var v26 int32 = -1
        _ = v26
    }
    v4 = (v4 + genFn0(1, 104, 0))
    var v27 float64 = 0.0
    fmt.Println(v27, v1, v2)
    switch "a" {
    case "hello":
        var v28 uint32 = 0
        _ = v28
    case "\n\t":
        v29 := GenS3{F0: 1}
        _ = v29
    default:
        if ((1 < 65535) && v2) {
            fmt.Println(v27)
            var v30 float64 = (v27 / 1.5)
            _ = v30
        } else {
            fmt.Println(v4, v1)
            v31 := []int{32768, -546, 2147483647}
            v32 := GenS2{F0: 65535, F1: 27, F2: 143}
            _ = v31
            _ = v32
        }
    }
    v33 := []int{4294967295, 42, 32767}
    var v34 int8 = 7
    var v35 int32 = 42
    fmt.Println(v27)
    var v36 int8 = v34
    var v37 int16 = 32767
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    _ = len(v3)
    if (42 >= 0) {
        v38 := GenS3{F0: 2}
        v39 := GenS2{F0: 2, F1: 870, F2: 42}
        _ = v38
        _ = v39
    } else {
        var v40 uint16 = 0
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}

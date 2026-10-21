package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int8; F1 int16; F2 uint64 }
func genFn0(p0 uint, p1 int16, p2 uint8) uint64 {
    v4 := []int{}
    var v5 int8 = -1
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 bool = (-344 <= -940)
    v2 := make(map[string]uint32)
    v3 := []uint32{2, 840, 2}
    for fi6 := 0; fi6 < 12; fi6++ {
        if len(v3) > 0 {
            v3[0] = 48
        }
        v7 := GenS1{F0: 0, F1: 914}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: 285}
        _ = fi8
        _ = v9
    }
    v10 := v2["hello"]
    fmt.Println(v10)
    switch -2147483648 {
    case 32767:
        v10 = (v10 | 286)
    case 4294967295:
        var v11 uint16 = 0
        _ = v11
    case 32768:
        for fi12 := 0; fi12 < 16; fi12++ {
            v2[("a" + "foo bar")] = v10
            var v13 int = (fi12 + (fi12 | (fi12 << ((4294967295) & 63))))
            _ = fi12
            _ = v13
        }
    }
    switch 2147483647 {
    case 32767:
        var v14 float64 = 0.0
        _ = v14
    case 4294967295:
        if false {
            var v15 float32 = (float32(0.1) * 0.1)
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 = v1
                if ((2147483647 == v10) && v1) {
                    var v17 int64 = -9223372036854775808
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 float32 = v15
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (2 < 2) {
                            v20 := GenS2{F0: ("hello" + "foo bar"), F1: 189, F2: 95}
                            var v21 int32 = 826
                            if (69 >= 2) {
                                switch "foo bar" {
                                case "hello":
                                    _ = len(v3)
                                case "":
                                    var v22 int64 = v17
                                    var v23 uint16 = 32767
                                    _ = v22
                                    _ = v23
                                default:
                                    var v24 string = "a"
                                    _ = v24
                                }
                                v19 += (v19 * v15)
                                switch (-(fi16)) {
                                case -9223372036854775808:
                                    var v25 uint = 42
                                    _ = v25
                                case -703:
                                    v26 := GenS1{F0: genFn0(18446744073709551615, 0, 1), F1: 32767}
                                    _ = v26
                                case 32768:
                                    fmt.Println(v19, v19)
                                }
                            } else {
                                fmt.Println(v15, fi16)
                                v27 := GenS1{F0: 537, F1: 65535}
                                _ = v27
                            }
                            _ = v20
                            _ = v21
                        } else {
                            v28 := GenS3{F0: 127, F1: 42, F2: 0}
                            v29 := GenS3{F0: 1, F1: 32767, F2: 42}
                            _ = v28
                            _ = v29
                        }
                        _ = fi18
                        _ = v19
                    }
                    switch "a" {
                    case "hello":
                        var v30 int = (fi16 + 32768)
                        _ = v30
                    case "":
                        var v31 int64 = v17
                        _ = v31
                    case "a":
                        delete(v2, "foo bar")
                    }
                    _ = v17
                } else {
                    var v32 string = "foo bar"
                    _ = v32
                }
                _ = fi16
            }
            switch -765 {
            case 0:
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 int32 = 2147483647
                    _ = len(v3)
                    _ = fi33
                    _ = v34
                }
            }
            _ = v15
        }
    case 32768:
        var v35 float32 = 3.14159
        _ = v35
    }
    v10 = (v10 & (v10 - v10))
    var v36 string = "hello"
    v10 = 2147483647
    if false {
        fmt.Println(v10, v10)
    } else {
        v10 = (^(v10))
        delete(v2, ("hello" + "foo bar"))
    }
    v10 = (v10 / (((v10 * 926)) | 1))
    var v37 uint32 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}

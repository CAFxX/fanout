package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }

func main() {
    var v1 uint64 = 731
    var v2 int64 = 53
    v3 := []string{"hello", "a", "a", "a"}
    var v4 int32 = int32(v2)
    var v5 int = -2147483648
    v6 := make(map[string]uint32)
    v4 += v4
    v7 := GenS1{F0: "", F1: 255}
    if (2147483648 <= 4) {
        if true {
            for fi8 := 0; fi8 < 16; fi8++ {
                if len(v3) > 0 {
                    v3[0] = ("" + "")
                }
                for fi9 := 0; fi9 < 19; fi9++ {
                    switch ("a" + "foo bar") {
                    case "a":
                        v2 = (v2 << ((4294967295) & 63))
                    default:
                        v2 += int64(v1)
                    }
                    var v10 int16 = 42
                    var v11 uint = 2
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                fmt.Println(v2)
                _ = fi8
            }
            switch "foo bar" {
            case "a":
                var v12 int32 = (v4 + (v4 % (((v4 << ((0) & 63))) | 1)))
                _ = v12
            case "hello":
                v13 := GenS2{F0: -914, F1: 42, F2: -1}
                _ = v13
            default:
                v14 := GenS2{F0: 0, F1: v4, F2: v2}
                _ = v14
            }
            var v15 uint = 346
            _ = v15
        } else {
            v16 := GenS1{F0: "hello", F1: 0}
            for fi17 := 0; fi17 < 14; fi17++ {
                if (!true) {
                    var v18 uint8 = 128
                    _ = v18
                } else {
                    var v19 string = "\n\t"
                    if ((936 < 2) && false) {
                        v5 -= fi17
                        if (false || false) {
                            fmt.Println(v5, fi17)
                        } else {
                            if ((-(v5)) < (fi17 ^ 9223372036854775807)) {
                                if (v2 < 9223372036854775807) {
                                    v19 = (v19 + "hello")
                                    if (!true) {
                                        var v20 uint8 = 1
                                        delete(v6, "a")
                                        var v21 float64 = 1.5
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    v4 = (^(v4))
                                    var v23 int16 = 42
                                    var v24 float64 = -2.25
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                v25 := GenS3{F0: -1, F1: 59, F2: 214}
                                _ = v25
                            }
                            v26 := GenS2{F0: -1, F1: (v4 % ((371) | 1)), F2: (v2 & v2)}
                            v3 = append(v3, ((("" + ("hello" + v19)) + ((v19 + v19) + "")) + v19))
                            _ = v26
                        }
                    } else {
                        fmt.Println(fi17, v2)
                        var v27 int16 = 32767
                        _ = v27
                    }
                    _ = v19
                }
                fmt.Println(fi17, v4)
                for fi28 := 0; fi28 < 9; fi28++ {
                    v4 = (v4 | 42)
                    var v29 float64 = ((0.0 / float64(-2.25)) + (((0.0 + ((((0.1 * 1.5) / 1e10) / 3.14159) - 1e10)) + (1.5 / (1e10 * 0.1))) * 1e10))
                    _ = fi28
                    _ = v29
                }
                _ = fi17
            }
            _ = v16
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            if (!true) {
                var v31 string = "hello"
                _ = len(v3)
                _ = v31
            } else {
                v2 += v2
            }
            var v32 bool = false
            var v33 int64 = -1
            _ = fi30
            _ = v32
            _ = v33
        }
        v34 := GenS3{F0: -1, F1: v4, F2: 1000}
        _ = v34
    }
    v35 := GenS3{F0: -1, F1: (v4 / ((1) | 1)), F2: -993}
    v5 = (^(v5))
    var v36 bool = (32767 <= 531)
    var v37 int16 = -1
    var v38 uint = 65535
    fmt.Println(v5, v5, v38)
    v5 = (v5 - (v5 << ((351) & 63)))
    v39 := GenS1{F0: ("" + ""), F1: 65535}
    v1 += 0
    fmt.Println(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

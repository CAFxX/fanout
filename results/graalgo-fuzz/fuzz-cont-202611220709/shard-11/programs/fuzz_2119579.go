package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint }
func genFn0(p0 string, p1 uint32) bool {
    p0 = p0
    p0 = "a"
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 string = ("a" + "\n\t")
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            if false {
                v3 = (v3 + ("" + ("\n\t" + ("" + ((v3 + v3) + v3)))))
                if true {
                    v3 = v3
                }
                v3 = "a"
            } else {
                var v6 uint8 = 1
                var v7 uint = 42
                v2[("\n\t" + v3)] = (!(1 < -128))
                _ = v6
                _ = v7
            }
            v3 = ((v3 + ("foo bar" + ("hello" + (v3 + (v3 + ""))))) + "")
            _ = fi5
        }
        var v8 uint16 = 694
        var v9 uint32 = 2147483647
        _ = fi4
        _ = v8
        _ = v9
    }
    var v10 uint8 = 37
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v3)
        _ = fi11
    }
    if (0 == -692) {
        var v12 int64 = 9223372036854775807
        _ = v12
    } else {
        if false {
            v10 -= 223
            if false {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 bool = (100 >= 42)
                    v10 += (v10 * (v10 / (((v10 * 42)) | 1)))
                    v15 := []int{32767, (-(fi13))}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    if ((fi16 != fi16) && (2 <= 0)) {
                        v17 := GenS1{F0: 42}
                        v18 := GenS1{F0: 2}
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(fi16)
                    }
                    _ = fi16
                }
                if ((0.0 / -2.25) != 0.0) {
                    switch v3 {
                    case "hello":
                        _ = len(v2)
                    default:
                        var v19 int16 = 188
                        _ = v19
                    }
                } else {
                    var v20 bool = false
                    _ = v20
                }
            } else {
                var v21 float32 = float32(float64(float32(0.0)))
                fmt.Println(v21, v3)
                v3 = ("" + v3)
                _ = v21
            }
        } else {
            var v22 int64 = -9223372036854775808
            v23 := GenS1{F0: 267}
            v24 := GenS1{F0: 608}
            _ = v22
            _ = v23
            _ = v24
        }
        if (1 > 0) {
            var v25 float32 = float32(float64(float32(((float64((float32(((((0.1 / (float64((float32(3.14159) + 0.0)) + 0.0)) + (1.5 - float64((((0.0 - 3.14159) - (1.5 + 1e10)) + 1.5)))) + 1e10) + (3.14159 / 3.14159))) / 1e10)) / 1.0) - 1.5))))
            var v26 int64 = 42
            var v27 int16 = -32768
            _ = v25
            _ = v26
            _ = v27
        }
        if ((float32((float64(0.1) * 0.1)) <= 0.1) && false) {
            if (701 != 201) {
                _ = len(v2)
            }
            if true {
                var v28 float64 = (0.1 + (3.14159 + 0.1))
                var v29 string = (((((v3 + ("\n\t" + v3)) + v3) + (((v3 + v3) + v3) + v3)) + "hello") + "a")
                _ = v28
                _ = v29
            }
        } else {
            if (-1 != int32(v10)) {
                switch 1 {
                case 9223372036854775807:
                    fmt.Println(v10, v10)
                case 2147483647:
                    for fi30 := 0; fi30 < 18; fi30++ {
                        var v31 uint = 4294967295
                        v32 := GenS1{F0: 952}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                default:
                    fmt.Println(v10, v3)
                }
                v33 := GenS2{F0: 70, F1: 255}
                _ = v33
            } else {
                v10 += 128
                fmt.Println(v3, v3)
            }
            v3 = "a"
        }
    }
    fmt.Println(v10, v3)
    for fi34 := 0; fi34 < 1; fi34++ {
        v10 += v10
        for fi35 := 0; fi35 < 15; fi35++ {
            v36 := GenS3{F0: -32768, F1: 42}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 string = ((v3 + "\n\t") + (v3 + v3))
    v38 := GenS2{F0: int8(v10), F1: 255}
    v37 = v3
    for fi39 := 0; fi39 < 6; fi39++ {
        v40 := []int{fi39, fi39, fi39, fi39}
        _ = fi39
        _ = v40
    }
    var v41 bool = (!(((18446744073709551615 >= 625) || ((2 == 255) && false)) && true))
    if (2147483647 == -622) {
        v10 = v10
        v3 = ("a" + "a")
    } else {
        var v42 float64 = ((-2.25 + 3.14159) + 0.0)
        fmt.Println(v42, v3, v10)
        var v43 int32 = -1
        _ = v42
        _ = v43
    }
    v10 -= (v10 * 150)
    var v44 float32 = (0.1 + 0.1)
    v10 = v10
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v41
    _ = v44
}

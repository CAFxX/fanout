package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint64) uint8 {
    var v7 float64 = 1.5
    v8 := []int{32767, 42, 561}
    v7 += v7
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 float64 = ((0.0 * ((1e10 - 1.5) + 1e10)) * 1.5)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    v5 := []uint32{}
    var v6 float32 = 0.1
    v1 = 0.1
    for fi9 := 0; fi9 < 6; fi9++ {
        delete(v2, 42)
        if false {
            v6 -= v6
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            if true {
                v1 += ((v1 + 0.0) - 3.14159)
            } else {
                v11 := GenS1{F0: false, F1: -32768}
                switch fi9 {
                case -1:
                    v5 = append(v5, 337)
                default:
                    var v12 uint64 = 212
                    _ = len(v2)
                    _ = v12
                }
                _ = v11
            }
            var v13 int = (fi10 % ((32767) | 1))
            _ = fi10
            _ = v13
        }
        _ = fi9
    }
    if ((("a" + "hello") + ("\n\t" + "hello")) >= "hello") {
        if false {
            v6 = v6
            var v14 int32 = 627
            _ = v14
        } else {
            if false {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v6 = (v6 - float32(v1))
                    fmt.Println(v1, v1)
                    _ = fi15
                }
                v1 = v1
            } else {
                var v16 int = 200
                v6 += v6
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 float32 = 0.0
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            var v19 string = "hello"
            fmt.Println(v19, v1)
            _ = v19
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS1{F0: true, F1: -60}
            v6 = (1.5 - 0.1)
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int32 = -1
            _ = fi22
            _ = v23
        }
    }
    var v24 uint = 513
    v24 = v24
    v6 = (-2.25 - (-2.25 / (v6 + v6)))
    if false {
        var v25 bool = true
        if v25 {
            var v26 int64 = 9223372036854775807
            _ = v26
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                if v25 {
                    v6 = ((v6 + v6) * (v6 + (((-2.25 + ((v6 * (0.1 / v6)) * ((((1e10 + v6) + float32(float64(float32(-2.25)))) / v6) + v6))) - 1e10) - -2.25)))
                    v29 := GenS2{F0: 0, F1: 0}
                    v25 = (!(v1 == 3.14159))
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        if (42 != 100) {
                            var v30 int64 = 1
                            _ = v30
                        } else {
                            var v31 string = ""
                            v32 := GenS2{F0: 476, F1: uint64(v24)}
                            _ = v31
                            _ = v32
                        }
                    case "":
                        var v33 float64 = v1
                        _ = v33
                    default:
                        v34 := GenS1{F0: (-157 != 0), F1: -654}
                        _ = len(v3)
                        _ = v34
                    }
                    if true {
                        v35 := GenS3{F0: 1, F1: -391}
                        var v36 uint8 = 127
                        _ = v35
                        _ = v36
                    }
                    _ = v29
                }
                fmt.Println(v24)
                v25 = ((1 > 1000) && (398 > 2))
                _ = fi28
            }
            _ = fi27
        }
        _ = v25
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v37 float32 = v6
    var v38 int = 2147483647
    var v39 int = v38
    var v40 string = "\n\t"
    var v41 string = "hello"
    v42 := GenS1{F0: (((3.14159 * (float32(v1) - (float32(-2.25) * 1e10))) + v37) > -2.25), F1: -1}
    var v43 float64 = v1
    var v44 uint64 = 369
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}

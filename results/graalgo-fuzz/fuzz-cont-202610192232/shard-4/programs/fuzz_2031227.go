package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0() uint64 {
    var v6 uint = 0
    v6 = (^(v6))
    if true {
        var v7 float64 = float64(0.0)
        fmt.Println(v6, v6)
        for fi8 := 0; fi8 < 8; fi8++ {
            v7 -= 1e10
            var v9 uint16 = 255
            v6 -= (^(v6))
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        v10 := []int{32767, 535, 127, 9223372036854775807}
        v11 := GenS1{F0: 5}
        var v12 int = 128
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if (2 <= 0) {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS2{F0: 1, F1: 152}
                if true {
                    v16 := GenS3{F0: false, F1: 373, F2: -128}
                    _ = v16
                } else {
                    v6 -= v6
                }
                fmt.Println(fi13, fi13)
                _ = fi14
                _ = v15
            }
            if false {
                v17 := GenS4{F0: 2}
                switch "" {
                case "\n\t":
                    var v18 uint32 = 325
                    _ = v18
                }
                var v19 int16 = -112
                _ = v17
                _ = v19
            } else {
                if false {
                    v20 := []int{}
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    fmt.Println(fi13)
                    _ = v20
                }
            }
        } else {
            v21 := GenS1{F0: 255}
            var v22 float64 = -2.25
            for fi23 := 0; fi23 < 11; fi23++ {
                switch fi13 {
                case 42:
                    v24 := GenS1{F0: 127}
                    _ = v24
                case 32767:
                    v25 := []int{(^(fi13))}
                    _ = v25
                case 0:
                    fmt.Println(v6, v6, fi23)
                }
                fmt.Println(fi23, v6)
                _ = fi23
            }
            _ = v21
            _ = v22
        }
        v26 := []int{-786}
        switch "foo bar" {
        case "":
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 int64 = 0
                v29 := GenS2{F0: -1, F1: 2}
                v26 = append(v26, (-(fi13)))
                _ = fi27
                _ = v28
                _ = v29
            }
        case "foo bar":
            if false {
                v30 := GenS3{F0: false, F1: 4294967295, F2: 1}
                _ = len(v26)
                _ = v30
            }
        case "a":
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        }
        _ = fi13
        _ = v26
    }
    _ = v6
    return 162
}
func genFn1() bool {
    if false {
        v31 := []int{4294967295}
        var v32 uint = 539
        var v33 uint = (v32 << ((1) & 63))
        _ = v31
        _ = v32
        _ = v33
    }
    v34 := []int{929}
    _ = v34
    return (4294967295 > 2)
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    var v3 bool = (-511 >= -1)
    var v4 uint32 = 4294967295
    var v5 int32 = 0
    v3 = (!false)
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint = 42
        _ = fi35
        _ = v36
    }
    v37 := []int{1}
    var v38 float32 = 0.1
    _ = len(v1)
    fmt.Println(v4, v3)
    switch (("foo bar" + "hello") + ("hello" + "\n\t")) {
    case "":
        if v3 {
            v39 := GenS2{F0: 127, F1: 311}
            v38 = float32(float64(((v38 - float32((1e10 - ((0.0 - float64(v38)) * 0.1)))) - (1e10 / v38))))
            _ = v39
        } else {
            v4 -= 1
            if true {
                v5 += (-(v5))
            }
        }
    case "foo bar":
        var v40 float64 = 0.0
        _ = v40
    default:
        delete(v1, "")
    }
    var v41 int32 = (-(v2))
    var v42 float64 = (1.5 - (1e10 * 1.5))
    var v43 string = "foo bar"
    switch 128 {
    case 0:
        v4 += v4
    default:
        if len(v37) > 0 {
            v37[0] = -1
        }
    }
    if v3 {
        v41 = v5
    }
    v44 := GenS5{F0: 534, F1: 32767}
    v41 -= 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}

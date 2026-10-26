package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 uint64, p1 bool, p2 int64) int8 {
    if true {
        v7 := GenS1{F0: 52, F1: ("hello" + ("foo bar" + "foo bar"))}
        p0 = p0
        v8 := GenS1{F0: 1, F1: ("" + "hello")}
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 2
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                p0 = (p0 - 955)
            }
            p1 = (3.14159 == 1.5)
            _ = fi10
        }
        v11 := []int{-1, 2147483647}
        _ = v11
        _ = v9
    }
    var v12 uint = 42
    switch "hello" {
    case "foo bar":
        v13 := GenS2{F0: 128}
        _ = v13
    case "":
        if (!p1) {
            var v14 uint16 = 985
            if false {
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 uint16 = (^(v14))
                    var v17 float32 = float32(3.14159)
                    p1 = (!((-(fi15)) > fi15))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 int32 = 164
                _ = v18
            }
            _ = v14
        } else {
            v12 = v12
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 127
}

func main() {
    var v1 uint8 = 108
    v2 := make(map[uint32]uint32)
    var v3 int64 = -1
    var v4 uint64 = 147
    var v5 float64 = ((3.14159 * float64((float32((0.0 + 0.0)) / (1e10 + (float32(1.5) + (-2.25 + (float32((1e10 + 1e10)) - 0.0))))))) + 3.14159)
    v6 := make(map[string]uint32)
    var v19 uint8 = 128
    v3 += v3
    v20 := GenS2{F0: 32767}
    switch -385 {
    case 32767:
        var v21 int16 = -417
        _ = v21
    default:
        v5 = (float64((1e10 * float32((3.14159 + v5)))) - v5)
    }
    v22 := GenS2{F0: 119}
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := []int{}
        if ((-831 > (^(fi23))) && (false || ((fi23 / (((fi23 ^ (fi23 + fi23))) | 1)) >= (-(fi23))))) {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 uint = 42
                fmt.Println(v3, v3)
                _ = fi25
                _ = v26
            }
        } else {
            var v27 string = "hello"
            v28 := GenS3{F0: 0, F1: 42}
            _ = v27
            _ = v28
        }
        _ = fi23
        _ = v24
    }
    switch "foo bar" {
    case "a":
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS3{F0: 1, F1: 782}
            for fi31 := 0; fi31 < 9; fi31++ {
                switch fi31 {
                case -780:
                    var v32 int8 = 100
                    _ = v32
                case 32767:
                    v33 := v2[401]
                    if ((1e10 / (0.1 / 1.0)) != float32((v5 + (float64(1e10) + (3.14159 + float64((3.14159 - 3.14159))))))) {
                        var v34 float32 = 0.1
                        v35 := GenS3{F0: 2147483647, F1: 1}
                        v3 += 11
                        _ = v34
                        _ = v35
                    }
                    v36 := []int{fi29}
                    v1 = v19
                    var v37 uint8 = 156
                    if false {
                        switch fi29 {
                        case 0:
                            v38 := GenS2{F0: fi29}
                            _ = v38
                        }
                        var v39 string = "\n\t"
                        var v40 bool = true
                        _ = v39
                        _ = v40
                    }
                    if (!false) {
                        v41 := GenS3{F0: 360, F1: 4294967295}
                        _ = v41
                    } else {
                        v19 -= 128
                        v4 -= 1
                    }
                    v37 += (v37 >> ((65535) & 63))
                    _ = v33
                    _ = v36
                    _ = v37
                case 127:
                    v1 += (v1 - v1)
                default:
                    var v42 int32 = 2147483647
                    _ = v42
                }
                _ = fi31
            }
            var v43 float32 = 0.0
            _ = fi29
            _ = v30
            _ = v43
        }
    case "foo bar":
        if false {
            var v44 float64 = (((0.0 * 1e10) / v5) / v5)
            _ = v44
        }
    }
    fmt.Println(v3, v4)
    fmt.Println(v5)
    var v45 uint32 = 907
    fmt.Println(v45, v1, v5)
    v46 := GenS1{F0: 1, F1: ("a" + "foo bar")}
    v19 -= (v19 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}

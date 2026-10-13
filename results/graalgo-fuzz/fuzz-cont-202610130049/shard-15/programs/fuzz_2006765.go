package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint32 }
func genFn0() string {
    var v4 uint32 = 2147483647
    v5 := []int{9223372036854775807, 32767}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint16 = 32767
        switch "\n\t" {
        case "foo bar":
            var v8 int16 = int16(v4)
            _ = v8
        case "":
            if len(v5) > 0 {
                v5[0] = (fi6 - (fi6 << ((255) & 63)))
            }
        case "a":
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 int8 = 0
                var v11 int64 = 42
                v4 += v4
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            v4 -= 4294967295
        }
        v7 = v7
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    _ = v4
    _ = v5
    return (("" + "") + "foo bar")
}
func genFn1() int32 {
    var v12 int64 = -1
    v13 := []int{2147483647}
    _ = v12
    _ = v13
    return -385
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[string]string)
    v3 := []uint32{}
    switch -9223372036854775808 {
    case 4294967295:
        var v14 uint64 = 474
        _ = v14
    case -995:
        var v15 int16 = 0
        _ = v15
    case -352:
        var v16 uint32 = 4294967295
        _ = v16
    }
    if false {
        var v17 int = -944
        if len(v1) > 0 {
            v1[0] = 923
        }
        var v18 bool = (!(!(-962 == -1)))
        _ = v17
        _ = v18
    }
    v3 = append(v3, 65535)
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 float32 = -2.25
            _ = fi20
            _ = v21
        }
        var v22 uint32 = 816
        _ = fi19
        _ = v22
    }
    switch "hello" {
    case "a":
        var v23 int = -2147483648
        _ = v23
    case "":
        var v24 int = -418
        _ = v24
    }
    var v25 float64 = 3.14159
    fmt.Println(v25)
    var v26 uint64 = 4294967295
    var v27 uint64 = v26
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v27, v27, v26)
    case "hello":
        if (65535 >= 42) {
            var v28 bool = false
            _ = v28
        }
    }
    if true {
        v29 := GenS1{F0: 255, F1: 255}
        if false {
            for fi30 := 0; fi30 < 5; fi30++ {
                v25 -= v25
                for fi31 := 0; fi31 < 19; fi31++ {
                    if false {
                        if (!false) {
                            var v32 uint64 = 1
                            var v33 bool = (!true)
                            _ = v32
                            _ = v33
                        } else {
                            v25 -= v25
                        }
                        var v34 uint32 = 242
                        if true {
                            delete(v2, "")
                            delete(v2, "a")
                            fmt.Println(v26, fi30, v27)
                        }
                        _ = v34
                    }
                    v27 = v26
                    switch "a" {
                    case "a":
                        v26 += v27
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = 787
                        }
                    }
                    _ = fi31
                }
                fmt.Println(v25, v26, v27)
                _ = fi30
            }
            if true {
                v35 := GenS2{F0: false, F1: 980, F2: 32767}
                _ = v35
            } else {
                v3 = append(v3, 65535)
                var v36 int16 = 0
                _ = v36
            }
        } else {
            fmt.Println(v25, v25, v25)
            v27 -= v27
            switch "a" {
            case "hello":
                var v37 bool = false
                _ = v37
            default:
                if true {
                    var v38 uint16 = 32767
                    v39 := GenS3{F0: 1}
                    _ = v38
                    _ = v39
                }
            }
        }
        _ = v29
    } else {
        var v40 int8 = 7
        _ = v40
    }
    var v41 float64 = v25
    for fi42 := 0; fi42 < 15; fi42++ {
        v43 := GenS4{F0: "", F1: 0, F2: 1}
        v44 := GenS5{F0: 4294967295, F1: 1}
        for fi45 := 0; fi45 < 12; fi45++ {
            v41 = (3.14159 - ((v41 * 0.0) + ((v41 / (float64(((((((float32(v41) * 0.0) - 1.5) * 3.14159) * float32(float64(1.5))) / 1.5) / (3.14159 + 0.0))) / v25)) * 1.5)))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
            _ = fi45
        }
        _ = fi42
        _ = v43
        _ = v44
    }
    v1 = append(v1, 2147483647)
    for fi46 := 0; fi46 < 16; fi46++ {
        var v47 uint = 621
        _ = fi46
        _ = v47
    }
    v3 = append(v3, 2147483648)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v41
}

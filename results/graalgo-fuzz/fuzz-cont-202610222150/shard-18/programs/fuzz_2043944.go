package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint }
func genFn0(p0 int64, p1 int8) uint8 {
    for fi4 := 0; fi4 < 1; fi4++ {
        if (42 > 1) {
            var v5 float64 = 1.5
            _ = v5
        }
        for fi6 := 0; fi6 < 7; fi6++ {
            switch "hello" {
            case "a":
                var v7 uint16 = 1
                _ = v7
            case "hello":
                p1 = (-(p1))
            case "\n\t":
                var v8 uint64 = 845
                _ = v8
            }
            v9 := GenS1{F0: 488, F1: 0, F2: -743}
            v10 := GenS2{F0: 0, F1: ((p1 << ((1) & 63)) < -45)}
            _ = fi6
            _ = v10
            _ = v9
        }
        if true {
            switch "foo bar" {
            case "\n\t":
                if ((312 > 2147483648) || (false && (!true))) {
                    var v11 int16 = -1
                    v11 = (v11 % ((v11) | 1))
                    _ = v11
                }
            case "foo bar":
                var v12 uint8 = 2
                _ = v12
            case "":
                var v13 uint64 = 422
                _ = v13
            }
        }
        _ = fi4
    }
    v14 := []int{-2147483648}
    var v15 uint64 = 828
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 2
    v3 := make(map[int]string)
    v2 += genFn0(0, 127)
    fmt.Println(v2)
    var v16 uint = 42
    v17 := GenS2{F0: 621, F1: (1 >= 261)}
    fmt.Println(v16, v16)
    if (370 < 731) {
        v18 := []int{127}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint16 = 42
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        var v22 uint32 = 916
        var v23 uint = v16
        _ = v22
        _ = v23
    }
    fmt.Println(v16, v2)
    switch -1 {
    case 4294967295:
        v24 := GenS2{F0: 1, F1: ((v2 & v2) >= (^(v2)))}
        _ = v24
    case 42:
        switch -41 {
        case 1:
            for fi25 := 0; fi25 < 1; fi25++ {
                for fi26 := 0; fi26 < 15; fi26++ {
                    if ((-2.25 / (1e10 + -2.25)) != 1e10) {
                        var v27 int32 = 42
                        v3[(fi25 ^ -9223372036854775808)] = "a"
                        _ = v27
                    }
                    var v28 bool = false
                    _ = fi26
                    _ = v28
                }
                _ = fi25
            }
        }
    case 127:
        v16 = v16
    default:
        var v29 uint64 = 4294967295
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 float64 = 1.5
        for fi32 := 0; fi32 < 18; fi32++ {
            v31 -= (((0.0 + v31) / v31) * (float64(float32(v31)) + v31))
            if false {
                delete(v1, 999)
                v2 = v2
                if ("a" < "a") {
                    switch fi30 {
                    case -2147483648:
                        var v33 int64 = 0
                        _ = v33
                    case 4294967295:
                        v16 -= 965
                    }
                } else {
                    var v34 int = (fi32 / (((fi32 + fi30)) | 1))
                    _ = v34
                }
            }
            var v35 uint16 = 2
            if (-2147483648 <= 1000) {
                var v36 string = (("a" + ("\n\t" + "a")) + "a")
                _ = v36
            }
            if true {
                if (v16 > (v16 << (((v16 / ((0) | 1))) & 63))) {
                    v2 -= v2
                } else {
                    if false {
                        v37 := GenS3{F0: 0}
                        _ = v37
                    } else {
                        switch (-(fi30)) {
                        case 42:
                            var v38 uint32 = 693
                            _ = v38
                        case -1:
                            var v39 uint32 = 1
                            _ = v39
                        }
                    }
                    if false {
                        var v40 string = ""
                        switch "hello" {
                        case "":
                            fmt.Println(v40, fi32)
                        default:
                            v41 := GenS4{F0: 2147483647}
                            _ = v41
                        }
                        switch v40 {
                        case "hello":
                            v42 := []int{fi32}
                            _ = v42
                        }
                        _ = v40
                    }
                    _ = len(v1)
                }
            }
            _ = fi32
            _ = v35
        }
        _ = fi30
        _ = v31
    }
    if true {
        v2 = (v2 - v2)
        if (65535 > 42) {
            v43 := GenS5{F0: 2147483647, F1: v16}
            _ = v43
        } else {
            v2 = (v2 << ((v16) & 63))
            var v44 uint8 = (v2 ^ 128)
            _ = v44
        }
        switch "a" {
        case "\n\t":
            var v45 uint8 = genFn0(655, 42)
            _ = v45
        case "hello":
            for fi46 := 0; fi46 < 14; fi46++ {
                v47 := []int{fi46, -116, fi46, fi46}
                var v48 uint8 = genFn0(-349, 1)
                _ = fi46
                _ = v47
                _ = v48
            }
        }
    }
    for fi49 := 0; fi49 < 20; fi49++ {
        v2 += v2
        _ = fi49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}

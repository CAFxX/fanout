package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]bool)
    var v3 float32 = (v1 / v1)
    var v4 string = ""
    fmt.Println(v1, v1, v4)
    switch v4 {
    case "":
        var v5 int = -89
        _ = v5
    case "foo bar":
        v6 := []int{127, 32768}
        _ = v6
    case "\n\t":
        switch 213 {
        case 128:
            v7 := GenS1{F0: 127, F1: 1, F2: 405}
            _ = v7
        case -930:
            _ = len(v2)
        default:
            fmt.Println(v4, v3)
        }
    default:
        if false {
            v1 = (v1 * v1)
        } else {
            fmt.Println(v1, v1)
        }
    }
    var v8 uint32 = 2147483648
    if false {
        if true {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint64 = 18446744073709551615
                v8 += 1
                _ = fi9
                _ = v10
            }
            switch v4 {
            case "a":
                var v11 int16 = 415
                _ = v11
            case "\n\t":
                v12 := GenS2{F0: 1, F1: false, F2: false}
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint8 = 255
                var v15 uint16 = 65535
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    } else {
        switch "\n\t" {
        case "":
            var v16 int8 = 111
            _ = v16
        case "a":
            var v17 uint = uint(v8)
            _ = v17
        default:
            v3 = v3
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := []int{fi18}
            fmt.Println(v4)
            _ = fi18
            _ = v19
        }
        fmt.Println(v1, v1, v1)
    }
    v20 := GenS3{F0: 399, F1: 255}
    _ = len(v2)
    switch 354 {
    case 32767:
        var v21 uint16 = 1
        _ = v21
    }
    fmt.Println(v1, v4)
    switch "hello" {
    case "":
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int32 = -309
            v24 := GenS4{F0: (!false)}
            var v25 string = (("a" + v4) + (v4 + "a"))
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    case "hello":
        fmt.Println(v1)
    case "a":
        if ((!(!false)) || true) {
            v4 = (v4 + "\n\t")
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 float32 = (v3 - v3)
                var v28 string = "\n\t"
                _ = fi26
                _ = v27
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                switch fi29 {
                case 0:
                    for fi30 := 0; fi30 < 17; fi30++ {
                        fmt.Println(v1)
                        v4 = (v4 + v4)
                        if false {
                            v31 := GenS1{F0: 88, F1: 42, F2: 4294967295}
                            var v32 float32 = -2.25
                            _ = v31
                            _ = v32
                        }
                        _ = fi30
                    }
                case 32768:
                    v8 = v8
                }
                var v33 float64 = -2.25
                v2[v4] = (false || true)
                _ = fi29
                _ = v33
            }
            switch -9223372036854775808 {
            case -279:
                v4 = v4
            case 0:
                fmt.Println(v3, v1)
            case -2147483648:
                switch -2147483648 {
                case 9223372036854775807:
                    _ = len(v2)
                default:
                    v34 := GenS1{F0: 42, F1: 0, F2: 4294967295}
                    _ = v34
                }
            }
            v8 += 1
        }
    default:
        v1 = (v1 * v3)
    }
    v35 := GenS2{F0: -71, F1: (1 == 65535), F2: ((-9223372036854775808 <= -98) && false)}
    v36 := GenS4{F0: false}
    for fi37 := 0; fi37 < 3; fi37++ {
        fmt.Println(v3, v4)
        if (!true) {
            var v38 int64 = -1
            fmt.Println(v1)
            _ = v38
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v8
}

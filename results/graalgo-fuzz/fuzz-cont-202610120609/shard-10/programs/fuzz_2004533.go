package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    if (!true) {
        var v5 float64 = -2.25
        var v6 string = ("hello" + v3)
        _ = v5
        _ = v6
    }
    switch -469 {
    case -741:
        fmt.Println(v3, v2)
    default:
        fmt.Println(v3)
    }
    fmt.Println(v3)
    for fi7 := 0; fi7 < 4; fi7++ {
        switch v3 {
        case "foo bar":
            fmt.Println(v2, fi7)
        case "hello":
            if true {
                _ = len(v4)
            }
        case "a":
            fmt.Println(v3)
        default:
            var v8 int64 = -393
            _ = v8
        }
        fmt.Println(v2, fi7, fi7)
        fmt.Println(fi7)
        _ = fi7
    }
    if (!false) {
        var v9 uint32 = 712
        v10 := GenS1{F0: 90, F1: 411}
        _ = v10
        _ = v9
    } else {
        switch ("foo bar" + "foo bar") {
        case "":
            v2 = v2
        case "foo bar":
            v2 -= v2
        case "hello":
            v11 := GenS1{F0: -1, F1: 42}
            _ = v11
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 bool = (0.1 == (3.14159 + ((0.0 + float64(float32(1e10))) + 1.5)))
                    var v15 uint32 = 851
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        }
        var v16 float64 = 1.5
        _ = v16
    }
    v3 = v3
    if (!false) {
        v17 := GenS2{F0: 128, F1: -1}
        v18 := GenS3{F0: 10}
        _ = v17
        _ = v18
    } else {
        v2 += v2
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS1{F0: 508, F1: 42}
            _ = len(v4)
            switch -9223372036854775808 {
            case 1:
                v21 := GenS3{F0: -17}
                _ = v21
            case 128:
                if true {
                    v22 := []int{250, (-(fi19)), fi19}
                    v23 := GenS3{F0: 0}
                    _ = v22
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        var v25 int16 = 796
                        fmt.Println(v3, v2)
                        fmt.Println(v2, fi24)
                        _ = fi24
                        _ = v25
                    }
                }
            case 127:
                v26 := GenS2{F0: 109, F1: 588}
                _ = v26
            default:
                for fi27 := 0; fi27 < 1; fi27++ {
                    v1["foo bar"] = 2147483648
                    v3 = v3
                    var v28 uint32 = 4294967295
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi19
            _ = v20
        }
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(v3, v3, fi29)
        var v30 uint32 = 65535
        switch "\n\t" {
        case "a":
            v30 += v30
        case "foo bar":
            fmt.Println(fi29)
        case "":
            var v31 int64 = -49
            _ = v31
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}

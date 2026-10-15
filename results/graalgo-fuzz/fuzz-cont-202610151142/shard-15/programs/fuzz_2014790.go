package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 int32 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    v1 := []uint32{1}
    var v2 uint16 = 674
    v3 := []string{""}
    var v4 float32 = (0.1 / (3.14159 + 1.5))
    v5 := make(map[uint32]string)
    switch 365 {
    case 42:
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                var v7 uint = 2
                if (v4 > v4) {
                    v8 := GenS1{F0: (^(v7)), F1: 43}
                    var v9 int64 = -523
                    _ = v8
                    _ = v9
                }
                _ = v7
            } else {
                var v10 int64 = 42
                v11 := GenS1{F0: 42, F1: 42}
                fmt.Println(v4, v4)
                _ = v10
                _ = v11
            }
            switch fi6 {
            case 128:
                switch fi6 {
                case 4294967295:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        switch fi12 {
                        case 0:
                            v13 := GenS1{F0: 18446744073709551615, F1: uint64(v2)}
                            _ = v13
                        case 9223372036854775807:
                            var v14 string = ""
                            _ = v14
                        }
                        var v15 bool = (255 <= 2)
                        var v16 string = ("foo bar" + "a")
                        _ = fi12
                        _ = v15
                        _ = v16
                    }
                case -13:
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v18 := GenS2{F0: 42, F1: 286, F2: 4294967295}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            case 1:
                for fi19 := 0; fi19 < 10; fi19++ {
                    fmt.Println(fi19, v2, fi19)
                    var v20 uint64 = 2
                    v21 := GenS3{F0: -93, F1: 2, F2: 431}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case 32767:
                if (float32((0.1 + 0.0)) < v4) {
                    v2 += (v2 / (((v2 / ((v2) | 1))) | 1))
                    v4 += float32((1e10 - 0.0))
                    v2 = (v2 % ((32767) | 1))
                } else {
                    var v22 int = fi6
                    v2 = (v2 << ((1) & 63))
                    _ = v22
                }
            default:
                v23 := GenS3{F0: -1, F1: 255, F2: 2147483647}
                _ = v23
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint16 = (v2 * 32767)
                var v26 int32 = -56
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi6
        }
    case 32768:
        fmt.Println(v4)
    case 9223372036854775807:
        fmt.Println(v4, v4)
    }
    v2 = 65535
    for fi27 := 0; fi27 < 18; fi27++ {
        if true {
            v3 = append(v3, "")
            v28 := GenS2{F0: 2147483647, F1: 42, F2: 2147483647}
            var v29 string = ""
            _ = v28
            _ = v29
        } else {
            v4 -= 1e10
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 bool = true
                _ = fi30
                _ = v31
            }
            var v32 int8 = 43
            _ = v32
        }
        _ = fi27
    }
    delete(v5, 1)
    v4 = (v4 + ((v4 + 1e10) / (v4 / (float32((float64(((v4 - v4) / (v4 * (v4 + float32(1.5))))) - 1e10)) / (float32(-2.25) * ((((0.0 * -2.25) + (((v4 - float32(0.0)) / 0.1) - (v4 - 0.1))) + ((v4 - v4) * 1e10)) / (((-2.25 / 1.0) + v4) + v4)))))))
    v2 = (v2 & v2)
    var v33 int = 2147483647
    v33 = v33
    var v34 int32 = 1
    switch "\n\t" {
    case "\n\t":
        switch "a" {
        case "":
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS4{F0: "a", F1: 42}
                for fi37 := 0; fi37 < 13; fi37++ {
                    var v38 int = (-(v33))
                    _ = fi37
                    _ = v38
                }
                delete(v5, 42)
                _ = fi35
                _ = v36
            }
        }
    default:
        var v39 uint16 = (v2 + v2)
        _ = v39
    }
    if false {
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 int = v33
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi40
            _ = v41
        }
        v42 := GenS3{F0: -1, F1: 886, F2: -1}
        _ = v42
    }
    if true {
        v43 := GenS3{F0: 0, F1: 18, F2: 42}
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}

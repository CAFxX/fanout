package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 bool }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    delete(v3, "hello")
    var v4 float64 = 1e10
    if false {
        if false {
            v5 := GenS1{F0: true}
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint = 572
                v8 := v2[403]
                v9 := []int{}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 int8 = (-(v1))
            _ = v10
            _ = v5
        } else {
            fmt.Println(v1)
            v11 := GenS2{F0: 9223372036854775807, F1: 1, F2: (-(v1))}
            v1 -= v1
            _ = v11
        }
    } else {
        v12 := GenS3{F0: "hello"}
        _ = v12
    }
    var v13 int16 = -32768
    v14 := GenS2{F0: 224, F1: 32768, F2: (v1 + 57)}
    v15 := GenS1{F0: true}
    v4 = ((((v4 + float64(0.0)) / (3.14159 / (0.0 - v4))) * v4) - (0.0 * float64(0.0)))
    var v16 string = "\n\t"
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint32 = 65535
        if false {
            if (!false) {
                var v19 string = (v16 + (v16 + v16))
                _ = v19
            }
        }
        _ = fi17
        _ = v18
    }
    switch 128 {
    case 1:
        var v20 int8 = v1
        _ = v20
    case 563:
        v1 = (v1 * (v1 / ((79) | 1)))
    default:
        for fi21 := 0; fi21 < 4; fi21++ {
            switch "foo bar" {
            case "":
                v22 := GenS3{F0: v16}
                _ = v22
            case "\n\t":
                var v23 int16 = (v13 >> ((42) & 63))
                _ = v23
            case "hello":
                for fi24 := 0; fi24 < 4; fi24++ {
                    if (!true) {
                        var v25 float64 = (v4 / 3.14159)
                        v26 := GenS4{F0: ((0 != 2) && (true || true)), F1: 9223372036854775807, F2: 1}
                        _ = v25
                        _ = v26
                    } else {
                        v1 += v1
                        v16 = ((v16 + v16) + v16)
                    }
                    fmt.Println(fi21, v13)
                    delete(v3, "a")
                    switch "\n\t" {
                    case "a":
                        switch ((v16 + v16) + ("a" + "")) {
                        case "a":
                            if true {
                                if false {
                                    v13 += -1
                                    var v27 int = 127
                                    _ = v27
                                } else {
                                    if (true || ((v1 * 42) > v1)) {
                                        var v28 int16 = v13
                                        switch fi21 {
                                        case 145:
                                            var v29 uint32 = 0
                                            _ = v29
                                        case 128:
                                            _ = len(v2)
                                        default:
                                            v4 += float64(0.0)
                                        }
                                        _ = v28
                                    } else {
                                        if true {
                                            if (!((0.0 + float32(float64(float32(v4)))) <= (0.0 * 1e10))) {
                                                if (484 <= 0) {
                                                    v30 := GenS3{F0: v16}
                                                    if true {
                                                        v2[564] = (fi21 % (((fi24 / (((fi24 << ((1) & 63))) | 1))) | 1))
                                                        fmt.Println(v16)
                                                    }
                                                    v13 = int16(v1)
                                                    _ = v30
                                                } else {
                                                    v1 += 42
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        switch v16 {
                                                        case "":
                                                            delete(v3, "a")
                                                        case "foo bar":
                                                            var v31 float64 = v4
                                                            _ = v31
                                                        case "a":
                                                            v1 = (v1 % ((v1) | 1))
                                                        }
                                                    default:
                                                        var v32 int64 = -1
                                                        _ = v32
                                                    }
                                                }
                                            } else {
                                                var v33 uint64 = 1000
                                                if false {
                                                    v34 := GenS1{F0: (0 < (^(v33)))}
                                                    if true {
                                                        if true {
                                                            fmt.Println(fi21)
                                                        }
                                                    }
                                                    _ = v34
                                                } else {
                                                    var v35 int64 = 0
                                                    _ = v35
                                                }
                                                _ = v33
                                            }
                                            var v36 string = ("hello" + (("hello" + "\n\t") + (v16 + v16)))
                                            v13 = (v13 >> ((255) & 63))
                                            _ = v36
                                        }
                                        fmt.Println(v16, fi21, v1)
                                    }
                                }
                                v13 -= (^(v13))
                                v37 := GenS5{F0: 1, F1: (^(fi24))}
                                _ = v37
                            } else {
                                v1 -= -1
                                v16 = v16
                                v16 = "hello"
                            }
                        case "":
                            v1 = v1
                        default:
                            v13 = 32767
                        }
                    case "\n\t":
                        var v38 uint64 = 476
                        _ = v38
                    case "":
                        if false {
                            var v39 uint16 = 1
                            _ = v39
                        }
                    }
                    v40 := []int{fi24, (fi21 >> ((65535) & 63)), (fi24 ^ int(v13)), 127}
                    _ = fi24
                    _ = v40
                }
            default:
                var v41 int64 = 42
                _ = v41
            }
            var v42 uint = 42
            _ = fi21
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        v44 := GenS5{F0: 42, F1: fi43}
        _ = fi43
        _ = v44
    }
    switch (v16 + (v16 + v16)) {
    case "foo bar":
        delete(v3, v16)
    case "a":
        switch 32767 {
        case -1:
            v16 = v16
        default:
            switch 42 {
            case -294:
                v45 := GenS1{F0: false}
                _ = v45
            case 32767:
                _ = len(v3)
            case 108:
                var v46 uint64 = 4294967295
                _ = v46
            default:
                v13 = v13
            }
        }
    case "hello":
        var v47 int8 = (v1 / (((v1 << ((65535) & 63))) | 1))
        _ = v47
    default:
        v48 := GenS5{F0: 128, F1: 4294967295}
        _ = v48
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        switch -2147483648 {
        case 663:
            if (0 <= v13) {
                switch 42 {
                case 9223372036854775807:
                    var v50 float32 = 0.1
                    _ = v50
                case 2147483647:
                    var v51 uint64 = 18446744073709551615
                    _ = v51
                }
                v1 += v1
            } else {
                fmt.Println(v4, v13, v1)
                delete(v3, "")
                switch (fi49 + fi49) {
                case 986:
                    v13 = v13
                case 127:
                    v1 = 42
                default:
                    v4 += (v4 + 1.5)
                }
            }
        case 148:
            var v52 int8 = 127
            _ = v52
        case 4294967295:
            v53 := []int{(fi49 + (-(fi49)))}
            _ = v53
        }
        _ = fi49
    }
    for fi54 := 0; fi54 < 15; fi54++ {
        v4 = (-2.25 - v4)
        _ = fi54
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        var v56 uint32 = 42
        if true {
            for fi57 := 0; fi57 < 8; fi57++ {
                v58 := GenS2{F0: -325, F1: fi55, F2: int8(fi55)}
                if (false || (!true)) {
                    v59 := GenS6{F0: false}
                    _ = v59
                }
                _ = fi57
                _ = v58
            }
        } else {
            for fi60 := 0; fi60 < 13; fi60++ {
                if (-2.25 > ((1e10 * 1.5) + (0.0 / (1.5 * ((-2.25 * 0.1) - 1.5))))) {
                    var v61 float32 = ((float32(3.14159) - (3.14159 - float32(v4))) - (((float32(0.1) * (((0.1 * -2.25) + (((3.14159 * 3.14159) + -2.25) - float32((v4 + (v4 * (v4 - 0.0)))))) - 1e10)) - -2.25) + 1e10))
                    fmt.Println(fi60, v16, v4)
                    _ = v61
                }
                v62 := GenS4{F0: false, F1: 822, F2: 432}
                _ = fi60
                _ = v62
            }
            v13 += v13
        }
        _ = fi55
        _ = v56
    }
    var v63 float32 = (((float32((v4 * (((((v4 + v4) - v4) * 1.5) - v4) + (0.1 - 1.5)))) - (0.0 / (1e10 + (float32(3.14159) * -2.25)))) * float32(((v4 - 1.5) / v4))) - (0.0 - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}

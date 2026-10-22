package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int; F1 int8; F2 bool }
type GenS5 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS6 struct { F0 string; F1 int64; F2 int32 }

func main() {
    var v1 float32 = (1.5 / -2.25)
    var v2 uint8 = 127
    var v3 int32 = 1
    var v4 string = ""
    var v5 uint8 = (v2 << ((18446744073709551615) & 63))
    v6 := make(map[uint32]bool)
    var v7 uint = 2
    v8 := GenS1{F0: v7, F1: (v7 ^ 913), F2: 4294967295}
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := GenS1{F0: v7, F1: 0, F2: fi9}
                _ = fi10
                _ = v11
            }
        } else {
            if (!true) {
                var v12 float64 = (0.1 * -2.25)
                v4 = v4
                var v13 int = -390
                _ = v12
                _ = v13
            } else {
                fmt.Println(v3, v1)
                switch v4 {
                case "":
                    switch (fi9 | 127) {
                    case -1:
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 uint = v7
                            var v16 string = ("" + v4)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                case "foo bar":
                    fmt.Println(v1)
                case "hello":
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 float32 = 1.5
                        _ = fi17
                        _ = v18
                    }
                default:
                    switch (fi9 | 243) {
                    case 332:
                        v19 := GenS2{F0: 0}
                        _ = v19
                    case 0:
                        if false {
                            if false {
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    if (209 >= 4294967295) {
                                        v1 += (0.1 - 3.14159)
                                    }
                                    if false {
                                        v21 := []int{}
                                        _ = v21
                                    }
                                    v5 = 210
                                    _ = fi20
                                }
                                if false {
                                    v22 := []int{}
                                    fmt.Println(v4, v5)
                                    _ = v22
                                } else {
                                    switch 127 {
                                    case 0:
                                        v23 := GenS2{F0: 1000}
                                        _ = v23
                                    case -2147483648:
                                        for fi24 := 0; fi24 < 12; fi24++ {
                                            if (1 <= -9223372036854775808) {
                                                v4 = ((v4 + (v4 + (v4 + v4))) + ((("foo bar" + v4) + v4) + ("foo bar" + "hello")))
                                                if (false && (42 == (v5 - 244))) {
                                                    v3 += (v3 & v3)
                                                    _ = len(v6)
                                                }
                                            } else {
                                                var v25 int16 = 1000
                                                if false {
                                                    if true {
                                                        switch (fi9 / ((fi9) | 1)) {
                                                        case 0:
                                                            var v26 string = v4
                                                            var v27 int16 = 38
                                                            if false {
                                                                v26 = "a"
                                                            } else {
                                                                if ((9223372036854775807 < -9223372036854775808) && (!(!true))) {
                                                                    if true {
                                                                        v28 := []int{fi9, (fi24 - -95), fi24}
                                                                        v26 = ("hello" + v4)
                                                                        fmt.Println(v1, v3, v3)
                                                                        _ = v28
                                                                    }
                                                                    v7 += 4294967295
                                                                    v29 := GenS2{F0: 1000}
                                                                    var v30 int8 = 42
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                            }
                                                            _ = v26
                                                            _ = v27
                                                        case 4294967295:
                                                            fmt.Println(v7, fi24)
                                                        default:
                                                            var v31 float32 = v1
                                                            _ = v31
                                                        }
                                                        v32 := []int{int(v3), -752, int(v25)}
                                                        _ = v32
                                                    }
                                                } else {
                                                    v33 := GenS3{F0: 809, F1: -9223372036854775808}
                                                    if (!true) {
                                                        if true {
                                                            switch ((v4 + "foo bar") + v4) {
                                                            case "hello":
                                                                if (598 == 2147483648) {
                                                                    v34 := []int{fi24, fi24}
                                                                    _ = v34
                                                                } else {
                                                                    var v35 int64 = 9223372036854775807
                                                                    var v36 uint64 = 71
                                                                    _ = v35
                                                                    _ = v36
                                                                }
                                                            default:
                                                                var v37 float32 = (((v1 / v1) * ((v1 - (0.0 * v1)) + (float32(3.14159) - 1.5))) / ((1e10 * v1) - ((v1 * v1) * v1)))
                                                                _ = v37
                                                            }
                                                        }
                                                        v38 := []int{(fi9 & (fi24 + int(v5))), fi24, (fi24 + 42), (-(fi9))}
                                                        _ = v38
                                                    }
                                                    fmt.Println(v25, v5, fi9)
                                                    v5 = 255
                                                    switch fi9 {
                                                    case 128:
                                                        var v39 bool = (!true)
                                                        _ = v39
                                                    case 4294967295:
                                                        var v40 uint16 = 0
                                                        _ = v40
                                                    case -629:
                                                        v41 := GenS4{F0: fi9, F1: 0, F2: false}
                                                        _ = v41
                                                    default:
                                                        v7 -= v7
                                                    }
                                                    _ = v33
                                                }
                                                var v42 int8 = 127
                                                _ = v25
                                                _ = v42
                                            }
                                            v1 = (0.1 + v1)
                                            v43 := GenS1{F0: v7, F1: (v7 / (((v7 - 255)) | 1)), F2: 128}
                                            _ = fi24
                                            _ = v43
                                        }
                                    case 128:
                                        switch (fi9 | (fi9 - (fi9 >> ((v7) & 63)))) {
                                        case 590:
                                            var v44 int64 = -265
                                            _ = v44
                                        case -1:
                                            v5 += (v5 - 128)
                                        case 1:
                                            v5 = (v2 << (((v7 / (((v7 | 42)) | 1))) & 63))
                                        }
                                    }
                                }
                            }
                            for fi45 := 0; fi45 < 4; fi45++ {
                                v46 := GenS2{F0: 42}
                                switch "a" {
                                case "\n\t":
                                    v47 := GenS1{F0: (v7 | v7), F1: v7, F2: (fi45 | 56)}
                                    _ = v47
                                case "hello":
                                    v3 = int32(fi45)
                                case "":
                                    if false {
                                        var v48 uint32 = 42
                                        _ = v48
                                    }
                                }
                                v49 := GenS4{F0: fi9, F1: 9, F2: true}
                                _ = fi45
                                _ = v46
                                _ = v49
                            }
                        }
                    default:
                        fmt.Println(v2)
                    }
                }
                v50 := GenS5{F0: (v2 % (((v2 * v2)) | 1)), F1: 0, F2: 2}
                _ = v50
            }
            fmt.Println(v7, v1)
        }
        if false {
            v4 = (((v4 + ("" + v4)) + ("foo bar" + (v4 + "a"))) + v4)
            var v51 uint32 = 556
            _ = v51
        } else {
            v52 := []int{}
            _ = v52
        }
        v53 := v6[8]
        _ = fi9
        _ = v53
    }
    var v54 int = 32768
    var v55 float64 = 1.5
    v56 := GenS2{F0: 784}
    var v57 uint64 = 0
    if false {
        delete(v6, 42)
        v58 := GenS6{F0: v4, F1: 42, F2: (v3 >> (((v7 - (v7 / ((2) | 1)))) & 63))}
        _ = v58
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        v54 = fi59
        _ = fi59
    }
    var v60 float32 = (0.0 * v1)
    v60 -= v60
    for fi61 := 0; fi61 < 5; fi61++ {
        v4 = "hello"
        _ = fi61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v60
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int8; F2 int32 }
type GenS5 struct { F0 bool; F1 bool; F2 uint64 }
type GenS6 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 2
    var v3 uint16 = 65535
    var v4 uint16 = v3
    if true {
        var v5 int64 = 9223372036854775807
        _ = v5
    } else {
        if false {
            switch "hello" {
            case "":
                v4 += uint16(v2)
            case "a":
                if false {
                    v6 := GenS1{F0: 128, F1: (v3 * (v4 + (v3 ^ 859))), F2: -9223372036854775808}
                    v4 += (v4 - (v4 << ((v2) & 63)))
                    v7 := []int{}
                    _ = v6
                    _ = v7
                } else {
                    v8 := GenS2{F0: 1, F1: 2147483647, F2: 1}
                    var v9 int = 4294967295
                    switch "" {
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "\n\t":
                        delete(v1, (v9 & (v9 + 0)))
                    }
                    _ = v8
                    _ = v9
                }
            case "\n\t":
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch (fi10 + fi10) {
                    case 128:
                        if true {
                            v11 := GenS1{F0: 1, F1: v3, F2: -9223372036854775808}
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 11; fi12++ {
                                switch "hello" {
                                case "":
                                    var v13 int = (-(fi12))
                                    _ = v13
                                case "hello":
                                    var v14 uint8 = 2
                                    _ = v14
                                case "\n\t":
                                    var v15 string = "foo bar"
                                    _ = v15
                                default:
                                    if ((fi12 | (fi10 / ((fi10) | 1))) >= int(v3)) {
                                        fmt.Println(v2)
                                        switch "a" {
                                        case "":
                                            v16 := GenS3{F0: 1}
                                            _ = v16
                                        case "hello":
                                            if true {
                                                switch (fi12 | fi10) {
                                                case -1:
                                                    switch (-(fi12)) {
                                                    case -281:
                                                        v2 += 42
                                                    case 9223372036854775807:
                                                        v17 := GenS1{F0: 1, F1: (v4 + (v3 >> ((v2) & 63))), F2: -9223372036854775808}
                                                        _ = v17
                                                    case 42:
                                                        if false {
                                                            _ = len(v1)
                                                            v18 := []int{fi12}
                                                            _ = v18
                                                        }
                                                    default:
                                                        v4 += (^(v3))
                                                    }
                                                case -719:
                                                    switch (^(fi12)) {
                                                    case 0:
                                                        _ = len(v1)
                                                    }
                                                case 32768:
                                                    v19 := GenS2{F0: fi12, F1: 199, F2: fi12}
                                                    _ = v19
                                                }
                                                v20 := v1[32767]
                                                _ = v20
                                            } else {
                                                var v21 uint = (v2 % (((v2 % (((v2 - 1)) | 1))) | 1))
                                                var v22 int8 = 23
                                                if false {
                                                    var v23 uint8 = 1
                                                    _ = v23
                                                }
                                                if false {
                                                    var v24 int32 = int32(v4)
                                                    v25 := GenS3{F0: 810}
                                                    v26 := GenS4{F0: v2, F1: (v22 << (((v21 & v2)) & 63)), F2: 42}
                                                    _ = v24
                                                    _ = v25
                                                    _ = v26
                                                }
                                                v27 := []int{(fi10 >> (((v21 * (v2 ^ (v21 / ((v2) | 1))))) & 63)), (fi10 / (((fi12 | fi10)) | 1))}
                                                _ = v21
                                                _ = v22
                                                _ = v27
                                            }
                                        case "a":
                                            if ((v3 % ((v4) | 1)) == 2) {
                                                v3 = 2
                                                v28 := v1[fi12]
                                                var v29 uint = v2
                                                _ = v28
                                                _ = v29
                                            } else {
                                                switch fi10 {
                                                case -537:
                                                    v4 -= 65535
                                                }
                                                if true {
                                                    var v30 uint8 = 255
                                                    _ = v30
                                                }
                                                fmt.Println(v2, v3, fi10)
                                            }
                                        default:
                                            v31 := []int{(fi10 + (fi10 | (^(fi12))))}
                                            _ = v31
                                        }
                                    }
                                }
                                var v32 uint64 = 18446744073709551615
                                fmt.Println(v32)
                                _ = fi12
                                _ = v32
                            }
                            var v33 uint8 = 127
                            v4 = 32767
                            _ = v33
                        }
                    case -807:
                        fmt.Println(v2, fi10, v3)
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        fmt.Println(fi10, v2, v2)
                        v2 -= 351
                        var v35 float64 = (1.5 + float64(1e10))
                        _ = fi34
                        _ = v35
                    }
                    if true {
                        for fi36 := 0; fi36 < 14; fi36++ {
                            v37 := GenS1{F0: 238, F1: (v3 + (v4 >> ((v2) & 63))), F2: 42}
                            _ = fi36
                            _ = v37
                        }
                        var v38 float32 = (0.1 - 0.0)
                        for fi39 := 0; fi39 < 14; fi39++ {
                            v40 := GenS3{F0: -1}
                            switch (^(fi39)) {
                            case 0:
                                var v41 int16 = 1000
                                _ = v41
                            case -2147483648:
                                v42 := []int{(fi39 - fi10)}
                                _ = v42
                            case 2147483647:
                                switch fi39 {
                                case -589:
                                    var v43 uint32 = 672
                                    _ = v43
                                case -1:
                                    v44 := GenS3{F0: -505}
                                    _ = v44
                                case 42:
                                    var v45 uint = 255
                                    _ = v45
                                }
                            default:
                                v46 := GenS4{F0: 410, F1: 0, F2: 70}
                                _ = v46
                            }
                            v47 := GenS1{F0: 177, F1: v4, F2: -1}
                            fmt.Println(fi10, v3)
                            _ = fi39
                            _ = v40
                            _ = v47
                        }
                        _ = v38
                    }
                    _ = fi10
                }
            default:
                var v48 int = 722
                _ = v48
            }
            for fi49 := 0; fi49 < 2; fi49++ {
                var v50 int = fi49
                v51 := GenS3{F0: 1}
                fmt.Println(v2, v2)
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        var v52 int16 = -1
        v2 = v2
        _ = v52
    }
    if false {
        if false {
            if ((v2 | v2) <= v2) {
                delete(v1, -1)
            }
        }
        v53 := []int{-1, int(v2), 355}
        v54 := GenS3{F0: int16(v2)}
        _ = v53
        _ = v54
    }
    v4 += v4
    v55 := GenS5{F0: false, F1: false, F2: 2}
    var v56 int = 2147483647
    fmt.Println(v2, v4)
    var v57 bool = false
    if v57 {
        switch (v56 >> ((v2) & 63)) {
        case 1:
            if (255 <= v4) {
                v57 = (true || ((float64(((((0.0 / 1.0) * -2.25) + 0.1) + float32(1e10))) <= ((1e10 * 1e10) / 1e10)) && (v3 >= 1)))
                if ((566 < 4294967295) && v57) {
                    var v58 float64 = float64(0.0)
                    for fi59 := 0; fi59 < 7; fi59++ {
                        v60 := GenS6{F0: true, F1: -62}
                        fmt.Println(v3)
                        fmt.Println(fi59, v58)
                        _ = fi59
                        _ = v60
                    }
                    fmt.Println(v56, v3)
                    _ = v58
                }
                var v61 int8 = 127
                _ = v61
            } else {
                v2 = v2
                v3 += (v4 << ((v2) & 63))
                if (-1 < -23) {
                    for fi62 := 0; fi62 < 5; fi62++ {
                        fmt.Println(v56)
                        _ = fi62
                    }
                    v57 = (0.1 == (3.14159 + 0.0))
                }
            }
        default:
            v63 := []int{v56}
            _ = v63
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
}

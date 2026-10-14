package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint32) uint {
    var v5 int = -655
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int16 = 1
        p0 += 1
        var v8 uint64 = 604
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v5 = v5
        fmt.Println(p0)
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    return 65535
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{-9223372036854775808}
    var v3 bool = (!true)
    var v4 int32 = 904
    if true {
        var v11 uint16 = (v1 - (v1 | 65535))
        if v3 {
            v11 = (v1 << ((805) & 63))
            var v12 int16 = 1
            switch 0 {
            case -193:
                if v3 {
                    if (!v3) {
                        var v13 string = "a"
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 128}
                        _ = v14
                    }
                } else {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := GenS1{F0: 2}
                        _ = fi15
                        _ = v16
                    }
                    var v17 bool = v3
                    v11 -= (v1 * (v1 / ((277) | 1)))
                    _ = v17
                }
            case 4294967295:
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 bool = v3
                    v12 -= v12
                    _ = fi18
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := GenS1{F0: 127}
                    var v22 float32 = 0.1
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            _ = v12
        }
        if (v3 && (-2.25 > float32(0.1))) {
            v4 = (v4 & 160)
            if (true || (96 == 32767)) {
                _ = len(v2)
                var v23 uint32 = 1
                _ = v23
            }
        }
        _ = v11
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 += (v1 + (v1 | 1))
        var v24 int64 = 9223372036854775807
        _ = v24
    }
    switch "" {
    case "foo bar":
        fmt.Println(v4, v3)
    default:
        switch "" {
        case "a":
            var v25 uint8 = 0
            _ = v25
        case "foo bar":
            var v26 int32 = (v4 & (v4 * v4))
            _ = v26
        default:
            var v27 int64 = 0
            _ = v27
        }
    }
    var v28 uint = 42
    var v29 int16 = 0
    v4 = (^(v4))
    if v3 {
        v30 := GenS2{F0: (v28 + v28), F1: 127, F2: 2}
        _ = v30
    } else {
        if (v3 && (!((v1 | (^(v1))) == (v1 % (((v1 << ((63) & 63))) | 1))))) {
            var v31 string = ""
            var v32 uint64 = 1
            if true {
                if true {
                    var v33 uint = genFn0(1)
                    var v34 uint16 = v1
                    _ = v33
                    _ = v34
                }
            }
            _ = v31
            _ = v32
        } else {
            _ = len(v2)
            v28 = v28
        }
        v2 = append(v2, 128)
        var v35 uint = v28
        _ = v35
    }
    var v36 uint32 = uint32(v1)
    v37 := GenS3{F0: 42, F1: 1, F2: -1}
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 bool = v3
        v36 -= (v36 & (v36 + uint32(v1)))
        _ = fi38
        _ = v39
    }
    var v40 uint8 = 74
    var v41 uint8 = (v40 - (v40 / (((v40 - v40)) | 1)))
    if ((42 <= int64(v1)) || v3) {
        v42 := GenS3{F0: 100, F1: 1, F2: 0}
        fmt.Println(v36, v36)
        v2 = append(v2, 9223372036854775807)
        _ = v42
    } else {
        var v43 uint = (v28 * (v28 - 705))
        var v44 int8 = 0
        if false {
            v45 := GenS4{F0: 183, F1: v40}
            for fi46 := 0; fi46 < 5; fi46++ {
                switch fi46 {
                case -1:
                    switch fi46 {
                    case 0:
                        var v47 int64 = 1
                        _ = v47
                    case 32767:
                        if v3 {
                            fmt.Println(fi46, v41, v36)
                        }
                    case 32768:
                        v29 -= int16(v1)
                    default:
                        v41 -= 42
                    }
                case -2147483648:
                    var v48 float32 = -2.25
                    _ = v48
                default:
                    var v49 float64 = (3.14159 / (1e10 / float64(0.1)))
                    _ = v49
                }
                if ((^(v40)) > 1) {
                    switch "" {
                    case "hello":
                        v50 := GenS1{F0: 212}
                        _ = v50
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "a":
                            v1 = v1
                        }
                    default:
                        var v51 int16 = v29
                        _ = v51
                    }
                }
                _ = fi46
            }
            _ = v45
        }
        _ = v43
        _ = v44
    }
    switch "\n\t" {
    case "":
        v41 -= v40
    case "foo bar":
        for fi52 := 0; fi52 < 14; fi52++ {
            v41 += v40
            switch (fi52 | (fi52 << ((v28) & 63))) {
            case 59:
                var v53 bool = true
                _ = v53
            case -42:
                var v54 uint64 = 95
                _ = v54
            }
            fmt.Println(v4, v28, v3)
            _ = fi52
        }
    case "\n\t":
        v55 := GenS4{F0: 9223372036854775807, F1: (v40 / (((^(v41))) | 1))}
        _ = v55
    }
    var v56 uint32 = uint32(v28)
    switch 9223372036854775807 {
    case 128:
        switch 32768 {
        case -482:
            var v57 uint64 = 98
            _ = v57
        case -1:
            v58 := GenS5{F0: v36}
            _ = v58
        case -2147483648:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v59 uint32 = (v56 & (v36 + v36))
            _ = v59
        }
    case 501:
        var v60 int64 = int64(v28)
        _ = v60
    default:
        v29 = (v29 + v29)
    }
    v2 = append(v2, -376)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v56
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int32; F1 uint; F2 uint16 }
func genFn0() bool {
    var v4 string = "foo bar"
    switch 4294967295 {
    case 32767:
        var v5 uint16 = 255
        _ = v5
    case -2147483648:
        fmt.Println(v4)
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint16 = 613
        var v8 int = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    v4 = v4
    _ = v4
    return false
}
func genFn1(p0 int16) bool {
    v9 := GenS1{F0: 9223372036854775807}
    if ("" >= "a") {
        p0 = (p0 ^ (p0 << ((985) & 63)))
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS2{F0: 569}
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 int32 = 1
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        var v14 uint64 = 4294967295
        switch -2147483648 {
        case -509:
            if false {
                fmt.Println(v14, v14)
                v14 += 6
                var v15 uint = 4294967295
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 10; fi16++ {
                if (((1e10 + (float32((0.1 * ((1.5 / float64(((3.14159 - float32(3.14159)) - 1e10))) * 0.1))) * 3.14159)) <= -2.25) || true) {
                    if (!((32767 >= 2) || (!(226 != 2)))) {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            v14 += (v14 + 1)
                            v14 = v14
                            fmt.Println(fi17, fi16, v14)
                            switch (fi16 << ((874) & 63)) {
                            case 4294967295:
                                v18 := GenS2{F0: v14}
                                _ = v18
                            default:
                                v19 := GenS2{F0: (v14 << ((uint(v14)) & 63))}
                                _ = v19
                            }
                            var v20 int32 = 930
                            switch (("a" + "") + "hello") {
                            case "a":
                                var v21 int = (fi17 * 9223372036854775807)
                                _ = v21
                            case "hello":
                                var v22 int32 = v20
                                _ = v22
                            case "foo bar":
                                v23 := GenS3{F0: 6, F1: 457, F2: 25}
                                _ = v23
                            default:
                                switch fi17 {
                                case 766:
                                    v24 := GenS4{F0: int64(p0), F1: (fi16 >> ((255) & 63)), F2: -1}
                                    _ = v24
                                case -561:
                                    var v25 uint8 = 147
                                    _ = v25
                                case -1:
                                    p0 -= 179
                                default:
                                    v14 += v14
                                }
                            }
                            var v26 uint32 = 2147483647
                            if (42 == v14) {
                                v26 += (v26 << ((576) & 63))
                                v27 := []int{int(v14)}
                                v14 -= v14
                                _ = v27
                            }
                            fmt.Println(v26, v26, p0)
                            _ = fi17
                            _ = v20
                            _ = v26
                        }
                        v28 := []int{}
                        _ = v28
                    }
                    v14 = v14
                    var v29 float64 = 3.14159
                    _ = v29
                } else {
                    v30 := []int{fi16, fi16, (fi16 & fi16)}
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v32 := GenS5{F0: -9223372036854775808, F1: v14}
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                v33 := GenS6{F0: 239, F1: 65535, F2: 42}
                if false {
                    v34 := []int{2147483647, (fi16 % ((fi16) | 1)), fi16, -660}
                    _ = v34
                }
                _ = fi16
                _ = v33
            }
        }
        _ = v14
    }
    v35 := GenS7{F0: 102, F1: 32767}
    _ = p0
    _ = v35
    _ = v9
    return (true && (128 < 42))
}

func main() {
    var v1 int16 = 118
    var v2 string = "hello"
    var v3 uint = 255
    v36 := []int{-9223372036854775808}
    switch -1 {
    case 775:
        if len(v36) > 0 {
            v36[0] = 128
        }
    case 327:
        var v37 uint64 = 1
        _ = v37
    case 32767:
        v3 = v3
    }
    if len(v36) > 0 {
        v36[0] = 9223372036854775807
    }
    v3 = 160
    var v38 uint64 = 18446744073709551615
    if (true || true) {
        v39 := GenS6{F0: 0, F1: 0, F2: 1}
        switch 32768 {
        case 42:
            var v40 uint = 0
            _ = v40
        }
        _ = v39
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        v42 := GenS8{F0: -100, F1: v3, F2: 32767}
        for fi43 := 0; fi43 < 1; fi43++ {
            var v44 uint8 = 18
            _ = fi43
            _ = v44
        }
        fmt.Println(v1)
        _ = fi41
        _ = v42
    }
    if (1 != (v1 | v1)) {
        var v45 int32 = 664
        _ = v45
    } else {
        v46 := GenS6{F0: 1, F1: 141, F2: 566}
        _ = v46
    }
    v3 -= v3
    var v47 uint32 = 0
    v47 = v47
    v48 := GenS4{F0: 42, F1: 2147483647, F2: -122}
    v47 = uint32(v3)
    fmt.Println(v1, v38, v1)
    for fi49 := 0; fi49 < 16; fi49++ {
        if (false && true) {
            if len(v36) > 0 {
                v36[0] = fi49
            }
        } else {
            if true {
                v36 = append(v36, 1)
                v38 = (v38 * v38)
            } else {
                switch (v2 + "foo bar") {
                case "foo bar":
                    var v50 uint32 = (v47 % ((v47) | 1))
                    _ = v50
                case "hello":
                    var v51 float64 = float64(0.1)
                    _ = v51
                case "":
                    var v52 uint8 = 128
                    _ = v52
                }
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
                v38 = v38
            }
            fmt.Println(v2)
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v47
    _ = v48
}

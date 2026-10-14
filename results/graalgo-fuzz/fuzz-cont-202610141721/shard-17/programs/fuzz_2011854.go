package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint; F2 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool; F1 int; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int16; F2 string }
func genFn0() uint8 {
    var v6 int64 = -9223372036854775808
    v7 := []int{}
    if (0.0 <= 3.14159) {
        var v8 int16 = -995
        v6 = int64(v8)
        for fi9 := 0; fi9 < 1; fi9++ {
            if (0 == 255) {
                if len(v7) > 0 {
                    v7[0] = (fi9 - (^(fi9)))
                }
            } else {
                v8 = -1
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := GenS1{F0: fi10}
                    _ = fi10
                    _ = v11
                }
                switch "hello" {
                case "foo bar":
                    v12 := GenS1{F0: fi9}
                    _ = v12
                case "\n\t":
                    v13 := GenS2{F0: 42, F1: 65535, F2: 42}
                    _ = v13
                }
            }
            _ = fi9
        }
        _ = v8
    }
    _ = v6
    _ = v7
    return 222
}

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 1
    var v3 bool = true
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    var v14 int64 = 42
    var v15 string = (((("hello" + ("foo bar" + "")) + "a") + "\n\t") + "a")
    switch v15 {
    case "":
        _ = len(v4)
    case "\n\t":
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 float32 = float32((v1 - v1))
            switch fi16 {
            case 1:
                var v18 float32 = v17
                _ = v18
            case 0:
                v19 := []int{fi16, fi16}
                _ = v19
            case -9223372036854775808:
                v14 = v14
            }
            v3 = v3
            _ = fi16
            _ = v17
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint = 255
        var v22 uint16 = 65535
        _ = fi20
        _ = v21
        _ = v22
    }
    switch 582 {
    case 4294967295:
        v1 -= v1
    case 32767:
        if (3.14159 < 1.5) {
            switch (((v15 + ("a" + ("a" + v15))) + v15) + v15) {
            case "foo bar":
                var v23 int8 = 100
                _ = v23
            default:
                switch -1 {
                case 854:
                    v2 += v2
                default:
                    fmt.Println(v3, v2)
                }
            }
            var v24 string = v15
            _ = v24
        } else {
            fmt.Println(v2)
            v15 = (v15 + v15)
            for fi25 := 0; fi25 < 4; fi25++ {
                if (255 <= 2) {
                    switch ((v15 + "\n\t") + v15) {
                    case "":
                        for fi26 := 0; fi26 < 13; fi26++ {
                            v27 := []int{(fi26 / (((fi26 >> ((4294967295) & 63))) | 1)), (fi26 % ((int(v14)) | 1)), fi26, (fi26 + 128)}
                            v28 := v4[fi26]
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    case "\n\t":
                        for fi29 := 0; fi29 < 13; fi29++ {
                            switch 42 {
                            case -873:
                                var v30 int = fi29
                                _ = v30
                            default:
                                var v31 int = (fi29 - fi25)
                                _ = v31
                            }
                            _ = fi29
                        }
                    default:
                        for fi32 := 0; fi32 < 9; fi32++ {
                            if (false && v3) {
                                var v33 float64 = ((0.0 / v1) / 3.14159)
                                var v34 uint16 = 66
                                _ = v33
                                _ = v34
                            } else {
                                v3 = v3
                            }
                            var v35 int = 0
                            v36 := GenS3{F0: (fi32 ^ (fi32 ^ 127)), F1: v2, F2: genFn0()}
                            var v37 int8 = 42
                            v38 := GenS4{F0: (176 == genFn0()), F1: fi32, F2: v37}
                            _ = fi32
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    }
                    fmt.Println(v1)
                    for fi39 := 0; fi39 < 10; fi39++ {
                        fmt.Println(v1, v2, v1)
                        _ = fi39
                    }
                }
                var v40 uint64 = 1
                var v41 int64 = (v14 + (v14 / ((v14) | 1)))
                _ = fi25
                _ = v40
                _ = v41
            }
        }
    default:
        if (false || true) {
            delete(v5, 65535)
            v42 := []int{42, 32768, 128, 42}
            if ((float64(1.5) != v1) && (0 <= -1)) {
                for fi43 := 0; fi43 < 4; fi43++ {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        var v45 float64 = (v1 - v1)
                        v46 := GenS4{F0: v3, F1: fi43, F2: 100}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    var v47 int16 = 440
                    if len(v42) > 0 {
                        v42[0] = (fi43 << ((92) & 63))
                    }
                    _ = fi43
                    _ = v47
                }
            }
            _ = v42
        }
    }
    v48 := GenS1{F0: 2147483647}
    if (!v3) {
        var v49 int64 = (v14 * v14)
        _ = v49
    } else {
        var v50 bool = (v3 || true)
        switch 128 {
        case 1:
            switch v15 {
            case "":
                _ = len(v5)
            case "foo bar":
                var v51 uint64 = 332
                _ = v51
            case "\n\t":
                switch -1 {
                case 42:
                    switch 4294967295 {
                    case -171:
                        if v3 {
                            v50 = v50
                            v2 -= (^(v2))
                        } else {
                            if true {
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    v50 = (1 == 4294967295)
                                    v2 = (v2 / ((0) | 1))
                                    if (42 != 104) {
                                        v53 := GenS5{F0: 188}
                                        _ = v53
                                    } else {
                                        var v54 uint8 = 1
                                        _ = v54
                                    }
                                    _ = fi52
                                }
                                var v55 float32 = (-2.25 + (0.1 / 1.0))
                                if true {
                                    var v56 float32 = v55
                                    var v57 int64 = (v14 / (((v14 & (v14 * 0))) | 1))
                                    fmt.Println(v56)
                                    _ = v56
                                    _ = v57
                                }
                                _ = v55
                            } else {
                                for fi58 := 0; fi58 < 2; fi58++ {
                                    v59 := GenS6{F0: 1}
                                    if v3 {
                                        v60 := GenS3{F0: 1, F1: v2, F2: 0}
                                        fmt.Println(v1)
                                        _ = v60
                                    } else {
                                        var v61 uint64 = 1
                                        fmt.Println(v1)
                                        v62 := GenS3{F0: fi58, F1: (v2 % (((v2 >> ((255) & 63))) | 1)), F2: 255}
                                        _ = v61
                                        _ = v62
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                                var v63 float32 = 0.1
                                var v64 int16 = 32767
                                _ = v63
                                _ = v64
                            }
                        }
                    case 1:
                        v65 := GenS5{F0: 344}
                        _ = v65
                    default:
                        for fi66 := 0; fi66 < 1; fi66++ {
                            var v67 uint16 = 393
                            if v50 {
                                if (v1 > (float64(3.14159) - v1)) {
                                    v68 := GenS1{F0: -9223372036854775808}
                                    if (v2 <= v2) {
                                        switch v15 {
                                        case "foo bar":
                                            fmt.Println(v67, v67, v15)
                                        }
                                    }
                                    _ = v68
                                } else {
                                    var v69 int8 = 110
                                    _ = v69
                                }
                                if v50 {
                                    v4[fi66] = 613
                                    switch (fi66 | (fi66 * -9223372036854775808)) {
                                    case 4294967295:
                                        v70 := GenS3{F0: (fi66 % (((fi66 >> ((255) & 63))) | 1)), F1: (v2 * v2), F2: 128}
                                        _ = v70
                                    case 127:
                                        var v71 uint = 129
                                        _ = v71
                                    }
                                }
                            } else {
                                v1 = v1
                                v72 := GenS1{F0: fi66}
                                _ = v72
                            }
                            fmt.Println(v67, v14, v50)
                            _ = fi66
                            _ = v67
                        }
                    }
                }
            default:
                v2 = (v2 | 0)
            }
        }
        if ((true || true) || true) {
            v50 = (v3 && (!v50))
            fmt.Println(v1)
        } else {
            var v73 int = -1
            v1 = v1
            _ = v73
        }
        _ = v50
    }
    var v74 uint16 = 2
    v75 := GenS7{F0: v2, F1: -32768, F2: v15}
    v74 -= (v74 & v74)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v74
    _ = v75
}

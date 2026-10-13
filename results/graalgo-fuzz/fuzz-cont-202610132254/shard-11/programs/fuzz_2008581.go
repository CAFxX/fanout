package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
type GenS5 struct { F0 uint32; F1 int64; F2 int8 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 string; F1 int; F2 uint32 }
func genFn0(p0 int) uint {
    p0 -= (p0 ^ (p0 * (^(p0))))
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]string)
    var v3 string = "foo bar"
    var v4 uint64 = 4294967295
    v4 -= v4
    var v5 float64 = 0.1
    if false {
        if (("" + ("a" + v3)) == (v3 + v3)) {
            if true {
                v1 -= (v1 << ((0) & 63))
                if false {
                    for fi6 := 0; fi6 < 10; fi6++ {
                        v7 := []int{(fi6 * (v1 - (v1 >> ((15) & 63)))), v1, fi6}
                        v3 = v3
                        _ = fi6
                        _ = v7
                    }
                } else {
                    v1 = v1
                    if (false || false) {
                        fmt.Println(v4, v1)
                    }
                    if true {
                        for fi8 := 0; fi8 < 19; fi8++ {
                            var v9 int64 = -490
                            for fi10 := 0; fi10 < 6; fi10++ {
                                switch (v3 + (v3 + v3)) {
                                case "foo bar":
                                    var v11 uint32 = 2147483648
                                    _ = v11
                                case "hello":
                                    var v12 uint = 255
                                    _ = v12
                                default:
                                    v4 += (v4 << ((genFn0(-625)) & 63))
                                }
                                fmt.Println(v9)
                                v9 = (v9 ^ v9)
                                _ = fi10
                            }
                            var v13 uint32 = 686
                            _ = fi8
                            _ = v13
                            _ = v9
                        }
                        var v14 float32 = -2.25
                        _ = v14
                    } else {
                        if ((27 == -128) || false) {
                            if (360 <= 829) {
                                var v15 int = 128
                                for fi16 := 0; fi16 < 17; fi16++ {
                                    v17 := []int{v1}
                                    v18 := GenS1{F0: false}
                                    var v19 int = (v15 >> ((2) & 63))
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                _ = v15
                            } else {
                                var v20 int64 = 42
                                switch v1 {
                                case 127:
                                    if (!((!true) && false)) {
                                        if false {
                                            v20 = v20
                                        }
                                        switch 127 {
                                        case 4294967295:
                                            if (70 <= 0) {
                                                v1 += -307
                                                var v21 bool = true
                                                var v22 int32 = 0
                                                _ = v21
                                                _ = v22
                                            }
                                        case -2147483648:
                                            var v23 uint16 = 0
                                            _ = v23
                                        case 32768:
                                            var v24 uint8 = 65
                                            _ = v24
                                        }
                                        if true {
                                            fmt.Println(v20, v5)
                                            var v25 bool = (false || true)
                                            v26 := v2[(v1 / ((0) | 1))]
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                case 0:
                                    v27 := GenS1{F0: true}
                                    _ = v27
                                case -1:
                                    var v28 uint16 = 65535
                                    _ = v28
                                default:
                                    v29 := GenS1{F0: (true && false)}
                                    _ = v29
                                }
                                _ = v20
                            }
                            var v30 uint32 = 487
                            var v31 uint32 = v30
                            _ = v30
                            _ = v31
                        } else {
                            switch (v3 + v3) {
                            case "\n\t":
                                v4 -= v4
                            case "hello":
                                fmt.Println(v5, v3, v1)
                            }
                            var v32 uint8 = 214
                            switch (v1 + 32768) {
                            case 1:
                                v33 := GenS1{F0: (!(v3 >= ((v3 + "foo bar") + "")))}
                                _ = v33
                            case -646:
                                _ = len(v2)
                            }
                            _ = v32
                        }
                        var v34 bool = ((v4 % ((v4) | 1)) == v4)
                        _ = v34
                    }
                }
            } else {
                for fi35 := 0; fi35 < 1; fi35++ {
                    delete(v2, fi35)
                    for fi36 := 0; fi36 < 1; fi36++ {
                        var v37 uint32 = 261
                        v38 := GenS1{F0: false}
                        v39 := GenS2{F0: (-(fi35))}
                        delete(v2, (fi35 << ((2) & 63)))
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = fi35
                }
                v1 = (v1 + (v1 % ((-9223372036854775808) | 1)))
                if true {
                    if (898 != 2) {
                        for fi40 := 0; fi40 < 3; fi40++ {
                            var v41 uint16 = 2
                            if (v1 <= fi40) {
                                v42 := GenS3{F0: 65535, F1: (fi40 * (fi40 + (v1 + v1))), F2: v4}
                                v43 := GenS3{F0: 818, F1: 1, F2: v4}
                                if ((v41 | (v41 % ((1) | 1))) != (v41 / (((v41 - 32767)) | 1))) {
                                    v44 := GenS4{F0: v41, F1: 0, F2: 0}
                                    v45 := GenS3{F0: 739, F1: (-(fi40)), F2: 659}
                                    if true {
                                        v5 -= 1e10
                                        if false {
                                            var v46 int = (fi40 ^ (fi40 - -114))
                                            _ = v46
                                        }
                                        v47 := GenS3{F0: 2147483647, F1: v1, F2: v4}
                                        _ = v47
                                    } else {
                                        if false {
                                            for fi48 := 0; fi48 < 5; fi48++ {
                                                v5 = (float64(((1e10 + float32(v5)) / (((-2.25 + float32(float64(0.1))) / 3.14159) - 0.0))) - v5)
                                                _ = fi48
                                            }
                                        } else {
                                            var v49 uint8 = 1
                                            _ = v49
                                        }
                                        v50 := GenS4{F0: (v41 | (^(v41))), F1: 42, F2: 127}
                                        _ = v50
                                    }
                                    _ = v44
                                    _ = v45
                                }
                                _ = v42
                                _ = v43
                            }
                            _ = fi40
                            _ = v41
                        }
                        for fi51 := 0; fi51 < 12; fi51++ {
                            fmt.Println(v1)
                            v52 := GenS2{F0: fi51}
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        var v53 int32 = -421
                        v54 := GenS2{F0: -863}
                        _ = v53
                        _ = v54
                    }
                    var v55 string = v3
                    v1 = v1
                    _ = v55
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v3 = (v3 + (v3 + (v3 + ("" + "\n\t"))))
                    default:
                        v56 := GenS1{F0: true}
                        _ = v56
                    }
                    var v57 int16 = -1
                    if (v5 >= 0.1) {
                        var v58 float32 = (0.1 - 0.0)
                        fmt.Println(v4, v5)
                        _ = v58
                    } else {
                        v4 = 18446744073709551615
                    }
                    _ = v57
                }
            }
            for fi59 := 0; fi59 < 2; fi59++ {
                v60 := GenS2{F0: fi59}
                _ = fi59
                _ = v60
            }
            v61 := GenS5{F0: 42, F1: 42, F2: 42}
            _ = v61
        }
        v5 = (float64(float32((v5 * v5))) * (1.5 / (v5 - 3.14159)))
    } else {
        fmt.Println(v4)
    }
    v62 := []int{(v1 - 127)}
    fmt.Println(v4, v1, v5)
    var v63 int16 = -1
    v64 := GenS6{F0: -128, F1: 572}
    v5 -= (float64(0.1) - (v5 + (1e10 * 0.0)))
    v65 := GenS7{F0: v3, F1: v1, F2: uint32(v4)}
    var v66 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}

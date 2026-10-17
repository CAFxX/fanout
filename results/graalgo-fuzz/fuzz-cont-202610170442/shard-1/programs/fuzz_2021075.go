package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int16; F1 bool }

func main() {
    v1 := []uint32{2147483648, 651, 0}
    var v2 int64 = 210
    v3 := []int{1}
    var v4 uint8 = 1
    var v5 int16 = -32768
    var v6 uint = 0
    var v7 int = -500
    var v8 uint64 = 623
    switch ("foo bar" + "foo bar") {
    case "a":
        var v9 bool = false
        _ = v9
    case "hello":
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(v4)
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint16 = 255
                v13 := GenS1{F0: 2, F1: 42, F2: v5}
                _ = fi11
                _ = v12
                _ = v13
            }
            if true {
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 float64 = 0.0
                    if false {
                        v2 = int64(v6)
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                v2 += v2
            }
            _ = fi10
        }
    case "":
        var v16 float32 = (-2.25 / 3.14159)
        _ = v16
    default:
        if (512 <= 562) {
            v17 := GenS2{F0: 522}
            var v18 uint8 = (^(v4))
            if true {
                switch "\n\t" {
                case "a":
                    var v19 uint8 = (v4 / (((v18 ^ 42)) | 1))
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 13; fi20++ {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            var v22 int64 = -61
                            if len(v3) > 0 {
                                v3[0] = (fi20 + fi21)
                            }
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                }
            } else {
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 int16 = v5
                    _ = fi23
                    _ = v24
                }
                if true {
                    if false {
                        if true {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 int64 = (v2 / ((-899) | 1))
                                if (!false) {
                                    var v27 int16 = (v5 ^ (v5 >> ((v6) & 63)))
                                    v8 = (v8 << ((v6) & 63))
                                    v7 = 2147483647
                                    _ = v27
                                }
                                if true {
                                    var v28 uint16 = 200
                                    for fi29 := 0; fi29 < 20; fi29++ {
                                        fmt.Println(v2, v4)
                                        var v30 int32 = -2147483648
                                        switch fi29 {
                                        case 9223372036854775807:
                                            v31 := GenS3{F0: (v8 % (((v8 & uint64(v28))) | 1)), F1: (!true), F2: v28}
                                            _ = v31
                                        default:
                                            v32 := GenS3{F0: (^(v8)), F1: true, F2: (v28 % ((v28) | 1))}
                                            _ = v32
                                        }
                                        _ = fi29
                                        _ = v30
                                    }
                                    _ = v28
                                } else {
                                    var v33 bool = (127 > 42)
                                    _ = v33
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                        fmt.Println(v7, v4, v2)
                    } else {
                        var v34 uint64 = (^(v8))
                        _ = v34
                    }
                    _ = len(v3)
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v18 -= v4
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            v36 := GenS4{F0: 42, F1: (v4 >> ((42) & 63))}
                            _ = v36
                        default:
                            var v37 bool = false
                            _ = v37
                        }
                        _ = fi35
                    }
                }
            }
            _ = v17
            _ = v18
        } else {
            for fi38 := 0; fi38 < 14; fi38++ {
                v39 := GenS5{F0: 384}
                _ = fi38
                _ = v39
            }
            for fi40 := 0; fi40 < 19; fi40++ {
                for fi41 := 0; fi41 < 10; fi41++ {
                    if ((v4 >> ((v6) & 63)) >= 255) {
                        v6 += v6
                    } else {
                        v42 := GenS6{F0: v6, F1: 1, F2: (^(v8))}
                        _ = v42
                    }
                    _ = fi41
                }
                _ = fi40
            }
            var v43 int16 = (v5 - v5)
            _ = v43
        }
    }
    var v44 int8 = 100
    var v45 int32 = 931
    v4 += (^(v4))
    for fi46 := 0; fi46 < 9; fi46++ {
        v2 += 9223372036854775807
        _ = fi46
    }
    var v47 string = ((("a" + ("a" + ("foo bar" + "a"))) + "") + "")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v48 int32 = v45
    v3 = append(v3, (v7 ^ (v7 ^ 128)))
    var v49 float64 = (1e10 - float64(1.5))
    _ = len(v3)
    for fi50 := 0; fi50 < 17; fi50++ {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        var v51 uint64 = uint64(v6)
        v6 += v6
        _ = fi50
        _ = v51
    }
    v52 := GenS7{F0: (v5 & v5), F1: ((^(v6)) != v6)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0(p0 uint) int8 {
    v6 := []int{-9223372036854775808, 2147483647, -337}
    var v7 float32 = 1e10
    _ = len(v6)
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{2147483647, 894, 4294967295, 2}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    switch 754 {
    case -9223372036854775808:
        var v8 uint64 = 1
        _ = v8
    case -1:
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
    case 2147483647:
        v1 = append(v1, 42)
    default:
        for fi10 := 0; fi10 < 14; fi10++ {
            if true {
                v11 := GenS1{F0: -1, F1: 2147483647, F2: 42}
                _ = len(v3)
                var v12 int32 = 42
                _ = v11
                _ = v12
            }
            _ = fi10
        }
    }
    var v13 string = ("a" + "a")
    v13 = (v13 + v13)
    v13 = v13
    var v14 string = v13
    v14 = "hello"
    v13 = "foo bar"
    if false {
        var v15 bool = true
        _ = len(v1)
        fmt.Println(v14)
        _ = v15
    }
    var v16 int8 = 127
    for fi17 := 0; fi17 < 7; fi17++ {
        v13 = (v14 + "a")
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: 699, F1: 2147483647, F2: 65535}
            if (65535 > 65535) {
                fmt.Println(v16, fi17, v14)
                v16 -= v16
            } else {
                if true {
                    switch v14 {
                    case "hello":
                        v16 -= (v16 >> ((42) & 63))
                    case "":
                        v1 = append(v1, 65535)
                    case "foo bar":
                        v20 := GenS1{F0: 0, F1: 0, F2: 2}
                        var v21 uint8 = 128
                        if (430 != 32767) {
                            var v22 uint16 = 255
                            var v23 float32 = (3.14159 * float32(3.14159))
                            if (false && ((v23 * v23) >= 0.1)) {
                                v13 = (v13 + (v13 + v13))
                                v24 := GenS2{F0: (v21 & v21), F1: 100, F2: 805}
                                _ = v24
                            } else {
                                v16 += v16
                                fmt.Println(fi17, v21)
                                v22 += v22
                            }
                            _ = v22
                            _ = v23
                        } else {
                            v25 := v4[(v14 + "foo bar")]
                            var v26 int16 = 0
                            var v27 uint = 2
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = v20
                        _ = v21
                    }
                } else {
                    var v28 uint = 777
                    _ = v28
                }
                if false {
                    if (!false) {
                        var v29 uint = 609
                        _ = v29
                    }
                    var v30 int16 = 74
                    v1 = append(v1, 42)
                    if ((v30 % ((42) | 1)) == -1) {
                        var v31 int64 = -1
                        if true {
                            v31 = v31
                            v14 = "a"
                            fmt.Println(fi17, v30, fi18)
                            var v32 uint32 = 65535
                            _ = v32
                        } else {
                            v33 := GenS2{F0: 255, F1: (v16 + (v16 << ((598) & 63))), F2: v31}
                            v34 := GenS2{F0: 255, F1: (v16 >> ((779) & 63)), F2: v31}
                            v30 -= v30
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch -449 {
                            case 553:
                                delete(v2, fi18)
                            }
                            _ = v33
                            _ = v34
                        }
                        var v35 uint16 = 65535
                        v16 += genFn0(519)
                        v35 = v35
                        _ = v31
                        _ = v35
                    }
                    if false {
                        v36 := GenS3{F0: (v16 >> ((2) & 63))}
                        switch fi17 {
                        case 2147483647:
                            v37 := GenS3{F0: v16}
                            _ = v37
                        default:
                            v14 = ((("hello" + "\n\t") + ("hello" + v14)) + v14)
                        }
                        switch (fi18 - fi17) {
                        case 128:
                            var v38 uint64 = 42
                            _ = v38
                        case 361:
                            var v39 uint16 = 65535
                            _ = v39
                        case 765:
                            if false {
                                var v40 bool = (2147483647 <= 2147483647)
                                var v41 float64 = float64(-2.25)
                                v30 -= v30
                                if v40 {
                                    v40 = v40
                                    v42 := GenS3{F0: v16}
                                    _ = v42
                                } else {
                                    v5[2] = (fi18 ^ fi18)
                                }
                                v41 = v41
                                _ = v40
                                _ = v41
                            }
                        }
                        _ = v36
                    } else {
                        var v43 uint8 = 217
                        v44 := GenS4{F0: (v30 | (v30 & (v30 * (v30 << ((1) & 63))))), F1: (fi18 << ((0) & 63))}
                        var v45 int16 = (v30 ^ -1)
                        v46 := GenS5{F0: 1, F1: int8(v45)}
                        _ = v43
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    _ = v30
                }
                fmt.Println(fi17, fi17)
            }
            if len(v1) > 0 {
                v1[0] = 42
            }
            _ = fi18
            _ = v19
        }
        var v47 uint16 = 42
        _ = fi17
        _ = v47
    }
    v14 = "hello"
    v48 := v5[2147483648]
    v49 := GenS2{F0: 127, F1: (v16 - v16), F2: -119}
    v50 := GenS2{F0: 2, F1: genFn0(638), F2: -183}
    v16 -= 1
    v14 = v14
    v13 = (v13 + (v13 + v14))
    v16 = (^(v16))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}

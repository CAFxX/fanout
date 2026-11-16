package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 string; F1 int64 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 127
    var v4 uint8 = uint8(v1)
    var v5 bool = false
    var v6 uint16 = 510
    var v7 int16 = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: int32(v7), F1: 2147483647}
        v10 := GenS2{F0: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v5)
    if v5 {
        if ((-1 < (^(v7))) && false) {
            if (18446744073709551615 <= 4294967295) {
                v5 = v5
                var v11 int = 9223372036854775807
                _ = v11
            } else {
                if v5 {
                    var v12 int64 = int64(v7)
                    v1 = v1
                    _ = v12
                } else {
                    var v13 bool = v5
                    v14 := []int{742}
                    var v15 string = "\n\t"
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v16 := GenS2{F0: 2147483648}
                _ = v16
            }
        }
        v17 := GenS3{F0: ((4294967295 >= 258) || (-809 <= 188)), F1: (v7 / ((v7) | 1)), F2: v6}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 6; fi18++ {
            if v5 {
                if ((v6 + v6) < v6) {
                    v7 = v7
                }
                switch ("a" + "") {
                case "hello":
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v1 = (^(v1))
                        switch fi19 {
                        case -2147483648:
                            var v20 uint64 = v1
                            _ = v20
                        case -789:
                            v21 := []int{(fi19 & (fi19 << ((42) & 63)))}
                            _ = v21
                        case 346:
                            fmt.Println(v1)
                        default:
                            v7 -= v7
                        }
                        var v22 uint = 1
                        _ = fi19
                        _ = v22
                    }
                default:
                    if (v5 && (v7 != v7)) {
                        var v23 int8 = -1
                        if v5 {
                            v6 = v6
                            var v24 int32 = -996
                            switch -398 {
                            case 4294967295:
                                var v25 int8 = -63
                                _ = v25
                            }
                            _ = v24
                        }
                        _ = v23
                    } else {
                        if false {
                            var v26 uint32 = uint32(v1)
                            var v27 string = ""
                            fmt.Println(fi18)
                            _ = v26
                            _ = v27
                        }
                        var v28 int32 = 241
                        _ = v28
                    }
                }
            } else {
                v1 = 42
                if false {
                    var v29 string = "\n\t"
                    fmt.Println(v29, v6, v7)
                    _ = v29
                } else {
                    var v30 float32 = (3.14159 / 3.14159)
                    var v31 uint32 = 42
                    v32 := v2[v31]
                    _ = v30
                    _ = v31
                    _ = v32
                }
                switch "foo bar" {
                case "":
                    v7 -= 0
                default:
                    switch ("" + "hello") {
                    case "\n\t":
                        if (!(false || (!v5))) {
                            var v33 bool = v5
                            fmt.Println(v33, v1, v5)
                            _ = v33
                        } else {
                            v34 := GenS3{F0: v5, F1: (v7 % (((v7 << ((4294967295) & 63))) | 1)), F2: v6}
                            v35 := []int{(fi18 / ((496) | 1)), (fi18 << ((42) & 63)), fi18, 122}
                            v36 := GenS4{F0: 2147483647, F1: v7}
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    default:
                        v37 := []int{fi18}
                        _ = v37
                    }
                }
            }
            if (!false) {
                var v38 int16 = v7
                var v39 uint = 233
                _ = v38
                _ = v39
            }
            _ = fi18
        }
        v40 := GenS2{F0: 1}
        _ = v40
    }
    switch -799 {
    case 128:
        var v41 int8 = 100
        _ = v41
    case 4294967295:
        v42 := []int{1, 32768, 32768}
        _ = v42
    }
    v3 += 106
    for fi43 := 0; fi43 < 14; fi43++ {
        for fi44 := 0; fi44 < 10; fi44++ {
            var v45 uint8 = (v4 ^ (^(v3)))
            var v46 uint8 = 187
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi43
    }
    fmt.Println(v4, v3)
    if true {
        var v47 int64 = 0
        v2[uint32(v6)] = 2
        _ = v47
    } else {
        for fi48 := 0; fi48 < 12; fi48++ {
            var v49 int32 = -1
            if false {
                for fi50 := 0; fi50 < 2; fi50++ {
                    var v51 bool = (v5 && v5)
                    _ = fi50
                    _ = v51
                }
                switch "hello" {
                case "hello":
                    v52 := GenS5{F0: (("foo bar" + "hello") + "\n\t"), F1: 539}
                    _ = v52
                case "\n\t":
                    var v53 int64 = -9223372036854775808
                    _ = v53
                case "":
                    v4 = v3
                }
            } else {
                var v54 int8 = 42
                _ = v54
            }
            if ((v7 >> ((2) & 63)) >= -310) {
                switch (("hello" + "foo bar") + "foo bar") {
                case "\n\t":
                    var v55 float32 = 0.1
                    _ = v55
                case "":
                    switch fi48 {
                    case -9223372036854775808:
                        var v56 float64 = -2.25
                        _ = v56
                    case 32768:
                        fmt.Println(v6, v4)
                    case 390:
                        var v57 bool = false
                        _ = v57
                    }
                }
            } else {
                var v58 uint = 42
                for fi59 := 0; fi59 < 8; fi59++ {
                    switch 4294967295 {
                    case 32768:
                        fmt.Println(v3, v4)
                    case 128:
                        v2[0] = uint32(v3)
                    default:
                        v6 -= 727
                    }
                    var v60 string = "a"
                    v5 = v5
                    var v61 uint8 = (^(v4))
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                _ = v58
            }
            _ = fi48
            _ = v49
        }
        fmt.Println(v4, v7)
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        fmt.Println(fi62)
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

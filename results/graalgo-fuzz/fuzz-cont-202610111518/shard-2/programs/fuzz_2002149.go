package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 uint }
func genFn0(p0 uint32, p1 uint8) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        p0 = 231
        var v5 uint8 = (p1 ^ (p1 / ((1) | 1)))
        _ = fi4
        _ = v5
    }
    v6 := []int{32767, -298}
    v7 := GenS1{F0: 355, F1: 56}
    p0 += 955
    v8 := GenS2{F0: true, F1: (32768 < 32767), F2: 418}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return uint(p1)
}

func main() {
    var v1 uint8 = 0
    var v2 bool = false
    var v3 int8 = -128
    var v9 int8 = (v3 * (v3 * 0))
    v1 = v1
    var v10 uint16 = 2
    v1 += 195
    fmt.Println(v10)
    for fi11 := 0; fi11 < 18; fi11++ {
        if v2 {
            v12 := GenS3{F0: -1, F1: (v10 / (((v10 & v10)) | 1))}
            v13 := GenS2{F0: v2, F1: (32767 != -32768), F2: 42}
            v14 := []int{int(v10), (-(fi11))}
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := GenS3{F0: 1, F1: 1}
        var v16 uint = 0
        _ = fi11
        _ = v15
        _ = v16
    }
    v17 := []int{2147483647, -1}
    var v18 int8 = 41
    if v2 {
        switch 4294967295 {
        case 4294967295:
            for fi19 := 0; fi19 < 14; fi19++ {
                if ((^(v10)) < v10) {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        fmt.Println(fi20, v9)
                        v21 := GenS4{F0: 0, F1: (^(v18))}
                        v9 -= v3
                        _ = fi20
                        _ = v21
                    }
                }
                if v2 {
                    if true {
                        var v22 int16 = 42
                        v9 = v18
                        _ = v22
                    }
                    var v23 uint = 4294967295
                    _ = v23
                } else {
                    fmt.Println(v3, v9, v10)
                    for fi24 := 0; fi24 < 18; fi24++ {
                        if false {
                            switch ("a" + "a") {
                            case "":
                                v25 := GenS4{F0: 2, F1: (v9 % (((v3 * -123)) | 1))}
                                _ = v25
                            }
                            var v26 uint16 = 356
                            if false {
                                v1 = 0
                                switch ("a" + "a") {
                                case "":
                                    v3 -= 127
                                default:
                                    var v27 uint16 = (v26 ^ (v26 - 279))
                                    _ = v27
                                }
                                var v28 int64 = 9223372036854775807
                                v29 := GenS5{F0: 0, F1: ("hello" + ((("" + "") + "foo bar") + ("\n\t" + ("hello" + "hello")))), F2: 81}
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int8 = v9
                                if v2 {
                                    v31 := GenS3{F0: 42, F1: (v10 * (^(v26)))}
                                    _ = v31
                                }
                                var v32 int = fi24
                                _ = v30
                                _ = v32
                            }
                            _ = v26
                        }
                        if len(v17) > 0 {
                            v17[0] = fi19
                        }
                        _ = fi24
                    }
                    var v33 string = ("" + "")
                    _ = v33
                }
                v9 += v18
                _ = fi19
            }
        default:
            fmt.Println(v18, v18)
        }
        if len(v17) > 0 {
            v17[0] = 128
        }
    } else {
        v9 = (v18 >> ((42) & 63))
        switch ("foo bar" + ("a" + "")) {
        case "":
            if (-32768 <= 1000) {
                var v34 int64 = -768
                _ = v34
            }
        case "\n\t":
            fmt.Println(v2, v1, v18)
        }
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}

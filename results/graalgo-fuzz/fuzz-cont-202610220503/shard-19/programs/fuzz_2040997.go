package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int; F2 uint16 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint64 = 42
    var v2 string = "a"
    var v3 int32 = -2147483648
    var v4 int16 = 32767
    var v5 uint = 885
    var v6 uint8 = 2
    v7 := []int{2147483647}
    var v8 int8 = -112
    _ = len(v7)
    v7 = append(v7, 746)
    var v9 uint64 = (v1 ^ (v1 * v1))
    if false {
        if false {
            var v10 int32 = (v3 % ((int32(v5)) | 1))
            if len(v7) > 0 {
                v7[0] = 0
            }
            if true {
                v7 = append(v7, -653)
            }
            _ = v10
        } else {
            var v11 string = "hello"
            var v12 int16 = 769
            v13 := GenS1{F0: v6, F1: (v12 > 42)}
            _ = v11
            _ = v12
            _ = v13
        }
        if (!((v5 >> (((^(v5))) & 63)) > v5)) {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS2{F0: (fi14 + fi14), F1: -775}
                v8 = int8(v3)
                v16 := GenS3{F0: 2147483648, F1: 255}
                _ = fi14
                _ = v15
                _ = v16
            }
            if false {
                switch 1 {
                case -224:
                    if false {
                        switch 127 {
                        case -2147483648:
                            var v17 float32 = 0.0
                            _ = v17
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                default:
                    var v18 int32 = 1
                    _ = v18
                }
                fmt.Println(v8)
                if len(v7) > 0 {
                    v7[0] = int(v6)
                }
            } else {
                for fi19 := 0; fi19 < 13; fi19++ {
                    if true {
                        for fi20 := 0; fi20 < 12; fi20++ {
                            v21 := GenS3{F0: 239, F1: v6}
                            v22 := GenS4{F0: 871, F1: 128, F2: 32767}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        fmt.Println(v2, v2)
                    }
                    v23 := GenS1{F0: 223, F1: ((^(fi19)) != (fi19 - fi19))}
                    _ = fi19
                    _ = v23
                }
                if true {
                    v7 = append(v7, 764)
                } else {
                    if false {
                        v4 = v4
                    }
                    for fi24 := 0; fi24 < 6; fi24++ {
                        v25 := GenS5{F0: (fi24 + fi24), F1: (false && (false && true))}
                        v26 := GenS5{F0: 58, F1: (((v5 ^ 4294967295) < (v5 - 255)) && false)}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
            }
            var v27 int16 = -558
            _ = v27
        } else {
            var v28 float32 = (-2.25 + -2.25)
            v29 := GenS2{F0: 67, F1: v3}
            _ = v28
            _ = v29
        }
    }
    var v30 int16 = 0
    for fi31 := 0; fi31 < 3; fi31++ {
        if len(v7) > 0 {
            v7[0] = (fi31 ^ 32767)
        }
        fmt.Println(v9)
        _ = len(v7)
        _ = fi31
    }
    if ((42 != 0) && (v9 > v9)) {
        var v32 int = -131
        _ = v32
    } else {
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 int64 = -187
            switch fi33 {
            case 128:
                for fi35 := 0; fi35 < 16; fi35++ {
                    v4 -= v30
                    _ = fi35
                }
            case -201:
                v36 := GenS6{F0: (^(v5)), F1: 255}
                _ = v36
            case 2147483647:
                v5 = v5
            }
            _ = fi33
            _ = v34
        }
    }
    fmt.Println(v5, v8, v5)
    _ = len(v7)
    v37 := GenS4{F0: 42, F1: 2147483647, F2: uint16(v5)}
    var v38 string = ""
    v7 = append(v7, -53)
    for fi39 := 0; fi39 < 18; fi39++ {
        v4 += (v4 ^ v30)
        v9 = v1
        _ = fi39
    }
    var v40 uint8 = v6
    v41 := GenS7{F0: (v40 << ((v5) & 63)), F1: -51}
    v8 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

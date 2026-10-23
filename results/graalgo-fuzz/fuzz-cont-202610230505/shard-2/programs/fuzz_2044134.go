package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0(p0 int32, p1 bool) int16 {
    p0 += p0
    v5 := GenS1{F0: p1, F1: (796 > 65535), F2: (42 <= -2147483648)}
    var v6 uint8 = 202
    v7 := GenS1{F0: (337 > 1), F1: (p1 || (!(0 < -1))), F2: (!(p1 && (1.5 < 1.5)))}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -354
}
func genFn1(p0 int32, p1 string) bool {
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 float32 = (1e10 - -2.25)
                p1 = p1
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: true, F1: true, F2: false}
            _ = v11
        }
        _ = fi8
    }
    var v12 int8 = -92
    v13 := []int{42, 127, 23}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    var v3 uint32 = 1
    v4 := []int{-9223372036854775808, -9223372036854775808, -1}
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            v1 = (v1 % ((v1) | 1))
            var v16 int32 = 638
            v17 := GenS2{F0: -1}
            _ = fi15
            _ = v16
            _ = v17
        }
        v1 += v1
        if false {
            v1 += 42
            v4 = append(v4, fi14)
        } else {
            var v18 uint32 = (^(v3))
            _ = v18
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint32 = (v3 ^ v3)
        switch "hello" {
        case "\n\t":
            for fi21 := 0; fi21 < 18; fi21++ {
                v22 := GenS1{F0: (false && true), F1: true, F2: false}
                _ = fi21
                _ = v22
            }
        case "":
            for fi23 := 0; fi23 < 5; fi23++ {
                _ = len(v4)
                fmt.Println(v3)
                _ = fi23
            }
        default:
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS1{F0: true, F1: (!(false || false)), F2: (!(-274 > -1))}
                fmt.Println(v3, fi19)
                switch 9223372036854775807 {
                case 226:
                    fmt.Println(v20)
                }
                _ = fi24
                _ = v25
            }
        }
        var v26 uint8 = 1
        _ = fi19
        _ = v20
        _ = v26
    }
    v27 := GenS3{F0: 0}
    if (127 > -629) {
        v3 = (v3 | uint32(v1))
    } else {
        v28 := GenS3{F0: 50}
        v29 := GenS2{F0: 42}
        var v30 string = "\n\t"
        _ = v28
        _ = v29
        _ = v30
    }
    v31 := GenS4{F0: -1, F1: int64(v3), F2: 2}
    fmt.Println(v3, v3)
    var v32 string = "\n\t"
    fmt.Println(v32, v32, v32)
    if ((3.14159 * 0.1) == 3.14159) {
        fmt.Println(v3, v1)
        for fi33 := 0; fi33 < 18; fi33++ {
            if len(v4) > 0 {
                v4[0] = 349
            }
            v4 = append(v4, (fi33 % (((fi33 ^ fi33)) | 1)))
            var v34 int64 = -9223372036854775808
            _ = fi33
            _ = v34
        }
    } else {
        v35 := GenS5{F0: (0 == 0)}
        if true {
            v32 = ""
        } else {
            switch v32 {
            case "a":
                if (!true) {
                    v36 := GenS5{F0: ((127 <= 42) && true)}
                    _ = v36
                }
            }
            if false {
                v3 = v3
            } else {
                var v37 uint64 = 1
                var v38 uint8 = 2
                v39 := GenS2{F0: 127}
                _ = v37
                _ = v38
                _ = v39
            }
            var v40 uint = uint(v3)
            _ = v40
        }
        v41 := GenS6{F0: 42, F1: v3, F2: ("hello" + v32)}
        _ = v35
        _ = v41
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        for fi43 := 0; fi43 < 3; fi43++ {
            v44 := GenS3{F0: 1}
            fmt.Println(fi43)
            _ = fi43
            _ = v44
        }
        v45 := GenS7{F0: 0}
        _ = fi42
        _ = v45
    }
    var v46 int16 = 1
    fmt.Println(v46)
    switch 2147483647 {
    case 9223372036854775807:
        for fi47 := 0; fi47 < 11; fi47++ {
            switch "" {
            case "hello":
                v3 -= (^(v3))
            case "a":
                var v48 float64 = 1e10
                _ = v48
            default:
                v3 = v3
            }
            _ = fi47
        }
    default:
        v49 := GenS5{F0: (-2.25 >= (1e10 * 0.1))}
        _ = v49
    }
    v50 := GenS1{F0: (true && true), F1: (((!(!(51 > 135))) && (!true)) || (-2147483648 < -176)), F2: (true && true)}
    v46 -= -32768
    v51 := GenS8{F0: 4294967295, F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v46
    _ = v50
    _ = v51
}

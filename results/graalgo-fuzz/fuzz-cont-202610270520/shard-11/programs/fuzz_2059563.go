package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0() uint64 {
    var v4 string = ("hello" + "foo bar")
    v5 := []int{2147483647, -9223372036854775808, 9223372036854775807}
    if (false || false) {
        fmt.Println(v4)
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6)
            if (!(42 <= 0)) {
                var v7 bool = (0 == 42)
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 float32 = 0.0
                    var v10 int8 = 1
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 19; fi11++ {
                    fmt.Println(v4, v4)
                    _ = fi11
                }
                _ = v7
            }
            var v12 uint32 = 147
            _ = fi6
            _ = v12
        }
    }
    switch -1 {
    case 2147483647:
        v4 = ("hello" + "\n\t")
    case 128:
        v4 = (v4 + ("" + v4))
    }
    v13 := GenS1{F0: 0, F1: 128, F2: 1}
    _ = v13
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint64) string {
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 int32 = int32(fi15)
            _ = fi15
            _ = v16
        }
        v17 := []int{(fi14 - (^(fi14))), fi14}
        _ = fi14
        _ = v17
    }
    switch -1 {
    case -962:
        p0 = 1
    }
    _ = p0
    return ("a" + "a")
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = uint16(v1)
    var v3 int8 = 116
    var v18 int32 = -879
    switch 32768 {
    case 32767:
        v1 = v1
    case 174:
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := []int{(fi19 ^ fi19), fi19}
            _ = fi19
            _ = v20
        }
    case 4294967295:
        v21 := []int{32767}
        _ = v21
    default:
        v18 += v18
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS1{F0: 128, F1: 127, F2: uint(v2)}
        _ = fi22
        _ = v23
    }
    var v24 float32 = ((-2.25 - 3.14159) / 1e10)
    v25 := []int{}
    switch 9223372036854775807 {
    case 218:
        v18 += -1
    case 4294967295:
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS1{F0: 2, F1: 127, F2: 0}
            v28 := GenS1{F0: 128, F1: 128, F2: v1}
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        if false {
            for fi29 := 0; fi29 < 14; fi29++ {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                var v30 bool = true
                v3 += v3
                _ = fi29
                _ = v30
            }
            if len(v25) > 0 {
                v25[0] = 32768
            }
            var v31 uint16 = 32767
            _ = v31
        } else {
            fmt.Println(v3)
            var v32 bool = true
            _ = v32
        }
    }
    if true {
        v33 := GenS2{F0: 635, F1: 180}
        v18 = v18
        var v34 int = 2147483647
        _ = v33
        _ = v34
    } else {
        for fi35 := 0; fi35 < 7; fi35++ {
            v25 = append(v25, int(v1))
            _ = fi35
        }
    }
    switch 32767 {
    case 32768:
        var v36 uint8 = 42
        _ = v36
    }
    if ((genFn0() == 1) || (4294967295 != (^(v1)))) {
        if ((v2 >= (v2 % ((v2) | 1))) || ((1 != 127) && (1 > 42))) {
            if false {
                v37 := GenS3{F0: 923, F1: -1, F2: 241}
                if len(v25) > 0 {
                    v25[0] = 0
                }
                _ = v37
            }
        } else {
            v38 := GenS4{F0: "hello"}
            v39 := GenS1{F0: 1, F1: 255, F2: v1}
            var v40 uint16 = (v2 ^ v2)
            _ = v38
            _ = v39
            _ = v40
        }
        switch 42 {
        case 127:
            if (true || true) {
                var v41 uint64 = 0
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = v41
            }
        case 32767:
            for fi42 := 0; fi42 < 17; fi42++ {
                v18 = v18
                v3 = (v3 / ((v3) | 1))
                _ = fi42
            }
        default:
            v2 = (v2 % ((128) | 1))
        }
    } else {
        v43 := GenS5{F0: 2, F1: 128}
        v44 := GenS6{F0: v2}
        _ = v43
        _ = v44
    }
    var v45 string = ""
    fmt.Println(v24)
    v24 -= 1.5
    var v46 int8 = int8(v1)
    v47 := GenS4{F0: ("foo bar" + v45)}
    fmt.Println(v24, v1, v46)
    v48 := GenS5{F0: genFn0(), F1: -9223372036854775808}
    v49 := GenS7{F0: v1, F1: 2, F2: 961}
    var v50 int32 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}

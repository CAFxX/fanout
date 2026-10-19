package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int; F1 int64; F2 uint16 }
func genFn0() int64 {
    var v7 string = ("hello" + ((("hello" + "a") + "a") + "a"))
    fmt.Println(v7)
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 0
    v3 := []int{127, 2147483647}
    v4 := make(map[int]string)
    var v5 int32 = 1
    var v6 uint8 = 128
    v5 -= (-(v5))
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 42, F1: -66}
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int = fi10
            var v12 uint8 = 127
            if len(v3) > 0 {
                v3[0] = fi8
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: 0, F1: -1}
            v15 := GenS2{F0: 100}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi8
        _ = v9
    }
    var v16 uint16 = 42
    for fi17 := 0; fi17 < 7; fi17++ {
        if (-666 >= (v5 % (((v5 ^ (v5 & v5))) | 1))) {
            v18 := GenS3{F0: v1, F1: ("foo bar" + "")}
            var v19 uint16 = (v16 / ((v16) | 1))
            var v20 uint64 = 2
            _ = v18
            _ = v19
            _ = v20
        } else {
            v21 := GenS4{F0: 42, F1: 65535}
            for fi22 := 0; fi22 < 15; fi22++ {
                v2 = (v2 / ((int8(v16)) | 1))
                _ = fi22
            }
            _ = len(v4)
            _ = v21
        }
        switch fi17 {
        case -9223372036854775808:
            for fi23 := 0; fi23 < 10; fi23++ {
                v24 := GenS4{F0: 255, F1: uint(v1)}
                v25 := GenS3{F0: (v1 | (^(v1))), F1: "foo bar"}
                _ = fi23
                _ = v24
                _ = v25
            }
        case -2147483648:
            _ = len(v4)
        case -844:
            v4[fi17] = ("" + "")
        default:
            switch 128 {
            case 4294967295:
                fmt.Println(v6, v1)
            case 9223372036854775807:
                fmt.Println(v5, v1)
            default:
                for fi26 := 0; fi26 < 20; fi26++ {
                    switch fi26 {
                    case 127:
                        if len(v3) > 0 {
                            v3[0] = int(v2)
                        }
                    default:
                        var v27 int64 = genFn0()
                        v28 := GenS4{F0: 42, F1: 0}
                        v29 := v4[fi26]
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                }
            }
        }
        _ = fi17
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 int8 = v2
        var v32 uint = 42
        if ((v1 - 18446744073709551615) > (v1 * 126)) {
            for fi33 := 0; fi33 < 20; fi33++ {
                delete(v4, (fi30 + int(v6)))
                _ = fi33
            }
        } else {
            v5 = v5
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    var v34 int16 = 0
    v35 := GenS5{F0: 794, F1: 1, F2: (v16 << ((4294967295) & 63))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -1 {
    case 0:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 402:
        switch 9223372036854775807 {
        case 1:
            v36 := GenS5{F0: 4294967295, F1: genFn0(), F2: v16}
            _ = v36
        }
    case -2147483648:
        v6 -= 127
    default:
        if false {
            delete(v4, 32768)
        } else {
            var v37 int16 = (-(v34))
            _ = v37
        }
    }
    v5 += v5
    fmt.Println(v2, v2)
    v4[4294967295] = "hello"
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int64 = -9223372036854775808
        _ = fi38
        _ = v39
    }
    v34 = int16(v5)
    v2 += (v2 | v2)
    var v40 uint8 = (v6 | (v6 << ((278) & 63)))
    v6 -= (v6 ^ v40)
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(v34, v1)
        for fi42 := 0; fi42 < 17; fi42++ {
            var v43 int32 = v5
            v2 -= (v2 * -1)
            v6 = 2
            _ = fi42
            _ = v43
        }
        v16 += (v16 | (v16 >> ((995) & 63)))
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

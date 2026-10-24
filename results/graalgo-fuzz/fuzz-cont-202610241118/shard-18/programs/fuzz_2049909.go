package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
func genFn0(p0 int, p1 uint32) uint16 {
    fmt.Println(p1)
    v6 := GenS1{F0: 551, F1: 127}
    switch (p0 % (((p0 & 0)) | 1)) {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int8 = 67
            v9 := []int{fi7, p0, (-(p0)), fi7}
            p0 -= fi7
            _ = fi7
            _ = v8
            _ = v9
        }
    case 42:
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: 42, F1: -105}
            p1 = p1
            _ = fi10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 float32 = ((1.5 - ((-2.25 * 1.5) + 1e10)) - 0.1)
            _ = fi12
            _ = v13
        }
    }
    p0 += p0
    v14 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = p1
    _ = v14
    _ = v6
    return 2
}
func genFn1(p0 int8) int8 {
    fmt.Println(p0)
    var v15 uint32 = 65535
    _ = p0
    _ = v15
    return (p0 >> ((1) & 63))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 string = "hello"
    var v5 int8 = 0
    fmt.Println(v5, v5)
    switch ((((v4 + (("hello" + v4) + (v4 + ("hello" + (v4 + v4))))) + v4) + (("a" + (v4 + "foo bar")) + v4)) + v4) {
    case "foo bar":
        v4 = ""
    case "":
        if false {
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS2{F0: -1, F1: 255, F2: 4294967295}
                fmt.Println(v4)
                v18 := GenS2{F0: -9223372036854775808, F1: 42, F2: (-(fi16))}
                _ = fi16
                _ = v17
                _ = v18
            }
            v5 = v5
            var v19 uint64 = 0
            _ = v19
        } else {
            v20 := GenS3{F0: 233, F1: 244, F2: 127}
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 uint16 = 32767
                var v23 uint16 = (v22 ^ (^(v22)))
                for fi24 := 0; fi24 < 1; fi24++ {
                    if false {
                        var v25 int16 = -1
                        v2[2147483648] = true
                        if true {
                            v22 = 916
                        }
                        _ = len(v2)
                        v26 := GenS4{F0: (v5 | (v5 * 44)), F1: 1, F2: int16(v22)}
                        _ = v25
                        _ = v26
                    }
                    _ = fi24
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            v27 := GenS5{F0: 186}
            _ = v20
            _ = v27
        }
    default:
        var v28 uint = 0
        _ = v28
    }
    var v29 uint16 = 862
    if true {
        v30 := v3[127]
        v5 += v5
        v31 := []int{(v30 ^ (v30 - (v30 >> ((0) & 63)))), v30, (v30 - (v30 << ((613) & 63))), int(v29)}
        _ = v30
        _ = v31
    } else {
        v29 += genFn0(4294967295, 437)
    }
    fmt.Println(v4)
    if (false && false) {
        for fi32 := 0; fi32 < 20; fi32++ {
            v5 += (v5 * genFn1((v5 << ((18446744073709551615) & 63))))
            if true {
                var v33 float64 = (0.0 * float64(0.1))
                _ = v33
            } else {
                fmt.Println(v29, v5)
                v34 := GenS3{F0: uint32(v29), F1: 200, F2: 255}
                _ = v34
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 float32 = 0.0
                if (!false) {
                    fmt.Println(v29, fi35)
                    delete(v1, (v4 + (((v4 + ("foo bar" + "hello")) + v4) + v4)))
                    var v37 uint8 = 128
                    v29 -= (v29 % ((genFn0(fi35, 115)) | 1))
                    var v38 uint8 = (v37 + 36)
                    var v39 int32 = 565
                    v4 = v4
                    var v40 int8 = -1
                    _ = v37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = fi35
                _ = v36
            }
            _ = fi32
        }
        for fi41 := 0; fi41 < 11; fi41++ {
            v4 = v4
            v42 := GenS5{F0: uint(v29)}
            _ = fi41
            _ = v42
        }
        fmt.Println(v5, v4, v4)
    }
    if (1 == 2) {
        var v43 uint8 = 2
        _ = v43
    } else {
        if (((!false) || true) || (128 != 9223372036854775807)) {
            for fi44 := 0; fi44 < 13; fi44++ {
                v5 += v5
                _ = fi44
            }
            fmt.Println(v29, v5)
        }
        fmt.Println(v5)
        switch 32768 {
        case -292:
            v45 := GenS6{F0: -63}
            _ = v45
        case -780:
            for fi46 := 0; fi46 < 10; fi46++ {
                var v47 float32 = 1e10
                _ = fi46
                _ = v47
            }
        case -2147483648:
            if false {
                var v48 int32 = 183
                var v49 int = 0
                _ = v48
                _ = v49
            }
        default:
            var v50 int32 = -2147483648
            _ = v50
        }
    }
    var v51 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}

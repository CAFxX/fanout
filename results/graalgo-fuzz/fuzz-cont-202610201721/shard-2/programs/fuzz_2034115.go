package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 bool }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 uint64) int8 {
    var v7 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: 2147483647}
            if (true || true) {
                v11 := GenS1{F0: 4294967295}
                v12 := GenS1{F0: 42}
                v7 = (v7 << ((829) & 63))
                _ = v11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        v13 := []int{}
        _ = fi8
        _ = v13
    }
    var v14 uint64 = (p0 | p0)
    var v15 float64 = float64(0.1)
    v16 := []int{127, 166, -9223372036854775808}
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    _ = v7
    return -128
}

func main() {
    var v1 int64 = -371
    var v2 uint8 = 127
    var v3 int16 = -32768
    v4 := make(map[int]string)
    var v5 int16 = v3
    var v6 int64 = (v1 << ((116) & 63))
    v17 := GenS1{F0: 0}
    fmt.Println(v3, v1)
    if false {
        v18 := GenS2{F0: (v5 ^ (v3 / (((v3 - -1)) | 1))), F1: "hello", F2: (v1 == (v1 % (((v1 % ((-1) | 1))) | 1)))}
        v2 += (v2 ^ (v2 & (v2 - 2)))
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS2{F0: (v3 ^ (v5 & 912)), F1: ("\n\t" + "hello"), F2: (369 == 2)}
        var v21 uint64 = 1
        if true {
            var v22 float32 = 0.1
            if ((!(!false)) && (true || false)) {
                for fi23 := 0; fi23 < 7; fi23++ {
                    delete(v4, fi23)
                    _ = fi23
                }
            }
            _ = v22
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint = 235
        var v26 int8 = 42
        v27 := GenS1{F0: 2147483647}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    switch ("" + (("foo bar" + "\n\t") + "foo bar")) {
    case "":
        var v28 uint32 = 1
        _ = v28
    case "a":
        var v29 float64 = (-2.25 - 1.5)
        _ = v29
    default:
        for fi30 := 0; fi30 < 1; fi30++ {
            if (1.5 <= 0.0) {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v1 = (v6 % ((-1) | 1))
                    v32 := GenS3{F0: 1, F1: fi31}
                    if ((^(v2)) > v2) {
                        v3 -= (-(v5))
                        var v33 float64 = (0.0 * 3.14159)
                        switch "a" {
                        case "\n\t":
                            var v34 uint32 = 42
                            _ = v34
                        default:
                            v35 := []int{-779, (^(fi30))}
                            _ = v35
                        }
                        v36 := GenS1{F0: 2147483648}
                        _ = v33
                        _ = v36
                    }
                    _ = fi31
                    _ = v32
                }
            }
            if (2147483647 >= 1) {
                if ((2147483647 > -638) && (!true)) {
                    var v37 int16 = v3
                    switch "a" {
                    case "a":
                        v4[(fi30 << ((0) & 63))] = "\n\t"
                    case "\n\t":
                        v38 := GenS1{F0: 4294967295}
                        _ = v38
                    }
                    _ = v37
                } else {
                    v1 = (v6 % ((v1) | 1))
                    for fi39 := 0; fi39 < 11; fi39++ {
                        var v40 int = (fi39 & int(v5))
                        _ = fi39
                        _ = v40
                    }
                    fmt.Println(v2, v5, v3)
                }
                if (42 < 155) {
                    v6 += (-(v6))
                    switch 32767 {
                    case 42:
                        for fi41 := 0; fi41 < 2; fi41++ {
                            var v42 int = fi41
                            v2 = v2
                            _ = fi41
                            _ = v42
                        }
                    case -365:
                        var v43 float32 = 3.14159
                        _ = v43
                    }
                } else {
                    fmt.Println(v3)
                    v2 += v2
                    v44 := GenS1{F0: 2147483647}
                    _ = v44
                }
            }
            _ = fi30
        }
    }
    v1 -= (v1 % (((v6 ^ v1)) | 1))
    var v45 int16 = (v3 >> ((58) & 63))
    v46 := v4[9223372036854775807]
    if ((-666 <= -2147483648) || false) {
        v47 := GenS2{F0: 1, F1: "hello", F2: false}
        v48 := []int{32767, 32767}
        v49 := GenS1{F0: 129}
        _ = v47
        _ = v48
        _ = v49
    } else {
        var v50 float32 = 1e10
        if true {
            var v51 int8 = -128
            v51 += 100
            var v52 int = 2147483647
            _ = v51
            _ = v52
        }
        v6 = (v1 / ((v6) | 1))
        _ = v50
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        v5 -= v5
        _ = fi53
    }
    if (127 != 96) {
        v54 := []int{-625, 127, -965, 32768}
        var v55 uint32 = 2
        _ = v54
        _ = v55
    }
    var v56 uint8 = (v2 / ((v2) | 1))
    switch 32768 {
    case 9223372036854775807:
        v57 := GenS3{F0: -1, F1: -1}
        _ = v57
    case 32767:
        fmt.Println(v3, v5, v46)
    }
    if false {
        for fi58 := 0; fi58 < 12; fi58++ {
            v59 := []int{(fi58 | 9223372036854775807), -9223372036854775808, fi58, (fi58 * fi58)}
            var v60 uint8 = (v56 / (((^(v2))) | 1))
            var v61 uint32 = 4294967295
            _ = fi58
            _ = v59
            _ = v60
            _ = v61
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v56
    _ = v6
}

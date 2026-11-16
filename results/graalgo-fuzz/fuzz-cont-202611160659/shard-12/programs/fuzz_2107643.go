package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint8; F1 string; F2 int }
type GenS6 struct { F0 uint; F1 string }
func genFn0() int8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        switch (fi6 / ((-9223372036854775808) | 1)) {
        case 9223372036854775807:
            v7 := []int{42, (fi6 - 0)}
            _ = v7
        case -880:
            v8 := GenS1{F0: (fi6 % (((fi6 ^ fi6)) | 1)), F1: 486}
            _ = v8
        case -1:
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
        v10 := GenS1{F0: 32767, F1: 1}
        _ = fi6
        _ = v10
    }
    switch 4294967295 {
    case 2147483647:
        var v11 int8 = -106
        _ = v11
    case -9223372036854775808:
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{fi12, fi12, (fi12 + (-(fi12))), (fi12 ^ fi12)}
            v14 := GenS1{F0: -2147483648, F1: int64(fi12)}
            var v15 uint64 = 2
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    case 534:
        v16 := []int{-2147483648, 335, 127}
        _ = v16
    default:
        if (-70 <= 42) {
            var v17 int = 503
            var v18 int32 = 42
            _ = v17
            _ = v18
        } else {
            var v19 int16 = -662
            _ = v19
        }
    }
    if ((1e10 + ((float64((0.0 - ((1e10 + 0.1) + (1e10 / -2.25)))) * -2.25) / 1.0)) == 3.14159) {
        var v20 uint32 = 2147483647
        _ = v20
    } else {
        v21 := []int{1, -9223372036854775808, -9223372036854775808}
        v22 := GenS1{F0: 9223372036854775807, F1: 1}
        var v23 uint8 = 0
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        if (true && true) {
            fmt.Println(fi24)
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 float64 = 3.14159
            v26 -= v26
            v27 := GenS1{F0: fi24, F1: 535}
            var v28 int64 = 355
            if true {
                fmt.Println(fi24, v26, v28)
            }
            v29 := []int{fi25, (fi25 << ((1) & 63)), fi25}
            if (!(true || true)) {
                if false {
                    switch ("a" + "hello") {
                    case "a":
                        v26 += ((float64((1.5 / 0.1)) / -2.25) / 1.5)
                    case "":
                        v30 := GenS1{F0: 1, F1: (v28 ^ (v28 | 507))}
                        _ = v30
                    default:
                        var v31 int8 = -21
                        _ = v31
                    }
                } else {
                    var v32 bool = (true || false)
                    v33 := GenS1{F0: fi24, F1: 1}
                    _ = v32
                    _ = v33
                }
            } else {
                if false {
                    v26 -= (float64(float32((((v26 / (1.5 / (3.14159 / v26))) / v26) / (-2.25 - v26)))) - v26)
                    var v34 int8 = 0
                    v28 = v28
                    _ = v34
                } else {
                    v29 = append(v29, (^(fi24)))
                }
                fmt.Println(fi24)
            }
            v26 -= (v26 / ((float64((1e10 + 1e10)) + v26) * (0.1 * v26)))
            var v35 string = "\n\t"
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
            _ = v29
            _ = v35
        }
        _ = fi24
    }
    if (2147483647 > 1) {
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint32 = 2147483647
            var v38 string = "foo bar"
            fmt.Println(v37, fi36)
            _ = fi36
            _ = v37
            _ = v38
        }
        v39 := []int{32768, 9223372036854775807}
        var v40 float64 = 1e10
        _ = v39
        _ = v40
    } else {
        var v41 uint = 1
        _ = v41
    }
    return 1
}

func main() {
    v1 := []int{2147483647, 1}
    var v2 uint16 = 32767
    v3 := make(map[int]bool)
    var v4 uint32 = 4294967295
    v5 := make(map[int]string)
    var v42 float32 = 1.5
    if (false || true) {
        fmt.Println(v42, v42)
        if false {
            v43 := GenS2{F0: 42, F1: ("hello" + (("" + "foo bar") + (("a" + "") + (("" + "foo bar") + "foo bar"))))}
            _ = v43
        }
        v44 := v3[42]
        _ = v44
    }
    fmt.Println(v2)
    v45 := GenS2{F0: 42, F1: "foo bar"}
    var v46 bool = (0 <= 276)
    var v47 bool = v46
    for fi48 := 0; fi48 < 16; fi48++ {
        if v46 {
            _ = len(v3)
        } else {
            for fi49 := 0; fi49 < 11; fi49++ {
                v50 := GenS1{F0: fi48, F1: 1}
                v2 = 1
                _ = fi49
                _ = v50
            }
        }
        _ = fi48
    }
    if (((float64(v42) - 1.5) * 3.14159) >= 0.0) {
        if (v42 >= ((1.5 + v42) + ((float32((0.1 + (float64((float32(1.5) + ((v42 - v42) - float32(float64(v42))))) + 0.0))) - 3.14159) - 1.5))) {
            fmt.Println(v42)
            var v51 uint8 = 127
            _ = v51
        } else {
            v42 -= 0.0
            fmt.Println(v42)
        }
    } else {
        v52 := GenS3{F0: v46, F1: 0, F2: v2}
        _ = v52
    }
    if false {
        v4 = v4
        v53 := GenS4{F0: uint32(v2), F1: (v2 - 42), F2: -1}
        v54 := GenS3{F0: v47, F1: 0, F2: 396}
        _ = v53
        _ = v54
    }
    v55 := GenS3{F0: true, F1: 570, F2: v2}
    var v56 int = 1
    fmt.Println(v47, v4, v56)
    v4 += 2147483648
    var v57 int16 = -32768
    for fi58 := 0; fi58 < 20; fi58++ {
        if (1 > (v57 - (v57 / ((-32768) | 1)))) {
            if v46 {
                if len(v1) > 0 {
                    v1[0] = v56
                }
                v59 := GenS5{F0: 1, F1: "hello", F2: (v56 * (v56 * v56))}
                fmt.Println(fi58, v47, v57)
                _ = v59
            }
            _ = len(v1)
            var v60 int32 = int32(v4)
            _ = v60
        }
        fmt.Println(v46, fi58)
        _ = fi58
    }
    if v47 {
        if len(v1) > 0 {
            v1[0] = v56
        }
        v61 := GenS6{F0: 2, F1: ((("a" + "foo bar") + "a") + "hello")}
        var v62 uint32 = v4
        _ = v61
        _ = v62
    } else {
        var v63 int64 = -935
        var v64 int = (-(v56))
        if v46 {
            var v65 int8 = -110
            v47 = (434 >= 18446744073709551615)
            if v46 {
                var v66 bool = true
                var v67 uint16 = (v2 % (((v2 - 255)) | 1))
                _ = v66
                _ = v67
            } else {
                v46 = true
                v57 += (v57 >> ((65535) & 63))
            }
            _ = v65
        } else {
            switch "a" {
            case "foo bar":
                fmt.Println(v63, v47, v46)
            case "hello":
                for fi68 := 0; fi68 < 15; fi68++ {
                    v56 = (-(v56))
                    var v69 float64 = 1.5
                    _ = fi68
                    _ = v69
                }
            default:
                var v70 uint32 = v4
                _ = v70
            }
        }
        _ = v63
        _ = v64
    }
    var v71 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v71
}

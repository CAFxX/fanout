package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32) uint16 {
    v4 := GenS1{F0: 42, F1: 42, F2: 319}
    v5 := GenS1{F0: 42, F1: 32767, F2: 1}
    p0 = 2147483648
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1() bool {
    switch ("hello" + "a") {
    case "":
        if (0 != -9223372036854775808) {
            var v6 string = ("foo bar" + ("foo bar" + "hello"))
            var v7 int32 = -2147483648
            v8 := GenS1{F0: 128, F1: 376, F2: -1}
            _ = v6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        if ((0 >= 2147483648) && (!(-975 == 0))) {
            var v10 int32 = 179
            _ = v10
        } else {
            var v11 int64 = 1
            v11 -= (v11 + v11)
            _ = v11
        }
        _ = fi9
    }
    var v12 uint32 = 0
    _ = v12
    return (false && (!false))
}

func main() {
    var v1 float32 = ((1e10 * float32((3.14159 / 1.5))) * (3.14159 / -2.25))
    var v2 float64 = 1e10
    var v3 uint8 = 59
    v3 = (v3 << ((uint(v3)) & 63))
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := []int{(fi13 + fi13), (fi13 + fi13), (fi13 & (fi13 - (^(fi13)))), 0}
        var v15 int = (-(fi13))
        for fi16 := 0; fi16 < 11; fi16++ {
            if len(v14) > 0 {
                v14[0] = 127
            }
            var v17 uint32 = 65535
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v18 := GenS2{F0: 779}
    fmt.Println(v3)
    switch 375 {
    case 42:
        v3 -= v3
    case 2147483647:
        v2 = 1.5
    default:
        v1 -= 1e10
    }
    var v19 uint16 = 1
    v20 := []int{127}
    var v21 uint = 927
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := GenS3{F0: (766 <= (v21 % ((v21) | 1))), F1: -1, F2: 0}
        if (1 > (v19 - genFn0(0))) {
            if len(v20) > 0 {
                v20[0] = fi22
            }
            if (("" >= "a") && true) {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                var v24 uint = (v21 & 65535)
                _ = v24
            } else {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            }
        }
        _ = fi22
        _ = v23
    }
    v3 += v3
    if true {
        v21 = 1
        fmt.Println(v21, v19, v1)
        var v25 uint = (v21 * v21)
        _ = v25
    }
    if true {
        _ = len(v20)
        switch int(v3) {
        case 1:
            if false {
                if (!true) {
                    var v26 int32 = 0
                    _ = v26
                } else {
                    v27 := GenS3{F0: (false || (9223372036854775807 != 42)), F1: -9223372036854775808, F2: 128}
                    _ = v27
                }
            }
        case 32767:
            var v28 uint64 = 18446744073709551615
            _ = v28
        case -9223372036854775808:
            v3 += v3
        default:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        }
        var v29 uint64 = 4294967295
        _ = v29
    } else {
        switch "\n\t" {
        case "\n\t":
            var v30 uint64 = 1
            _ = v30
        default:
            var v31 int8 = 0
            _ = v31
        }
    }
    if len(v20) > 0 {
        v20[0] = 466
    }
    if (!(1 > -9223372036854775808)) {
        for fi32 := 0; fi32 < 10; fi32++ {
            _ = len(v20)
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int32 = -598
                var v35 uint = v21
                var v36 uint16 = v19
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            v19 += (v19 % ((v19) | 1))
            _ = fi32
        }
        if (v21 != 0) {
            v37 := GenS3{F0: ((float64(v1) / ((3.14159 + ((v2 * (1e10 + (float64(v1) / float64(1e10)))) * (v2 * (3.14159 - v2)))) / v2)) <= v2), F1: 1, F2: (^(v3))}
            switch 942 {
            case -337:
                v3 += (v3 >> ((v21) & 63))
            case 2147483647:
                v19 -= (v19 * uint16(v21))
            case 4294967295:
                var v38 uint32 = 202
                _ = v38
            }
            var v39 uint32 = 732
            _ = v37
            _ = v39
        } else {
            switch 32768 {
            case 42:
                fmt.Println(v19)
            case 32768:
                v1 += v1
            default:
                var v40 uint64 = 672
                _ = v40
            }
        }
    } else {
        if true {
            v41 := GenS4{F0: 85}
            _ = v41
        }
        if len(v20) > 0 {
            v20[0] = 127
        }
        v2 -= 0.1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}

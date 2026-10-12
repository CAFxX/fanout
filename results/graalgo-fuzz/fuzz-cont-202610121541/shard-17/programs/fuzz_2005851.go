package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 bool, p1 int16, p2 int64) uint8 {
    p0 = true
    var v5 bool = (!p0)
    for fi6 := 0; fi6 < 5; fi6++ {
        switch "" {
        case "hello":
            v7 := []int{fi6, (fi6 % ((fi6) | 1)), fi6, fi6}
            _ = v7
        case "a":
            var v8 uint16 = 216
            _ = v8
        default:
            v5 = (!(((1 >= 255) || true) || true))
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 128
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int8 = -1
    v3 := []uint32{2, 65535}
    var v4 float64 = 1.5
    var v9 int16 = 32767
    v10 := GenS1{F0: 1, F1: -88, F2: v2}
    _ = len(v3)
    v4 = -2.25
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (((1 >= 97) || (!(754 > 885))) || (-2.25 > ((((1.5 / (1e10 / (float32((float64((1e10 / float32((v4 / (v4 * v4))))) - v4)) * (((1e10 * (((0.1 - 1e10) + 3.14159) * float32((v4 * v4)))) * 1e10) - (0.1 + 1.5))))) - 1e10) + -2.25) + -2.25))) {
        var v11 int64 = 944
        if (float64(1e10) != (-2.25 + v4)) {
            v1 = (v1 << ((164) & 63))
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := GenS1{F0: v11, F1: 0, F2: v2}
                var v14 uint64 = 42
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int = (fi12 ^ (fi12 / ((fi15) | 1)))
                    v17 := GenS1{F0: 1, F1: 452, F2: v2}
                    v2 += (v2 & -47)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                for fi19 := 0; fi19 < 9; fi19++ {
                    v1 += v1
                    _ = fi19
                }
                var v20 uint64 = 347
                _ = fi18
                _ = v20
            }
            if false {
                v1 += v1
                fmt.Println(v2, v4, v2)
                fmt.Println(v1, v4, v2)
            } else {
                fmt.Println(v4, v2)
                v2 = 1
            }
        }
        fmt.Println(v11, v2, v1)
        _ = v11
    } else {
        v1 += (v1 * v1)
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS2{F0: "hello", F1: 1}
            _ = fi21
            _ = v22
        }
        v2 = v2
    }
    switch ("" + "foo bar") {
    case "foo bar":
        fmt.Println(v9, v9)
    case "\n\t":
        v2 = v2
    case "a":
        fmt.Println(v1, v9)
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if true {
            var v23 uint = 255
            v24 := GenS3{F0: 923, F1: ((((("" + "hello") + ((("foo bar" + "hello") + ("" + "\n\t")) + "foo bar")) + "\n\t") + "foo bar") + ("" + "\n\t"))}
            _ = v23
            _ = v24
        }
        v4 -= (((1.5 - v4) - ((v4 - v4) - ((float64((1e10 - 3.14159)) + (-2.25 * (0.1 * 0.1))) - (1.5 + v4)))) + ((v4 + v4) + v4))
    } else {
        switch ("" + (((("\n\t" + "hello") + "a") + ("foo bar" + "\n\t")) + "")) {
        case "":
            var v25 int16 = (v9 / (((v9 + (v9 << ((4294967295) & 63)))) | 1))
            _ = v25
        case "\n\t":
            for fi26 := 0; fi26 < 17; fi26++ {
                v4 += v4
                _ = fi26
            }
        case "a":
            fmt.Println(v1, v2)
        default:
            if true {
                v3 = append(v3, 2147483648)
                var v27 int16 = v9
                for fi28 := 0; fi28 < 15; fi28++ {
                    fmt.Println(v4)
                    _ = fi28
                }
                _ = v27
            } else {
                if false {
                    v4 -= v4
                    v1 = (v1 / ((v1) | 1))
                } else {
                    v9 -= (v9 + 42)
                }
                var v29 uint32 = 4294967295
                var v30 int64 = 996
                _ = v29
                _ = v30
            }
        }
    }
    v31 := GenS2{F0: ("a" + "hello"), F1: 4294967295}
    if len(v3) > 0 {
        v3[0] = 1
    }
    if true {
        if true {
            switch 42 {
            case 9223372036854775807:
                v1 = 2
            default:
                var v32 int8 = v2
                _ = v32
            }
        }
    }
    var v33 int = -9223372036854775808
    var v34 uint64 = v1
    switch v33 {
    case 127:
        var v35 uint = 65535
        _ = v35
    case 0:
        if (4294967295 >= 2147483648) {
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "hello":
                var v36 uint64 = 313
                _ = v36
            }
            if len(v3) > 0 {
                v3[0] = 42
            }
        } else {
            fmt.Println(v2, v2)
            var v37 uint = 4294967295
            if ((v9 ^ (v9 % ((v9) | 1))) < (^(v9))) {
                var v38 uint = 2
                fmt.Println(v4, v2)
                fmt.Println(v2)
                _ = v38
            } else {
                var v39 int8 = (v2 ^ v2)
                _ = v39
            }
            _ = v37
        }
    case -9223372036854775808:
        if (2 != 962) {
            fmt.Println(v34)
            var v40 float64 = 3.14159
            _ = v40
        } else {
            if (((!(false || (42 != 298))) || false) || ((1.5 + -2.25) != 3.14159)) {
                _ = len(v3)
                var v41 uint32 = 138
                if true {
                    _ = len(v3)
                }
                _ = v41
            } else {
                var v42 int64 = 42
                var v43 int64 = v42
                for fi44 := 0; fi44 < 8; fi44++ {
                    var v45 string = ""
                    _ = fi44
                    _ = v45
                }
                _ = v42
                _ = v43
            }
            v2 = (v2 % (((v2 & 42)) | 1))
        }
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v47 uint32 = 133
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}

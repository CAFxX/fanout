package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
type GenS5 struct { F0 uint; F1 int16 }
func genFn0(p0 uint32, p1 uint8, p2 int8) uint {
    var v6 string = ("" + "")
    switch "" {
    case "":
        v7 := []int{32768}
        _ = v7
    case "foo bar":
        p2 += (p2 / ((p2) | 1))
    case "a":
        v8 := []int{-447, -9223372036854775808, 32767, -1}
        _ = v8
    }
    var v9 string = v6
    var v10 uint32 = (p0 ^ 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 int64) uint32 {
    for fi11 := 0; fi11 < 14; fi11++ {
        if (42 == 728) {
            var v12 uint64 = 2
            v12 += (v12 >> ((42) & 63))
            _ = v12
        }
        var v13 int8 = 84
        p0 = p0
        _ = fi11
        _ = v13
    }
    var v14 uint = 42
    _ = p0
    _ = v14
    return 2147483648
}

func main() {
    var v1 int64 = 42
    var v2 float64 = 1e10
    v3 := make(map[string]uint32)
    var v4 int16 = 42
    v5 := make(map[string]bool)
    var v15 float32 = -2.25
    for fi16 := 0; fi16 < 18; fi16++ {
        if true {
            for fi17 := 0; fi17 < 18; fi17++ {
                v3[(("a" + "hello") + "")] = 4294967295
                v18 := GenS1{F0: 4294967295, F1: 42, F2: int32(v1)}
                v2 = (((float64(v15) / (v2 - v2)) * float64(((v15 / v15) / ((v15 - v15) + 1e10)))) + 3.14159)
                _ = fi17
                _ = v18
            }
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 uint = 18446744073709551615
        _ = fi19
        _ = v20
    }
    _ = len(v5)
    v4 += (v4 | (v4 - (v4 << ((209) & 63))))
    var v21 int32 = -2147483648
    var v22 string = (("" + (("foo bar" + "hello") + "")) + "")
    switch -9223372036854775808 {
    case 1:
        if false {
            v23 := GenS1{F0: 852, F1: 2, F2: v21}
            _ = v23
        }
    case -89:
        if (!(-2147483648 < (v21 ^ (-(v21))))) {
            switch v22 {
            case "foo bar":
                delete(v5, ("\n\t" + v22))
            case "hello":
                v1 += 740
            case "":
                fmt.Println(v4, v4, v2)
            }
            fmt.Println(v2)
        } else {
            v4 -= v4
            v24 := []int{-9223372036854775808, 4294967295}
            v1 = v1
            _ = v24
        }
    default:
        if (!(true && false)) {
            v22 = (v22 + v22)
            if ((!(90 >= 32767)) || (false || true)) {
                for fi25 := 0; fi25 < 11; fi25++ {
                    fmt.Println(v2, v15)
                    _ = fi25
                }
                v2 = (v2 * v2)
            } else {
                var v26 uint = 138
                v22 = v22
                _ = v26
            }
            if (((!true) || ((int16(v21) != 1) && (18446744073709551615 != genFn0(2147483647, 0, -128)))) && false) {
                v21 = (v21 * (v21 | -2147483648))
                v21 = (-(v21))
                var v27 uint32 = 1
                _ = v27
            } else {
                var v28 uint32 = 4294967295
                var v29 string = v22
                v30 := GenS1{F0: 1, F1: uint64(v28), F2: (v21 >> ((42) & 63))}
                _ = v28
                _ = v29
                _ = v30
            }
        } else {
            v31 := []int{42, -757, 9223372036854775807}
            var v32 int = 128
            var v33 int8 = 42
            _ = v31
            _ = v32
            _ = v33
        }
    }
    if true {
        var v34 int32 = 2147483647
        for fi35 := 0; fi35 < 18; fi35++ {
            switch (fi35 << ((255) & 63)) {
            case -1:
                v2 += v2
            case -76:
                if true {
                    if true {
                        v36 := []int{}
                        v37 := GenS2{F0: genFn0(2, 115, 42), F1: 255, F2: 127}
                        v1 += (v1 * 9223372036854775807)
                        _ = v36
                        _ = v37
                    }
                    v3[(v22 + "foo bar")] = 42
                    v38 := GenS3{F0: 1}
                    _ = v38
                }
            case 9223372036854775807:
                var v39 float32 = ((0.1 / 0.1) - (((0.1 / (float32(0.0) - float32((v2 + (1e10 + (((v2 * v2) / v2) - v2)))))) / v15) - (3.14159 * (v15 - 3.14159))))
                _ = v39
            default:
                v40 := GenS4{F0: int16(fi35), F1: (v1 - v1), F2: 1000}
                _ = v40
            }
            _ = fi35
        }
        _ = v34
    } else {
        v41 := v3["hello"]
        _ = v41
    }
    fmt.Println(v1, v2)
    for fi42 := 0; fi42 < 1; fi42++ {
        v43 := []int{(fi42 % (((^(fi42))) | 1)), (fi42 * (-(fi42))), fi42, fi42}
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
        _ = fi42
        _ = v43
    }
    switch 756 {
    case 159:
        v44 := GenS5{F0: 18446744073709551615, F1: v4}
        _ = v44
    case -711:
        v45 := []int{127}
        _ = v45
    default:
        v22 = "hello"
    }
    if true {
        v22 = v22
        for fi46 := 0; fi46 < 4; fi46++ {
            v15 = v15
            for fi47 := 0; fi47 < 9; fi47++ {
                if false {
                    if false {
                        switch (fi46 / ((fi46) | 1)) {
                        case -1:
                            v21 -= -444
                        case 42:
                            switch (^(fi47)) {
                            case -457:
                                var v48 string = "foo bar"
                                v3[(v48 + v22)] = 289
                                _ = v48
                            default:
                                fmt.Println(v22, v22, v4)
                            }
                        case -9223372036854775808:
                            switch (v22 + v22) {
                            case "foo bar":
                                var v49 uint64 = 4294967295
                                _ = v49
                            default:
                                v50 := []int{(fi47 % ((fi47) | 1))}
                                _ = v50
                            }
                        default:
                            fmt.Println(v2, v15, v4)
                        }
                    }
                }
                _ = fi47
            }
            _ = fi46
        }
        var v51 int64 = (v1 | (v1 << ((genFn0(65535, 180, -31)) & 63)))
        _ = v51
    }
    v2 -= v2
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}

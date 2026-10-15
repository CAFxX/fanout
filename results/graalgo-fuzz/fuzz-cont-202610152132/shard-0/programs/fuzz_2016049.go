package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint16, p1 int16) int8 {
    var v5 uint32 = 92
    v6 := []int{2147483647}
    var v7 uint = 255
    v8 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 bool) int32 {
    if (-886 == 248) {
        if p1 {
            var v9 int32 = -1
            _ = v9
        } else {
            var v10 int8 = -78
            if (!(127 != 113)) {
                var v11 float64 = 3.14159
                _ = v11
            } else {
                fmt.Println(v10)
                fmt.Println(v10, p0)
                v12 := GenS1{F0: 9223372036854775807}
                _ = v12
            }
            _ = v10
        }
    } else {
        if p0 {
            v13 := GenS1{F0: 1}
            var v14 uint16 = 2
            v14 += 65535
            _ = v13
            _ = v14
        }
        var v15 int64 = 9223372036854775807
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        switch "foo bar" {
        case "foo bar":
            v17 := []int{(^(fi16)), fi16, 2147483647, 0}
            _ = v17
        case "a":
            if (false && (0 < 42)) {
                switch "hello" {
                case "a":
                    var v18 int8 = int8(fi16)
                    _ = v18
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    if (false && (2147483647 <= 2)) {
                        if (-487 < 2147483647) {
                            switch "hello" {
                            case "":
                                var v20 int32 = 2147483647
                                _ = v20
                            case "\n\t":
                                var v21 uint32 = 42
                                _ = v21
                            case "foo bar":
                                var v22 float32 = -2.25
                                switch "" {
                                case "\n\t":
                                    v23 := GenS2{F0: -344, F1: 42}
                                    _ = v23
                                }
                                v22 += ((v22 * (v22 / ((3.14159 - v22) * (3.14159 * 1e10)))) + v22)
                                _ = v22
                            }
                            if (genFn0(447, 32767) > -91) {
                                if true {
                                    var v24 int64 = 0
                                    v25 := []int{}
                                    p1 = (!false)
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 float64 = 1.5
                                    fmt.Println(fi19)
                                    _ = v26
                                }
                                v27 := []int{(fi19 * -628), 1, fi19, 120}
                                p0 = (!(42 == 2147483647))
                                p1 = p1
                                if p0 {
                                    p0 = false
                                }
                                var v28 uint8 = 208
                                p0 = (!(1 <= 228))
                                p0 = (0 >= 262)
                                _ = v27
                                _ = v28
                            }
                            var v29 int16 = -417
                            v30 := GenS1{F0: 1}
                            if (p1 || true) {
                                var v31 float32 = (-2.25 * -2.25)
                                _ = v31
                            }
                            switch "a" {
                            case "\n\t":
                                var v32 int32 = -1
                                _ = v32
                            case "a":
                                v29 += v29
                            case "hello":
                                switch fi19 {
                                case 9223372036854775807:
                                    v33 := GenS3{F0: fi16, F1: 4294967295}
                                    _ = v33
                                case 367:
                                    var v34 string = ("hello" + "\n\t")
                                    _ = v34
                                default:
                                    var v35 uint = 18446744073709551615
                                    _ = v35
                                }
                                v36 := GenS2{F0: 0, F1: 2}
                                _ = v36
                            default:
                                var v37 int = (fi16 | int(v29))
                                _ = v37
                            }
                            _ = v29
                            _ = v30
                        }
                    }
                    var v38 int = (fi16 * fi19)
                    _ = fi19
                    _ = v38
                }
                if (127 == 100) {
                    v39 := GenS2{F0: -2147483648, F1: 0}
                    p1 = (2 >= 42)
                    var v40 int16 = 0
                    _ = v39
                    _ = v40
                } else {
                    var v41 uint64 = 42
                    v41 -= 511
                    _ = v41
                }
            } else {
                p0 = false
                var v42 float32 = 0.1
                var v43 string = "\n\t"
                _ = v42
                _ = v43
            }
        default:
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 int8 = genFn0(1, -1)
                _ = fi44
                _ = v45
            }
        }
        var v46 uint16 = 42
        _ = fi16
        _ = v46
    }
    v47 := []int{127, -199}
    var v48 int32 = 445
    fmt.Println(v48)
    _ = p0
    _ = p1
    _ = v47
    _ = v48
    return (v48 / (((v48 / ((v48) | 1))) | 1))
}

func main() {
    var v1 uint16 = 908
    var v2 bool = ((-128 != 100) && (false || (65535 <= 65535)))
    v3 := make(map[string]int)
    v4 := []string{"foo bar", (("foo bar" + ("foo bar" + ("foo bar" + ""))) + "foo bar")}
    v49 := GenS4{F0: 127, F1: 128}
    v50 := GenS3{F0: 128, F1: 18446744073709551615}
    var v51 int16 = 42
    v4 = append(v4, "foo bar")
    var v52 int64 = int64(v51)
    for fi53 := 0; fi53 < 15; fi53++ {
        var v54 uint = 0
        _ = len(v4)
        _ = fi53
        _ = v54
    }
    fmt.Println(v1, v51)
    if v2 {
        v55 := GenS2{F0: 2147483647, F1: 2}
        v56 := GenS4{F0: 1, F1: 1}
        v52 = -1
        _ = v55
        _ = v56
    } else {
        var v57 string = "hello"
        if ((v51 << ((42) & 63)) != v51) {
            delete(v3, v57)
        }
        _ = v57
    }
    v1 = (v1 & v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v58 := GenS3{F0: int(v1), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v58
}

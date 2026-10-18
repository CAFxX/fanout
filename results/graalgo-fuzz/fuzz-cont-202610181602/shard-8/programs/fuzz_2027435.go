package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int32; F2 int16 }
func genFn0(p0 int64, p1 int16, p2 int64) int16 {
    p1 -= 1
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 2, F1: p1, F2: -811}
    var v8 uint64 = 0
    if ((true || false) || true) {
        fmt.Println(p2, p1, p1)
        switch "" {
        case "a":
            for fi9 := 0; fi9 < 1; fi9++ {
                switch (fi9 >> ((4294967295) & 63)) {
                case 42:
                    if (true && true) {
                        v10 := []int{fi9}
                        v11 := GenS2{F0: 1000}
                        if (!true) {
                            p2 = -642
                            v12 := GenS3{F0: 692}
                            _ = v12
                        } else {
                            p2 -= p0
                            switch "\n\t" {
                            case "hello":
                                var v13 uint32 = 601
                                _ = v13
                            case "":
                                var v14 float32 = 0.0
                                _ = v14
                            }
                        }
                        _ = v10
                        _ = v11
                    } else {
                        if (p1 != 32767) {
                            switch fi9 {
                            case 9223372036854775807:
                                v15 := []int{(fi9 / (((-(fi9))) | 1)), fi9}
                                _ = v15
                            case 0:
                                var v16 uint32 = 2147483647
                                _ = v16
                            case 32767:
                                for fi17 := 0; fi17 < 13; fi17++ {
                                    var v18 uint = 2
                                    if true {
                                        var v19 int32 = -2147483648
                                        var v20 string = "foo bar"
                                        switch (^(fi17)) {
                                        case 963:
                                            v21 := GenS1{F0: v18, F1: (p1 & p1), F2: p1}
                                            _ = v21
                                        case -119:
                                            switch v20 {
                                            case "":
                                                v22 := GenS3{F0: v19}
                                                _ = v22
                                            case "a":
                                                switch -9223372036854775808 {
                                                case 128:
                                                    switch 42 {
                                                    case -9223372036854775808:
                                                        switch (v20 + v20) {
                                                        case "a":
                                                            var v23 float32 = ((0.1 * 3.14159) * 3.14159)
                                                            _ = v23
                                                        case "hello":
                                                            fmt.Println(p1, p0)
                                                        case "":
                                                            v20 = "\n\t"
                                                        }
                                                    case 2147483647:
                                                        v24 := GenS4{F0: 255}
                                                        _ = v24
                                                    default:
                                                        v25 := []int{fi9, 32767, (-(fi17))}
                                                        _ = v25
                                                    }
                                                default:
                                                    fmt.Println(v18)
                                                }
                                            case "hello":
                                                var v26 bool = (!(0 <= 386))
                                                _ = v26
                                            default:
                                                switch (fi17 >> (((^(v18))) & 63)) {
                                                case -577:
                                                    var v27 bool = false
                                                    _ = v27
                                                case 0:
                                                    v28 := []int{(fi17 / ((0) | 1)), (^(fi17)), fi17}
                                                    _ = v28
                                                }
                                            }
                                        case 2147483647:
                                            p0 = 0
                                        default:
                                            switch fi9 {
                                            case 4294967295:
                                                v19 += 247
                                            }
                                        }
                                        _ = v19
                                        _ = v20
                                    } else {
                                        var v29 int = (^(fi17))
                                        _ = v29
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                            default:
                                var v30 string = ("\n\t" + "")
                                _ = v30
                            }
                        }
                    }
                case 32767:
                    v8 += (v8 ^ (v8 * 2))
                case -1:
                    fmt.Println(fi9, p2, v8)
                default:
                    for fi31 := 0; fi31 < 4; fi31++ {
                        var v32 uint8 = 122
                        var v33 float64 = (1e10 * 1e10)
                        var v34 int8 = 1
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                }
                v35 := []int{fi9, fi9, (fi9 >> ((42) & 63))}
                v36 := GenS4{F0: 0}
                _ = fi9
                _ = v35
                _ = v36
            }
        case "":
            for fi37 := 0; fi37 < 2; fi37++ {
                var v38 uint = 65535
                fmt.Println(v8, p2, p1)
                if ((0 > 32767) && true) {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        p2 -= (p0 ^ 297)
                        _ = fi39
                    }
                    var v40 uint32 = 4294967295
                    _ = v40
                } else {
                    var v41 bool = (("a" + (("\n\t" + ("foo bar" + "a")) + "")) >= ((("\n\t" + ("a" + ((("" + ("a" + ((("\n\t" + ("hello" + "hello")) + ("" + "hello")) + ("foo bar" + "")))) + ("foo bar" + ("" + "hello"))) + (("a" + "\n\t") + "a")))) + "\n\t") + "\n\t"))
                    _ = v41
                }
                _ = fi37
                _ = v38
            }
        case "hello":
            var v42 int32 = 2147483647
            _ = v42
        default:
            v43 := []int{32767, -594, 4294967295, int(p1)}
            _ = v43
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 % (((^(p1))) | 1))
}

func main() {
    var v1 int64 = 0
    var v2 int16 = -1
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    var v5 uint = 65535
    var v6 uint16 = 968
    var v44 int = 4294967295
    var v45 float64 = (0.0 / (((float64((1e10 - 1.5)) * (float64(((3.14159 / 1e10) / (3.14159 + 0.1))) - 3.14159)) - v3) / 1.0))
    var v46 uint32 = 423
    if (!(false || true)) {
        var v47 int16 = -32768
        _ = v47
    } else {
        v5 += 119
        v48 := v4["\n\t"]
        fmt.Println(v5)
        _ = v48
    }
    v49 := GenS5{F0: v2, F1: -1, F2: v2}
    v2 = genFn0((v1 - v1), (v2 & (-(v2))), v1)
    v50 := GenS1{F0: v5, F1: genFn0(v1, (v2 & genFn0((v1 | int64(v5)), (v2 >> (((^(v5))) & 63)), v1)), (v1 + 9223372036854775807)), F2: genFn0(int64(v2), 32767, v1)}
    var v51 int16 = int16(v5)
    var v52 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 int; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
func genFn0(p0 int8, p1 uint32) uint32 {
    p0 = 100
    for fi4 := 0; fi4 < 3; fi4++ {
        if false {
            v5 := GenS1{F0: 65535, F1: 677, F2: 0}
            _ = v5
        } else {
            v6 := []int{fi4}
            switch "hello" {
            case "":
                if false {
                    if (!(!(!true))) {
                        var v7 int8 = p0
                        fmt.Println(fi4, p0)
                        _ = v7
                    }
                    v8 := GenS1{F0: 65535, F1: 549, F2: 0}
                    for fi9 := 0; fi9 < 17; fi9++ {
                        _ = len(v6)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        p1 = p1
                        p0 += (-(p0))
                        _ = fi9
                    }
                    _ = v8
                }
            case "hello":
                var v10 int64 = -1
                _ = v10
            }
            var v11 int = fi4
            _ = v11
            _ = v6
        }
        var v12 int8 = (p0 >> ((4294967295) & 63))
        _ = fi4
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint = 679
        _ = fi13
        _ = v14
    }
    var v15 uint = 18446744073709551615
    p0 = p0
    _ = p0
    _ = p1
    _ = v15
    return (p1 / (((p1 >> ((255) & 63))) | 1))
}
func genFn1() bool {
    var v16 float32 = (3.14159 - float32(((-2.25 * 0.0) / 1.0)))
    fmt.Println(v16)
    _ = v16
    return false
}

func main() {
    var v1 float32 = (float32((3.14159 + 1e10)) - (3.14159 * 0.0))
    var v2 uint8 = 0
    v3 := []string{"hello", (("hello" + "hello") + "a"), "\n\t", "a"}
    var v17 int64 = -937
    v3 = append(v3, (("hello" + "a") + "\n\t"))
    v18 := GenS2{F0: 431, F1: (v17 + (v17 & (v17 / ((v17) | 1)))), F2: 2}
    switch "a" {
    case "\n\t":
        var v19 uint64 = 4294967295
        _ = v19
    }
    if ((!(!true)) || (1 >= 1)) {
        if (!true) {
            var v20 int64 = v17
            for fi21 := 0; fi21 < 4; fi21++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v20, v2)
                _ = fi21
            }
            switch 127 {
            case 42:
                for fi22 := 0; fi22 < 7; fi22++ {
                    if true {
                        if (v2 < v2) {
                            var v23 string = "hello"
                            v24 := GenS2{F0: (fi22 >> ((751) & 63)), F1: (v17 + 1), F2: 4294967295}
                            _ = v23
                            _ = v24
                        } else {
                            switch ("hello" + "hello") {
                            case "":
                                v2 -= 3
                            case "\n\t":
                                if true {
                                    v1 -= -2.25
                                    if (float64((v1 + (v1 + v1))) <= -2.25) {
                                        v1 += 0.1
                                        v2 = (v2 % (((^(v2))) | 1))
                                    } else {
                                        if false {
                                            for fi25 := 0; fi25 < 20; fi25++ {
                                                v2 = (v2 >> ((541) & 63))
                                                var v26 uint32 = 2147483648
                                                switch "hello" {
                                                case "foo bar":
                                                    v1 = 3.14159
                                                }
                                                _ = fi25
                                                _ = v26
                                            }
                                        }
                                        if (0 <= 0) {
                                            var v27 float64 = 0.0
                                            _ = v27
                                        }
                                        for fi28 := 0; fi28 < 3; fi28++ {
                                            if (!true) {
                                                v2 += v2
                                            }
                                            _ = fi28
                                        }
                                    }
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v30 := GenS3{F0: 80, F1: "hello", F2: 2}
                                        v31 := GenS1{F0: 277, F1: 2147483647, F2: 42}
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            case "hello":
                                fmt.Println(v20)
                            default:
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + "hello")
                                }
                            }
                        }
                        if false {
                            for fi32 := 0; fi32 < 20; fi32++ {
                                fmt.Println(v17, v2, v20)
                                _ = fi32
                            }
                        }
                    } else {
                        v33 := GenS4{F0: fi22}
                        v2 -= 0
                        if false {
                            fmt.Println(v1, v2)
                            if false {
                                for fi34 := 0; fi34 < 20; fi34++ {
                                    v17 -= (-(v17))
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v20 += (v17 * (v17 | 42))
                                    _ = fi34
                                }
                            }
                            for fi35 := 0; fi35 < 3; fi35++ {
                                v20 += 9223372036854775807
                                v36 := GenS5{F0: 630}
                                fmt.Println(fi35, v1)
                                _ = fi35
                                _ = v36
                            }
                        } else {
                            v3 = append(v3, "")
                        }
                        _ = v33
                    }
                    var v37 string = "\n\t"
                    _ = fi22
                    _ = v37
                }
            case 4294967295:
                var v38 float32 = float32(((1.5 / 1.0) * (0.1 / 1.0)))
                _ = v38
            case 2147483647:
                var v39 float32 = float32(1e10)
                _ = v39
            default:
                fmt.Println(v2)
            }
            _ = v20
        } else {
            if ((0 < 0) || true) {
                if (!false) {
                    v3 = append(v3, "\n\t")
                } else {
                    if len(v3) > 0 {
                        v3[0] = (("a" + "a") + "\n\t")
                    }
                    switch "\n\t" {
                    case "hello":
                        switch 42 {
                        case 4294967295:
                            v2 = (v2 >> ((800) & 63))
                        case 215:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 535:
                            var v40 uint64 = 2
                            _ = v40
                        default:
                            v41 := GenS5{F0: -1}
                            _ = v41
                        }
                    case "\n\t":
                        fmt.Println(v1, v17, v2)
                    case "a":
                        if ((1.5 - float32(-2.25)) > (v1 / 1.0)) {
                            var v42 uint32 = 2147483648
                            v42 += v42
                            var v43 uint16 = 0
                            _ = v42
                            _ = v43
                        } else {
                            if (true || ((1e10 / (((1.5 * (3.14159 / float64(v1))) / (3.14159 - 1.5)) / ((1e10 * 1e10) * ((0.0 - ((0.0 - 0.1) - float64(1.5))) + 1e10)))) >= (((1.5 - (1e10 * (0.1 + ((0.1 * (0.1 + 1.5)) + -2.25)))) * 0.0) + 0.0))) {
                                var v44 float32 = 1e10
                                fmt.Println(v2, v17)
                                _ = v44
                            }
                        }
                    default:
                        v2 -= 255
                    }
                    var v45 bool = ((2 != 4294967295) && (128 != (v2 - (v2 / (((v2 - 128)) | 1)))))
                    _ = v45
                }
            } else {
                for fi46 := 0; fi46 < 1; fi46++ {
                    for fi47 := 0; fi47 < 17; fi47++ {
                        v2 -= (v2 ^ v2)
                        if ("hello" != ("foo bar" + "a")) {
                            fmt.Println(v2)
                            _ = len(v3)
                        }
                        _ = fi47
                    }
                    _ = fi46
                }
                _ = len(v3)
            }
            var v48 int64 = (v17 * v17)
            var v49 int32 = 42
            _ = v48
            _ = v49
        }
        fmt.Println(v17, v2, v17)
        _ = len(v3)
    }
    v17 += (v17 >> ((2) & 63))
    if len(v3) > 0 {
        v3[0] = (("a" + "a") + "\n\t")
    }
    v2 -= (v2 + (v2 % ((69) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}

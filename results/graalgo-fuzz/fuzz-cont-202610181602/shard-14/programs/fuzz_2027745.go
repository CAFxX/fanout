package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint16 }
func genFn0(p0 uint, p1 uint64, p2 uint) string {
    var v6 int64 = -1
    switch ("a" + "foo bar") {
    case "hello":
        var v7 uint64 = (p1 / ((p1) | 1))
        _ = v7
    }
    var v8 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return ("\n\t" + ("" + ("a" + "foo bar")))
}
func genFn1(p0 uint64, p1 uint) int64 {
    v9 := GenS1{F0: 42}
    switch 1 {
    case 42:
        fmt.Println(p1)
    case 2147483647:
        p1 += 18446744073709551615
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            p0 -= p0
            p0 += p0
            _ = fi10
        }
    }
    v11 := GenS1{F0: 2147483648}
    _ = p0
    _ = p1
    _ = v11
    _ = v9
    return -805
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    v4 := []string{("hello" + "")}
    var v5 bool = false
    fmt.Println(v1)
    v12 := GenS2{F0: 0, F1: -1}
    if (1000 == 32767) {
        v13 := GenS3{F0: 1000, F1: ("\n\t" >= "foo bar"), F2: 42}
        _ = v13
    } else {
        var v14 string = "hello"
        switch -187 {
        case 42:
            var v15 int = 407
            _ = v15
        case -579:
            fmt.Println(v5, v14, v1)
        default:
            v1 = (v1 % (((v1 * v1)) | 1))
        }
        _ = v14
    }
    v4 = append(v4, "a")
    var v16 float64 = 0.1
    v17 := v2[(("a" + ("foo bar" + "")) + "")]
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS4{F0: 1}
        v17 = v17
        delete(v2, ("foo bar" + (("a" + ("hello" + ("foo bar" + ("foo bar" + "\n\t")))) + "")))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        if (v17 && (((fi20 ^ (fi20 * 9223372036854775807)) > (fi20 % (((fi20 * (^(fi20)))) | 1))) || v17)) {
            v21 := GenS4{F0: 42}
            var v22 uint32 = v1
            _ = v21
            _ = v22
        } else {
            v23 := GenS5{F0: -439, F1: 135}
            v24 := GenS4{F0: 65535}
            v25 := GenS6{F0: true, F1: (v1 / (((v1 - v1)) | 1))}
            _ = v23
            _ = v24
            _ = v25
        }
        var v26 int8 = 1
        var v27 bool = v17
        _ = fi20
        _ = v26
        _ = v27
    }
    v5 = (-97 == 128)
    var v28 uint64 = 1
    if v5 {
        v5 = ((((1.5 + 0.0) + float32(v16)) + 0.1) <= 3.14159)
        v29 := v2["hello"]
        for fi30 := 0; fi30 < 14; fi30++ {
            if v5 {
                var v31 float64 = (v16 / (((0.0 + 0.0) + 3.14159) + 1.5))
                _ = v31
            }
            var v32 float32 = 1.5
            v5 = ((v1 + v1) < 675)
            _ = fi30
            _ = v32
        }
        _ = v29
    }
    if v5 {
        if v5 {
            v33 := GenS3{F0: -848, F1: true, F2: 0}
            if v17 {
                if v5 {
                    v16 += (v16 / 3.14159)
                } else {
                    fmt.Println(v17)
                }
            } else {
                for fi34 := 0; fi34 < 2; fi34++ {
                    v17 = (42 == 747)
                    for fi35 := 0; fi35 < 11; fi35++ {
                        if v17 {
                            if (1 <= 1) {
                                v28 += v28
                                fmt.Println(fi35, v1, fi35)
                                var v36 uint32 = 61
                                _ = v36
                            } else {
                                var v37 string = "foo bar"
                                _ = v37
                            }
                            var v38 string = ("a" + "\n\t")
                            var v39 int = fi35
                            _ = v38
                            _ = v39
                        } else {
                            v40 := GenS5{F0: 1, F1: int16(v28)}
                            if ((!v17) || false) {
                                v5 = ((((1000 <= 1000) && v5) && (v5 && false)) && v5)
                            }
                            v1 += v1
                            _ = v40
                        }
                        _ = fi35
                    }
                    _ = fi34
                }
                v3[v1] = 4294967295
                v16 = 0.1
            }
            _ = v33
        } else {
            fmt.Println(v1)
            var v41 bool = ((!true) || v5)
            _ = v41
        }
        v28 = v28
        v42 := v3[(v1 << ((255) & 63))]
        _ = v42
    }
    if v5 {
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 5; fi44++ {
                var v45 float32 = -2.25
                if v17 {
                    var v46 uint64 = (v28 >> ((uint(v28)) & 63))
                    var v47 float64 = v16
                    _ = v46
                    _ = v47
                } else {
                    v48 := v3[(v1 / (((v1 % ((v1) | 1))) | 1))]
                    var v49 int = (fi44 / ((int(v48)) | 1))
                    switch "\n\t" {
                    case "":
                        v5 = false
                    case "hello":
                        var v50 uint = 2
                        _ = v50
                    case "a":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    fmt.Println(v1, v49, v48)
                    v48 += 2147483648
                    fmt.Println(v48, v45, v48)
                    var v51 float64 = v16
                    _ = v48
                    _ = v49
                    _ = v51
                }
                var v52 int64 = 180
                switch fi44 {
                case 1:
                    v16 += (v16 - float64(float32(v16)))
                default:
                    v53 := GenS7{F0: -670}
                    _ = v53
                }
                if (684 <= 4294967295) {
                    v28 = v28
                    if ((v28 & (v28 / ((v28) | 1))) < 527) {
                        if v5 {
                            var v54 uint = 18446744073709551615
                            var v55 uint64 = 42
                            _ = v54
                            _ = v55
                        }
                        v56 := GenS7{F0: 1}
                        _ = len(v3)
                        _ = v56
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    switch (fi44 * fi43) {
                    case 0:
                        switch 4294967295 {
                        case 0:
                            delete(v3, 2)
                        case -2147483648:
                            v4 = append(v4, ("a" + ("\n\t" + "a")))
                        case -102:
                            var v57 int32 = 805
                            _ = v57
                        }
                    case 119:
                        switch (-(fi44)) {
                        case 0:
                            v4 = append(v4, "hello")
                            _ = len(v4)
                        case 1:
                            v58 := GenS8{F0: fi43, F1: 2}
                            _ = v58
                        case 59:
                            var v59 int8 = 49
                            _ = v59
                        }
                        v1 -= v1
                        v1 -= v1
                    default:
                        v60 := GenS4{F0: 32767}
                        _ = v60
                    }
                    v45 = v45
                }
                _ = fi44
                _ = v45
                _ = v52
            }
            v61 := GenS6{F0: false, F1: 2}
            var v62 uint16 = 42
            _ = fi43
            _ = v61
            _ = v62
        }
        v3[65535] = v1
    } else {
        for fi63 := 0; fi63 < 4; fi63++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi63
        }
        var v64 int16 = 42
        _ = v64
    }
    var v65 int = 127
    fmt.Println(v28)
    v16 += v16
    var v66 int = v65
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
}

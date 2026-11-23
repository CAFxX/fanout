package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int }
func genFn0(p0 string) uint16 {
    var v4 string = (p0 + p0)
    v5 := GenS1{F0: 2, F1: 400, F2: -385}
    v4 = (v4 + ((v4 + p0) + v4))
    v6 := GenS1{F0: 283, F1: 0, F2: 32767}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 76
}
func genFn1(p0 uint16) int8 {
    switch 2147483647 {
    case 32767:
        p0 = (p0 << ((65535) & 63))
    }
    fmt.Println(p0)
    p0 += p0
    for fi7 := 0; fi7 < 20; fi7++ {
        if ((float32((0.0 / (3.14159 - float64(1.5)))) > -2.25) || (!(647 != 42))) {
            fmt.Println(fi7, fi7)
            if true {
                p0 = 2
            } else {
                p0 = 0
            }
            if (127 <= 128) {
                if true {
                    v8 := []int{4294967295}
                    v9 := GenS1{F0: 4294967295, F1: 0, F2: -1}
                    _ = v8
                    _ = v9
                }
                v10 := []int{(fi7 % (((fi7 / (((^(fi7))) | 1))) | 1)), (fi7 | (fi7 % (((fi7 + 42)) | 1)))}
                _ = v10
            } else {
                v11 := []int{(fi7 * -1)}
                if len(v11) > 0 {
                    v11[0] = 9223372036854775807
                }
                _ = v11
            }
        }
        p0 = p0
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := []int{127}
            v14 := GenS2{F0: 42}
            v13 = append(v13, fi12)
            var v15 uint64 = 18446744073709551615
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    v16 := []int{1, -728}
    _ = p0
    _ = v16
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 917
    var v3 uint = 65535
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int32 = 1
        if false {
            var v19 uint8 = 128
            v1[2147483648] = ""
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    switch -628 {
    case 32768:
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 uint32 = 2147483648
            v21 -= v21
            v22 := GenS2{F0: 207}
            _ = fi20
            _ = v21
            _ = v22
        }
    case 4294967295:
        v23 := GenS3{F0: 127, F1: (v3 % ((v3) | 1)), F2: 751}
        _ = v23
    case 2147483647:
        v1[661] = "foo bar"
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v2 -= uint16(v3)
        v25 := GenS3{F0: 128, F1: v3, F2: 17}
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: genFn0(("" + ("a" + "foo bar"))), F1: 1000}
    fmt.Println(v3, v3)
    v2 = (v2 >> (((v3 - v3)) & 63))
    var v27 uint32 = 2147483648
    v28 := GenS5{F0: v27}
    var v29 int16 = 1
    for fi30 := 0; fi30 < 18; fi30++ {
        if (!(true || true)) {
            fmt.Println(v27, fi30, v29)
            var v31 int8 = -128
            if false {
                if ((1e10 <= (((1.5 - 1e10) / float32(1.5)) * 1e10)) || true) {
                    v32 := GenS4{F0: (v2 & (v2 - (v2 % ((255) | 1)))), F1: (v29 & v29)}
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        if ((v31 ^ (v31 | (^(v31)))) <= v31) {
                            fmt.Println(v3, v3)
                        } else {
                            v34 := GenS2{F0: 822}
                            _ = v34
                        }
                        if (false || (false && false)) {
                            switch fi33 {
                            case 579:
                                var v35 uint16 = 255
                                _ = v35
                            default:
                                var v36 int = -9223372036854775808
                                _ = v36
                            }
                            var v37 int64 = 1
                            var v38 bool = ("foo bar" != "hello")
                            v39 := GenS2{F0: (v37 / (((v37 - (^(v37)))) | 1))}
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = fi33
                    }
                }
            } else {
                fmt.Println(v3, v27, v27)
                switch "foo bar" {
                case "foo bar":
                    var v40 int = fi30
                    _ = v40
                default:
                    var v41 int = fi30
                    _ = v41
                }
                v42 := []int{int(v27)}
                _ = v42
            }
            _ = v31
        }
        switch (fi30 << ((731) & 63)) {
        case 9223372036854775807:
            fmt.Println(fi30, v29, v29)
        }
        v43 := GenS6{F0: 0, F1: 2147483647}
        _ = fi30
        _ = v43
    }
    v44 := GenS3{F0: 131, F1: v3, F2: 4294967295}
    v3 += 151
    var v45 int64 = -1
    for fi46 := 0; fi46 < 18; fi46++ {
        delete(v1, v27)
        var v47 uint8 = 128
        var v48 float64 = (float64((0.1 + float32(float64(0.0)))) * 1e10)
        _ = fi46
        _ = v47
        _ = v48
    }
    v49 := v1[v27]
    v27 -= v27
    if false {
        v2 += 42
        for fi50 := 0; fi50 < 15; fi50++ {
            v51 := []int{fi50, (fi50 - fi50), fi50}
            v49 = (v49 + v49)
            var v52 int16 = v29
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    if false {
        v49 = (v49 + v49)
        for fi53 := 0; fi53 < 12; fi53++ {
            var v54 int8 = 127
            v55 := GenS4{F0: v2, F1: v29}
            fmt.Println(v29)
            _ = fi53
            _ = v54
            _ = v55
        }
    } else {
        v27 = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v44
    _ = v45
    _ = v49
}

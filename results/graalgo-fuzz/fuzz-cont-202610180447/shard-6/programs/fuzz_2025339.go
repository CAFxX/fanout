package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int64 }
func genFn0(p0 int32, p1 uint32, p2 int) int32 {
    var v4 int64 = -1
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int64 = -1
        var v7 bool = true
        v8 := GenS1{F0: (p0 % (((p0 ^ -2147483648)) | 1)), F1: 1, F2: 127}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        var v9 bool = (true && (1.5 >= 3.14159))
        _ = v9
    default:
        if (!false) {
            v4 += 330
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: (p0 & (p0 / (((^(p0))) | 1))), F1: -128, F2: 128}
                v4 -= 0
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: p0, F1: int8(p0), F2: 255}
            _ = v12
        } else {
            v13 := GenS1{F0: p0, F1: 9, F2: 118}
            v14 := GenS2{F0: (p2 * (p2 % ((p2) | 1))), F1: (p1 & (p1 % ((2147483647) | 1)))}
            _ = v13
            _ = v14
        }
    }
    switch p2 {
    case -1:
        v15 := []int{32767, p2, (p2 ^ 4294967295)}
        _ = v15
    case 32768:
        if true {
            if ((true && (-888 >= 1)) && ("foo bar" < ("\n\t" + (("hello" + ("a" + "a")) + "a")))) {
                if (42 < p2) {
                    v16 := GenS3{F0: -68, F1: (p0 | (p0 ^ p0))}
                    _ = v16
                }
                v17 := []int{(p2 + 32767), int(p0), (-(p2))}
                _ = v17
            }
            p0 += (p0 ^ 295)
        } else {
            if false {
                var v18 int32 = p0
                _ = v18
            }
        }
    case 9223372036854775807:
        v19 := []int{(p2 ^ int(p1))}
        _ = v19
    }
    switch p2 {
    case 32767:
        v20 := GenS2{F0: -2147483648, F1: p1}
        _ = v20
    case -9223372036854775808:
        v21 := []int{}
        _ = v21
    default:
        if false {
            var v22 float64 = (3.14159 + (1.5 / float64(0.1)))
            p0 = (^(p0))
            _ = v22
        } else {
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 float64 = ((0.0 + (float64(1.5) + (1.5 + 1e10))) + 3.14159)
                var v25 float64 = v24
                _ = fi23
                _ = v24
                _ = v25
            }
            if true {
                p1 = (p1 / ((843) | 1))
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 + (p0 & (^(p0))))
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = -2.25
    v3 := make(map[uint32]bool)
    v26 := GenS1{F0: (v1 ^ v1), F1: -128, F2: 77}
    var v27 float32 = (v2 * (v2 + -2.25))
    var v28 bool = true
    fmt.Println(v28, v28, v2)
    v1 -= v1
    var v29 string = "foo bar"
    for fi30 := 0; fi30 < 7; fi30++ {
        _ = len(v3)
        _ = len(v3)
        v31 := GenS4{F0: v29, F1: 678, F2: -511}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        switch ((v29 + "a") + (v29 + v29)) {
        case "foo bar":
            v27 = -2.25
        case "hello":
            fmt.Println(v28)
        }
        _ = fi32
    }
    v33 := v3[272]
    var v34 uint32 = 2147483647
    var v35 uint8 = 42
    var v36 uint64 = 597
    fmt.Println(v27, v27, v36)
    if v33 {
        v2 -= 0.0
        v2 = ((((((-2.25 - float32(((1e10 - float64((1.5 / v27))) * (-2.25 - (float64(3.14159) - (((1e10 / (1.5 + 1.5)) - 1.5) / 1.0)))))) + -2.25) + v27) * (float32(0.0) + v27)) / float32(3.14159)) * v2)
        v37 := GenS5{F0: (v34 << ((839) & 63))}
        _ = v37
    } else {
        if v33 {
            fmt.Println(v28, v2, v27)
            if v33 {
                var v38 int32 = v1
                v33 = v33
                _ = v38
            } else {
                var v39 int16 = 42
                switch 0 {
                case 288:
                    if (!false) {
                        var v40 int64 = 0
                        v41 := GenS1{F0: (^(v1)), F1: int8(v39), F2: (v35 & uint8(v36))}
                        v42 := []int{int(v39), -9223372036854775808, -9223372036854775808, 393}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                case 9223372036854775807:
                    fmt.Println(v33, v28)
                }
                v43 := []int{32767}
                _ = v39
                _ = v43
            }
        }
        var v44 uint64 = (^(v36))
        _ = v44
    }
    switch (((("" + (v29 + v29)) + v29) + (v29 + "foo bar")) + v29) {
    case "\n\t":
        if v28 {
            _ = len(v3)
            var v45 uint64 = (^(v36))
            for fi46 := 0; fi46 < 16; fi46++ {
                for fi47 := 0; fi47 < 10; fi47++ {
                    v28 = v28
                    switch fi46 {
                    case -72:
                        var v48 string = (v29 + ((v29 + "\n\t") + v29))
                        _ = v48
                    default:
                        fmt.Println(fi46)
                    }
                    _ = fi47
                }
                _ = fi46
            }
            _ = v45
        } else {
            v33 = (!(!v33))
            for fi49 := 0; fi49 < 15; fi49++ {
                if (!(v2 < (v2 * float32(((1.5 + (0.1 / float64(0.1))) * 0.1))))) {
                    var v50 uint16 = 495
                    v33 = (!false)
                    v51 := GenS6{F0: 1, F1: -1}
                    _ = v50
                    _ = v51
                }
                for fi52 := 0; fi52 < 11; fi52++ {
                    fmt.Println(v29, v36)
                    _ = fi52
                }
                _ = fi49
            }
            var v53 float32 = v27
            _ = v53
        }
    case "hello":
        for fi54 := 0; fi54 < 19; fi54++ {
            for fi55 := 0; fi55 < 4; fi55++ {
                if false {
                    var v56 int64 = 42
                    _ = v56
                }
                _ = fi55
            }
            fmt.Println(v28, fi54, v35)
            v27 = v2
            _ = fi54
        }
    default:
        switch 42 {
        case 128:
            v28 = v28
        case 4294967295:
            v2 = v27
        case -9223372036854775808:
            switch "\n\t" {
            case "a":
                for fi57 := 0; fi57 < 7; fi57++ {
                    for fi58 := 0; fi58 < 12; fi58++ {
                        v2 -= 1.5
                        v35 = (v35 & 255)
                        _ = fi58
                    }
                    var v59 uint16 = 32767
                    for fi60 := 0; fi60 < 10; fi60++ {
                        fmt.Println(v1, fi57, v34)
                        var v61 uint8 = v35
                        if v28 {
                            fmt.Println(v2)
                            v62 := GenS2{F0: (fi60 + -1), F1: (v34 | v34)}
                            v35 += (v61 / (((v35 ^ v61)) | 1))
                            delete(v3, v34)
                            v59 = 42
                            var v63 uint32 = (v34 | v34)
                            _ = v62
                            _ = v63
                        }
                        _ = fi60
                        _ = v61
                    }
                    _ = fi57
                    _ = v59
                }
            case "foo bar":
                var v64 bool = (v33 && false)
                _ = v64
            default:
                v65 := GenS2{F0: 9223372036854775807, F1: v34}
                _ = v65
            }
        default:
            v66 := []int{4294967295, -2147483648}
            _ = v66
        }
    }
    v3[1] = ((127 <= 0) || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}

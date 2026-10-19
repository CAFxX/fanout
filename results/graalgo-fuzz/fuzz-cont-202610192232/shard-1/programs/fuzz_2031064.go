package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint; F1 uint }
func genFn0(p0 int32, p1 int64, p2 uint64) uint64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5)
        p2 = (p2 / ((p2) | 1))
        _ = fi5
    }
    if false {
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(p0)
            p0 += (p0 | -945)
            if ((!(false || false)) || false) {
                var v7 int32 = (p0 + (-(p0)))
                _ = v7
            }
            _ = fi6
        }
        v8 := GenS1{F0: 246, F1: p2, F2: uint32(p2)}
        switch "\n\t" {
        case "a":
            fmt.Println(p0)
        case "hello":
            var v9 uint8 = 0
            _ = v9
        case "\n\t":
            v10 := GenS2{F0: -9223372036854775808}
            _ = v10
        default:
            p0 -= (p0 / (((p0 << ((114) & 63))) | 1))
        }
        _ = v8
    } else {
        switch 42 {
        case 958:
            p1 = p1
        case -728:
            var v11 float32 = 1e10
            _ = v11
        case 127:
            p2 -= p2
        default:
            switch "foo bar" {
            case "hello":
                v12 := []int{128, 0, int(p0), -1}
                _ = v12
            }
        }
    }
    fmt.Println(p0, p0)
    v13 := GenS2{F0: (p1 >> ((255) & 63))}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p2
}
func genFn1(p0 string, p1 uint, p2 int32) int8 {
    for fi14 := 0; fi14 < 6; fi14++ {
        p1 = 888
        _ = fi14
    }
    if false {
        p2 = p2
    } else {
        if true {
            if (p2 < p2) {
                p2 += 771
                v15 := GenS2{F0: 268}
                var v16 uint8 = 42
                _ = v15
                _ = v16
            } else {
                var v17 bool = (2 < uint64(p1))
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if false {
                var v19 int8 = 0
                v20 := GenS1{F0: 32767, F1: genFn0((p2 & 42), 0, 18446744073709551615), F2: 2147483648}
                _ = v19
                _ = v20
            } else {
                v21 := []int{}
                var v22 uint = 65535
                v22 = (v22 ^ p1)
                _ = v21
                _ = v22
            }
            var v23 float32 = 1.5
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 int16 = 32767
                _ = fi24
                _ = v25
            }
            _ = fi18
            _ = v23
        }
        var v26 int16 = 1
        _ = v26
    }
    var v27 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return -48
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v1 += (((float32(0.0) - ((0.1 * v1) * (v1 + float32(-2.25)))) * float32(3.14159)) * (-2.25 + float32(-2.25)))
    if false {
        v28 := GenS3{F0: 32767, F1: 1, F2: 2}
        for fi29 := 0; fi29 < 15; fi29++ {
            for fi30 := 0; fi30 < 14; fi30++ {
                v31 := GenS3{F0: fi30, F1: 1, F2: 2147483648}
                var v32 float32 = 0.0
                var v33 int = fi30
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi29
        }
        var v34 uint32 = 65535
        _ = v28
        _ = v34
    } else {
        var v35 bool = ((2 >= 948) || (((1 < 698) && (!(!((606 < 1) || (genFn0(2147483647, -9223372036854775808, 633) == 42))))) && true))
        _ = v35
    }
    v1 = ((v1 - (1.5 - (v1 + float32(-2.25)))) + float32(-2.25))
    switch 1 {
    case -2147483648:
        var v36 int8 = 42
        _ = v36
    case -1:
        v37 := GenS4{F0: 0}
        _ = v37
    }
    var v38 float32 = 3.14159
    v38 = 3.14159
    var v39 bool = true
    var v40 string = "hello"
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS2{F0: 449}
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        v44 := v2[("" + v40)]
        switch v40 {
        case "\n\t":
            v45 := v2["a"]
            _ = v45
        case "a":
            for fi46 := 0; fi46 < 6; fi46++ {
                var v47 float32 = v38
                fmt.Println(v40, fi43, fi43)
                _ = fi46
                _ = v47
            }
        default:
            fmt.Println(v1, v38)
        }
        fmt.Println(fi43, v40)
        _ = fi43
        _ = v44
    }
    var v48 uint = 42
    fmt.Println(v40, v1)
    if v39 {
        for fi49 := 0; fi49 < 2; fi49++ {
            if (484 != 42) {
                fmt.Println(v40)
                v38 += v1
                var v50 int = fi49
                _ = v50
            } else {
                delete(v3, (fi49 & 22))
                var v51 uint16 = 255
                _ = v51
            }
            _ = fi49
        }
    } else {
        switch "a" {
        case "foo bar":
            v48 += (v48 / ((v48) | 1))
        case "\n\t":
            if true {
                if true {
                    var v52 uint = v48
                    if v39 {
                        if true {
                            var v53 int = -2147483648
                            v54 := GenS2{F0: 420}
                            v55 := []int{9223372036854775807}
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                    } else {
                        for fi56 := 0; fi56 < 7; fi56++ {
                            var v57 uint8 = 0
                            _ = fi56
                            _ = v57
                        }
                        v58 := GenS5{F0: 2, F1: 0, F2: -265}
                        _ = v58
                    }
                    _ = v52
                }
                if false {
                    fmt.Println(v40)
                    var v59 int64 = 0
                    switch 32767 {
                    case 778:
                        v60 := GenS6{F0: 0, F1: v48}
                        _ = v60
                    case 1:
                        fmt.Println(v40, v48)
                    default:
                        var v61 int64 = (v59 ^ -709)
                        _ = v61
                    }
                    _ = v59
                }
                var v62 string = ("hello" + (v40 + v40))
                _ = v62
            } else {
                v40 = ""
                v63 := []int{2147483647, 32768}
                for fi64 := 0; fi64 < 11; fi64++ {
                    v48 += 65535
                    _ = fi64
                }
                _ = v63
            }
        case "":
            for fi65 := 0; fi65 < 10; fi65++ {
                switch 32767 {
                case 32768:
                    if true {
                        v39 = (2 <= 1)
                        var v66 uint16 = 255
                        _ = v66
                    }
                case -540:
                    v39 = ((float64(1e10) / 1.0) < 1e10)
                }
                switch fi65 {
                case 0:
                    _ = len(v2)
                default:
                    for fi67 := 0; fi67 < 7; fi67++ {
                        var v68 int8 = -105
                        v1 -= ((v38 * ((v1 * (v1 - ((v38 - (3.14159 / v1)) / 1.5))) / v1)) + ((1.5 - (1e10 / float32(-2.25))) * ((((0.0 - ((v1 / (v1 * v1)) - v1)) - v1) * v1) / float32(float64((v1 + (v38 - v38)))))))
                        _ = len(v2)
                        _ = fi67
                        _ = v68
                    }
                }
                var v69 uint8 = 127
                _ = fi65
                _ = v69
            }
        default:
            v48 = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v48
}

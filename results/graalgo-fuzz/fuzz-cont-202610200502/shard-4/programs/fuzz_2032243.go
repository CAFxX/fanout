package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint16; F1 uint32 }
func genFn0(p0 bool, p1 int16, p2 uint32) uint64 {
    var v6 uint32 = 65535
    p2 = p2
    fmt.Println(p2)
    v7 := []int{32767}
    if len(v7) > 0 {
        v7[0] = 4294967295
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 int32, p2 bool) uint8 {
    fmt.Println(p2, p2)
    p0 = p0
    var v8 float64 = -2.25
    v8 += v8
    v9 := []int{32767, -546}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint32 = 2
    var v2 float32 = -2.25
    v3 := make(map[int]string)
    var v4 uint64 = 0
    var v5 uint32 = v1
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int = (-(fi11))
            _ = fi11
            _ = v12
        }
        v5 -= v1
        var v13 int = (fi10 << ((0) & 63))
        _ = fi10
        _ = v13
    }
    var v14 uint = 65535
    v15 := GenS1{F0: 9223372036854775807, F1: 127, F2: ("hello" != ("foo bar" + ("" + (((("\n\t" + "\n\t") + ("foo bar" + "foo bar")) + "hello") + "a"))))}
    v2 += ((float32(-2.25) * float32((((1e10 / -2.25) / (float64((v2 - ((0.0 / float32(((3.14159 / ((1e10 / (0.1 * (float64(((1e10 * (float32(0.1) * ((((((0.1 / 1.0) - float32(1.5)) - 0.0) + float32((((1e10 / 1e10) / 0.1) - 1e10))) * (v2 * v2)) - v2))) - v2)) * float64(((0.0 / ((-2.25 + 1e10) - v2)) * (1.5 + v2)))))) * -2.25)) * (1e10 * (1e10 / 3.14159))))) + float32((-2.25 / ((-2.25 * (0.1 / float64(v2))) + float64((float32(float64(v2)) * v2)))))))) - (float64(((v2 - 3.14159) + (((v2 + 0.0) + v2) / float32((1.5 + float64(v2)))))) - 0.1))) + 0.0))) / (v2 / v2))
    v1 = (^(v5))
    v4 = (v4 + 4294967295)
    var v16 int8 = 42
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := []int{(fi17 / (((fi17 * (fi17 + fi17))) | 1)), fi17, (fi17 ^ 327)}
        if false {
            v5 += v1
            v2 += 1e10
        } else {
            if true {
                v19 := GenS2{F0: 65535, F1: -2147483648}
                _ = v19
            } else {
                v20 := GenS2{F0: 0, F1: 42}
                switch ("\n\t" + ((("\n\t" + ("\n\t" + "")) + "hello") + "a")) {
                case "\n\t":
                    var v21 int = fi17
                    _ = v21
                case "foo bar":
                    v16 = (v16 % ((v16) | 1))
                case "hello":
                    for fi22 := 0; fi22 < 9; fi22++ {
                        v18 = append(v18, fi17)
                        _ = fi22
                    }
                default:
                    v2 = 0.1
                }
                _ = v20
            }
            v23 := GenS3{F0: uint8(v14), F1: (v16 << (((v14 | (v14 & (v14 | v14)))) & 63))}
            var v24 int64 = 9223372036854775807
            _ = v23
            _ = v24
        }
        v25 := GenS4{F0: 32767, F1: false}
        _ = fi17
        _ = v18
        _ = v25
    }
    var v26 uint16 = 372
    var v27 uint16 = v26
    switch 0 {
    case -799:
        fmt.Println(v16, v26)
    case 320:
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 float64 = float64((1e10 - v2))
            v3[(fi28 << (((v14 & (v14 - 42))) & 63))] = "\n\t"
            v14 = uint(v1)
            _ = fi28
            _ = v29
        }
    case 127:
        v30 := []int{-370, int(v4), 32767}
        _ = v30
    default:
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := []int{(fi31 / (((fi31 ^ fi31)) | 1))}
            _ = len(v32)
            v33 := GenS5{F0: 2, F1: v27}
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    var v34 bool = ((0.1 * (v2 * (0.1 - 3.14159))) > v2)
    for fi35 := 0; fi35 < 16; fi35++ {
        fmt.Println(v2)
        if v34 {
            v36 := GenS5{F0: 0, F1: v26}
            switch fi35 {
            case -1:
                var v37 bool = (!(657 < 42))
                _ = v37
            case 2147483647:
                switch "hello" {
                case "hello":
                    fmt.Println(fi35, v4)
                case "":
                    var v38 uint64 = (v4 | uint64(v5))
                    _ = v38
                }
            case 32768:
                var v39 uint64 = v4
                _ = v39
            default:
                if v34 {
                    v40 := GenS6{F0: 42, F1: int16(v16)}
                    _ = v40
                }
            }
            for fi41 := 0; fi41 < 12; fi41++ {
                v3[fi35] = "a"
                var v42 bool = v34
                var v43 int16 = 1
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = v36
        } else {
            var v44 uint8 = 128
            v45 := []int{fi35}
            _ = v44
            _ = v45
        }
        _ = fi35
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        v27 = (v27 + 65535)
        v47 := []int{fi46, (^(fi46)), (-(fi46)), 32768}
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        for fi49 := 0; fi49 < 20; fi49++ {
            fmt.Println(v4)
            _ = fi49
        }
        v27 -= 255
        delete(v3, fi48)
        _ = fi48
    }
    v50 := GenS7{F0: (^(v27)), F1: 814}
    if (!(892 != 32768)) {
        _ = len(v3)
        switch ("foo bar" + "") {
        case "foo bar":
            v1 += 42
        case "a":
            switch -1 {
            case 0:
                var v51 uint32 = v5
                _ = v51
            case 1:
                v5 = (^(v5))
            default:
                for fi52 := 0; fi52 < 3; fi52++ {
                    v53 := GenS7{F0: v27, F1: v5}
                    switch fi52 {
                    case 150:
                        v54 := v3[fi52]
                        _ = v54
                    case 502:
                        for fi55 := 0; fi55 < 15; fi55++ {
                            _ = len(v3)
                            _ = fi55
                        }
                    default:
                        v56 := GenS4{F0: 18, F1: v34}
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                }
            }
        default:
            var v57 int8 = (v16 / ((v16) | 1))
            _ = v57
        }
        if (2147483647 <= 9223372036854775807) {
            v2 = (-2.25 - 1.5)
        } else {
            v3[-935] = "a"
            var v58 int32 = 42
            switch "foo bar" {
            case "\n\t":
                v4 -= v4
            }
            _ = v58
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v50
}

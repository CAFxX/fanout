package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 int }
func genFn0(p0 int16) int64 {
    if (690 == 4294967295) {
        switch "a" {
        case "hello":
            var v4 int = -134
            _ = v4
        }
        v5 := GenS1{F0: 100}
        _ = v5
    }
    var v6 int16 = (-(p0))
    v6 = v6
    switch 32767 {
    case 42:
        var v7 float64 = float64(1.5)
        _ = v7
    case 32768:
        if (2 != 2) {
            var v8 uint64 = 611
            p0 = 716
            _ = v8
        }
    case -939:
        fmt.Println(p0, p0)
    default:
        switch "" {
        case "hello":
            v9 := []int{953, 128, 720, 0}
            _ = v9
        default:
            var v10 uint32 = 2147483647
            _ = v10
        }
    }
    _ = p0
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("foo bar" + "a"), "foo bar"}
    v3 := []int{}
    v3 = append(v3, 127)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int32 = -2147483648
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 1, F1: 0}
    if false {
        v14 := v1[1]
        v14 = v14
        if (v14 && (42 > 1)) {
            var v15 bool = (true || (-128 < 0))
            var v16 uint8 = 2
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v15
            _ = v16
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                var v17 string = ""
                fmt.Println(v17)
                switch 532 {
                case 127:
                    v18 := GenS1{F0: 0}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v14 = v14
                        var v20 int64 = -868
                        for fi21 := 0; fi21 < 11; fi21++ {
                            v14 = (v17 < (v17 + v17))
                            v3 = append(v3, fi21)
                            v22 := GenS1{F0: -128}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                _ = v17
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v23 := GenS3{F0: 186, F1: 1, F2: -128}
                _ = v23
            }
        }
        _ = v14
    } else {
        _ = len(v1)
    }
    if (495 != 165) {
        v24 := GenS4{F0: (42 > 2147483647)}
        if true {
            var v25 float32 = float32(-2.25)
            v26 := GenS2{F0: 42, F1: -1}
            _ = v25
            _ = v26
        }
        v27 := GenS5{F0: 927}
        _ = v24
        _ = v27
    }
    if true {
        v28 := GenS4{F0: true}
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint8 = 255
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    if ((((-2.25 / 1.0) < float32(((0.1 - 3.14159) + 0.1))) || true) && true) {
        for fi31 := 0; fi31 < 15; fi31++ {
            switch ("foo bar" + "foo bar") {
            case "a":
                var v32 int16 = 1000
                _ = v32
            case "foo bar":
                var v33 uint32 = 654
                _ = v33
            case "":
                var v34 uint16 = 65535
                _ = v34
            default:
                switch "" {
                case "":
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v35 int32 = 42
                        v36 := GenS5{F0: fi31}
                        _ = v35
                        _ = v36
                    } else {
                        v1[1] = (255 <= 42)
                        var v37 string = "a"
                        if false {
                            var v38 uint16 = 2
                            _ = v38
                        } else {
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            var v39 float64 = ((0.1 - float64((float32(((float64((3.14159 - (0.1 * float32((float64(float32(3.14159)) + float64(((0.0 - float32((1.5 * (-2.25 / (float64(0.1) + 3.14159))))) + (-2.25 - float32(-2.25))))))))) / -2.25) / 1.0)) * (1e10 / ((0.1 / ((1.5 - (((1e10 - float32(1e10)) + 0.0) * 0.1)) - (-2.25 - 0.0))) + 1.5))))) * 1.5)
                            _ = v39
                        }
                        _ = v37
                    }
                case "foo bar":
                    switch "hello" {
                    case "foo bar":
                        for fi40 := 0; fi40 < 6; fi40++ {
                            if len(v3) > 0 {
                                v3[0] = fi40
                            }
                            _ = fi40
                        }
                    }
                }
            }
            var v41 int8 = 127
            var v42 int16 = 0
            _ = fi31
            _ = v41
            _ = v42
        }
        v43 := GenS6{F0: 64, F1: -9223372036854775808}
        var v44 int32 = -2147483648
        _ = v43
        _ = v44
    }
    var v45 float32 = (1.5 + 1.5)
    if (!true) {
        v45 = v45
    }
    if false {
        _ = len(v2)
        v46 := GenS3{F0: 2, F1: 2147483647, F2: 0}
        _ = v46
    }
    switch -133 {
    case -1:
        var v47 int = -474
        _ = v47
    case 794:
        v2 = append(v2, ("foo bar" + "foo bar"))
    case -9223372036854775808:
        var v48 uint8 = 255
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        if len(v3) > 0 {
            v3[0] = 0
        }
        _ = fi49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v45
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0() int64 {
    var v6 int8 = 42
    fmt.Println(v6)
    if (-67 != 127) {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 613}
            if false {
                v9 := []int{fi7, 128}
                v10 := GenS1{F0: 0}
                var v11 uint16 = 42
                _ = v10
                _ = v11
                _ = v9
            } else {
                v12 := []int{}
                if true {
                    var v13 int32 = -1
                    _ = v13
                } else {
                    if len(v12) > 0 {
                        v12[0] = fi7
                    }
                }
                _ = v12
            }
            v6 = -16
            _ = fi7
            _ = v8
        }
        var v14 int8 = (-(v6))
        _ = v14
    }
    v6 = v6
    v6 = (v6 >> ((655) & 63))
    _ = v6
    return 974
}
func genFn1() int16 {
    if true {
        var v15 bool = (!true)
        _ = v15
    } else {
        var v16 uint8 = 255
        for fi17 := 0; fi17 < 2; fi17++ {
            switch (-(fi17)) {
            case -9223372036854775808:
                switch "foo bar" {
                case "a":
                    switch 9223372036854775807 {
                    case 128:
                        v16 += (v16 | (v16 & (v16 & (v16 & v16))))
                    case 127:
                        v18 := []int{(-(fi17)), 127, fi17, fi17}
                        _ = v18
                    case 32768:
                        for fi19 := 0; fi19 < 7; fi19++ {
                            switch "" {
                            case "foo bar":
                                var v20 int8 = 101
                                v21 := GenS1{F0: -1}
                                _ = v20
                                _ = v21
                            default:
                                v22 := []int{(fi17 % (((-(fi19))) | 1))}
                                _ = v22
                            }
                            v16 -= 127
                            _ = fi19
                        }
                    }
                default:
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := []int{fi17, fi17}
                        _ = fi23
                        _ = v24
                    }
                }
            default:
                var v25 float64 = (1.5 * 0.1)
                _ = v25
            }
            var v26 float32 = (float32(1.5) - float32(-2.25))
            if false {
                switch (^(fi17)) {
                case -1:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 uint32 = 377
                        v29 := GenS1{F0: -954}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                case 1:
                    v16 = (^(v16))
                default:
                    v30 := GenS1{F0: 42}
                    _ = v30
                }
            }
            _ = fi17
            _ = v26
        }
        _ = v16
    }
    v31 := GenS1{F0: -834}
    var v32 uint16 = 2
    v32 += 0
    fmt.Println(v32)
    _ = v31
    _ = v32
    return -965
}

func main() {
    var v1 bool = false
    var v2 uint8 = 0
    var v3 string = "hello"
    var v4 float32 = 1e10
    v5 := []int{127, 42, 4294967295, 9223372036854775807}
    if (42 >= (v2 % (((v2 | (v2 - (^(v2))))) | 1))) {
        switch -65 {
        case 127:
            if (!(0.0 == (1e10 / (1e10 - 1.5)))) {
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 int = fi33
                    _ = fi33
                    _ = v34
                }
                if false {
                    var v35 uint32 = 784
                    var v36 float64 = (0.0 / (0.1 / 1.5))
                    _ = v35
                    _ = v36
                } else {
                    var v37 string = (v3 + v3)
                    var v38 uint64 = 0
                    _ = v37
                    _ = v38
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case -9223372036854775808:
            v1 = true
        case 128:
            switch 128 {
            case 32768:
                fmt.Println(v1, v3, v1)
            case 0:
                var v39 float64 = 0.1
                _ = v39
            }
        default:
            if true {
                if v1 {
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                    for fi40 := 0; fi40 < 18; fi40++ {
                        _ = len(v5)
                        _ = fi40
                    }
                    if (32767 != 0) {
                        switch int(v2) {
                        case 9223372036854775807:
                            switch 2147483647 {
                            case -516:
                                var v41 uint = 981
                                _ = v41
                            }
                        default:
                            var v42 uint64 = 867
                            _ = v42
                        }
                        v2 = (^(v2))
                        var v43 int16 = genFn1()
                        _ = v43
                    }
                } else {
                    v44 := GenS1{F0: -162}
                    v4 -= v4
                    _ = v44
                }
            }
        }
        var v45 uint32 = 828
        _ = v45
    } else {
        for fi46 := 0; fi46 < 2; fi46++ {
            v47 := GenS1{F0: 42}
            v4 = (float32(3.14159) / v4)
            _ = fi46
            _ = v47
        }
        v4 = v4
    }
    v4 -= v4
    var v48 uint64 = 2
    v49 := GenS2{F0: v1, F1: 39}
    var v50 float32 = float32(3.14159)
    switch (v3 + "\n\t") {
    case "hello":
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 uint16 = 2
            _ = fi51
            _ = v52
        }
    case "a":
        for fi53 := 0; fi53 < 2; fi53++ {
            v3 = v3
            _ = fi53
        }
    default:
        v50 = (v50 + v50)
    }
    for fi54 := 0; fi54 < 10; fi54++ {
        for fi55 := 0; fi55 < 19; fi55++ {
            if len(v5) > 0 {
                v5[0] = 0
            }
            if v1 {
                _ = len(v5)
            }
            if len(v5) > 0 {
                v5[0] = (fi55 >> ((2) & 63))
            }
            _ = fi55
        }
        _ = fi54
    }
    _ = len(v5)
    for fi56 := 0; fi56 < 18; fi56++ {
        v2 = (v2 * (v2 & 42))
        _ = fi56
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        switch v3 {
        case "hello":
            var v58 uint8 = v2
            _ = v58
        case "a":
            v3 = (v3 + (v3 + v3))
        }
        v59 := GenS3{F0: 2, F1: -808, F2: 255}
        _ = fi57
        _ = v59
    }
    var v60 int8 = 8
    if true {
        var v61 float64 = 0.0
        _ = v61
    }
    fmt.Println(v3)
    switch v3 {
    case "a":
        for fi62 := 0; fi62 < 20; fi62++ {
            v60 = 42
            _ = fi62
        }
    case "foo bar":
        v1 = v1
    default:
        switch v3 {
        case "":
            v5 = append(v5, -9223372036854775808)
        case "foo bar":
            v63 := GenS2{F0: (-1 != 0), F1: (v60 - (v60 - v60))}
            _ = v63
        case "\n\t":
            v64 := GenS3{F0: 2147483647, F1: int32(v60), F2: 984}
            _ = v64
        default:
            var v65 int = -2
            _ = v65
        }
    }
    v66 := GenS4{F0: (579 != 0), F1: 2147483648, F2: 42}
    var v67 string = "foo bar"
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v60:", v60)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v60
    _ = v66
    _ = v67
}

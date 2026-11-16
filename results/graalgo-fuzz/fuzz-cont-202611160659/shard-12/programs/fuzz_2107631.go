package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 int; F2 int8 }
type GenS3 struct { F0 string; F1 int64; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0() int16 {
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int8 = 8
            v9 := []int{(fi7 / ((2147483647) | 1)), fi7}
            if (-865 < int16(v8)) {
                if ((!false) && false) {
                    fmt.Println(v8, fi7)
                    v8 = (v8 >> ((18446744073709551615) & 63))
                    v10 := GenS1{F0: 18446744073709551615, F1: v8, F2: 66}
                    _ = v10
                } else {
                    v8 += 127
                    var v11 int32 = -375
                    _ = v11
                }
                v12 := GenS1{F0: 0, F1: -128, F2: 928}
                var v13 int64 = -335
                _ = v12
                _ = v13
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        v14 := []int{}
        _ = v14
    } else {
        var v15 bool = false
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if ((float32(0.1) - 1.5) == (0.0 + float32(1e10))) {
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 uint = 941
                _ = fi17
                _ = v18
            }
            var v19 int32 = -850
            _ = v19
        } else {
            var v20 int = (fi16 & 32767)
            v21 := GenS1{F0: 1, F1: 1, F2: 2147483647}
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float32 = 3.14159
            var v24 uint16 = 899
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := []int{}
        _ = fi16
        _ = v25
    }
    switch -137 {
    case 943:
        var v26 int8 = 1
        _ = v26
    case 0:
        var v27 uint8 = 1
        _ = v27
    case 128:
        if true {
            switch -9223372036854775808 {
            case -1:
                var v28 int32 = -2147483648
                _ = v28
            case 128:
                for fi29 := 0; fi29 < 13; fi29++ {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        fmt.Println(fi29, fi30)
                        if true {
                            var v31 int64 = -9223372036854775808
                            v32 := []int{fi29}
                            fmt.Println(fi29, v31, fi29)
                            _ = v31
                            _ = v32
                        } else {
                            var v33 int8 = 0
                            v34 := []int{fi30, (fi30 % (((fi29 + fi29)) | 1)), fi29, 0}
                            fmt.Println(v33, v33)
                            _ = v33
                            _ = v34
                        }
                        _ = fi30
                    }
                    _ = fi29
                }
            case 0:
                var v35 int16 = 1
                _ = v35
            default:
                v36 := []int{0, -2147483648, 140, 4294967295}
                _ = v36
            }
            for fi37 := 0; fi37 < 4; fi37++ {
                for fi38 := 0; fi38 < 7; fi38++ {
                    v39 := []int{fi37, fi38}
                    v40 := GenS2{F0: "a", F1: (fi38 | (^(fi38))), F2: 42}
                    v41 := GenS2{F0: "\n\t", F1: fi38, F2: -1}
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                var v42 uint16 = 228
                var v43 int32 = 1
                _ = fi37
                _ = v42
                _ = v43
            }
        }
    default:
        v44 := GenS3{F0: "foo bar", F1: -1, F2: 9223372036854775807}
        _ = v44
    }
    var v45 int32 = 42
    switch -2147483648 {
    case 4294967295:
        v46 := GenS1{F0: 2, F1: 127, F2: 2147483648}
        _ = v46
    }
    _ = v45
    return 675
}
func genFn1(p0 int64, p1 uint) uint16 {
    v47 := GenS4{F0: 868}
    v48 := []int{128, 440, -9223372036854775808}
    switch "a" {
    case "\n\t":
        _ = len(v48)
    case "a":
        switch "\n\t" {
        case "\n\t":
            var v49 float64 = 3.14159
            _ = v49
        case "hello":
            for fi50 := 0; fi50 < 17; fi50++ {
                v48 = append(v48, int(p0))
                for fi51 := 0; fi51 < 3; fi51++ {
                    var v52 uint64 = 18446744073709551615
                    _ = fi51
                    _ = v52
                }
                _ = fi50
            }
        case "foo bar":
            for fi53 := 0; fi53 < 20; fi53++ {
                var v54 uint16 = 255
                var v55 int64 = p0
                v48 = append(v48, 127)
                _ = fi53
                _ = v54
                _ = v55
            }
        default:
            var v56 float64 = (((1e10 * float64((((0.1 * (float32(float64(3.14159)) + ((((float32(3.14159) / 0.1) / 1.0) - 3.14159) + float32(((3.14159 + 1e10) + (0.0 - ((1.5 + 3.14159) * 0.0))))))) + 1.5) + (float32((1.5 * ((3.14159 / 0.1) * float64(float32(1.5))))) - 0.0)))) - 0.1) / 3.14159)
            _ = v56
        }
    default:
        p0 = p0
    }
    p0 += (p0 | p0)
    _ = p0
    _ = p1
    _ = v47
    _ = v48
    return 1
}

func main() {
    var v1 uint64 = 0
    var v2 int64 = -9223372036854775808
    var v3 int64 = (v2 - int64(v1))
    var v4 float32 = (float32(float64(1.5)) * 1.5)
    v5 := []string{"\n\t"}
    v6 := []string{("\n\t" + (("hello" + ("hello" + ("a" + ("" + "\n\t")))) + ("hello" + "")))}
    v57 := GenS1{F0: 700, F1: 127, F2: 0}
    var v58 uint64 = v1
    var v59 string = "a"
    switch ("" + ("" + ("hello" + v59))) {
    case "":
        if len(v5) > 0 {
            v5[0] = v59
        }
    }
    v60 := GenS5{F0: v2}
    var v61 uint8 = 127
    v61 = v61
    for fi62 := 0; fi62 < 11; fi62++ {
        v63 := GenS6{F0: 996, F1: 246, F2: 18446744073709551615}
        for fi64 := 0; fi64 < 9; fi64++ {
            v6 = append(v6, (("foo bar" + ("\n\t" + (v59 + v59))) + "\n\t"))
            v4 = v4
            _ = fi64
        }
        var v65 int16 = 0
        _ = fi62
        _ = v63
        _ = v65
    }
    var v66 uint32 = 41
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch -621 {
    case 180:
        var v67 uint64 = 4294967295
        _ = v67
    case 2147483647:
        for fi68 := 0; fi68 < 17; fi68++ {
            v1 -= v58
            _ = fi68
        }
    case 42:
        var v69 int8 = 1
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v66
}

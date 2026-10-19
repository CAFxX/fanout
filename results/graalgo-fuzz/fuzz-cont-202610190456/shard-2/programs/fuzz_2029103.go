package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint; F1 int; F2 uint }
type GenS5 struct { F0 int32; F1 uint64; F2 int64 }
type GenS6 struct { F0 int16; F1 int32; F2 int }
func genFn0() int8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        if ("a" >= "hello") {
            switch 32767 {
            case -799:
                var v7 int16 = 0
                _ = v7
            case 2147483647:
                v8 := []int{(fi6 & -945), (fi6 & (fi6 / ((-872) | 1)))}
                _ = v8
            case -768:
                for fi9 := 0; fi9 < 16; fi9++ {
                    fmt.Println(fi9, fi9)
                    var v10 uint16 = 643
                    fmt.Println(fi6)
                    _ = fi9
                    _ = v10
                }
            }
        } else {
            var v11 int32 = 1
            if (false || true) {
                v12 := GenS1{F0: (false && true)}
                _ = v12
            } else {
                v13 := []int{(fi6 >> ((42) & 63)), (fi6 ^ (fi6 & fi6)), (fi6 ^ -322), fi6}
                v13 = append(v13, -2147483648)
                _ = v13
            }
            var v14 int16 = 337
            _ = v11
            _ = v14
        }
        v15 := GenS1{F0: false}
        var v16 uint64 = 18446744073709551615
        _ = fi6
        _ = v15
        _ = v16
    }
    var v17 int8 = 100
    v18 := []int{0}
    v17 += -1
    var v19 bool = (0 >= -9223372036854775808)
    _ = v17
    _ = v18
    _ = v19
    return v17
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 string = "a"
    var v5 float32 = -2.25
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            if (!true) {
                switch v1 {
                case "hello":
                    v22 := GenS1{F0: (v4 < ("a" + v1))}
                    _ = v22
                case "":
                    v5 += v5
                }
                if (v5 > v5) {
                    v1 = ((v1 + ("" + v1)) + ("foo bar" + v4))
                    var v23 int = (fi21 % ((1) | 1))
                    v1 = ""
                    _ = v23
                }
            }
            v24 := []int{fi21}
            _ = len(v24)
            _ = fi21
            _ = v24
        }
        if false {
            var v25 uint32 = 150
            switch (fi20 - fi20) {
            case 128:
                var v26 int32 = 42
                _ = v26
            case 9223372036854775807:
                var v27 uint = 1
                _ = v27
            }
            switch 2147483647 {
            case 32767:
                fmt.Println(v25)
            case -2147483648:
                if false {
                    v28 := v2[(v25 + 2)]
                    var v29 bool = ("hello" <= (v4 + ("" + v4)))
                    v30 := GenS2{F0: v29}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    var v31 uint32 = v25
                    _ = v31
                }
            }
            _ = v25
        }
        if true {
            for fi32 := 0; fi32 < 11; fi32++ {
                v4 = "foo bar"
                v33 := []int{fi32}
                v1 = ("hello" + v4)
                _ = fi32
                _ = v33
            }
            var v34 float32 = v5
            for fi35 := 0; fi35 < 4; fi35++ {
                _ = len(v2)
                _ = len(v3)
                _ = fi35
            }
            _ = v34
        } else {
            if (65535 == 267) {
                v36 := []int{}
                var v37 uint8 = 251
                _ = v36
                _ = v37
            }
            v38 := []int{fi20, (fi20 + -11)}
            switch 32768 {
            case 178:
                switch ((((v1 + "") + "") + "") + v1) {
                case "\n\t":
                    if false {
                        v39 := GenS1{F0: (true && true)}
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 16; fi40++ {
                            var v41 int8 = 0
                            fmt.Println(fi40)
                            v4 = v1
                            _ = fi40
                            _ = v41
                        }
                        v1 = ("\n\t" + v1)
                        fmt.Println(v5)
                    }
                case "a":
                    var v42 float32 = v5
                    _ = v42
                case "foo bar":
                    v43 := GenS3{F0: fi20, F1: 2}
                    _ = v43
                default:
                    for fi44 := 0; fi44 < 6; fi44++ {
                        v1 = v4
                        if (!(1 >= -1)) {
                            v4 = v1
                            delete(v2, 2147483648)
                        } else {
                            v45 := GenS4{F0: 1, F1: (fi20 + (fi20 + (fi44 * -2147483648))), F2: 2}
                            v4 = v4
                            fmt.Println(fi20)
                            _ = v45
                        }
                        fmt.Println(fi20, v5, fi20)
                        _ = fi44
                    }
                }
            case 2147483647:
                v46 := GenS5{F0: 0, F1: 0, F2: 0}
                _ = v46
            default:
                if len(v38) > 0 {
                    v38[0] = fi20
                }
            }
            _ = v38
        }
        _ = fi20
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 uint8 = 78
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        for fi50 := 0; fi50 < 18; fi50++ {
            switch (fi50 + fi50) {
            case 127:
                v4 = v1
            }
            v5 = (1e10 * ((v5 / 3.14159) + v5))
            v4 = v4
            var v51 int = fi49
            var v52 int = (fi50 * fi49)
            var v53 float32 = v5
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
        }
        v4 = ((v1 + "hello") + v4)
        if true {
            fmt.Println(v5)
        } else {
            var v54 float64 = 1.5
            switch fi49 {
            case -359:
                v54 += float64(1e10)
            case 42:
                var v55 int64 = -716
                _ = v55
            }
            _ = v54
        }
        _ = fi49
    }
    v1 = v4
    if true {
        if true {
            var v56 uint8 = 2
            v5 = v5
            _ = v56
        } else {
            v1 = "a"
            v5 = (float32(-2.25) - float32(((((-2.25 * 1.5) + (float64(-2.25) * 1.5)) / 3.14159) + ((float64(3.14159) - (0.0 / (1e10 / (3.14159 * (0.1 / 1e10))))) + (0.0 + -2.25)))))
            if false {
                v2[42] = 360
                fmt.Println(v1, v1, v5)
                var v57 int16 = 767
                _ = v57
            }
        }
    } else {
        var v58 int64 = 1
        var v59 int = -9223372036854775808
        v60 := GenS6{F0: -811, F1: 4, F2: 676}
        _ = v58
        _ = v59
        _ = v60
    }
    v61 := GenS3{F0: 128, F1: 681}
    if ((42 != 220) || true) {
        v1 = v1
    } else {
        if ("\n\t" < ("foo bar" + ((((v1 + v1) + ("foo bar" + v1)) + (v1 + v4)) + ((v4 + (v4 + "foo bar")) + "hello")))) {
            var v62 int64 = -378
            _ = v62
        } else {
            for fi63 := 0; fi63 < 9; fi63++ {
                v64 := []int{}
                if len(v64) > 0 {
                    v64[0] = fi63
                }
                v65 := v2[323]
                _ = fi63
                _ = v64
                _ = v65
            }
            if true {
                var v66 int32 = 2147483647
                _ = v66
            } else {
                var v67 bool = (2 < 4294967295)
                _ = v67
            }
        }
    }
    delete(v2, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}

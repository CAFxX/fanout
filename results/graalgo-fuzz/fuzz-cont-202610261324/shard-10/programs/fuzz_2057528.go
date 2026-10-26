package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8, p1 uint64) uint16 {
    var v7 uint = 65535
    v8 := GenS1{F0: p1, F1: ("hello" + "foo bar"), F2: p0}
    p0 = (p0 / (((p0 * 255)) | 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint64, p1 string) bool {
    p1 = p1
    p0 = p0
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 int16 = 1
    var v2 float32 = 1.5
    var v3 string = ("\n\t" + ("a" + "a"))
    var v4 float64 = 3.14159
    var v5 string = (v3 + v3)
    var v6 float64 = 0.0
    v5 = (("foo bar" + v5) + v5)
    v1 -= (v1 << ((2) & 63))
    if true {
        v9 := []int{-1, 2147483647, 715}
        var v10 uint16 = 65535
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 float32 = float32(v6)
            var v13 uint64 = 2
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    }
    v4 += 3.14159
    switch "\n\t" {
    case "\n\t":
        if (!(true && (false || true))) {
            for fi14 := 0; fi14 < 17; fi14++ {
                if (true || ((false && false) || ((128 >= 192) && (-2147483648 < -48)))) {
                    fmt.Println(v3, v2, v6)
                    if true {
                        var v15 float64 = 1.5
                        v16 := []int{fi14, (fi14 % ((fi14) | 1))}
                        var v17 int = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        switch (fi14 / ((fi14) | 1)) {
                        case 127:
                            var v18 int16 = -195
                            _ = v18
                        case 42:
                            var v19 int64 = 0
                            _ = v19
                        case 32768:
                            v3 = "foo bar"
                        }
                        v3 = v3
                        v1 -= (-(v1))
                    }
                    fmt.Println(v5, v3)
                } else {
                    if false {
                        v20 := []int{}
                        switch (-(fi14)) {
                        case 0:
                            var v21 int16 = v1
                            _ = v21
                        case 127:
                            var v22 uint32 = 284
                            _ = v22
                        case -811:
                            for fi23 := 0; fi23 < 4; fi23++ {
                                if false {
                                    var v24 int = 2147483647
                                    v3 = ""
                                    var v25 uint16 = 255
                                    v6 = v4
                                    _ = v24
                                    _ = v25
                                }
                                var v26 string = v3
                                _ = fi23
                                _ = v26
                            }
                        }
                        if true {
                            if false {
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    fmt.Println(v1)
                                    v1 += (v1 / (((-(v1))) | 1))
                                    _ = fi27
                                }
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    var v29 uint8 = 0
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                v30 := GenS2{F0: v1, F1: 333, F2: 32767}
                                _ = v30
                            }
                            switch fi14 {
                            case 42:
                                fmt.Println(v2)
                            case -667:
                                v1 += v1
                            case 127:
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    v32 := GenS3{F0: -80}
                                    v33 := GenS2{F0: 38, F1: fi14, F2: v1}
                                    var v34 uint64 = 1
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            default:
                                fmt.Println(v1)
                            }
                            fmt.Println(v2)
                        }
                        _ = v20
                    } else {
                        if false {
                            fmt.Println(v2, v1)
                            fmt.Println(v4, fi14, v3)
                            var v35 uint16 = 1
                            _ = v35
                        }
                    }
                }
                _ = fi14
            }
            v36 := GenS2{F0: v1, F1: -2147483648, F2: 134}
            _ = v36
        }
    case "foo bar":
        if true {
            v37 := GenS4{F0: 694, F1: 1}
            v38 := GenS5{F0: 1}
            v1 = 981
            _ = v37
            _ = v38
        } else {
            if false {
                v39 := GenS6{F0: 137}
                if (((0 <= 326) && (0 != 1)) || false) {
                    var v40 uint64 = 1
                    var v41 float32 = float32(v6)
                    v6 = 0.0
                    _ = v40
                    _ = v41
                } else {
                    var v42 int8 = 0
                    if true {
                        if false {
                            v43 := []int{447, 128, 2147483647}
                            _ = v43
                        }
                        v5 = "a"
                        v44 := []int{171, 32767, 4294967295}
                        _ = v44
                    } else {
                        v1 += (v1 >> ((2) & 63))
                        var v45 float32 = v2
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 14; fi46++ {
                        var v47 int64 = -1
                        v4 += v4
                        _ = fi46
                        _ = v47
                    }
                    _ = v42
                }
                v6 = 0.1
                _ = v39
            } else {
                v48 := []int{32767}
                for fi49 := 0; fi49 < 12; fi49++ {
                    for fi50 := 0; fi50 < 17; fi50++ {
                        v5 = "\n\t"
                        var v51 uint64 = 18446744073709551615
                        _ = fi50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 14; fi52++ {
                        if (!(!(true && true))) {
                            switch ("a" + v5) {
                            case "":
                                switch v5 {
                                case "a":
                                    var v53 uint8 = 42
                                    _ = v53
                                case "hello":
                                    fmt.Println(v1, v3)
                                case "\n\t":
                                    var v54 int64 = -803
                                    _ = v54
                                }
                            }
                            if false {
                                v5 = v3
                                v2 = float32(1.5)
                                v1 = v1
                            } else {
                                v6 -= 1.5
                                var v55 float32 = float32(((1e10 / (v6 - 1e10)) - v6))
                                _ = v55
                            }
                            fmt.Println(fi49, v5, v1)
                            var v56 uint8 = 0
                            v6 += v4
                            var v57 bool = false
                            _ = v56
                            _ = v57
                        } else {
                            var v58 int8 = -1
                            var v59 bool = false
                            v60 := GenS2{F0: (v1 % ((v1) | 1)), F1: (fi49 >> ((802) & 63)), F2: (v1 / (((^(v1))) | 1))}
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        if (72 >= 243) {
                            if true {
                                switch "a" {
                                case "\n\t":
                                    if (341 < 177) {
                                        v1 = v1
                                    }
                                case "hello":
                                    if len(v48) >= 2 {
                                        v48 = v48[1:len(v48)]
                                    }
                                case "a":
                                    if len(v48) > 0 {
                                        v48[0] = fi52
                                    }
                                default:
                                    v3 = "\n\t"
                                }
                                v2 += 1e10
                                if false {
                                    if len(v48) > 0 {
                                        v48[0] = fi52
                                    }
                                    v48 = append(v48, (fi49 % (((-(fi52))) | 1)))
                                }
                            } else {
                                v4 -= (0.0 * v4)
                            }
                        }
                        var v61 uint64 = 705
                        var v62 int8 = 100
                        var v63 int64 = -1
                        if (true && false) {
                            if (!(!false)) {
                                var v64 int32 = -2147483648
                                _ = v64
                            } else {
                                v48 = append(v48, (fi52 >> ((0) & 63)))
                                switch v3 {
                                case "\n\t":
                                    v1 = v1
                                }
                            }
                        }
                        _ = fi52
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    _ = fi49
                }
                fmt.Println(v3)
                _ = v48
            }
            v65 := []int{}
            _ = v65
        }
    }
    var v66 uint = 42
    switch 127 {
    case 32768:
        v4 = 0.1
    default:
        switch v3 {
        case "foo bar":
            var v67 float32 = 1e10
            _ = v67
        case "hello":
            v68 := GenS3{F0: 42}
            _ = v68
        }
    }
    v3 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}

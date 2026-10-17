package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int32; F2 uint32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }
func genFn0(p0 uint32, p1 int8, p2 int16) bool {
    var v7 uint64 = 0
    var v8 float64 = ((float64(1.5) / ((-2.25 + 1.5) + (1e10 / 1.0))) * 1.5)
    var v9 uint8 = 255
    var v10 uint32 = p0
    if (!false) {
        v11 := []int{127, -9223372036854775808, 238, 659}
        switch 32767 {
        case -1:
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 float64 = (float64(0.1) * 3.14159)
                v14 := GenS1{F0: -72}
                for fi15 := 0; fi15 < 8; fi15++ {
                    v16 := GenS1{F0: (p1 + (p1 >> ((42) & 63)))}
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        case 9223372036854775807:
            fmt.Println(p1)
        default:
            switch -2147483648 {
            case 0:
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 int64 = 9223372036854775807
                    _ = fi17
                    _ = v18
                }
            case -9223372036854775808:
                for fi19 := 0; fi19 < 16; fi19++ {
                    v11 = append(v11, 128)
                    var v20 float64 = v8
                    v10 = 2147483647
                    _ = fi19
                    _ = v20
                }
            }
        }
        _ = v11
    } else {
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v8)
            var v22 string = "\n\t"
            switch fi21 {
            case -452:
                switch v22 {
                case "hello":
                    v7 += v7
                case "":
                    fmt.Println(fi21, v8)
                case "foo bar":
                    v23 := GenS2{F0: "a", F1: 641, F2: v10}
                    _ = v23
                default:
                    switch (-(fi21)) {
                    case -117:
                        var v24 int64 = int64(v7)
                        _ = v24
                    case -850:
                        var v25 int32 = -605
                        _ = v25
                    }
                }
            case -1:
                v26 := GenS3{F0: "foo bar", F1: 0}
                _ = v26
            case 405:
                v27 := []int{}
                _ = v27
            default:
                var v28 bool = true
                _ = v28
            }
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int, p1 int8) string {
    for fi29 := 0; fi29 < 3; fi29++ {
        p1 -= p1
        var v30 int = fi29
        if false {
            v30 = (fi29 << ((65535) & 63))
            p0 = v30
            v31 := []int{p0, 2147483647}
            _ = v31
        }
        _ = fi29
        _ = v30
    }
    var v32 int = p0
    _ = p0
    _ = p1
    _ = v32
    return ("" + "foo bar")
}

func main() {
    v1 := []uint32{826, 0, 4294967295}
    var v2 uint16 = 42
    v3 := []string{"", "foo bar", ("" + "hello"), (("\n\t" + "a") + "")}
    var v4 int8 = 3
    var v5 int = 63
    var v6 uint16 = 2
    var v33 string = "foo bar"
    var v34 float64 = 3.14159
    v1 = append(v1, 4294967295)
    v6 = 32767
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            v3 = append(v3, v33)
            v37 := GenS3{F0: "a", F1: 906}
            var v38 int = v5
            _ = fi36
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 3; fi39++ {
            var v40 int16 = -32768
            switch 32768 {
            case 0:
                if (true && (false && (-768 >= v40))) {
                    v41 := GenS1{F0: v4}
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v6 -= 42
                        var v42 uint64 = 714
                        fmt.Println(v5, fi35, v2)
                        var v43 float32 = 1.5
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(v6, v6, v34)
                    }
                    _ = v41
                }
            case 4294967295:
                fmt.Println(v6, v4)
            default:
                v44 := GenS4{F0: 823, F1: 607, F2: (v2 + (v6 | (v2 / ((v2) | 1))))}
                _ = v44
            }
            _ = fi39
            _ = v40
        }
        _ = fi35
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v34 = ((v34 + v34) - -2.25)
    switch v5 {
    case -681:
        switch (v33 + v33) {
        case "a":
            var v45 uint16 = (v6 >> ((208) & 63))
            _ = v45
        case "hello":
            v46 := GenS3{F0: ("" + "a"), F1: 0}
            _ = v46
        case "":
            var v47 string = ("a" + ("" + ("" + ("" + v33))))
            _ = v47
        default:
            switch "a" {
            case "foo bar":
                v48 := GenS2{F0: ((v33 + ((v33 + v33) + ((v33 + v33) + v33))) + v33), F1: int32(v5), F2: 545}
                _ = v48
            case "\n\t":
                v33 = v33
            }
        }
    case 127:
        var v49 uint16 = (v6 << ((0) & 63))
        _ = v49
    case 1:
        v50 := GenS3{F0: ("a" + (("" + (v33 + v33)) + "a")), F1: 255}
        _ = v50
    default:
        if (false && false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) > 0 {
                v1[0] = 2
            }
        } else {
            switch v33 {
            case "hello":
                var v51 int16 = 1000
                _ = v51
            case "a":
                fmt.Println(v5)
            case "":
                if (0 < 32767) {
                    switch v33 {
                    case "":
                        fmt.Println(v5)
                    default:
                        fmt.Println(v6, v2, v4)
                    }
                    if ("\n\t" != (v33 + v33)) {
                        var v52 uint = 65535
                        v52 = v52
                        v1 = append(v1, 2)
                        _ = v52
                    } else {
                        v53 := GenS3{F0: v33, F1: 255}
                        fmt.Println(v6)
                        if (!(65535 >= 2)) {
                            _ = len(v1)
                            fmt.Println(v34, v4)
                        } else {
                            if true {
                                fmt.Println(v5, v34, v6)
                                v34 = 0.0
                                switch 188 {
                                case -331:
                                    v33 = v33
                                case -173:
                                    switch v33 {
                                    case "a":
                                        v54 := GenS5{F0: v5}
                                        _ = v54
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                            }
                            v6 = 255
                            if true {
                                v2 -= (v6 & v2)
                                var v55 int = (-(v5))
                                _ = v55
                            } else {
                                v56 := GenS6{F0: 42}
                                v2 += v6
                                _ = v56
                            }
                        }
                        _ = v53
                    }
                }
            }
        }
    }
    fmt.Println(v4, v2)
    v6 = (v2 + 32767)
    v6 += (v6 / (((v2 ^ v6)) | 1))
    var v57 int8 = v4
    var v58 uint = 2
    fmt.Println(v4)
    v6 -= v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}

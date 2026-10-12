package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS6 struct { F0 string }
func genFn0(p0 uint, p1 uint16) uint64 {
    var v4 uint8 = 255
    v5 := GenS1{F0: 943}
    var v6 string = ""
    var v7 uint = 369
    var v8 string = v6
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint32, p1 uint8) int {
    var v9 int16 = -817
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            if ((-202 <= -1) && (!(229 >= 121))) {
                v11 := []int{}
                fmt.Println(p1, p0)
                _ = v11
            } else {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v9 = -471
                    fmt.Println(fi12)
                    _ = fi12
                }
                v13 := GenS1{F0: fi10}
                v14 := GenS1{F0: 1}
                _ = v13
                _ = v14
            }
            _ = fi10
        }
        if true {
            switch 127 {
            case 128:
                if false {
                    if (-1 > 0) {
                        v15 := []int{32767, 4294967295}
                        v16 := GenS2{F0: 128, F1: 941, F2: v9}
                        _ = v15
                        _ = v16
                    } else {
                        p0 += p0
                    }
                    p1 -= (p1 - p1)
                    v17 := []int{}
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        fmt.Println(fi18, p0, p0)
                        _ = fi18
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        p0 += (p0 | p0)
                        v20 := GenS3{F0: 1}
                        v21 := []int{}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            case -840:
                var v22 int8 = 42
                _ = v22
            case -1:
                if ((42 <= 1) || (1 != -655)) {
                    v23 := GenS4{F0: ""}
                    var v24 bool = true
                    _ = v23
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        p0 += (p0 >> ((255) & 63))
                        if true {
                            var v26 int32 = -1
                            if false {
                                v27 := []int{fi25}
                                switch "hello" {
                                case "\n\t":
                                    v26 += 545
                                case "hello":
                                    p1 = (^(p1))
                                case "":
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        if ((-2.25 - (float64(3.14159) + 1.5)) >= 0.1) {
                                            v9 = v9
                                            v29 := GenS1{F0: fi25}
                                            _ = v29
                                        }
                                        fmt.Println(p0)
                                        _ = fi28
                                    }
                                default:
                                    v27 = append(v27, (fi25 / ((fi25) | 1)))
                                }
                                v30 := GenS2{F0: (p1 << ((uint(p1)) & 63)), F1: 2, F2: (v9 * (v9 | (v9 ^ int16(fi25))))}
                                _ = v27
                                _ = v30
                            } else {
                                var v31 uint16 = 255
                                var v32 uint64 = 1
                                var v33 uint32 = (p0 & p0)
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            v9 = int16(p1)
                            _ = v26
                        } else {
                            p0 = p0
                            for fi34 := 0; fi34 < 19; fi34++ {
                                v35 := []int{fi34, fi34, (fi25 + (fi34 >> ((42) & 63)))}
                                _ = fi34
                                _ = v35
                            }
                        }
                        var v36 bool = (!true)
                        _ = fi25
                        _ = v36
                    }
                    p0 += p0
                    switch 577 {
                    case 0:
                        var v37 int8 = 1
                        _ = v37
                    default:
                        fmt.Println(p1, p1)
                    }
                }
            }
            v9 = (v9 << ((1) & 63))
            for fi38 := 0; fi38 < 5; fi38++ {
                v9 -= v9
                _ = fi38
            }
        } else {
            v39 := []int{128, 1}
            _ = v39
        }
        switch ("" + "\n\t") {
        case "\n\t":
            fmt.Println(v9, p0)
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 127
}

func main() {
    var v1 int16 = 578
    v2 := []uint32{2, 308}
    v3 := []string{"a"}
    v1 += 42
    var v40 int16 = v1
    v40 -= (^(v1))
    for fi41 := 0; fi41 < 16; fi41++ {
        var v42 int8 = 100
        var v43 uint64 = 0
        _ = fi41
        _ = v42
        _ = v43
    }
    v40 += (v1 / ((v1) | 1))
    var v44 int8 = 42
    _ = len(v2)
    fmt.Println(v1)
    v44 = -70
    v40 = v1
    fmt.Println(v1)
    for fi45 := 0; fi45 < 11; fi45++ {
        v46 := GenS3{F0: -9223372036854775808}
        _ = fi45
        _ = v46
    }
    var v47 int64 = 1
    if false {
        for fi48 := 0; fi48 < 10; fi48++ {
            if (false && (0.1 > 0.0)) {
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "hello")
                }
            } else {
                v44 -= (-(v44))
                var v49 float64 = ((-2.25 + 1e10) * 3.14159)
                v50 := GenS5{F0: 2, F1: 188, F2: 4294967295}
                _ = v49
                _ = v50
            }
            v51 := GenS1{F0: genFn1(4294967295, 128)}
            _ = fi48
            _ = v51
        }
        for fi52 := 0; fi52 < 19; fi52++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi52
        }
    } else {
        fmt.Println(v1)
        switch ("\n\t" + ("a" + ("\n\t" + "foo bar"))) {
        case "":
            if true {
                v1 = (v1 % ((v1) | 1))
                for fi53 := 0; fi53 < 11; fi53++ {
                    for fi54 := 0; fi54 < 17; fi54++ {
                        if ((1.5 - 3.14159) > 0.0) {
                            var v55 float64 = 3.14159
                            v47 -= v47
                            _ = v55
                        }
                        _ = fi54
                    }
                    v40 = v40
                    v1 -= -32768
                    _ = fi53
                }
                var v56 float32 = 0.1
                _ = v56
            } else {
                v1 -= (v1 ^ (v40 >> ((18446744073709551615) & 63)))
                v1 = v40
            }
        case "a":
            fmt.Println(v40, v40)
        case "hello":
            if (false && ((!(v47 == (v47 + (^(v47))))) || ((!false) || true))) {
                var v57 uint = 18446744073709551615
                if true {
                    switch ("\n\t" + (("hello" + ("a" + ("hello" + "\n\t"))) + "hello")) {
                    case "\n\t":
                        v57 += (^(v57))
                    }
                } else {
                    if true {
                        v40 += (-(v40))
                        for fi58 := 0; fi58 < 11; fi58++ {
                            for fi59 := 0; fi59 < 7; fi59++ {
                                v60 := GenS6{F0: ""}
                                if (v47 > v47) {
                                    if ((fi59 - fi59) >= -682) {
                                        fmt.Println(fi58, fi59, v47)
                                    }
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        _ = len(v2)
                                    case -891:
                                        v61 := GenS3{F0: v47}
                                        _ = v61
                                    }
                                    fmt.Println(v47, fi59)
                                } else {
                                    v57 += 4294967295
                                    var v62 float64 = ((-2.25 * 1e10) - 3.14159)
                                    fmt.Println(v1, v62)
                                    _ = v62
                                }
                                _ = fi59
                                _ = v60
                            }
                            v40 = (^(v40))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi58
                        }
                    }
                    var v63 int32 = 2147483647
                    _ = v63
                }
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                _ = v57
            }
        default:
            fmt.Println(v40, v44)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v44
    _ = v47
}

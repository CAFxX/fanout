package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 uint8; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 int8; F1 int; F2 uint8 }
func genFn0(p0 uint32) int8 {
    v6 := GenS1{F0: 255, F1: 2, F2: (-32768 <= 32767)}
    if (-1 >= 50) {
        v7 := []int{761, -1}
        p0 = p0
        _ = v7
    }
    v8 := []int{0}
    _ = p0
    _ = v6
    _ = v8
    return -115
}

func main() {
    v1 := []string{"a", "\n\t", "", "\n\t"}
    var v2 string = ("foo bar" + "\n\t")
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    var v5 bool = true
    if true {
        var v9 string = v2
        var v10 uint = 18446744073709551615
        _ = v10
        _ = v9
    } else {
        var v11 int = -9223372036854775808
        if (!true) {
            v12 := v3[1]
            _ = v12
        } else {
            var v13 int64 = 42
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint32 = 710
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        if v5 {
            v16 := GenS2{F0: 255, F1: 2}
            if v5 {
                if (127 > 128) {
                    v5 = v5
                    v17 := GenS3{F0: 1, F1: 42, F2: false}
                    _ = v17
                } else {
                    var v18 string = "foo bar"
                    v2 = "\n\t"
                    _ = v18
                }
                fmt.Println(v5, v2)
            } else {
                v11 = 9223372036854775807
                var v19 int8 = 0
                for fi20 := 0; fi20 < 10; fi20++ {
                    v11 = 578
                    _ = fi20
                }
                _ = v19
            }
            _ = v16
        }
        _ = v11
    }
    if (-1 > -32768) {
        v5 = (42 > 18446744073709551615)
        var v21 uint32 = 4294967295
        _ = v21
    }
    v22 := GenS2{F0: 2, F1: 2}
    for fi23 := 0; fi23 < 2; fi23++ {
        fmt.Println(v5)
        _ = fi23
    }
    if true {
        switch v2 {
        case "a":
            fmt.Println(v5, v5)
        default:
            for fi24 := 0; fi24 < 5; fi24++ {
                v2 = ""
                _ = fi24
            }
        }
    } else {
        v25 := GenS4{F0: 395, F1: 93, F2: -9223372036854775808}
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        if v5 {
            if (!true) {
                v27 := GenS2{F0: 255, F1: 134}
                v2 = "hello"
                v28 := v3[4294967295]
                _ = v27
                _ = v28
            } else {
                v2 = ""
            }
            if (v5 && v5) {
                for fi29 := 0; fi29 < 17; fi29++ {
                    var v30 int32 = -2147483648
                    switch 128 {
                    case -972:
                        v31 := v3[65535]
                        var v32 uint64 = 2
                        fmt.Println(v5, v31, v30)
                        var v33 uint8 = 92
                        v2 = (("foo bar" + v2) + v2)
                        _ = v31
                        _ = v32
                        _ = v33
                    default:
                        if (255 <= 2) {
                            v34 := GenS5{F0: fi26, F1: 128}
                            _ = v34
                        } else {
                            var v35 int64 = 1
                            var v36 string = v2
                            v37 := GenS6{F0: 1}
                            v36 = "\n\t"
                            v2 = "foo bar"
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                    v5 = (3.14159 != ((float32((((float64(3.14159) * float64(((1e10 - ((0.1 * 0.0) / 0.1)) * 1.5))) - 1.5) / 3.14159)) - 0.1) + 1.5))
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v2, v5, fi26)
            } else {
                var v38 int = (-(fi26))
                v5 = (42 <= 255)
                v5 = v5
                _ = v38
            }
        } else {
            for fi39 := 0; fi39 < 10; fi39++ {
                if (0 >= 647) {
                    if v5 {
                        v40 := GenS6{F0: 4294967295}
                        switch 128 {
                        case -9223372036854775808:
                            v41 := v3[280]
                            if true {
                                v42 := v4["\n\t"]
                                v2 = ("a" + ("a" + ("" + (v2 + "a"))))
                                var v43 uint16 = 925
                                _ = v42
                                _ = v43
                            }
                            fmt.Println(v2, v41, fi26)
                            _ = v41
                        default:
                            v5 = v5
                        }
                        v44 := GenS7{F0: (fi39 >> ((65535) & 63)), F1: 255, F2: (-(fi26))}
                        _ = v40
                        _ = v44
                    }
                    _ = len(v4)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v45 := GenS8{F0: 0}
                    v5 = true
                    fmt.Println(v5, v2)
                    _ = v45
                }
                var v46 int = (fi26 >> ((4294967295) & 63))
                v4[v2] = 797
                v46 -= (fi26 - v46)
                _ = fi39
                _ = v46
            }
            fmt.Println(fi26, v2)
            v47 := GenS7{F0: (fi26 / (((-(fi26))) | 1)), F1: 2, F2: fi26}
            _ = v47
        }
        _ = fi26
    }
    if (!(1000 == -32768)) {
        v5 = false
    }
    v48 := GenS9{F0: 79, F1: -901}
    for fi49 := 0; fi49 < 5; fi49++ {
        for fi50 := 0; fi50 < 17; fi50++ {
            v5 = true
            if v5 {
                v51 := v3[2147483648]
                v2 = v2
                if ((fi49 / (((fi50 / (((fi50 + fi50)) | 1))) | 1)) <= fi50) {
                    if false {
                        v5 = v5
                        _ = len(v4)
                        switch 9223372036854775807 {
                        case 32:
                            var v52 int = (v51 | (v51 + (^(v51))))
                            _ = v52
                        case -647:
                            v3[0] = (-(fi50))
                        case -248:
                            var v53 uint16 = 516
                            _ = v53
                        }
                    }
                    var v54 uint32 = 2147483648
                    _ = v54
                } else {
                    v5 = (42 <= 0)
                    fmt.Println(fi50, fi50)
                    v55 := GenS10{F0: -1, F1: fi50, F2: 128}
                    _ = v55
                }
                _ = v51
            }
            _ = fi50
        }
        _ = fi49
    }
    switch "" {
    case "\n\t":
        v5 = (115 >= 0)
    default:
        switch v2 {
        case "":
            if v5 {
                for fi56 := 0; fi56 < 14; fi56++ {
                    v2 = v2
                    _ = fi56
                }
            }
        }
    }
    var v57 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v57
}

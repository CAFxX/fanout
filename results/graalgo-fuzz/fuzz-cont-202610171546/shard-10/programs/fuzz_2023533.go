package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
func genFn0(p0 uint) int {
    fmt.Println(p0)
    if false {
        v6 := []int{32767, 2147483647, 730}
        if len(v6) > 0 {
            v6[0] = 942
        }
        _ = v6
    }
    var v7 int16 = -1
    v7 += 0
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1(p0 uint8) uint8 {
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: fi8, F1: 42, F2: p0}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        p0 -= 2
        switch "foo bar" {
        case "hello":
            v11 := GenS1{F0: -321, F1: p0, F2: (^(p0))}
            _ = v11
        }
        _ = fi10
    }
    if true {
        v12 := []int{}
        var v13 int32 = 1
        _ = v12
        _ = v13
    }
    if (!(!((p0 & p0) == p0))) {
        v14 := GenS1{F0: 32767, F1: (p0 % ((91) | 1)), F2: p0}
        var v15 float32 = 1e10
        _ = v14
        _ = v15
    } else {
        fmt.Println(p0)
        v16 := GenS1{F0: -1, F1: p0, F2: p0}
        _ = v16
    }
    _ = p0
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (false && (0.0 > (0.1 + 0.1)))
    var v3 uint = 2
    var v4 uint8 = 2
    v5 := []int{}
    for fi17 := 0; fi17 < 5; fi17++ {
        if false {
            for fi18 := 0; fi18 < 10; fi18++ {
                v4 = 79
                fmt.Println(fi17)
                var v19 int = genFn0(uint(v4))
                _ = fi18
                _ = v19
            }
        } else {
            fmt.Println(v2, v4, fi17)
        }
        if v2 {
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint64 = 0
                v4 = genFn1((v4 - v4))
                _ = fi20
                _ = v21
            }
            fmt.Println(v4)
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int = (fi17 << ((v3) & 63))
                _ = fi22
                _ = v23
            }
        }
        v24 := GenS1{F0: 2147483647, F1: (v4 ^ (v4 + v4)), F2: genFn1(v4)}
        _ = fi17
        _ = v24
    }
    var v25 uint = 42
    for fi26 := 0; fi26 < 14; fi26++ {
        v2 = ((^(v4)) < (^(v4)))
        var v27 int32 = 1
        if v2 {
            v2 = (v2 || (42 >= 1))
            delete(v1, "\n\t")
            for fi28 := 0; fi28 < 17; fi28++ {
                v4 -= (v4 + (^(v4)))
                _ = len(v5)
                _ = fi28
            }
        } else {
            if (0 == 2147483647) {
                v25 = 2
            }
            if (!false) {
                var v29 uint = (v25 + v25)
                var v30 int16 = 32767
                for fi31 := 0; fi31 < 3; fi31++ {
                    var v32 uint64 = 18446744073709551615
                    var v33 int64 = -9223372036854775808
                    fmt.Println(fi31, v32)
                    fmt.Println(fi26)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v29
                _ = v30
            } else {
                for fi34 := 0; fi34 < 18; fi34++ {
                    v27 = -998
                    _ = fi34
                }
            }
            switch "hello" {
            case "\n\t":
                for fi35 := 0; fi35 < 10; fi35++ {
                    switch ("\n\t" + "\n\t") {
                    case "":
                        var v36 int16 = -85
                        _ = v36
                    case "hello":
                        if (3.14159 > 3.14159) {
                            v3 -= v25
                            if ((^(v3)) == (v3 ^ 255)) {
                                if false {
                                    v2 = (!v2)
                                    if (!v2) {
                                        var v37 bool = (560 != 65535)
                                        _ = v37
                                    }
                                }
                                var v38 int32 = 42
                                var v39 float32 = -2.25
                                if ((v4 % ((0) | 1)) <= (v4 >> (((v25 << ((255) & 63))) & 63))) {
                                    v2 = (4294967295 >= (v3 | 0))
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v38
                                _ = v39
                            } else {
                                if false {
                                    var v40 int = 1
                                    var v41 uint64 = 2
                                    _ = v40
                                    _ = v41
                                }
                            }
                            v42 := GenS1{F0: (fi35 + (fi26 & (fi35 & (-(fi35))))), F1: 20, F2: genFn1((v4 / (((v4 + v4)) | 1)))}
                            _ = v42
                        }
                    case "a":
                        v43 := GenS1{F0: genFn0((^(v3))), F1: 253, F2: v4}
                        _ = v43
                    default:
                        v44 := GenS1{F0: genFn0(2), F1: 255, F2: (v4 * v4)}
                        _ = v44
                    }
                    v3 = v3
                    _ = fi35
                }
            }
        }
        _ = fi26
        _ = v27
    }
    v3 += v3
    fmt.Println(v2, v4)
    var v45 int8 = -110
    switch int(v45) {
    case 32768:
        if len(v5) > 0 {
            v5[0] = 601
        }
    case 740:
        if (!true) {
            for fi46 := 0; fi46 < 2; fi46++ {
                fmt.Println(v4)
                v2 = (!true)
                v3 = (v3 / ((625) | 1))
                _ = fi46
            }
            switch "" {
            case "a":
                fmt.Println(v4, v3, v45)
            case "":
                v47 := GenS2{F0: -751, F1: -763}
                _ = v47
            }
        } else {
            v48 := GenS1{F0: 2147483647, F1: (v4 * (v4 >> (((v25 | v3)) & 63))), F2: v4}
            fmt.Println(v25, v3)
            _ = v48
        }
    default:
        if (!false) {
            v4 = v4
        }
    }
    v49 := GenS3{F0: -2147483648, F1: v4, F2: v45}
    if len(v5) > 0 {
        v5[0] = -929
    }
    var v50 bool = v2
    var v51 uint8 = (v4 % ((v4) | 1))
    v2 = (255 > 255)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}

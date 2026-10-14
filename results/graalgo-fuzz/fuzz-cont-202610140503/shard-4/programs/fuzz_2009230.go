package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint8, p1 string, p2 uint) uint32 {
    v4 := GenS1{F0: 255}
    var v5 int32 = -1
    p2 = p2
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint32 = 592
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS2{F0: fi8}
            var v11 int = fi8
            v12 := GenS1{F0: 0}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    v13 := []int{}
    _ = v13
    return 2
}

func main() {
    var v1 string = "\n\t"
    var v2 uint = 113
    var v3 uint16 = 32767
    var v14 bool = false
    v15 := GenS2{F0: 32767}
    fmt.Println(v2, v14, v14)
    if (true || v14) {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint8 = genFn1()
            _ = fi16
            _ = v17
        }
        var v18 int64 = 665
        _ = v18
    } else {
        var v19 int32 = 0
        v3 = (v3 / ((v3) | 1))
        v2 += uint(v3)
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        if v14 {
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS3{F0: 42}
                v23 := GenS2{F0: fi20}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 uint32 = 65535
            _ = v24
        }
        fmt.Println(fi20, v14)
        _ = fi20
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 bool = true
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 65535}
    v14 = false
    if v14 {
        v28 := []int{42, 32767, 501, 32768}
        if v14 {
            v29 := GenS3{F0: 420}
            v1 = (("\n\t" + "hello") + "a")
            _ = v29
        }
        switch 42 {
        case 781:
            if v14 {
                v30 := GenS1{F0: 2}
                _ = v30
            } else {
                fmt.Println(v2, v2, v2)
                for fi31 := 0; fi31 < 9; fi31++ {
                    v32 := GenS4{F0: "hello", F1: (false && v14)}
                    fmt.Println(v14)
                    fmt.Println(v1, fi31)
                    _ = fi31
                    _ = v32
                }
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            }
        }
        _ = v28
    }
    var v33 uint8 = 1
    switch (("foo bar" + ((v1 + v1) + "\n\t")) + (v1 + v1)) {
    case "\n\t":
        v34 := GenS5{F0: (v2 & v2)}
        _ = v34
    case "foo bar":
        v35 := []int{4294967295, 1, int(v2)}
        _ = v35
    case "":
        if v14 {
            for fi36 := 0; fi36 < 12; fi36++ {
                if v14 {
                    v37 := GenS2{F0: -9223372036854775808}
                    _ = v37
                } else {
                    var v38 string = ((v1 + v1) + v1)
                    _ = v38
                }
                switch 0 {
                case 4294967295:
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v14 = (v1 == "\n\t")
                        if (true || v14) {
                            fmt.Println(fi36, v33, fi36)
                            if v14 {
                                if v14 {
                                    v33 = (v33 * uint8(v3))
                                    switch (^(fi39)) {
                                    case -784:
                                        v40 := []int{fi39, fi39}
                                        _ = v40
                                    case -1:
                                        v1 = (v1 + (v1 + ""))
                                    case 0:
                                        v3 += v3
                                    }
                                }
                            }
                            if (!true) {
                                v41 := []int{}
                                if len(v41) > 0 {
                                    v41[0] = (fi36 / ((fi36) | 1))
                                }
                                _ = v41
                            }
                        } else {
                            v42 := []int{}
                            _ = v42
                        }
                        switch fi36 {
                        case -9223372036854775808:
                            var v43 string = (v1 + "hello")
                            _ = v43
                        default:
                            switch (fi36 | fi36) {
                            case 2147483647:
                                fmt.Println(v14, v3, v3)
                            default:
                                v44 := GenS6{F0: fi39, F1: 127}
                                _ = v44
                            }
                        }
                        _ = fi39
                    }
                case -2147483648:
                    for fi45 := 0; fi45 < 19; fi45++ {
                        var v46 int32 = 2147483647
                        if true {
                            v47 := GenS7{F0: (-(v46))}
                            var v48 uint = (v2 << (((v2 & v2)) & 63))
                            var v49 float32 = 0.1
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        _ = fi45
                        _ = v46
                    }
                case 9223372036854775807:
                    var v50 bool = (fi36 != fi36)
                    _ = v50
                }
                var v51 bool = ((^(v3)) == v3)
                _ = fi36
                _ = v51
            }
            var v52 float32 = 1.5
            if (!false) {
                v3 -= v3
            } else {
                fmt.Println(v2, v33)
                for fi53 := 0; fi53 < 10; fi53++ {
                    fmt.Println(v1, v2, v1)
                    if v14 {
                        if true {
                            var v54 int8 = -128
                            _ = v54
                        } else {
                            var v55 int = (fi53 * 0)
                            var v56 int = 42
                            switch v56 {
                            case 795:
                                v57 := GenS5{F0: (v2 / ((18446744073709551615) | 1))}
                                _ = v57
                            case 42:
                                var v58 uint8 = 1
                                _ = v58
                            case 1:
                                v1 = v1
                            default:
                                var v59 bool = v14
                                _ = v59
                            }
                            _ = v55
                            _ = v56
                        }
                        for fi60 := 0; fi60 < 9; fi60++ {
                            var v61 int64 = 0
                            v62 := GenS8{F0: (v33 / ((genFn1()) | 1)), F1: -31}
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                    }
                    var v63 uint = (v2 + (v2 | 65535))
                    _ = fi53
                    _ = v63
                }
                fmt.Println(v1)
            }
            _ = v52
        } else {
            var v64 uint32 = genFn0(v33, v1, v2)
            _ = v64
        }
    }
    var v65 uint64 = 18446744073709551615
    for fi66 := 0; fi66 < 3; fi66++ {
        fmt.Println(v14, v2, v1)
        _ = fi66
    }
    fmt.Println(v2, v65)
    if ((v33 != v33) && true) {
        v67 := []int{4294967295}
        v2 = 42
        _ = v67
    }
    if ((^(v2)) < (^(v2))) {
        fmt.Println(v2, v2)
        if v14 {
            if (-680 < 1) {
                fmt.Println(v1, v14)
            }
            var v68 int = 4294967295
            _ = v68
        } else {
            v69 := []int{127}
            _ = v69
        }
    }
    if (1.5 == 1e10) {
        fmt.Println(v65)
    } else {
        switch v1 {
        case "hello":
            var v70 int16 = -1
            _ = v70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v65
}

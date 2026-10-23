package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint8, p1 uint16, p2 int) int {
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            p2 = fi6
            _ = fi6
        }
        v7 := []int{}
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: (p1 | p1), F1: 42, F2: 1}
        var v10 bool = true
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: -1, F1: -190}
    v12 := GenS2{F0: 2147483647, F1: 200}
    switch ("hello" + "hello") {
    case "foo bar":
        p2 = p2
    case "hello":
        v13 := []int{(-(p2)), (p2 % (((p2 - p2)) | 1)), 32768}
        _ = v13
    case "\n\t":
        if false {
            switch "\n\t" {
            case "foo bar":
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch fi14 {
                    case 1:
                        p0 -= uint8(p1)
                    case 9223372036854775807:
                        p1 += 893
                    }
                    p0 -= (p0 >> ((42) & 63))
                    _ = fi14
                }
            }
        } else {
            if true {
                p1 += 1
                if (p1 == (^(p1))) {
                    p2 = -878
                    if (true && (!false)) {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            if (!false) {
                                p1 = (p1 + 828)
                            } else {
                                p1 = (p1 / ((652) | 1))
                            }
                            _ = fi15
                        }
                    }
                    p2 = (p2 ^ p2)
                }
            } else {
                var v16 int32 = -1
                _ = v16
            }
            p1 = (p1 << ((383) & 63))
            switch (("foo bar" + ("" + "\n\t")) + ("hello" + "foo bar")) {
            case "\n\t":
                if false {
                    if false {
                        fmt.Println(p0)
                        fmt.Println(p2)
                        p0 -= (p0 + p0)
                    }
                }
            case "foo bar":
                v17 := GenS3{F0: 1}
                _ = v17
            default:
                fmt.Println(p2, p0)
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p2
}
func genFn1(p0 uint) int8 {
    p0 += 18446744073709551615
    p0 += p0
    p0 += (^(p0))
    _ = p0
    return 127
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    var v3 int64 = -9223372036854775808
    var v4 int = -2147483648
    var v18 int = v4
    _ = len(v2)
    var v19 uint32 = 2
    var v20 int64 = (v3 + -838)
    if (v1 && (!(false || v1))) {
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                switch v4 {
                case -435:
                    fmt.Println(v20, v4)
                }
                _ = fi22
            }
            if (v1 || v1) {
                var v23 uint16 = 32767
                _ = v23
            } else {
                v1 = (1000 < -690)
                v19 = 1
            }
            _ = fi21
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := GenS4{F0: 42, F1: 467, F2: (v20 & (v3 | -356))}
            switch "hello" {
            case "a":
                fmt.Println(v1)
            case "":
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := GenS2{F0: -2147483648, F1: 42}
                    var v28 float32 = 0.0
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            default:
                var v29 int = 42
                _ = v29
            }
            _ = fi24
            _ = v25
        }
        var v30 uint8 = 127
        _ = v30
    } else {
        v31 := []int{v18}
        _ = v31
    }
    v1 = (!(-9223372036854775808 >= v20))
    if v1 {
        var v32 int32 = 2147483647
        _ = v32
    } else {
        if true {
            var v33 uint64 = 595
            _ = v33
        } else {
            var v34 uint = 4294967295
            var v35 string = ((((("\n\t" + ("" + ("hello" + ""))) + "") + ((("\n\t" + "hello") + "") + "")) + "\n\t") + "")
            _ = v34
            _ = v35
        }
    }
    switch "foo bar" {
    case "foo bar":
        for fi36 := 0; fi36 < 4; fi36++ {
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 int32 = -1
                if false {
                    var v39 string = ("hello" + "a")
                    _ = v39
                }
                v1 = true
                var v40 float64 = -2.25
                if (!true) {
                    v19 = 2147483648
                    if v1 {
                        v38 = -1
                        var v41 int16 = 0
                        _ = v41
                    } else {
                        v42 := GenS5{F0: (v19 + 769), F1: 0, F2: 540}
                        _ = v42
                    }
                    v43 := GenS2{F0: v38, F1: -761}
                    _ = v43
                } else {
                    fmt.Println(v4, fi36, v1)
                }
                v44 := v2[4294967295]
                _ = fi37
                _ = v38
                _ = v40
                _ = v44
            }
            _ = fi36
        }
    case "a":
        if true {
            for fi45 := 0; fi45 < 17; fi45++ {
                if false {
                    v46 := GenS6{F0: 42, F1: (genFn1(255) == -128)}
                    v2[2] = (1 < 2)
                    _ = v46
                } else {
                    fmt.Println(v18)
                    for fi47 := 0; fi47 < 5; fi47++ {
                        if ((954 < 2) && (-123 <= 42)) {
                            var v48 uint = 883
                            v49 := GenS7{F0: 523, F1: -1}
                            v4 = fi45
                            _ = v48
                            _ = v49
                        }
                        var v50 int64 = -39
                        if true {
                            var v51 uint16 = 32767
                            v52 := []int{}
                            _ = v51
                            _ = v52
                        } else {
                            var v53 string = "hello"
                            var v54 bool = ((v53 >= ((("" + v53) + (((v53 + v53) + v53) + "")) + v53)) && ((v50 * v3) < -1))
                            v55 := v2[(v19 & v19)]
                            var v56 int64 = v20
                            var v57 int64 = (^(v50))
                            var v58 bool = (0 > 0)
                            v59 := GenS6{F0: 0, F1: v55}
                            _ = v53
                            _ = v54
                            _ = v55
                            _ = v56
                            _ = v57
                            _ = v58
                            _ = v59
                        }
                        _ = fi47
                        _ = v50
                    }
                }
                switch ("" + "\n\t") {
                case "foo bar":
                    var v60 float32 = -2.25
                    _ = v60
                }
                var v61 string = "a"
                _ = fi45
                _ = v61
            }
        }
    }
    v62 := GenS6{F0: -1, F1: v1}
    switch ("a" + "\n\t") {
    case "\n\t":
        for fi63 := 0; fi63 < 15; fi63++ {
            var v64 int = (v4 * (fi63 + (v4 << ((42) & 63))))
            _ = fi63
            _ = v64
        }
    case "a":
        if v1 {
            v20 = 716
            v3 += (v20 & v3)
        }
    case "":
        if v1 {
            for fi65 := 0; fi65 < 7; fi65++ {
                switch (v4 - -9223372036854775808) {
                case 128:
                    delete(v2, 249)
                }
                v1 = true
                _ = fi65
            }
            if false {
                v66 := GenS3{F0: 0}
                for fi67 := 0; fi67 < 7; fi67++ {
                    if v1 {
                        v4 += 127
                    } else {
                        var v68 uint = 1
                        var v69 int = genFn0(1, 1, (v18 & v18))
                        _ = v68
                        _ = v69
                    }
                    var v70 int32 = 592
                    var v71 float64 = float64((1e10 + (((1e10 / 1.5) + 1.5) * 1.5)))
                    _ = fi67
                    _ = v70
                    _ = v71
                }
                var v72 bool = ((^(v20)) < 993)
                _ = v66
                _ = v72
            }
        }
    }
    fmt.Println(v4)
    v73 := GenS1{F0: 1, F1: 66, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v62
    _ = v73
}

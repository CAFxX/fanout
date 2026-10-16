package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 string }
func genFn0(p0 int32, p1 bool, p2 uint) int16 {
    p2 = (^(p2))
    var v6 uint = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 621
    var v3 string = ""
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    switch v2 {
    case 9223372036854775807:
        v2 += -9223372036854775808
    case 454:
        var v7 int32 = 2147483647
        _ = v7
    case 42:
        v8 := GenS1{F0: (v2 % (((^(v2))) | 1)), F1: 127, F2: 42}
        _ = v8
    }
    v9 := v1[(v2 / ((v2) | 1))]
    switch v3 {
    case "a":
        for fi10 := 0; fi10 < 15; fi10++ {
            if true {
                v2 += v9
                if (((0.0 < 3.14159) && (((-96 <= fi10) && false) || (0 <= 574))) && ((-2.25 + -2.25) <= (float64(0.1) / -2.25))) {
                    if true {
                        var v11 string = v3
                        switch (v9 << ((2) & 63)) {
                        case 32768:
                            var v12 string = v11
                            _ = v12
                        case 2147483647:
                            var v13 int8 = -128
                            _ = v13
                        default:
                            switch 523 {
                            case -9223372036854775808:
                                for fi14 := 0; fi14 < 7; fi14++ {
                                    fmt.Println(v2, v3)
                                    fmt.Println(fi14, fi10)
                                    v3 = (v11 + (("\n\t" + ("a" + v3)) + v3))
                                    _ = fi14
                                }
                            case 127:
                                switch v2 {
                                case -9223372036854775808:
                                    var v15 uint64 = 42
                                    _ = v15
                                case 32767:
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        var v17 int8 = int8(fi10)
                                        _ = fi16
                                        _ = v17
                                    }
                                default:
                                    v3 = "hello"
                                }
                            }
                        }
                        v1[(v9 & -302)] = v2
                        _ = v11
                    } else {
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v9 -= (v9 * 127)
                            var v19 int = (-(v2))
                            switch v3 {
                            case "a":
                                switch ((v3 + "\n\t") + v3) {
                                case "hello":
                                    switch v3 {
                                    case "\n\t":
                                        var v20 float32 = 1.5
                                        _ = v20
                                    }
                                }
                            case "":
                                v3 = "a"
                            case "\n\t":
                                delete(v5, "foo bar")
                                var v21 float32 = float32(-2.25)
                                _ = v21
                            }
                            switch v3 {
                            case "\n\t":
                                if false {
                                    v3 = v3
                                    delete(v4, v2)
                                } else {
                                    v9 -= (fi10 % (((v2 * 251)) | 1))
                                    switch v2 {
                                    case 39:
                                        if false {
                                            var v22 string = v3
                                            fmt.Println(fi18, fi18)
                                            var v23 uint = 255
                                            _ = v22
                                            _ = v23
                                        } else {
                                            v24 := GenS2{F0: -1}
                                            _ = v24
                                        }
                                        var v25 string = (v3 + "hello")
                                        _ = v25
                                    case -67:
                                        v9 = 1
                                    case 32767:
                                        fmt.Println(v9, v3)
                                    }
                                    v2 += 32767
                                    var v26 uint32 = 2147483647
                                    v27 := GenS3{F0: 0, F1: 1, F2: 0}
                                    v28 := GenS4{F0: 0}
                                    v4[v19] = (v3 + "")
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                            var v29 uint64 = 42
                            _ = fi18
                            _ = v19
                            _ = v29
                        }
                        var v30 uint = 65535
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v32 := GenS4{F0: 1}
                        _ = fi31
                        _ = v32
                    }
                } else {
                    switch (v3 + (v3 + (v3 + (v3 + v3)))) {
                    case "":
                        if true {
                            for fi33 := 0; fi33 < 15; fi33++ {
                                var v34 string = v3
                                _ = fi33
                                _ = v34
                            }
                            var v35 uint16 = 734
                            _ = v35
                        } else {
                            switch (v9 >> ((18446744073709551615) & 63)) {
                            case 2147483647:
                                v36 := GenS1{F0: v2, F1: 130, F2: 128}
                                _ = v36
                            case 0:
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    v38 := GenS2{F0: 42}
                                    fmt.Println(v2)
                                    fmt.Println(fi37, fi10)
                                    _ = fi37
                                    _ = v38
                                }
                            default:
                                if (-1 != 32767) {
                                    v39 := GenS1{F0: (fi10 & -9223372036854775808), F1: 0, F2: 0}
                                    _ = v39
                                } else {
                                    for fi40 := 0; fi40 < 3; fi40++ {
                                        v41 := GenS1{F0: v9, F1: 42, F2: 255}
                                        if (-2147483648 >= 453) {
                                            var v42 string = (v3 + (("foo bar" + ("foo bar" + "\n\t")) + (v3 + ((v3 + v3) + v3))))
                                            v2 -= (v2 - v9)
                                            _ = v42
                                        } else {
                                            if (!(false || false)) {
                                                var v43 float64 = 0.1
                                                v9 += (v2 - (-(fi10)))
                                                var v44 bool = false
                                                switch (fi10 % (((-(v9))) | 1)) {
                                                case 127:
                                                    v43 += (1.5 / v43)
                                                case -9223372036854775808:
                                                    fmt.Println(v2)
                                                case -1:
                                                    v45 := GenS1{F0: fi40, F1: 128, F2: 76}
                                                    _ = v45
                                                }
                                                v46 := GenS5{F0: 0, F1: (v9 << ((107) & 63))}
                                                v43 = 0.0
                                                switch (fi40 >> ((947) & 63)) {
                                                case 4294967295:
                                                    v47 := GenS2{F0: genFn0(778, (!v44), 859)}
                                                    _ = v47
                                                case 42:
                                                    v44 = v44
                                                case -191:
                                                    _ = len(v1)
                                                }
                                                _ = v43
                                                _ = v44
                                                _ = v46
                                            }
                                            v9 -= (v9 * v9)
                                            if (!(2 <= 18446744073709551615)) {
                                                if true {
                                                    var v48 int = 842
                                                    v2 = fi10
                                                    v2 = (fi40 + (fi10 ^ (v9 << ((1) & 63))))
                                                    _ = v48
                                                } else {
                                                    v3 = "hello"
                                                    v9 = (v9 | (v2 / ((128) | 1)))
                                                    fmt.Println(v9, v2)
                                                    v9 = (fi40 | (v9 / (((^(v2))) | 1)))
                                                }
                                            } else {
                                                var v49 int32 = 2147483647
                                                v2 -= fi10
                                                var v50 uint16 = 664
                                                v50 = 65535
                                                v49 = v49
                                                v51 := []int{}
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                        }
                                        delete(v1, (^(v2)))
                                        _ = fi40
                                        _ = v41
                                    }
                                }
                            }
                        }
                    }
                }
            } else {
                var v52 int8 = -1
                _ = v52
            }
            _ = fi10
        }
    }
    var v53 int32 = 2147483647
    v53 += 0
    v54 := GenS6{F0: "hello"}
    fmt.Println(v3, v53, v3)
    v55 := GenS3{F0: 595, F1: 2147483648, F2: (-(v53))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v9
}

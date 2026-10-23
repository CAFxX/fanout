package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 int16) uint32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int16 = 5
        var v8 uint16 = 2
        v9 := GenS1{F0: (v7 << ((65535) & 63))}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if false {
            var v11 string = ("a" + "hello")
            _ = v11
        }
        p0 -= (^(p0))
        var v12 uint16 = 42
        _ = fi10
        _ = v12
    }
    var v13 int8 = 100
    _ = p0
    _ = v13
    return 924
}
func genFn1(p0 string, p1 int, p2 string) uint16 {
    switch (p1 >> ((0) & 63)) {
    case -2147483648:
        switch 141 {
        case 128:
            if true {
                var v14 uint = 1
                v15 := GenS2{F0: 18446744073709551615}
                var v16 uint32 = 1
                _ = v14
                _ = v15
                _ = v16
            }
        case -2147483648:
            if false {
                v17 := GenS3{F0: 9223372036854775807, F1: 124}
                for fi18 := 0; fi18 < 19; fi18++ {
                    fmt.Println(fi18)
                    v19 := []int{(fi18 - p1), p1}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
        }
    }
    p2 = p0
    v20 := []int{p1}
    var v21 int64 = -9223372036854775808
    var v22 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    return (v22 - v22)
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 uint = 2
    switch v1 {
    case "a":
        v23 := []int{int(v5), -217, -925}
        _ = v23
    default:
        var v24 int = -853
        _ = v24
    }
    v25 := []int{1}
    fmt.Println(v1, v1)
    fmt.Println(v1)
    v5 -= 240
    v5 -= 255
    if false {
        v26 := GenS1{F0: 42}
        var v27 string = v1
        switch 9223372036854775807 {
        case -942:
            if true {
                v28 := GenS4{F0: (v1 + "a")}
                v2[0] = v1
                _ = v28
            } else {
                v27 = "hello"
                fmt.Println(v5, v1, v27)
            }
        case 9223372036854775807:
            if true {
                var v29 float32 = 1.5
                v30 := v3[v27]
                var v31 bool = false
                _ = v29
                _ = v30
                _ = v31
            } else {
                var v32 float32 = (float32(-2.25) / (0.0 - 0.1))
                v33 := GenS3{F0: 9223372036854775807, F1: 255}
                _ = v32
                _ = v33
            }
        default:
            if (634 >= 1) {
                switch -1 {
                case -767:
                    v1 = v1
                case 506:
                    for fi34 := 0; fi34 < 6; fi34++ {
                        v1 = "a"
                        switch "a" {
                        case "\n\t":
                            if false {
                                fmt.Println(v5)
                                if false {
                                    for fi35 := 0; fi35 < 11; fi35++ {
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        if true {
                                            v36 := GenS2{F0: 42}
                                            v37 := GenS2{F0: 2}
                                            var v38 float32 = (float32(0.1) * 1.5)
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        } else {
                                            switch fi34 {
                                            case 1:
                                                switch -9223372036854775808 {
                                                case 347:
                                                    v39 := v3[(((v27 + v27) + v27) + v1)]
                                                    _ = v39
                                                case -1:
                                                    fmt.Println(v1)
                                                case 32768:
                                                    switch -1 {
                                                    case -9223372036854775808:
                                                        v1 = v27
                                                    case 313:
                                                        var v40 float64 = 1.5
                                                        _ = v40
                                                    }
                                                }
                                            case -1:
                                                v5 += (v5 / (((v5 >> (((v5 | 726)) & 63))) | 1))
                                            case 32767:
                                                v27 = (v27 + v27)
                                                var v41 uint32 = 1
                                                _ = v41
                                            }
                                            var v42 uint32 = genFn0(450)
                                            v27 = "hello"
                                            _ = v42
                                        }
                                        _ = fi35
                                    }
                                } else {
                                    var v43 int16 = 593
                                    switch (v1 + (v1 + v27)) {
                                    case "":
                                        if true {
                                            v1 = v27
                                        }
                                    case "a":
                                        v44 := GenS2{F0: 830}
                                        _ = v44
                                    default:
                                        v45 := GenS3{F0: -1, F1: genFn1(v1, 4294967295, "a")}
                                        _ = v45
                                    }
                                    _ = v43
                                }
                            }
                        default:
                            if true {
                                v46 := v3["hello"]
                                _ = v46
                            }
                        }
                        v27 = "a"
                        _ = fi34
                    }
                case 128:
                    v27 = ((v27 + v1) + v27)
                default:
                    if (("foo bar" + "a") < "hello") {
                        var v47 int = 719
                        _ = v47
                    }
                }
                if false {
                    for fi48 := 0; fi48 < 12; fi48++ {
                        for fi49 := 0; fi49 < 19; fi49++ {
                            v27 = v1
                            _ = fi49
                        }
                        _ = fi48
                    }
                    var v50 int8 = 100
                    if (244 == 42) {
                        for fi51 := 0; fi51 < 5; fi51++ {
                            var v52 uint16 = 0
                            v53 := v3[v1]
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    } else {
                        switch ("\n\t" + v1) {
                        case "foo bar":
                            v1 = ""
                        case "\n\t":
                            fmt.Println(v50, v27, v27)
                        }
                    }
                    _ = v50
                } else {
                    var v54 float64 = 3.14159
                    _ = v54
                }
            }
        }
        _ = v26
        _ = v27
    } else {
        for fi55 := 0; fi55 < 9; fi55++ {
            if (442 < 737) {
                v56 := GenS5{F0: (("foo bar" + v1) + v1), F1: (false || false)}
                if len(v25) > 0 {
                    v25[0] = 9223372036854775807
                }
                _ = v56
            }
            _ = fi55
        }
    }
    v1 = ("a" + (v1 + (v1 + v1)))
    for fi57 := 0; fi57 < 12; fi57++ {
        if (-24 >= 100) {
            v58 := v4[(("a" + v1) + ((v1 + ("\n\t" + "foo bar")) + ("\n\t" + (v1 + v1))))]
            var v59 int32 = -2147483648
            _ = v58
            _ = v59
        } else {
            if (!true) {
                v25 = append(v25, (fi57 * (fi57 * (-(fi57)))))
                if true {
                    if (42 > 100) {
                        var v60 uint = (^(v5))
                        var v61 int = fi57
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        _ = v60
                        _ = v61
                    } else {
                        _ = len(v4)
                    }
                    v62 := GenS6{F0: "hello", F1: v1, F2: -29}
                    v63 := GenS1{F0: 42}
                    _ = v62
                    _ = v63
                }
                var v64 bool = false
                _ = v64
            } else {
                v65 := GenS1{F0: 42}
                _ = v65
            }
        }
        switch fi57 {
        case 32767:
            fmt.Println(v5, v5)
        case 42:
            delete(v2, 237)
        case 127:
            v1 = v1
        default:
            fmt.Println(fi57, v5, v5)
        }
        _ = fi57
    }
    fmt.Println(v1)
    v66 := GenS3{F0: 947, F1: 867}
    fmt.Println(v5, v5)
    var v67 int = 742
    var v68 uint64 = 635
    v25 = append(v25, -2147483648)
    v5 = 530
    fmt.Println(v67, v68, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v68
}

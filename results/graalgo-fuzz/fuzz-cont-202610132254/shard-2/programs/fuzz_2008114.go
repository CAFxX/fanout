package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32 }
func genFn0(p0 int64) int8 {
    if false {
        for fi4 := 0; fi4 < 5; fi4++ {
            fmt.Println(fi4)
            _ = fi4
        }
    }
    v5 := []int{4294967295, 1}
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: -960, F1: 2, F2: 0}
        fmt.Println(p0, fi6)
        p0 += 0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 86
}

func main() {
    var v1 uint = 42
    var v2 int16 = -32768
    var v3 int32 = -2147483648
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS2{F0: 65535, F1: 2}
        switch fi8 {
        case 127:
            if true {
                var v10 int16 = (v2 | 32767)
                var v11 string = "foo bar"
                _ = v10
                _ = v11
            } else {
                var v12 uint32 = 697
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 int8 = -9
                    if false {
                        var v15 float64 = 1e10
                        _ = v15
                    } else {
                        if true {
                            v16 := []int{}
                            _ = v16
                        } else {
                            fmt.Println(v2, v12, v14)
                            if false {
                                var v17 int = (^(fi8))
                                if false {
                                    v18 := GenS1{F0: -910, F1: 52, F2: 1}
                                    v19 := []int{(-(fi8)), fi13, v17}
                                    _ = v18
                                    _ = v19
                                }
                                _ = v17
                            }
                        }
                    }
                    v20 := GenS1{F0: v3, F1: 2, F2: 280}
                    v21 := []int{fi13, int(v12), fi13}
                    _ = fi13
                    _ = v14
                    _ = v20
                    _ = v21
                }
                _ = v12
            }
        case 4294967295:
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := []int{4294967295, 127, fi8, 42}
                fmt.Println(fi22, v1)
                _ = fi22
                _ = v23
            }
        default:
            v24 := []int{(fi8 % (((fi8 + fi8)) | 1))}
            _ = v24
        }
        v1 = v1
        _ = fi8
        _ = v9
    }
    v1 -= (v1 & 916)
    v25 := []int{0, 32767, 1, -688}
    if (0 <= 626) {
        v1 += v1
        switch -9223372036854775808 {
        case 128:
            v26 := GenS1{F0: (^(v3)), F1: 2, F2: 1}
            _ = v26
        default:
            var v27 int = -2147483648
            _ = v27
        }
        v25 = append(v25, -1)
    } else {
        v28 := GenS2{F0: 886, F1: 42}
        if true {
            if ("hello" != "") {
                var v29 bool = (-86 >= genFn0(-9223372036854775808))
                var v30 int64 = 0
                _ = v29
                _ = v30
            }
            v3 = v3
        } else {
            v25 = append(v25, 0)
        }
        _ = v28
    }
    v2 += 42
    switch -1 {
    case 127:
        var v31 bool = false
        _ = v31
    case -1:
        if (("hello" + "\n\t") > ("hello" + ("" + ("\n\t" + "hello")))) {
            for fi32 := 0; fi32 < 16; fi32++ {
                v3 = v3
                v3 += (-(v3))
                _ = fi32
            }
            if (888 != 2) {
                v33 := GenS1{F0: -2147483648, F1: 255, F2: 42}
                var v34 int8 = 100
                var v35 float64 = ((float64(-2.25) / float64(0.1)) / 3.14159)
                _ = v33
                _ = v34
                _ = v35
            } else {
                if (!((696 <= uint64(v1)) && (42 < v3))) {
                    v36 := GenS2{F0: 718, F1: 5}
                    switch (("foo bar" + ("foo bar" + "a")) + "hello") {
                    case "hello":
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v38 := GenS1{F0: v3, F1: 0, F2: 0}
                                _ = fi37
                                _ = v38
                            }
                        case "":
                            for fi39 := 0; fi39 < 13; fi39++ {
                                fmt.Println(v2, v3)
                                _ = fi39
                            }
                        case "hello":
                            v2 = (v2 & (v2 - (v2 - (v2 ^ 42))))
                        default:
                            var v40 int32 = (^(v3))
                            _ = v40
                        }
                    case "":
                        var v41 float32 = 1.5
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 20; fi42++ {
                        var v43 uint16 = 0
                        if (!true) {
                            v44 := GenS3{F0: -19, F1: 121, F2: 42}
                            var v45 int16 = (v2 % (((v2 + 1)) | 1))
                            _ = v44
                            _ = v45
                        }
                        _ = fi42
                        _ = v43
                    }
                    _ = v36
                } else {
                    fmt.Println(v1)
                }
                v46 := GenS1{F0: v3, F1: 89, F2: 32767}
                _ = v46
            }
        } else {
            if (((!false) && (!true)) && (true || (100 >= genFn0(201)))) {
                v47 := GenS2{F0: 2147483647, F1: 46}
                v1 += (v1 + v1)
                var v48 bool = true
                _ = v47
                _ = v48
            } else {
                _ = len(v25)
            }
            var v49 uint = 65535
            _ = v49
        }
    case 526:
        if (true || false) {
            switch (("foo bar" + ("hello" + (("" + "hello") + "hello"))) + ("foo bar" + ("foo bar" + ("foo bar" + "hello")))) {
            case "foo bar":
                for fi50 := 0; fi50 < 9; fi50++ {
                    v51 := GenS4{F0: ("a" + ("\n\t" + "\n\t"))}
                    var v52 int8 = 77
                    v25 = append(v25, fi50)
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            case "":
                switch -2147483648 {
                case 2147483647:
                    for fi53 := 0; fi53 < 20; fi53++ {
                        switch fi53 {
                        case 653:
                            v54 := GenS2{F0: 2147483648, F1: 255}
                            _ = v54
                        case 32768:
                            var v55 int = (fi53 & (^(fi53)))
                            _ = v55
                        case 32767:
                            var v56 uint64 = 0
                            _ = v56
                        }
                        _ = fi53
                    }
                case 128:
                    var v57 uint8 = 42
                    _ = v57
                default:
                    var v58 float64 = 0.1
                    _ = v58
                }
            case "a":
                var v59 string = ((("" + "foo bar") + "") + ("foo bar" + "foo bar"))
                _ = v59
            default:
                var v60 uint64 = 0
                _ = v60
            }
            v2 = v2
            v1 = (^(v1))
        } else {
            var v61 int64 = 9223372036854775807
            _ = v61
        }
    }
    var v62 int16 = (-(v2))
    v63 := GenS5{F0: (2 <= 32767), F1: (v3 % (((-(v3))) | 1))}
    v62 = (v2 + -473)
    for fi64 := 0; fi64 < 1; fi64++ {
        v25 = append(v25, (fi64 - fi64))
        v2 -= (v2 % ((179) | 1))
        fmt.Println(fi64, v62, fi64)
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v62
    _ = v63
}

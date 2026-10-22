package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int16; F2 string }
type GenS5 struct { F0 int16 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint64 = 1
    var v5 float64 = 1e10
    _ = len(v3)
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v5 -= v5
            _ = fi7
        }
        for fi8 := 0; fi8 < 9; fi8++ {
            if true {
                var v9 int16 = 1
                v5 += -2.25
                fmt.Println(v9, v5)
                _ = v9
            } else {
                var v10 int16 = -729
                v4 += (v4 / ((18446744073709551615) | 1))
                _ = v10
            }
            if (42 >= 445) {
                fmt.Println(v5, fi6)
                switch ("a" + (("foo bar" + ("\n\t" + "foo bar")) + "foo bar")) {
                case "\n\t":
                    v3[fi6] = 2147483648
                }
            }
            _ = fi8
        }
        _ = fi6
    }
    fmt.Println(v1, v1)
    fmt.Println(v1)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint64 = v4
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v5)
    fmt.Println(v5, v5)
    switch 195 {
    case 413:
        v13 := GenS1{F0: ((42 == 544) && v1), F1: ""}
        _ = v13
    }
    v4 += 687
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := []int{fi14}
        v16 := GenS1{F0: v1, F1: "foo bar"}
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 uint16 = 65535
            v18 = (v18 * (v18 & (v18 - (v18 << ((65535) & 63)))))
            v19 := GenS2{F0: v18, F1: 0}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    if (771 == 4294967295) {
        if (17 == -92) {
            if false {
                switch 2147483647 {
                case -12:
                    switch ("\n\t" + ("foo bar" + ("\n\t" + ""))) {
                    case "\n\t":
                        v20 := GenS3{F0: 59}
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 11; fi21++ {
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v5 = -2.25
                                _ = fi22
                            }
                            for fi23 := 0; fi23 < 13; fi23++ {
                                v5 = 1e10
                                v24 := GenS4{F0: 32767, F1: 0, F2: (("foo bar" + ("\n\t" + "foo bar")) + "foo bar")}
                                if (!true) {
                                    var v25 float32 = (0.1 / 0.1)
                                    v25 += ((v25 - (float32(v5) / 1e10)) * v25)
                                    v25 = v25
                                    v26 := GenS1{F0: ((-9223372036854775808 == -131) && false), F1: "hello"}
                                    _ = v25
                                    _ = v26
                                }
                                _ = len(v3)
                                v27 := []int{37}
                                _ = fi23
                                _ = v24
                                _ = v27
                            }
                            _ = fi21
                        }
                    }
                }
                v4 = v4
            } else {
                switch "foo bar" {
                case "hello":
                    for fi28 := 0; fi28 < 15; fi28++ {
                        v4 += v4
                        var v29 uint16 = 855
                        v30 := []int{fi28, fi28, fi28}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "foo bar":
                    var v31 uint = 255
                    _ = v31
                }
                fmt.Println(v1, v1, v5)
            }
            delete(v3, int(v4))
        }
        fmt.Println(v4)
        v1 = true
    } else {
        for fi32 := 0; fi32 < 4; fi32++ {
            if (false || v1) {
                v33 := []int{(-(fi32))}
                delete(v3, fi32)
                fmt.Println(v4, v4, fi32)
                _ = v33
            } else {
                for fi34 := 0; fi34 < 9; fi34++ {
                    if v1 {
                        v1 = v1
                        var v35 float32 = 1.5
                        var v36 int16 = -669
                        v36 = -32768
                        switch "hello" {
                        case "foo bar":
                            v36 = (v36 & (v36 << ((487) & 63)))
                        }
                        _ = v35
                        _ = v36
                    } else {
                        var v37 bool = true
                        _ = v37
                    }
                    var v38 uint32 = 2
                    switch fi34 {
                    case 849:
                        v39 := GenS2{F0: 496, F1: 500}
                        _ = v39
                    }
                    v40 := GenS5{F0: int16(v38)}
                    _ = fi34
                    _ = v38
                    _ = v40
                }
                v41 := GenS4{F0: 1, F1: 1, F2: "hello"}
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 uint16 = 32767
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 uint8 = 127
                v1 = v1
                v46 := []int{4294967295, 460}
                var v47 uint32 = 2147483648
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
            var v48 uint64 = (^(v4))
            _ = fi32
            _ = v48
        }
        var v49 uint = 0
        _ = v49
    }
    _ = len(v3)
    v50 := GenS1{F0: v1, F1: "foo bar"}
    var v51 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}

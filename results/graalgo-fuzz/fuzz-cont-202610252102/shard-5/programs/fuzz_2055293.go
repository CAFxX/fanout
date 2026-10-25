package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32) int64 {
    p1 -= uint32(p0)
    fmt.Println(p0, p1)
    p1 = (p1 | (p1 % ((42) | 1)))
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int64, p1 uint32) int {
    if (127 > 255) {
        var v6 uint64 = 0
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                p1 = p1
                fmt.Println(p0, v6, v6)
                v6 -= (^(v6))
                _ = fi8
            }
            var v9 float32 = ((1e10 * 0.0) - 0.1)
            v10 := GenS1{F0: -32768}
            _ = fi7
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: -32768}
        _ = v11
        _ = v6
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "hello" {
        case "":
            p1 = p1
        case "a":
            v13 := GenS2{F0: (("" + "hello") + (("hello" + "foo bar") + "foo bar")), F1: 18446744073709551615, F2: 93}
            _ = v13
        case "\n\t":
            switch (("" + ("hello" + ("hello" + ("\n\t" + "hello")))) + "") {
            case "":
                for fi14 := 0; fi14 < 12; fi14++ {
                    if (((42 > 0) || (!false)) || false) {
                        switch fi14 {
                        case 101:
                            switch (fi12 + fi12) {
                            case 32767:
                                fmt.Println(fi12, fi14, p0)
                            case 127:
                                var v15 uint64 = 0
                                var v16 int32 = 0
                                _ = v15
                                _ = v16
                            }
                        case 2147483647:
                            p0 -= (-(p0))
                        case -1:
                            v17 := GenS3{F0: 32, F1: 1}
                            _ = v17
                        }
                    }
                    fmt.Println(p1, fi14)
                    v18 := []int{32768, (fi12 % (((fi14 % ((127) | 1))) | 1)), fi14, 42}
                    fmt.Println(p1)
                    _ = fi14
                    _ = v18
                }
            }
        default:
            switch fi12 {
            case 317:
                fmt.Println(p0, fi12)
            case 4294967295:
                v19 := GenS1{F0: 42}
                _ = v19
            case 32768:
                if (!true) {
                    p1 += p1
                    v20 := []int{fi12}
                    if len(v20) > 0 {
                        v20[0] = (fi12 - (fi12 ^ fi12))
                    }
                    _ = v20
                } else {
                    v21 := []int{127, (^(fi12))}
                    for fi22 := 0; fi22 < 9; fi22++ {
                        switch fi22 {
                        case 4294967295:
                            _ = len(v21)
                            var v23 int32 = -1
                            var v24 uint = 42
                            _ = v23
                            _ = v24
                        case 9223372036854775807:
                            var v25 int = 498
                            _ = v25
                        }
                        _ = fi22
                    }
                    v26 := GenS1{F0: -668}
                    _ = v21
                    _ = v26
                }
            }
        }
        var v27 uint8 = 69
        v28 := GenS4{F0: 127, F1: -56}
        _ = fi12
        _ = v27
        _ = v28
    }
    _ = p0
    _ = p1
    return 436
}

func main() {
    var v1 uint64 = 0
    var v2 int32 = 2147483647
    var v3 uint64 = v1
    var v4 uint8 = 1
    var v5 bool = true
    if ((733 >= 2) && (!v5)) {
        for fi29 := 0; fi29 < 5; fi29++ {
            fmt.Println(v1)
            _ = fi29
        }
    }
    var v30 uint32 = 1
    var v31 int16 = 640
    if v5 {
        switch 42 {
        case 800:
            for fi32 := 0; fi32 < 6; fi32++ {
                switch ("hello" + "foo bar") {
                case "\n\t":
                    v33 := GenS5{F0: 1, F1: v4}
                    _ = v33
                }
                _ = fi32
            }
        case 2147483647:
            v34 := GenS6{F0: 609, F1: (-(v31))}
            _ = v34
        case 985:
            v35 := []int{128, 4294967295, 419}
            _ = v35
        default:
            fmt.Println(v4, v3)
        }
        v36 := []int{-9223372036854775808}
        _ = v36
    } else {
        var v37 int32 = v2
        switch 42 {
        case -171:
            if (0 >= int16(v3)) {
                if (!(2147483647 == genFn1(42, 2))) {
                    if (4294967295 > 18446744073709551615) {
                        v31 = (v31 % (((v31 - 625)) | 1))
                        var v38 uint = 712
                        _ = v38
                    }
                    v1 = v1
                }
                v39 := []int{28, 0}
                _ = v39
            }
        }
        _ = v37
    }
    switch 203 {
    case 4294967295:
        v40 := []int{461, 32767}
        _ = v40
    case 42:
        if (32767 <= -2147483648) {
            for fi41 := 0; fi41 < 14; fi41++ {
                for fi42 := 0; fi42 < 14; fi42++ {
                    v43 := GenS7{F0: fi42, F1: (v30 & v30)}
                    var v44 float64 = float64(1e10)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = fi41
            }
            v4 = (v4 | 42)
            var v45 float64 = 0.1
            _ = v45
        } else {
            var v46 uint16 = 65535
            v47 := GenS3{F0: 247, F1: 2}
            _ = v46
            _ = v47
        }
    case 128:
        var v48 uint = 65535
        _ = v48
    }
    v2 = v2
    v49 := []int{32767, -2147483648, 386}
    v2 = (-(v2))
    if v5 {
        var v50 int32 = v2
        _ = v50
    } else {
        if len(v49) > 0 {
            v49[0] = 0
        }
        v1 -= v1
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        fmt.Println(v3)
        fmt.Println(v2, v5)
        v2 -= v2
        _ = fi51
    }
    fmt.Println(v5, v4)
    if v5 {
        for fi52 := 0; fi52 < 5; fi52++ {
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
            v31 = (-(v31))
            _ = fi52
        }
        var v53 uint64 = (v1 / ((v3) | 1))
        fmt.Println(v4, v1, v30)
        _ = v53
    }
    v5 = (!(0 != 2))
    if true {
        if (!true) {
            v54 := GenS8{F0: 42}
            fmt.Println(v1)
            _ = v54
        } else {
            var v55 float64 = -2.25
            v2 = 2147483647
            _ = v55
        }
    }
    v31 = v31
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
    _ = v5
}

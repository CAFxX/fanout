package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int8; F1 int8; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 int64) uint8 {
    v4 := []int{}
    var v5 uint16 = 1
    fmt.Println(p0, p0)
    var v6 bool = (!true)
    switch "a" {
    case "":
        if false {
            switch ("\n\t" + "a") {
            case "a":
                var v7 int32 = -56
                _ = v7
            }
            fmt.Println(p0, p0, p0)
        }
    case "\n\t":
        v6 = v6
    case "hello":
        var v8 bool = (42 == -260)
        _ = v8
    default:
        var v9 int32 = -568
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 int) int64 {
    var v10 bool = false
    v10 = (0 == 866)
    var v11 int8 = -128
    p0 = p0
    _ = p0
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 float64 = 0.0
    var v2 int16 = 774
    var v3 bool = true
    if ("foo bar" <= "a") {
        var v12 uint16 = 1
        _ = v12
    } else {
        var v13 float64 = v1
        _ = v13
    }
    if false {
        if ((4294967295 <= -2147483648) && false) {
            v1 = float64(1.5)
            fmt.Println(v3)
            v1 = 0.0
        }
        v14 := GenS1{F0: 1, F1: 1, F2: 69}
        var v15 bool = v3
        _ = v14
        _ = v15
    } else {
        if v3 {
            if (v3 || (42 < 2147483647)) {
                v3 = (v3 && ((!(2 >= 2147483648)) && v3))
                v3 = (!(0 > (v2 & 1)))
            } else {
                for fi16 := 0; fi16 < 20; fi16++ {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := []int{fi17}
                        _ = fi17
                        _ = v18
                    }
                    v19 := []int{fi16}
                    switch fi16 {
                    case 9223372036854775807:
                        var v20 uint16 = 937
                        _ = v20
                    case 2147483647:
                        v1 = (0.0 / float64((1e10 + 3.14159)))
                    case -2147483648:
                        switch ("" + "a") {
                        case "":
                            switch "" {
                            case "":
                                var v21 int64 = genFn1(fi16)
                                _ = v21
                            case "\n\t":
                                _ = len(v19)
                            }
                        case "a":
                            var v22 float32 = (1.5 + 3.14159)
                            _ = v22
                        }
                    default:
                        switch "" {
                        case "a":
                            v23 := GenS1{F0: 1, F1: 1, F2: 740}
                            _ = v23
                        case "\n\t":
                            switch (("\n\t" + ((((("\n\t" + "a") + "hello") + "\n\t") + ("hello" + "")) + "")) + "\n\t") {
                            case "hello":
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            case "":
                                var v24 uint8 = 28
                                _ = v24
                            case "foo bar":
                                v3 = v3
                            }
                        case "foo bar":
                            var v25 float32 = 1e10
                            _ = v25
                        default:
                            var v26 int32 = -101
                            _ = v26
                        }
                    }
                    _ = fi16
                    _ = v19
                }
                var v27 int8 = -128
                _ = v27
            }
            v28 := GenS2{F0: 255, F1: 68}
            _ = v28
        } else {
            v29 := []int{}
            _ = v29
        }
    }
    v30 := GenS3{F0: -128, F1: 0, F2: 42}
    v3 = v3
    v2 -= (v2 - (v2 << ((255) & 63)))
    switch -985 {
    case 32768:
        var v31 int64 = 0
        _ = v31
    case 1:
        v3 = (false || (238 <= 1))
    case 596:
        v1 = v1
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := GenS4{F0: 569}
        if v3 {
            v34 := []int{(fi32 - (fi32 << ((2) & 63))), (fi32 & fi32), -1, (fi32 << ((511) & 63))}
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = v34
        } else {
            v2 += 1
            if v3 {
                if v3 {
                    v35 := []int{}
                    _ = v35
                }
                v36 := []int{}
                _ = v36
            } else {
                v37 := GenS1{F0: 944, F1: 0, F2: genFn1((fi32 ^ (fi32 << ((501) & 63))))}
                _ = v37
            }
        }
        v38 := []int{128, fi32}
        _ = fi32
        _ = v33
        _ = v38
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 bool = false
        for fi41 := 0; fi41 < 6; fi41++ {
            fmt.Println(v40, fi39)
            v42 := GenS3{F0: -91, F1: 83, F2: -1}
            v1 += -2.25
            var v43 uint32 = 476
            switch (fi41 | fi41) {
            case -354:
                if (127 < 36) {
                    v3 = ((^(v43)) >= v43)
                    if v40 {
                        fmt.Println(v43, fi39)
                        if v3 {
                            var v44 uint8 = 127
                            fmt.Println(v44, v1)
                            if v40 {
                                var v45 uint64 = 822
                                v40 = false
                                v45 = v45
                                _ = v45
                            }
                            _ = v44
                        } else {
                            var v46 int16 = (v2 << ((1) & 63))
                            v43 = (v43 + (v43 / (((v43 % (((v43 + 4294967295)) | 1))) | 1)))
                            v43 = 968
                            _ = v46
                        }
                        if v3 {
                            switch (fi39 << ((762) & 63)) {
                            case -373:
                                switch (^(fi41)) {
                                case -258:
                                    v40 = (298 >= 105)
                                }
                            case -678:
                                v1 += (((0.0 + 1.5) * 3.14159) * v1)
                            default:
                                if false {
                                    v1 = ((1.5 * float64(0.0)) - (v1 / (0.0 / ((float64(float32(0.1)) + (v1 + v1)) / 1.0))))
                                }
                            }
                            fmt.Println(v43)
                        } else {
                            v40 = (!(v2 < v2))
                            var v47 uint64 = 271
                            _ = v47
                        }
                    }
                    if false {
                        v48 := []int{(-(fi39))}
                        if v3 {
                            v40 = v40
                            fmt.Println(v43, v2, v2)
                            v43 = (v43 - (v43 - v43))
                            v2 = (v2 << ((396) & 63))
                            var v49 string = ("foo bar" + "")
                            _ = v49
                        } else {
                            if v40 {
                                var v50 int = fi39
                                _ = len(v48)
                                _ = v50
                            }
                        }
                        _ = v48
                    }
                } else {
                    v51 := GenS3{F0: 100, F1: -1, F2: 42}
                    var v52 uint16 = 65535
                    if ((fi41 & fi39) == 0) {
                        v43 += (v43 % ((2) | 1))
                        v1 = ((v1 + float64(float32(1e10))) + (v1 * ((float64((3.14159 + (-2.25 * -2.25))) + v1) / v1)))
                        v53 := []int{}
                        _ = v53
                    }
                    v54 := GenS4{F0: -740}
                    var v55 int16 = (-(v2))
                    var v56 int64 = 824
                    _ = v51
                    _ = v52
                    _ = v54
                    _ = v55
                    _ = v56
                }
            default:
                v57 := GenS5{F0: 1}
                v58 := GenS6{F0: 130, F1: -2147483648, F2: 9223372036854775807}
                _ = v57
                _ = v58
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        v59 := []int{fi39, fi39}
        _ = fi39
        _ = v40
        _ = v59
    }
    if false {
        v3 = v3
        fmt.Println(v1)
        var v60 uint32 = 829
        _ = v60
    }
    var v61 string = "hello"
    if v3 {
        for fi62 := 0; fi62 < 5; fi62++ {
            var v63 int64 = -1
            if (0 <= 18446744073709551615) {
                if v3 {
                    v64 := GenS3{F0: -64, F1: 0, F2: -1}
                    _ = v64
                }
                v1 = 0.0
            }
            _ = fi62
            _ = v63
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v61
}

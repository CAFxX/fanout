package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int; F1 int16; F2 string }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int32; F1 int64; F2 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint; F1 uint16; F2 bool }
func genFn0(p0 int64, p1 uint64, p2 int) uint32 {
    v5 := GenS1{F0: 65535}
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(p2, p2, fi6)
        fmt.Println(p2)
        _ = fi6
    }
    if (!true) {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS2{F0: 333, F1: 18446744073709551615}
                v10 := GenS3{F0: 2}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        p1 -= (p1 >> ((1) & 63))
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            p1 = (p1 - (p1 / (((p1 & p1)) | 1)))
            _ = fi11
        }
        p2 += 4294967295
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{}
            if false {
                if (662 != 42) {
                    if (!true) {
                        p0 = (p0 * p0)
                    }
                    v14 := GenS4{F0: 2147483647, F1: 90, F2: 255}
                    _ = v14
                } else {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    for fi15 := 0; fi15 < 19; fi15++ {
                        p2 = fi12
                        _ = fi15
                    }
                }
            } else {
                for fi16 := 0; fi16 < 11; fi16++ {
                    p2 = (fi16 << ((255) & 63))
                    var v17 uint32 = 65535
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 1000
    v2 := []string{"\n\t", "foo bar", "\n\t"}
    var v3 uint64 = 0
    v4 := make(map[int]int)
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 float32 = (0.1 * 1.5)
        _ = fi18
        _ = v19
    }
    v3 += (v3 ^ v3)
    v20 := GenS5{F0: 127, F1: (v1 & (v1 + (v1 * (v1 - (v1 - (v1 * v1)))))), F2: ((("foo bar" + "foo bar") + "a") + ((("a" + "") + (("hello" + "") + "\n\t")) + ("\n\t" + ("" + ("hello" + "hello")))))}
    if true {
        var v21 uint32 = 2147483648
        _ = v21
    } else {
        v22 := GenS6{F0: 2147483647, F1: 100}
        _ = v22
    }
    v3 -= (v3 / (((v3 * v3)) | 1))
    var v23 int64 = 42
    if (2147483648 < 1) {
        switch "" {
        case "":
            if ((((!(!(!((1e10 - (float64(float32(1.5)) * 1.5)) > 1.5)))) || false) && true) || (42 != 4294967295)) {
                var v24 bool = (true && true)
                _ = v24
            } else {
                v23 += 9223372036854775807
            }
        case "\n\t":
            var v25 int = 4294967295
            _ = v25
        case "foo bar":
            v26 := GenS7{F0: 2147483647, F1: -9223372036854775808, F2: 4294967295}
            _ = v26
        default:
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(v3, v3)
                var v28 bool = true
                _ = fi27
                _ = v28
            }
        }
        switch "a" {
        case "foo bar":
            v29 := GenS8{F0: 0}
            _ = v29
        default:
            fmt.Println(v23, v1, v3)
        }
    }
    var v30 int64 = -9223372036854775808
    var v31 int8 = 127
    if (true || true) {
        switch int(v1) {
        case 4294967295:
            if (false && false) {
                v32 := v4[793]
                _ = v32
            }
        case -2147483648:
            switch "" {
            case "foo bar":
                if false {
                    fmt.Println(v23)
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if false {
                            if ((42 != int32(v1)) || ((0.1 - 0.1) == -2.25)) {
                                var v34 int = fi33
                                if true {
                                    var v35 int64 = v23
                                    fmt.Println(v23)
                                    _ = v35
                                }
                                _ = v34
                            }
                            var v36 string = "a"
                            v37 := GenS9{F0: 255, F1: 18446744073709551615}
                            _ = v36
                            _ = v37
                        } else {
                            var v38 uint32 = 49
                            v1 -= (v1 >> ((18446744073709551615) & 63))
                            _ = v38
                        }
                        switch fi33 {
                        case 184:
                            var v39 string = ("\n\t" + ("foo bar" + (("hello" + "") + "hello")))
                            _ = v39
                        case 9223372036854775807:
                            v30 -= (v23 << ((1) & 63))
                        case 128:
                            v30 += v23
                        default:
                            fmt.Println(v31)
                        }
                        _ = fi33
                    }
                    v40 := GenS6{F0: -2147483648, F1: v31}
                    _ = v40
                }
            case "a":
                v41 := GenS7{F0: -1, F1: 1, F2: 42}
                _ = v41
            }
        default:
            v42 := GenS5{F0: 32767, F1: -32768, F2: "\n\t"}
            _ = v42
        }
        v43 := GenS1{F0: 255}
        v44 := GenS8{F0: 2147483647}
        _ = v43
        _ = v44
    }
    var v45 int64 = v30
    v3 = v3
    if (0 >= 4294967295) {
        v46 := GenS3{F0: genFn0(-439, (^(v3)), 32767)}
        switch "hello" {
        case "foo bar":
            for fi47 := 0; fi47 < 4; fi47++ {
                v48 := GenS9{F0: 42, F1: 255}
                _ = fi47
                _ = v48
            }
        case "hello":
            v23 += 1
        case "\n\t":
            switch 4294967295 {
            case -9223372036854775808:
                for fi49 := 0; fi49 < 5; fi49++ {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        v23 -= (v45 + v30)
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        _ = fi50
                    }
                    fmt.Println(v30)
                    _ = fi49
                }
            case -2147483648:
                v51 := GenS10{F0: v3}
                _ = v51
            case 1:
                var v52 int = 652
                _ = v52
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v46
    } else {
        fmt.Println(v30, v23, v45)
        v53 := GenS11{F0: 2, F1: 255, F2: (255 != 42)}
        _ = v53
    }
    v54 := GenS11{F0: 368, F1: 398, F2: (true && (!(0 != 2147483647)))}
    v55 := GenS11{F0: 255, F1: 32767, F2: false}
    fmt.Println(v45, v1)
    v23 -= (^(v30))
    if (true || (true || false)) {
        var v56 uint8 = 0
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v45
    _ = v54
    _ = v55
}

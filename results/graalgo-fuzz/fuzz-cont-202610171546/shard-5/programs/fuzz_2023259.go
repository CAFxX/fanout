package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int64; F1 int16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int32; F2 int8 }
func genFn0() int64 {
    var v6 int = -9223372036854775808
    if (0 > 2147483648) {
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v6 += fi8
                v9 := GenS1{F0: -17, F1: ""}
                if ((fi7 - 65) >= (^(v6))) {
                    switch ("" + ((("a" + "") + "a") + "a")) {
                    case "foo bar":
                        v6 = fi7
                    case "\n\t":
                        v6 = (fi8 >> ((443) & 63))
                    case "":
                        if false {
                            fmt.Println(fi7, fi7, fi7)
                            v10 := []int{0}
                            _ = v10
                        } else {
                            var v11 int8 = -4
                            v12 := GenS2{F0: (546 <= 65535), F1: 65535}
                            _ = v11
                            _ = v12
                        }
                        v6 -= -1
                        if false {
                            v6 = (v6 / (((fi7 << ((695) & 63))) | 1))
                            var v13 float64 = 0.1
                            _ = v13
                        } else {
                            v14 := GenS3{F0: -1, F1: -32768, F2: 42}
                            _ = v14
                        }
                    default:
                        v15 := GenS3{F0: 869, F1: 32767, F2: 0}
                        _ = v15
                    }
                    v6 = 1
                    var v16 int64 = 1
                    fmt.Println(fi8, v6, v16)
                    _ = v16
                } else {
                    v17 := GenS4{F0: 325, F1: 1}
                    if true {
                        v18 := GenS3{F0: 9223372036854775807, F1: 42, F2: -24}
                        _ = v18
                    } else {
                        v19 := GenS4{F0: 856, F1: 512}
                        if (-1 < 1) {
                            var v20 int64 = 0
                            v6 -= -2147483648
                            switch "\n\t" {
                            case "foo bar":
                                v6 = (fi8 & 42)
                            default:
                                var v21 uint16 = 698
                                _ = v21
                            }
                            _ = v20
                        }
                        v22 := GenS3{F0: -716, F1: -1, F2: 0}
                        _ = v19
                        _ = v22
                    }
                    v23 := GenS3{F0: -9223372036854775808, F1: -1, F2: 37}
                    var v24 int64 = 0
                    _ = v17
                    _ = v23
                    _ = v24
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    }
    if false {
        for fi25 := 0; fi25 < 13; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                fmt.Println(fi26, fi26)
                v27 := GenS4{F0: 2, F1: -1}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    if true {
        for fi28 := 0; fi28 < 5; fi28++ {
            v6 += (v6 + (v6 * (fi28 << ((466) & 63))))
            v6 = (v6 | v6)
            switch (fi28 & 128) {
            case 9223372036854775807:
                switch v6 {
                case 4294967295:
                    for fi29 := 0; fi29 < 15; fi29++ {
                        v6 -= fi29
                        v30 := []int{552, fi28}
                        var v31 float64 = 0.0
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case 32768:
                    var v32 uint = 1
                    _ = v32
                case -180:
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if ((((("\n\t" + "\n\t") + "") + "\n\t") + ("a" + "\n\t")) < "hello") {
                            var v34 int8 = -1
                            var v35 float32 = float32(3.14159)
                            _ = v34
                            _ = v35
                        }
                        _ = fi33
                    }
                default:
                    v36 := GenS5{F0: 128}
                    _ = v36
                }
            case 32767:
                var v37 int16 = 42
                _ = v37
            }
            _ = fi28
        }
        var v38 string = (("hello" + "hello") + ("hello" + "foo bar"))
        v39 := []int{(-(v6))}
        _ = v38
        _ = v39
    }
    _ = v6
    return 42
}
func genFn1(p0 int64) uint8 {
    var v40 int64 = p0
    var v41 float64 = float64(-2.25)
    v42 := GenS4{F0: 18446744073709551615, F1: 42}
    _ = p0
    _ = v40
    _ = v41
    _ = v42
    return 34
}

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 255
    var v3 int16 = -1
    v4 := []string{(("hello" + "foo bar") + ("a" + (("foo bar" + (("\n\t" + (("foo bar" + "a") + "a")) + "\n\t")) + ("hello" + (("\n\t" + ("\n\t" + ("hello" + ("foo bar" + "\n\t")))) + ("\n\t" + ("hello" + "foo bar"))))))), "", ((("" + "") + "foo bar") + "hello")}
    v5 := make(map[int]uint32)
    for fi43 := 0; fi43 < 13; fi43++ {
        v2 = (v2 * (^(v2)))
        if true {
            v3 = (v3 - v3)
            v1 = v1
        } else {
            v3 = v3
            if len(v4) > 0 {
                v4[0] = ""
            }
            fmt.Println(fi43)
        }
        _ = fi43
    }
    v2 -= (v2 * 65535)
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    var v44 string = ""
    if false {
        v45 := GenS6{F0: 97, F1: -1, F2: 42}
        v3 -= v3
        _ = v45
    } else {
        switch (((v44 + v44) + "foo bar") + v44) {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            v2 = 502
        }
    }
    switch 1 {
    case 791:
        var v46 float32 = v1
        _ = v46
    case 673:
        v47 := GenS6{F0: 0, F1: 0, F2: -46}
        _ = v47
    default:
        v4 = append(v4, v44)
    }
    var v48 bool = false
    delete(v5, 521)
    var v49 int8 = -1
    v49 = (v49 | (v49 - int8(v2)))
    if (!true) {
        v3 -= v3
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi50 := 0; fi50 < 15; fi50++ {
        v2 -= v2
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v48
    _ = v49
    _ = v5
}

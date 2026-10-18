package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint64; F1 string; F2 int }
type GenS5 struct { F0 int16; F1 int32; F2 int32 }
type GenS6 struct { F0 int32; F1 int16 }
func genFn0(p0 uint64) string {
    var v4 uint = 65535
    v4 += v4
    _ = p0
    _ = v4
    return "foo bar"
}

func main() {
    v1 := []int{-1, -9223372036854775808}
    var v2 int64 = -9223372036854775808
    var v3 int = -705
    v5 := GenS1{F0: 65535, F1: 127}
    if (!(12 == 0)) {
        v2 = v2
    } else {
        v6 := GenS1{F0: 42, F1: 206}
        v2 = (v2 / ((v2) | 1))
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS2{F0: true}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        switch "\n\t" {
        case "foo bar":
            v2 += (v2 * (v2 - (v2 & (v2 * (v2 ^ (^(v2)))))))
        case "a":
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := GenS3{F0: 0, F1: -1}
                var v12 uint16 = 65535
                switch "hello" {
                case "\n\t":
                    var v13 int = 337
                    _ = v13
                default:
                    var v14 int8 = -1
                    _ = v14
                }
                fmt.Println(v12)
                if len(v1) > 0 {
                    v1[0] = fi10
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi9
    }
    v1 = append(v1, (v3 >> ((4294967295) & 63)))
    var v15 float64 = 1.5
    var v16 uint64 = 612
    _ = len(v1)
    v16 += v16
    for fi17 := 0; fi17 < 17; fi17++ {
        v16 -= (^(v16))
        v1 = append(v1, (fi17 / ((32767) | 1)))
        _ = fi17
    }
    v18 := GenS4{F0: (v16 & 0), F1: "hello", F2: (^(v3))}
    var v19 int8 = -3
    if false {
        var v20 int = (-(v3))
        if false {
            var v21 string = ("hello" + "a")
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v22 uint64 = (v16 / (((^(v16))) | 1))
            _ = v21
            _ = v22
        }
        var v23 int = v20
        _ = v20
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v3 -= (fi24 & fi24)
        _ = fi24
    }
    if true {
        v3 = int(v16)
        for fi25 := 0; fi25 < 17; fi25++ {
            switch 74 {
            case 42:
                switch "a" {
                case "foo bar":
                    var v26 int8 = (v19 + (-(v19)))
                    _ = v26
                case "a":
                    _ = len(v1)
                }
            default:
                switch "hello" {
                case "\n\t":
                    var v27 uint = 255
                    _ = v27
                case "a":
                    for fi28 := 0; fi28 < 2; fi28++ {
                        _ = len(v1)
                        if false {
                            switch v3 {
                            case 2147483647:
                                fmt.Println(fi28, v16)
                            case 1:
                                if (0 == (v19 | v19)) {
                                    fmt.Println(v19, fi25, v19)
                                    if true {
                                        if (!(0.0 < 3.14159)) {
                                            if (1.5 == (-2.25 / -2.25)) {
                                                var v29 uint = 1
                                                if (!false) {
                                                    v30 := GenS1{F0: 2147483648, F1: 2}
                                                    v15 += v15
                                                    v31 := GenS5{F0: -949, F1: 635, F2: -713}
                                                    _ = v30
                                                    _ = v31
                                                }
                                                _ = v29
                                            }
                                        } else {
                                            var v32 int8 = 126
                                            _ = v32
                                        }
                                    }
                                    v3 = (v3 * (fi28 / ((fi25) | 1)))
                                }
                            }
                        }
                        v3 = -528
                        _ = fi28
                    }
                }
            }
            _ = fi25
        }
        switch v3 {
        case 1:
            v33 := GenS6{F0: 0, F1: -1}
            _ = v33
        case -2147483648:
            for fi34 := 0; fi34 < 14; fi34++ {
                switch ("" + ("" + ("a" + ("" + ("foo bar" + "foo bar"))))) {
                case "foo bar":
                    for fi35 := 0; fi35 < 18; fi35++ {
                        v1 = append(v1, fi35)
                        _ = fi35
                    }
                case "hello":
                    for fi36 := 0; fi36 < 15; fi36++ {
                        if len(v1) > 0 {
                            v1[0] = v3
                        }
                        v37 := GenS5{F0: 0, F1: 1, F2: 1}
                        v38 := GenS6{F0: 701, F1: 32767}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                fmt.Println(fi34, v19)
                v19 -= (v19 & v19)
                _ = fi34
            }
        case 32768:
            v39 := GenS4{F0: 4294967295, F1: "foo bar", F2: 4294967295}
            _ = v39
        }
    }
    if (678 >= (v2 - v2)) {
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 int64 = v2
            var v42 float64 = -2.25
            for fi43 := 0; fi43 < 15; fi43++ {
                _ = len(v1)
                _ = fi43
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        switch "foo bar" {
        case "\n\t":
            var v44 float32 = ((-2.25 - -2.25) * float32(((v15 + 3.14159) - v15)))
            _ = v44
        default:
            var v45 uint64 = v16
            _ = v45
        }
        if false {
            for fi46 := 0; fi46 < 16; fi46++ {
                var v47 int16 = -764
                _ = fi46
                _ = v47
            }
            for fi48 := 0; fi48 < 17; fi48++ {
                for fi49 := 0; fi49 < 10; fi49++ {
                    if false {
                        v50 := GenS5{F0: 42, F1: -1, F2: 493}
                        _ = v50
                    }
                    _ = fi49
                }
                var v51 int64 = v2
                _ = fi48
                _ = v51
            }
            v2 = v2
        } else {
            v19 = (v19 | v19)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}

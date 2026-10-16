package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8) uint8 {
    p0 -= p0
    for fi4 := 0; fi4 < 17; fi4++ {
        p0 += p0
        _ = fi4
    }
    var v5 bool = (!false)
    var v6 uint16 = 417
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return (p0 << ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := []uint32{350, 2}
    switch ("foo bar" + "") {
    case "a":
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int32 = -2147483648
            v9 := GenS1{F0: 68}
            v10 := GenS2{F0: 1, F1: 776}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    case "\n\t":
        if true {
            var v11 int16 = 32767
            _ = v11
        }
    case "foo bar":
        var v12 uint16 = 65535
        _ = v12
    default:
        var v13 int8 = -33
        _ = v13
    }
    var v14 int64 = -9223372036854775808
    v14 = (v14 | (-(v14)))
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS3{F0: 0}
        switch ("hello" + ("a" + "\n\t")) {
        case "a":
            switch 32767 {
            case 127:
                delete(v1, 1)
            case 128:
                fmt.Println(v14, fi15)
            case 0:
                var v17 uint = 1
                _ = v17
            default:
                v14 -= (-(v14))
            }
        default:
            if true {
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 int8 = -128
                    _ = fi18
                    _ = v19
                }
            } else {
                if (0 == 687) {
                    var v20 float32 = 0.1
                    _ = v20
                } else {
                    switch fi15 {
                    case -723:
                        delete(v2, "hello")
                    case -2147483648:
                        v21 := GenS3{F0: 1}
                        _ = v21
                    case -164:
                        var v22 int64 = -389
                        _ = v22
                    }
                }
            }
        }
        var v23 int32 = 624
        _ = fi15
        _ = v16
        _ = v23
    }
    v14 -= (v14 / ((v14) | 1))
    for fi24 := 0; fi24 < 13; fi24++ {
        switch ("a" + (("hello" + "foo bar") + ("hello" + ("\n\t" + ("foo bar" + "a"))))) {
        case "foo bar":
            v14 = (v14 - (v14 % ((0) | 1)))
        case "hello":
            if false {
                for fi25 := 0; fi25 < 3; fi25++ {
                    _ = len(v3)
                    _ = fi25
                }
                if true {
                    var v26 bool = (!true)
                    _ = v26
                } else {
                    v27 := GenS1{F0: 127}
                    _ = v27
                }
            }
        case "":
            v28 := GenS1{F0: 1}
            _ = v28
        default:
            fmt.Println(fi24, v14)
        }
        fmt.Println(fi24)
        _ = fi24
    }
    delete(v1, 2147483647)
    delete(v1, 4294967295)
    var v29 uint64 = 42
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    v14 += (v14 + (v14 & v14))
    if len(v3) > 0 {
        v3[0] = 1
    }
    if true {
        v29 -= v29
        var v30 string = ("\n\t" + ("" + "hello"))
        if (-1 < 32768) {
            v31 := GenS2{F0: -410, F1: v29}
            if (776 <= 65535) {
                _ = len(v2)
                var v32 int = 127
                _ = v32
            }
            var v33 int64 = 1
            _ = v31
            _ = v33
        } else {
            if (-1 != 42) {
                if (42 <= 4294967295) {
                    delete(v1, 42)
                } else {
                    if ((v29 << ((231) & 63)) < (v29 << ((424) & 63))) {
                        if (float64(3.14159) >= (1.5 - 3.14159)) {
                            var v34 float32 = -2.25
                            v30 = v30
                            var v35 string = "foo bar"
                            _ = v34
                            _ = v35
                        } else {
                            switch v30 {
                            case "foo bar":
                                v30 = (v30 + v30)
                            default:
                                if (genFn0(5) < 127) {
                                    for fi36 := 0; fi36 < 4; fi36++ {
                                        if ((fi36 - (fi36 >> ((18446744073709551615) & 63))) < (fi36 ^ (fi36 - (-(fi36))))) {
                                            for fi37 := 0; fi37 < 18; fi37++ {
                                                v38 := GenS2{F0: 1000, F1: (^(v29))}
                                                _ = fi37
                                                _ = v38
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        fmt.Println(v30, v14)
                                        var v39 uint16 = 0
                                        _ = fi36
                                        _ = v39
                                    }
                                    var v40 uint = 789
                                    fmt.Println(v29, v30, v29)
                                    _ = v40
                                } else {
                                    var v41 int64 = (-(v14))
                                    _ = v41
                                }
                            }
                        }
                    }
                }
                var v42 uint64 = v29
                _ = v42
            }
            v14 = (v14 ^ (v14 % (((v14 | 1)) | 1)))
            v29 = 133
        }
        _ = v30
    } else {
        v29 += v29
    }
    var v43 uint8 = 42
    v44 := v2["hello"]
    var v45 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v43
    _ = v44
    _ = v45
}

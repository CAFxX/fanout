package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
func genFn0() uint64 {
    var v5 int16 = -1
    v6 := []int{9223372036854775807, -9223372036854775808}
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int16 = -1
        v10 := GenS1{F0: 0, F1: v5}
        if true {
            v11 := GenS1{F0: 9223372036854775807, F1: (v9 % (((v9 * v9)) | 1))}
            for fi12 := 0; fi12 < 14; fi12++ {
                switch -9223372036854775808 {
                case 1:
                    var v13 string = ("" + "a")
                    _ = v13
                default:
                    v14 := GenS2{F0: 446}
                    _ = v14
                }
                fmt.Println(v7, fi8, v9)
                _ = fi12
            }
            v15 := GenS3{F0: (v9 / ((-1) | 1)), F1: 1, F2: 2147483647}
            _ = v11
            _ = v15
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v16 string = "foo bar"
    _ = v16
    _ = v5
    _ = v6
    _ = v7
    return 581
}
func genFn1(p0 int32) string {
    var v17 int32 = -2147483648
    for fi18 := 0; fi18 < 17; fi18++ {
        v17 -= (p0 * v17)
        v19 := GenS3{F0: -406, F1: 2, F2: 872}
        _ = fi18
        _ = v19
    }
    p0 -= (v17 >> ((312) & 63))
    _ = p0
    _ = v17
    return "\n\t"
}

func main() {
    var v1 float64 = 0.0
    var v2 int = 2147483647
    var v3 int8 = 101
    v4 := make(map[string]string)
    var v20 uint = 275
    for fi21 := 0; fi21 < 20; fi21++ {
        if false {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS1{F0: int64(fi22), F1: -32768}
                delete(v4, "\n\t")
                if (((v3 - 100) > (v3 & v3)) || true) {
                    v2 = 571
                } else {
                    delete(v4, "foo bar")
                }
                _ = fi22
                _ = v23
            }
        }
        var v24 string = ("foo bar" + "")
        switch (v24 + v24) {
        case "hello":
            if false {
                var v25 float64 = v1
                _ = v25
            } else {
                v26 := GenS3{F0: 32767, F1: 0, F2: 42}
                _ = v26
            }
        case "\n\t":
            var v27 uint = v20
            _ = v27
        }
        _ = fi21
        _ = v24
    }
    v20 = v20
    var v28 uint32 = 0
    v28 -= (v28 * (^(v28)))
    fmt.Println(v3)
    switch "\n\t" {
    case "":
        v28 += (v28 + 421)
    }
    if (!(false && true)) {
        if ("a" < "foo bar") {
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := []int{(fi29 & 0), fi29, int(v3), (fi29 - v2)}
                if false {
                    switch (("\n\t" + "foo bar") + "") {
                    case "foo bar":
                        v28 = (v28 * v28)
                    case "hello":
                        for fi31 := 0; fi31 < 12; fi31++ {
                            switch -1 {
                            case 127:
                                fmt.Println(v28)
                            case 1:
                                fmt.Println(v3)
                            }
                            _ = fi31
                        }
                    case "a":
                        if (!(!false)) {
                            var v32 uint16 = 255
                            fmt.Println(v20)
                            _ = v32
                        }
                    default:
                        if (42 == (v28 - (^(v28)))) {
                            fmt.Println(v2)
                            v2 = v2
                        }
                    }
                }
                var v33 uint64 = 4294967295
                _ = fi29
                _ = v30
                _ = v33
            }
        }
    } else {
        v34 := GenS1{F0: 0, F1: -788}
        switch v2 {
        case -218:
            v28 = (^(v28))
        case 0:
            v35 := v4["\n\t"]
            _ = v35
        case -1:
            _ = len(v4)
        }
        switch "" {
        case "a":
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := GenS2{F0: (fi36 + v2)}
                v1 += float64(((-2.25 / 1.5) * (-2.25 / -2.25)))
                v38 := GenS4{F0: (v28 - (^(v28))), F1: 536, F2: 0}
                _ = fi36
                _ = v37
                _ = v38
            }
        case "\n\t":
            switch (v2 / (((v2 / (((v2 & (v2 | (v2 % ((-987) | 1))))) | 1))) | 1)) {
            case -1:
                switch "a" {
                case "a":
                    fmt.Println(v3, v2)
                case "":
                    var v39 uint = 0
                    _ = v39
                case "hello":
                    v28 -= v28
                }
            default:
                if false {
                    v40 := GenS5{F0: 1}
                    var v41 uint64 = 4294967295
                    v42 := GenS6{F0: -30}
                    _ = v40
                    _ = v41
                    _ = v42
                }
            }
        default:
            v2 = (-(v2))
        }
        _ = v34
    }
    fmt.Println(v3, v1, v1)
    v2 -= v2
    v43 := GenS5{F0: (v3 << ((v20) & 63))}
    for fi44 := 0; fi44 < 17; fi44++ {
        v28 = (^(v28))
        var v45 int64 = 458
        switch fi44 {
        case 165:
            var v46 int = fi44
            _ = v46
        case -9223372036854775808:
            var v47 float64 = v1
            _ = v47
        case 0:
            var v48 float64 = ((((float64(-2.25) * v1) - 1.5) + float64(0.1)) + v1)
            _ = v48
        default:
            delete(v4, "")
        }
        _ = fi44
        _ = v45
    }
    v49 := v4["foo bar"]
    v1 = v1
    v1 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v49
}

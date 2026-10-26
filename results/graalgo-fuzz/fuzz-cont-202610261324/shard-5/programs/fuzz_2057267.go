package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 2; fi5++ {
        switch ("foo bar" + "\n\t") {
        case "":
            v6 := GenS1{F0: 278, F1: "", F2: 255}
            _ = v6
        case "hello":
            var v7 uint8 = 127
            _ = v7
        }
        _ = fi5
    }
    if true {
        for fi8 := 0; fi8 < 6; fi8++ {
            if (0 > 0) {
                var v9 float32 = 1.5
                if ((fi8 >> ((847) & 63)) < 32767) {
                    fmt.Println(v9)
                }
                var v10 string = (("" + ("hello" + "foo bar")) + ("\n\t" + ("\n\t" + (("a" + "") + (("a" + ((("foo bar" + "\n\t") + "a") + "\n\t")) + "a")))))
                _ = v10
                _ = v9
            } else {
                v11 := []int{fi8, (fi8 ^ (fi8 * -624)), 851, (fi8 * 2147483647)}
                var v12 float32 = 1.5
                var v13 int16 = -32768
                _ = v11
                _ = v12
                _ = v13
            }
            fmt.Println(fi8)
            var v14 int64 = 9223372036854775807
            _ = fi8
            _ = v14
        }
        var v15 int32 = 2147483647
        _ = v15
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            if false {
                v17 := GenS2{F0: 100, F1: -2147483648, F2: 586}
                v18 := GenS3{F0: 32767, F1: 18446744073709551615, F2: 42}
                var v19 uint16 = 65535
                _ = v17
                _ = v18
                _ = v19
            }
            var v20 string = ("foo bar" + ("foo bar" + "a"))
            var v21 int32 = -1
            _ = fi16
            _ = v20
            _ = v21
        }
    }
    return 362
}
func genFn1() int32 {
    var v22 int32 = 2147483647
    var v23 uint8 = 2
    var v24 int32 = v22
    if (0 <= 2) {
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 int16 = -1
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := []int{(fi27 % (((fi27 | fi27)) | 1)), fi27}
            v22 += v22
            switch "\n\t" {
            case "\n\t":
                for fi29 := 0; fi29 < 17; fi29++ {
                    _ = len(v28)
                    _ = fi29
                }
            case "foo bar":
                if false {
                    v24 -= (v24 | (v22 & 0))
                    v28 = append(v28, fi27)
                    var v30 uint = 826
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        if (1000 < 42) {
                            var v32 uint8 = v23
                            var v33 uint = 0
                            _ = len(v28)
                            _ = v32
                            _ = v33
                        }
                        switch (fi31 / ((fi31) | 1)) {
                        case 4294967295:
                            if (false || true) {
                                var v34 uint16 = 42
                                v35 := GenS3{F0: 32767, F1: 255, F2: 442}
                                var v36 float32 = -2.25
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        }
                        v37 := GenS3{F0: 1000, F1: 255, F2: 2}
                        _ = fi31
                        _ = v37
                    }
                }
            default:
                v38 := GenS4{F0: 65535}
                _ = v38
            }
            _ = fi27
            _ = v28
        }
        if false {
            for fi39 := 0; fi39 < 15; fi39++ {
                fmt.Println(v24)
                if (!false) {
                    switch (("hello" + "\n\t") + ("foo bar" + ("" + "\n\t"))) {
                    case "hello":
                        switch "a" {
                        case "hello":
                            v40 := []int{-429, (fi39 & 32768)}
                            _ = v40
                        }
                    case "a":
                        v24 = -1
                    case "foo bar":
                        fmt.Println(v24, v24)
                    default:
                        v24 = 1
                    }
                    v41 := []int{fi39, fi39}
                    v23 += 0
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 5; fi42++ {
                        var v43 int32 = -307
                        v44 := GenS1{F0: 42, F1: "", F2: 701}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    for fi45 := 0; fi45 < 7; fi45++ {
                        var v46 bool = false
                        v47 := GenS5{F0: 0}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                }
                fmt.Println(v22, fi39, v24)
                _ = fi39
            }
            v23 = v23
        }
    } else {
        var v48 uint = 1
        _ = v48
    }
    var v49 bool = false
    _ = v22
    _ = v23
    _ = v24
    _ = v49
    return -906
}

func main() {
    v1 := []int{4294967295, 32768, 1, 42}
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    v4 := []string{"\n\t"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v50 := GenS4{F0: 18446744073709551615}
    var v51 float32 = 1.5
    var v52 string = "\n\t"
    var v53 uint16 = 65535
    var v54 bool = false
    v52 = v52
    delete(v3, "hello")
    v52 = v52
    switch 42 {
    case 127:
        v51 += 1e10
    default:
        v54 = false
    }
    for fi55 := 0; fi55 < 15; fi55++ {
        for fi56 := 0; fi56 < 4; fi56++ {
            fmt.Println(v53)
            if len(v4) > 0 {
                v4[0] = v52
            }
            _ = fi56
        }
        v57 := v3[v52]
        _ = fi55
        _ = v57
    }
    v58 := GenS3{F0: -595, F1: 0, F2: 255}
    if false {
        v53 += v53
        for fi59 := 0; fi59 < 11; fi59++ {
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi59
        }
        fmt.Println(v52)
    } else {
        if (true || v54) {
            var v60 float64 = 0.1
            fmt.Println(v53, v53, v54)
            _ = v60
        } else {
            if (1 == 151) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi61 := 0; fi61 < 15; fi61++ {
                    for fi62 := 0; fi62 < 6; fi62++ {
                        fmt.Println(v54, v51, v53)
                        _ = fi62
                    }
                    _ = fi61
                }
                v51 -= (v51 - 0.1)
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v51, v51, v52)
                v53 += v53
            }
        }
        v53 = (v53 << ((42) & 63))
    }
    if v54 {
        for fi63 := 0; fi63 < 17; fi63++ {
            switch 2147483647 {
            case 32768:
                v53 -= 2
            case -624:
                v51 -= 3.14159
            }
            delete(v3, "hello")
            v54 = (int8(v53) == 0)
            _ = fi63
        }
        v52 = "foo bar"
        for fi64 := 0; fi64 < 20; fi64++ {
            fmt.Println(v51)
            v53 += v53
            switch "hello" {
            case "hello":
                fmt.Println(v51)
            default:
                delete(v2, v52)
            }
            _ = fi64
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v58
}

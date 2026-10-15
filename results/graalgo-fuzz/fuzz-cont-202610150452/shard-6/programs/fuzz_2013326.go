package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint8 }
func genFn0() string {
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: (32767 > 42), F1: 864}
    fmt.Println(v7)
    var v9 int8 = -63
    _ = v7
    _ = v8
    _ = v9
    return ("a" + "foo bar")
}
func genFn1() uint8 {
    var v10 int8 = 127
    var v11 int32 = 964
    _ = v10
    _ = v11
    return 128
}

func main() {
    v1 := []string{"hello", ""}
    var v2 uint64 = 1
    var v3 float64 = ((float64(1.5) / -2.25) + -2.25)
    v4 := []int{}
    var v5 float64 = 0.0
    var v6 int64 = 9223372036854775807
    var v12 int64 = v6
    if (!true) {
        fmt.Println(v5, v3, v3)
        v6 += -540
        _ = len(v4)
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        _ = len(v4)
        v14 := GenS1{F0: (v3 >= v3), F1: 1000}
        _ = fi13
        _ = v14
    }
    switch "foo bar" {
    case "hello":
        v15 := GenS1{F0: true, F1: 0}
        _ = v15
    case "\n\t":
        var v16 uint = 0
        _ = v16
    case "":
        switch "foo bar" {
        case "":
            var v17 uint32 = 179
            _ = v17
        case "foo bar":
            for fi18 := 0; fi18 < 8; fi18++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v19 uint8 = genFn1()
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi18
                _ = v19
            }
        case "hello":
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(v3, v3, fi20)
                switch "" {
                case "a":
                    fmt.Println(v6)
                }
                _ = fi20
            }
        default:
            if false {
                v12 = v12
                v6 -= 9223372036854775807
                var v21 string = "a"
                _ = v21
            } else {
                var v22 uint16 = 32767
                _ = v22
            }
        }
    default:
        var v23 uint64 = (v2 * (v2 / ((596) | 1)))
        _ = v23
    }
    v12 = v12
    var v24 int16 = -1
    var v25 uint32 = 42
    for fi26 := 0; fi26 < 19; fi26++ {
        v1 = append(v1, "")
        if (v24 > 42) {
            if (!(!false)) {
                v6 -= v12
                switch "\n\t" {
                case "a":
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v25 -= (v25 % ((4294967295) | 1))
                        _ = fi27
                    }
                case "":
                    var v28 int32 = -212
                    _ = v28
                case "\n\t":
                    for fi29 := 0; fi29 < 15; fi29++ {
                        v5 -= v5
                        fmt.Println(v25)
                        _ = fi29
                    }
                default:
                    v30 := GenS2{F0: (fi26 / ((fi26) | 1)), F1: "hello", F2: "a"}
                    _ = v30
                }
                fmt.Println(v24)
            }
            var v31 uint = 827
            _ = v31
        }
        if false {
            if ((("hello" + "") <= ("foo bar" + ("hello" + "hello"))) && (!((^(v25)) >= 2147483648))) {
                var v32 uint64 = v2
                _ = v32
            } else {
                fmt.Println(v6, v3)
                if false {
                    v24 += (v24 & 0)
                }
            }
            if true {
                for fi33 := 0; fi33 < 2; fi33++ {
                    v34 := GenS1{F0: true, F1: 1}
                    v1 = append(v1, "")
                    var v35 string = "a"
                    if (v6 < -71) {
                        switch (fi33 % ((fi26) | 1)) {
                        case 42:
                            v2 -= v2
                        }
                        if false {
                            v3 += float64(float32(0.1))
                        }
                    } else {
                        var v36 int8 = 1
                        v4 = append(v4, fi26)
                        if true {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v37 := GenS3{F0: v36}
                            _ = v37
                        }
                        _ = v36
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            v2 = uint64(v25)
        } else {
            for fi38 := 0; fi38 < 2; fi38++ {
                v4 = append(v4, 9223372036854775807)
                var v39 int64 = (v6 ^ (v12 % ((42) | 1)))
                switch "a" {
                case "hello":
                    switch (fi38 >> ((1) & 63)) {
                    case -680:
                        var v40 uint64 = v2
                        _ = v40
                    default:
                        v41 := GenS1{F0: (!(-906 != v24)), F1: (v24 - (v24 % ((v24) | 1)))}
                        _ = v41
                    }
                }
                _ = fi38
                _ = v39
            }
            v12 += (^(v6))
        }
        _ = fi26
    }
    v42 := GenS4{F0: 1, F1: v2, F2: v6}
    v25 = v25
    v12 -= v12
    switch 42 {
    case 517:
        switch 142 {
        case 0:
            fmt.Println(v25, v6, v2)
        default:
            v6 = (v6 << ((123) & 63))
        }
    case 42:
        v43 := GenS1{F0: true, F1: (v24 % (((v24 ^ 0)) | 1))}
        _ = v43
    case 512:
        v2 = v2
    default:
        var v44 uint = 65535
        _ = v44
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v45 uint16 = 65535
    v12 -= v12
    if ((-2147483648 < 0) || false) {
        if false {
            v45 = v45
            fmt.Println(v6)
            if (uint8(v25) != 128) {
                switch 2147483647 {
                case 127:
                    var v46 uint64 = 2
                    _ = v46
                default:
                    switch 32767 {
                    case -990:
                        v45 += v45
                    case 32768:
                        v45 -= (v45 - 482)
                    case -182:
                        v24 = v24
                    }
                }
                v47 := GenS5{F0: 0}
                _ = v47
            } else {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                for fi48 := 0; fi48 < 7; fi48++ {
                    if true {
                        var v49 uint64 = v2
                        _ = v49
                    }
                    v1 = append(v1, "a")
                    _ = fi48
                }
            }
        }
    } else {
        var v50 uint16 = 42
        var v51 uint32 = 4294967295
        if true {
            v24 += int16(v2)
        }
        _ = v50
        _ = v51
    }
    var v52 int = 889
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v52
    _ = v6
}

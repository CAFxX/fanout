package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int16; F1 int8 }
func genFn0() uint16 {
    var v7 uint8 = 255
    var v8 int32 = 743
    if (("" + ("hello" + "\n\t")) < "foo bar") {
        if false {
            var v9 int16 = -249
            _ = v9
        } else {
            var v10 uint = 18446744073709551615
            _ = v10
        }
        var v11 int = 4294967295
        _ = v11
    }
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int16) bool {
    if (!false) {
        p0 = p0
    } else {
        var v12 uint = 0
        v13 := []int{}
        switch "foo bar" {
        case "\n\t":
            p0 += p0
        default:
            p0 = (p0 >> (((v12 << ((4294967295) & 63))) & 63))
        }
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: "hello", F1: true}
    v15 := []int{-9223372036854775808, 242, 128}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: "a", F1: ((!(!((3.14159 * (1.5 / 1e10)) <= 3.14159))) || true)}
        if true {
            switch "foo bar" {
            case "hello":
                v18 := GenS1{F0: "foo bar", F1: (304 != 4294967295)}
                _ = v18
            case "foo bar":
                p0 = (^(p0))
            case "a":
                if (1 != 737) {
                    v19 := GenS2{F0: (p0 % (((^(p0))) | 1))}
                    _ = v19
                } else {
                    fmt.Println(fi16, fi16)
                    var v20 int32 = 1
                    v20 -= 804
                    _ = v20
                }
            }
            var v21 int = (^(fi16))
            v21 += v21
            _ = v21
        }
        var v22 int32 = 1
        _ = fi16
        _ = v17
        _ = v22
    }
    _ = p0
    _ = v14
    _ = v15
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -1
    var v3 uint8 = 143
    v4 := []string{}
    var v5 float32 = (3.14159 / 1.5)
    v6 := make(map[string]uint32)
    fmt.Println(v3)
    var v23 int16 = 32767
    v24 := GenS2{F0: v23}
    if (65535 >= 314) {
        switch "" {
        case "hello":
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 int8 = 0
                switch (fi25 | fi25) {
                case 32767:
                    var v27 int64 = 0
                    _ = v27
                default:
                    switch "foo bar" {
                    case "":
                        v26 += v26
                    }
                }
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 int32 = -2147483648
                    switch "\n\t" {
                    case "\n\t":
                        v1[2] = ""
                    case "":
                        if (!true) {
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = (("a" + (("\n\t" + "foo bar") + "a")) + "a")
                                }
                                v5 = (v5 + 1.5)
                                v6["foo bar"] = 0
                            } else {
                                var v30 int32 = (v29 & (^(v29)))
                                _ = len(v4)
                                v30 = int32(v26)
                                _ = v30
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            delete(v6, "a")
                        }
                    }
                    if false {
                        _ = len(v4)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi28
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
        case "a":
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 uint16 = genFn0()
                var v33 int16 = v23
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 uint64 = 42
                    v4 = append(v4, (("" + "\n\t") + "hello"))
                    switch (fi34 ^ 42) {
                    case 32768:
                        v23 = v23
                    case 127:
                        v36 := GenS1{F0: "", F1: (!((v3 << ((348) & 63)) == v3))}
                        _ = v36
                    case -782:
                        v37 := GenS3{F0: int32(v3), F1: v3}
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        default:
            v3 = (v3 - (v3 & (v3 << ((255) & 63))))
        }
    }
    _ = len(v6)
    v23 += 42
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 int8 = -81
        v40 := v1[2147483647]
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 int = (v2 / ((261) | 1))
    if (246 >= 1) {
        for fi42 := 0; fi42 < 12; fi42++ {
            v23 += v23
            var v43 uint = 42
            _ = fi42
            _ = v43
        }
    } else {
        if (genFn0() <= 32767) {
            v44 := v6["foo bar"]
            _ = v44
        }
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := v1[2147483647]
            var v47 uint16 = 666
            _ = fi45
            _ = v46
            _ = v47
        }
        if (!true) {
            v48 := GenS1{F0: "", F1: (false && (9223372036854775807 < -35))}
            switch -1 {
            case -9223372036854775808:
                v49 := GenS4{F0: "\n\t", F1: v3}
                _ = v49
            case 127:
                var v50 uint64 = 4294967295
                _ = v50
            }
            var v51 int32 = 42
            _ = v48
            _ = v51
        } else {
            v52 := GenS5{F0: 4294967295, F1: 42}
            fmt.Println(v41, v5)
            _ = v52
        }
    }
    _ = len(v6)
    var v53 int8 = 106
    var v54 bool = (1 == 2147483648)
    if false {
        var v55 int64 = 9223372036854775807
        _ = v55
    } else {
        v4 = append(v4, "a")
        v56 := GenS2{F0: v23}
        v41 = v41
        _ = v56
    }
    v57 := GenS6{F0: (v23 << ((2) & 63)), F1: (v53 << ((473) & 63))}
    v53 = v53
    var v58 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v53
    _ = v54
    _ = v57
    _ = v58
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32, p1 uint64) bool {
    v5 := GenS1{F0: 0}
    fmt.Println(p1, p0)
    v6 := []int{127}
    if false {
        var v7 int64 = -9223372036854775808
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(p1, v7, p1)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (!(805 >= 340))
}
func genFn1(p0 int32, p1 bool) uint64 {
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            if (p0 != p0) {
                p0 = (p0 - p0)
            }
            var v9 int8 = -128
            var v10 float64 = float64(1e10)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 int64 = 42
    var v12 bool = true
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := []string{}
    var v4 int64 = 942
    v4 = 9223372036854775807
    switch -555 {
    case 979:
        switch "hello" {
        case "\n\t":
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int8 = 2
                v14 = -128
                fmt.Println(fi13, fi13)
                _ = fi13
                _ = v14
            }
        }
    }
    if ((true || ((!(true && true)) || true)) || true) {
        fmt.Println(v4)
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            v4 = v4
            for fi16 := 0; fi16 < 12; fi16++ {
                v4 = (v4 / (((v4 | v4)) | 1))
                v4 -= v4
                v3 = append(v3, "hello")
                _ = fi16
            }
            v4 = (v4 ^ v4)
            _ = fi15
        }
        var v17 int64 = v4
        if false {
            v17 += (^(v4))
        }
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        switch "a" {
        case "hello":
            for fi19 := 0; fi19 < 16; fi19++ {
                if true {
                    v20 := GenS2{F0: -1, F1: 1, F2: ((42 <= 1) && true)}
                    var v21 int32 = -364
                    var v22 int16 = -944
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: 319}
                v4 = v4
                _ = fi19
                _ = v23
            }
        default:
            fmt.Println(fi18, v4)
        }
        _ = fi18
    }
    v4 = 42
    v4 += 42
    if (float32(1.5) == 1e10) {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int = (^(fi24))
            _ = fi24
            _ = v25
        }
        v26 := GenS3{F0: 42}
        _ = v26
    } else {
        v27 := GenS4{F0: 1}
        _ = v27
    }
    switch "a" {
    case "":
        v28 := GenS5{F0: "\n\t", F1: v4}
        _ = v28
    case "foo bar":
        if (!true) {
            var v29 int64 = (v4 ^ (v4 >> ((0) & 63)))
            v30 := GenS3{F0: 4294967295}
            _ = v29
            _ = v30
        } else {
            switch (("foo bar" + ("hello" + "foo bar")) + "\n\t") {
            case "\n\t":
                v31 := GenS4{F0: 0}
                _ = v31
            case "a":
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS6{F0: 42, F1: (fi32 & 127)}
                    _ = fi32
                    _ = v33
                }
            case "":
                if true {
                    switch 844 {
                    case -299:
                        v34 := GenS7{F0: 65535, F1: 4294967295}
                        _ = v34
                    }
                    v35 := GenS1{F0: 273}
                    if false {
                        if (32767 == 271) {
                            if (true || false) {
                                v36 := GenS1{F0: 18446744073709551615}
                                for fi37 := 0; fi37 < 4; fi37++ {
                                    var v38 int = fi37
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                                _ = v36
                            }
                            v4 += (v4 / ((0) | 1))
                            fmt.Println(v4)
                        }
                        var v39 int8 = 0
                        v40 := GenS4{F0: 2}
                        _ = v39
                        _ = v40
                    }
                    _ = v35
                }
            }
            fmt.Println(v4)
            v4 = (v4 / (((^(v4))) | 1))
        }
    default:
        if (2 < 187) {
            var v41 int64 = v4
            v41 -= (v4 << ((255) & 63))
            switch (("" + "hello") + "foo bar") {
            case "foo bar":
                var v42 uint32 = 356
                _ = v42
            }
            _ = v41
        } else {
            var v43 string = "\n\t"
            switch 42 {
            case -1:
                var v44 int16 = 1000
                _ = v44
            default:
                for fi45 := 0; fi45 < 2; fi45++ {
                    var v46 float32 = (3.14159 - 0.1)
                    delete(v1, "\n\t")
                    v47 := GenS7{F0: 2, F1: genFn1(42, (0 >= 18446744073709551615))}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}

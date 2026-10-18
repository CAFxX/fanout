package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int }
func genFn0(p0 int) int8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{}
        p0 -= (p0 << ((255) & 63))
        var v6 uint = 1
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int16 = -747
    p0 -= p0
    var v8 uint = 42
    _ = p0
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]int)
    var v3 string = "hello"
    v3 = (v3 + (v3 + v3))
    v3 = ("foo bar" + v3)
    v9 := GenS1{F0: v3}
    switch "a" {
    case "a":
        if (!false) {
            v10 := []int{-9223372036854775808, 9223372036854775807}
            v1 = -9223372036854775808
            switch 128 {
            case 128:
                v1 = v1
            case 9223372036854775807:
                _ = len(v2)
            case 310:
                var v11 uint64 = 0
                _ = v11
            default:
                v12 := GenS2{F0: 255, F1: -1}
                _ = v12
            }
            _ = v10
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 uint64 = 2
                v1 += v1
                switch v3 {
                case "":
                    if (fi13 > fi13) {
                        _ = len(v2)
                        v1 = 908
                    }
                case "\n\t":
                    var v15 bool = false
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            fmt.Println(v1, v1)
            v1 += v1
        }
    case "hello":
        v1 -= (v1 * 482)
    }
    v1 = v1
    for fi16 := 0; fi16 < 8; fi16++ {
        switch "hello" {
        case "foo bar":
            switch (fi16 + fi16) {
            case -599:
                var v17 float64 = (-2.25 - 1e10)
                _ = v17
            case 32768:
                v3 = "a"
            case 32767:
                v3 = ((v3 + "foo bar") + (v3 + ("\n\t" + v3)))
            }
        case "":
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS2{F0: 2, F1: -618}
                if (65535 < 255) {
                    fmt.Println(fi16, v1)
                    if false {
                        v1 += (v1 | 9223372036854775807)
                        var v20 int16 = 264
                        v20 += (v20 % ((v20) | 1))
                        v3 = v3
                        _ = v20
                    } else {
                        v1 = v1
                    }
                } else {
                    fmt.Println(fi18, v3, v3)
                }
                _ = fi18
                _ = v19
            }
        case "\n\t":
            fmt.Println(v3, v1, v1)
        }
        v21 := []int{fi16, -9223372036854775808}
        _ = fi16
        _ = v21
    }
    v22 := GenS3{F0: -32768, F1: (v1 & v1)}
    v23 := GenS4{F0: 32767, F1: true, F2: (v1 + (v1 % ((-1) | 1)))}
    v24 := GenS5{F0: -10, F1: 1}
    var v25 int16 = -1
    _ = len(v2)
    if true {
        switch v3 {
        case "foo bar":
            for fi26 := 0; fi26 < 13; fi26++ {
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 float32 = ((-2.25 - 1.5) / 1.0)
                    fmt.Println(v1, fi26, fi27)
                    v29 := GenS2{F0: 18446744073709551615, F1: 1}
                    if false {
                        if false {
                            if (false && (!false)) {
                                var v30 float64 = 0.1
                                var v31 bool = true
                                _ = v30
                                _ = v31
                            } else {
                                v3 = v3
                            }
                        }
                        v32 := GenS6{F0: fi26}
                        var v33 int64 = v1
                        _ = v32
                        _ = v33
                    }
                    v1 = (v1 / ((int64(v25)) | 1))
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                var v34 uint = 2
                if true {
                    var v35 uint = (^(v34))
                    fmt.Println(v34)
                    for fi36 := 0; fi36 < 6; fi36++ {
                        var v37 uint32 = 2147483647
                        v25 = v25
                        v25 += 0
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
                _ = fi26
                _ = v34
            }
        case "\n\t":
            v38 := GenS2{F0: 0, F1: 0}
            _ = v38
        case "":
            if (965 < 42) {
                var v39 uint32 = 443
                if true {
                    v39 = 933
                    if false {
                        v2[(v3 + "a")] = 4294967295
                        if (true || ((v1 > (v1 >> ((255) & 63))) && true)) {
                            switch (v3 + v3) {
                            case "foo bar":
                                var v40 string = (v3 + v3)
                                _ = v40
                            default:
                                v1 = v1
                            }
                            v39 = v39
                            if true {
                                v39 = (^(v39))
                            }
                        } else {
                            var v41 int32 = -2147483648
                            if (1 <= 42) {
                                var v42 bool = (!false)
                                var v43 int = 32767
                                _ = v42
                                _ = v43
                            }
                            _ = v41
                        }
                        fmt.Println(v3)
                    }
                    if false {
                        v25 = (v25 / (((v25 & (v25 & v25))) | 1))
                        v44 := GenS2{F0: 167, F1: 42}
                        v1 = (v1 + v1)
                        _ = v44
                    }
                } else {
                    for fi45 := 0; fi45 < 13; fi45++ {
                        var v46 bool = false
                        var v47 int = fi45
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    v48 := GenS6{F0: 128}
                    _ = v48
                }
                v39 = 406
                _ = v39
            }
        }
        fmt.Println(v3, v25, v1)
    }
    if false {
        var v49 uint8 = 42
        _ = v49
    } else {
        var v50 uint8 = 128
        _ = v50
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v9
}

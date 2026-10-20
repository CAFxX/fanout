package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 675
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    if ((1 == 1) || (false && (32767 >= 65535))) {
        fmt.Println(v2)
        v5 := []int{4294967295}
        if true {
            if len(v5) > 0 {
                v5[0] = -1
            }
        } else {
            if true {
                for fi6 := 0; fi6 < 2; fi6++ {
                    switch (-(fi6)) {
                    case -689:
                        v5 = append(v5, (fi6 + (fi6 + fi6)))
                    case 4294967295:
                        v2 -= v2
                    case 2147483647:
                        var v7 uint = 0
                        _ = v7
                    }
                    _ = fi6
                }
                if false {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        switch ("foo bar" + ("\n\t" + ("foo bar" + "hello"))) {
                        case "\n\t":
                            if true {
                                v9 := GenS1{F0: false, F1: -128}
                                var v10 int8 = -115
                                switch (-(fi8)) {
                                case 127:
                                    var v11 bool = true
                                    _ = v11
                                case 32767:
                                    fmt.Println(v2, v10, v10)
                                }
                                _ = v10
                                _ = v9
                            } else {
                                _ = len(v4)
                                v12 := GenS1{F0: (2 >= 276), F1: -1}
                                v2 = -741
                                _ = v12
                            }
                        case "a":
                            for fi13 := 0; fi13 < 8; fi13++ {
                                var v14 uint32 = 507
                                _ = fi13
                                _ = v14
                            }
                        default:
                            fmt.Println(fi8, v2)
                        }
                        var v15 uint32 = 19
                        _ = fi8
                        _ = v15
                    }
                    if true {
                        var v16 int16 = -851
                        var v17 float32 = -2.25
                        _ = v16
                        _ = v17
                    }
                } else {
                    v2 = (v2 ^ (v2 % ((v2) | 1)))
                }
            } else {
                v2 = v2
            }
        }
        _ = v5
    } else {
        var v18 uint32 = 0
        _ = v18
    }
    v19 := GenS2{F0: -60, F1: 513}
    v2 = -1
    v20 := GenS3{F0: 421}
    if (127 < 42) {
        delete(v3, "a")
    }
    fmt.Println(v2)
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 bool = false
        _ = fi21
        _ = v22
    }
    var v23 int16 = 1000
    for fi24 := 0; fi24 < 20; fi24++ {
        if ((-(fi24)) <= (^(fi24))) {
            v2 = 42
        }
        if true {
            v23 = (-(v23))
            for fi25 := 0; fi25 < 13; fi25++ {
                switch "hello" {
                case "\n\t":
                    if ((-128 >= 0) && false) {
                        switch "\n\t" {
                        case "hello":
                            v26 := GenS3{F0: 42}
                            _ = v26
                        case "":
                            var v27 int32 = 128
                            _ = v27
                        case "foo bar":
                            v2 += (v2 / (((v2 >> ((918) & 63))) | 1))
                        }
                        var v28 bool = false
                        _ = v28
                    } else {
                        v23 = (-(v23))
                    }
                case "":
                    v2 = (v2 | (^(v2)))
                }
                delete(v1, 737)
                _ = fi25
            }
        } else {
            v29 := GenS2{F0: fi24, F1: 1}
            _ = v29
        }
        _ = fi24
    }
    var v30 uint = 763
    v2 -= (v2 * v2)
    var v31 uint = (v30 % (((v30 - (v30 / ((v30) | 1)))) | 1))
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 int32 = -959
        switch (-(fi32)) {
        case 4294967295:
            v3[""] = (fi32 >> ((695) & 63))
        case -1:
            fmt.Println(v30, v30, v31)
        case 127:
            v31 -= v31
        default:
            fmt.Println(v2, v23, fi32)
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v2)
    v30 += (v31 | v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}

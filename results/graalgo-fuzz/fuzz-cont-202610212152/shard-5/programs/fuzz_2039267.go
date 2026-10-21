package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 2
    v3 := []uint32{1, 65535}
    if (!((42 <= 255) && (42 == -542))) {
        fmt.Println(v2)
    }
    if (!true) {
        fmt.Println(v2)
        v4 := GenS1{F0: 768}
        _ = v4
    }
    var v5 string = "\n\t"
    v5 = "foo bar"
    v2 -= (^(v2))
    _ = len(v1)
    var v6 uint = (v2 - (v2 ^ (v2 >> (((v2 * (v2 << ((v2) & 63)))) & 63))))
    if (-1 < 1) {
        var v7 int64 = -9223372036854775808
        _ = v7
    } else {
        fmt.Println(v5, v6)
    }
    var v8 uint8 = 135
    if false {
        switch 819 {
        case 0:
            var v9 uint64 = 2
            _ = v9
        case -283:
            for fi10 := 0; fi10 < 16; fi10++ {
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 int16 = -742
                    var v13 int16 = v12
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS2{F0: 31}
            v17 := GenS1{F0: 18446744073709551615}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 uint64 = 859
        if true {
            v19 := GenS3{F0: "\n\t", F1: 102}
            v20 := GenS1{F0: (v18 % (((v18 << ((v6) & 63))) | 1))}
            var v21 bool = (!true)
            _ = v19
            _ = v20
            _ = v21
        } else {
            var v22 string = ("foo bar" + v5)
            var v23 string = "a"
            _ = v22
            _ = v23
        }
        _ = fi14
        _ = v18
    }
    var v24 int8 = 1
    v8 = v8
    if ((((((("\n\t" + (v5 + v5)) + v5) + "foo bar") + v5) + v5) + ((v5 + "hello") + v5)) >= (v5 + v5)) {
        if false {
            switch 221 {
            case 32767:
                if (32767 == 255) {
                    if false {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v26 := v1[(v5 + v5)]
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v5, fi25)
                            default:
                                for fi27 := 0; fi27 < 17; fi27++ {
                                    v1[((v26 + v5) + v26)] = (v5 + (((v5 + "foo bar") + ((v26 + v26) + ((("\n\t" + v26) + v5) + v26))) + v26))
                                    var v28 uint8 = v8
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            v5 = "\n\t"
                            _ = fi25
                            _ = v26
                        }
                        switch -1 {
                        case 32768:
                            v3 = append(v3, 13)
                        case -284:
                            var v29 int8 = v24
                            _ = v29
                        }
                        fmt.Println(v5)
                    } else {
                        var v30 uint = (v2 + 4294967295)
                        switch 0 {
                        case 127:
                            v6 = (v6 * (v30 - (^(v2))))
                        }
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 9; fi31++ {
                        if ((-2.25 != -2.25) && (450 > 2)) {
                            v32 := GenS4{F0: 0, F1: (fi31 ^ 42)}
                            fmt.Println(fi31, v6)
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 19; fi33++ {
                                v24 = v24
                                var v34 int16 = 0
                                _ = fi33
                                _ = v34
                            }
                            _ = len(v3)
                        }
                        _ = fi31
                    }
                } else {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 float32 = 3.14159
                        for fi37 := 0; fi37 < 20; fi37++ {
                            v3 = append(v3, 2)
                            v8 = v8
                            if (!true) {
                                v5 = "foo bar"
                                switch ("" + (v5 + v5)) {
                                case "a":
                                    switch fi37 {
                                    case -1:
                                        var v38 uint16 = 2
                                        _ = v38
                                    case 127:
                                        switch -9223372036854775808 {
                                        case -670:
                                            if false {
                                                v39 := GenS1{F0: 0}
                                                fmt.Println(fi35, v5)
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v6 -= v6
                                                }
                                                v40 := GenS1{F0: 4294967295}
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    }
                                case "hello":
                                    _ = len(v3)
                                case "\n\t":
                                    v3 = append(v3, 428)
                                    switch v5 {
                                    case "":
                                        var v41 uint32 = 1
                                        _ = v41
                                    default:
                                        v6 += 255
                                    }
                                    if true {
                                        v42 := GenS5{F0: 0, F1: (fi37 & -308)}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v42
                                    }
                                }
                                v2 = 18446744073709551615
                            }
                            _ = fi37
                        }
                        fmt.Println(v8)
                        _ = fi35
                        _ = v36
                    }
                }
            case -1:
                v2 -= (v6 * v6)
            }
            if false {
                if false {
                    var v43 uint16 = 255
                    v1[v5] = "hello"
                    _ = v43
                } else {
                    v1[(v5 + (v5 + v5))] = ("a" + v5)
                    var v44 uint32 = 0
                    _ = v44
                }
            } else {
                fmt.Println(v8, v5)
                var v45 uint64 = 4294967295
                fmt.Println(v24)
                _ = v45
            }
            var v46 uint64 = 42
            _ = v46
        } else {
            delete(v1, "a")
        }
    } else {
        if (42 != 11) {
            v5 = (((v5 + v5) + (v5 + ("" + ("\n\t" + ("hello" + v5))))) + v5)
            v24 -= 92
        } else {
            fmt.Println(v8, v2, v24)
        }
        v6 += v2
    }
    v5 = (("foo bar" + "\n\t") + "\n\t")
    var v47 bool = true
    fmt.Println(v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v47
    _ = v5
    _ = v6
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 int; F1 uint; F2 int }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    var v1 int16 = -32768
    v2 := []uint32{65535}
    var v3 int32 = -509
    v1 = v1
    v1 -= v1
    v2 = append(v2, 729)
    for fi4 := 0; fi4 < 7; fi4++ {
        switch fi4 {
        case 2147483647:
            v5 := GenS1{F0: 738, F1: ("" + ""), F2: 100}
            _ = v5
        default:
            for fi6 := 0; fi6 < 12; fi6++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    if (((!(!true)) && true) || (!false)) {
                        if false {
                            v7 := GenS1{F0: 2, F1: ("foo bar" + (("\n\t" + "a") + ("" + ("a" + ("\n\t" + ("a" + "hello")))))), F2: 100}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v7
                        }
                        var v8 int = fi6
                        _ = v8
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                        var v9 uint32 = 42
                        _ = v9
                    }
                }
                v2 = append(v2, 0)
                _ = fi6
            }
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            switch 127 {
            case 9223372036854775807:
                v3 = (v3 % ((42) | 1))
            case -9223372036854775808:
                if false {
                    if (1 < 42) {
                        switch "foo bar" {
                        case "\n\t":
                            var v11 int32 = (-(v3))
                            _ = v11
                        case "":
                            v1 -= -901
                        default:
                            v12 := GenS2{F0: (-(fi10)), F1: 65535, F2: (fi10 / (((fi10 >> ((1) & 63))) | 1))}
                            _ = v12
                        }
                    }
                    var v13 uint64 = 1
                    var v14 int8 = 42
                    v15 := GenS1{F0: 0, F1: "hello", F2: (v14 << ((1) & 63))}
                    v14 += (v14 * 21)
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                if (0 <= (v3 % ((-2147483648) | 1))) {
                    if false {
                        if true {
                            v16 := GenS1{F0: 930, F1: "", F2: 15}
                            _ = v16
                        } else {
                            fmt.Println(fi4, v1, v3)
                        }
                        fmt.Println(v3, v1, fi4)
                        v17 := GenS3{F0: false, F1: false, F2: (true && true)}
                        switch fi10 {
                        case 9223372036854775807:
                            v1 = (v1 >> ((4294967295) & 63))
                        case -1:
                            fmt.Println(fi10, fi10)
                        case 42:
                            var v18 int = fi4
                            _ = v18
                        }
                        _ = v17
                    } else {
                        v19 := GenS4{F0: ("a" + ("hello" + "\n\t")), F1: -562}
                        _ = v19
                    }
                    v20 := GenS2{F0: (fi4 - -640), F1: 0, F2: fi10}
                    _ = v20
                } else {
                    fmt.Println(fi4, v1)
                }
            }
            v21 := GenS1{F0: 2147483647, F1: "hello", F2: 1}
            v3 += v3
            _ = fi10
            _ = v21
        }
        _ = fi4
    }
    var v22 int32 = v3
    v1 = (v1 ^ 1000)
    var v23 uint64 = 42
    for fi24 := 0; fi24 < 13; fi24++ {
        _ = len(v2)
        _ = fi24
    }
    var v25 uint8 = 194
    if (0 > 0) {
        v2 = append(v2, 1)
        v26 := GenS5{F0: -1, F1: -332}
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 int64 = 0
            var v29 float64 = 0.0
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        if false {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi30
    }
    var v31 int8 = 86
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v31
}

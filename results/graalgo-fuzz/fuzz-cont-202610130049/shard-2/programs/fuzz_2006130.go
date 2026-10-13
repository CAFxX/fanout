package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := []string{"\n\t"}
    v4 := []string{"foo bar", "", "a"}
    var v5 int8 = -123
    var v6 uint8 = 142
    switch 2147483647 {
    case 32767:
        var v7 int16 = -32768
        _ = v7
    default:
        if len(v3) > 0 {
            v3[0] = ("" + ((("a" + "a") + "foo bar") + "foo bar"))
        }
    }
    v6 -= 12
    if true {
        v6 = (v6 * 200)
        if true {
            switch -9223372036854775808 {
            case -1:
                var v8 uint = 65535
                _ = v8
            default:
                if true {
                    var v9 uint32 = 425
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                v5 -= -1
                if ((^(v6)) <= 96) {
                    v11 := GenS1{F0: 362}
                    var v12 uint64 = 962
                    _ = v11
                    _ = v12
                }
                fmt.Println(v5)
                _ = fi10
            }
            if false {
                switch 32767 {
                case 485:
                    var v13 uint = 4294967295
                    _ = v13
                case 131:
                    var v14 uint64 = 0
                    _ = v14
                case -1:
                    if (true && ((false || false) && false)) {
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v6 -= v6
                            v5 += 1
                            _ = fi15
                        }
                        v3 = append(v3, (("" + "foo bar") + ("foo bar" + "")))
                    }
                default:
                    fmt.Println(v6, v5)
                }
            }
        } else {
            switch int(v5) {
            case 2147483647:
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS1{F0: -2147483648}
                    switch ("\n\t" + ("a" + "foo bar")) {
                    case "a":
                        var v18 float64 = 0.1
                        _ = v18
                    case "\n\t":
                        v19 := GenS2{F0: (v6 << ((255) & 63)), F1: 42}
                        _ = v19
                    case "":
                        for fi20 := 0; fi20 < 13; fi20++ {
                            if (((^(v6)) == 128) && ((v6 + 14) == (v6 / (((^(v6))) | 1)))) {
                                v5 = v5
                                var v21 uint16 = 360
                                _ = v21
                            }
                            _ = fi20
                        }
                    default:
                        if false {
                            _ = len(v3)
                        }
                    }
                    _ = fi16
                    _ = v17
                }
            case -2147483648:
                switch -883 {
                case -9223372036854775808:
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v4 = append(v4, "")
                        delete(v1, 507)
                        var v23 bool = false
                        _ = fi22
                        _ = v23
                    }
                default:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        var v25 int64 = -817
                        var v26 uint16 = 197
                        delete(v1, 1)
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
            case 32767:
                fmt.Println(v5)
            default:
                var v27 int64 = -9223372036854775808
                _ = v27
            }
            fmt.Println(v5, v5)
        }
    }
    v28 := GenS3{F0: false, F1: -2147483648}
    fmt.Println(v6, v5)
    v6 -= (v6 * v6)
    delete(v2, 4294967295)
    switch 2147483647 {
    case 254:
        v5 -= 0
    }
    if false {
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 uint16 = 1
                _ = fi30
                _ = v31
            }
            v32 := GenS4{F0: fi29, F1: (v6 + 255), F2: 872}
            _ = fi29
            _ = v32
        }
        var v33 uint16 = 36
        _ = v33
    } else {
        v3 = append(v3, "a")
        var v34 float32 = 0.1
        _ = len(v2)
        _ = v34
    }
    v5 = (v5 << ((42) & 63))
    if true {
        fmt.Println(v6, v6)
    }
    var v35 uint16 = 65535
    fmt.Println(v6, v35, v35)
    v6 += (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}

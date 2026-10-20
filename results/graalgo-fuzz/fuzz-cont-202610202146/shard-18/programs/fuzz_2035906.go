package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 int32) string {
    for fi6 := 0; fi6 < 20; fi6++ {
        switch fi6 {
        case 4294967295:
            v7 := GenS1{F0: -1}
            _ = v7
        case -9223372036854775808:
            var v8 int64 = 648
            _ = v8
        case -2147483648:
            v9 := GenS1{F0: -32768}
            _ = v9
        default:
            switch (fi6 ^ fi6) {
            case -1:
                var v10 int8 = -128
                _ = v10
            case 975:
                fmt.Println(fi6)
            case 128:
                p0 -= 0
            }
        }
        if true {
            switch (fi6 | (^(fi6))) {
            case -317:
                v11 := GenS1{F0: 32767}
                _ = v11
            case 32767:
                p0 += (p0 / ((-176) | 1))
            case 660:
                p0 += -1
            }
        }
        _ = fi6
    }
    v12 := GenS2{F0: 507, F1: 233}
    for fi13 := 0; fi13 < 4; fi13++ {
        p0 = 42
        p0 = -564
        p0 = p0
        _ = fi13
    }
    var v14 int16 = 32767
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 float32 = -2.25
        v17 := []int{fi15}
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v12
    _ = v14
    return "hello"
}
func genFn1(p0 int64, p1 uint16, p2 int64) int64 {
    p2 = 0
    v18 := []int{289, -1}
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return p2
}

func main() {
    v1 := []string{"a"}
    var v2 uint8 = 219
    v3 := []string{}
    var v4 uint = 0
    var v5 bool = (1 >= 2147483647)
    for fi19 := 0; fi19 < 13; fi19++ {
        if true {
            fmt.Println(v4)
            var v20 int8 = 1
            _ = v20
        }
        var v21 uint8 = v2
        v22 := GenS1{F0: 42}
        _ = fi19
        _ = v21
        _ = v22
    }
    switch "" {
    case "\n\t":
        var v23 uint32 = uint32(v2)
        _ = v23
    case "a":
        var v24 uint8 = (^(v2))
        _ = v24
    case "":
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 int8 = 127
                _ = fi26
                _ = v27
            }
            var v28 int = (^(fi25))
            _ = fi25
            _ = v28
        }
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if v5 {
        _ = len(v3)
    }
    var v29 uint64 = 225
    v30 := GenS1{F0: -1}
    if v5 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v4 = (v4 * v4)
    var v31 uint32 = 99
    if (!true) {
        fmt.Println(v2)
    } else {
        var v32 int16 = -1
        _ = v32
    }
    switch 358 {
    case 32768:
        var v33 uint32 = v31
        _ = v33
    case -761:
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 int32 = 816
            switch fi34 {
            case 4294967295:
                if v5 {
                    if v5 {
                        for fi36 := 0; fi36 < 9; fi36++ {
                            fmt.Println(v2, v35)
                            _ = fi36
                        }
                        switch 127 {
                        case 4294967295:
                            for fi37 := 0; fi37 < 10; fi37++ {
                                if len(v1) > 0 {
                                    v1[0] = (("" + "hello") + ((("\n\t" + "") + ("hello" + ("\n\t" + ""))) + "\n\t"))
                                }
                                v38 := GenS2{F0: (^(v29)), F1: (fi37 ^ -264)}
                                var v39 int16 = -564
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        case -9223372036854775808:
                            fmt.Println(v2)
                        case -244:
                            v2 += v2
                        default:
                            fmt.Println(v5, v35)
                        }
                        switch "hello" {
                        case "":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    } else {
                        if v5 {
                            v5 = ((v31 * 54) > v31)
                            v5 = false
                        }
                        if v5 {
                            v40 := GenS3{F0: ("hello" + ("" + (("hello" + ("foo bar" + "a")) + "a"))), F1: ("foo bar" + "")}
                            _ = v40
                        }
                    }
                    var v41 uint8 = v2
                    _ = v41
                }
            default:
                fmt.Println(v29)
            }
            _ = fi34
            _ = v35
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

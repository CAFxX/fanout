package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int64 }
func genFn0() uint16 {
    var v7 uint16 = 255
    v7 = (v7 | v7)
    _ = v7
    return (^(v7))
}

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 int16 = 42
    var v4 int64 = 1
    var v5 uint16 = 2
    v6 := []string{("" + ("\n\t" + ("" + "hello"))), ("\n\t" + "\n\t"), (("foo bar" + ("a" + "foo bar")) + "a")}
    v8 := GenS1{F0: ((v4 / ((v4) | 1)) <= (-(v4))), F1: 42, F2: (v5 % ((13) | 1))}
    switch 970 {
    case 128:
        var v9 int64 = 787
        _ = v9
    case 127:
        for fi10 := 0; fi10 < 13; fi10++ {
            v3 -= v3
            var v11 uint32 = 2147483648
            _ = fi10
            _ = v11
        }
    default:
        var v12 int32 = 946
        _ = v12
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v3 -= (-(v3))
        _ = fi13
    }
    v3 -= 42
    v14 := GenS1{F0: (!v1), F1: 65535, F2: v5}
    var v15 int64 = 367
    if (4294967295 > 378) {
        v16 := GenS1{F0: true, F1: 1, F2: (v5 >> ((0) & 63))}
        v15 -= -1
        _ = v16
    }
    var v17 int32 = 1
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS2{F0: 2, F1: ("hello" + "a")}
        v17 += (v17 ^ -2147483648)
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 float64 = 0.0
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    if v1 {
        v22 := GenS3{F0: -128, F1: 2}
        v5 += v5
        _ = v22
    } else {
        var v23 uint32 = 1
        if (!v1) {
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS4{F0: v15}
                _ = fi24
                _ = v25
            }
            v26 := GenS4{F0: v4}
            _ = v26
        } else {
            v6 = append(v6, ("\n\t" + ""))
            v4 += -1
        }
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = (("\n\t" + ("\n\t" + "\n\t")) + ("foo bar" + "\n\t"))
    }
    var v27 uint16 = 255
    for fi28 := 0; fi28 < 1; fi28++ {
        switch "\n\t" {
        case "":
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        case "hello":
            v4 = v15
        case "a":
            switch "" {
            case "\n\t":
                if true {
                    var v29 uint64 = 0
                    switch fi28 {
                    case 837:
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v4 += v4
                            _ = fi30
                        }
                    default:
                        var v31 uint32 = 65535
                        _ = v31
                    }
                    _ = v29
                } else {
                    if ((fi28 * fi28) <= fi28) {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            if (1 == 2) {
                                v3 = 32767
                            }
                            var v33 int32 = v17
                            _ = fi32
                            _ = v33
                        }
                        if ((!v1) || ("hello" == "hello")) {
                            if v1 {
                                if v1 {
                                    v1 = v1
                                    var v34 float64 = 1e10
                                    _ = v34
                                }
                                fmt.Println(v17, v27, v3)
                            }
                            v35 := GenS4{F0: v4}
                            fmt.Println(v27, v3)
                            _ = v35
                        }
                        var v36 float64 = 1.5
                        _ = v36
                    }
                    switch (("hello" + "a") + "") {
                    case "foo bar":
                        if len(v2) > 0 {
                            v2[0] = ("hello" + "foo bar")
                        }
                    case "hello":
                        switch "a" {
                        case "hello":
                            v37 := GenS2{F0: 2, F1: ""}
                            _ = v37
                        case "a":
                            var v38 uint8 = 42
                            _ = v38
                        }
                    default:
                        v1 = (!(v3 < (-(v3))))
                    }
                }
            case "":
                for fi39 := 0; fi39 < 2; fi39++ {
                    v40 := GenS3{F0: -1, F1: 42}
                    switch (fi39 & 9223372036854775807) {
                    case 2147483647:
                        v17 += v17
                    case 0:
                        var v41 uint8 = 42
                        _ = v41
                    default:
                        v3 -= 32767
                    }
                    var v42 uint16 = v5
                    _ = fi39
                    _ = v40
                    _ = v42
                }
            case "a":
                _ = len(v2)
            default:
                for fi43 := 0; fi43 < 2; fi43++ {
                    v27 = (v27 >> ((uint(v5)) & 63))
                    switch (("foo bar" + "hello") + "\n\t") {
                    case "\n\t":
                        var v44 float32 = 1e10
                        var v45 float32 = float32((0.1 * (0.1 + 0.1)))
                        v45 = 1.5
                        _ = v44
                        _ = v45
                    case "":
                        var v46 uint64 = 470
                        _ = v46
                    }
                    _ = fi43
                }
            }
        default:
            for fi47 := 0; fi47 < 7; fi47++ {
                v15 = v15
                _ = fi47
            }
        }
        if v1 {
            v48 := GenS1{F0: v1, F1: 4294967295, F2: (v5 - genFn0())}
            for fi49 := 0; fi49 < 13; fi49++ {
                var v50 float64 = ((0.1 / 1e10) - 3.14159)
                v4 += v15
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi49
                _ = v50
            }
            v51 := GenS2{F0: 255, F1: ("hello" + "hello")}
            _ = v48
            _ = v51
        }
        for fi52 := 0; fi52 < 16; fi52++ {
            fmt.Println(v15, v1, v27)
            _ = fi52
        }
        _ = fi28
    }
    for fi53 := 0; fi53 < 2; fi53++ {
        v54 := GenS2{F0: 163, F1: "foo bar"}
        v5 = (v27 | (^(v27)))
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

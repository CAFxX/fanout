package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }

func main() {
    var v1 uint8 = 1
    v2 := []string{"a", ("" + "hello")}
    v3 := []uint32{663, 783, 42, 2147483648}
    var v4 string = ""
    fmt.Println(v1, v1)
    switch v4 {
    case "a":
        fmt.Println(v1, v1)
    case "":
        for fi5 := 0; fi5 < 2; fi5++ {
            switch 42 {
            case -504:
                switch fi5 {
                case -308:
                    v4 = "hello"
                case 747:
                    if (1 > -317) {
                        if false {
                            v6 := GenS1{F0: "", F1: "a", F2: v4}
                            _ = v6
                        }
                        v7 := GenS1{F0: ("" + v4), F1: v4, F2: ((((((v4 + (v4 + "")) + v4) + "") + v4) + v4) + v4)}
                        v8 := GenS1{F0: "", F1: v4, F2: v4}
                        _ = v7
                        _ = v8
                    } else {
                        v1 -= v1
                        switch v4 {
                        case "":
                            v4 = v4
                        case "a":
                            var v9 uint8 = v1
                            _ = v9
                        case "\n\t":
                            _ = len(v3)
                        default:
                            v10 := GenS1{F0: "a", F1: (v4 + "a"), F2: v4}
                            _ = v10
                        }
                    }
                case 2147483647:
                    var v11 bool = true
                    _ = v11
                }
            case 2147483647:
                fmt.Println(fi5)
            case -2147483648:
                if (!false) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v12 := GenS1{F0: v4, F1: v4, F2: v4}
                    v1 += (v1 - v1)
                    _ = v12
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 = "foo bar"
                }
            }
            switch 706 {
            case 42:
                v1 += 33
            case -613:
                v4 = v4
            default:
                fmt.Println(fi5, v4, fi5)
            }
            _ = fi5
        }
    }
    v1 = (v1 & v1)
    fmt.Println(v4)
    v1 += 42
    v4 = v4
    fmt.Println(v4, v4)
    switch 9223372036854775807 {
    case -1:
        var v13 int16 = 0
        _ = v13
    case 32768:
        fmt.Println(v1, v1)
    }
    if (!(0 < -558)) {
        var v14 bool = true
        _ = v14
    } else {
        var v15 int16 = -32768
        var v16 float64 = (-2.25 * float64(0.0))
        _ = v15
        _ = v16
    }
    var v17 uint32 = 42
    var v18 int32 = 387
    v1 += 0
    if ((565 < -9223372036854775808) || ((v17 >= v17) && false)) {
        var v19 int64 = 697
        _ = v19
    } else {
        for fi20 := 0; fi20 < 12; fi20++ {
            switch (fi20 - fi20) {
            case 42:
                v21 := GenS1{F0: v4, F1: "\n\t", F2: "hello"}
                _ = v21
            case 9223372036854775807:
                var v22 uint64 = uint64(v17)
                _ = v22
            case 0:
                var v23 uint64 = 42
                _ = v23
            }
            _ = fi20
        }
        v24 := GenS1{F0: "\n\t", F1: v4, F2: "hello"}
        v18 -= 1
        _ = v24
    }
    switch ("hello" + ("foo bar" + (v4 + "a"))) {
    case "hello":
        for fi25 := 0; fi25 < 14; fi25++ {
            if (false || false) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v26 uint = 4294967295
                v17 = (v17 & 1)
                v18 = v18
                _ = v26
            }
            var v27 bool = (!true)
            v3 = append(v3, (^(v17)))
            _ = fi25
            _ = v27
        }
    case "foo bar":
        if (128 > 742) {
            v17 = (^(v17))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = ""
        }
    default:
        var v28 bool = ((v17 > 876) || true)
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

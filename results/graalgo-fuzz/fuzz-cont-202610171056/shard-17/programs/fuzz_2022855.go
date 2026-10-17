package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint16 }

func main() {
    var v1 uint8 = 42
    v2 := []int{1, -9223372036854775808, 128, 2147483647}
    v3 := []int{1, 777, 32768, 32767}
    var v4 int32 = -1
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint32 = 368
        if true {
            _ = len(v3)
            v7 := GenS1{F0: 81, F1: v6, F2: 856}
            v8 := GenS2{F0: "a", F1: (fi5 / ((9223372036854775807) | 1))}
            _ = v7
            _ = v8
        } else {
            v2 = append(v2, (fi5 * fi5))
            v6 -= 647
        }
        var v9 float64 = -2.25
        _ = fi5
        _ = v6
        _ = v9
    }
    if (true || true) {
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            if (true || (!false)) {
                var v11 bool = (944 > 964)
                if (2 > 2) {
                    var v12 int32 = v4
                    var v13 uint64 = 2
                    v14 := GenS1{F0: (v13 * v13), F1: 2, F2: 753}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = v11
            } else {
                var v15 uint32 = 2
                switch fi10 {
                case -720:
                    if (42 >= 42) {
                        v4 -= v4
                        switch (fi10 << ((uint(v1)) & 63)) {
                        case 127:
                            v1 = (v1 ^ 19)
                        case 4294967295:
                            for fi16 := 0; fi16 < 9; fi16++ {
                                v4 = (v4 | (v4 << ((255) & 63)))
                                _ = fi16
                            }
                        }
                        var v17 int64 = 42
                        _ = v17
                    } else {
                        var v18 float64 = float64(-2.25)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 uint16 = 255
                            if ((0.1 <= v18) || (("a" + "hello") > "hello")) {
                                var v21 uint8 = v1
                                fmt.Println(v1, v4, v1)
                                var v22 int64 = 0
                                _ = v21
                                _ = v22
                            }
                            switch 759 {
                            case -235:
                                v2 = append(v2, fi10)
                            case -2147483648:
                                var v23 uint64 = 1
                                _ = v23
                            case 2147483647:
                                v4 = v4
                            default:
                                var v24 int = (fi19 >> ((65535) & 63))
                                _ = v24
                            }
                            var v25 int = (fi10 - fi19)
                            var v26 int64 = 636
                            _ = fi19
                            _ = v20
                            _ = v25
                            _ = v26
                        }
                        _ = v18
                    }
                default:
                    v15 = (v15 | v15)
                }
                _ = v15
            }
            _ = fi10
        }
        v27 := GenS2{F0: "a", F1: 127}
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch ("foo bar" + "") {
        case "":
            switch ("" + ("\n\t" + (("a" + "\n\t") + ("hello" + (("foo bar" + ("foo bar" + "foo bar")) + "\n\t"))))) {
            case "foo bar":
                v29 := GenS3{F0: 2147483648}
                _ = v29
            case "":
                if ((false && false) || ((9223372036854775807 == -735) && false)) {
                    var v30 int64 = 1
                    if false {
                        var v31 uint32 = 73
                        _ = v31
                    }
                    _ = v30
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        case "a":
            v1 -= (v1 - v1)
        case "hello":
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(v1, v4)
                _ = fi32
            }
        }
        _ = fi28
    }
    var v33 uint32 = 0
    v1 = uint8(v33)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v33 -= v33
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 2; fi35++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v36 := GenS4{F0: -290, F1: 2, F2: 32767}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v37 int64 = -9223372036854775808
    if (v4 >= (v4 << ((18446744073709551615) & 63))) {
        v37 -= (^(v37))
    }
    var v38 bool = true
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v38 {
        fmt.Println(v33, v4)
    }
    v1 -= (v1 - (v1 / ((107) | 1)))
    if ((1e10 <= 0.1) || (v38 && (v4 > (v4 / ((int32(v37)) | 1))))) {
        fmt.Println(v38, v37)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
}

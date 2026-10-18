package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32 }

func main() {
    var v1 int32 = 0
    var v2 uint32 = 726
    v3 := []int{-2147483648}
    var v4 string = ("" + "\n\t")
    var v5 uint64 = 84
    v6 := GenS1{F0: 127, F1: 728, F2: -1}
    v4 = (v4 + (v4 + v4))
    v2 += (v2 - (v2 & v2))
    var v7 uint64 = v5
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(fi8, v1)
                switch v4 {
                case "":
                    if ((v5 & v5) <= (v5 ^ v5)) {
                        var v10 int = (-(fi8))
                        _ = v10
                    } else {
                        v1 -= int32(fi8)
                        v5 -= (^(v7))
                        var v11 int8 = 1
                        _ = v11
                    }
                case "foo bar":
                    if true {
                        switch (v4 + v4) {
                        case "a":
                            switch (v4 + ("hello" + (v4 + (v4 + v4)))) {
                            case "\n\t":
                                switch fi9 {
                                case 42:
                                    v12 := GenS1{F0: 1, F1: (v2 & v2), F2: fi8}
                                    _ = v12
                                default:
                                    var v13 string = (("" + "foo bar") + v4)
                                    _ = v13
                                }
                            default:
                                var v14 uint8 = 84
                                _ = v14
                            }
                            v3 = append(v3, fi8)
                        case "":
                            if false {
                                v15 := GenS1{F0: -106, F1: v2, F2: fi8}
                                switch "a" {
                                case "\n\t":
                                    var v16 string = v4
                                    _ = v16
                                default:
                                    fmt.Println(v4)
                                }
                                _ = v15
                            } else {
                                v4 = (v4 + "hello")
                                v7 = (v7 - (v7 & (v7 ^ (v7 / ((v5) | 1)))))
                            }
                        default:
                            var v17 uint64 = (v7 | v5)
                            _ = v17
                        }
                        if true {
                            v1 += v1
                        } else {
                            if true {
                                switch fi8 {
                                case 1:
                                    fmt.Println(v5, v7, v4)
                                case -2147483648:
                                    v1 += (v1 % ((v1) | 1))
                                case -983:
                                    if ((v2 >> ((611) & 63)) > 418) {
                                        v1 = 174
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v18 uint64 = (v5 - (v5 & 1))
                                        v4 = (v4 + "\n\t")
                                        _ = v18
                                    }
                                }
                                switch v4 {
                                case "foo bar":
                                    var v19 uint32 = 9
                                    _ = v19
                                default:
                                    v7 = (v5 >> ((1) & 63))
                                }
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            var v20 uint64 = v5
                            if len(v3) > 0 {
                                v3[0] = 32767
                            }
                            _ = v20
                        }
                        if false {
                            var v21 int16 = 42
                            v5 = 18446744073709551615
                            if (1 <= 255) {
                                if false {
                                    v22 := GenS1{F0: -46, F1: 2147483647, F2: 42}
                                    _ = v22
                                } else {
                                    if ((1 == 1) || ((187 < 0) && (-105 >= 100))) {
                                        var v23 string = v4
                                        v23 = v4
                                        _ = v23
                                    }
                                }
                                v24 := GenS2{F0: (fi9 ^ (fi9 * fi8)), F1: 32767}
                                _ = v24
                            }
                            _ = v21
                        }
                    }
                default:
                    v25 := GenS3{F0: 255, F1: (fi8 - fi8), F2: (((v4 + (v4 + "")) + "foo bar") + (((("a" + v4) + (v4 + "")) + v4) + "foo bar"))}
                    _ = v25
                }
                v7 = v7
                _ = fi9
            }
            var v26 uint64 = 749
            var v27 int8 = -128
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 14; fi28++ {
                fmt.Println(fi8, v1)
                if (-446 != -1) {
                    var v29 int = int(v1)
                    _ = v29
                } else {
                    var v30 int32 = v1
                    v3 = append(v3, (fi8 / ((fi28) | 1)))
                    _ = v30
                }
                var v31 string = v4
                _ = fi28
                _ = v31
            }
            v5 += v5
            var v32 bool = false
            _ = v32
        }
        v33 := GenS1{F0: 26, F1: v2, F2: (fi8 + (fi8 << ((42) & 63)))}
        _ = fi8
        _ = v33
    }
    _ = len(v3)
    for fi34 := 0; fi34 < 8; fi34++ {
        v7 -= 1
        v3 = append(v3, fi34)
        _ = fi34
    }
    if (((v2 & (v2 | v2)) > v2) && (-128 < -1)) {
        var v35 uint16 = 32767
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        switch ("" + v4) {
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "foo bar":
            if true {
                var v36 uint64 = 1
                _ = v36
            } else {
                for fi37 := 0; fi37 < 12; fi37++ {
                    fmt.Println(v2, v35, v35)
                    v4 = v4
                    _ = fi37
                }
                if (true && (((1.5 * (((-2.25 - -2.25) / 3.14159) - 3.14159)) == 0.0) && true)) {
                    v2 = 4294967295
                    var v38 int16 = 1000
                    v35 = v35
                    _ = v38
                }
                switch "\n\t" {
                case "hello":
                    for fi39 := 0; fi39 < 14; fi39++ {
                        switch v4 {
                        case "foo bar":
                            _ = len(v3)
                        }
                        var v40 uint = 408
                        for fi41 := 0; fi41 < 12; fi41++ {
                            if len(v3) > 0 {
                                v3[0] = -2147483648
                            }
                            if (true || true) {
                                switch (fi41 >> ((42) & 63)) {
                                case 2147483647:
                                    var v42 uint = (^(v40))
                                    _ = v42
                                }
                            }
                            _ = fi41
                        }
                        _ = fi39
                        _ = v40
                    }
                default:
                    v3 = append(v3, -932)
                }
            }
        }
        _ = v35
    } else {
        for fi43 := 0; fi43 < 18; fi43++ {
            fmt.Println(v7)
            var v44 int16 = -1
            for fi45 := 0; fi45 < 10; fi45++ {
                _ = len(v3)
                var v46 int32 = int32(fi43)
                v47 := GenS4{F0: ("hello" + v4)}
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = fi43
            _ = v44
        }
    }
    v48 := GenS5{F0: false, F1: v1}
    var v49 int64 = -567
    for fi50 := 0; fi50 < 6; fi50++ {
        v2 = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi50
    }
    fmt.Println(v5)
    fmt.Println(v4, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}

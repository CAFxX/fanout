package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int16; F2 uint64 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 string; F1 uint64; F2 uint }
type GenS11 struct { F0 int32 }

func main() {
    v1 := []int{1, 9223372036854775807, 128}
    var v2 int = 32767
    var v3 int16 = 0
    v4 := []int{v2, 1, v2}
    switch "hello" {
    case "hello":
        var v5 int64 = -1
        _ = v5
    default:
        if true {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint32 = 4294967295
                v7 = v7
                switch -1 {
                case -468:
                    v8 := GenS1{F0: fi6, F1: v3, F2: fi6}
                    _ = v8
                case -2147483648:
                    var v9 string = "a"
                    _ = v9
                case 32767:
                    switch (^(v2)) {
                    case -1:
                        var v10 int = fi6
                        _ = v10
                    case 32768:
                        v4 = append(v4, 4294967295)
                    default:
                        v3 += (v3 - v3)
                    }
                default:
                    fmt.Println(v2, v7)
                }
                _ = fi6
                _ = v7
            }
            v11 := GenS2{F0: 4294967295, F1: 2147483647}
            var v12 int32 = -2147483648
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS3{F0: 1, F1: 9223372036854775807}
                var v15 uint = 0
                v16 := GenS4{F0: 2147483647, F1: ("" + "foo bar")}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            fmt.Println(v2)
            if ((!false) || true) {
                if false {
                    switch (("" + "a") + ("" + "\n\t")) {
                    case "foo bar":
                        v17 := GenS2{F0: 2, F1: 2147483647}
                        _ = v17
                    }
                } else {
                    var v18 int = -9223372036854775808
                    _ = v18
                }
                v2 += v2
                fmt.Println(v2, v3)
            } else {
                fmt.Println(v3, v2)
                var v19 int16 = v3
                switch v2 {
                case -2147483648:
                    v20 := GenS5{F0: 1, F1: 127}
                    _ = v20
                case 128:
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        if false {
                            if false {
                                switch (("hello" + "\n\t") + ("\n\t" + "")) {
                                case "":
                                    v2 += (-(v2))
                                case "a":
                                    var v21 int16 = -1
                                    _ = v21
                                }
                                switch v2 {
                                case 635:
                                    switch v2 {
                                    case 32767:
                                        var v22 string = "foo bar"
                                        _ = v22
                                    default:
                                        v1 = append(v1, (v2 % (((v2 + (-(v2)))) | 1)))
                                    }
                                case 9223372036854775807:
                                    _ = len(v4)
                                default:
                                    v3 -= v3
                                }
                            } else {
                                v23 := GenS2{F0: 2, F1: 65535}
                                for fi24 := 0; fi24 < 13; fi24++ {
                                    var v25 uint8 = 132
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v23
                            }
                            v19 += v19
                            v26 := GenS1{F0: -217, F1: v19, F2: (v2 - v2)}
                            _ = v26
                        }
                    case "hello":
                        if false {
                            for fi27 := 0; fi27 < 4; fi27++ {
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        v3 -= v19
                                    }
                                    v4 = append(v4, (v2 << ((42) & 63)))
                                    for fi28 := 0; fi28 < 10; fi28++ {
                                        v29 := GenS2{F0: 4294967295, F1: 345}
                                        fmt.Println(v2, fi27, v3)
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                                var v30 string = "a"
                                _ = fi27
                                _ = v30
                            }
                            v19 -= v19
                        }
                    }
                }
                _ = v19
            }
        }
    }
    _ = len(v4)
    switch (v2 >> ((0) & 63)) {
    case -9223372036854775808:
        var v31 int = (^(v2))
        _ = v31
    case 32768:
        v3 = (v3 - v3)
    case 42:
        v32 := GenS6{F0: 127}
        _ = v32
    default:
        for fi33 := 0; fi33 < 13; fi33++ {
            v2 = -806
            var v34 string = "foo bar"
            if ((42 != 42) || (true || false)) {
                v35 := GenS4{F0: 4294967295, F1: ("\n\t" + (v34 + v34))}
                _ = v35
            }
            _ = fi33
            _ = v34
        }
    }
    if (127 >= -128) {
        v4 = append(v4, (v2 % ((1) | 1)))
    } else {
        for fi36 := 0; fi36 < 15; fi36++ {
            fmt.Println(v3, fi36, fi36)
            v37 := GenS7{F0: -799, F1: (-(v3)), F2: 0}
            _ = fi36
            _ = v37
        }
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 bool = (478 > -2147483648)
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        for fi41 := 0; fi41 < 14; fi41++ {
            if ((-2 > 0) || (false && (v3 >= v3))) {
                v42 := GenS8{F0: (!false), F1: 2}
                v43 := GenS1{F0: 976, F1: v3, F2: 4294967295}
                var v44 uint8 = 1
                _ = v42
                _ = v43
                _ = v44
            }
            var v45 bool = (false && (true && (!((v3 | -587) <= (v3 | v3)))))
            v46 := GenS9{F0: 748, F1: true}
            _ = fi41
            _ = v45
            _ = v46
        }
        v3 = (-(v3))
        _ = fi40
    }
    v47 := GenS10{F0: "foo bar", F1: 4294967295, F2: 2}
    for fi48 := 0; fi48 < 16; fi48++ {
        v3 = -581
        if false {
            v49 := GenS2{F0: 42, F1: 329}
            switch (fi48 << ((255) & 63)) {
            case -897:
                v3 += v3
            case -791:
                var v50 int64 = 1
                _ = v50
            case 32768:
                var v51 uint16 = 1
                _ = v51
            }
            _ = v49
        }
        v2 = -186
        _ = fi48
    }
    v52 := GenS11{F0: 0}
    var v53 int64 = 0
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v52
    _ = v53
}

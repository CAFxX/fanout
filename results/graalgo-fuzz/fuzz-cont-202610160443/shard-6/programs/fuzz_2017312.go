package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int16; F2 bool }
func genFn0(p0 int, p1 uint16, p2 uint16) uint16 {
    v6 := []int{p0}
    var v7 string = "\n\t"
    v7 = "a"
    v8 := GenS1{F0: p0, F1: (-(p0)), F2: 1000}
    switch p0 {
    case 32768:
        var v9 bool = true
        _ = v9
    default:
        v10 := GenS1{F0: (-(p0)), F1: -2147483648, F2: -704}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 765
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    v3 := []string{"", "a", "foo bar", ""}
    var v4 int = -1
    var v5 int32 = 2147483647
    v11 := GenS2{F0: 2147483648, F1: 181, F2: (!true)}
    var v12 float64 = 0.0
    switch "\n\t" {
    case "foo bar":
        v13 := GenS3{F0: 0, F1: int16(v4), F2: 745}
        _ = v13
    case "":
        v14 := GenS2{F0: 452, F1: 23, F2: true}
        _ = v14
    default:
        if (128 > 128) {
            v12 += v12
            var v15 int = (v4 >> ((42) & 63))
            _ = v15
        }
    }
    v16 := GenS4{F0: v4, F1: -154}
    v17 := v1["hello"]
    var v18 uint16 = 333
    v19 := GenS1{F0: 152, F1: v4, F2: 42}
    var v20 int = v4
    fmt.Println(v20)
    for fi21 := 0; fi21 < 10; fi21++ {
        switch ("a" + "a") {
        case "foo bar":
            switch -962 {
            case 32768:
                fmt.Println(v17, v17, v4)
            case 128:
                if (!true) {
                    var v22 int32 = (v5 << ((40) & 63))
                    v23 := GenS3{F0: 32767, F1: 42, F2: (^(v17))}
                    _ = v22
                    _ = v23
                }
            }
        case "a":
            fmt.Println(v18)
        case "hello":
            _ = len(v2)
        default:
            v4 += (v4 - (v20 << ((1) & 63)))
        }
        v5 -= (v5 >> ((983) & 63))
        _ = fi21
    }
    v24 := GenS2{F0: (v17 << ((1) & 63)), F1: 2, F2: (!false)}
    switch ("" + ("\n\t" + "")) {
    case "foo bar":
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint16 = (v18 / (((v18 - (v18 << ((18446744073709551615) & 63)))) | 1))
            _ = fi25
            _ = v26
        }
    case "a":
        if false {
            if true {
                delete(v1, "hello")
            }
        } else {
            v18 = (v18 >> ((2) & 63))
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 bool = (!true)
                v28 = v28
                _ = fi27
                _ = v28
            }
            if true {
                switch 32767 {
                case 864:
                    _ = len(v1)
                default:
                    v29 := v1["a"]
                    _ = v29
                }
                if (0 > 1) {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        switch (v20 ^ -9223372036854775808) {
                        case -330:
                            var v31 uint = 55
                            _ = v31
                        case 32768:
                            v32 := GenS4{F0: (fi30 / (((v4 / (((fi30 & -12)) | 1))) | 1)), F1: 1000}
                            _ = v32
                        case 32767:
                            var v33 int64 = -205
                            _ = v33
                        default:
                            if (2 != 1) {
                                fmt.Println(v5, fi30)
                                fmt.Println(v20)
                                v34 := GenS2{F0: (v17 & 0), F1: v17, F2: (750 <= 2)}
                                _ = v34
                            } else {
                                for fi35 := 0; fi35 < 9; fi35++ {
                                    fmt.Println(fi30, v18, v12)
                                    v12 += 0.0
                                    var v36 uint64 = 0
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        }
                        _ = fi30
                    }
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v18 = v18
                        var v38 uint = 255
                        v39 := GenS1{F0: 2147483647, F1: -9223372036854775808, F2: 1000}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 14; fi40++ {
                        if false {
                            for fi41 := 0; fi41 < 4; fi41++ {
                                fmt.Println(v5, v12, v4)
                                _ = fi41
                            }
                            v17 += 0
                        } else {
                            v12 -= v12
                            v12 = 1e10
                            v4 = (-(v4))
                        }
                        switch "foo bar" {
                        case "\n\t":
                            for fi42 := 0; fi42 < 5; fi42++ {
                                v43 := v1[(("\n\t" + ("hello" + "foo bar")) + (("a" + "") + ((("foo bar" + "a") + "foo bar") + "foo bar")))]
                                v5 = 2147483647
                                if false {
                                    fmt.Println(v5)
                                    var v44 uint = 4294967295
                                    v3 = append(v3, (("\n\t" + "foo bar") + "a"))
                                    _ = v44
                                }
                                v45 := GenS5{F0: (v18 * (v18 | v18)), F1: 2}
                                _ = fi42
                                _ = v43
                                _ = v45
                            }
                        }
                        _ = fi40
                    }
                } else {
                    v46 := GenS1{F0: 0, F1: 127, F2: 32767}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v46
                }
            } else {
                delete(v1, "hello")
                for fi47 := 0; fi47 < 15; fi47++ {
                    _ = len(v3)
                    _ = fi47
                }
                v48 := GenS6{F0: false}
                _ = v48
            }
        }
    case "\n\t":
        v49 := GenS7{F0: 1, F1: -32768, F2: (42 >= -86)}
        _ = v49
    }
    var v50 int = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 bool; F1 int64; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint32; F2 uint8 }
func genFn0(p0 int) int {
    switch (p0 | 2147483647) {
    case 32767:
        v5 := GenS1{F0: 0, F1: 519}
        _ = v5
    case 813:
        v6 := []int{(p0 & p0), 2147483647}
        _ = v6
    default:
        p0 += p0
    }
    var v7 int16 = 1000
    v8 := []int{-9223372036854775808, p0}
    _ = p0
    _ = v7
    _ = v8
    return (p0 | p0)
}

func main() {
    v1 := []int{}
    var v2 uint64 = 0
    var v3 int16 = 32767
    var v4 int = 9223372036854775807
    if (-879 < (v4 * 9223372036854775807)) {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: (fi9 >> ((22) & 63)), F1: 4294967295}
            fmt.Println(v4, v3)
            var v11 int64 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v14 := GenS2{F0: v2}
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                v4 = v4
                var v16 string = ("" + "hello")
                var v17 float32 = 0.1
                _ = fi15
                _ = v16
                _ = v17
            }
            fmt.Println(fi12, v3)
            _ = fi12
        }
    } else {
        v2 += (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))
        if len(v1) > 0 {
            v1[0] = genFn0((v4 * 2147483647))
        }
    }
    var v18 uint16 = 42
    v3 = v3
    v4 = v4
    var v19 uint8 = 215
    fmt.Println(v3, v19, v19)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v20 int64 = 42
    if false {
        for fi21 := 0; fi21 < 5; fi21++ {
            if len(v1) > 0 {
                v1[0] = (fi21 + (v4 >> ((83) & 63)))
            }
            _ = fi21
        }
        var v22 int16 = v3
        _ = v22
    } else {
        if ((v4 < genFn0(genFn0(genFn0((v4 ^ v4))))) && true) {
            var v23 int8 = 77
            _ = v23
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            if (!true) {
                fmt.Println(v19, v18)
            } else {
                var v25 int16 = (v3 >> ((107) & 63))
                if false {
                    v18 -= (v18 | (v18 - (^(v18))))
                    var v26 int32 = 0
                    _ = v26
                }
                switch (fi24 + 2147483647) {
                case -9223372036854775808:
                    switch "foo bar" {
                    case "\n\t":
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 int8 = -128
                            switch 0 {
                            case 32767:
                                switch 9223372036854775807 {
                                case 127:
                                    if false {
                                        var v29 int32 = 1
                                        _ = len(v1)
                                        _ = v29
                                    } else {
                                        switch -9223372036854775808 {
                                        case 42:
                                            var v30 int32 = -987
                                            _ = v30
                                        case 4294967295:
                                            v31 := GenS2{F0: v2}
                                            _ = v31
                                        }
                                        var v32 uint16 = (v18 << ((18446744073709551615) & 63))
                                        _ = v32
                                    }
                                    v33 := GenS1{F0: fi27, F1: 255}
                                    _ = v33
                                case 32768:
                                    switch fi27 {
                                    case 896:
                                        var v34 uint64 = v2
                                        _ = v34
                                    case -9223372036854775808:
                                        v35 := GenS3{F0: 1, F1: (v25 < (v3 % (((v25 * (-(v3)))) | 1)))}
                                        _ = v35
                                    case -580:
                                        if true {
                                            v36 := GenS3{F0: 175, F1: false}
                                            _ = v36
                                        }
                                        v18 = 1
                                        switch ("hello" + ("foo bar" + "hello")) {
                                        case "hello":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "":
                                            var v37 int = fi24
                                            _ = v37
                                        case "a":
                                            _ = len(v1)
                                        default:
                                            v38 := GenS1{F0: (^(fi27)), F1: 207}
                                            _ = v38
                                        }
                                    }
                                case 4294967295:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case -1:
                                fmt.Println(fi24, v20, v19)
                            case -879:
                                fmt.Println(fi27)
                            }
                            if true {
                                v4 = -1
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                    case "foo bar":
                        for fi39 := 0; fi39 < 10; fi39++ {
                            v1 = append(v1, genFn0(v4))
                            if true {
                                v20 = (v20 / ((v20) | 1))
                            } else {
                                v3 += (v25 - 873)
                                v40 := GenS4{F0: (!(!(!(!false)))), F1: (v20 ^ 988), F2: v2}
                                fmt.Println(v4)
                                var v41 uint = 2
                                _ = v40
                                _ = v41
                            }
                            var v42 float64 = 1e10
                            _ = fi39
                            _ = v42
                        }
                    case "":
                        fmt.Println(v25, v18, v25)
                    }
                default:
                    v43 := GenS5{F0: 0, F1: 156}
                    _ = v43
                }
                _ = v25
            }
            for fi44 := 0; fi44 < 19; fi44++ {
                var v45 int32 = -2147483648
                _ = fi44
                _ = v45
            }
            _ = fi24
        }
    }
    v4 += int(v18)
    var v46 bool = (false || false)
    v20 = v20
    var v47 string = ""
    var v48 uint16 = (^(v18))
    if v46 {
        _ = len(v1)
        v4 = (v4 + v4)
    }
    v49 := GenS6{F0: (42 < 42), F1: 4294967295, F2: (v19 >> ((500) & 63))}
    var v50 uint8 = (v19 | 127)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}

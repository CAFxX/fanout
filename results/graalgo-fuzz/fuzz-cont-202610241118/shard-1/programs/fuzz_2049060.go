package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint64 }
func genFn0(p0 string) int {
    v7 := GenS1{F0: 61, F1: p0}
    var v8 bool = (42 <= 9223372036854775807)
    v9 := GenS2{F0: 2, F1: 710}
    for fi10 := 0; fi10 < 15; fi10++ {
        v8 = v8
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS2{F0: 325, F1: 4294967295}
            var v13 int16 = int16(fi10)
            if v8 {
                var v14 float32 = 1.5
                v15 := []int{200, (fi10 - fi10)}
                fmt.Println(fi11, v13)
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    v3 := []int{-354, 9223372036854775807, 846, 2147483647}
    var v4 uint = 564
    var v5 int64 = 1
    v6 := []string{}
    var v16 float64 = 3.14159
    _ = len(v3)
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v17 := v2[-602]
        var v18 int32 = 1
        _ = v17
        _ = v18
    } else {
        v5 += (-(v5))
        if false {
            _ = len(v3)
            v4 = 955
            var v19 int8 = -128
            _ = v19
        } else {
            var v20 int = genFn0(("" + "a"))
            v21 := GenS1{F0: v1, F1: "\n\t"}
            if false {
                v22 := GenS2{F0: v4, F1: (v20 % ((genFn0("\n\t")) | 1))}
                for fi23 := 0; fi23 < 13; fi23++ {
                    if (459 != -32768) {
                        var v24 uint8 = v1
                        _ = v24
                    } else {
                        v1 = 1
                        v25 := GenS3{F0: 16, F1: v1, F2: 4294967295}
                        var v26 uint64 = 405
                        _ = v25
                        _ = v26
                    }
                    var v27 int64 = v5
                    _ = fi23
                    _ = v27
                }
                _ = v22
            }
            _ = v20
            _ = v21
        }
        switch ((("a" + ("\n\t" + "\n\t")) + ("hello" + "\n\t")) + "") {
        case "foo bar":
            if false {
                fmt.Println(v16, v4, v4)
            }
        case "\n\t":
            fmt.Println(v4)
        case "":
            var v28 uint8 = (v1 & 2)
            _ = v28
        }
    }
    v1 = (^(v1))
    fmt.Println(v4)
    var v29 int8 = 42
    v1 = (^(v1))
    var v30 int = 2147483647
    v29 = v29
    for fi31 := 0; fi31 < 16; fi31++ {
        v29 += (v29 + int8(v1))
        if true {
            for fi32 := 0; fi32 < 15; fi32++ {
                if false {
                    switch (v30 | fi32) {
                    case 625:
                        v33 := GenS3{F0: 1, F1: (v1 + (v1 - v1)), F2: 207}
                        _ = v33
                    case -1:
                        v30 -= v30
                    case -622:
                        if true {
                            if (true || false) {
                                var v34 bool = false
                                var v35 uint64 = 0
                                v34 = v34
                                switch (fi32 >> (((v4 ^ (v4 % (((^(v4))) | 1)))) & 63)) {
                                case 4294967295:
                                    v29 = (v29 >> (((v4 % ((v4) | 1))) & 63))
                                    if v34 {
                                        if v34 {
                                            var v36 int16 = 42
                                            if (!false) {
                                                if false {
                                                    v37 := GenS4{F0: (!(!(v35 > (v35 << ((uint(v35)) & 63))))), F1: (1e10 <= 0.1)}
                                                    v38 := GenS3{F0: 0, F1: (v1 << ((v4) & 63)), F2: v35}
                                                    v39 := GenS2{F0: v4, F1: fi31}
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                }
                                            } else {
                                                var v40 uint = (v4 / ((1) | 1))
                                                v6 = append(v6, "")
                                                var v41 uint = (v40 | 65535)
                                                v1 -= 128
                                                v42 := GenS5{F0: true, F1: -399, F2: (v29 << ((2) & 63))}
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                            if v34 {
                                                if (100 == -70) {
                                                    v34 = (v34 || false)
                                                    if (!(!v34)) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        switch (v30 & (fi32 >> (((^(v4))) & 63))) {
                                                        case 2147483647:
                                                            var v43 bool = (v16 >= (((float64(1e10) / 0.1) * ((v16 - 0.1) + 1.5)) + v16))
                                                            _ = v43
                                                        case -2147483648:
                                                            fmt.Println(v16, fi32)
                                                        }
                                                        var v44 int32 = 526
                                                        var v45 uint8 = v1
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    v46 := GenS5{F0: v34, F1: -2147483648, F2: v29}
                                                    _ = v46
                                                } else {
                                                    switch int(v29) {
                                                    case 32767:
                                                        var v47 bool = v34
                                                        _ = v47
                                                    case -9223372036854775808:
                                                        v34 = (v16 <= (1.5 + v16))
                                                    case 127:
                                                        var v48 uint16 = uint16(v1)
                                                        _ = v48
                                                    }
                                                    fmt.Println(fi31, fi31)
                                                    var v49 uint16 = 0
                                                    _ = v49
                                                }
                                            } else {
                                                var v50 bool = false
                                                _ = v50
                                            }
                                            _ = v36
                                        }
                                        if v34 {
                                            var v51 int8 = -128
                                            v16 = v16
                                            var v52 uint16 = 32767
                                            _ = v51
                                            _ = v52
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        _ = len(v6)
                                        var v53 int = (v30 << ((750) & 63))
                                        var v54 uint32 = 642
                                        _ = v53
                                        _ = v54
                                    }
                                case -289:
                                    var v55 uint16 = uint16(v4)
                                    _ = v55
                                }
                                _ = v34
                                _ = v35
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    if ((true && (2147483648 > 2)) && (-32768 != 0)) {
                                        fmt.Println(fi31)
                                    } else {
                                        v56 := GenS1{F0: 28, F1: "a"}
                                        _ = v56
                                    }
                                case "a":
                                    var v57 uint64 = 0
                                    _ = v57
                                case "":
                                    fmt.Println(v5, fi31, v16)
                                }
                                var v58 int16 = -32768
                                var v59 int = fi32
                                _ = v58
                                _ = v59
                            }
                            var v60 int64 = (v5 ^ (^(v5)))
                            if false {
                                v61 := GenS5{F0: true, F1: 687, F2: 127}
                                _ = v61
                            } else {
                                if (2147483648 < 804) {
                                    fmt.Println(v30, v5, v4)
                                }
                                fmt.Println(v5, fi32)
                            }
                            _ = v60
                        } else {
                            v62 := GenS1{F0: 127, F1: ("a" + ("" + ""))}
                            _ = v62
                        }
                        var v63 uint16 = 680
                        v64 := v2[(v30 >> (((v4 % (((^(v4))) | 1))) & 63))]
                        _ = v63
                        _ = v64
                    }
                }
                v4 += 0
                _ = fi32
            }
            if len(v6) > 0 {
                v6[0] = ("a" + "foo bar")
            }
        } else {
            for fi65 := 0; fi65 < 11; fi65++ {
                var v66 uint32 = 65535
                if ((false && (v1 == (v1 * (^(v1))))) && (-128 >= v29)) {
                    v67 := GenS6{F0: 1, F1: 269}
                    switch (v30 * -9223372036854775808) {
                    case 366:
                        v68 := GenS4{F0: (!((-2.25 + (-2.25 * 1e10)) >= (1.5 + 1.5))), F1: (false || false)}
                        _ = v68
                    case -1:
                        v1 = v1
                    case -622:
                        var v69 uint32 = (v66 << ((v4) & 63))
                        _ = v69
                    }
                    _ = v67
                }
                v70 := GenS6{F0: 42, F1: 2}
                _ = fi65
                _ = v66
                _ = v70
            }
            var v71 float32 = 1.5
            _ = v71
        }
        var v72 int8 = (v29 / (((v29 / (((v29 & -128)) | 1))) | 1))
        _ = fi31
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

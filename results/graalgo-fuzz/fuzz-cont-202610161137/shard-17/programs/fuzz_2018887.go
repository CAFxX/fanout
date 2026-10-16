package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint; F2 bool }

func main() {
    v1 := []int{432, 9223372036854775807, -9223372036854775808}
    var v2 int32 = -2147483648
    var v3 uint = 1
    v4 := []int{}
    var v5 string = ("foo bar" + (("hello" + ("hello" + "hello")) + "foo bar"))
    var v6 int64 = 9223372036854775807
    if ((^(v6)) >= v6) {
        v2 += (v2 / (((v2 ^ (v2 * v2))) | 1))
        if (127 == 42) {
            var v7 uint16 = 42
            v8 := GenS1{F0: 65535}
            fmt.Println(v2, v2, v5)
            _ = v7
            _ = v8
        }
        fmt.Println(v6)
    }
    var v9 int8 = -1
    v2 = v2
    var v10 int8 = v9
    switch v5 {
    case "foo bar":
        v1 = append(v1, 631)
    case "":
        switch ("" + (v5 + "")) {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "a":
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: 269}
                v2 -= 42
                _ = fi12
                _ = v13
            }
            if true {
                if (42 != 127) {
                    if true {
                        switch v5 {
                        case "":
                            var v14 int = 42
                            _ = v14
                        }
                    } else {
                        fmt.Println(v9, v5)
                        for fi15 := 0; fi15 < 15; fi15++ {
                            _ = len(v4)
                            fmt.Println(fi15, v5, fi11)
                            v16 := GenS1{F0: 2147483648}
                            _ = fi15
                            _ = v16
                        }
                        v17 := GenS1{F0: 2}
                        _ = v17
                    }
                    if false {
                        switch "a" {
                        case "":
                            var v18 uint8 = 255
                            _ = v18
                        case "a":
                            if false {
                                v6 += v6
                                var v19 int64 = v6
                                if (!false) {
                                    fmt.Println(v10)
                                }
                                _ = v19
                            }
                        case "hello":
                            var v20 int32 = v2
                            _ = v20
                        }
                    } else {
                        if false {
                            var v21 uint64 = 1
                            v22 := GenS1{F0: 2147483647}
                            for fi23 := 0; fi23 < 19; fi23++ {
                                fmt.Println(fi11, v9, v3)
                                _ = fi23
                            }
                            _ = v21
                            _ = v22
                        } else {
                            if true {
                                for fi24 := 0; fi24 < 13; fi24++ {
                                    var v25 uint32 = 433
                                    if (112 != (v9 & (v10 + int8(v2)))) {
                                        v25 = 4294967295
                                        v26 := GenS1{F0: v25}
                                        v27 := GenS2{F0: (v25 != 1)}
                                        v28 := GenS2{F0: (false && (65535 <= (v25 + (^(v25)))))}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if ((!false) && false) {
                                            v3 = 983
                                            v10 += int8(fi24)
                                            if false {
                                                fmt.Println(v3)
                                                if false {
                                                    var v29 bool = (!(v5 < (((v5 + ((((("" + "hello") + "") + "\n\t") + (((v5 + ("hello" + "\n\t")) + "") + v5)) + v5)) + "hello") + "a")))
                                                    var v30 uint8 = 68
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    var v31 int = 0
                                                    _ = v31
                                                }
                                            } else {
                                                if true {
                                                    switch (fi24 << ((v3) & 63)) {
                                                    case 2147483647:
                                                        v4 = append(v4, -833)
                                                    case 9223372036854775807:
                                                        v3 = 4294967295
                                                    }
                                                    fmt.Println(v3)
                                                    fmt.Println(fi11)
                                                }
                                                v3 = 255
                                            }
                                            v6 -= (v6 / (((v6 + 42)) | 1))
                                        } else {
                                            v6 += v6
                                            var v32 uint8 = 35
                                            v33 := GenS3{F0: v9}
                                            _ = v32
                                            _ = v33
                                        }
                                        if false {
                                            v3 = v3
                                        } else {
                                            v10 -= 86
                                        }
                                        v34 := GenS3{F0: (v9 ^ int8(v3))}
                                        _ = v34
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                for fi35 := 0; fi35 < 20; fi35++ {
                                    var v36 int8 = 0
                                    switch -1 {
                                    case 32768:
                                        var v37 uint32 = 65535
                                        _ = v37
                                    case 931:
                                        if (!(128 >= 42)) {
                                            if true {
                                                if true {
                                                    _ = len(v4)
                                                    v10 = (v9 >> ((v3) & 63))
                                                    fmt.Println(fi11, v9, v3)
                                                }
                                                var v38 bool = ((v6 > v6) || true)
                                                fmt.Println(fi35)
                                                v39 := GenS4{F0: 65535, F1: 255, F2: uint64(v3)}
                                                _ = v38
                                                _ = v39
                                            }
                                            if (1 <= 0) {
                                                _ = len(v1)
                                                v5 = (v5 + ("\n\t" + (v5 + v5)))
                                            }
                                            v40 := GenS4{F0: 330, F1: 32767, F2: 754}
                                            _ = v40
                                        }
                                    case -9223372036854775808:
                                        switch fi11 {
                                        case 140:
                                            v2 -= (v2 / (((^(v2))) | 1))
                                        default:
                                            if (427 == 1) {
                                                fmt.Println(v9)
                                                _ = len(v1)
                                            } else {
                                                var v41 uint16 = 42
                                                v5 = "foo bar"
                                                _ = v41
                                            }
                                        }
                                    default:
                                        if (true || false) {
                                            if true {
                                                var v42 int64 = v6
                                                v10 -= v9
                                                v5 = v5
                                                _ = v42
                                            }
                                            _ = len(v1)
                                        }
                                    }
                                    switch v5 {
                                    case "":
                                        v10 -= v10
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            }
                            if (("hello" + v5) != (v5 + ("" + v5))) {
                                var v43 int16 = -773
                                _ = v43
                            } else {
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    if (1.5 != 3.14159) {
                                        if ((!(1.5 < 3.14159)) && true) {
                                            fmt.Println(v10)
                                        }
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = (-(fi44))
                                    }
                                    _ = fi44
                                }
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    v46 := GenS1{F0: 4294967295}
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v47 float64 = (3.14159 / 3.14159)
                        fmt.Println(v9, v9, v10)
                        _ = v47
                    }
                    v10 -= v9
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
        }
    default:
        if (false || (!(2 < 169))) {
            if ((v6 | (v6 << ((v3) & 63))) > int64(v3)) {
                if false {
                    v48 := GenS4{F0: 2147483648, F1: 0, F2: 1}
                    var v49 uint32 = 1
                    if (true && false) {
                        v50 := GenS1{F0: (v49 & 530)}
                        v51 := GenS4{F0: v49, F1: 42, F2: 4294967295}
                        var v52 int64 = 9223372036854775807
                        _ = v50
                        _ = v51
                        _ = v52
                    } else {
                        v53 := GenS5{F0: 30, F1: v3, F2: (!true)}
                        _ = v53
                    }
                    _ = v48
                    _ = v49
                }
            }
            v3 = 4294967295
            var v54 string = "foo bar"
            _ = v54
        } else {
            _ = len(v4)
            v55 := GenS4{F0: 2147483648, F1: 255, F2: 42}
            if false {
                var v56 uint = v3
                _ = v56
            }
            _ = v55
        }
    }
    var v57 int8 = v9
    v6 -= v6
    if len(v4) > 0 {
        v4[0] = -613
    }
    _ = len(v1)
    v5 = (v5 + v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v9
}

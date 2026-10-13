package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int16; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "\n\t"
    var v3 int32 = 2147483647
    fmt.Println(v3, v2)
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := []int{fi4, -9223372036854775808, 1, fi4}
        if true {
            v2 = v2
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        if false {
            if (false || (65535 != 253)) {
                v6 := GenS1{F0: "hello", F1: 756, F2: false}
                switch fi4 {
                case 4294967295:
                    fmt.Println(v2, v2)
                }
                _ = v6
            }
            fmt.Println(v3)
            fmt.Println(v3, fi4)
        }
        _ = fi4
        _ = v5
    }
    switch -1 {
    case 842:
        if (65535 >= 32767) {
            var v7 uint = 65535
            v8 := GenS1{F0: "a", F1: 65535, F2: false}
            _ = v7
            _ = v8
        } else {
            v9 := []int{-968}
            _ = v9
        }
    }
    if (!(255 > 0)) {
        for fi10 := 0; fi10 < 19; fi10++ {
            v1[1] = true
            _ = fi10
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            switch (v2 + "foo bar") {
            case "hello":
                if true {
                    v12 := GenS2{F0: 592, F1: 0}
                    var v13 int = 9223372036854775807
                    var v14 uint8 = 1
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    switch "\n\t" {
                    case "":
                        for fi15 := 0; fi15 < 12; fi15++ {
                            v2 = v2
                            _ = fi15
                        }
                    case "\n\t":
                        if false {
                            for fi16 := 0; fi16 < 4; fi16++ {
                                var v17 int64 = -9223372036854775808
                                var v18 int8 = 0
                                switch (v2 + (("a" + (v2 + "foo bar")) + v2)) {
                                case "hello":
                                    fmt.Println(fi16)
                                case "foo bar":
                                    switch (-(fi11)) {
                                    case 42:
                                        v3 += v3
                                    case -1:
                                        var v19 uint32 = 142
                                        _ = v19
                                    }
                                case "a":
                                    var v20 float32 = 0.1
                                    _ = v20
                                default:
                                    v3 += (-(v3))
                                }
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            if false {
                                switch (fi11 << ((553) & 63)) {
                                case 685:
                                    var v21 uint8 = 242
                                    _ = v21
                                case 998:
                                    v22 := GenS3{F0: 42, F1: 898, F2: 0}
                                    _ = v22
                                case 32768:
                                    var v23 int64 = -1
                                    _ = v23
                                default:
                                    if (false && true) {
                                        for fi24 := 0; fi24 < 12; fi24++ {
                                            if true {
                                                var v25 bool = ((0 > 18446744073709551615) && false)
                                                _ = v25
                                            }
                                            v26 := GenS3{F0: 1, F1: 816, F2: 1000}
                                            _ = fi24
                                            _ = v26
                                        }
                                        fmt.Println(v2, v2)
                                    } else {
                                        v3 += (v3 ^ (v3 & (v3 >> ((1) & 63))))
                                        for fi27 := 0; fi27 < 7; fi27++ {
                                            v2 = ((("\n\t" + "hello") + v2) + (v2 + "foo bar"))
                                            var v28 int32 = 42
                                            _ = fi27
                                            _ = v28
                                        }
                                        fmt.Println(v2, v2, v2)
                                    }
                                }
                                for fi29 := 0; fi29 < 5; fi29++ {
                                    v30 := GenS3{F0: 65535, F1: 273, F2: -1}
                                    _ = fi29
                                    _ = v30
                                }
                                fmt.Println(fi11, fi11)
                            }
                            var v31 string = (v2 + v2)
                            _ = v31
                        } else {
                            if false {
                                switch ("a" + v2) {
                                case "a":
                                    for fi32 := 0; fi32 < 20; fi32++ {
                                        var v33 string = v2
                                        switch fi11 {
                                        case 2147483647:
                                            v34 := []int{-1, (fi32 ^ fi11), fi32}
                                            _ = v34
                                        default:
                                            switch (v33 + v2) {
                                            case "foo bar":
                                                fmt.Println(v2, fi11)
                                            case "a":
                                                if false {
                                                    v3 = (v3 / (((v3 >> ((65535) & 63))) | 1))
                                                } else {
                                                    fmt.Println(v3, v3, v3)
                                                }
                                            }
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                default:
                                    for fi35 := 0; fi35 < 17; fi35++ {
                                        switch fi35 {
                                        case -9223372036854775808:
                                            if true {
                                                v3 -= v3
                                                var v36 uint16 = 1
                                                v2 = ("" + "a")
                                                _ = v36
                                            }
                                        case 4294967295:
                                            v37 := []int{(fi11 / ((int(v3)) | 1)), (^(fi11))}
                                            _ = v37
                                        case -318:
                                            var v38 uint = 65535
                                            _ = v38
                                        default:
                                            fmt.Println(v3, v3)
                                        }
                                        var v39 int64 = 1
                                        if (((^(v3)) <= v3) || (821 != 0)) {
                                            if true {
                                                var v40 int8 = 89
                                                _ = v40
                                            }
                                            fmt.Println(fi11)
                                            if true {
                                                v41 := GenS4{F0: true}
                                                _ = v41
                                            }
                                            fmt.Println(v39, v3, v39)
                                        }
                                        _ = fi35
                                        _ = v39
                                    }
                                }
                                var v42 float32 = 1.5
                                v42 -= (v42 + v42)
                                _ = v42
                            }
                        }
                    case "hello":
                        var v43 float64 = 1e10
                        _ = v43
                    default:
                        v44 := GenS5{F0: fi11, F1: 32767, F2: 68}
                        _ = v44
                    }
                    for fi45 := 0; fi45 < 3; fi45++ {
                        v46 := GenS6{F0: -48, F1: 625}
                        _ = fi45
                        _ = v46
                    }
                }
            }
            var v47 string = ""
            fmt.Println(v3, fi11)
            _ = fi11
            _ = v47
        }
        var v48 int64 = -1
        _ = v48
    }
    var v49 float32 = 1e10
    var v50 uint8 = 128
    var v51 bool = false
    switch 906 {
    case 42:
        v51 = v51
    default:
        v51 = v51
    }
    delete(v1, -867)
    for fi52 := 0; fi52 < 19; fi52++ {
        var v53 int = int(v50)
        v49 -= (v49 + (0.1 - v49))
        _ = fi52
        _ = v53
    }
    for fi54 := 0; fi54 < 13; fi54++ {
        var v55 float32 = (v49 + (0.1 - (1.5 - float32(1.5))))
        v55 = v49
        _ = fi54
        _ = v55
    }
    switch -783 {
    case -171:
        var v56 int32 = (^(v3))
        _ = v56
    default:
        v50 -= (v50 - (v50 << ((339) & 63)))
    }
    v2 = v2
    var v57 string = ""
    v58 := v1[-9223372036854775808]
    fmt.Println(v50, v51, v58)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v51
    _ = v57
    _ = v58
}

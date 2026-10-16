package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint16; F1 bool }
func genFn0() uint {
    v6 := GenS1{F0: false, F1: 2147483647, F2: 182}
    var v7 int16 = -32768
    var v8 float64 = 0.1
    v8 += v8
    var v9 uint16 = 421
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 64
}
func genFn1() uint16 {
    if true {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(fi10)
            v11 := GenS1{F0: true, F1: fi10, F2: 2}
            var v12 int = fi10
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint8 = 1
        v14 := []int{-818, -1}
        _ = v13
        _ = v14
    } else {
        var v15 float64 = 1.5
        v16 := GenS1{F0: false, F1: 32768, F2: 42}
        v17 := []int{}
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    var v19 float32 = float32(-2.25)
    v19 += (((0.0 * float32(0.1)) / 1.0) * v19)
    _ = v19
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.1
    var v3 uint16 = 998
    var v4 uint16 = v3
    var v5 int32 = -1
    v3 += v4
    switch "\n\t" {
    case "hello":
        fmt.Println(v2, v2)
    case "\n\t":
        if true {
            var v20 int16 = 42
            fmt.Println(v2, v20, v20)
            v3 = (v3 ^ (v3 * 42))
            _ = v20
        } else {
            fmt.Println(v4, v2, v5)
            var v21 float64 = (v2 * 1e10)
            _ = v21
        }
    case "":
        v5 = v5
    default:
        switch 773 {
        case 127:
            if false {
                if true {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v4 -= (v4 + (v3 << ((42) & 63)))
                        if true {
                            v23 := []int{(fi22 % (((fi22 | 0)) | 1)), int(v4), (-(fi22)), int(v3)}
                            var v24 uint64 = 60
                            _ = v23
                            _ = v24
                        } else {
                            if false {
                                switch fi22 {
                                case 0:
                                    v2 = (v2 + 0.1)
                                case 9223372036854775807:
                                    var v25 int32 = v5
                                    _ = v25
                                case 128:
                                    switch "hello" {
                                    case "hello":
                                        v2 = v2
                                    case "":
                                        v3 = (v3 << ((0) & 63))
                                    case "foo bar":
                                        if false {
                                            v3 = (^(v4))
                                        }
                                    default:
                                        v5 = v5
                                    }
                                default:
                                    var v26 uint8 = 2
                                    _ = v26
                                }
                                v27 := []int{(fi22 ^ fi22), 1, (fi22 >> ((3) & 63)), 32767}
                                _ = v27
                            } else {
                                var v28 bool = false
                                v29 := GenS1{F0: (897 < 2), F1: (fi22 - fi22), F2: 127}
                                _ = v28
                                _ = v29
                            }
                            v3 -= (v4 >> ((genFn0()) & 63))
                            v30 := []int{(^(fi22)), fi22, fi22}
                            _ = v30
                        }
                        _ = fi22
                    }
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        if (!false) {
                            switch -9223372036854775808 {
                            case 202:
                                v5 = 2147483647
                            case -1:
                                v31 := []int{}
                                _ = v31
                            case 537:
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    v4 -= 65535
                                    _ = fi32
                                }
                            }
                        }
                    case "a":
                        var v33 int = 4294967295
                        _ = v33
                    case "hello":
                        var v34 int64 = 0
                        _ = v34
                    }
                    v5 = (v5 & v5)
                } else {
                    for fi35 := 0; fi35 < 1; fi35++ {
                        fmt.Println(v4)
                        fmt.Println(v2, v3, fi35)
                        _ = fi35
                    }
                    if false {
                        v4 -= v3
                    } else {
                        v36 := GenS1{F0: true, F1: 999, F2: 42}
                        _ = v36
                    }
                    v4 = v4
                }
                var v37 uint8 = 71
                _ = v37
            }
        case 1:
            fmt.Println(v5)
        case 32767:
            v38 := GenS1{F0: true, F1: 9223372036854775807, F2: 128}
            _ = v38
        default:
            if ("foo bar" == "foo bar") {
                v39 := GenS2{F0: 917, F1: "foo bar"}
                _ = v39
            } else {
                for fi40 := 0; fi40 < 3; fi40++ {
                    switch fi40 {
                    case 128:
                        v41 := []int{fi40, -9223372036854775808}
                        _ = v41
                    }
                    switch -2147483648 {
                    case -2147483648:
                        for fi42 := 0; fi42 < 2; fi42++ {
                            var v43 uint = 2
                            _ = fi42
                            _ = v43
                        }
                    case 128:
                        var v44 int8 = 100
                        _ = v44
                    }
                    switch ("foo bar" + "a") {
                    case "hello":
                        v3 = (^(v4))
                    case "":
                        v45 := []int{fi40, 0, (-(fi40)), -9223372036854775808}
                        _ = v45
                    }
                    _ = fi40
                }
            }
        }
    }
    v46 := GenS3{F0: (v4 << ((1) & 63)), F1: -165}
    var v47 int16 = -423
    var v48 float64 = 0.1
    if true {
        var v49 string = "foo bar"
        for fi50 := 0; fi50 < 8; fi50++ {
            if false {
                for fi51 := 0; fi51 < 13; fi51++ {
                    var v52 uint64 = 4294967295
                    _ = fi51
                    _ = v52
                }
                fmt.Println(v49)
            }
            if (fi50 != (-(fi50))) {
                switch 128 {
                case 42:
                    if false {
                        if false {
                            fmt.Println(v47)
                            v48 += v2
                            v53 := GenS4{F0: 0, F1: int64(v47)}
                            _ = v53
                        } else {
                            for fi54 := 0; fi54 < 13; fi54++ {
                                var v55 uint64 = 945
                                v56 := GenS5{F0: (^(v4)), F1: (false || (42 == uint32(v4)))}
                                v4 = v3
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                        }
                        switch (v49 + ("a" + v49)) {
                        case "hello":
                            v57 := []int{(^(fi50)), (fi50 / ((fi50) | 1))}
                            _ = v57
                        default:
                            if (false || true) {
                                var v58 uint = 2
                                var v59 float64 = v2
                                _ = v58
                                _ = v59
                            }
                        }
                    }
                case -2147483648:
                    v47 += 42
                case -220:
                    for fi60 := 0; fi60 < 13; fi60++ {
                        switch -2147483648 {
                        case 585:
                            v5 -= (v5 << ((4294967295) & 63))
                        }
                        var v61 uint8 = 42
                        _ = fi60
                        _ = v61
                    }
                }
                fmt.Println(v2, fi50)
                v62 := []int{}
                _ = v62
            }
            _ = fi50
        }
        _ = v49
    }
    v63 := GenS2{F0: 4294967295, F1: "hello"}
    v64 := GenS3{F0: genFn1(), F1: -9223372036854775808}
    v1[0] = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v63
    _ = v64
}

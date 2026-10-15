package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 string; F1 int32; F2 int16 }
type GenS5 struct { F0 uint; F1 uint64; F2 int16 }
type GenS6 struct { F0 int64 }
func genFn0() uint64 {
    if true {
        var v5 int8 = -1
        v6 := GenS1{F0: 0}
        v7 := GenS1{F0: 4294967295}
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 bool = (-197 >= -1)
        switch 0 {
        case 2147483647:
            fmt.Println(v8)
        case 1:
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := []int{42, fi9, fi9, -2147483648}
                _ = fi9
                _ = v10
            }
        case 343:
            var v11 uint32 = 790
            _ = v11
        default:
            v8 = false
        }
        _ = v8
    }
    if false {
        var v12 int16 = -673
        var v13 bool = (true || (255 >= 0))
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 9223372036854775807}
        var v15 uint16 = 223
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: "foo bar", F1: 439}
    _ = v16
    return 814
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 0.0
    v3 := make(map[int]string)
    var v4 float32 = float32(v2)
    if false {
        if false {
            var v17 float64 = v2
            fmt.Println(v4)
            _ = v17
        } else {
            v2 += 1.5
        }
        if true {
            v18 := GenS2{F0: ("foo bar" + ("foo bar" + "")), F1: 71}
            _ = v18
        }
    } else {
        fmt.Println(v4, v1, v4)
        for fi19 := 0; fi19 < 10; fi19++ {
            v2 += v2
            var v20 uint8 = 202
            if false {
                if (!true) {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        v2 = (float64((v1 / ((v4 / (v1 * 1.5)) / -2.25))) * v2)
                        _ = fi21
                    }
                    var v22 uint64 = 42
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 uint64 = (v22 * 1)
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
                switch -9223372036854775808 {
                case 128:
                    fmt.Println(v1)
                }
                var v25 float64 = v2
                _ = v25
            } else {
                v1 = (((3.14159 * (0.0 / float32(1e10))) - v1) + v1)
            }
            _ = fi19
            _ = v20
        }
    }
    var v26 int32 = 42
    v27 := GenS1{F0: 128}
    for fi28 := 0; fi28 < 9; fi28++ {
        if ((fi28 << ((0) & 63)) != (fi28 - (fi28 & (fi28 ^ (fi28 + (^(fi28))))))) {
            switch fi28 {
            case -2147483648:
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 int64 = -154
                    fmt.Println(v30)
                    switch (^(fi28)) {
                    case 32767:
                        v2 = v2
                    case 42:
                        switch ("" + "hello") {
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                v31 := v3[fi28]
                                _ = v31
                            case "\n\t":
                                v2 -= 1.5
                            }
                        case "\n\t":
                            v2 = v2
                        }
                    case 9223372036854775807:
                        var v32 int = 250
                        _ = v32
                    default:
                        v33 := GenS3{F0: 32767, F1: v26, F2: 0}
                        _ = v33
                    }
                    v34 := []int{(fi28 >> ((65535) & 63)), fi29, (fi29 * 4294967295), fi29}
                    var v35 int32 = v26
                    v26 += -1
                    _ = fi29
                    _ = v30
                    _ = v34
                    _ = v35
                }
            case 128:
                for fi36 := 0; fi36 < 3; fi36++ {
                    v37 := GenS4{F0: "foo bar", F1: 425, F2: 1}
                    fmt.Println(v26, fi36)
                    _ = fi36
                    _ = v37
                }
            default:
                v38 := GenS1{F0: fi28}
                _ = v38
            }
            var v39 uint64 = 42
            v2 = (v2 + float64(1e10))
            _ = v39
        } else {
            v40 := []int{(fi28 % ((9223372036854775807) | 1)), (fi28 - (fi28 / (((-(fi28))) | 1))), fi28}
            _ = v40
        }
        var v41 uint = 42
        _ = fi28
        _ = v41
    }
    v42 := GenS5{F0: 333, F1: 42, F2: 1}
    v1 += -2.25
    var v43 uint64 = 250
    v44 := []int{4294967295, 1}
    v26 += (v26 / (((v26 & (v26 % (((v26 / ((-192) | 1))) | 1)))) | 1))
    for fi45 := 0; fi45 < 9; fi45++ {
        v46 := GenS4{F0: "\n\t", F1: (v26 & (v26 ^ (v26 / ((v26) | 1)))), F2: 0}
        switch -855 {
        case 42:
            var v47 int8 = -1
            _ = v47
        default:
            var v48 uint64 = v43
            _ = v48
        }
        v49 := GenS6{F0: -37}
        _ = fi45
        _ = v46
        _ = v49
    }
    var v50 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v50
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint8; F1 uint16; F2 bool }
type GenS5 struct { F0 int8 }
func genFn0() uint16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 15; fi8++ {
        if (255 == 0) {
            var v9 int = fi8
            _ = v9
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (42 != 1) {
            var v10 int64 = 42
            _ = v10
        }
        _ = fi8
    }
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := make(map[int]bool)
    var v4 float64 = (0.1 / 3.14159)
    v5 := make(map[string]int)
    var v6 uint64 = 42
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v2, v6, v2)
        _ = fi11
    }
    v12 := v3[2147483647]
    if v12 {
        if true {
            v13 := GenS1{F0: v6, F1: -9223372036854775808}
            _ = v13
        } else {
            v3[-790] = true
        }
        switch ((("hello" + "a") + "\n\t") + ("hello" + ("hello" + "\n\t"))) {
        case "hello":
            switch "" {
            case "\n\t":
                v12 = ((v2 < (v2 << ((255) & 63))) && (2 != 0))
            }
        case "":
            if false {
                switch ((((("\n\t" + "hello") + ("foo bar" + "")) + ("hello" + "foo bar")) + "hello") + "\n\t") {
                case "a":
                    v6 = 0
                case "":
                    v14 := GenS1{F0: 1, F1: -838}
                    _ = v14
                case "\n\t":
                    v6 -= (v6 - (v6 + v6))
                }
            }
        case "\n\t":
            if v12 {
                for fi15 := 0; fi15 < 2; fi15++ {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS1{F0: 2, F1: 0}
                        v4 += v4
                        _ = fi16
                        _ = v17
                    }
                    switch ("foo bar" + ("foo bar" + "\n\t")) {
                    case "foo bar":
                        switch (fi15 - (fi15 - (fi15 / (((-(fi15))) | 1)))) {
                        case 32767:
                            for fi18 := 0; fi18 < 15; fi18++ {
                                var v19 uint8 = 0
                                v20 := []int{}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        default:
                            switch "hello" {
                            case "a":
                                for fi21 := 0; fi21 < 18; fi21++ {
                                    if (1 <= 9223372036854775807) {
                                        var v22 uint8 = (v2 >> ((65535) & 63))
                                        var v23 uint64 = 2
                                        var v24 uint16 = 255
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(v6)
                                        case "\n\t":
                                            if v12 {
                                                _ = len(v1)
                                                if true {
                                                    v6 -= (v23 / ((v6) | 1))
                                                }
                                            }
                                        }
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                    v2 += (v2 + (v2 >> ((1) & 63)))
                                    var v25 bool = (-1 != 9223372036854775807)
                                    _ = fi21
                                    _ = v25
                                }
                            case "foo bar":
                                var v26 float64 = v4
                                _ = v26
                            default:
                                v4 = v4
                            }
                        }
                    case "\n\t":
                        v4 = v4
                    case "a":
                        var v27 int64 = int64(v6)
                        _ = v27
                    }
                    if false {
                        var v28 uint64 = v6
                        v4 -= ((float64((0.0 * 0.1)) + (v4 - 3.14159)) + float64(3.14159))
                        v6 = 4294967295
                        _ = v28
                    } else {
                        var v29 int = int(v6)
                        v30 := GenS2{F0: v12}
                        _ = v29
                        _ = v30
                    }
                    _ = fi15
                }
            }
        }
        var v31 int16 = -32768
        _ = v31
    }
    var v32 float64 = 0.1
    v4 = -2.25
    v33 := []int{2147483647, -1}
    v34 := GenS1{F0: (v6 | v6), F1: 42}
    for fi35 := 0; fi35 < 17; fi35++ {
        v3[-173] = (!(v12 && ("\n\t" == "a")))
        _ = fi35
    }
    var v36 int8 = -22
    if true {
        v36 = v36
    } else {
        if (v2 > v2) {
            if (!v12) {
                v37 := GenS3{F0: 1, F1: (1.5 <= (-2.25 - ((-2.25 - (((1e10 / -2.25) * 0.1) * ((1.5 * ((3.14159 - 0.1) + (0.0 + (1.5 + ((1e10 * 1e10) - 3.14159))))) * 1.5))) * (1.5 * (3.14159 * 1.5)))))}
                _ = v37
            }
        } else {
            switch "\n\t" {
            case "a":
                switch "" {
                case "a":
                    v38 := GenS1{F0: (v6 << ((402) & 63)), F1: 0}
                    _ = v38
                case "\n\t":
                    var v39 int8 = v36
                    _ = v39
                default:
                    v6 += (v6 | v6)
                }
            case "":
                var v40 string = ("\n\t" + (("foo bar" + "") + ((("hello" + "foo bar") + ((("hello" + "") + "a") + "a")) + "foo bar")))
                _ = v40
            case "foo bar":
                var v41 uint8 = v2
                _ = v41
            }
        }
    }
    v36 += v36
    if ((("foo bar" + ("\n\t" + ("" + ("" + "\n\t")))) + "\n\t") > "hello") {
        if false {
            v36 -= (v36 % (((v36 + v36)) | 1))
            v6 = v6
        }
        v42 := GenS4{F0: (v2 + 5), F1: 65535, F2: (v12 || (!(!false)))}
        v4 = (v32 + -2.25)
        _ = v42
    } else {
        _ = len(v33)
        if false {
            v43 := GenS5{F0: (v36 % (((v36 & 42)) | 1))}
            if len(v33) > 0 {
                v33[0] = -410
            }
            _ = v43
        } else {
            var v44 int8 = (v36 % ((0) | 1))
            for fi45 := 0; fi45 < 13; fi45++ {
                var v46 int32 = 2147483647
                switch "a" {
                case "a":
                    if (930 < v6) {
                        var v47 int8 = v44
                        for fi48 := 0; fi48 < 16; fi48++ {
                            var v49 uint32 = 2147483648
                            v46 -= v46
                            v44 = v47
                            _ = fi48
                            _ = v49
                        }
                        switch "" {
                        case "a":
                            for fi50 := 0; fi50 < 8; fi50++ {
                                var v51 float32 = 1.5
                                var v52 int16 = -32768
                                fmt.Println(v2)
                                delete(v1, (fi50 ^ (fi45 * fi45)))
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                        case "foo bar":
                            fmt.Println(v2, v4, v2)
                        }
                        _ = v47
                    }
                case "":
                    v1[fi45] = fi45
                default:
                    for fi53 := 0; fi53 < 5; fi53++ {
                        delete(v1, fi53)
                        v54 := v1[1]
                        v4 -= (1e10 * ((v4 * -2.25) * 1.5))
                        _ = fi53
                        _ = v54
                    }
                }
                if len(v33) > 0 {
                    v33[0] = fi45
                }
                _ = fi45
                _ = v46
            }
            var v55 uint = 282
            _ = v44
            _ = v55
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

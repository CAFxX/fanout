package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 string }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint8; F1 bool; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 string }
func genFn0() int32 {
    var v4 int8 = 127
    if false {
        if (32767 < 196) {
            v5 := GenS1{F0: (^(v4))}
            var v6 uint32 = 2147483648
            _ = v5
            _ = v6
        }
        v7 := GenS2{F0: -1}
        v4 -= (v4 - (v4 % ((v4) | 1)))
        _ = v7
    } else {
        var v8 bool = false
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint64 = 1
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS3{F0: fi9, F1: v10, F2: "a"}
            v4 += v4
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 float64 = ((0.1 + 1e10) - ((1e10 / 0.1) - (3.14159 / (float64(0.1) / 1.0))))
    v4 += v4
    _ = v13
    _ = v4
    return 1
}

func main() {
    var v1 int = 2147483647
    var v2 uint = 0
    v3 := make(map[int]bool)
    if true {
        switch v1 {
        case -781:
            v2 += 0
        case -291:
            if false {
                v2 = v2
                var v14 uint8 = 67
                var v15 uint8 = (v14 / ((v14) | 1))
                _ = v14
                _ = v15
            } else {
                if true {
                    switch v1 {
                    case 32768:
                        v16 := GenS4{F0: "a", F1: 18446744073709551615}
                        _ = v16
                    }
                    if (true || false) {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v18 := []int{v1, (v1 % (((fi17 + -1)) | 1)), (v1 ^ fi17), (fi17 ^ (fi17 << (((v2 - (v2 | (v2 >> (((v2 + v2)) & 63))))) & 63)))}
                            if false {
                                _ = len(v18)
                            } else {
                                v3[fi17] = true
                                v19 := GenS1{F0: 98}
                                _ = v19
                            }
                            _ = fi17
                            _ = v18
                        }
                        var v20 uint32 = 975
                        var v21 uint = (v2 - v2)
                        _ = v20
                        _ = v21
                    } else {
                        v1 += v1
                        for fi22 := 0; fi22 < 17; fi22++ {
                            if true {
                                fmt.Println(v1)
                            }
                            _ = fi22
                        }
                    }
                } else {
                    if (42 <= 427) {
                        var v23 uint64 = 2
                        if true {
                            var v24 float64 = (((float64(0.1) + (float64(1.5) + (1e10 + 1.5))) / ((-2.25 - float64(0.0)) - float64((3.14159 - 0.1)))) - float64(0.1))
                            v25 := GenS2{F0: 0}
                            _ = v24
                            _ = v25
                        } else {
                            if true {
                                v26 := GenS2{F0: 587}
                                if (true && (202 <= 0)) {
                                    fmt.Println(v1, v1)
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        switch ("\n\t" + "\n\t") {
                                        case "foo bar":
                                            switch 519 {
                                            case -958:
                                                for fi28 := 0; fi28 < 7; fi28++ {
                                                    v29 := v3[0]
                                                    _ = fi28
                                                    _ = v29
                                                }
                                            }
                                        default:
                                            var v30 int64 = 1
                                            _ = v30
                                        }
                                        var v31 int8 = 127
                                        _ = fi27
                                        _ = v31
                                    }
                                    v32 := GenS5{F0: 0, F1: (2147483647 <= 2147483647), F2: "foo bar"}
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    v23 += v23
                                    v23 = v23
                                    _ = fi33
                                }
                                _ = v26
                            } else {
                                var v34 int64 = -1
                                fmt.Println(v2, v23, v1)
                                _ = v34
                            }
                        }
                        _ = v23
                    }
                }
                v1 = v1
            }
        case -9223372036854775808:
            for fi35 := 0; fi35 < 5; fi35++ {
                switch fi35 {
                case 88:
                    var v36 bool = false
                    _ = v36
                case 9223372036854775807:
                    if true {
                        var v37 uint16 = 42
                        _ = v37
                    }
                default:
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v1 = v1
                        _ = fi38
                    }
                }
                _ = fi35
            }
        }
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            for fi40 := 0; fi40 < 16; fi40++ {
                var v41 float32 = 0.1
                switch (fi39 + -573) {
                case 1:
                    v3[fi40] = ("a" > "")
                }
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
        switch (v1 % ((v1) | 1)) {
        case 0:
            if (!true) {
                v42 := []int{(-(v1)), v1, v1}
                _ = v42
            }
        case 988:
            fmt.Println(v1, v2)
        case 1:
            v43 := GenS4{F0: (("a" + "hello") + "hello"), F1: 287}
            _ = v43
        }
        v2 += 0
    }
    v44 := []int{}
    switch "" {
    case "foo bar":
        v2 = v2
    case "":
        v45 := GenS1{F0: 81}
        _ = v45
    default:
        v2 = 18446744073709551615
    }
    v46 := GenS6{F0: 32}
    fmt.Println(v2)
    var v47 int = -1
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 uint8 = 42
        _ = fi48
        _ = v49
    }
    v50 := GenS7{F0: 42, F1: 1, F2: "a"}
    v44 = append(v44, (v47 - v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v46
    _ = v47
    _ = v50
}

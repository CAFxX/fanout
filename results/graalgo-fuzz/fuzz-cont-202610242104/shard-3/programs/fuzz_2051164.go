package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int8; F2 bool }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = 120
    var v3 uint = 42
    switch 77 {
    case 32767:
        if (255 <= 127) {
            switch -2147483648 {
            case -2147483648:
                fmt.Println(v1)
            case 872:
                if (65535 <= 1) {
                    switch -9223372036854775808 {
                    case 0:
                        var v4 uint32 = 0
                        _ = v4
                    case 32767:
                        fmt.Println(v3, v2)
                    }
                    for fi5 := 0; fi5 < 9; fi5++ {
                        switch fi5 {
                        case 42:
                            var v6 float64 = 1e10
                            _ = v6
                        default:
                            switch ("" + "a") {
                            case "":
                                v2 -= (v2 | 0)
                            default:
                                v7 := GenS1{F0: 42, F1: 832}
                                _ = v7
                            }
                        }
                        fmt.Println(v3, v1)
                        var v8 int64 = 555
                        _ = fi5
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v10 := []int{(-(fi9))}
                        if true {
                            for fi11 := 0; fi11 < 2; fi11++ {
                                v12 := GenS1{F0: 0, F1: 1}
                                _ = fi11
                                _ = v12
                            }
                        } else {
                            v13 := GenS1{F0: 2147483647, F1: 42}
                            if ((true || true) || (!true)) {
                                v14 := GenS1{F0: 88, F1: 710}
                                v3 -= (v1 >> ((v1) & 63))
                                _ = v14
                            }
                            _ = v13
                        }
                        _ = fi9
                        _ = v10
                    }
                } else {
                    v15 := []int{899, -1, 476}
                    v15 = append(v15, -186)
                    _ = v15
                }
            case 311:
                v16 := GenS1{F0: 42, F1: -2147483648}
                _ = v16
            default:
                v2 -= 1
            }
            v1 -= 590
            var v17 uint32 = 2
            _ = v17
        } else {
            fmt.Println(v2, v2, v2)
            if (42 <= 429) {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v3 += 4294967295
                    v19 := []int{(^(fi18))}
                    v3 -= 42
                    _ = fi18
                    _ = v19
                }
                v1 += (^(v1))
            } else {
                if true {
                    v20 := GenS2{F0: false, F1: 0, F2: -2147483648}
                    v21 := []int{}
                    var v22 uint = 1
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS1{F0: -1, F1: 1}
                    _ = v23
                }
                v2 += (v2 * v2)
            }
        }
    case 127:
        fmt.Println(v1, v2)
    default:
        fmt.Println(v1)
    }
    v24 := GenS2{F0: (v2 < (^(v2))), F1: 795, F2: -609}
    if true {
        if false {
            for fi25 := 0; fi25 < 18; fi25++ {
                fmt.Println(fi25, v2)
                _ = fi25
            }
            switch 9223372036854775807 {
            case 388:
                fmt.Println(v1, v1, v2)
            case 713:
                v26 := GenS2{F0: true, F1: 65535, F2: 1}
                _ = v26
            case 171:
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch -262 {
                    case 439:
                        v1 -= 0
                    case -958:
                        v1 = 141
                    case 4294967295:
                        switch "" {
                        case "foo bar":
                            v28 := []int{}
                            _ = v28
                        }
                    default:
                        v29 := GenS1{F0: 556, F1: 42}
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 10; fi30++ {
                        v31 := []int{fi30, (-(fi30)), fi27, (fi30 ^ -2147483648)}
                        if (true && (100 == 0)) {
                            v32 := GenS3{F0: ""}
                            if false {
                                var v33 int8 = 113
                                var v34 int8 = (v33 % ((v33) | 1))
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(fi30, v2, fi27)
                            _ = v32
                        }
                        var v35 uint32 = 2147483648
                        _ = fi30
                        _ = v31
                        _ = v35
                    }
                    _ = fi27
                }
            default:
                fmt.Println(v3, v3)
            }
            fmt.Println(v1)
        }
    }
    var v36 uint8 = 127
    for fi37 := 0; fi37 < 9; fi37++ {
        v38 := GenS4{F0: 759, F1: 54, F2: false}
        var v39 uint64 = 2
        _ = fi37
        _ = v38
        _ = v39
    }
    if false {
        var v40 float32 = 1e10
        for fi41 := 0; fi41 < 12; fi41++ {
            for fi42 := 0; fi42 < 11; fi42++ {
                if ((v40 + v40) < v40) {
                    var v43 float32 = 0.0
                    if ((false && (42 >= -21)) && ((42 < 1) || true)) {
                        var v44 int64 = (v2 >> (((v1 | (v1 + v1))) & 63))
                        v3 += (v1 & (v1 << (((v1 | (v3 | (v3 - (v1 & v1))))) & 63)))
                        _ = v44
                    }
                    v45 := GenS1{F0: -1, F1: 0}
                    _ = v43
                    _ = v45
                }
                v1 = 18446744073709551615
                v46 := []int{(fi41 & fi41), (fi41 & (fi42 - 819))}
                _ = len(v46)
                _ = fi42
                _ = v46
            }
            switch ("\n\t" + "a") {
            case "foo bar":
                v47 := GenS3{F0: "a"}
                _ = v47
            case "a":
                v1 = v3
            case "hello":
                v48 := []int{32768, fi41, (fi41 ^ (fi41 / ((fi41) | 1))), fi41}
                _ = v48
            }
            for fi49 := 0; fi49 < 9; fi49++ {
                v50 := []int{}
                _ = fi49
                _ = v50
            }
            _ = fi41
        }
        _ = v40
    } else {
        fmt.Println(v2)
        if (false || (1 <= uint64(v3))) {
            var v51 float64 = 1e10
            _ = v51
        } else {
            v3 += v3
        }
        var v52 uint = (^(v1))
        _ = v52
    }
    v53 := []int{9223372036854775807, 32767}
    var v54 bool = ((0.0 != (-2.25 * ((0.0 * 0.0) + ((0.1 - (1e10 * 0.1)) + 0.0)))) || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v53
    _ = v54
}

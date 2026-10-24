package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 int8; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint32; F2 int8 }
func genFn0() int {
    v4 := GenS1{F0: 126, F1: "foo bar", F2: 338}
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{42, 128, fi5}
        var v7 int16 = -32768
        switch (^(fi5)) {
        case -480:
            v7 += 890
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float64 = -2.25
    if false {
        var v9 int16 = 32767
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(fi10, v8, fi10)
            var v11 float32 = 0.1
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = v4
    _ = v8
    return -117
}
func genFn1() int16 {
    v12 := []int{32768, -2147483648, 4294967295}
    _ = len(v12)
    switch 923 {
    case 4294967295:
        v13 := GenS1{F0: -737, F1: "foo bar", F2: 0}
        _ = v13
    default:
        if true {
            v12 = append(v12, 128)
            switch "\n\t" {
            case "a":
                var v14 int32 = -2147483648
                _ = v14
            default:
                var v15 int = 1
                _ = v15
            }
            if len(v12) > 0 {
                v12[0] = genFn0()
            }
        } else {
            var v16 int32 = -2147483648
            _ = v16
        }
    }
    v17 := GenS2{F0: 0, F1: 641, F2: 1}
    _ = v12
    _ = v17
    return -178
}

func main() {
    var v1 int64 = 822
    v2 := make(map[string]int)
    var v3 uint32 = 440
    var v18 int32 = 42
    v1 = v1
    var v19 string = ""
    v20 := GenS3{F0: 0, F1: 4294967295, F2: 42}
    if (v3 != 0) {
        for fi21 := 0; fi21 < 2; fi21++ {
            v1 -= v1
            _ = fi21
        }
        v1 -= -148
    } else {
        v3 = (^(v3))
        v18 = (v18 ^ -83)
        if (v19 < v19) {
            v22 := GenS1{F0: (v1 | (v1 - -9223372036854775808)), F1: "", F2: 698}
            _ = v22
        } else {
            v23 := GenS3{F0: 691, F1: v3, F2: 1}
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 int16 = genFn1()
                v26 := GenS2{F0: 99, F1: 4294967295, F2: 18446744073709551615}
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 int16 = 163
            _ = v23
            _ = v27
        }
    }
    var v28 uint = 975
    if true {
        delete(v2, ("" + v19))
        var v29 uint64 = 440
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 uint32 = uint32(v29)
            delete(v2, "")
            v28 += 1
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        var v32 uint8 = 15
        var v33 uint = 335
        var v34 float64 = 1e10
        _ = v32
        _ = v33
        _ = v34
    }
    if false {
        var v35 int64 = 0
        _ = v35
    } else {
        for fi36 := 0; fi36 < 1; fi36++ {
            if false {
                fmt.Println(fi36)
                v19 = v19
                if (32767 >= 1) {
                    for fi37 := 0; fi37 < 7; fi37++ {
                        v18 = 42
                        _ = fi37
                    }
                    for fi38 := 0; fi38 < 12; fi38++ {
                        fmt.Println(v1, v1, v3)
                        switch (^(fi36)) {
                        case 983:
                            v3 -= (v3 + (^(v3)))
                        case 32767:
                            switch ((v19 + v19) + v19) {
                            case "foo bar":
                                var v39 uint8 = 1
                                _ = v39
                            }
                        }
                        var v40 float64 = 1e10
                        _ = fi38
                        _ = v40
                    }
                    if (128 == 0) {
                        if false {
                            v18 = v18
                        }
                        for fi41 := 0; fi41 < 5; fi41++ {
                            v42 := GenS2{F0: 42, F1: v3, F2: (v28 + 232)}
                            _ = fi41
                            _ = v42
                        }
                        var v43 int16 = 32767
                        _ = v43
                    } else {
                        var v44 uint32 = (v3 >> (((^(v28))) & 63))
                        _ = v44
                    }
                } else {
                    for fi45 := 0; fi45 < 4; fi45++ {
                        var v46 int32 = -1
                        _ = fi45
                        _ = v46
                    }
                    fmt.Println(v18)
                }
            }
            _ = fi36
        }
        v1 += (v1 - (v1 ^ (v1 / ((-954) | 1))))
    }
    v18 -= v18
    v47 := GenS3{F0: 747, F1: 2147483648, F2: -1}
    if true {
        if false {
            v3 = (^(v3))
        }
        v2[(v19 + v19)] = -1
        fmt.Println(v19, v3)
    }
    switch 698 {
    case -2147483648:
        fmt.Println(v18, v28, v1)
    case 9223372036854775807:
        for fi48 := 0; fi48 < 9; fi48++ {
            v49 := GenS2{F0: 0, F1: 2147483648, F2: (v28 / (((v28 * 9)) | 1))}
            _ = fi48
            _ = v49
        }
    }
    var v50 uint8 = 127
    var v51 float64 = 1.5
    for fi52 := 0; fi52 < 13; fi52++ {
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 bool = false
            _ = fi53
            _ = v54
        }
        v55 := []int{}
        var v56 int16 = 880
        _ = fi52
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v47
    _ = v50
    _ = v51
}

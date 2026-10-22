package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 int16, p1 uint16) uint {
    v6 := GenS1{F0: 42, F1: 2}
    var v7 int32 = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        v7 = (v7 - v7)
        var v9 bool = (true || (false || true))
        v7 -= v7
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8) bool {
    v10 := []int{}
    var v11 int64 = 103
    var v12 uint32 = 234
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 42
    var v3 int = v2
    var v4 int64 = 1
    var v5 int = 1
    delete(v1, (v3 << ((4294967295) & 63)))
    var v13 float64 = (-2.25 + 1.5)
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "foo bar" {
        case "a":
            if (2 > 1) {
                switch v5 {
                case -9223372036854775808:
                    var v15 uint = 179
                    _ = v15
                case 128:
                    v16 := GenS1{F0: 18446744073709551615, F1: 182}
                    _ = v16
                case 723:
                    var v17 float32 = 1e10
                    _ = v17
                default:
                    var v18 uint64 = 0
                    _ = v18
                }
            } else {
                fmt.Println(v3, v13)
            }
        case "\n\t":
            v5 = (v5 | v2)
        case "foo bar":
            switch (-(v3)) {
            case 127:
                var v19 uint64 = 2
                _ = v19
            case 0:
                fmt.Println(v2, v13)
            case 354:
                v20 := GenS1{F0: 0, F1: 58}
                _ = v20
            }
        }
        var v21 int64 = v4
        if (false || (373 >= 568)) {
            var v22 bool = false
            v13 = v13
            v23 := GenS1{F0: 467, F1: 181}
            _ = v22
            _ = v23
        } else {
            var v24 int32 = -1
            var v25 int64 = (v4 % (((v4 | 42)) | 1))
            v26 := GenS1{F0: 0, F1: 121}
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi14
        _ = v21
    }
    v27 := GenS1{F0: 0, F1: 1}
    v28 := GenS1{F0: 1, F1: 42}
    var v29 int16 = 984
    if (2 < 209) {
        v3 -= v5
    }
    if true {
        v30 := []int{4294967295}
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 int64 = 0
            switch "foo bar" {
            case "":
                var v33 int8 = -102
                _ = v33
            case "\n\t":
                delete(v1, (v5 << ((42) & 63)))
            case "foo bar":
                for fi34 := 0; fi34 < 17; fi34++ {
                    var v35 string = "a"
                    switch -329 {
                    case 729:
                        v36 := GenS1{F0: 691, F1: 2}
                        _ = v36
                    case -694:
                        switch (v35 + "foo bar") {
                        case "hello":
                            var v37 uint8 = 103
                            _ = v37
                        }
                    default:
                        if (!true) {
                            v35 = v35
                            _ = len(v1)
                            v3 -= v3
                            _ = len(v30)
                        } else {
                            var v38 int16 = 0
                            var v39 uint8 = 128
                            v40 := GenS1{F0: 622, F1: 255}
                            v41 := v1[128]
                            _ = v38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    }
                    v42 := GenS2{F0: 32768, F1: 42}
                    _ = fi34
                    _ = v35
                    _ = v42
                }
            default:
                var v43 int16 = (v29 - v29)
                _ = v43
            }
            _ = fi31
            _ = v32
        }
        var v44 int16 = (-(v29))
        _ = v30
        _ = v44
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            v46 := []int{v3, (-(v2)), 9223372036854775807}
            if (float32(v13) == 3.14159) {
                fmt.Println(v2, v13)
            }
            fmt.Println(v5, v2)
            _ = fi45
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        switch fi47 {
        case -917:
            v48 := GenS1{F0: 1, F1: 98}
            _ = v48
        case -1:
            if (!false) {
                var v49 bool = false
                v50 := GenS2{F0: v5, F1: 0}
                _ = v49
                _ = v50
            } else {
                if true {
                    v13 = v13
                    switch "hello" {
                    case "":
                        var v51 int = 32767
                        _ = v51
                    case "foo bar":
                        var v52 uint8 = 7
                        _ = v52
                    }
                } else {
                    v53 := GenS3{F0: 255, F1: 2}
                    var v54 uint = 1
                    _ = v53
                    _ = v54
                }
                fmt.Println(v13, v4, v29)
                var v55 uint16 = 65535
                _ = v55
            }
        case 404:
            var v56 int32 = -1
            _ = v56
        default:
            v57 := []int{}
            _ = v57
        }
        _ = fi47
    }
    v13 += 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

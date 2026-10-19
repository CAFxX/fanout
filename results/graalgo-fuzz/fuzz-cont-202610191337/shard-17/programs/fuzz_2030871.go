package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 string; F1 int32; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint = 0
    v4 := GenS1{F0: v3, F1: 127}
    var v5 uint = (v3 & 65535)
    fmt.Println(v3)
    delete(v1, 266)
    fmt.Println(v3, v5)
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: fi6}
        var v8 int64 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    if (694 > -1) {
        delete(v2, 876)
    } else {
        var v9 uint64 = 767
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if (0 != 307) {
            v11 := []int{fi10, (fi10 >> ((v5) & 63))}
            delete(v1, 65535)
            _ = v11
        }
        v12 := GenS2{F0: -9223372036854775808, F1: 101, F2: (!false)}
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(fi13)
            _ = fi13
        }
        _ = fi10
        _ = v12
    }
    v14 := GenS3{F0: -32768, F1: 1, F2: 133}
    for fi15 := 0; fi15 < 3; fi15++ {
        if false {
            if false {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 int = fi16
                    _ = fi16
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := GenS3{F0: -1, F1: 2, F2: 1}
                    _ = fi18
                    _ = v19
                }
                fmt.Println(v3)
            }
            v2[736] = 42
            var v20 bool = (!false)
            _ = v20
        } else {
            fmt.Println(fi15)
        }
        _ = fi15
    }
    fmt.Println(v3, v5)
    v21 := []int{-2147483648, -2147483648, 0, 127}
    if true {
        for fi22 := 0; fi22 < 1; fi22++ {
            switch fi22 {
            case 9223372036854775807:
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            default:
                var v23 uint64 = 4294967295
                _ = v23
            }
            var v24 uint64 = 42
            var v25 bool = (1e10 >= 1.5)
            _ = fi22
            _ = v24
            _ = v25
        }
        fmt.Println(v5)
        v5 += v5
    } else {
        var v26 uint16 = 242
        v3 = 4294967295
        var v27 bool = (42 > 2)
        _ = v26
        _ = v27
    }
    var v28 bool = (42 == 161)
    switch -536 {
    case 1:
        v29 := GenS3{F0: 42, F1: 4294967295, F2: 42}
        _ = v29
    case 193:
        if true {
            var v30 uint = (^(v5))
            if ((v28 || v28) || (true || false)) {
                if len(v21) > 0 {
                    v21[0] = -360
                }
            } else {
                v31 := GenS1{F0: v3, F1: 1}
                var v32 float64 = float64((-2.25 * 1.5))
                if v28 {
                    if len(v21) > 0 {
                        v21[0] = 42
                    }
                    v21 = append(v21, -1)
                }
                _ = v31
                _ = v32
            }
            switch "" {
            case "hello":
                var v33 int16 = -32768
                _ = v33
            }
            _ = v30
        } else {
            if (4294967295 > 2) {
                switch "\n\t" {
                case "hello":
                    v28 = v28
                default:
                    v34 := GenS4{F0: 0, F1: -1}
                    _ = v34
                }
                if (1e10 != float64(0.0)) {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        if v28 {
                            var v36 string = "foo bar"
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 20; fi37++ {
                            v38 := GenS2{F0: 42, F1: 1, F2: (v28 && (-2.25 != -2.25))}
                            fmt.Println(v5, v28)
                            v5 = 4294967295
                            fmt.Println(fi35, v28)
                            fmt.Println(fi35)
                            v5 = v5
                            _ = len(v21)
                            _ = fi37
                            _ = v38
                        }
                        v39 := GenS3{F0: 42, F1: 4294967295, F2: 2}
                        _ = fi35
                        _ = v39
                    }
                    var v40 float64 = (-2.25 - 0.1)
                    _ = v40
                }
            } else {
                _ = len(v21)
                if (v28 && v28) {
                    fmt.Println(v3, v5)
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 uint8 = 140
                        _ = fi41
                        _ = v42
                    }
                }
            }
            var v43 uint32 = 0
            _ = v43
        }
    default:
        v28 = (750 <= 1)
    }
    v28 = v28
    v44 := GenS2{F0: -724, F1: -1, F2: (2 < 65535)}
    if false {
        v45 := GenS5{F0: 664, F1: 1}
        switch 0 {
        case -734:
            var v46 int = 0
            _ = v46
        default:
            fmt.Println(v3, v3)
        }
        v5 = v3
        _ = v45
    } else {
        v5 += v5
        v47 := GenS6{F0: "hello", F1: -1, F2: 0}
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}

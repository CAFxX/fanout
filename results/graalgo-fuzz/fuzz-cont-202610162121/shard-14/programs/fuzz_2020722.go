package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 int16) int8 {
    switch ("foo bar" + ("foo bar" + "foo bar")) {
    case "":
        if (!false) {
            p1 = 1000
            v7 := []int{32767}
            _ = v7
        } else {
            p1 = (^(p1))
        }
    case "a":
        if false {
            for fi8 := 0; fi8 < 12; fi8++ {
                fmt.Println(p0, fi8, p0)
                v9 := []int{(-(fi8)), fi8, fi8, fi8}
                _ = fi8
                _ = v9
            }
            switch 4294967295 {
            case -878:
                v10 := []int{32767, 32767}
                _ = v10
            case 32767:
                var v11 int64 = 42
                _ = v11
            }
        }
    case "\n\t":
        if true {
            v12 := []int{-2147483648, 4294967295, -9223372036854775808, 42}
            switch "\n\t" {
            case "foo bar":
                p0 = p0
            case "\n\t":
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 bool = false
                    if true {
                        v15 := GenS1{F0: v14, F1: (^(p1)), F2: -1}
                        _ = v15
                    } else {
                        v12 = append(v12, (fi13 & fi13))
                        fmt.Println(fi13, v14)
                        for fi16 := 0; fi16 < 14; fi16++ {
                            var v17 int32 = 2147483647
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            case "a":
                switch -1 {
                case 42:
                    var v18 int8 = 1
                    _ = v18
                }
            default:
                var v19 string = ("hello" + "\n\t")
                _ = v19
            }
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = v12
        }
    default:
        v20 := GenS2{F0: p1, F1: 42}
        _ = v20
    }
    p1 = (p1 * (p1 << ((42) & 63)))
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 253
    var v4 uint = 370
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    var v21 uint = 4294967295
    v21 = (^(v21))
    var v22 int8 = 0
    v23 := GenS3{F0: -613, F1: 32767, F2: v3}
    var v24 int32 = 0
    if (v22 < v22) {
        var v25 uint16 = 255
        _ = v25
    }
    fmt.Println(v3, v21, v3)
    var v26 int32 = (v24 % ((v24) | 1))
    v27 := GenS4{F0: "a", F1: (v3 + (v3 * (v3 + 1)))}
    fmt.Println(v21)
    v26 = v24
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 float64 = 1.5
        _ = fi28
        _ = v29
    }
    switch 42 {
    case -9223372036854775808:
        if (2 > 1) {
            if true {
                v30 := []int{-2147483648}
                if (!(995 < 255)) {
                    for fi31 := 0; fi31 < 11; fi31++ {
                        var v32 int8 = -1
                        _ = fi31
                        _ = v32
                    }
                } else {
                    fmt.Println(v3, v21)
                }
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                _ = v30
            } else {
                v21 = (v4 << (((v21 << (((^(v4))) & 63))) & 63))
                v24 += -1
                switch 32767 {
                case 42:
                    var v33 bool = false
                    _ = v33
                case 232:
                    fmt.Println(v4, v22)
                default:
                    if false {
                        v34 := []int{1, 127}
                        switch "hello" {
                        case "":
                            var v35 float64 = 1e10
                            _ = v35
                        default:
                            v34 = append(v34, -271)
                        }
                        _ = v34
                    }
                }
            }
            v36 := GenS5{F0: -1, F1: -998}
            _ = v36
        }
    case 32767:
        if (255 < 1) {
            _ = len(v6)
            for fi37 := 0; fi37 < 17; fi37++ {
                var v38 uint = v4
                _ = fi37
                _ = v38
            }
        } else {
            for fi39 := 0; fi39 < 11; fi39++ {
                v40 := GenS6{F0: v22}
                v3 = 2
                v41 := []int{(^(fi39)), fi39, (fi39 - -1), fi39}
                _ = fi39
                _ = v40
                _ = v41
            }
            var v42 uint64 = 1
            _ = v42
        }
    default:
        for fi43 := 0; fi43 < 6; fi43++ {
            v44 := []int{-58}
            var v45 int = fi43
            v26 = v24
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    var v46 string = "foo bar"
    v4 += 0
    var v47 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

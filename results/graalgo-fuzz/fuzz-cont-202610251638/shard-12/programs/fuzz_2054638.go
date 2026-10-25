package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 bool }
type GenS5 struct { F0 int; F1 int16; F2 uint64 }
func genFn0(p0 int, p1 uint64, p2 bool) int8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: -9223372036854775808, F1: -1}
        _ = fi7
        _ = v8
    }
    p0 = int(p1)
    v9 := []int{p0}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 111
}
func genFn1(p0 string, p1 uint64, p2 int32) bool {
    fmt.Println(p1)
    var v10 int16 = -635
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return false
}

func main() {
    var v1 uint8 = 50
    v2 := make(map[int]bool)
    var v3 int16 = -1
    var v4 int16 = 42
    var v5 uint = 42
    v6 := []int{1, 42, 970}
    v11 := GenS1{F0: 42, F1: 42}
    if (-9223372036854775808 <= -54) {
        switch "\n\t" {
        case "":
            v5 = (v5 ^ uint(v1))
        case "a":
            v12 := GenS1{F0: 774, F1: -1}
            _ = v12
        case "foo bar":
            if (-100 < 100) {
                v2[-263] = (!true)
                switch "" {
                case "a":
                    var v13 uint8 = 42
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 18; fi14++ {
                        var v15 int64 = 0
                        switch (("\n\t" + "\n\t") + "\n\t") {
                        case "foo bar":
                            v16 := GenS1{F0: v15, F1: v15}
                            _ = v16
                        case "":
                            v15 = v15
                        default:
                            v15 = v15
                        }
                        v17 := GenS1{F0: (v15 ^ v15), F1: v15}
                        _ = fi14
                        _ = v15
                        _ = v17
                    }
                }
                var v18 uint32 = 65535
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 14; fi19++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi19
            }
        }
    } else {
        v5 -= 255
        var v20 uint32 = 1
        _ = v20
    }
    switch "foo bar" {
    case "hello":
        v4 = v3
    }
    v3 = v4
    if (true || (false || true)) {
        v4 = (v3 / ((v4) | 1))
        delete(v2, int(v4))
    } else {
        var v21 uint64 = 1
        if true {
            var v22 uint32 = 2147483647
            _ = len(v2)
            if (449 < -88) {
                var v23 int16 = v4
                fmt.Println(v22)
                v24 := v2[32767]
                _ = v23
                _ = v24
            }
            _ = v22
        } else {
            var v25 uint = (v5 | v5)
            for fi26 := 0; fi26 < 13; fi26++ {
                v5 -= 1
                v5 -= v5
                var v27 uint8 = (^(v1))
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        v28 := GenS2{F0: true, F1: int16(v5), F2: 1}
        _ = v21
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 17; fi29++ {
            switch 9223372036854775807 {
            case 32767:
                v3 = (v3 ^ (v4 & v3))
            case -688:
                if ((fi29 ^ 42) <= (fi29 % (((fi29 | fi29)) | 1))) {
                    v30 := GenS3{F0: 9223372036854775807}
                    _ = v30
                } else {
                    var v31 int8 = 100
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                delete(v2, 2147483647)
                v3 += v3
                _ = fi32
            }
            _ = fi29
        }
        delete(v2, 127)
    }
    v5 = 1
    v33 := GenS2{F0: true, F1: 1000, F2: 65535}
    v34 := v2[234]
    var v35 int32 = 0
    switch -842 {
    case 127:
        var v36 uint32 = 2147483648
        _ = v36
    }
    var v37 int16 = int16(v35)
    v38 := GenS4{F0: (^(v35)), F1: 42, F2: v34}
    v4 = 310
    v39 := GenS5{F0: 2147483647, F1: 42, F2: 1}
    fmt.Println(v3, v35)
    var v40 int32 = 692
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

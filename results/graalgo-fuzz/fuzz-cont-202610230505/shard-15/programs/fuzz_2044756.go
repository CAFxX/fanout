package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = 80
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 uint8 = 42
    var v5 float32 = (float32(0.1) - 1e10)
    for fi6 := 0; fi6 < 19; fi6++ {
        delete(v3, 210)
        v2[""] = (fi6 << ((18446744073709551615) & 63))
        _ = fi6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            if (127 < (-(v1))) {
                _ = len(v2)
            } else {
                var v9 uint16 = 65535
                var v10 int64 = -185
                _ = v10
                _ = v9
            }
            v1 += (v1 | v1)
            switch fi8 {
            case 2147483647:
                var v11 bool = (v1 <= (v1 & v1))
                _ = v11
            case -864:
                v12 := GenS1{F0: 1}
                _ = v12
            case -138:
                v1 += v1
            default:
                var v13 int32 = 698
                _ = v13
            }
            _ = fi8
        }
        v4 -= v4
        var v14 int16 = int16(fi7)
        _ = fi7
        _ = v14
    }
    v3[0] = -2147483648
    if false {
        v15 := v3[129]
        var v16 float32 = v5
        _ = v15
        _ = v16
    } else {
        if false {
            if false {
                v4 += (v4 - v4)
                for fi17 := 0; fi17 < 6; fi17++ {
                    switch fi17 {
                    case 4294967295:
                        v1 += v1
                    case 0:
                        var v18 int = fi17
                        _ = v18
                    }
                    v1 = (v1 & (v1 | -1))
                    v19 := GenS1{F0: 362}
                    _ = fi17
                    _ = v19
                }
                v20 := GenS1{F0: 4294967295}
                _ = v20
            }
            switch ("foo bar" + "hello") {
            case "a":
                var v21 uint32 = 65535
                _ = v21
            case "foo bar":
                fmt.Println(v1, v4, v1)
            case "":
                if true {
                    fmt.Println(v4)
                    switch "foo bar" {
                    case "":
                        fmt.Println(v4, v5, v5)
                    case "a":
                        v22 := GenS1{F0: 4294967295}
                        _ = v22
                    }
                    v2["\n\t"] = 32767
                } else {
                    var v23 int64 = 435
                    v24 := GenS2{F0: "", F1: -22}
                    for fi25 := 0; fi25 < 7; fi25++ {
                        var v26 uint = 4294967295
                        v27 := GenS2{F0: "foo bar", F1: fi25}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = v23
                    _ = v24
                }
            }
            var v28 uint32 = 4294967295
            _ = v28
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 uint32 = 2147483648
                    _ = fi30
                    _ = v31
                }
                switch fi29 {
                case 541:
                    v4 += (v4 % (((v4 / ((v4) | 1))) | 1))
                case -1:
                    for fi32 := 0; fi32 < 20; fi32++ {
                        var v33 int = (-(fi32))
                        _ = fi32
                        _ = v33
                    }
                case -9223372036854775808:
                    var v34 int8 = 12
                    _ = v34
                }
                v5 = -2.25
                _ = fi29
            }
            if ((!(!false)) || true) {
                var v35 int = -910
                _ = v35
            } else {
                for fi36 := 0; fi36 < 16; fi36++ {
                    for fi37 := 0; fi37 < 17; fi37++ {
                        v38 := []int{(fi36 >> ((65535) & 63))}
                        if true {
                            v5 = 3.14159
                        } else {
                            v4 += (v4 - v4)
                            _ = len(v3)
                            if (((42 >= -275) && ((32767 != 504) || false)) || false) {
                                v4 += v4
                            }
                        }
                        v1 = (v1 << ((479) & 63))
                        v4 = 0
                        var v39 int64 = -820
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    var v40 uint64 = 4294967295
                    _ = fi36
                    _ = v40
                }
                v1 += v1
            }
        }
        var v41 int8 = v1
        var v42 int16 = 1000
        _ = v41
        _ = v42
    }
    if true {
        v43 := GenS3{F0: v4, F1: v4, F2: v1}
        _ = v43
    }
    v1 = 0
    fmt.Println(v4, v4)
    for fi44 := 0; fi44 < 7; fi44++ {
        fmt.Println(v5, v5, fi44)
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

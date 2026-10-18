package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 uint64, p1 int8) uint32 {
    v7 := []int{9223372036854775807, 128, 0, 9223372036854775807}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 int32 = 42
    var v6 int64 = -9223372036854775808
    var v8 uint16 = 32767
    fmt.Println(v6, v6)
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: 0, F1: 493}
            var v12 uint64 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    if true {
        switch 128 {
        case -125:
            if (!(true && true)) {
                var v13 int8 = -1
                v14 := []int{127, -862, -45, int(v13)}
                var v15 bool = true
                _ = v13
                _ = v14
                _ = v15
            } else {
                var v16 float32 = (1.5 - 0.0)
                v16 = float32(float64((0.0 / v16)))
                _ = v16
            }
        case 4294967295:
            v5 += (v5 << ((823) & 63))
        case 774:
            v5 -= int32(v8)
        }
    } else {
        var v17 int8 = -53
        for fi18 := 0; fi18 < 18; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := []int{127, fi19, (^(fi18))}
                _ = fi19
                _ = v20
            }
            var v21 uint32 = 65535
            _ = fi18
            _ = v21
        }
        v22 := GenS1{F0: 255, F1: 4294967295}
        _ = v17
        _ = v22
    }
    if false {
        var v23 uint16 = 255
        for fi24 := 0; fi24 < 15; fi24++ {
            switch fi24 {
            case 4294967295:
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS1{F0: 125, F1: 314}
                    v2["a"] = 71
                    switch fi24 {
                    case -630:
                        if false {
                            v1[(fi25 >> ((65535) & 63))] = (^(fi24))
                            if true {
                                if (!false) {
                                    if ((v6 << ((65535) & 63)) <= v6) {
                                        var v27 int8 = -1
                                        v28 := []int{fi25, fi24, 128, fi24}
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                var v29 int16 = 1
                                _ = v29
                            }
                        } else {
                            v30 := GenS2{F0: 2, F1: (!true), F2: 969}
                            _ = v30
                        }
                    }
                    _ = fi25
                    _ = v26
                }
            case 32768:
                v31 := []int{(-(fi24)), 42, fi24}
                _ = v31
            default:
                for fi32 := 0; fi32 < 15; fi32++ {
                    switch ("foo bar" + "a") {
                    case "a":
                        if true {
                            v33 := []int{(fi24 | fi24), (fi24 / ((fi24) | 1))}
                            var v34 int8 = -1
                            _ = v33
                            _ = v34
                        }
                    default:
                        v35 := GenS2{F0: 354, F1: (!(421 == uint64(v8))), F2: 915}
                        _ = v35
                    }
                    _ = fi32
                }
            }
            switch fi24 {
            case 423:
                v6 -= v6
            case 42:
                var v36 int32 = v5
                _ = v36
            default:
                var v37 uint8 = 42
                _ = v37
            }
            v6 -= (v6 - (v6 ^ -9223372036854775808))
            _ = fi24
        }
        for fi38 := 0; fi38 < 18; fi38++ {
            v5 += -2147483648
            v39 := GenS2{F0: 18446744073709551615, F1: (128 != 1), F2: 2}
            _ = fi38
            _ = v39
        }
        _ = v23
    }
    v8 += v8
    if (!((0 <= 128) && ((0.0 * (0.0 * 1.5)) > float64(1.5)))) {
        var v40 uint64 = 2
        switch "foo bar" {
        case "foo bar":
            v6 = 816
        case "":
            if true {
                v8 = 65535
            }
        }
        _ = v40
    } else {
        for fi41 := 0; fi41 < 13; fi41++ {
            var v42 bool = (1 >= 1)
            _ = fi41
            _ = v42
        }
    }
    var v43 uint = 4294967295
    v6 += v6
    fmt.Println(v8)
    var v44 float64 = 3.14159
    v5 -= v5
    switch "" {
    case "foo bar":
        fmt.Println(v8, v43)
    default:
        for fi45 := 0; fi45 < 11; fi45++ {
            v44 += (((((v44 - ((v44 + v44) / (float64(1.5) * 0.1))) - 1.5) + 0.1) - v44) * v44)
            var v46 int32 = v5
            if true {
                v5 = -471
                switch "\n\t" {
                case "hello":
                    var v47 string = "hello"
                    _ = v47
                default:
                    var v48 uint8 = 255
                    _ = v48
                }
            } else {
                var v49 uint32 = 2147483647
                _ = v49
            }
            _ = fi45
            _ = v46
        }
    }
    v50 := GenS3{F0: false, F1: 42, F2: v43}
    var v51 float32 = -2.25
    var v52 uint = v43
    for fi53 := 0; fi53 < 20; fi53++ {
        fmt.Println(v52)
        var v54 string = (("hello" + "\n\t") + "\n\t")
        _ = fi53
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v8
}

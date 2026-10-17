package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint; F1 int16 }
func genFn0(p0 uint) uint32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 0}
        p0 += 62
        _ = fi7
        _ = v8
    }
    if true {
        if true {
            p0 = p0
            fmt.Println(p0)
            p0 = (p0 * p0)
        } else {
            v9 := GenS2{F0: true, F1: 2}
            switch "foo bar" {
            case "a":
                v10 := GenS3{F0: 2}
                _ = v10
            case "foo bar":
                for fi11 := 0; fi11 < 18; fi11++ {
                    fmt.Println(fi11, p0)
                    _ = fi11
                }
            case "hello":
                v12 := GenS3{F0: 0}
                _ = v12
            default:
                v13 := []int{}
                _ = v13
            }
            var v14 uint16 = 429
            _ = v14
            _ = v9
        }
        if (9223372036854775807 == 127) {
            v15 := GenS2{F0: (0 <= 0), F1: 1}
            for fi16 := 0; fi16 < 3; fi16++ {
                p0 = p0
                var v17 uint32 = 65535
                _ = fi16
                _ = v17
            }
            switch 675 {
            case 32768:
                if true {
                    p0 = (p0 >> (((p0 % ((p0) | 1))) & 63))
                }
            case -9223372036854775808:
                fmt.Println(p0)
            case -1:
                v18 := []int{-97, 32768, 128}
                _ = v18
            }
            _ = v15
        } else {
            var v19 int64 = -1
            v19 = -758
            _ = v19
        }
        if (!true) {
            p0 = p0
        } else {
            if false {
                if (468 < 2147483647) {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        if false {
                            var v21 uint64 = 646
                            var v22 uint64 = v21
                            _ = v21
                            _ = v22
                        }
                        if (1 < 46) {
                            v23 := []int{}
                            p0 += 2
                            if true {
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                fmt.Println(fi20)
                            }
                            _ = v23
                        }
                        var v24 int16 = 775
                        _ = fi20
                        _ = v24
                    }
                } else {
                    v25 := GenS4{F0: -95}
                    _ = v25
                }
            }
        }
    }
    if (680 > 42) {
        if false {
            if true {
                v26 := GenS3{F0: 18446744073709551615}
                for fi27 := 0; fi27 < 1; fi27++ {
                    fmt.Println(fi27, p0)
                    _ = fi27
                }
                _ = v26
            } else {
                v28 := GenS5{F0: 255, F1: true}
                var v29 int8 = -1
                for fi30 := 0; fi30 < 15; fi30++ {
                    v29 += v29
                    _ = fi30
                }
                _ = v28
                _ = v29
            }
            var v31 uint16 = 666
            _ = v31
        }
        if (1.5 >= 3.14159) {
            v32 := []int{735, 996, 510}
            _ = v32
        } else {
            p0 -= 42
        }
        p0 = (p0 & 255)
    } else {
        p0 += (^(p0))
        var v33 int32 = -2147483648
        _ = v33
    }
    switch 447 {
    case 128:
        var v34 bool = (-417 == 180)
        _ = v34
    case 32768:
        v35 := []int{9223372036854775807, 127}
        _ = v35
    }
    _ = p0
    return 2147483648
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 18
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    if (197 > int16(v1)) {
        var v36 bool = false
        if false {
            var v37 uint16 = 920
            _ = v37
        } else {
            switch "a" {
            case "":
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 bool = ((!v36) || ((v4 / (((^(v4))) | 1)) == (v4 * (^(v4)))))
                    _ = fi38
                    _ = v39
                }
            case "a":
                v40 := GenS5{F0: 65535, F1: v36}
                _ = v40
            }
            v41 := GenS5{F0: uint16(v1), F1: (!false)}
            v42 := GenS1{F0: 1}
            _ = v41
            _ = v42
        }
        _ = v36
    }
    v43 := GenS5{F0: 65535, F1: true}
    v1 += (v1 | (v1 << ((0) & 63)))
    v44 := GenS2{F0: true, F1: 670}
    v1 = v1
    switch "hello" {
    case "hello":
        v45 := []int{-9223372036854775808, 9223372036854775807, 9223372036854775807, 1000}
        _ = v45
    case "a":
        if false {
            v46 := GenS5{F0: 892, F1: (!false)}
            if false {
                v1 += (v1 / (((^(v1))) | 1))
            } else {
                v47 := []int{0, int(v4)}
                for fi48 := 0; fi48 < 4; fi48++ {
                    var v49 int8 = -54
                    _ = fi48
                    _ = v49
                }
                _ = v47
            }
            _ = v46
        } else {
            v50 := []int{}
            v4 = genFn0(uint(v4))
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        v52 := GenS2{F0: (true && (9223372036854775807 != fi51)), F1: 790}
        v4 = uint32(v1)
        _ = fi51
        _ = v52
    }
    v1 = v1
    v4 -= 2147483648
    v4 = (v4 >> ((624) & 63))
    if true {
        var v53 uint8 = 127
        if true {
            v4 = (v4 / ((uint32(v1)) | 1))
            var v54 int64 = -333
            v54 -= (v54 ^ (v54 & (v54 & 42)))
            _ = v54
        }
        _ = v53
    } else {
        for fi55 := 0; fi55 < 10; fi55++ {
            v1 += 1
            v56 := v6[""]
            v57 := GenS6{F0: 0, F1: 0}
            _ = fi55
            _ = v56
            _ = v57
        }
        fmt.Println(v4, v4)
    }
    switch "a" {
    case "foo bar":
        var v58 uint32 = 1
        _ = v58
    case "a":
        v4 = 42
    default:
        v59 := GenS4{F0: -128}
        _ = v59
    }
    v60 := GenS1{F0: 0}
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v60
}

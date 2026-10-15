package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0(p0 uint, p1 int64, p2 uint) int8 {
    v4 := []int{32767, -198, 32768, -362}
    v5 := GenS1{F0: -614, F1: false, F2: 886}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]bool)
    v3 := []int{4294967295}
    v2[-9223372036854775808] = (false && false)
    if (true || (800 >= 2147483648)) {
        for fi6 := 0; fi6 < 11; fi6++ {
            if true {
                if true {
                    v1 += (v1 & 115)
                }
                var v7 uint = 42
                _ = v7
            }
            v1 += v1
            _ = fi6
        }
    } else {
        var v8 string = (("a" + "hello") + "hello")
        var v9 int8 = 1
        _ = v8
        _ = v9
    }
    fmt.Println(v1)
    for fi10 := 0; fi10 < 3; fi10++ {
        _ = len(v2)
        _ = fi10
    }
    v11 := GenS1{F0: -687, F1: (-2147483648 > -1), F2: 835}
    v1 = (v1 >> ((150) & 63))
    v1 = v1
    for fi12 := 0; fi12 < 17; fi12++ {
        switch fi12 {
        case 0:
            var v13 uint8 = 163
            _ = v13
        default:
            var v14 uint16 = (v1 << ((1) & 63))
            _ = v14
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            v1 -= v1
            var v16 uint32 = 2
            _ = fi15
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS1{F0: 32767, F1: (1 >= -9223372036854775808), F2: -2147483648}
    if (42 != -32768) {
        if false {
            v1 = (v1 / ((v1) | 1))
            var v18 int32 = 474
            if true {
                var v19 bool = false
                for fi20 := 0; fi20 < 10; fi20++ {
                    _ = len(v3)
                    for fi21 := 0; fi21 < 5; fi21++ {
                        fmt.Println(v19, fi20)
                        v22 := GenS1{F0: int16(fi20), F1: true, F2: (^(v18))}
                        switch "hello" {
                        case "":
                            var v23 uint64 = 568
                            _ = v23
                        case "foo bar":
                            var v24 int = fi20
                            _ = v24
                        case "\n\t":
                            if v19 {
                                var v25 int32 = v18
                                _ = v25
                            }
                        default:
                            var v26 int32 = (v18 + (v18 + v18))
                            _ = v26
                        }
                        _ = fi21
                        _ = v22
                    }
                    _ = fi20
                }
                _ = v19
            }
            _ = v18
        }
        v1 = v1
    }
    v1 -= v1
    for fi27 := 0; fi27 < 17; fi27++ {
        fmt.Println(fi27, v1)
        switch fi27 {
        case 4294967295:
            v28 := GenS1{F0: 42, F1: (!false), F2: int32(fi27)}
            _ = v28
        case -392:
            v29 := GenS1{F0: 880, F1: false, F2: 42}
            _ = v29
        case -9223372036854775808:
            v1 = v1
        }
        _ = fi27
    }
    fmt.Println(v1)
    for fi30 := 0; fi30 < 7; fi30++ {
        fmt.Println(v1, v1)
        v31 := GenS2{F0: fi30, F1: v1, F2: 262}
        v1 -= v1
        _ = fi30
        _ = v31
    }
    if false {
        fmt.Println(v1)
        var v32 int16 = int16(v1)
        v1 = 0
        _ = v32
    } else {
        if ("\n\t" != "a") {
            var v33 float32 = 1.5
            switch "\n\t" {
            case "a":
                if false {
                    v34 := v2[-9223372036854775808]
                    fmt.Println(v1, v1)
                    _ = v34
                }
            }
            _ = v33
        } else {
            if (-530 < 9223372036854775807) {
                v2[9223372036854775807] = (589 >= 1)
                for fi35 := 0; fi35 < 6; fi35++ {
                    v1 = (v1 / ((65535) | 1))
                    v1 -= (^(v1))
                    v1 = 65535
                    _ = fi35
                }
            }
            v36 := v2[1]
            var v37 bool = (v36 || ("" >= ""))
            _ = v36
            _ = v37
        }
        var v38 bool = false
        _ = v38
    }
    if true {
        v1 += (v1 + 32767)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
}

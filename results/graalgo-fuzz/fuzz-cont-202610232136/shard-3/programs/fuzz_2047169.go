package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint8) string {
    p0 -= (p0 - (p0 << ((18446744073709551615) & 63)))
    p0 += (p0 / (((p0 + p0)) | 1))
    if (!false) {
        v6 := []int{685}
        _ = v6
    }
    _ = p0
    return ""
}
func genFn1() uint8 {
    var v7 bool = (!false)
    var v8 int = 4294967295
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 2
    v3 := []uint32{v2, (v2 >> ((261) & 63)), 1}
    v4 := make(map[uint32]uint32)
    var v5 int32 = 878
    v5 = (v5 >> ((18446744073709551615) & 63))
    if (v2 <= (^(v2))) {
        v4[(v2 % ((v2) | 1))] = v2
        if false {
            v10 := v4[v2]
            var v11 uint32 = (v10 & (v10 / ((v10) | 1)))
            _ = v10
            _ = v11
        } else {
            fmt.Println(v1, v2)
            for fi12 := 0; fi12 < 18; fi12++ {
                v4[884] = (v2 & 2147483648)
                v1 += 1e10
                _ = fi12
            }
        }
    }
    v13 := GenS1{F0: 903, F1: (v2 / ((v2) | 1)), F2: 296}
    fmt.Println(v2, v1, v5)
    v5 = 2147483647
    v14 := GenS1{F0: 42, F1: 4294967295, F2: 42}
    v15 := GenS2{F0: 255, F1: -128}
    for fi16 := 0; fi16 < 1; fi16++ {
        _ = len(v4)
        _ = fi16
    }
    v17 := v4[4294967295]
    var v18 int = 744
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint16 = 2
        _ = fi19
        _ = v20
    }
    if false {
        fmt.Println(v18, v17, v18)
        if (!(((v5 + v5) > v5) && true)) {
            if true {
                fmt.Println(v2)
            }
            if ((!true) || (!true)) {
                v4[832] = v2
                var v21 int = (v18 & v18)
                _ = v21
            }
        } else {
            v2 = (v17 | 193)
            if true {
                v1 -= ((-2.25 - -2.25) - (v1 - v1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                for fi22 := 0; fi22 < 8; fi22++ {
                    switch 127 {
                    case -9223372036854775808:
                        v18 -= (fi22 % ((32768) | 1))
                    case -538:
                        v23 := GenS1{F0: 1, F1: v17, F2: 482}
                        _ = v23
                    case -560:
                        var v24 string = "\n\t"
                        _ = v24
                    default:
                        v25 := GenS2{F0: 33, F1: 42}
                        _ = v25
                    }
                    _ = fi22
                }
            }
            var v26 uint = 122
            _ = v26
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := v4[v2]
            _ = fi27
            _ = v28
        }
    } else {
        v29 := GenS1{F0: 397, F1: (v2 + 2147483647), F2: 0}
        switch -415 {
        case -1:
            if false {
                var v30 uint16 = 255
                _ = v30
            }
        default:
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 bool = true
                _ = len(v4)
                v18 = v18
                _ = fi31
                _ = v32
            }
        }
        switch (("hello" + "\n\t") + (("a" + ("" + "")) + ("" + (("\n\t" + "hello") + "a")))) {
        case "":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v18, v18)
            case "hello":
                delete(v4, v2)
            case "a":
                v5 = (v5 ^ v5)
            default:
                var v33 int64 = 0
                _ = v33
            }
        case "foo bar":
            v34 := GenS2{F0: 279, F1: 42}
            _ = v34
        }
        _ = v29
    }
    if (400 == 326) {
        for fi35 := 0; fi35 < 4; fi35++ {
            v2 -= (v17 * (v17 | v17))
            var v36 uint = 716
            _ = fi35
            _ = v36
        }
        var v37 uint64 = 42
        delete(v4, (^(v17)))
        _ = v37
    } else {
        fmt.Println(v5, v18)
        fmt.Println(v5, v2)
    }
    v2 = v2
    _ = len(v3)
    var v38 uint32 = (v17 % (((v17 % ((v2) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}

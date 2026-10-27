package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0() uint {
    var v7 uint = 65535
    fmt.Println(v7)
    fmt.Println(v7)
    v7 = 2
    v8 := []int{32767, 42, 42}
    _ = v7
    _ = v8
    return 651
}
func genFn1(p0 uint8, p1 int8) uint16 {
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            p0 += 1
            if false {
                var v11 float64 = 1.5
                var v12 uint = 42
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(p1, p0)
    switch int(p1) {
    case 42:
        var v13 float64 = -2.25
        _ = v13
    case -9223372036854775808:
        v14 := GenS1{F0: 263, F1: -163, F2: (^(p0))}
        _ = v14
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 uint32 = 783
    v2 := []int{-2147483648}
    var v3 uint32 = 2147483647
    v4 := []int{}
    var v5 int16 = -32768
    v6 := []int{4294967295, -9223372036854775808}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 = append(v4, 9223372036854775807)
    v15 := GenS1{F0: 2, F1: 42, F2: 0}
    for fi16 := 0; fi16 < 6; fi16++ {
        if (4294967295 >= 805) {
            _ = len(v2)
            v5 += (-(v5))
            for fi17 := 0; fi17 < 7; fi17++ {
                v5 = (-(v5))
                _ = fi17
            }
        }
        v5 -= (v5 << ((genFn0()) & 63))
        var v18 int32 = -1
        _ = fi16
        _ = v18
    }
    var v19 bool = true
    if false {
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS1{F0: 18446744073709551615, F1: -1, F2: 107}
            for fi22 := 0; fi22 < 17; fi22++ {
                v1 += v1
                _ = fi22
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi20
            _ = v21
        }
    } else {
        v23 := GenS2{F0: genFn1(2, -128), F1: -1, F2: (^(v3))}
        _ = v23
    }
    v24 := GenS3{F0: ("hello" + ("\n\t" + "a"))}
    var v25 uint32 = (v3 << ((2) & 63))
    if (v19 || ("a" <= "hello")) {
        var v26 float64 = 1.5
        v27 := GenS3{F0: "a"}
        v2 = append(v2, -922)
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 int8, p1 string, p2 int64) int8 {
    v6 := GenS1{F0: 1}
    p2 = (p2 / ((9223372036854775807) | 1))
    var v7 uint32 = 42
    var v8 int16 = 0
    var v9 int64 = (p2 % ((558) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1() uint32 {
    v10 := []int{2147483647}
    var v11 bool = false
    var v12 int32 = 1
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS2{F0: v11}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    _ = v10
    _ = v11
    _ = v12
    return 891
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 uint64 = 1
    var v4 int16 = 825
    v5 := []int{}
    for fi16 := 0; fi16 < 7; fi16++ {
        v4 = (v4 ^ v4)
        _ = fi16
    }
    v17 := GenS2{F0: (v3 > (v3 << ((2) & 63)))}
    fmt.Println(v4, v3)
    for fi18 := 0; fi18 < 20; fi18++ {
        v3 = (^(v3))
        var v19 int32 = 781
        _ = fi18
        _ = v19
    }
    _ = len(v5)
    v3 += v3
    if (!(("\n\t" + "a") <= ("foo bar" + "\n\t"))) {
        fmt.Println(v3)
    } else {
        var v20 int64 = 9223372036854775807
        _ = v20
    }
    v4 += -900
    if true {
        if true {
            var v21 uint64 = (^(v3))
            v22 := v1[("a" + ("hello" + ("\n\t" + (("a" + "") + ""))))]
            if true {
                v21 = v21
                fmt.Println(v3)
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS3{F0: 1, F1: v3, F2: 2147483647}
                    _ = fi23
                    _ = v24
                }
            }
            _ = v21
            _ = v22
        }
        var v25 bool = (4294967295 > 84)
        _ = v25
    } else {
        v4 = v4
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v4, v4)
    }
    var v26 uint16 = 255
    if false {
        v27 := GenS4{F0: (1 > -1)}
        v26 -= 131
        v28 := GenS5{F0: 1, F1: (((!true) || ((false && (!true)) || (2 < 2))) || (false && (!true))), F2: -24}
        _ = v27
        _ = v28
    }
    switch ("a" + "") {
    case "\n\t":
        var v29 uint8 = 186
        _ = v29
    }
    v30 := GenS3{F0: genFn1(), F1: v3, F2: 462}
    v26 = (v26 << ((2) & 63))
    v26 += (v26 & (^(v26)))
    var v31 uint16 = (v26 % ((v26) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

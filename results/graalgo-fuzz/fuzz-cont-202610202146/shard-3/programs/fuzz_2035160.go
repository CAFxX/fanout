package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
func genFn0() int64 {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int16 = -1
            v10 := []int{fi8}
            switch "hello" {
            case "\n\t":
                v9 = 42
            case "hello":
                v9 = (v9 << ((2) & 63))
            case "":
                fmt.Println(fi7)
            }
            v9 = (v9 | 406)
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int32 = 719
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: 2, F1: -867, F2: 702}
            if (!false) {
                var v14 int64 = 1
                fmt.Println(v14, fi7)
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        _ = fi7
        _ = v11
    }
    var v15 float32 = 1.5
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint32 = 4294967295
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS1{F0: 176, F1: 42, F2: 0}
            _ = fi19
            _ = v20
        }
        var v21 bool = (!(42 < 4294967295))
        v22 := []int{fi18}
        _ = fi18
        _ = v21
        _ = v22
    }
    _ = v15
    return -9223372036854775808
}
func genFn1() uint32 {
    v23 := GenS1{F0: 2, F1: -1, F2: 539}
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := []int{(fi24 / ((fi24) | 1))}
        _ = fi24
        _ = v25
    }
    _ = v23
    return 478
}

func main() {
    var v1 int32 = -1
    v2 := []int{127}
    var v3 int8 = -128
    var v4 int8 = (v3 << ((4294967295) & 63))
    var v5 uint64 = 1
    var v6 uint32 = 350
    var v26 int64 = -9223372036854775808
    v27 := GenS1{F0: 2, F1: (v26 - (v26 & (v26 - -9223372036854775808))), F2: 18446744073709551615}
    if len(v2) > 0 {
        v2[0] = 128
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v4, v26)
    case "":
        if true {
            v26 = 42
        } else {
            var v28 uint = 1
            switch "foo bar" {
            case "\n\t":
                v29 := GenS1{F0: v28, F1: v26, F2: 18446744073709551615}
                _ = v29
            case "a":
                var v30 bool = (genFn0() > (v26 | genFn0()))
                _ = v30
            }
            for fi31 := 0; fi31 < 12; fi31++ {
                var v32 int64 = -9223372036854775808
                var v33 uint = 249
                fmt.Println(v1, v33)
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v28
        }
    case "hello":
        for fi34 := 0; fi34 < 8; fi34++ {
            v26 += (v26 % (((v26 / ((-440) | 1))) | 1))
            v6 = (v6 & 2147483647)
            _ = fi34
        }
    default:
        var v35 int16 = 42
        _ = v35
    }
    v36 := GenS1{F0: 117, F1: v26, F2: 18446744073709551615}
    var v37 int64 = 962
    for fi38 := 0; fi38 < 4; fi38++ {
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS1{F0: 255, F1: (v26 >> ((uint(v5)) & 63)), F2: uint(v5)}
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v4 += v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

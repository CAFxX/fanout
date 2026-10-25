package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 2147483647
    var v3 int64 = -314
    var v4 int8 = -128
    var v5 string = ("" + "\n\t")
    v2 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 127, F1: 65535, F2: 4294967295}
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint16 = 866
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: 65535, F1: 0}
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 uint = 0
    v12 := GenS2{F0: v11, F1: 0}
    if false {
        fmt.Println(v11)
    } else {
        var v13 uint8 = 1
        if (true && (!false)) {
            v14 := []int{942, -9223372036854775808, -9223372036854775808, -662}
            v15 := GenS2{F0: v11, F1: 1}
            _ = v14
            _ = v15
        }
        _ = v13
    }
    fmt.Println(v3)
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int32 = int32(v3)
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(fi18, v5, fi18)
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := []int{(fi19 - 4294967295), fi19, fi19, int(v11)}
        _ = fi19
        _ = v20
    }
    v3 = (v3 + (v3 / ((v3) | 1)))
    v11 = v11
    fmt.Println(v4)
    var v21 int64 = v3
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS2{F0: v11, F1: 2}
        switch (fi22 / ((fi22) | 1)) {
        case 0:
            if ((^(v4)) < v4) {
                for fi24 := 0; fi24 < 17; fi24++ {
                    v11 = 0
                    _ = fi24
                }
            }
        case 127:
            v3 = (^(v3))
        case -2147483648:
            var v25 uint = v11
            _ = v25
        default:
            v26 := GenS3{F0: 255, F1: (((v4 - v4) > (v4 / ((v4) | 1))) && true), F2: false}
            _ = v26
        }
        v4 = (^(v4))
        _ = fi22
        _ = v23
    }
    v11 = 872
    fmt.Println(v2, v2, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

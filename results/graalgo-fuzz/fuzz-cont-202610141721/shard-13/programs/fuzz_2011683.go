package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 2
    var v2 uint8 = 194
    var v3 uint64 = 0
    var v4 int32 = 1
    v4 = (v4 & v4)
    switch 879 {
    case -740:
        v5 := []int{}
        _ = v5
    case -41:
        fmt.Println(v4)
    default:
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(v3, fi6, v4)
            v1 = v1
            for fi7 := 0; fi7 < 8; fi7++ {
                v3 = 0
                if false {
                    v3 -= v3
                    v8 := GenS1{F0: fi6, F1: 134, F2: 0}
                    v4 += -1
                    _ = v8
                } else {
                    v9 := GenS1{F0: fi7, F1: 212, F2: -32768}
                    v10 := GenS2{F0: fi7}
                    v11 := []int{(fi7 & fi6), fi7}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v12 := []int{}
                _ = fi7
                _ = v12
            }
            _ = fi6
        }
    }
    v4 = -187
    v13 := []int{-919}
    v1 = v1
    fmt.Println(v2, v3)
    var v14 int8 = -49
    for fi15 := 0; fi15 < 15; fi15++ {
        if len(v13) > 0 {
            v13[0] = (fi15 | fi15)
        }
        v16 := GenS3{F0: "a"}
        _ = fi15
        _ = v16
    }
    if true {
        v4 = (v4 + v4)
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v17 int32 = v4
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS3{F0: "a"}
        switch fi18 {
        case 904:
            fmt.Println(v2, v17)
        case 0:
            v14 = -1
        default:
            v3 += v3
        }
        _ = fi18
        _ = v19
    }
    var v20 int64 = -9223372036854775808
    v17 += v17
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            switch "foo bar" {
            case "foo bar":
                v23 := GenS1{F0: int(v17), F1: -249, F2: int16(v14)}
                _ = v23
            case "":
                if false {
                    v3 = (v3 >> ((v1) & 63))
                    var v24 uint32 = 937
                    v2 += uint8(v3)
                    fmt.Println(fi21)
                    _ = v24
                }
            }
            v13 = append(v13, fi22)
            _ = fi22
        }
        v25 := GenS1{F0: fi21, F1: 42, F2: 962}
        _ = fi21
        _ = v25
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int32 }
func genFn0() int {
    if (-1 >= 510) {
        switch 128 {
        case 1:
            v6 := GenS1{F0: 18446744073709551615}
            _ = v6
        case 622:
            var v7 uint32 = 1
            _ = v7
        default:
            v8 := []int{704, 4294967295}
            _ = v8
        }
        var v9 string = "foo bar"
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int64 = 1
            switch fi10 {
            case 32767:
                switch "foo bar" {
                case "":
                    v12 := GenS1{F0: 1}
                    _ = v12
                case "\n\t":
                    v9 = v9
                case "hello":
                    fmt.Println(v11, v11, fi10)
                default:
                    v11 = (v11 + -9223372036854775808)
                }
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                v11 = (v11 << ((990) & 63))
                fmt.Println(v9, fi10)
                v14 := []int{}
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        v15 := GenS2{F0: -669, F1: 255, F2: 255}
        _ = v15
    }
    v16 := GenS3{F0: 2147483647}
    v17 := []int{17, 402}
    _ = v16
    _ = v17
    return -9223372036854775808
}
func genFn1(p0 int32, p1 int8, p2 int) uint {
    p1 = p1
    var v18 int8 = p1
    if false {
        fmt.Println(v18, v18, v18)
    } else {
        v19 := []int{(p2 ^ 4294967295), p2, p2, genFn0()}
        switch genFn0() {
        case 923:
            if false {
                p2 += (p2 / (((-(p2))) | 1))
                if (true || false) {
                    var v20 uint8 = 255
                    _ = v20
                }
                v21 := GenS2{F0: -501, F1: 0, F2: 42}
                _ = v21
            }
        case 42:
            var v22 int16 = 0
            _ = v22
        default:
            var v23 int16 = -724
            _ = v23
        }
        p2 = (p2 - genFn0())
        _ = v19
    }
    fmt.Println(p1)
    p0 += (p0 | -394)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 4294967295
}

func main() {
    var v1 int32 = -1
    var v2 int64 = -779
    var v3 uint32 = 42
    var v4 uint32 = v3
    var v5 int16 = int16(v1)
    var v24 uint32 = 2
    switch -977 {
    case 2147483647:
        var v25 int8 = -1
        _ = v25
    case 128:
        v2 = int64(v5)
    }
    v24 += 284
    v26 := GenS3{F0: v1}
    v24 -= 965
    for fi27 := 0; fi27 < 17; fi27++ {
        v24 = 327
        v28 := []int{fi27, fi27, fi27}
        _ = fi27
        _ = v28
    }
    var v29 int16 = v5
    switch 15 {
    case 4294967295:
        v30 := GenS2{F0: v2, F1: 83, F2: 229}
        _ = v30
    case -2147483648:
        v31 := GenS1{F0: 1}
        _ = v31
    case -1:
        v32 := []int{}
        _ = v32
    }
    var v33 string = "foo bar"
    v4 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}

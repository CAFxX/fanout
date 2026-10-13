package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 bool; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint64) int8 {
    switch 815 {
    case 989:
        var v6 bool = true
        _ = v6
    case 32768:
        switch -181 {
        case -366:
            var v7 int32 = 1
            _ = v7
        case 9223372036854775807:
            p0 += (p0 % ((75) | 1))
        case -2147483648:
            for fi8 := 0; fi8 < 16; fi8++ {
                p0 += p0
                v9 := GenS1{F0: 65, F1: 18446744073709551615, F2: (p0 / (((p0 ^ p0)) | 1))}
                _ = fi8
                _ = v9
            }
        }
    }
    if false {
        v10 := []int{-2147483648, 888}
        _ = len(v10)
        _ = v10
    } else {
        v11 := []int{-741, -469, 128}
        var v12 float64 = 1.5
        for fi13 := 0; fi13 < 12; fi13++ {
            v11 = append(v11, -1)
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    v14 := GenS1{F0: 608, F1: 276, F2: (p0 >> (((p0 + (p0 % ((p0) | 1)))) & 63))}
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := []int{}
        var v17 uint8 = 2
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v14
    return 28
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]string)
    var v3 bool = (("foo bar" <= "foo bar") || false)
    var v4 string = "\n\t"
    var v5 uint16 = 74
    var v18 float64 = 0.1
    var v19 uint64 = 4294967295
    var v20 uint = 984
    fmt.Println(v20, v18, v5)
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := []int{}
        _ = fi21
        _ = v22
    }
    var v23 int16 = 526
    v19 += (v19 >> (((^(v20))) & 63))
    for fi24 := 0; fi24 < 13; fi24++ {
        v23 = (v23 << (((v20 % ((4294967295) | 1))) & 63))
        switch ("hello" + "hello") {
        case "a":
            fmt.Println(v23, v5)
        }
        _ = fi24
    }
    if (42 == -1) {
        if v3 {
            v25 := GenS1{F0: uint32(v19), F1: (v20 + v20), F2: (v20 >> ((0) & 63))}
            v26 := []int{21, 2147483647, 2147483647}
            var v27 uint16 = (v5 << ((4294967295) & 63))
            _ = v25
            _ = v26
            _ = v27
        } else {
            fmt.Println(v19)
        }
        if ((-128 < 67) && (genFn0(v20, v19) < 0)) {
            v19 = v19
        }
        if false {
            v28 := GenS2{F0: -2147483648, F1: v3, F2: v20}
            _ = v28
        } else {
            switch "hello" {
            case "a":
                var v29 int32 = 2147483647
                _ = v29
            case "":
                v30 := v2[-547]
                _ = v30
            case "hello":
                switch 128 {
                case 32768:
                    if true {
                        var v31 int16 = v23
                        var v32 string = ("foo bar" + v4)
                        v33 := []int{0, 2147483647, 309}
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case 4294967295:
                    for fi34 := 0; fi34 < 8; fi34++ {
                        fmt.Println(v18, v18)
                        v3 = (fi34 > fi34)
                        _ = fi34
                    }
                }
            }
            var v35 int64 = int64(v19)
            _ = v35
        }
    } else {
        var v36 int8 = 42
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        v5 = (^(v5))
        v3 = true
        v18 += 1e10
        _ = fi37
    }
    fmt.Println(v18)
    v38 := GenS3{F0: 42}
    var v39 int16 = -266
    v40 := GenS3{F0: (v5 << (((v20 | (v20 + v20))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint16; F1 int64 }
func genFn0() uint16 {
    switch 127 {
    case 127:
        if false {
            var v7 uint8 = 127
            v8 := []int{-241, 956, -2147483648, 42}
            _ = v7
            _ = v8
        } else {
            var v9 uint32 = 949
            v10 := []int{9223372036854775807, 128, -1, -1}
            _ = v10
            _ = v9
        }
    case 75:
        var v11 float32 = -2.25
        _ = v11
    }
    switch 1 {
    case 32767:
        if false {
            switch "\n\t" {
            case "":
                switch -9223372036854775808 {
                case -9223372036854775808:
                    v12 := GenS1{F0: -9223372036854775808}
                    _ = v12
                case 0:
                    var v13 uint16 = 184
                    _ = v13
                case -1:
                    v14 := []int{127, 690, -680, 849}
                    _ = v14
                }
            case "foo bar":
                var v15 float32 = 0.1
                _ = v15
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(fi16)
                var v17 uint16 = 1
                _ = fi16
                _ = v17
            }
        }
    case 257:
        v18 := []int{0, -9223372036854775808, 42}
        _ = v18
    }
    switch "hello" {
    case "hello":
        var v19 int32 = -1
        _ = v19
    case "a":
        if (-1 != 1) {
            v20 := []int{}
            _ = v20
        } else {
            v21 := GenS2{F0: 886, F1: 128, F2: 87}
            var v22 string = "hello"
            var v23 int32 = -252
            _ = v21
            _ = v22
            _ = v23
        }
    }
    var v24 uint32 = 2147483648
    for fi25 := 0; fi25 < 4; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            switch fi26 {
            case -293:
                fmt.Println(fi26)
                v27 := GenS3{F0: -910}
                _ = v27
            }
            v28 := GenS4{F0: (v24 & 0)}
            _ = fi26
            _ = v28
        }
        _ = fi25
    }
    _ = v24
    return 32767
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    var v3 uint64 = 418
    var v4 uint8 = 2
    v5 := make(map[int]uint32)
    var v6 uint = 4294967295
    var v29 float32 = 0.0
    fmt.Println(v1, v4, v1)
    v1 += v1
    if false {
        v1 -= -280
        fmt.Println(v3, v6, v6)
        v30 := []int{190, 880, 302, 605}
        _ = v30
    }
    switch "hello" {
    case "":
        fmt.Println(v6, v1, v3)
    }
    v4 = v4
    var v31 uint16 = 746
    v32 := GenS5{F0: ("\n\t" < (("hello" + "a") + "hello"))}
    v33 := GenS6{F0: (!true), F1: 42}
    for fi34 := 0; fi34 < 1; fi34++ {
        v3 += v3
        fmt.Println(v29, v1)
        if false {
            v6 -= (v6 ^ v6)
        } else {
            if false {
                var v35 uint16 = (v31 + (v31 * v31))
                fmt.Println(v31, v3, v1)
                _ = v35
            } else {
                fmt.Println(fi34, v29, v1)
                if (-1 > 0) {
                    v36 := []int{}
                    _ = v36
                }
                fmt.Println(v6, v31)
            }
            v37 := GenS7{F0: (v31 - (v31 >> (((v6 | 2)) & 63))), F1: -87}
            var v38 uint = 365
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

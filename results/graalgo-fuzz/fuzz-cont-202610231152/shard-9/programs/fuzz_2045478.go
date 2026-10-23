package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 bool }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 string = "a"
    var v4 float32 = (0.1 * 1e10)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int32 = 1
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 42
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 1000
        v1[-376] = 2147483648
        var v10 uint8 = 161
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{117}
    if true {
        if true {
            var v12 float32 = float32(0.0)
            if (!(("foo bar" + v3) < ((v3 + "\n\t") + v3))) {
                v13 := GenS1{F0: (^(v7)), F1: 127, F2: 942}
                var v14 string = (v3 + "")
                v15 := v2[int(v7)]
                _ = v13
                _ = v14
                _ = v15
            }
            switch -1 {
            case -434:
                switch 901 {
                case 0:
                    v16 := GenS1{F0: (^(v7)), F1: 42, F2: -2147483648}
                    _ = v16
                }
            }
            _ = v12
        } else {
            for fi17 := 0; fi17 < 11; fi17++ {
                switch (v3 + "") {
                case "hello":
                    v2[(fi17 % ((fi17) | 1))] = v3
                case "foo bar":
                    var v18 int = -473
                    _ = v18
                case "a":
                    v19 := GenS2{F0: -1, F1: -97, F2: 65535}
                    _ = v19
                }
                _ = fi17
            }
        }
        var v20 uint16 = 2
        _ = v20
    }
    v3 = (v3 + (((("foo bar" + "a") + v3) + "a") + v3))
    switch 2147483647 {
    case 127:
        var v21 string = v3
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(v3)
        _ = fi22
    }
    if false {
        var v23 int = 32768
        _ = v23
    } else {
        if true {
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := GenS1{F0: v7, F1: 127, F2: 1}
                _ = fi24
                _ = v25
            }
            fmt.Println(v3, v3, v4)
        }
        v4 -= v4
        switch -747 {
        case 42:
            fmt.Println(v4, v3)
        default:
            v4 -= v4
        }
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 bool = ((4294967295 <= 65535) && true)
        v3 = v3
        _ = fi26
        _ = v27
    }
    if true {
        fmt.Println(v7, v3)
    } else {
        v28 := GenS3{F0: true, F1: -1, F2: (1 != 100)}
        if false {
            var v29 uint = 2
            var v30 int32 = 2147483647
            fmt.Println(v3)
            _ = v29
            _ = v30
        } else {
            v3 = ("" + v3)
            _ = len(v11)
        }
        _ = v28
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v7 = 4294967295
        _ = fi31
    }
    if false {
        for fi32 := 0; fi32 < 2; fi32++ {
            if true {
                fmt.Println(v3, v7)
                fmt.Println(v3)
            }
            _ = fi32
        }
        v33 := GenS4{F0: 432, F1: -819}
        v7 -= v7
        _ = v33
    } else {
        v7 = (v7 >> ((18446744073709551615) & 63))
        v4 = (v4 / v4)
    }
    v7 += v7
    v34 := v2[128]
    switch 32767 {
    case 1:
        switch 32767 {
        case -1:
            var v35 int16 = 1
            _ = v35
        case -144:
            switch 374 {
            case -9223372036854775808:
                var v36 int = -588
                _ = v36
            case 127:
                v7 = (v7 % ((v7) | 1))
            case -973:
                v7 = v7
            }
        case 32767:
            v37 := GenS5{F0: 577}
            _ = v37
        }
    default:
        v38 := GenS6{F0: 0, F1: 65535, F2: 371}
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v7
}

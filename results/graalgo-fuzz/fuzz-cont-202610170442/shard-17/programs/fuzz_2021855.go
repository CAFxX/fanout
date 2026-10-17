package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    if (true || ((3.14159 + (((float64(1.5) + 3.14159) / -2.25) / 0.1)) != (0.0 / 0.1))) {
        v4 := GenS1{F0: 0, F1: -32768, F2: -32768}
        switch "hello" {
        case "a":
            var v5 int8 = -128
            _ = v5
        }
        v6 := []int{4294967295}
        _ = v4
        _ = v6
    }
    var v7 int32 = 42
    v8 := GenS1{F0: 1, F1: 1000, F2: -32768}
    var v9 int32 = (v7 - -942)
    if true {
        var v10 uint8 = 60
        _ = v10
    } else {
        v7 += v7
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int32 = (-(v7))
            if true {
                if false {
                    fmt.Println(v9, v12)
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v9 += v12
                        _ = fi13
                    }
                } else {
                    var v14 uint64 = 42
                    var v15 int32 = (v12 % ((v12) | 1))
                    var v16 uint8 = 184
                    _ = v14
                    _ = v15
                    _ = v16
                }
            } else {
                v12 -= v7
                fmt.Println(v9, v12, v9)
            }
            _ = fi11
            _ = v12
        }
        var v17 uint32 = 42
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v7 -= -2147483648
        v7 += v9
        _ = fi18
    }
    fmt.Println(v9)
    switch 1 {
    case 128:
        v9 -= -221
    default:
        v9 = v9
    }
    v19 := GenS1{F0: 240, F1: 42, F2: 42}
    switch "a" {
    case "":
        switch "hello" {
        case "foo bar":
            v1[4294967295] = 2147483647
        default:
            if ((v7 % ((v7) | 1)) > v9) {
                v9 -= (v7 >> ((42) & 63))
                v20 := GenS2{F0: -1, F1: 4294967295}
                _ = v20
            } else {
                var v21 uint = 1
                if true {
                    v21 = (v21 & (v21 & v21))
                    v22 := GenS3{F0: 927, F1: -1, F2: 32767}
                    v21 = v21
                    _ = v22
                } else {
                    fmt.Println(v21)
                    var v23 int64 = 0
                    _ = v23
                }
                _ = v21
            }
        }
    case "foo bar":
        v24 := []int{-2147483648, 127}
        _ = v24
    case "hello":
        v25 := GenS1{F0: 42, F1: 0, F2: -682}
        _ = v25
    default:
        var v26 int16 = 319
        _ = v26
    }
    fmt.Println(v9)
    if (1 <= 18446744073709551615) {
        v9 -= v9
        var v27 uint = 18446744073709551615
        if true {
            switch "" {
            case "\n\t":
                for fi28 := 0; fi28 < 1; fi28++ {
                    v7 += (v9 % ((v7) | 1))
                    v29 := GenS1{F0: 127, F1: 32767, F2: 42}
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        v27 += (v27 ^ v27)
                    case "a":
                        v30 := []int{fi28, fi28}
                        _ = v30
                    case "\n\t":
                        var v31 uint32 = 0
                        _ = v31
                    default:
                        v32 := []int{}
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                }
            case "hello":
                fmt.Println(v9, v27, v7)
            }
        } else {
            for fi33 := 0; fi33 < 1; fi33++ {
                v7 += v7
                if true {
                    v34 := v2[fi33]
                    _ = v34
                }
                fmt.Println(fi33, v27, v7)
                _ = fi33
            }
        }
        _ = v27
    }
    switch "" {
    case "foo bar":
        v7 = v7
    case "\n\t":
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 string = ("hello" + "\n\t")
            switch "a" {
            case "foo bar":
                v36 = v36
            default:
                v37 := GenS2{F0: 1, F1: 255}
                _ = v37
            }
            var v38 bool = true
            _ = fi35
            _ = v36
            _ = v38
        }
    case "":
        v39 := GenS2{F0: 11, F1: 42}
        _ = v39
    default:
        var v40 int32 = v7
        _ = v40
    }
    switch 468 {
    case 42:
        fmt.Println(v9, v9)
    case -9223372036854775808:
        for fi41 := 0; fi41 < 8; fi41++ {
            v42 := GenS2{F0: 0, F1: 42}
            if (!true) {
                v7 -= (v7 - -657)
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 string = ""
                    var v45 uint32 = 2147483648
                    var v46 uint8 = 1
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                for fi47 := 0; fi47 < 10; fi47++ {
                    v48 := GenS4{F0: 42}
                    var v49 int8 = 1
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            }
            fmt.Println(v7, v7, v7)
            _ = fi41
            _ = v42
        }
    default:
        fmt.Println(v7, v7)
    }
    v50 := GenS2{F0: -1, F1: 18446744073709551615}
    v51 := GenS5{F0: 42}
    switch 2147483647 {
    case 127:
        v52 := v2[4294967295]
        _ = v52
    case -1:
        v7 = v9
    case 32768:
        switch "" {
        case "a":
            var v53 uint8 = 250
            _ = v53
        }
    }
    var v54 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v54
    _ = v7
    _ = v8
    _ = v9
}

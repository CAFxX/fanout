package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int32; F1 int; F2 int64 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int }
func genFn0(p0 string) uint {
    fmt.Println(p0)
    v4 := GenS1{F0: 42}
    p0 = p0
    switch ("a" + p0) {
    case "\n\t":
        p0 = "a"
    case "foo bar":
        fmt.Println(p0)
    case "":
        for fi5 := 0; fi5 < 18; fi5++ {
            p0 = p0
            if (!false) {
                if ((fi5 - fi5) < fi5) {
                    for fi6 := 0; fi6 < 10; fi6++ {
                        fmt.Println(fi5, p0)
                        fmt.Println(p0, fi6, p0)
                        _ = fi6
                    }
                    fmt.Println(p0, fi5)
                    var v7 int32 = 0
                    _ = v7
                }
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := GenS2{F0: 65535, F1: true, F2: 42}
                    v10 := []int{}
                    v10 = append(v10, (fi5 - (fi8 / ((127) | 1))))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                p0 = p0
            } else {
                var v11 float32 = 1.5
                _ = v11
            }
            v12 := GenS1{F0: 166}
            _ = fi5
            _ = v12
        }
    }
    _ = p0
    _ = v4
    return 4294967295
}

func main() {
    var v1 int32 = -1
    v2 := []int{-2147483648, -1, 9223372036854775807, 304}
    var v3 int16 = 1000
    v13 := GenS3{F0: 0}
    v3 += -111
    var v14 uint32 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 32768 {
    case 1:
        switch 625 {
        case 879:
            switch -9223372036854775808 {
            case 9223372036854775807:
                if false {
                    var v15 uint = 729
                    var v16 int8 = 59
                    _ = v15
                    _ = v16
                } else {
                    v1 = -2147483648
                }
            }
        default:
            v14 -= v14
        }
    case 0:
        v1 += v1
    default:
        var v17 uint64 = 0
        _ = v17
    }
    v18 := GenS3{F0: (v14 | v14)}
    for fi19 := 0; fi19 < 11; fi19++ {
        if ((true && (0 <= v14)) && true) {
            if len(v2) > 0 {
                v2[0] = (fi19 >> ((uint(v14)) & 63))
            }
            var v20 int8 = -1
            _ = v20
        }
        v3 -= 42
        if (!true) {
            fmt.Println(v3)
            v21 := GenS4{F0: 119, F1: true}
            v1 = v1
            _ = v21
        } else {
            fmt.Println(v3, fi19, fi19)
            v22 := GenS3{F0: 65535}
            _ = v22
        }
        _ = fi19
    }
    fmt.Println(v3, v3, v3)
    var v23 uint32 = v14
    switch "hello" {
    case "a":
        var v24 int16 = v3
        _ = v24
    case "\n\t":
        if (v1 <= 708) {
            switch ("" + ("hello" + "foo bar")) {
            case "\n\t":
                v25 := GenS5{F0: 42, F1: -2147483648}
                _ = v25
            case "hello":
                var v26 uint = genFn0("\n\t")
                _ = v26
            case "foo bar":
                var v27 float32 = ((1e10 / 1.0) + 0.0)
                _ = v27
            }
            var v28 uint32 = (v14 % ((v14) | 1))
            fmt.Println(v14, v3)
            _ = v28
        } else {
            v29 := GenS3{F0: v14}
            _ = v29
        }
    case "foo bar":
        v30 := GenS6{F0: v1, F1: 0, F2: 0}
        _ = v30
    }
    fmt.Println(v14)
    var v31 bool = ((0 >= 127) || true)
    v3 = -785
    fmt.Println(v3)
    switch 32767 {
    case 2147483647:
        for fi32 := 0; fi32 < 9; fi32++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch 551 {
            case 1:
                if v31 {
                    fmt.Println(fi32, fi32)
                    v33 := GenS2{F0: 4294967295, F1: (0 <= 255), F2: v3}
                    _ = v33
                } else {
                    v34 := GenS7{F0: (v3 - -874), F1: v3}
                    _ = v34
                }
            case 2147483647:
                var v35 int16 = (v3 | 32767)
                _ = v35
            case 128:
                v31 = v31
            }
            v36 := GenS8{F0: fi32}
            _ = fi32
            _ = v36
        }
    case 1:
        fmt.Println(v31, v23, v1)
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v3 = v3
        _ = fi37
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v38 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v38
}

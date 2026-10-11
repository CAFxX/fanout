package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 bool; F1 bool; F2 int64 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint64 = 710
    var v2 int64 = 285
    var v3 int = 128
    if true {
        var v4 int16 = 1000
        v5 := []int{v3, (v3 & v3), int(v2)}
        switch (v3 << ((2) & 63)) {
        case -1:
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 uint32 = 947
                v8 := GenS1{F0: v2, F1: -1}
                _ = fi6
                _ = v7
                _ = v8
            }
        default:
            var v9 int64 = (v2 << ((649) & 63))
            _ = v9
        }
        _ = v4
        _ = v5
    } else {
        switch ("hello" + "a") {
        case "a":
            v1 -= v1
        default:
            var v10 uint16 = 2
            _ = v10
        }
        v11 := []int{}
        fmt.Println(v3, v1)
        _ = v11
    }
    v2 = v2
    switch "foo bar" {
    case "hello":
        var v12 float64 = (1e10 + 1.5)
        _ = v12
    case "a":
        v13 := []int{}
        _ = v13
    default:
        var v14 bool = (false && false)
        _ = v14
    }
    var v15 string = ""
    var v16 bool = true
    v17 := GenS2{F0: v2, F1: uint16(v1)}
    fmt.Println(v2)
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint64 = (v1 / ((v1) | 1))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        switch 1 {
        case -289:
            var v21 uint16 = 432
            _ = v21
        default:
            if v16 {
                var v22 uint8 = 203
                for fi23 := 0; fi23 < 17; fi23++ {
                    v15 = v15
                    _ = fi23
                }
                var v24 int64 = (v2 + (v2 >> ((309) & 63)))
                _ = v22
                _ = v24
            }
        }
        switch -345 {
        case 4294967295:
            v16 = v16
        }
        _ = fi20
    }
    switch v3 {
    case -264:
        v15 = (("\n\t" + v15) + "foo bar")
    case 4294967295:
        v25 := []int{}
        _ = v25
    default:
        if (1 != 1) {
            var v26 int8 = int8(v2)
            _ = v26
        }
    }
    fmt.Println(v1, v2)
    v27 := GenS3{F0: 100, F1: 32767, F2: v1}
    for fi28 := 0; fi28 < 20; fi28++ {
        switch v3 {
        case 56:
            v29 := []int{(v3 >> ((2) & 63)), fi28, 337, (fi28 << ((18446744073709551615) & 63))}
            _ = v29
        case 128:
            v30 := []int{fi28, -9223372036854775808, 32768, (v3 >> ((674) & 63))}
            _ = v30
        }
        _ = fi28
    }
    switch -9223372036854775808 {
    case 0:
        var v31 int32 = -254
        _ = v31
    }
    fmt.Println(v2)
    switch 1 {
    case 32768:
        v32 := GenS4{F0: v16, F1: v16, F2: 0}
        _ = v32
    default:
        for fi33 := 0; fi33 < 2; fi33++ {
            v34 := GenS3{F0: -73, F1: 9223372036854775807, F2: 363}
            switch (v3 + 535) {
            case 127:
                if (897 < 233) {
                    v35 := []int{}
                    _ = v35
                }
            case 2147483647:
                for fi36 := 0; fi36 < 8; fi36++ {
                    var v37 int = v3
                    v38 := GenS5{F0: v2}
                    v37 += (fi36 / (((-(v37))) | 1))
                    v3 = (v37 - (fi36 * (fi36 * v37)))
                    if v16 {
                        fmt.Println(v16)
                    }
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
            fmt.Println(v3, v1, v1)
            _ = fi33
            _ = v34
        }
    }
    v16 = ((v3 + (^(v3))) < (v3 | v3))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
}

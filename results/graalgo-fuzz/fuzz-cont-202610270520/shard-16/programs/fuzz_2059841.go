package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint32 = 1
    var v2 int16 = -1
    var v3 uint64 = 229
    var v4 int = 2147483647
    var v5 string = "hello"
    var v6 int8 = 85
    var v7 int16 = v2
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{v4, (v4 | v4), (v4 | (^(fi8))), fi8}
        switch 904 {
        case 4294967295:
            var v10 uint64 = (v3 % ((291) | 1))
            _ = v10
        }
        v11 := GenS1{F0: 0, F1: 24}
        _ = fi8
        _ = v11
        _ = v9
    }
    v12 := []int{890, v4}
    if ((1e10 / -2.25) > 3.14159) {
        switch "a" {
        case "\n\t":
            v4 -= -2147483648
        case "foo bar":
            if (!false) {
                v7 += v2
            }
        }
        var v13 uint16 = 225
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v13
    }
    if (true || (!(v1 < v1))) {
        var v14 uint32 = v1
        v5 = v5
        _ = v14
    }
    var v15 bool = true
    v15 = (255 < 0)
    if (!(v2 == -1)) {
        v16 := GenS2{F0: 1}
        if (false && (v2 >= (v2 + int16(v4)))) {
            v2 = (v2 >> ((38) & 63))
            v6 = v6
            for fi17 := 0; fi17 < 12; fi17++ {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS2{F0: 1}
                    v20 := GenS3{F0: v3, F1: 0, F2: uint16(v1)}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi17
            }
        } else {
            if true {
                v15 = v15
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 uint64 = (v3 % (((v3 >> ((2) & 63))) | 1))
                    _ = fi21
                    _ = v22
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        }
        _ = v16
    }
    _ = len(v12)
    v23 := GenS4{F0: 42}
    var v24 uint32 = v1
    if (-573 == -9223372036854775808) {
        if true {
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 int16 = (v2 << ((65535) & 63))
                var v27 int16 = 1
                v28 := GenS5{F0: -2147483648}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            fmt.Println(v3, v24)
            fmt.Println(v4, v3)
        }
        if v15 {
            if v15 {
                fmt.Println(v6, v7)
            }
            var v29 uint8 = 1
            _ = v29
        }
        switch int(v7) {
        case -537:
            var v30 string = "\n\t"
            _ = v30
        default:
            _ = len(v12)
        }
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 string = "\n\t"
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            v2 -= v7
            v5 = v5
            _ = fi34
        }
        _ = fi33
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v1)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

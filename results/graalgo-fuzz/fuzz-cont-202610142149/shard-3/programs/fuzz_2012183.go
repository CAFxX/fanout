package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 int, p1 int, p2 int8) uint32 {
    fmt.Println(p1)
    if (!false) {
        v4 := []int{-498, p1, p1, p1}
        var v5 uint = 0
        p1 += 127
        _ = v4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int8) uint32 {
    var v6 string = ("a" + "a")
    v6 = (v6 + v6)
    v7 := GenS1{F0: 65535}
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 0
    v3 := []int{-2147483648, 2147483647}
    fmt.Println(v2)
    switch 32768 {
    case 661:
        v2 -= (v2 & v2)
    case 4294967295:
        if (1 != -1) {
            v2 -= 1
            var v8 int = -913
            if false {
                if (525 <= 241) {
                    var v9 int32 = 1
                    var v10 uint16 = 32767
                    v11 := GenS2{F0: 8, F1: 127}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int = 0
                    _ = len(v1)
                    v2 -= (v2 + (v2 << ((18446744073709551615) & 63)))
                    _ = fi12
                    _ = v13
                }
            }
            _ = v8
        } else {
            _ = len(v1)
            fmt.Println(v2)
        }
    case 127:
        v14 := GenS3{F0: 630}
        _ = v14
    default:
        for fi15 := 0; fi15 < 7; fi15++ {
            v2 = v2
            v1[213] = (false && (1 == 42))
            _ = fi15
        }
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS4{F0: 127, F1: "\n\t"}
        _ = fi16
        _ = v17
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(v2)
        var v19 uint32 = v2
        v2 = genFn0((fi18 >> ((580) & 63)), fi18, 42)
        _ = fi18
        _ = v19
    }
    var v20 string = "\n\t"
    v20 = (v20 + (v20 + "a"))
    var v21 uint16 = 634
    v2 = v2
    v22 := GenS3{F0: 0}
    v23 := GenS4{F0: -1, F1: v20}
    switch v20 {
    case "\n\t":
        v24 := GenS5{F0: 123, F1: (v2 << ((255) & 63))}
        _ = v24
    }
    var v25 string = v20
    for fi26 := 0; fi26 < 8; fi26++ {
        switch (fi26 << ((65535) & 63)) {
        case 42:
            if len(v3) > 0 {
                v3[0] = fi26
            }
        case -1:
            var v27 int = -2147483648
            _ = v27
        case 127:
            switch (-(fi26)) {
            case 42:
                if (!(true && true)) {
                    var v28 int64 = 9223372036854775807
                    _ = v28
                } else {
                    if false {
                        var v29 bool = (9223372036854775807 <= 0)
                        v25 = (v25 + v20)
                        v30 := v1[uint32(v21)]
                        _ = v29
                        _ = v30
                    }
                }
            case 426:
                v2 = v2
            case -9223372036854775808:
                v31 := GenS1{F0: 380}
                _ = v31
            default:
                fmt.Println(v2)
            }
        default:
            for fi32 := 0; fi32 < 15; fi32++ {
                v33 := GenS6{F0: 224, F1: 160, F2: (fi26 - (-(fi32)))}
                _ = fi32
                _ = v33
            }
        }
        v20 = "a"
        for fi34 := 0; fi34 < 6; fi34++ {
            v20 = (("\n\t" + v25) + v25)
            var v35 string = "\n\t"
            switch (((v20 + v35) + v25) + ((v35 + "\n\t") + v35)) {
            case "\n\t":
                v36 := GenS1{F0: 18446744073709551615}
                _ = v36
            case "hello":
                v1[303] = (42 < -749)
            default:
                v2 -= v2
            }
            _ = fi34
            _ = v35
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}

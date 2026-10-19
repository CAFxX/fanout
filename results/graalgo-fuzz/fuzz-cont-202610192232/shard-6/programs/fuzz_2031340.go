package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint64) bool {
    if true {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "foo bar":
            p0 += (p0 + (p0 >> ((329) & 63)))
        case "\n\t":
            var v4 string = "foo bar"
            _ = v4
        case "":
            p0 -= 0
        }
        var v5 int64 = -1
        p0 += p0
        _ = v5
    } else {
        p0 = (^(p0))
        p0 = (^(p0))
        var v6 bool = false
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    return ((p0 == (p0 ^ (p0 + (p0 & (^(p0)))))) && true)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 99
    var v3 uint64 = 18446744073709551615
    v3 += v3
    if ((127 >= v2) || (856 <= 1)) {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = 0
            var v9 uint = 834
            _ = len(v1)
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 string = "\n\t"
        _ = v10
    }
    switch 2147483647 {
    case 2147483647:
        for fi11 := 0; fi11 < 12; fi11++ {
            if false {
                var v12 int = fi11
                switch 462 {
                case 32767:
                    fmt.Println(v2, v3)
                case 128:
                    v2 = (v2 / ((v2) | 1))
                }
                if true {
                    fmt.Println(v12, v2)
                } else {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        fmt.Println(v12, fi13)
                        _ = fi13
                    }
                    v12 -= (^(v12))
                }
                _ = v12
            }
            v14 := GenS1{F0: (^(v2)), F1: 431, F2: (18446744073709551615 <= (v3 | 18446744073709551615))}
            if (65535 >= 858) {
                fmt.Println(v3, v2, fi11)
                v15 := GenS1{F0: 83, F1: 372, F2: (42 == 1)}
                v3 = v3
                _ = v15
            }
            _ = fi11
            _ = v14
        }
    case -517:
        switch "a" {
        case "a":
            v2 -= v2
        case "hello":
            var v16 int16 = 0
            _ = v16
        }
    case 353:
        if (2 <= 1) {
            fmt.Println(v3, v2)
            fmt.Println(v2, v2)
        }
    }
    delete(v1, 1)
    v2 -= (^(v2))
    v17 := GenS1{F0: v2, F1: 32767, F2: false}
    v3 = 1
    var v18 int32 = -2147483648
    if (-922 != -896) {
        v3 = v3
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(fi20)
                _ = fi20
            }
            v3 += v3
            _ = fi19
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS2{F0: (v3 & v3), F1: 912}
            switch 127 {
            case 1:
                delete(v1, 112)
            case -9223372036854775808:
                v23 := GenS3{F0: -9223372036854775808, F1: v3}
                _ = v23
            case 2147483647:
                var v24 uint16 = 65535
                _ = v24
            }
            _ = fi21
            _ = v22
        }
    }
    var v25 uint = 709
    v25 = v25
    switch 352 {
    case -951:
        v2 -= (^(v2))
    default:
        v26 := GenS3{F0: -9223372036854775808, F1: v3}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint) string {
    var v4 int = 4294967295
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 197, F1: 42}
        p0 = p0
        _ = fi5
        _ = v6
    }
    if true {
        var v7 int32 = 2147483647
        v4 = (v4 / ((4294967295) | 1))
        _ = v7
    } else {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: 42, F1: -864}
                p0 -= (p0 % ((p0) | 1))
                _ = fi8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            fmt.Println(v4)
            switch "foo bar" {
            case "\n\t":
                var v11 int = -706
                _ = v11
            case "":
                p0 = p0
            }
            v12 := GenS1{F0: -1, F1: 1}
            _ = fi10
            _ = v12
        }
    }
    v13 := GenS1{F0: 32767, F1: 42}
    if false {
        var v14 float64 = 3.14159
        v15 := []int{}
        p0 = p0
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v13
    _ = v4
    return ("\n\t" + ("foo bar" + ("hello" + "")))
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = -2147483648
    var v3 bool = (("\n\t" != ("foo bar" + "foo bar")) && (false && true))
    v16 := []int{32768, 9223372036854775807, 128, 32768}
    var v17 bool = v3
    v2 = (v2 ^ v2)
    if false {
        fmt.Println(v1, v3)
        if (!v17) {
            fmt.Println(v2, v17)
            v3 = ((v2 != v2) && v17)
        }
        _ = len(v16)
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if false {
        switch "hello" {
        case "\n\t":
            _ = len(v16)
        case "":
            for fi18 := 0; fi18 < 12; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := GenS2{F0: 65535}
                    v21 := GenS1{F0: 32767, F1: int32(v1)}
                    fmt.Println(fi19, v2, fi18)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                v2 += 0
                for fi22 := 0; fi22 < 8; fi22++ {
                    var v23 uint8 = 127
                    var v24 int64 = -9223372036854775808
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi18
            }
        case "hello":
            v2 += (v2 ^ (v2 * v2))
        default:
            v25 := GenS3{F0: 1}
            _ = v25
        }
    }
    fmt.Println(v1, v1, v2)
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            switch (-(fi26)) {
            case 2147483647:
                v1 -= (v1 % ((v1) | 1))
            }
            var v28 bool = (170 < 32767)
            fmt.Println(v1, v3, v28)
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
}

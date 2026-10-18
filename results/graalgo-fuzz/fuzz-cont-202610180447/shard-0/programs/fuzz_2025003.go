package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint64 = 816
    var v2 uint32 = 2147483647
    var v3 int8 = -7
    var v4 uint64 = (v1 << ((141) & 63))
    var v5 int16 = int16(v4)
    v4 -= (v1 + v4)
    var v6 uint16 = 1
    if true {
        switch 9223372036854775807 {
        case 1:
            var v7 int16 = (v5 ^ 1000)
            _ = v7
        }
        fmt.Println(v5)
        for fi8 := 0; fi8 < 18; fi8++ {
            switch 127 {
            case 1:
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := []int{(fi8 - (fi9 % ((32768) | 1))), (fi9 * 4294967295)}
                    v3 = 42
                    switch (fi9 % ((303) | 1)) {
                    case 127:
                        v11 := GenS1{F0: 255, F1: 297}
                        _ = v11
                    case 893:
                        v2 += (v2 ^ v2)
                    case -765:
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        switch ("foo bar" + "a") {
                        case "":
                            v2 += v2
                        case "foo bar":
                            v12 := GenS1{F0: uint8(v2), F1: 1}
                            _ = v12
                        }
                        _ = len(v10)
                    }
                    _ = len(v10)
                    var v13 uint64 = v1
                    var v14 float64 = 0.1
                    fmt.Println(v2, fi9)
                    fmt.Println(v13)
                    var v15 int = 4294967295
                    v4 -= (v4 >> ((643) & 63))
                    _ = fi9
                    _ = v10
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case -274:
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 uint32 = v2
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi8
        }
    } else {
        v3 = (v3 * -112)
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 int8 = v3
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v2 = v2
        _ = fi20
    }
    v6 = 255
    fmt.Println(v1)
    v3 = (v3 % (((v3 & v3)) | 1))
    v2 = (v2 + v2)
    switch "a" {
    case "hello":
        v1 -= v4
    case "":
        v21 := GenS1{F0: 128, F1: 18446744073709551615}
        _ = v21
    case "\n\t":
        v4 = (v4 + (v4 & (^(v4))))
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

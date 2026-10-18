package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 uint32) uint {
    v5 := GenS1{F0: uint64(p1)}
    for fi6 := 0; fi6 < 5; fi6++ {
        switch "" {
        case "foo bar":
            p1 = p1
        }
        _ = fi6
    }
    v7 := []int{(p0 >> ((18446744073709551615) & 63)), 0, (p0 ^ (p0 - 1)), p0}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 255
}

func main() {
    var v1 uint = 568
    var v2 uint16 = 692
    v3 := []string{"a", ("\n\t" + "hello"), ""}
    var v4 int64 = 59
    fmt.Println(v1)
    fmt.Println(v4)
    v4 = (v4 * v4)
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v2 += (v2 << (((v1 % (((^(v1))) | 1))) & 63))
            var v10 uint16 = (^(v2))
            _ = fi9
            _ = v10
        }
        fmt.Println(v1, v1, v4)
        _ = fi8
    }
    if false {
        if (127 == -128) {
            v1 += (v1 | (v1 / ((v1) | 1)))
        }
        v4 = v4
    } else {
        v2 -= v2
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(v1, v4, fi11)
            switch (fi11 + -901) {
            case 51:
                switch fi11 {
                case 9223372036854775807:
                    for fi12 := 0; fi12 < 2; fi12++ {
                        fmt.Println(fi11)
                        switch fi11 {
                        case 802:
                            fmt.Println(v4)
                        case 32767:
                            v2 -= (v2 << ((v1) & 63))
                        case -495:
                            v2 = (v2 >> ((v1) & 63))
                        }
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = fi12
                    }
                default:
                    v2 = 65535
                }
            case -2147483648:
                switch fi11 {
                case 1:
                    var v13 uint = genFn0(fi11, 42)
                    _ = v13
                default:
                    var v14 uint = v1
                    _ = v14
                }
            default:
                v2 = (v2 << (((^(v1))) & 63))
            }
            v4 += (v4 | (v4 << ((2) & 63)))
            _ = fi11
        }
        var v15 bool = false
        _ = v15
    }
    var v16 int = 32768
    v17 := GenS2{F0: 0}
    v18 := GenS3{F0: 0, F1: (v16 ^ v16)}
    v2 = 1
    var v19 bool = false
    v20 := GenS1{F0: 566}
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint8 = 2
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

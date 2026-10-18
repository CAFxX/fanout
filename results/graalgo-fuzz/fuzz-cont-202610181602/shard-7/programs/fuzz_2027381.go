package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int64, p1 uint64) int8 {
    v5 := GenS1{F0: 548}
    p1 -= 1
    v6 := []int{-468, -529}
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = 99
            _ = fi8
            _ = v9
        }
        fmt.Println(p1, p1, p1)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 42
    v2 := []string{"foo bar", ("foo bar" + (("hello" + (("" + "a") + ((("a" + "a") + "hello") + ("hello" + (("a" + ("\n\t" + "foo bar")) + ""))))) + ("hello" + "a"))), "hello"}
    v3 := []int{127, -2147483648, 2147483647, 0}
    var v4 uint16 = 1
    for fi10 := 0; fi10 < 20; fi10++ {
        v1 = v1
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int16 = 1000
            switch 1 {
            case 128:
                v12 -= (v12 & v12)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "" {
                case "":
                    v12 -= (v12 | (v12 - (v12 | v12)))
                case "\n\t":
                    v13 := GenS1{F0: 65535}
                    _ = v13
                case "a":
                    if ((fi11 / ((fi11) | 1)) < -9223372036854775808) {
                        v3 = append(v3, fi10)
                        if (v12 == -32768) {
                            if len(v2) > 0 {
                                v2[0] = ((("\n\t" + "foo bar") + "") + "a")
                            }
                        } else {
                            v4 = (v4 * 42)
                        }
                    } else {
                        fmt.Println(fi11, v4)
                    }
                }
            case 0:
                v1 = (v1 + (v1 | v1))
            }
            _ = fi11
            _ = v12
        }
        v14 := GenS1{F0: 0}
        _ = fi10
        _ = v14
    }
    v4 -= (v4 << ((255) & 63))
    if true {
        v15 := GenS2{F0: 436, F1: 407}
        _ = v15
    }
    v16 := GenS2{F0: 369, F1: 32767}
    v4 += 2
    switch "hello" {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 32768
        }
    case "":
        v17 := GenS1{F0: 2}
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v4)
        var v19 int16 = 42
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 18446744073709551615}
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int16 = 1
        v23 := GenS2{F0: 267, F1: (-(v22))}
        fmt.Println(v22, v22, v22)
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        switch 851 {
        case 4294967295:
            var v25 int = 2147483647
            _ = v25
        case -886:
            var v26 int = 32767
            _ = v26
        default:
            if len(v3) > 0 {
                v3[0] = (fi24 / ((fi24) | 1))
            }
        }
        _ = fi24
    }
    if (!false) {
        var v27 int = 127
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v27
    } else {
        v1 = 42
    }
    if false {
        var v28 uint32 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        v1 = v1
        _ = fi29
    }
    fmt.Println(v1, v4)
    var v30 int = 1
    v31 := GenS3{F0: 0, F1: 34}
    v30 -= (v30 >> ((1) & 63))
    v30 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}

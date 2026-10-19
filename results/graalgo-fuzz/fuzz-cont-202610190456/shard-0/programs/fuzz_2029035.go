package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int32; F1 uint32; F2 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 65535
    v3 := []int{}
    v4 := []uint32{1, 42, 11}
    var v5 float32 = 0.1
    v6 := []int{}
    fmt.Println(v1)
    switch (((("\n\t" + (("" + "") + "hello")) + ("\n\t" + "a")) + "a") + "\n\t") {
    case "hello":
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: uint64(v2), F1: 2, F2: ("a" + ("foo bar" + "\n\t"))}
                _ = fi8
                _ = v9
            }
            switch fi7 {
            case 128:
                var v10 int = (fi7 | fi7)
                _ = v10
            case 4294967295:
                v2 = v2
            default:
                v11 := GenS2{F0: 226}
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 17; fi12++ {
                    switch "foo bar" {
                    case "\n\t":
                        v2 -= (v2 ^ 0)
                    }
                    v13 := GenS3{F0: (v1 * 42), F1: 2147483648}
                    _ = fi12
                    _ = v13
                }
            }
            _ = fi7
        }
    case "a":
        v14 := GenS2{F0: 18}
        _ = v14
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 float32 = float32(0.1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch -212 {
        case -462:
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v16, v1, v16)
                if (0 == 789) {
                    var v18 int64 = (v1 << ((42) & 63))
                    _ = v18
                } else {
                    v16 = 0.1
                    switch fi17 {
                    case 142:
                        v6 = append(v6, (^(fi17)))
                    case 4294967295:
                        v1 += (v1 / ((-9223372036854775808) | 1))
                    }
                    v19 := GenS4{F0: fi17, F1: -94}
                    _ = v19
                }
                _ = fi17
            }
        }
        _ = fi15
        _ = v16
    }
    v20 := GenS1{F0: 1, F1: 421, F2: ""}
    v1 = (v1 & 175)
    fmt.Println(v1, v1)
    v5 = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            if true {
                v2 = 2
                v4 = append(v4, 2147483648)
            }
            _ = fi22
        }
        _ = fi21
    }
    v23 := GenS5{F0: 1, F1: 2147483647, F2: 105}
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    switch 2147483647 {
    case -306:
        v2 += v2
    case 4294967295:
        v5 -= v5
    case 0:
        v1 += (v1 ^ (-(v1)))
    default:
        if true {
            v1 = v1
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 uint16 = 2
                var v26 int16 = 32767
                fmt.Println(v5, fi24)
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            _ = len(v4)
            var v27 int = 96
            for fi28 := 0; fi28 < 10; fi28++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi28
            }
            _ = v27
        }
    }
    _ = len(v3)
    var v29 float32 = (v5 * -2.25)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

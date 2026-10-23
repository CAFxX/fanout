package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint }
type GenS7 struct { F0 int16; F1 uint16 }
func genFn0() int8 {
    v7 := GenS1{F0: 18446744073709551615}
    var v8 int64 = -109
    v8 = v8
    _ = v7
    _ = v8
    return -73
}

func main() {
    var v1 uint32 = 0
    var v2 int8 = 1
    v3 := make(map[string]string)
    v4 := []uint32{(v1 << ((1) & 63))}
    var v5 float64 = 3.14159
    v6 := []int{895, -9223372036854775808}
    v9 := GenS2{F0: 32767}
    if false {
        for fi10 := 0; fi10 < 10; fi10++ {
            v2 -= (v2 + 81)
            switch (^(fi10)) {
            case 42:
                var v11 float64 = 0.1
                _ = v11
            case -872:
                for fi12 := 0; fi12 < 18; fi12++ {
                    if ((803 != 65535) && (3.14159 > 3.14159)) {
                        switch 127 {
                        case 2147483647:
                            v1 += v1
                        case -2147483648:
                            if (999 == 18446744073709551615) {
                                v5 += float64(float32(v5))
                                v2 = genFn0()
                                fmt.Println(v1)
                            }
                        case -296:
                            v13 := GenS3{F0: 18446744073709551615}
                            _ = v13
                        }
                        switch "foo bar" {
                        case "hello":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "a":
                            var v14 uint64 = 71
                            _ = v14
                        }
                    } else {
                        v5 = v5
                    }
                    _ = fi12
                }
            default:
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS4{F0: 777}
                    _ = fi15
                    _ = v16
                }
            }
            v5 = v5
            _ = fi10
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS5{F0: 2}
            v1 += v1
            v2 -= int8(fi17)
            _ = fi17
            _ = v18
        }
    } else {
        if (((false && ((621 < 0) && false)) || (v1 < 733)) || false) {
            v19 := v3[("foo bar" + "")]
            _ = v19
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                if len(v6) > 0 {
                    v6[0] = fi20
                }
                _ = fi20
            }
        }
        switch -686 {
        case 128:
            v6 = append(v6, 1)
        case 60:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -732:
            fmt.Println(v1, v5, v2)
        default:
            for fi21 := 0; fi21 < 12; fi21++ {
                v5 = float64(((0.1 - 3.14159) * 3.14159))
                _ = fi21
            }
        }
        v2 += 0
    }
    var v22 uint32 = 661
    fmt.Println(v22, v2)
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch "" {
        case "foo bar":
            v3[(("" + "") + "hello")] = ("" + ("hello" + "a"))
        case "":
            v1 += (^(v22))
        case "\n\t":
            switch "\n\t" {
            case "hello":
                v23 := GenS6{F0: 0, F1: 63, F2: 599}
                _ = v23
            case "\n\t":
                fmt.Println(v2, v5, v22)
            case "a":
                for fi24 := 0; fi24 < 12; fi24++ {
                    _ = len(v3)
                    v25 := GenS6{F0: 0, F1: -128, F2: 18446744073709551615}
                    if len(v6) > 0 {
                        v6[0] = 32767
                    }
                    _ = fi24
                    _ = v25
                }
            default:
                v26 := GenS4{F0: 0}
                _ = v26
            }
        default:
            _ = len(v6)
        }
    }
    fmt.Println(v1, v5)
    v27 := GenS3{F0: 65535}
    v1 -= (v1 - 4294967295)
    v28 := GenS7{F0: int16(v2), F1: 962}
    var v29 int8 = genFn0()
    var v30 string = "\n\t"
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    v31 := GenS3{F0: 42}
    var v32 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

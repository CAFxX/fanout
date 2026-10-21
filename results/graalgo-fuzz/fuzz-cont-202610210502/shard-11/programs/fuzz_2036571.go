package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint32; F2 string }

func main() {
    var v1 uint32 = 957
    var v2 uint16 = 65535
    var v3 uint = 2
    var v4 uint64 = 717
    var v5 int = 127
    var v6 uint8 = 15
    if true {
        v5 = (v5 & v5)
    }
    fmt.Println(v1, v2, v1)
    v7 := []int{(^(v5)), -9223372036854775808, (v5 >> ((42) & 63))}
    fmt.Println(v5, v4, v1)
    v4 += (v4 & 2)
    fmt.Println(v4)
    if len(v7) > 0 {
        v7[0] = v5
    }
    v8 := GenS1{F0: 1}
    fmt.Println(v3, v4, v5)
    for fi9 := 0; fi9 < 16; fi9++ {
        if true {
            var v10 uint = (v3 - v3)
            switch "a" {
            case "hello":
                for fi11 := 0; fi11 < 19; fi11++ {
                    if false {
                        fmt.Println(v1, v6, v4)
                        v7 = append(v7, (fi9 & (fi11 ^ (fi11 << (((v10 / ((42) | 1))) & 63)))))
                        v3 -= v3
                    }
                    v1 += uint32(v4)
                    fmt.Println(fi9)
                    _ = fi11
                }
            case "":
                if true {
                    switch "" {
                    case "\n\t":
                        v7 = append(v7, int(v2))
                    }
                    if (((v6 < (v6 % (((^(v6))) | 1))) || false) || false) {
                        fmt.Println(v5, v1)
                        for fi12 := 0; fi12 < 3; fi12++ {
                            var v13 uint32 = (v1 & (v1 & uint32(v4)))
                            fmt.Println(v10)
                            _ = fi12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 18; fi14++ {
                            var v15 int8 = -1
                            var v16 float32 = -2.25
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        v5 += 32768
                        for fi17 := 0; fi17 < 15; fi17++ {
                            switch "a" {
                            case "\n\t":
                                var v18 uint8 = (v6 * v6)
                                _ = v18
                            case "a":
                                v5 = fi9
                            case "foo bar":
                                switch "" {
                                case "":
                                    switch fi17 {
                                    case 832:
                                        v19 := GenS1{F0: 0}
                                        _ = v19
                                    case -2147483648:
                                        v2 = v2
                                    }
                                default:
                                    v20 := GenS2{F0: -1}
                                    _ = v20
                                }
                                v5 += (-(fi17))
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            default:
                                v21 := GenS1{F0: 666}
                                _ = v21
                            }
                            _ = fi17
                        }
                        v4 += v4
                    }
                    switch v5 {
                    case 67:
                        v22 := GenS2{F0: -937}
                        _ = v22
                    case 42:
                        v23 := GenS1{F0: 0}
                        _ = v23
                    case -1:
                        fmt.Println(v10, v1)
                    }
                } else {
                    var v24 uint8 = v6
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v6, fi9, v10)
                    case "hello":
                        var v25 uint8 = 128
                        _ = v25
                    case "":
                        switch "foo bar" {
                        case "a":
                            switch (v5 + (fi9 & -1)) {
                            case -9223372036854775808:
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    var v27 bool = ((v10 | 2) >= 2)
                                    _ = fi26
                                    _ = v27
                                }
                            default:
                                if (42 < -816) {
                                    var v28 float32 = 1.5
                                    _ = v28
                                }
                            }
                        case "":
                            v29 := GenS3{F0: "\n\t"}
                            _ = v29
                        }
                    }
                    _ = v24
                }
            }
            v30 := GenS4{F0: (-710 <= 325), F1: v1, F2: (("" + "") + "\n\t")}
            _ = v10
            _ = v30
        }
        v7 = append(v7, (v5 / ((-9223372036854775808) | 1)))
        _ = fi9
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            v33 := GenS1{F0: -690}
            v5 = 4294967295
            v34 := GenS3{F0: "foo bar"}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 1
    v3 := make(map[int]uint32)
    switch v2 {
    case -2147483648:
        if (2 == 42) {
            v4 := GenS1{F0: ("hello" + ("\n\t" + ("foo bar" + "foo bar"))), F1: 0}
            _ = v4
        }
    case 867:
        var v5 int32 = 969
        _ = v5
    }
    var v6 int32 = 1
    fmt.Println(v6)
    var v7 int16 = 1
    v6 += (v6 - v6)
    switch (v2 >> ((255) & 63)) {
    case 49:
        var v8 string = ""
        _ = v8
    case 2147483647:
        v9 := []int{9223372036854775807, (v2 % ((v2) | 1))}
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS2{F0: 291, F1: (v6 ^ (-(v6))), F2: 0}
            if false {
                v7 += (-(v7))
                v13 := []int{fi11, fi11}
                if ((v6 / ((0) | 1)) <= v6) {
                    var v14 uint64 = 4294967295
                    delete(v3, v2)
                    if false {
                        if true {
                            switch "" {
                            case "foo bar":
                                fmt.Println(v2)
                            case "":
                                v1[uint32(v14)] = true
                            case "hello":
                                if ((v7 - (-(v7))) < (v7 * -732)) {
                                    if ((1e10 < 1.5) && ((v7 >> ((104) & 63)) >= v7)) {
                                        if len(v13) > 0 {
                                            v13[0] = 2147483647
                                        }
                                        var v15 int = 1
                                        _ = v15
                                    }
                                    v14 -= 2
                                    var v16 int16 = -32768
                                    _ = v16
                                } else {
                                    fmt.Println(v7, v7, v14)
                                }
                                v6 += v6
                            default:
                                v17 := GenS1{F0: "\n\t", F1: 42}
                                _ = v17
                            }
                            if false {
                                v13 = append(v13, int(v6))
                                var v18 uint32 = 2
                                _ = v18
                            }
                        }
                    }
                    fmt.Println(fi11, v2, v2)
                    var v19 int8 = 21
                    _ = v14
                    _ = v19
                }
                _ = v13
            }
            v20 := GenS1{F0: "a", F1: 100}
            _ = fi11
            _ = v12
            _ = v20
        }
        fmt.Println(v2, v6)
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS2{F0: 0, F1: 1, F2: 128}
            v23 := []int{}
            var v24 int32 = (v6 ^ (v6 ^ (^(v6))))
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi10
    }
    v6 -= v6
    v6 += v6
    var v25 uint16 = 0
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint = 2
        for fi28 := 0; fi28 < 19; fi28++ {
            if true {
                if (18446744073709551615 < 754) {
                    v27 += (^(v27))
                } else {
                    fmt.Println(v6, v6, v6)
                }
                v2 -= (fi28 & int(v6))
            } else {
                var v29 float64 = float64(-2.25)
                var v30 int8 = 42
                v31 := []int{(v2 | fi28), fi28}
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        if (uint16(v27) <= v25) {
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS3{F0: (("hello" + "\n\t") + ((("hello" + "a") + ("a" + ("hello" + "hello"))) + ("hello" + ""))), F1: -4}
                v34 := GenS4{F0: 42, F1: 127, F2: 0}
                _ = fi32
                _ = v33
                _ = v34
            }
        } else {
            v25 = v25
            v2 += v2
            var v35 uint = (v27 >> ((v27) & 63))
            _ = v35
        }
        _ = fi26
        _ = v27
    }
    v2 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}

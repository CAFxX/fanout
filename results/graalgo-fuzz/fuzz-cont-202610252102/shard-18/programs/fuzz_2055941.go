package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 uint }
func genFn0(p0 uint) int16 {
    var v5 float32 = 0.1
    v5 = v5
    v5 = float32(0.0)
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 uint64 = 684
    var v2 float32 = (0.0 / 0.1)
    var v3 int32 = 2147483647
    v4 := []string{}
    var v6 uint32 = 639
    switch ("hello" + "") {
    case "a":
        switch "hello" {
        case "\n\t":
            var v7 int16 = 407
            _ = v7
        case "hello":
            v8 := GenS1{F0: 255}
            _ = v8
        case "foo bar":
            var v9 uint32 = (v6 << ((18446744073709551615) & 63))
            _ = v9
        default:
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v10 float32 = (v2 / (-2.25 - ((v2 / 1.0) + v2)))
                if false {
                    v11 := GenS2{F0: (^(v6)), F1: 2}
                    _ = v11
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v12 := GenS3{F0: -130}
                    var v13 uint = 916
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        }
    case "\n\t":
        var v14 uint32 = (v6 - v6)
        _ = v14
    default:
        if true {
            for fi15 := 0; fi15 < 5; fi15++ {
                fmt.Println(v2, v6, fi15)
                v6 = (v6 / ((v6) | 1))
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                _ = fi15
            }
            if (!(1 == 32767)) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if true {
                    var v16 string = "foo bar"
                    _ = v16
                }
                fmt.Println(v2, v2, v2)
            } else {
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 string = "a"
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v3 = (v3 | v3)
                        v18 = (v18 + "a")
                        _ = fi19
                    }
                    _ = fi17
                    _ = v18
                }
                var v20 uint8 = 2
                v4 = append(v4, "\n\t")
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 int64 = int64(v6)
                for fi23 := 0; fi23 < 16; fi23++ {
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    var v24 uint64 = (v1 * v1)
                    _ = fi23
                    _ = v24
                }
                v25 := GenS4{F0: 4294967295, F1: -1}
                _ = fi21
                _ = v22
                _ = v25
            }
            v26 := GenS5{F0: 65535}
            _ = v26
        }
    }
    fmt.Println(v2, v2)
    v27 := GenS4{F0: 951, F1: 0}
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println(v2, v3, v3)
    for fi28 := 0; fi28 < 16; fi28++ {
        v3 = (v3 / ((-818) | 1))
        v3 -= (v3 * (v3 / ((v3) | 1)))
        v29 := GenS4{F0: 255, F1: -113}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        if len(v4) > 0 {
            v4[0] = ("a" + "a")
        }
        if false {
            if true {
                var v31 uint16 = 831
                _ = v31
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 int32 = (v3 % (((v3 & int32(fi32))) | 1))
                v34 := GenS2{F0: v6, F1: 1}
                _ = fi32
                _ = v33
                _ = v34
            }
            var v35 uint64 = v1
            _ = v35
        }
        _ = len(v4)
        _ = fi30
    }
    v36 := GenS1{F0: 120}
    v37 := GenS6{F0: 1, F1: 809}
    fmt.Println(v1, v1, v2)
    if ((false || (1 <= 42)) || (true && true)) {
        v1 -= 863
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v38 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v6
}

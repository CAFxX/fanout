package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 int32, p1 int16) uint16 {
    v6 := GenS1{F0: 18446744073709551615}
    if (false || (p1 <= p1)) {
        p0 -= (p0 - (p0 + p0))
        p0 = p0
    }
    var v7 int64 = -1
    if true {
        var v8 float32 = 0.1
        v9 := []int{}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}
func genFn1() int64 {
    var v10 uint32 = 65535
    v11 := GenS2{F0: 100}
    _ = v10
    _ = v11
    return 497
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint64 = 1
    var v3 int8 = 42
    var v4 string = "hello"
    v5 := []uint32{(v1 / ((65535) | 1)), 0}
    v1 += (v1 - v1)
    v5 = append(v5, (v1 | v1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v12 int16 = -871
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v1 = v1
        _ = fi13
    }
    if (float32(3.14159) != -2.25) {
        v2 = (^(v2))
        v4 = v4
        if (true || (1 >= 2)) {
            switch v4 {
            case "":
                v2 -= 2
            case "foo bar":
                if (!(false && true)) {
                    fmt.Println(v2, v4, v4)
                }
            case "hello":
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := GenS2{F0: (v3 + (v3 >> ((65535) & 63)))}
                    var v16 string = v4
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            v1 -= v1
            switch 9223372036854775807 {
            case 4294967295:
                v17 := GenS2{F0: (^(v3))}
                _ = v17
            case 1:
                for fi18 := 0; fi18 < 14; fi18++ {
                    if (true || (0 >= 476)) {
                        v3 += 127
                        v19 := GenS2{F0: (v3 ^ -39)}
                        if len(v5) > 0 {
                            v5[0] = (v1 & (v1 ^ v1))
                        }
                        _ = v19
                    }
                    _ = fi18
                }
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v20 int16 = 42
            _ = v20
        }
    } else {
        v4 = (((((v4 + v4) + "hello") + v4) + ("hello" + v4)) + v4)
        fmt.Println(v2, v1, v2)
        v4 = "a"
    }
    if (733 >= 9223372036854775807) {
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := GenS2{F0: v3}
                var v24 uint = 4294967295
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 float32 = 3.14159
            fmt.Println(v4)
            _ = fi21
            _ = v25
        }
    } else {
        fmt.Println(v3, v1, v3)
        if (!(65535 > 1)) {
            fmt.Println(v3)
        }
    }
    v1 = v1
    if (false && (!true)) {
        if true {
            v26 := GenS2{F0: 1}
            _ = v26
        } else {
            switch ("\n\t" + "a") {
            case "":
                for fi27 := 0; fi27 < 8; fi27++ {
                    if true {
                        var v28 int16 = -32768
                        _ = v28
                    } else {
                        var v29 uint16 = 65535
                        _ = v29
                    }
                    _ = fi27
                }
            case "\n\t":
                v30 := GenS3{F0: -414}
                _ = v30
            }
            switch -9223372036854775808 {
            case -2147483648:
                v1 -= (v1 ^ v1)
            case 1:
                var v31 float64 = ((float64(-2.25) * (3.14159 + (3.14159 * 1.5))) + (0.0 / 1.0))
                _ = v31
            case 42:
                v2 = v2
            default:
                v32 := GenS4{F0: (v3 << ((156) & 63)), F1: 0}
                _ = v32
            }
        }
    } else {
        _ = len(v5)
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS1{F0: 65535}
            _ = fi33
            _ = v34
        }
    }
    var v35 uint = 0
    if len(v5) > 0 {
        v5[0] = (v1 + 65535)
    }
    v36 := GenS5{F0: 42, F1: v1, F2: 160}
    var v37 int16 = 42
    _ = len(v5)
    v35 = (v35 - v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}

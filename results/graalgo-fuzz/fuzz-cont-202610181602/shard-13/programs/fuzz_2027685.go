package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int = 1
    v2 := []string{("a" + "\n\t"), ("a" + (("" + "\n\t") + "a")), "hello"}
    v3 := []uint32{625}
    var v4 uint64 = 4294967295
    var v5 int8 = 1
    var v6 float32 = float32((-2.25 + -2.25))
    switch "" {
    case "a":
        if true {
            v3 = append(v3, 42)
            v7 := GenS1{F0: 490, F1: 127, F2: v4}
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = v7
        } else {
            switch v1 {
            case -400:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v8 := GenS1{F0: v4, F1: 0, F2: v4}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS2{F0: "", F1: (^(v4))}
            _ = fi10
            _ = v11
        }
        switch "foo bar" {
        case "foo bar":
            var v12 float64 = ((3.14159 - 1e10) + (3.14159 - 1e10))
            _ = v12
        case "hello":
            v13 := GenS2{F0: "a", F1: 4294967295}
            _ = v13
        }
        _ = fi9
    }
    if true {
        switch "foo bar" {
        case "foo bar":
            v14 := GenS2{F0: ("a" + "hello"), F1: v4}
            _ = v14
        case "a":
            if true {
                v1 -= -1
            }
        default:
            v1 -= (v1 << ((4294967295) & 63))
        }
    }
    v15 := GenS2{F0: ("a" + ""), F1: v4}
    switch v1 {
    case 32768:
        v4 = (v4 & v4)
    case 42:
        v16 := GenS2{F0: "hello", F1: v4}
        _ = v16
    case 150:
        v5 = v5
    default:
        v6 -= ((3.14159 * (v6 - ((-2.25 + v6) - 0.1))) - v6)
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        switch v1 {
        case 4294967295:
            switch (v1 & v1) {
            case 815:
                if len(v2) > 0 {
                    v2[0] = ""
                }
            }
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if true {
                v1 = v1
                if true {
                    var v19 uint32 = 4294967295
                    if (true || (65535 <= uint16(v19))) {
                        if (1 >= 579) {
                            v4 = v4
                            v6 += ((0.1 - (-2.25 - v6)) + v6)
                        } else {
                            v20 := GenS2{F0: "foo bar", F1: (v4 + (v4 | 342))}
                            _ = v20
                        }
                    } else {
                        if ((false || true) && false) {
                            var v21 int = -1
                            v19 = (v19 + v19)
                            var v22 int8 = v5
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = v19
                } else {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    v6 = (v6 * -2.25)
                    var v23 float32 = 1e10
                    _ = v23
                }
            } else {
                v24 := GenS3{F0: 42}
                _ = len(v3)
                var v25 uint = 1
                _ = v24
                _ = v25
            }
            var v26 int16 = -712
            var v27 float64 = ((1e10 - 1.5) / (1e10 + (0.1 - 0.0)))
            var v28 int16 = 42
            var v29 float64 = ((v27 / v27) / (v27 - v27))
            v5 = (v5 - v5)
            _ = fi18
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi17
    }
    v4 = (v4 % (((v4 % ((v4) | 1))) | 1))
    var v30 uint64 = 2
    for fi31 := 0; fi31 < 14; fi31++ {
        if true {
            if (32767 != 32767) {
                if (967 < 809) {
                    v30 = 657
                    v1 -= v1
                    fmt.Println(v4, v6, v4)
                }
            }
            if len(v3) > 0 {
                v3[0] = 0
            }
            var v32 float32 = v6
            _ = v32
        } else {
            for fi33 := 0; fi33 < 6; fi33++ {
                if (v6 < 0.0) {
                    v1 -= fi31
                } else {
                    var v34 uint64 = 42
                    _ = v34
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi33
            }
            fmt.Println(v30, fi31, fi31)
        }
        v5 += -117
        _ = fi31
    }
    v4 -= (v4 % (((v4 + 18446744073709551615)) | 1))
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

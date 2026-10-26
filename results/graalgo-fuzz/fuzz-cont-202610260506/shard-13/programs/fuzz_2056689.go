package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 string }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 string) int32 {
    v6 := []int{-101}
    var v7 uint8 = 1
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{2, 0, 294, 2147483648}
    var v2 int16 = -1
    var v3 uint8 = 110
    var v4 uint64 = 0
    var v5 float32 = 1.5
    v8 := GenS1{F0: 919, F1: 650, F2: (("\n\t" + ((("" + "a") + ("hello" + "a")) + (("\n\t" + ("hello" + ((("hello" + "a") + "a") + "a"))) + "foo bar"))) + "\n\t")}
    v4 -= v4
    switch ("hello" + "foo bar") {
    case "a":
        v1 = append(v1, 42)
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        if (!(255 < 1)) {
            v9 := GenS1{F0: -235, F1: 384, F2: "hello"}
            v3 = v3
            _ = v9
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "hello" {
            case "":
                for fi10 := 0; fi10 < 15; fi10++ {
                    v5 = v5
                    _ = fi10
                }
            case "\n\t":
                var v11 uint16 = 42
                _ = v11
            case "hello":
                var v12 int32 = 1
                _ = v12
            default:
                v13 := GenS2{F0: true, F1: false, F2: 542}
                _ = v13
            }
        }
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 bool = false
        _ = fi14
        _ = v15
    }
    var v16 string = ""
    if (2147483647 > 42) {
        if (0 != -293) {
            fmt.Println(v5)
        } else {
            if true {
                var v17 uint32 = 65535
                _ = v17
            } else {
                switch -2147483648 {
                case 335:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v20 := GenS3{F0: 1, F1: 100, F2: (v3 | v3)}
                            _ = fi19
                            _ = v20
                        }
                        _ = fi18
                    }
                case 4294967295:
                    v5 -= v5
                case -214:
                    var v21 uint32 = 2
                    _ = v21
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v22 uint16 = 1
                _ = v22
            }
            _ = len(v1)
            var v23 string = ("hello" + v16)
            _ = v23
        }
        fmt.Println(v3, v4)
        v4 = v4
    } else {
        if false {
            fmt.Println(v16, v3)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 string = v16
                    v1 = append(v1, 4294967295)
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            v1 = append(v1, 0)
        }
    }
    if (v2 == (v2 * (v2 >> ((18446744073709551615) & 63)))) {
        v27 := GenS3{F0: uint(v3), F1: 23, F2: (v3 | (v3 * (v3 & v3)))}
        var v28 int16 = (v2 % (((v2 % (((v2 - -1)) | 1))) | 1))
        v4 += v4
        _ = v27
        _ = v28
    } else {
        if (!(!(!((((v5 * 0.0) > (float32(0.1) - v5)) && false) && true)))) {
            fmt.Println(v4, v2)
            if len(v1) > 0 {
                v1[0] = 2
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v29 := GenS2{F0: (!true), F1: true, F2: 55}
            if true {
                var v30 uint = 42
                v5 = float32(0.0)
                var v31 string = (v16 + v16)
                _ = v30
                _ = v31
            }
            v2 -= v2
            _ = v29
        }
    }
    if (-1 == -9223372036854775808) {
        _ = len(v1)
    }
    var v32 string = "\n\t"
    var v33 int8 = 100
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v8
}

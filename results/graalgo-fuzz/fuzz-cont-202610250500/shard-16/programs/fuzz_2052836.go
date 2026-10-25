package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }

func main() {
    v1 := []int{32768}
    var v2 int32 = -366
    var v3 string = (("foo bar" + "hello") + ("a" + (("foo bar" + "foo bar") + "")))
    v4 := GenS1{F0: 0, F1: 42, F2: 65535}
    v3 = v3
    v2 += -1
    if false {
        switch "foo bar" {
        case "a":
            var v5 uint16 = 42
            _ = v5
        case "":
            fmt.Println(v3)
        default:
            v2 -= (v2 % ((-177) | 1))
        }
        switch v3 {
        case "\n\t":
            fmt.Println(v2)
        case "":
            v2 = 1
        default:
            var v6 uint32 = 877
            _ = v6
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 uint32 = 117
    v8 := GenS1{F0: 255, F1: 573, F2: 386}
    for fi9 := 0; fi9 < 4; fi9++ {
        v1 = append(v1, (fi9 + (fi9 % (((-(fi9))) | 1))))
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        v7 = 2147483648
        _ = fi9
    }
    if (!true) {
        if false {
            var v10 uint32 = 428
            var v11 int32 = v2
            if (1000 != 121) {
                v12 := GenS1{F0: 65535, F1: 255, F2: 2}
                var v13 uint32 = v10
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v11
        }
        switch -127 {
        case -828:
            v14 := GenS2{F0: 2, F1: -48, F2: 42}
            _ = v14
        case 9223372036854775807:
            if (v7 < (v7 | 320)) {
                var v15 int = 42
                for fi16 := 0; fi16 < 6; fi16++ {
                    fmt.Println(v2)
                    _ = fi16
                }
                if (0 >= 100) {
                    var v17 uint32 = 530
                    var v18 uint16 = 65535
                    var v19 int = (v15 % (((v15 - 127)) | 1))
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v15)
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(v15, v15, v15)
                        _ = fi20
                    }
                }
                _ = v15
            } else {
                for fi21 := 0; fi21 < 10; fi21++ {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        _ = len(v1)
                        _ = len(v1)
                        _ = fi22
                    }
                    _ = fi21
                }
                fmt.Println(v2, v2)
            }
        case 770:
            v1 = append(v1, 1)
        }
        var v23 string = ("" + (v3 + ((v3 + (v3 + v3)) + ("hello" + ("" + (v3 + "hello"))))))
        _ = v23
    } else {
        v2 -= -67
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v2 += v2
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 string }
func genFn0(p0 int16, p1 uint16) uint32 {
    var v5 int16 = 955
    switch "\n\t" {
    case "hello":
        v6 := []int{int(v5), 9223372036854775807, 4294967295, 32767}
        _ = v6
    case "":
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v5
    return 2147483648
}

func main() {
    var v1 float32 = 1.5
    var v2 uint16 = 32767
    var v3 uint64 = 42
    var v4 uint16 = v2
    if true {
        v3 = (v3 << ((1) & 63))
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: -1, F1: 194}
            if (((fi7 % (((fi7 + 32767)) | 1)) != fi7) || (true || (0 <= -1))) {
                if (!false) {
                    switch -1 {
                    case -496:
                        v2 = 532
                    default:
                        var v9 string = ("foo bar" + "hello")
                        _ = v9
                    }
                }
            } else {
                var v10 bool = false
                switch "hello" {
                case "hello":
                    if true {
                        v11 := GenS2{F0: genFn0(-96, (v4 << ((777) & 63))), F1: (fi7 + (^(fi7))), F2: -1}
                        v12 := GenS3{F0: 1, F1: v3, F2: 745}
                        _ = v11
                        _ = v12
                    }
                case "":
                    var v13 uint32 = 1
                    _ = v13
                }
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        v2 -= uint16(v3)
    }
    var v14 int = 32767
    fmt.Println(v14)
    var v15 uint32 = 2
    if (9223372036854775807 == -1) {
        if (-500 != v14) {
            fmt.Println(v1, v4, v1)
        } else {
            v1 = float32(3.14159)
        }
    }
    if true {
        var v16 int64 = -753
        _ = v16
    } else {
        v4 = v4
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        if (false || false) {
            for fi18 := 0; fi18 < 4; fi18++ {
                if false {
                    v19 := []int{fi17, v14, (fi18 - fi17)}
                    var v20 uint = 1
                    _ = len(v19)
                    _ = v19
                    _ = v20
                }
                v4 -= (v2 / ((219) | 1))
                var v21 uint16 = v4
                _ = fi18
                _ = v21
            }
            v22 := GenS4{F0: 629, F1: 255, F2: "hello"}
            _ = v22
        }
        v3 = 0
        v23 := []int{(fi17 & 887), (fi17 / ((int(v3)) | 1))}
        _ = fi17
        _ = v23
    }
    var v24 bool = (false || (65535 == 255))
    v25 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}

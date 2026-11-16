package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    var v1 float32 = (0.0 + -2.25)
    var v2 uint16 = 42
    var v3 int16 = 652
    var v4 uint8 = 42
    fmt.Println(v4, v1, v4)
    v5 := []int{}
    var v6 int64 = -1
    v7 := GenS1{F0: "", F1: (0.1 != 1e10), F2: v2}
    fmt.Println(v3, v3)
    if (v6 > (^(v6))) {
        v3 += (v3 << ((255) & 63))
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: "a", F1: false, F2: 604}
        if false {
            var v10 uint8 = v4
            _ = v10
        }
        if (v6 == 42) {
            var v11 uint32 = 2147483648
            var v12 int64 = (v6 ^ (v6 / ((v6) | 1)))
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    if (float32((1e10 + 0.1)) <= v1) {
        var v13 bool = ((1 < 128) || (((v3 + (v3 * v3)) == (v3 / (((v3 * v3)) | 1))) || (!((3.14159 - v1) > 0.0))))
        fmt.Println(v3, v2, v13)
        _ = v13
    } else {
        if (!true) {
            v14 := GenS1{F0: (((("hello" + ((("\n\t" + "a") + "") + ("foo bar" + "\n\t"))) + ("hello" + "a")) + "\n\t") + ("hello" + "foo bar")), F1: (!(!(!(0 < 0)))), F2: (v2 * (^(v2)))}
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 uint16 = (v2 ^ v2)
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            v17 := GenS2{F0: 1, F1: v2}
            v18 := GenS3{F0: "", F1: 42, F2: 595}
            var v19 string = "hello"
            _ = v17
            _ = v18
            _ = v19
        }
    }
    var v20 uint8 = v4
    if (638 >= 622) {
        v2 += (v2 | (v2 ^ v2))
        _ = len(v5)
    } else {
        v21 := GenS4{F0: (v6 & v6), F1: v4, F2: 736}
        _ = v21
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        if true {
            switch (fi22 ^ (fi22 ^ fi22)) {
            case 32767:
                if true {
                    v4 += 42
                    switch "" {
                    case "foo bar":
                        v5 = append(v5, (fi22 + 1))
                    case "":
                        for fi23 := 0; fi23 < 13; fi23++ {
                            _ = len(v5)
                            var v24 uint16 = (v2 << ((65535) & 63))
                            _ = fi23
                            _ = v24
                        }
                    default:
                        fmt.Println(fi22)
                    }
                } else {
                    v4 += (v20 % ((v20) | 1))
                    _ = len(v5)
                }
            case -1:
                var v25 int64 = (v6 << ((1) & 63))
                _ = v25
            }
        }
        var v26 uint8 = v20
        var v27 bool = (4294967295 >= 53)
        _ = fi22
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint }

func main() {
    var v1 int8 = 1
    var v2 uint8 = 77
    v3 := []uint32{645, 42, 2147483647, 0}
    var v4 bool = false
    var v5 int8 = v1
    var v6 int = 2147483647
    for fi7 := 0; fi7 < 13; fi7++ {
        switch ((("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + "hello") + "foo bar") {
        case "\n\t":
            for fi8 := 0; fi8 < 3; fi8++ {
                v2 = v2
                v5 = (v1 - (v5 % ((0) | 1)))
                var v9 uint64 = 27
                _ = fi8
                _ = v9
            }
        case "":
            var v10 int32 = -2147483648
            _ = v10
        case "foo bar":
            var v11 int = (v6 >> ((1) & 63))
            _ = v11
        default:
            if (((v4 && (v4 || v4)) || ("a" == "\n\t")) || (65535 > 42)) {
                v1 += v1
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch v6 {
        case -9223372036854775808:
            v4 = false
        case -792:
            v12 := GenS1{F0: true, F1: 9223372036854775807, F2: (v6 - (fi7 << ((4294967295) & 63)))}
            _ = v12
        case -395:
            if (42 >= 4294967295) {
                switch (fi7 >> ((4294967295) & 63)) {
                case 896:
                    v5 = (v1 ^ 96)
                case -2147483648:
                    fmt.Println(v5, v6, v1)
                }
            } else {
                if false {
                    v4 = v4
                    var v13 uint64 = 0
                    _ = v13
                }
                var v14 float32 = 0.1
                for fi15 := 0; fi15 < 19; fi15++ {
                    switch (v6 ^ (^(fi15))) {
                    case 2147483647:
                        if (677 <= 0) {
                            var v16 uint16 = 777
                            _ = v16
                        } else {
                            if v4 {
                                v1 -= v1
                            }
                        }
                        v5 = (v1 & 0)
                        if (("\n\t" + "foo bar") >= "a") {
                            var v17 uint64 = 42
                            _ = v17
                        }
                    default:
                        v4 = v4
                    }
                    _ = fi15
                }
                _ = v14
            }
        }
        _ = fi7
    }
    var v18 int32 = -2147483648
    v4 = (!(1.5 < (1e10 - -2.25)))
    var v19 uint16 = 255
    v4 = false
    _ = len(v3)
    v1 = (v1 | (v1 * v5))
    v5 = (v1 % ((v1) | 1))
    var v20 uint32 = 65535
    for fi21 := 0; fi21 < 8; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 int = -1
            v24 := GenS2{F0: (v5 % ((int8(v20)) | 1))}
            v2 = v2
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    v25 := GenS3{F0: uint16(v2), F1: 463}
    switch "\n\t" {
    case "foo bar":
        v26 := GenS1{F0: true, F1: 9223372036854775807, F2: v6}
        _ = v26
    }
    v27 := GenS2{F0: (v5 - v5)}
    for fi28 := 0; fi28 < 7; fi28++ {
        v2 = (v2 << ((uint(v20)) & 63))
        _ = fi28
    }
    if (!(!true)) {
        v29 := GenS1{F0: true, F1: int64(v6), F2: (v6 % ((v6) | 1))}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

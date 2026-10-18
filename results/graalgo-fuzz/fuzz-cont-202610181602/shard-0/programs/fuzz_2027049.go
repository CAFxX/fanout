package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
type GenS3 struct { F0 int64; F1 uint8; F2 int }

func main() {
    var v1 int64 = -341
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    v4 := []string{"a", ("a" + (("foo bar" + "") + ("\n\t" + "\n\t")))}
    var v5 bool = (-2.25 > 0.0)
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 42, F1: -13}
        v5 = (743 > 2147483647)
        v8 := GenS1{F0: 2, F1: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    if v5 {
        v1 = v1
        v9 := GenS1{F0: 1, F1: 0}
        _ = v9
    }
    fmt.Println(v5)
    for fi10 := 0; fi10 < 3; fi10++ {
        v1 += v1
        _ = fi10
    }
    if (-469 < -979) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v11 float64 = (-2.25 + -2.25)
        if (285 != 459) {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v5)
                v13 := GenS1{F0: 2, F1: -2147483648}
                var v14 int64 = -9223372036854775808
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 uint64 = 463
            _ = v15
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v11
    } else {
        switch 128 {
        case 0:
            if false {
                var v16 uint32 = 4294967295
                if false {
                    v17 := GenS1{F0: 42, F1: 902}
                    _ = v17
                }
                _ = v16
            } else {
                v1 -= (v1 ^ v1)
            }
        }
    }
    v18 := GenS2{F0: 255, F1: 0, F2: ("foo bar" + ("\n\t" + "foo bar"))}
    if v5 {
        var v19 uint8 = 226
        for fi20 := 0; fi20 < 5; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS3{F0: (v1 * (v1 << ((687) & 63))), F1: (^(v19)), F2: fi21}
                v1 = v1
                var v23 uint = 924
                _ = fi21
                _ = v22
                _ = v23
            }
            v1 += (v1 % (((v1 * 1)) | 1))
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 float64 = (0.1 - (1.5 / -2.25))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v19 -= 128
                _ = fi24
                _ = v25
            }
            _ = fi20
        }
        _ = v19
    }
    var v26 uint64 = 2
    var v27 uint64 = v26
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

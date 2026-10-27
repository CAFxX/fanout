package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }

func main() {
    var v1 int = 1
    var v2 int16 = 32767
    var v3 uint16 = 2
    v1 = -2147483648
    if true {
        var v4 int32 = 42
        _ = v4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int16 = 32767
        switch (fi5 | (fi5 ^ (fi5 + v1))) {
        case 9223372036854775807:
            var v7 int = 32768
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    v8 := []int{-1}
    v9 := GenS1{F0: 9223372036854775807}
    v10 := GenS2{F0: (v3 ^ 255)}
    if true {
        v3 -= 0
        v11 := GenS2{F0: v3}
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint8 = 0
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS3{F0: 47, F1: fi12, F2: 2}
                var v16 uint64 = 4294967295
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    if (5 != -1) {
        if len(v8) > 0 {
            v8[0] = (v1 - v1)
        }
        if true {
            if len(v8) > 0 {
                v8[0] = v1
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                v3 -= (v3 + v3)
                _ = fi17
            }
            v3 = (v3 % ((42) | 1))
        } else {
            if ((!((("foo bar" + "hello") < "") && ((42 == v2) && ((v3 + 219) != (v3 ^ v3))))) || false) {
                var v18 int = 544
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = len(v8)
                _ = v18
            }
            var v19 uint32 = 4294967295
            if true {
                fmt.Println(v2)
            } else {
                fmt.Println(v3, v3, v3)
                var v20 uint32 = v19
                if false {
                    v1 = (v1 | (v1 + v1))
                } else {
                    var v21 int = 32768
                    v20 += (v19 >> ((255) & 63))
                    for fi22 := 0; fi22 < 6; fi22++ {
                        fmt.Println(v3)
                        if len(v8) > 0 {
                            v8[0] = 0
                        }
                        _ = fi22
                    }
                    _ = v21
                }
                _ = v20
            }
            _ = v19
        }
    } else {
        for fi23 := 0; fi23 < 8; fi23++ {
            if false {
                v3 += (v3 / ((42) | 1))
                v24 := GenS1{F0: -88}
                v1 = -907
                _ = v24
            } else {
                var v25 string = "foo bar"
                var v26 uint64 = 1
                var v27 uint32 = 1
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi23
        }
        var v28 uint64 = 42
        _ = v28
    }
    var v29 int8 = 100
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}

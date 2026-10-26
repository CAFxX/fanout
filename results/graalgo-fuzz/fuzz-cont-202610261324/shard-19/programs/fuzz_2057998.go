package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int8; F2 bool }

func main() {
    var v1 int16 = -239
    var v2 int64 = -1
    var v3 bool = false
    v3 = v3
    for fi4 := 0; fi4 < 13; fi4++ {
        switch "\n\t" {
        case "a":
            switch (fi4 << ((0) & 63)) {
            case -1:
                v1 = v1
            default:
                var v5 int16 = (-(v1))
                _ = v5
            }
        case "":
            v1 -= v1
        }
        if true {
            if false {
                var v6 int64 = v2
                _ = v6
            } else {
                fmt.Println(v1, fi4, fi4)
                for fi7 := 0; fi7 < 15; fi7++ {
                    switch fi4 {
                    case 32767:
                        var v8 int16 = -32768
                        _ = v8
                    case 4294967295:
                        v9 := []int{(fi4 ^ (fi7 >> ((1) & 63))), -357, fi7}
                        _ = v9
                    case 280:
                        var v10 int32 = 0
                        _ = v10
                    }
                    v11 := GenS1{F0: 13}
                    switch (fi7 - (fi4 << ((4294967295) & 63))) {
                    case 498:
                        var v12 int64 = v2
                        fmt.Println(v12)
                        _ = v12
                    case -425:
                        var v13 int32 = 0
                        _ = v13
                    }
                    _ = fi7
                    _ = v11
                }
                v14 := []int{32768, (fi4 + (fi4 | (fi4 >> ((1) & 63)))), (fi4 / (((-(fi4))) | 1))}
                _ = v14
            }
            v15 := GenS1{F0: 128}
            var v16 int = (fi4 << ((924) & 63))
            _ = v15
            _ = v16
        } else {
            v17 := []int{}
            var v18 float32 = (float32(0.1) + float32(1e10))
            var v19 uint8 = 255
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi4
    }
    if (!(65535 > 426)) {
        v20 := GenS2{F0: 0, F1: 205}
        fmt.Println(v1, v3, v2)
        for fi21 := 0; fi21 < 11; fi21++ {
            v2 = (v2 ^ (v2 << ((221) & 63)))
            var v22 int32 = 477
            v1 -= (-(v1))
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        var v23 uint16 = 745
        v23 -= v23
        switch -9223372036854775808 {
        case 4294967295:
            v24 := GenS3{F0: v3, F1: 2}
            _ = v24
        }
        _ = v23
    }
    v3 = false
    v2 = 9223372036854775807
    fmt.Println(v3, v1, v2)
    for fi25 := 0; fi25 < 14; fi25++ {
        if (65535 == 1) {
            switch 1 {
            case 0:
                v26 := GenS2{F0: -100, F1: 2}
                _ = v26
            default:
                var v27 int64 = v2
                _ = v27
            }
        } else {
            v28 := []int{(fi25 | fi25), int(v2)}
            _ = v28
        }
        if v3 {
            var v29 uint64 = 42
            v2 = (v2 >> ((42) & 63))
            _ = v29
        }
        _ = fi25
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v1 = (^(v1))
        if (1 > -74) {
            var v31 float32 = (1.5 - 1e10)
            v32 := GenS4{F0: 41, F1: 126, F2: ((v2 != (-(v2))) || v3)}
            v3 = v3
            _ = v31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 10; fi33++ {
                v1 += -38
                fmt.Println(fi30, fi30)
                _ = fi33
            }
        }
        v34 := []int{32767}
        _ = fi30
        _ = v34
    }
    var v35 int32 = -957
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}

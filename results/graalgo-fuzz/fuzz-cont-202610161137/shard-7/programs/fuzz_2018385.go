package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 uint = 255
    var v4 int = 467
    v4 -= -1
    var v5 uint16 = 276
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{}
        v8 := GenS1{F0: v5, F1: 65535, F2: 582}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (!((v4 + (v4 >> ((v3) & 63))) > v4)) {
        switch (v4 - (v4 % ((v4) | 1))) {
        case 128:
            var v9 int8 = -1
            _ = v9
        case -48:
            v4 -= v4
        case -2147483648:
            var v10 int8 = -52
            _ = v10
        default:
            v11 := []int{(v4 - v4), v4, v4, 42}
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS1{F0: (v5 & 65535), F1: 2, F2: 1}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                v5 = 632
                v16 := GenS2{F0: 42}
                _ = fi15
                _ = v16
            }
            fmt.Println(v5, fi14, fi14)
            v17 := GenS1{F0: v5, F1: 881, F2: 845}
            _ = fi14
            _ = v17
        }
    }
    if false {
        fmt.Println(v3)
    } else {
        v3 = (v3 + v3)
        v18 := []int{v4, v4, (v4 & (v4 + (v4 + -9223372036854775808)))}
        _ = v18
    }
    switch v4 {
    case 127:
        v19 := GenS1{F0: v5, F1: 586, F2: 1}
        _ = v19
    case 9223372036854775807:
        switch v4 {
        case -326:
            var v20 string = ("hello" + "hello")
            _ = v20
        case -139:
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS1{F0: 1, F1: 65535, F2: 2}
                _ = len(v2)
                v23 := GenS1{F0: (v5 - 426), F1: 2147483648, F2: 0}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    fmt.Println(v4)
    switch ("a" + "foo bar") {
    case "":
        fmt.Println(v3, v4)
    case "hello":
        v24 := GenS3{F0: -55, F1: (v5 * v5), F2: false}
        _ = v24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        if false {
            v3 = v3
            var v26 float64 = 1.5
            var v27 float32 = (-2.25 + ((1e10 * -2.25) + 1.5))
            _ = v26
            _ = v27
        }
        var v28 float32 = float32(0.0)
        _ = fi25
        _ = v28
    }
    var v29 uint16 = (v5 | (v5 + v5))
    switch 796 {
    case 509:
        v4 = (v4 % (((v4 << (((v3 & (v3 & v3))) & 63))) | 1))
    case 9223372036854775807:
        if true {
            switch (v4 - 42) {
            case -121:
                v3 += v3
            case 127:
                var v30 uint64 = 42
                _ = v30
            }
        } else {
            switch (v4 % (((v4 << (((v3 / (((v3 - 42)) | 1))) & 63))) | 1)) {
            case 2147483647:
                var v31 bool = ("\n\t" <= "\n\t")
                _ = v31
            case -415:
                v5 += v29
            default:
                var v32 int8 = -1
                _ = v32
            }
            v3 -= (v3 << ((v3) & 63))
        }
    default:
        if false {
            v33 := GenS3{F0: -1, F1: (v29 - v5), F2: false}
            if false {
                var v34 int32 = 835
                var v35 int64 = 1
                _ = v34
                _ = v35
            } else {
                switch v4 {
                case 0:
                    fmt.Println(v3, v4)
                }
                v29 += (^(v5))
                v36 := []int{32768, (v4 / (((v4 * (v4 << ((v3) & 63)))) | 1)), v4, 9223372036854775807}
                _ = v36
            }
            _ = v33
        }
    }
    delete(v2, "\n\t")
    fmt.Println(v3, v5, v5)
    v29 = v29
    var v37 int64 = 42
    v3 = (v3 << (((^(v3))) & 63))
    v38 := GenS4{F0: "hello", F1: (v37 ^ -9223372036854775808)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}

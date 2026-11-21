package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{127, -9223372036854775808, 714}
    var v3 uint8 = 0
    var v4 uint8 = 127
    var v5 float32 = ((((3.14159 - float32((-2.25 - (1e10 / ((-2.25 - 1e10) / 1.5))))) / -2.25) / (1.5 + 1e10)) * 1.5)
    v4 = 33
    for fi6 := 0; fi6 < 3; fi6++ {
        switch "\n\t" {
        case "hello":
            var v7 bool = true
            _ = v7
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = (fi6 * (-(fi6)))
            }
        case "foo bar":
            if (true || (-374 > -2147483648)) {
                fmt.Println(v3)
                var v8 uint16 = 65535
                if (false && false) {
                    v9 := GenS1{F0: 0, F1: 42, F2: 1}
                    v5 += v5
                    var v10 string = "\n\t"
                    _ = v10
                    _ = v9
                }
                _ = v8
            } else {
                if ((0.1 * 0.1) >= 1e10) {
                    v3 += v4
                    _ = len(v2)
                }
            }
        default:
            v4 -= (v3 / (((v4 + (v3 + v4))) | 1))
        }
        _ = fi6
    }
    if (true || (2 > 4294967295)) {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                v2 = append(v2, fi12)
                _ = fi12
            }
            _ = fi11
        }
        switch 32768 {
        case -9223372036854775808:
            _ = len(v2)
        default:
            switch -2147483648 {
            case 1:
                var v13 uint32 = 1
                _ = v13
            case 9223372036854775807:
                fmt.Println(v3, v5)
            case 4294967295:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        v14 := GenS1{F0: 0, F1: -1, F2: 0}
        _ = v14
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            switch -831 {
            case -31:
                fmt.Println(v3)
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                switch "a" {
                case "foo bar":
                    var v16 uint32 = 2
                    _ = v16
                case "hello":
                    switch fi15 {
                    case 4294967295:
                        fmt.Println(v5, v4)
                    }
                }
                delete(v1, fi15)
                _ = fi15
            }
        } else {
            v5 -= 0.0
            switch 32767 {
            case 990:
                var v17 uint64 = 0
                _ = v17
            }
        }
    }
    if true {
        delete(v1, 1)
        var v18 uint8 = 112
        var v19 float64 = 1e10
        _ = v18
        _ = v19
    }
    var v20 bool = false
    v4 -= v4
    var v21 int8 = -128
    if (!v20) {
        if (((v20 || false) && true) && (false && (("hello" + "\n\t") != "\n\t"))) {
            fmt.Println(v21, v4, v20)
            var v22 uint32 = 2147483647
            for fi23 := 0; fi23 < 8; fi23++ {
                switch fi23 {
                case 32767:
                    v21 += v21
                case -143:
                    var v24 float32 = ((v5 * v5) * 1.5)
                    _ = v24
                case 681:
                    var v25 float64 = (1.5 + ((float64((float32(float64(v5)) + 1e10)) / 1.0) + -2.25))
                    _ = v25
                }
                fmt.Println(v21, v3)
                var v26 uint8 = 128
                _ = fi23
                _ = v26
            }
            _ = v22
        }
        var v27 int16 = 129
        _ = v27
    } else {
        var v28 uint16 = 60
        v5 -= v5
        _ = v28
    }
    var v29 float32 = (v5 * -2.25)
    delete(v1, -1)
    fmt.Println(v4, v20, v20)
    v4 -= 255
    fmt.Println(v29, v3)
    var v30 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

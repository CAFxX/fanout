package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16, p1 int32) int8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p1, p0)
        _ = fi5
    }
    p0 = (-(p0))
    var v6 string = ""
    v7 := []int{4294967295}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: p1, F1: v6, F2: 268}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 uint16, p2 uint8) uint {
    p0 = uint(p2)
    if false {
        v11 := []int{-2147483648, 0, -893}
        switch 4294967295 {
        case 9223372036854775807:
            v12 := GenS1{F0: 2147483647, F1: (("hello" + "foo bar") + (("a" + "hello") + ("foo bar" + (("" + "\n\t") + "")))), F2: 0}
            _ = v12
        case -1:
            p0 += (p0 % ((0) | 1))
        default:
            var v13 bool = true
            _ = v13
        }
        _ = v11
    } else {
        switch -2147483648 {
        case 32768:
            v14 := []int{32768, 45, 0, 72}
            _ = v14
        case 127:
            var v15 uint64 = 42
            _ = v15
        case 805:
            var v16 int16 = 32767
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    v3 := make(map[uint32]uint32)
    var v4 int64 = 42
    if true {
        v17 := GenS1{F0: int32(v4), F1: "foo bar", F2: 730}
        _ = v17
    }
    v2 += 9223372036854775807
    v2 += v2
    v3[2147483647] = 4294967295
    var v18 float64 = 3.14159
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v18 = 1e10
            v21 := GenS1{F0: 0, F1: "foo bar", F2: 42}
            _ = fi20
            _ = v21
        }
        switch (fi19 + (fi19 / ((9223372036854775807) | 1))) {
        case 1:
            var v22 uint64 = 2
            _ = v22
        case 4294967295:
            v18 = ((((0.0 + float64(0.0)) * float64(-2.25)) + (((float64(-2.25) * 1e10) - float64(0.1)) / ((float64(float32(v18)) + float64(3.14159)) * ((v18 * v18) * v18)))) / 3.14159)
        default:
            var v23 uint64 = 0
            _ = v23
        }
        if true {
            fmt.Println(v18)
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := []int{(-(fi19)), fi24, (fi19 % (((fi24 | fi19)) | 1)), fi24}
                fmt.Println(v2)
                v2 += 0
                _ = fi24
                _ = v25
            }
            if (-2.25 >= 1.5) {
                for fi26 := 0; fi26 < 18; fi26++ {
                    _ = len(v1)
                    _ = fi26
                }
                v27 := GenS2{F0: 0}
                _ = len(v3)
                _ = v27
            } else {
                fmt.Println(v2, v18)
            }
        } else {
            var v28 float64 = float64(1e10)
            _ = v28
        }
        _ = fi19
    }
    var v29 int8 = 0
    v4 += v2
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS2{F0: 18446744073709551615}
        for fi32 := 0; fi32 < 18; fi32++ {
            switch "\n\t" {
            case "":
                var v33 uint8 = 128
                _ = v33
            case "a":
                switch (^(fi32)) {
                case 1000:
                    v29 -= (v29 / ((genFn0(1, -2147483648)) | 1))
                }
            case "\n\t":
                fmt.Println(v2, v4)
            }
            if false {
                v34 := []int{127, -9223372036854775808, (fi32 | 128)}
                v4 = -1
                _ = v34
            } else {
                var v35 uint8 = 1
                switch "a" {
                case "\n\t":
                    v3[2] = 177
                }
                var v36 float64 = 0.0
                _ = v35
                _ = v36
            }
            _ = fi32
        }
        if true {
            v37 := []int{fi30, fi30, fi30, fi30}
            _ = v37
        }
        _ = fi30
        _ = v31
    }
    if (!true) {
        var v38 int64 = v2
        v2 = (v2 << ((65535) & 63))
        _ = v38
    }
    if false {
        var v39 int64 = (v4 * (v4 >> ((42) & 63)))
        v3[2147483648] = 391
        _ = v39
    }
    var v40 int64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
}

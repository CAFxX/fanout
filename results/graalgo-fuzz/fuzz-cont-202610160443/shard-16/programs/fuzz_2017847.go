package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]int)
    var v3 uint = 4294967295
    v4 := make(map[uint32]string)
    v3 = 828
    if true {
        var v5 int32 = int32(v1)
        v3 -= (v3 | v3)
        _ = v5
    }
    if ((!(!false)) || (!true)) {
        v1 -= 1
        v6 := []int{-89, 845}
        v3 = v3
        _ = v6
    } else {
        if ((v1 * (^(v1))) == v1) {
            switch 1 {
            case 1:
                var v7 float64 = (1.5 + 0.1)
                _ = v7
            case 338:
                if true {
                    v1 += v1
                    v2["foo bar"] = 1
                }
            case 32768:
                v3 += v3
            }
            v8 := GenS1{F0: 451, F1: 65535, F2: 4294967295}
            var v9 uint64 = 880
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 2, F1: 255, F2: -2147483648}
            _ = v10
        }
    }
    v3 -= v3
    if (true || (("hello" + ("foo bar" + ("" + "hello"))) == "a")) {
        v1 += v1
        if true {
            var v11 uint32 = 780
            var v12 uint16 = 239
            v13 := GenS2{F0: -90}
            _ = v11
            _ = v12
            _ = v13
        }
    } else {
        var v14 bool = (127 == 1)
        switch -1 {
        case 1:
            switch "hello" {
            case "\n\t":
                var v15 int16 = -509
                _ = v15
            }
        }
        _ = v14
    }
    var v16 float32 = -2.25
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            v3 = v3
            var v19 uint8 = 128
            _ = fi18
            _ = v19
        }
        _ = len(v4)
        switch (-(fi17)) {
        case 0:
            if (-2.25 > 0.1) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    if false {
                        v21 := GenS2{F0: 42}
                        _ = v21
                    } else {
                        fmt.Println(fi17, v1, fi20)
                    }
                    _ = fi20
                }
            } else {
                v22 := []int{(-(fi17)), -2147483648}
                v23 := GenS3{F0: fi17}
                _ = v22
                _ = v23
            }
        case 1:
            for fi24 := 0; fi24 < 11; fi24++ {
                if (42 > 158) {
                    if false {
                        if false {
                            v1 -= (v1 + (v1 << ((v3) & 63)))
                            fmt.Println(v1, v16)
                        }
                        v25 := []int{}
                        _ = v25
                    }
                    switch (fi17 + (fi24 % (((fi17 | fi17)) | 1))) {
                    case 32767:
                        var v26 bool = (!(0 <= 404))
                        _ = v26
                    default:
                        var v27 float32 = (((v16 + v16) + float32(0.1)) / 1.0)
                        _ = v27
                    }
                    var v28 uint32 = 980
                    _ = v28
                }
                v29 := GenS2{F0: 42}
                v3 -= v3
                _ = fi24
                _ = v29
            }
        case 32767:
            v16 = (v16 + ((((v16 + ((v16 + -2.25) + (v16 / v16))) - v16) / 0.1) - float32(1e10)))
        default:
            var v30 uint32 = 4294967295
            _ = v30
        }
        _ = fi17
    }
    v1 += (^(v1))
    var v31 uint16 = 65535
    fmt.Println(v3)
    v1 = (v1 ^ v1)
    if (uint8(v31) == 225) {
        var v32 int = -9223372036854775808
        var v33 uint8 = 127
        var v34 int32 = -2147483648
        _ = v32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}

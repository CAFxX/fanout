package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int) uint32 {
    switch "" {
    case "a":
        fmt.Println(p0)
    default:
        fmt.Println(p0)
    }
    v6 := []int{}
    _ = p0
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 uint8) uint32 {
    var v7 uint16 = 513
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            p1 += uint8(v7)
            var v10 int = (fi8 % (((fi9 ^ int(p0))) | 1))
            _ = fi9
            _ = v10
        }
        fmt.Println(v7)
        _ = fi8
    }
    if (2 == uint64(v7)) {
        v11 := []int{64, 2147483647}
        if ("foo bar" >= "foo bar") {
            p1 -= 128
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 921
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    var v4 uint = 650
    var v5 int64 = -9223372036854775808
    for fi12 := 0; fi12 < 3; fi12++ {
        if false {
            v4 -= v4
            v13 := []int{}
            switch v1 {
            case "\n\t":
                v4 += (v4 << ((v4) & 63))
            case "":
                var v14 uint64 = 0
                _ = v14
            }
            _ = v13
        } else {
            switch v1 {
            case "\n\t":
                var v15 int8 = -41
                _ = v15
            case "":
                var v16 int16 = -887
                _ = v16
            case "a":
                v17 := v3[-532]
                _ = v17
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int64 = (v5 ^ v5)
                _ = fi18
                _ = v19
            }
        }
        _ = fi12
    }
    v20 := []int{-33, 128, 2147483647, 523}
    var v21 int64 = 42
    v22 := GenS1{F0: genFn1(61, 1), F1: -1, F2: v5}
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            if false {
                var v25 int64 = -1
                _ = v25
            }
            v26 := GenS2{F0: int16(v4), F1: (v5 * (v21 - v21))}
            var v27 string = "hello"
            _ = fi24
            _ = v26
            _ = v27
        }
        v28 := GenS2{F0: 279, F1: (v5 - v21)}
        switch (v1 + v1) {
        case "hello":
            var v29 uint32 = genFn1(-1, uint8(v4))
            _ = v29
        case "\n\t":
            for fi30 := 0; fi30 < 17; fi30++ {
                switch (-(fi23)) {
                case -2147483648:
                    v20 = append(v20, fi30)
                case 128:
                    var v31 uint = 1
                    _ = v31
                }
                _ = fi30
            }
        }
        _ = fi23
        _ = v28
    }
    v2[-690] = (((((v1 + (((((((("foo bar" + "hello") + v1) + ((v1 + "hello") + (v1 + "a"))) + v1) + v1) + "\n\t") + "\n\t") + ("" + v1))) + "\n\t") + v1) + v1) + (v1 + v1))
    if true {
        if true {
            for fi32 := 0; fi32 < 4; fi32++ {
                if (1 == -1) {
                    if len(v20) > 0 {
                        v20[0] = -9223372036854775808
                    }
                }
                var v33 uint64 = 4294967295
                for fi34 := 0; fi34 < 5; fi34++ {
                    v35 := GenS3{F0: 32767, F1: 65535}
                    var v36 int = 786
                    switch 42 {
                    case 0:
                        var v37 uint32 = genFn0(fi32)
                        _ = v37
                    case 127:
                        v21 = 42
                    default:
                        var v38 int32 = 42
                        _ = v38
                    }
                    var v39 int = fi32
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v39
                }
                _ = fi32
                _ = v33
            }
            switch 0 {
            case 32767:
                if false {
                    if true {
                        v5 -= v5
                        delete(v3, 42)
                        var v40 uint32 = 2147483647
                        _ = v40
                    } else {
                        fmt.Println(v1)
                    }
                }
            default:
                for fi41 := 0; fi41 < 1; fi41++ {
                    var v42 float32 = 0.0
                    _ = fi41
                    _ = v42
                }
            }
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            for fi44 := 0; fi44 < 5; fi44++ {
                var v45 int16 = 171
                var v46 int8 = 0
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = fi43
        }
        var v47 uint = (v4 / (((v4 * 135)) | 1))
        _ = v47
    }
    var v48 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 int, p1 uint8) int64 {
    var v7 string = "a"
    var v8 uint = 1
    var v9 string = "hello"
    switch 127 {
    case 128:
        switch 596 {
        case 127:
            fmt.Println(p0, v9, v7)
        }
    case -788:
        v7 = "hello"
    }
    if false {
        for fi10 := 0; fi10 < 17; fi10++ {
            switch v7 {
            case "hello":
                v7 = (("a" + (v9 + (v7 + "hello"))) + v7)
            case "foo bar":
                fmt.Println(fi10, fi10)
            case "":
                v11 := []int{(^(fi10)), (^(p0)), -9223372036854775808}
                _ = v11
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int8) uint8 {
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int = fi12
        _ = fi12
        _ = v13
    }
    switch 9223372036854775807 {
    case 128:
        v14 := GenS1{F0: 0}
        _ = v14
    }
    _ = p0
    return 189
}

func main() {
    v1 := []string{("" + ""), "a"}
    var v2 float32 = (3.14159 + -2.25)
    v3 := []uint32{65535}
    var v4 int8 = 1
    var v5 uint64 = 2
    v6 := make(map[uint32]string)
    for fi15 := 0; fi15 < 7; fi15++ {
        delete(v6, uint32(v5))
        switch ("\n\t" + "foo bar") {
        case "hello":
            v16 := GenS2{F0: 0, F1: v5}
            _ = v16
        case "":
            v17 := v6[2]
            _ = v17
        }
        v2 = (0.0 + ((v2 + 1.5) - v2))
        _ = fi15
    }
    v2 -= 1e10
    switch int(v4) {
    case 128:
        v18 := GenS1{F0: -683}
        _ = v18
    case 32767:
        v4 -= 0
    default:
        _ = len(v6)
    }
    var v19 uint16 = 2
    fmt.Println(v19)
    for fi20 := 0; fi20 < 4; fi20++ {
        switch (fi20 >> ((18446744073709551615) & 63)) {
        case 42:
            switch -9223372036854775808 {
            case 490:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 127:
                switch (fi20 | -1) {
                case -2147483648:
                    switch "" {
                    case "\n\t":
                        v19 = (^(v19))
                    case "foo bar":
                        if false {
                            var v21 int8 = -128
                            _ = v21
                        }
                    }
                default:
                    fmt.Println(v4, v2, v19)
                }
            default:
                _ = len(v3)
            }
        case 951:
            var v22 float64 = 0.1
            _ = v22
        case 32767:
            v19 = (v19 % ((v19) | 1))
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            v5 += v5
            v24 := GenS3{F0: v5, F1: v5}
            _ = fi23
            _ = v24
        }
        switch fi20 {
        case 128:
            switch fi20 {
            case -502:
                if true {
                    v25 := GenS2{F0: 0, F1: (v5 ^ (v5 ^ (v5 & (v5 >> ((255) & 63)))))}
                    _ = v25
                }
            case -389:
                switch (fi20 << ((uint(v19)) & 63)) {
                case -1:
                    var v26 int = (fi20 ^ (-(fi20)))
                    _ = v26
                case 690:
                    v27 := v6[42]
                    _ = v27
                case -2147483648:
                    v3 = append(v3, 4294967295)
                default:
                    if true {
                        v28 := GenS2{F0: genFn0(631, 1), F1: v5}
                        switch ("\n\t" + ("\n\t" + "")) {
                        case "":
                            v29 := GenS4{F0: 42, F1: (fi20 | (fi20 / (((fi20 | (fi20 - (fi20 | 955)))) | 1))), F2: 56}
                            _ = v29
                        default:
                            fmt.Println(v4)
                        }
                        v30 := GenS4{F0: 323, F1: fi20, F2: genFn1(v4)}
                        _ = v28
                        _ = v30
                    }
                }
            }
        default:
            var v31 int = fi20
            _ = v31
        }
        _ = fi20
    }
    v32 := GenS2{F0: -676, F1: (^(v5))}
    if true {
        if false {
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int32 = -2147483648
                for fi35 := 0; fi35 < 2; fi35++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v36 float64 = float64(v2)
                    var v37 uint8 = 42
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = fi33
                _ = v34
            }
            var v38 uint32 = 2
            _ = v38
        }
    }
    v5 += v5
    if (255 <= 194) {
        _ = len(v3)
        v4 = -1
        var v39 float64 = 0.1
        _ = v39
    } else {
        var v40 uint64 = v5
        _ = v40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v42 := GenS3{F0: v5, F1: v5}
        v2 = 1.5
        _ = fi41
        _ = v42
    }
    v4 = 127
    if (-52 <= -9223372036854775808) {
        v6[42] = ("a" + "a")
    }
    fmt.Println(v19, v19)
    v19 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

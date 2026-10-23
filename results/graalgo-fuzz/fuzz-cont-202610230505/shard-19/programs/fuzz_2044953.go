package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint }
func genFn0() int {
    switch "\n\t" {
    case "\n\t":
        var v6 uint8 = 158
        _ = v6
    case "hello":
        var v7 uint32 = 2
        _ = v7
    }
    var v8 float64 = 1.5
    fmt.Println(v8)
    v8 = 1.5
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello", "\n\t", "a", ""}
    v3 := make(map[string]string)
    v4 := []uint32{0, 82, 0, 960}
    var v5 int8 = 1
    for fi9 := 0; fi9 < 12; fi9++ {
        switch fi9 {
        case -799:
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 42, F1: 0, F2: 825}
                _ = len(v3)
                fmt.Println(fi9, fi9, fi10)
                _ = fi10
                _ = v11
            }
        case -808:
            delete(v3, "")
        case 183:
            v5 = (v5 * (v5 % (((v5 + (v5 >> ((724) & 63)))) | 1)))
        }
        fmt.Println(fi9, v5)
        _ = fi9
    }
    v5 += (v5 / ((v5) | 1))
    if false {
        if (!(32767 < -1)) {
            if (1000 <= -32768) {
                switch -854 {
                case 4294967295:
                    fmt.Println(v5)
                case 665:
                    for fi12 := 0; fi12 < 7; fi12++ {
                        fmt.Println(fi12)
                        fmt.Println(fi12, v5)
                        _ = fi12
                    }
                }
                switch 42 {
                case -9223372036854775808:
                    if (0.0 >= (1.5 * 3.14159)) {
                        var v13 uint32 = 51
                        fmt.Println(v13)
                        _ = v13
                    }
                case 909:
                    var v14 uint8 = 154
                    _ = v14
                case 127:
                    var v15 bool = ((0.0 != 3.14159) || true)
                    _ = v15
                }
            }
            fmt.Println(v5)
            var v16 float64 = 0.0
            _ = v16
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v5 -= (v5 * (v5 ^ v5))
    if len(v2) > 0 {
        v2[0] = ""
    }
    v3[""] = ("a" + "hello")
    v17 := GenS2{F0: 42}
    var v18 float64 = ((((1e10 - (((((1.5 + (0.0 / float64((float32(1.5) - ((0.0 * (-2.25 / (0.1 / 1.0))) / 3.14159))))) + 0.0) * 1e10) + (1e10 + -2.25)) * (1e10 * (1e10 * ((-2.25 / (-2.25 + 1.5)) + 0.0))))) + float64(0.0)) + 1.5) + 0.1)
    fmt.Println(v5, v5)
    v5 -= v5
    var v19 uint = 42
    v18 -= ((((v18 - float64((-2.25 / 1.5))) - (((v18 * (3.14159 + v18)) + 3.14159) + 0.1)) + v18) + v18)
    if (42 <= 336) {
        for fi20 := 0; fi20 < 20; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS3{F0: 397, F1: v19}
                v19 += 2
                var v23 uint = 264
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
    }
    switch ((("a" + "foo bar") + "") + ("foo bar" + "")) {
    case "":
        v19 = 65535
    case "hello":
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 float64 = 1.5
            switch (fi24 << ((2) & 63)) {
            case 628:
                var v26 uint32 = 65535
                _ = v26
            case 977:
                var v27 int16 = -969
                _ = v27
            case 128:
                var v28 uint16 = 0
                _ = v28
            }
            _ = fi24
            _ = v25
        }
    case "foo bar":
        if false {
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 uint32 = 816
                _ = fi29
                _ = v30
            }
            var v31 uint8 = 191
            _ = v31
        } else {
            switch 128 {
            case 0:
                v5 -= -69
            }
            if false {
                switch 127 {
                case 42:
                    var v32 string = (("\n\t" + "a") + ("a" + ("\n\t" + "hello")))
                    _ = v32
                case -2147483648:
                    v19 = v19
                case 9223372036854775807:
                    switch "foo bar" {
                    case "a":
                        for fi33 := 0; fi33 < 14; fi33++ {
                            v19 += v19
                            v5 = v5
                            _ = fi33
                        }
                    }
                default:
                    fmt.Println(v19, v5, v19)
                }
            }
        }
    }
    delete(v3, (("" + "a") + "\n\t"))
    v34 := GenS2{F0: 0}
    v5 += (v5 % (((v5 | v5)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}

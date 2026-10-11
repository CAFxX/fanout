package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint }

func main() {
    v1 := []string{(((("" + ("a" + "a")) + "foo bar") + "hello") + "hello"), "", ("\n\t" + "\n\t"), ("foo bar" + "")}
    var v2 float32 = (float32((float64(0.0) / ((3.14159 / 1e10) / 3.14159))) - 1e10)
    v3 := make(map[uint32]bool)
    v4 := []int{0, -2147483648, 461}
    v5 := []int{351}
    var v6 int32 = -1
    switch 4294967295 {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(fi7, v2)
            _ = fi7
        }
    case -2147483648:
        if (0 >= 9223372036854775807) {
            for fi8 := 0; fi8 < 1; fi8++ {
                v2 -= v2
                _ = len(v3)
                var v9 int32 = 771
                _ = fi8
                _ = v9
            }
            v6 = (v6 / (((v6 / ((v6) | 1))) | 1))
        }
    case 42:
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(v2, fi10, fi10)
            switch (fi10 - fi10) {
            case 757:
                v2 -= 1e10
            case -2147483648:
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint = 0
                    v6 -= (v6 % ((0) | 1))
                    var v13 string = "foo bar"
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 584:
                if true {
                    v2 = (v2 - v2)
                } else {
                    v14 := GenS1{F0: (("a" + "a") + (((("foo bar" + (("foo bar" + ("" + ("a" + "hello"))) + "")) + "a") + "hello") + "")), F1: 785, F2: 32767}
                    v15 := GenS1{F0: "hello", F1: 668, F2: 0}
                    _ = v14
                    _ = v15
                }
            default:
                v16 := GenS2{F0: (v6 >> ((65535) & 63))}
                _ = v16
            }
            _ = fi10
        }
    }
    if (1000 <= -1) {
        var v17 string = ""
        _ = v17
    } else {
        switch -9223372036854775808 {
        case 4294967295:
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 int8 = 78
                _ = fi18
                _ = v19
            }
        case 244:
            for fi20 := 0; fi20 < 20; fi20++ {
                fmt.Println(fi20, fi20, fi20)
                _ = fi20
            }
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            v6 = v6
            _ = fi21
        }
        v22 := GenS1{F0: "hello", F1: 42, F2: 42}
        _ = v22
    }
    if true {
        v23 := GenS3{F0: "\n\t", F1: 18446744073709551615}
        _ = v23
    }
    var v24 float32 = -2.25
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v25 float32 = float32(1.5)
        _ = v25
    case -1:
        switch "foo bar" {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        if true {
            _ = len(v1)
            var v26 uint64 = 457
            v6 -= -2147483648
            _ = v26
        } else {
            v24 -= v24
            v24 += v2
            var v27 int = -75
            _ = v27
        }
    }
    var v28 uint = 843
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint8 = 127
        v6 = (v6 & 435)
        _ = fi29
        _ = v30
    }
    v28 = (v28 * v28)
    if (v6 <= (v6 >> ((v28) & 63))) {
        v31 := GenS1{F0: "foo bar", F1: 255, F2: 75}
        if ("\n\t" >= "\n\t") {
            var v32 float64 = 1e10
            var v33 int16 = 32767
            _ = v32
            _ = v33
        }
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

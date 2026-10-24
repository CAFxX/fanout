package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    switch ("" + "hello") {
    case "a":
        if false {
            for fi5 := 0; fi5 < 4; fi5++ {
                p0 = p0
                fmt.Println(p0, fi5)
                if true {
                    p0 = p0
                    p0 = p0
                    var v6 string = "a"
                    _ = v6
                } else {
                    switch "hello" {
                    case "\n\t":
                        v7 := GenS1{F0: 65535, F1: -1}
                        _ = v7
                    case "a":
                        v8 := GenS1{F0: 32767, F1: 1}
                        _ = v8
                    }
                }
                _ = fi5
            }
            switch -570 {
            case 1:
                var v9 float32 = 1e10
                _ = v9
            case 127:
                p0 = (^(p0))
            default:
                var v10 int = -398
                _ = v10
            }
        } else {
            v11 := GenS1{F0: 473, F1: -1}
            _ = v11
        }
    case "foo bar":
        if (!true) {
            fmt.Println(p0)
            for fi12 := 0; fi12 < 4; fi12++ {
                p0 -= (p0 ^ 161)
                _ = fi12
            }
        }
    }
    p0 = (p0 & p0)
    fmt.Println(p0)
    _ = p0
    return false
}
func genFn1() uint16 {
    var v13 float32 = 3.14159
    if (true || (!(-32768 < 1))) {
        v14 := []int{-2147483648, -9223372036854775808}
        _ = len(v14)
        _ = v14
    }
    var v15 uint8 = 43
    switch (("a" + "a") + "a") {
    case "a":
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 int8 = 0
            for fi18 := 0; fi18 < 7; fi18++ {
                v13 = ((-2.25 + (-2.25 / 1.0)) + 3.14159)
                v19 := GenS1{F0: 32767, F1: -762}
                v20 := []int{(fi16 << ((2) & 63)), 128, (fi18 >> ((0) & 63))}
                fmt.Println(v13, fi18, fi18)
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    _ = v13
    _ = v15
    return 307
}

func main() {
    var v1 bool = false
    var v2 uint16 = 416
    v3 := make(map[int]string)
    var v4 float32 = (float32(((1e10 / 1e10) + (((0.0 / 1.5) * 3.14159) / 0.1))) - -2.25)
    for fi21 := 0; fi21 < 5; fi21++ {
        switch fi21 {
        case 1:
            switch ("foo bar" + "a") {
            case "a":
                var v22 int64 = 369
                _ = v22
            case "hello":
                fmt.Println(v4, fi21)
            case "":
                fmt.Println(fi21, fi21, v4)
            default:
                v23 := GenS1{F0: v2, F1: -625}
                _ = v23
            }
        case 127:
            if (!(v1 || (!v1))) {
                if (!v1) {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v25 := GenS2{F0: -639, F1: 1}
                        var v26 int = (^(fi21))
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    var v27 float64 = float64(float32(((1e10 / (3.14159 + 3.14159)) - 3.14159)))
                    v28 := []int{fi21, fi21}
                    _ = v27
                    _ = v28
                }
            }
        case 4294967295:
            for fi29 := 0; fi29 < 11; fi29++ {
                v30 := GenS1{F0: (v2 ^ (^(v2))), F1: 22}
                _ = fi29
                _ = v30
            }
        default:
            v3[(fi21 / (((fi21 | 128)) | 1))] = "\n\t"
        }
        v1 = v1
        _ = fi21
    }
    v31 := GenS1{F0: v2, F1: 32767}
    var v32 uint16 = (v2 >> ((18446744073709551615) & 63))
    v33 := GenS3{F0: 202, F1: int64(v32)}
    v34 := []int{-2147483648, -2147483648}
    if v1 {
        fmt.Println(v2)
        fmt.Println(v1, v2)
    }
    if (v1 && v1) {
        v35 := GenS3{F0: 65535, F1: int64(v2)}
        _ = v35
    }
    var v36 uint = 472
    if false {
        fmt.Println(v2, v36, v2)
        var v37 int = 2147483647
        _ = v37
    } else {
        var v38 int64 = -209
        var v39 uint = v36
        v40 := GenS1{F0: v32, F1: -1}
        _ = v38
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        v1 = (!(1.5 != v4))
        fmt.Println(v4, fi41)
        _ = fi41
    }
    delete(v3, -481)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
}

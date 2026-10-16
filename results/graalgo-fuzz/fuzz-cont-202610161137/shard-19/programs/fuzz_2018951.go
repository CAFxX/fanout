package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 uint8, p2 uint8) uint16 {
    v5 := GenS1{F0: 2, F1: 4294967295, F2: 0}
    var v6 float32 = (3.14159 + float32((3.14159 - (-2.25 / 1.5))))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 629
}
func genFn1(p0 uint16, p1 int, p2 bool) int {
    v7 := []int{(p1 - p1)}
    if p2 {
        var v8 int8 = 66
        _ = v8
    }
    var v9 uint16 = genFn0(1, 139, 117)
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint = 229
        if true {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            var v12 int16 = -268
            var v13 uint16 = (^(p0))
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return p1
}

func main() {
    var v1 float32 = (0.1 / 3.14159)
    var v2 int16 = 1
    v3 := []string{"\n\t", "a", ((("hello" + "foo bar") + ("hello" + "a")) + ""), "hello"}
    var v4 uint = 1
    if true {
        if len(v3) > 0 {
            v3[0] = (("hello" + "foo bar") + ("" + (("a" + "") + ("\n\t" + ""))))
        }
        v1 = float32(1e10)
        var v14 int = 4294967295
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        switch (fi15 ^ fi15) {
        case 32768:
            v16 := GenS1{F0: 255, F1: 83, F2: 42}
            _ = v16
        case -1:
            v17 := GenS1{F0: 1, F1: 0, F2: 18446744073709551615}
            _ = v17
        }
        _ = fi15
    }
    var v18 int8 = 42
    v2 = -1
    switch 127 {
    case -9223372036854775808:
        var v19 uint = v4
        _ = v19
    case 1:
        v2 += 0
    case 9223372036854775807:
        switch "\n\t" {
        case "":
            for fi20 := 0; fi20 < 4; fi20++ {
                v4 -= v4
                v21 := GenS1{F0: 2, F1: 2, F2: uint64(v4)}
                v4 -= 18446744073709551615
                _ = fi20
                _ = v21
            }
        case "a":
            var v22 string = "hello"
            _ = v22
        case "hello":
            switch "" {
            case "hello":
                switch -853 {
                case 42:
                    var v23 uint16 = 42
                    _ = v23
                }
            case "":
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 float64 = (0.1 / 1e10)
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                v1 -= v1
            }
        default:
            v26 := GenS1{F0: 128, F1: 588, F2: 1}
            _ = v26
        }
    default:
        v27 := GenS1{F0: 255, F1: 2147483648, F2: 4294967295}
        _ = v27
    }
    v28 := GenS2{F0: ("\n\t" + "\n\t")}
    v29 := GenS1{F0: 2, F1: 0, F2: 4294967295}
    var v30 string = "\n\t"
    v18 = (v18 + (v18 * 0))
    switch 2147483647 {
    case 4294967295:
        var v31 bool = true
        _ = v31
    }
    var v32 uint16 = 255
    if len(v3) > 0 {
        v3[0] = v30
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v2 += v2
        if true {
            v1 = 1e10
            if false {
                fmt.Println(v18)
            } else {
                var v34 int8 = (v18 / ((v18) | 1))
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 16; fi35++ {
                if false {
                    var v36 int8 = (v18 / (((v18 ^ v18)) | 1))
                    var v37 uint16 = v32
                    var v38 int8 = (v36 * 1)
                    _ = v36
                    _ = v37
                    _ = v38
                } else {
                    var v39 string = (("" + (((((v30 + v30) + "") + "hello") + v30) + v30)) + (((("a" + (("a" + v30) + "hello")) + v30) + (v30 + "")) + ("" + v30)))
                    var v40 bool = ("hello" >= "")
                    fmt.Println(v30, v39)
                    _ = v39
                    _ = v40
                }
                _ = fi35
            }
        }
        for fi41 := 0; fi41 < 4; fi41++ {
            v4 = v4
            _ = fi41
        }
        _ = fi33
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v32 -= genFn0(v4, 2, 42)
    for fi42 := 0; fi42 < 1; fi42++ {
        if (false && (true || false)) {
            v32 = 586
        }
        if (18446744073709551615 <= 42) {
            for fi43 := 0; fi43 < 20; fi43++ {
                switch 0 {
                case 42:
                    fmt.Println(v30, v4, v18)
                }
                _ = fi43
            }
            v44 := GenS3{F0: (65535 > (v32 / ((445) | 1)))}
            _ = v44
        } else {
            fmt.Println(v1)
            _ = len(v3)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        fmt.Println(v2)
        _ = fi42
    }
    var v45 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v45
}

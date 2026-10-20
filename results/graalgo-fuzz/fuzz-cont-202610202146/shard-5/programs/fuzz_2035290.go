package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 uint16) int {
    v5 := GenS1{F0: "hello", F1: p0}
    if false {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0)
        case "a":
            switch "" {
            case "\n\t":
                fmt.Println(p0)
            default:
                v6 := []int{(p0 >> ((65535) & 63)), (p0 - (p0 * 127)), 42}
                _ = v6
            }
        }
        if false {
            switch (p0 ^ 1) {
            case 32767:
                if (!(1 <= 4294967295)) {
                    for fi7 := 0; fi7 < 8; fi7++ {
                        v8 := GenS1{F0: ("foo bar" + "foo bar"), F1: fi7}
                        for fi9 := 0; fi9 < 17; fi9++ {
                            p0 = (-(fi9))
                            switch ("a" + "a") {
                            case "a":
                                if (42 >= 4294967295) {
                                    var v10 uint16 = 546
                                    _ = v10
                                } else {
                                    switch ("hello" + (("foo bar" + ("" + "a")) + (("foo bar" + "hello") + "a"))) {
                                    case "hello":
                                        fmt.Println(p1, p0, p1)
                                    }
                                }
                            }
                            _ = fi9
                        }
                        var v11 uint = 731
                        _ = fi7
                        _ = v11
                        _ = v8
                    }
                }
            case -1:
                fmt.Println(p0, p1)
            case -2147483648:
                v12 := []int{(p0 % (((p0 | p0)) | 1)), p0, (-(p0)), p0}
                _ = v12
            }
            var v13 uint16 = (p1 | 2)
            fmt.Println(v13)
            _ = v13
        }
        var v14 float32 = (1.5 / (1.5 / (((1.5 * (1e10 * 1.5)) + ((float32((3.14159 + 3.14159)) * -2.25) / float32(-2.25))) - -2.25)))
        _ = v14
    }
    v15 := []int{54, p0}
    p0 = 9223372036854775807
    var v16 string = "foo bar"
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v5
    return p0
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 32767
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v17 bool = false
    fmt.Println(v17, v2)
    for fi18 := 0; fi18 < 5; fi18++ {
        v2 += (^(fi18))
        v17 = v17
        _ = fi18
    }
    v2 = v2
    var v19 uint = 88
    fmt.Println(v17, v19)
    v17 = false
    v17 = true
    var v20 int8 = 104
    switch "hello" {
    case "\n\t":
        switch v2 {
        case 4294967295:
            v21 := GenS2{F0: (v2 * (-(v2))), F1: 128}
            _ = v21
        default:
            v20 -= (^(v20))
        }
    case "hello":
        var v22 int = v2
        _ = v22
    case "a":
        var v23 uint16 = 523
        _ = v23
    default:
        if v17 {
            fmt.Println(v2)
        }
    }
    var v24 string = "hello"
    fmt.Println(v24)
    v20 = (v20 | (v20 - v20))
    v25 := GenS3{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}

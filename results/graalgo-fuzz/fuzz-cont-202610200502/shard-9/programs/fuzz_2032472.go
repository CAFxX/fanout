package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16 }

func main() {
    var v1 float32 = float32(((0.1 + (1e10 + 1e10)) + ((((1e10 + 0.1) / -2.25) * 1e10) - 3.14159)))
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    var v4 uint16 = 42
    var v5 float32 = 1.5
    v1 = 1.5
    v4 += v4
    v3[-9223372036854775808] = 0
    switch v2 {
    case "":
        v2 = ""
    case "hello":
        if false {
            var v6 uint = 474
            _ = v6
        } else {
            if true {
                v7 := GenS1{F0: ""}
                _ = v7
            } else {
                var v8 uint8 = 128
                _ = v8
            }
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 float32 = 3.14159
                    v12 := []int{}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: (v2 + v2)}
                _ = fi9
                _ = v13
            }
        }
    case "a":
        var v14 int64 = 42
        _ = v14
    default:
        if (!(174 <= 48)) {
            v4 -= v4
            v4 = v4
            var v15 string = "foo bar"
            _ = v15
        }
    }
    if true {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int32 = 2147483647
            v5 -= v5
            _ = fi16
            _ = v17
        }
        var v18 uint8 = 2
        var v19 uint = 65535
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{4294967295, (fi20 + 4294967295), fi20}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int = fi22
        if (-8 == -32768) {
            switch "foo bar" {
            case "foo bar":
                v4 = (v4 / ((1) | 1))
            default:
                v2 = v2
            }
        } else {
            switch 32767 {
            case 0:
                fmt.Println(v5, v2)
            default:
                if (false && false) {
                    var v24 uint8 = 1
                    _ = v24
                }
            }
            v25 := GenS1{F0: v2}
            switch "\n\t" {
            case "":
                v4 -= 1
            case "hello":
                v1 = ((-2.25 + float32(-2.25)) - 0.1)
            case "a":
                v26 := GenS1{F0: (v2 + v2)}
                _ = v26
            }
            _ = v25
        }
        if ("\n\t" >= ("hello" + v2)) {
            if true {
                v2 = v2
                switch v23 {
                case 348:
                    var v27 int16 = 924
                    _ = v27
                case -2147483648:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                default:
                    if true {
                        var v29 string = v2
                        v1 = v1
                        var v30 float32 = ((v5 + ((v5 + -2.25) - (v5 - 1.5))) / 0.1)
                        _ = v29
                        _ = v30
                    }
                }
                _ = len(v3)
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                v23 = (-(v23))
                fmt.Println(v4)
                v1 = v5
                _ = fi31
            }
            v32 := []int{}
            _ = v32
        } else {
            switch (v23 ^ (v23 + v23)) {
            case 1:
                var v33 uint64 = 1
                _ = v33
            case 0:
                v4 = v4
            case -579:
                v1 = (v5 - (((v1 + (v1 + float32(1.5))) / (v1 / v1)) - (v1 * 1e10)))
            }
            for fi34 := 0; fi34 < 5; fi34++ {
                v35 := GenS1{F0: v2}
                if true {
                    var v36 float32 = (float32((0.0 / 1e10)) / v1)
                    _ = v36
                } else {
                    v37 := GenS2{F0: 169, F1: (v2 + v2)}
                    fmt.Println(v2)
                    _ = v37
                }
                v38 := GenS3{F0: 42}
                _ = fi34
                _ = v35
                _ = v38
            }
        }
        _ = fi22
        _ = v23
    }
    v39 := []int{-2147483648, 1, -2147483648, 727}
    fmt.Println(v2, v2, v2)
    _ = len(v39)
    v4 = (v4 / ((v4) | 1))
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 int8 = 42
        _ = fi40
        _ = v41
    }
    if (!false) {
        v4 -= 211
        for fi42 := 0; fi42 < 13; fi42++ {
            var v43 int = (^(fi42))
            if len(v39) >= 2 {
                v39 = v39[1:len(v39)]
            }
            delete(v3, (fi42 / ((v43) | 1)))
            _ = fi42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}

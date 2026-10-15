package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint64; F2 string }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int = 32768
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 44}
        if ((!(!false)) || (!false)) {
            v3 += (^(v3))
        }
        _ = fi4
        _ = v5
    }
    if (false || false) {
        var v6 int32 = -858
        for fi7 := 0; fi7 < 20; fi7++ {
            v3 = (fi7 / ((v3) | 1))
            _ = fi7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            switch ((("hello" + "a") + "\n\t") + "hello") {
            case "foo bar":
                switch (fi8 / (((v3 << ((216) & 63))) | 1)) {
                case -730:
                    if (411 < -1) {
                        v9 := []int{(^(v3))}
                        _ = v9
                    } else {
                        switch (v3 / ((v3) | 1)) {
                        case 9223372036854775807:
                            var v10 float64 = (-2.25 + -2.25)
                            _ = v10
                        }
                        var v11 int64 = 0
                        fmt.Println(v3, v11)
                        _ = v11
                    }
                case -1:
                    var v12 uint32 = 381
                    _ = v12
                case 9223372036854775807:
                    v6 = (v6 + (v6 % ((416) | 1)))
                default:
                    var v13 uint32 = 2147483648
                    _ = v13
                }
            default:
                fmt.Println(v3)
            }
            if (255 == 255) {
                switch (fi8 / (((fi8 & 334)) | 1)) {
                case 4294967295:
                    v1[(v3 - 32768)] = ("\n\t" + "hello")
                case 32767:
                    v14 := GenS2{F0: -114, F1: 631, F2: ""}
                    _ = v14
                }
                switch "foo bar" {
                case "a":
                    fmt.Println(v3)
                }
                v15 := []int{v3, (v3 - fi8)}
                _ = v15
            }
            _ = fi8
        }
        _ = v6
    }
    if ("foo bar" < ((("hello" + "") + "hello") + ("foo bar" + "a"))) {
        var v16 uint16 = 255
        _ = v16
    }
    var v17 uint16 = 536
    v18 := []int{v3, v3, 127, -957}
    var v19 uint64 = 2
    if true {
        v20 := GenS3{F0: v19, F1: 42, F2: ("hello" + (("\n\t" + "hello") + "hello"))}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 bool = false
            _ = fi21
            _ = v22
        }
        switch "foo bar" {
        case "\n\t":
            if (42 < int32(v19)) {
                if true {
                    v2[v3] = true
                } else {
                    var v23 uint = 380
                    v24 := GenS4{F0: (("a" + "\n\t") + "\n\t")}
                    _ = len(v18)
                    _ = v23
                    _ = v24
                }
                var v25 uint32 = 2147483647
                _ = v25
            }
        case "a":
            fmt.Println(v19)
        default:
            v19 = (v19 & (v19 ^ v19))
        }
    }
    if true {
        var v26 bool = true
        _ = v26
    } else {
        v17 += 585
        v27 := GenS1{F0: 0}
        var v28 string = ""
        _ = v27
        _ = v28
    }
    var v29 int32 = 0
    if (false && (1 != 42)) {
        v17 += (v17 / ((v17) | 1))
        v30 := GenS5{F0: ((false && ((3.14159 == 0.1) && true)) && false), F1: true}
        v31 := GenS1{F0: 222}
        _ = v30
        _ = v31
    } else {
        var v32 uint64 = (v19 % (((v19 & v19)) | 1))
        _ = v32
    }
    v3 -= (v3 << ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 int, p1 uint) uint32 {
    fmt.Println(p0)
    p0 += p0
    v5 := []int{(^(p0))}
    for fi6 := 0; fi6 < 10; fi6++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            fmt.Println(p1, p0, p1)
        }
        _ = len(v5)
        var v7 string = ""
        _ = fi6
        _ = v7
    }
    switch (p0 ^ (^(p0))) {
    case 971:
        if (183 <= 0) {
            if ((-128 != 127) || false) {
                switch (p0 * (p0 + 642)) {
                case -2147483648:
                    for fi8 := 0; fi8 < 11; fi8++ {
                        p0 = p0
                        if true {
                            if (521 <= 776) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(p1, p0, p0)
                            } else {
                                p1 += p1
                            }
                            var v9 int32 = 737
                            _ = v9
                        }
                        var v10 uint8 = 42
                        _ = fi8
                        _ = v10
                    }
                case 0:
                    p0 = -560
                case -9223372036854775808:
                    switch p0 {
                    case 9223372036854775807:
                        for fi11 := 0; fi11 < 6; fi11++ {
                            p0 += 128
                            _ = len(v5)
                            _ = fi11
                        }
                    }
                }
                switch "\n\t" {
                case "a":
                    var v12 int = (p0 >> ((p1) & 63))
                    _ = v12
                case "":
                    var v13 string = ""
                    _ = v13
                }
            }
        }
    case 0:
        _ = len(v5)
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    v1 := []string{("hello" + ("hello" + "\n\t")), "foo bar", "foo bar", ("" + "a")}
    var v2 uint32 = 4294967295
    v3 := []string{}
    var v4 uint64 = 530
    switch "hello" {
    case "foo bar":
        var v14 int = 32768
        _ = v14
    }
    switch "foo bar" {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = "a"
        }
    case "a":
        _ = len(v3)
    }
    v2 = v2
    if (65535 < 916) {
        if false {
            var v15 bool = false
            _ = v15
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int = (^(fi16))
                switch fi16 {
                case 2147483647:
                    v18 := GenS1{F0: 65535, F1: 630}
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            var v19 uint8 = 0
            _ = v19
        }
        var v20 bool = true
        _ = v20
    }
    var v21 uint8 = 128
    var v22 float64 = 1e10
    v23 := GenS1{F0: 0, F1: -646}
    v4 += (v4 - (v4 * (v4 / (((v4 & v4)) | 1))))
    v24 := GenS1{F0: 2, F1: 543}
    fmt.Println(v4)
    v25 := GenS2{F0: v2, F1: "\n\t"}
    switch "hello" {
    case "\n\t":
        v1 = append(v1, "")
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v26 int32 = 1
    if false {
        var v27 uint32 = v2
        if (!(2 <= 42)) {
            fmt.Println(v4)
        }
        _ = v27
    }
    var v28 int64 = 806
    switch (("\n\t" + "") + "foo bar") {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "a"
        }
    default:
        fmt.Println(v2)
    }
    v2 = uint32(v21)
    fmt.Println(v2, v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}

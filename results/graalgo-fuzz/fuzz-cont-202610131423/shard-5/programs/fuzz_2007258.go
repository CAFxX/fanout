package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"\n\t", "foo bar"}
    v3 := make(map[int]bool)
    v1 = 42
    if false {
        var v4 int16 = 1
        _ = v4
    }
    var v5 int = -9223372036854775808
    switch 2147483647 {
    case -936:
        if (!false) {
            var v6 int16 = 0
            v6 = (v6 / ((32767) | 1))
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int32 = -1
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            var v9 int8 = -1
            v5 -= (v5 / (((v5 | v5)) | 1))
            _ = v9
        }
    case 9223372036854775807:
        for fi10 := 0; fi10 < 12; fi10++ {
            if (1 != 42) {
                fmt.Println(v1)
                var v11 uint8 = 255
                _ = v11
            } else {
                var v12 uint32 = 553
                fmt.Println(v1)
                _ = v12
            }
            _ = fi10
        }
    default:
        var v13 int32 = 995
        _ = v13
    }
    if (false || ((v5 | v5) > v5)) {
        if (1 == -867) {
            v5 -= (^(v5))
        }
    }
    if true {
        fmt.Println(v1, v1)
        v1 = v1
        var v14 int = v5
        _ = v14
    } else {
        fmt.Println(v1, v1)
    }
    var v15 bool = false
    if (42 < 546) {
        switch (v5 << ((701) & 63)) {
        case 0:
            v5 += -443
        case 1:
            switch (v5 & v5) {
            case -1:
                var v16 int = v5
                _ = v16
            }
        case 32767:
            var v17 int32 = 0
            _ = v17
        default:
            v1 -= v1
        }
    }
    var v18 uint16 = 0
    v1 += (v1 << ((18446744073709551615) & 63))
    var v19 uint32 = 1
    switch "foo bar" {
    case "":
        v20 := GenS1{F0: 1, F1: 2}
        _ = v20
    case "\n\t":
        var v21 int64 = (v1 % (((v1 / (((v1 % (((v1 << ((0) & 63))) | 1))) | 1))) | 1))
        _ = v21
    case "foo bar":
        v15 = v15
    }
    if v15 {
        v2 = append(v2, (((("\n\t" + "hello") + "foo bar") + "a") + ("" + ("" + ("" + ("" + ""))))))
    } else {
        for fi22 := 0; fi22 < 9; fi22++ {
            if (18446744073709551615 >= 0) {
                v15 = v15
                delete(v3, -9223372036854775808)
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 uint64 = 992
                    switch "a" {
                    case "hello":
                        v19 = (v19 % ((v19) | 1))
                    case "foo bar":
                        v19 = (v19 >> ((1) & 63))
                    case "\n\t":
                        fmt.Println(v19)
                    }
                    _ = fi23
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int = (fi25 / (((-(fi22))) | 1))
                v18 = 42
                var v27 float64 = -2.25
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}

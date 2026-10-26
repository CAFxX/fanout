package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint64; F2 bool }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := []string{(("" + "\n\t") + "\n\t")}
    if (v1 >= (v1 ^ v1)) {
        if (9223372036854775807 < -626) {
            for fi5 := 0; fi5 < 7; fi5++ {
                fmt.Println(v1)
                _ = fi5
            }
            if true {
                v1 += v1
            } else {
                var v6 uint8 = 128
                v7 := GenS1{F0: 65535, F1: 42}
                _ = v6
                _ = v7
            }
        }
        switch "foo bar" {
        case "":
            var v8 float32 = 1e10
            _ = v8
        case "foo bar":
            switch "foo bar" {
            case "hello":
                fmt.Println(v1)
            case "\n\t":
                v9 := v2[-447]
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 int64 = -839
                    _ = fi10
                    _ = v11
                }
            }
        case "hello":
            if true {
                var v12 int = 0
                _ = v12
            } else {
                v13 := GenS2{F0: "hello", F1: 2, F2: 4294967295}
                _ = v13
            }
        }
        v1 = v1
    } else {
        fmt.Println(v1)
    }
    delete(v2, 2147483647)
    if true {
        v1 -= v1
        v1 = (v1 + (v1 | (v1 >> ((255) & 63))))
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = -2147483648
            switch "" {
            case "foo bar":
                v15 = (v15 + (v15 >> ((0) & 63)))
            }
            if (623 == 2) {
                v16 := GenS1{F0: 42, F1: v15}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        v1 = v1
        v1 = (v1 * (v1 + 255))
    }
    v1 -= v1
    var v17 int64 = -9223372036854775808
    if (!(true && (0 == 65535))) {
        fmt.Println(v1)
        v17 = 42
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: 18446744073709551615, F1: 2147483647}
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS2{F0: ("\n\t" + "foo bar"), F1: 42, F2: uint64(v1)}
    if (-2.25 > 1.5) {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS3{F0: 18446744073709551615}
            _ = fi21
            _ = v22
        }
    }
    v17 = v17
    var v23 uint32 = 2
    v23 -= (v23 + v23)
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            fmt.Println(v1)
            v1 = (v1 - (v1 >> ((557) & 63)))
            _ = fi25
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v23, v23, v17)
            var v27 bool = ((332 == -930) && false)
            fmt.Println(v27, v17)
            if v27 {
                var v28 int64 = -347
                v3[771] = (fi24 / ((fi24) | 1))
                _ = v28
            } else {
                if (fi26 < fi24) {
                    switch (fi26 << ((255) & 63)) {
                    case 840:
                        fmt.Println(v1, v27, v17)
                    default:
                        v29 := GenS1{F0: 255, F1: -895}
                        _ = v29
                    }
                } else {
                    v23 -= 1
                }
                if (-1 >= -3) {
                    switch fi24 {
                    case -171:
                        v23 = (v23 + 2147483647)
                    }
                    var v30 int = -124
                    v17 = v17
                    _ = v30
                }
                v17 -= (v17 | (v17 / ((947) | 1)))
            }
            _ = fi26
            _ = v27
        }
        _ = fi24
    }
    var v31 int64 = (v17 ^ v17)
    v32 := GenS4{F0: "", F1: 0, F2: false}
    var v33 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}

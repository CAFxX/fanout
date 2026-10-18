package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", ("" + ("foo bar" + "hello")), "hello"}
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    fmt.Println(v1)
    v1 -= v1
    for fi6 := 0; fi6 < 16; fi6++ {
        v1 += (v1 % ((503) | 1))
        var v7 bool = false
        for fi8 := 0; fi8 < 9; fi8++ {
            v2[2147483648] = fi8
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v1 -= (v1 ^ v1)
    v1 = (v1 >> ((4294967295) & 63))
    _ = len(v4)
    v1 = 2
    fmt.Println(v1)
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint64 = 844
        v11 := GenS1{F0: (fi9 * fi9), F1: 2}
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "" {
    case "a":
        for fi12 := 0; fi12 < 5; fi12++ {
            switch "\n\t" {
            case "foo bar":
                switch fi12 {
                case -1:
                    v1 -= 42
                }
            case "hello":
                v13 := GenS1{F0: fi12, F1: 363}
                _ = v13
            }
            var v14 int16 = 32767
            _ = fi12
            _ = v14
        }
    case "\n\t":
        switch "hello" {
        case "hello":
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := GenS1{F0: (fi15 >> ((4294967295) & 63)), F1: uint(v1)}
                v1 = 2
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi15
                _ = v16
            }
        }
    case "":
        v17 := GenS2{F0: -9223372036854775808}
        _ = v17
    }
    v1 = v1
    for fi18 := 0; fi18 < 2; fi18++ {
        switch 0 {
        case 32767:
            var v19 int = 127
            _ = v19
        }
        if (1 >= 42) {
            if (!true) {
                v1 += (v1 ^ v1)
                fmt.Println(v1)
                switch "hello" {
                case "\n\t":
                    var v20 int = (fi18 * fi18)
                    _ = v20
                }
            }
            v1 += 0
            var v21 int = (fi18 / ((-822) | 1))
            _ = v21
        } else {
            var v22 string = "hello"
            _ = v22
        }
        _ = fi18
    }
    if true {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 int = 42
            var v25 int = fi23
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 int64 = -9223372036854775808
        v26 += v26
        _ = v26
    }
    var v27 uint32 = 357
    switch 0 {
    case 9223372036854775807:
        if (!(!(!true))) {
            v1 = (v1 >> ((392) & 63))
            for fi28 := 0; fi28 < 16; fi28++ {
                v29 := GenS3{F0: ""}
                _ = fi28
                _ = v29
            }
        } else {
            var v30 int8 = 42
            for fi31 := 0; fi31 < 16; fi31++ {
                _ = len(v3)
                for fi32 := 0; fi32 < 11; fi32++ {
                    delete(v2, (v27 & 65535))
                    _ = fi32
                }
                v1 = (^(v1))
                _ = fi31
            }
            _ = v30
        }
    case -9223372036854775808:
        if (128 <= 255) {
            v27 = 4294967295
        }
    }
    delete(v4, (^(v27)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

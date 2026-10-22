package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }

func main() {
    var v1 int = 753
    var v2 int16 = 42
    var v3 int64 = 300
    if (75 < 255) {
        for fi4 := 0; fi4 < 1; fi4++ {
            if true {
                var v5 uint16 = 1
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v5, v5)
                case "foo bar":
                    v6 := []int{(-(v1)), (v1 + -9223372036854775808)}
                    _ = v6
                }
                fmt.Println(fi4, v1, v2)
                _ = v5
            }
            _ = fi4
        }
        var v7 float32 = -2.25
        var v8 int16 = (v2 << ((4294967295) & 63))
        _ = v7
        _ = v8
    } else {
        v9 := []int{(v1 >> ((18446744073709551615) & 63)), (v1 / ((42) | 1)), v1, (v1 / ((v1) | 1))}
        v1 = (v1 / ((1) | 1))
        switch (v1 | 1) {
        case 4294967295:
            v10 := GenS1{F0: 18, F1: "a", F2: (false && true)}
            _ = v10
        case 42:
            fmt.Println(v3)
        case -1:
            v3 -= 42
        }
        _ = v9
    }
    v3 = -1
    for fi11 := 0; fi11 < 18; fi11++ {
        if false {
            v12 := GenS1{F0: -1, F1: "\n\t", F2: (22 != 0)}
            switch "a" {
            case "\n\t":
                if false {
                    v1 += fi11
                    v3 += 42
                } else {
                    fmt.Println(v1, v3)
                }
            }
            _ = v12
        }
        _ = fi11
    }
    if true {
        if (v1 > 9223372036854775807) {
            v1 = 2147483647
            if (!true) {
                var v13 int16 = (v2 & (v2 | (v2 % ((v2) | 1))))
                _ = v13
            } else {
                var v14 uint = 0
                v15 := GenS1{F0: -61, F1: "\n\t", F2: false}
                v14 = 0
                _ = v14
                _ = v15
            }
        }
    }
    var v16 uint = 621
    v17 := []int{32768, v1}
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int16 = 0
        v16 += 128
        _ = fi18
        _ = v19
    }
    v17 = append(v17, (v1 % (((v1 | v1)) | 1)))
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(v2, v1, v1)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}

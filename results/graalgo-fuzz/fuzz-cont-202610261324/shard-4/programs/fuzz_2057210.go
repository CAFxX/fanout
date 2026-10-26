package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int8 = -128
    var v2 int8 = (v1 / (((^(v1))) | 1))
    var v3 uint64 = 4294967295
    var v4 float64 = (-2.25 + -2.25)
    for fi5 := 0; fi5 < 6; fi5++ {
        switch "\n\t" {
        case "":
            switch "a" {
            case "":
                v6 := GenS1{F0: 0, F1: 2, F2: 4294967295}
                _ = v6
            case "hello":
                if (((1e10 / 1.5) / (-2.25 + 0.0)) > 0.0) {
                    var v7 int16 = -32768
                    v8 := []int{(fi5 % (((fi5 - (fi5 % ((fi5) | 1)))) | 1)), 1, fi5, fi5}
                    fmt.Println(v3)
                    _ = v7
                    _ = v8
                }
            case "foo bar":
                var v9 int = -867
                _ = v9
            }
        case "foo bar":
            switch "foo bar" {
            case "hello":
                if true {
                    v1 -= int8(v3)
                } else {
                    if false {
                        v10 := []int{(fi5 << ((18446744073709551615) & 63)), fi5, fi5}
                        _ = len(v10)
                        _ = len(v10)
                        _ = v10
                    } else {
                        var v11 float32 = 3.14159
                        if (!true) {
                            v12 := []int{}
                            v12 = append(v12, (fi5 - (fi5 % ((fi5) | 1))))
                            _ = v12
                        }
                        v1 = 1
                        _ = v11
                    }
                }
            }
        case "\n\t":
            var v13 int8 = (v1 + v2)
            _ = v13
        }
        _ = fi5
    }
    v14 := GenS1{F0: 1, F1: 0, F2: v3}
    var v15 int32 = 1
    for fi16 := 0; fi16 < 15; fi16++ {
        if false {
            var v17 uint32 = 2
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 int = (fi16 & fi16)
                v15 -= (v15 + (v15 % (((-(v15))) | 1)))
                v19 -= fi16
                if (2 >= 0) {
                    var v20 uint32 = (^(v17))
                    var v21 uint16 = 1
                    v17 += 2147483647
                    _ = v20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            v22 := GenS2{F0: -2147483648, F1: 65535}
            _ = v17
            _ = v22
        }
        _ = fi16
    }
    v1 = v1
    var v23 int = 2147483647
    v24 := GenS2{F0: v15, F1: 65535}
    v25 := GenS3{F0: v2}
    if (!false) {
        v4 += 0.1
        switch (v23 % ((42) | 1)) {
        case -569:
            v26 := []int{v23, 128, v23}
            _ = v26
        case -1:
            v3 -= v3
        }
    }
    var v27 uint64 = (^(v3))
    for fi28 := 0; fi28 < 12; fi28++ {
        if true {
            v29 := []int{v23, (fi28 + fi28), fi28, v23}
            _ = v29
        }
        switch ((("\n\t" + "") + "\n\t") + "") {
        case "\n\t":
            fmt.Println(v1, v3)
        default:
            v30 := []int{}
            _ = v30
        }
        var v31 float32 = (((1e10 * (-2.25 * 1.5)) * -2.25) / 1e10)
        _ = fi28
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}

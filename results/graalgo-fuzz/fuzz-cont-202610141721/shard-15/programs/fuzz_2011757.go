package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int64, p2 int32) int8 {
    p1 -= (p1 | (p1 << ((18446744073709551615) & 63)))
    var v5 float32 = ((0.1 * 0.1) / 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 100
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 bool = (true && false)
    v4 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(fi6, v3)
        _ = fi6
    }
    if (!(65535 == 993)) {
        var v7 int = -1
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 float32 = -2.25
            _ = fi8
            _ = v9
        }
        fmt.Println(v3)
        _ = v7
    } else {
        v3 = true
    }
    switch "hello" {
    case "hello":
        var v10 int64 = 9223372036854775807
        _ = v10
    case "\n\t":
        switch -407 {
        case -9223372036854775808:
            v3 = (1 < 2147483648)
        }
    case "a":
        if true {
            var v11 int = -9223372036854775808
            fmt.Println(v3, v3)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                v3 = ("hello" != "foo bar")
                if false {
                    fmt.Println(fi12)
                    v3 = false
                    switch (fi12 << ((0) & 63)) {
                    case -861:
                        fmt.Println(fi12, v3)
                    }
                }
                _ = fi12
            }
            v13 := GenS1{F0: 42}
            v3 = (!(255 > 255))
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int = fi14
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 838}
    switch "foo bar" {
    case "":
        if true {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 bool = false
                v3 = (((0 >= 42) || false) || (false && false))
                _ = fi17
                _ = v18
            }
            var v19 int = 4294967295
            _ = v19
        }
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        if (-537 != 2147483647) {
            v20 := GenS1{F0: 52}
            _ = v20
        } else {
            fmt.Println(v3)
        }
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS1{F0: 1}
        v23 := []int{}
        v24 := GenS2{F0: 42, F1: 1, F2: 1}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v3 = v3
        v3 = (-1 <= -32768)
        v26 := []int{fi25}
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := GenS3{F0: 0}
        _ = fi27
        _ = v28
    }
    v3 = ((!(0.1 == (1e10 + (0.1 * (0.1 * (0.1 / 1.0)))))) && (128 < 0))
    v3 = (!v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}

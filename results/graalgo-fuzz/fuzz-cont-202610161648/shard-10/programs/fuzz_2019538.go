package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16) uint32 {
    v7 := GenS1{F0: "foo bar"}
    v8 := []int{}
    var v9 uint = 2
    p0 -= (^(p0))
    fmt.Println(p0, v9)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 77
}
func genFn1(p0 int32) uint16 {
    v10 := GenS2{F0: "foo bar", F1: true, F2: -1}
    v11 := []int{}
    if (!true) {
        v12 := GenS2{F0: "a", F1: (4294967295 > 18446744073709551615), F2: (-(p0))}
        if ("a" < ("" + ("\n\t" + "\n\t"))) {
            var v13 float32 = 3.14159
            _ = v13
        } else {
            p0 = -1
            if len(v11) > 0 {
                v11[0] = 1
            }
            fmt.Println(p0)
        }
        _ = v12
    }
    _ = len(v11)
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 777
    var v3 int = int(v1)
    var v4 bool = true
    var v5 float32 = 3.14159
    v6 := []string{"", "\n\t", "foo bar", "hello"}
    for fi14 := 0; fi14 < 19; fi14++ {
        v3 += fi14
        switch "foo bar" {
        case "\n\t":
            v2 = v2
        case "foo bar":
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint32 = 0
                var v17 float32 = v5
                _ = fi15
                _ = v16
                _ = v17
            }
        case "a":
            switch v3 {
            case 128:
                v18 := GenS1{F0: "hello"}
                _ = v18
            case 127:
                v3 += fi14
            }
        default:
            v5 -= v5
        }
        v4 = (2 != 0)
        _ = fi14
    }
    v5 += v5
    if (true || (355 > -99)) {
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS2{F0: "a", F1: v4, F2: v1}
            if v4 {
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 int8 = 127
                    _ = fi21
                    _ = v22
                }
                v6 = append(v6, "a")
            }
            _ = fi19
            _ = v20
        }
        var v23 int = (v3 % ((v3) | 1))
        v24 := GenS3{F0: v3, F1: 2147483647}
        _ = v23
        _ = v24
    }
    if ((genFn0(2) != genFn0(10)) || v4) {
        for fi25 := 0; fi25 < 10; fi25++ {
            switch "" {
            case "\n\t":
                if true {
                    if v4 {
                        fmt.Println(v1, v1, v4)
                    }
                }
            case "foo bar":
                fmt.Println(v4, fi25, v3)
            case "hello":
                var v26 int64 = -590
                _ = v26
            }
            _ = fi25
        }
    }
    switch (v3 << ((255) & 63)) {
    case 926:
        if v4 {
            for fi27 := 0; fi27 < 4; fi27++ {
                _ = len(v6)
                v2 -= (v2 % (((^(v2))) | 1))
                if (v4 || v4) {
                    _ = len(v6)
                }
                _ = fi27
            }
        }
    case -9223372036854775808:
        fmt.Println(v2, v4, v1)
    case 32767:
        _ = len(v6)
    default:
        var v28 int64 = 42
        _ = v28
    }
    v29 := GenS4{F0: v2}
    var v30 uint16 = 255
    fmt.Println(v30, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

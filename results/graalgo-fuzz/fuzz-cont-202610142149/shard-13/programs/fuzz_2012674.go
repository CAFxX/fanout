package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
func genFn0(p0 int64) int {
    v7 := GenS1{F0: 390, F1: ("\n\t" + ("\n\t" + "hello"))}
    p0 += (p0 & 42)
    switch "a" {
    case "":
        var v8 float64 = ((float64(float32(1e10)) - -2.25) + 1e10)
        _ = v8
    case "a":
        v9 := GenS1{F0: 309, F1: "foo bar"}
        _ = v9
    default:
        for fi10 := 0; fi10 < 18; fi10++ {
            if true {
                p0 = (p0 ^ p0)
            } else {
                if true {
                    p0 = (p0 * p0)
                    fmt.Println(fi10, p0)
                }
            }
            var v11 int8 = -128
            if false {
                var v12 int = fi10
                var v13 float32 = 3.14159
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    p0 -= (p0 % (((p0 % ((p0) | 1))) | 1))
    p0 = (p0 >> ((985) & 63))
    _ = p0
    _ = v7
    return -670
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 805
    v3 := []string{"a", "foo bar"}
    v4 := make(map[int]int)
    var v5 uint = v2
    v6 := []int{-2147483648}
    for fi14 := 0; fi14 < 6; fi14++ {
        _ = len(v6)
        v2 = (v5 * 4294967295)
        v15 := GenS2{F0: 65535}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v5 += (^(v5))
        switch ("\n\t" + ("" + "foo bar")) {
        case "hello":
            var v17 int16 = 0
            _ = v17
        case "foo bar":
            switch 32767 {
            case 812:
                var v18 int32 = -919
                _ = v18
            case 437:
                v19 := GenS3{F0: 2, F1: (-(fi16))}
                _ = v19
            }
        }
        var v20 bool = (0 >= 42)
        _ = fi16
        _ = v20
    }
    switch (("" + ("hello" + "foo bar")) + ("" + "hello")) {
    case "\n\t":
        v4[1] = 32768
    }
    var v21 int16 = 978
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v2, v5)
        fmt.Println(fi22, fi22, v2)
        switch genFn0(-1) {
        case 32767:
            var v23 float64 = 0.1
            _ = v23
        default:
            fmt.Println(fi22)
        }
        _ = fi22
    }
    v2 += (v5 << ((127) & 63))
    v24 := GenS1{F0: 42, F1: "a"}
    fmt.Println(v2, v2)
    var v25 uint = 530
    v26 := GenS3{F0: 18446744073709551615, F1: 32768}
    var v27 uint = v25
    if (!false) {
        var v28 bool = false
        if (2147483647 == 65535) {
            if ((1 < 1) && v28) {
                switch genFn0(1) {
                case 4294967295:
                    var v29 bool = (!(-128 < 1))
                    _ = v29
                case 0:
                    v30 := GenS3{F0: 1, F1: 2147483647}
                    _ = v30
                case -226:
                    var v31 int8 = -128
                    _ = v31
                }
                for fi32 := 0; fi32 < 9; fi32++ {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        var v34 bool = ((v21 / (((^(v21))) | 1)) < (v21 ^ (^(v21))))
                        v25 = (v25 - 4294967295)
                        fmt.Println(v27, v28)
                        v35 := GenS4{F0: 27}
                        v36 := GenS1{F0: v21, F1: "hello"}
                        v21 = v21
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    if v28 {
                        fmt.Println(v21)
                        delete(v1, 2)
                        v37 := GenS5{F0: 0, F1: -1}
                        _ = v37
                    }
                    _ = fi32
                }
            }
        }
        if (("foo bar" + "hello") < "") {
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 string = "a"
                _ = fi38
                _ = v39
            }
        } else {
            v27 += (v27 * 1)
            if true {
                delete(v4, 9223372036854775807)
            }
        }
        _ = v28
    }
    var v40 uint16 = 879
    fmt.Println(v40, v2, v5)
    v41 := v4[-1]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}

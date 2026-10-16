package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int8; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int; F2 string }

func main() {
    var v1 int8 = 42
    var v2 uint64 = 164
    var v3 uint64 = v2
    var v4 uint32 = 42
    v5 := make(map[uint32]int)
    if true {
        var v6 uint = 164
        fmt.Println(v3, v6, v6)
        _ = v6
    }
    var v7 int64 = 754
    v3 = v2
    switch "\n\t" {
    case "foo bar":
        var v8 bool = ((1e10 / (0.0 + (1.5 - 0.0))) < ((-2.25 - 0.0) - -2.25))
        _ = v8
    case "\n\t":
        v3 += 694
    case "hello":
        fmt.Println(v7, v7)
    }
    switch "hello" {
    case "hello":
        v4 += (^(v4))
    case "":
        fmt.Println(v3, v2, v3)
    case "a":
        v9 := []int{128}
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        delete(v5, (^(v4)))
        var v11 int16 = -291
        _ = fi10
        _ = v11
    }
    v2 -= (v2 / (((v3 + (v3 << ((664) & 63)))) | 1))
    fmt.Println(v7, v1)
    if (false && (42 > int16(v2))) {
        for fi12 := 0; fi12 < 16; fi12++ {
            switch 190 {
            case -1:
                fmt.Println(fi12, v1, v4)
            case 138:
                v13 := []int{fi12, (^(fi12)), fi12, fi12}
                _ = v13
            }
            v14 := GenS1{F0: (("hello" + "") + "")}
            if false {
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch fi12 {
                    case 521:
                        v3 -= 818
                    }
                    _ = fi15
                }
                v5[2] = 699
            }
            _ = fi12
            _ = v14
        }
    } else {
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := []int{(^(fi16)), (fi16 % ((fi16) | 1)), (fi16 << ((18446744073709551615) & 63))}
            v1 = 1
            var v18 uint64 = v2
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := []int{128, -861, -405, 2147483647}
        _ = v19
    }
    var v20 uint32 = (v4 + v4)
    switch "a" {
    case "foo bar":
        v2 += (v3 >> ((112) & 63))
    case "a":
        if false {
            var v21 uint8 = 127
            _ = v21
        }
    default:
        v4 = (v4 - (v4 - v20))
    }
    fmt.Println(v20)
    var v22 uint32 = v20
    switch -9223372036854775808 {
    case 9223372036854775807:
        v22 -= v4
    case 4294967295:
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS1{F0: "\n\t"}
                if true {
                    if true {
                        v26 := GenS2{F0: 2147483647, F1: (v1 - v1), F2: 0}
                        _ = v26
                    }
                }
                _ = fi24
                _ = v25
            }
            if true {
                v20 += (v4 + 42)
                for fi27 := 0; fi27 < 16; fi27++ {
                    v28 := GenS3{F0: 18446744073709551615}
                    v29 := []int{}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v30 := []int{(fi23 * (fi23 ^ 4294967295)), (^(fi23)), (fi23 % ((32768) | 1))}
                _ = v30
            } else {
                var v31 string = "foo bar"
                v31 = "\n\t"
                v32 := []int{(fi23 % ((fi23) | 1))}
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 4; fi33++ {
                _ = len(v5)
                switch (("hello" + "a") + ("foo bar" + "")) {
                case "foo bar":
                    v34 := GenS1{F0: "foo bar"}
                    _ = v34
                case "a":
                    switch -9223372036854775808 {
                    case -1:
                        v4 = (v20 ^ (v22 * v20))
                    }
                default:
                    if (!((false || (false || true)) && true)) {
                        v35 := GenS2{F0: -1, F1: (-(v1)), F2: v1}
                        _ = v35
                    }
                }
                _ = fi33
            }
            _ = fi23
        }
    case 746:
        switch "" {
        case "foo bar":
            switch -299 {
            case 9223372036854775807:
                for fi36 := 0; fi36 < 1; fi36++ {
                    v20 -= 2147483647
                    v4 += uint32(v3)
                    _ = fi36
                }
            case 2147483647:
                if (653 > 1000) {
                    v37 := GenS3{F0: 42}
                    fmt.Println(v3, v22)
                    v38 := GenS3{F0: 1}
                    _ = v37
                    _ = v38
                }
            case -9223372036854775808:
                fmt.Println(v4, v20)
            }
        case "":
            v39 := v5[2147483648]
            _ = v39
        }
    default:
        var v40 float64 = 3.14159
        _ = v40
    }
    var v41 uint64 = 700
    v7 += v7
    for fi42 := 0; fi42 < 13; fi42++ {
        switch "a" {
        case "\n\t":
            v43 := GenS4{F0: 0, F1: fi42, F2: "a"}
            _ = v43
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v7
}

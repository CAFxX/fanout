package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 string }
func genFn0() string {
    var v7 bool = false
    v8 := GenS1{F0: 371}
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 int) uint8 {
    if false {
        v9 := []int{(p0 >> ((18446744073709551615) & 63)), p0, 1, p0}
        _ = v9
    } else {
        p0 = p0
        var v10 float32 = (1.5 / (1.5 * (1.5 + -2.25)))
        _ = v10
    }
    var v11 uint64 = 979
    v12 := []int{p0}
    var v13 bool = (100 > 0)
    if len(v12) > 0 {
        v12[0] = 1
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = -2.25
    var v3 int = -1
    v4 := make(map[uint32]int)
    var v5 uint8 = 1
    v6 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 6; fi14++ {
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "a":
            v15 := []int{}
            _ = v15
        case "hello":
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(v2, v2, fi16)
                var v17 int8 = 42
                switch "foo bar" {
                case "hello":
                    v2 += (1e10 + v2)
                case "a":
                    v18 := GenS2{F0: -920}
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        default:
            for fi19 := 0; fi19 < 2; fi19++ {
                switch ("\n\t" + ("\n\t" + "")) {
                case "":
                    v2 -= 1.5
                default:
                    if false {
                        var v20 int8 = -1
                        _ = v20
                    }
                }
                v21 := GenS1{F0: -692}
                v3 = fi19
                _ = fi19
                _ = v21
            }
        }
        _ = fi14
    }
    v3 -= v3
    var v22 bool = (-9223372036854775808 >= v3)
    v23 := GenS2{F0: 32767}
    for fi24 := 0; fi24 < 18; fi24++ {
        if true {
            for fi25 := 0; fi25 < 12; fi25++ {
                fmt.Println(v3, v2, v2)
                if (42 > 1) {
                    v22 = v22
                }
                _ = fi25
            }
            if (false || (v22 && false)) {
                v3 = 4294967295
            } else {
                var v26 int16 = -491
                if v22 {
                    var v27 uint16 = 42
                    fmt.Println(v27, v22)
                    var v28 uint64 = 0
                    _ = v27
                    _ = v28
                } else {
                    v29 := []int{v3, v3, fi24}
                    _ = v29
                }
                var v30 uint16 = 204
                _ = v26
                _ = v30
            }
        }
        v3 += 9223372036854775807
        fmt.Println(fi24)
        _ = fi24
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        if v22 {
            v32 := GenS3{F0: v22, F1: "foo bar"}
            _ = v32
        }
        _ = fi31
    }
    switch "\n\t" {
    case "":
        fmt.Println(v3, v5, v2)
    case "hello":
        switch "foo bar" {
        case "a":
            if (504 <= 0) {
                for fi33 := 0; fi33 < 11; fi33++ {
                    fmt.Println(v3, fi33, fi33)
                    _ = fi33
                }
                var v34 uint16 = 0
                _ = v34
            }
        }
    case "\n\t":
        v35 := []int{v3, v3}
        _ = v35
    default:
        switch "\n\t" {
        case "hello":
            var v36 int8 = 127
            _ = v36
        case "a":
            var v37 int16 = 1000
            _ = v37
        default:
            var v38 int64 = -1
            _ = v38
        }
    }
    _ = len(v6)
    if (18446744073709551615 == 235) {
        for fi39 := 0; fi39 < 20; fi39++ {
            v2 = (-2.25 + 0.1)
            fmt.Println(v3)
            switch 88 {
            case -709:
                var v40 bool = (42 >= 42)
                _ = v40
            case -2147483648:
                var v41 uint32 = 907
                _ = v41
            case -1:
                var v42 int8 = 100
                _ = v42
            }
            _ = fi39
        }
        var v43 uint = 1
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

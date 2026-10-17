package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 int; F2 int }
type GenS5 struct { F0 uint16; F1 bool; F2 uint32 }
func genFn0(p0 uint32, p1 int16, p2 int8) int32 {
    p0 = (^(p0))
    if (p1 == p1) {
        p2 = 50
        switch "a" {
        case "foo bar":
            v4 := GenS1{F0: false, F1: 881}
            _ = v4
        case "hello":
            p0 = (^(p0))
        }
    }
    v5 := []int{-189, -414, -839}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint = 830
        _ = len(v5)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (0 == 42), F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return -841
}

func main() {
    var v1 int = 1
    var v2 uint8 = 89
    v3 := make(map[uint32]int)
    v9 := GenS2{F0: 207, F1: 882}
    if false {
        var v10 int16 = -340
        _ = v10
    } else {
        var v11 uint = 1
        if ((!true) || false) {
            v12 := GenS1{F0: true, F1: 788}
            _ = v12
        }
        _ = v11
    }
    fmt.Println(v1)
    if (!(580 != 9223372036854775807)) {
        v1 = 127
        switch (v1 / (((v1 ^ (v1 * -143))) | 1)) {
        case -1:
            v2 = v2
        case 128:
            switch 4294967295 {
            case 377:
                v13 := GenS3{F0: 2147483647, F1: ("a" + "hello"), F2: 117}
                _ = v13
            }
        case -902:
            switch "hello" {
            case "":
                delete(v3, 65535)
            case "hello":
                for fi14 := 0; fi14 < 15; fi14++ {
                    v2 += (^(v2))
                    _ = fi14
                }
            }
        }
    } else {
        v2 -= v2
    }
    v1 -= v1
    for fi15 := 0; fi15 < 10; fi15++ {
        v1 -= (v1 - fi15)
        fmt.Println(v2, fi15, v2)
        _ = fi15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v2 += v2
        v1 -= v1
        _ = fi16
    }
    if ((v2 - v2) < (v2 ^ v2)) {
        v2 -= (v2 - (v2 >> ((42) & 63)))
    } else {
        if ((v2 >= (^(v2))) && (true && ("\n\t" != "foo bar"))) {
            v17 := GenS4{F0: (!false), F1: v1, F2: (v1 * 729)}
            for fi18 := 0; fi18 < 14; fi18++ {
                switch "a" {
                case "hello":
                    delete(v3, 42)
                case "foo bar":
                    fmt.Println(v2)
                case "a":
                    v1 = (-(v1))
                default:
                    v19 := GenS5{F0: 65535, F1: (504 >= 1), F2: 42}
                    _ = v19
                }
                v20 := []int{fi18, (v1 - (v1 * fi18))}
                var v21 uint64 = 0
                _ = fi18
                _ = v20
                _ = v21
            }
            _ = v17
        }
    }
    delete(v3, 2)
    fmt.Println(v2, v1)
    fmt.Println(v1, v1)
    for fi22 := 0; fi22 < 15; fi22++ {
        switch v1 {
        case 42:
            switch "foo bar" {
            case "a":
                v23 := v3[0]
                _ = v23
            }
        case 9223372036854775807:
            var v24 int32 = -2147483648
            _ = v24
        case -390:
            fmt.Println(v2, v1, fi22)
        }
        _ = fi22
    }
    var v25 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v9
}

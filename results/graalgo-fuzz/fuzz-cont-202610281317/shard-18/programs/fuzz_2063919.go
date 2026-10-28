package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }

func main() {
    var v1 int = 808
    var v2 string = "a"
    var v3 uint = 4294967295
    fmt.Println(v2, v3)
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4)
        var v5 uint8 = 127
        for fi6 := 0; fi6 < 15; fi6++ {
            switch "hello" {
            case "\n\t":
                v5 -= v5
            default:
                v7 := GenS1{F0: -1, F1: 699}
                _ = v7
            }
            v8 := GenS1{F0: 42, F1: fi6}
            if true {
                v9 := []int{(fi4 ^ (fi4 + (fi4 & fi6)))}
                _ = v9
            }
            _ = fi6
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    switch (^(v1)) {
    case 0:
        var v10 int8 = -112
        _ = v10
    case 128:
        v11 := GenS2{F0: 32767}
        _ = v11
    case -2147483648:
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS2{F0: 32767}
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS3{F0: -640, F1: 42, F2: 255}
                v16 := []int{(fi12 >> ((v3) & 63)), 9223372036854775807}
                v2 = v2
                _ = fi14
                _ = v15
                _ = v16
            }
            switch v1 {
            case 32768:
                v17 := []int{(v1 >> ((1) & 63)), 32767, (fi12 % (((v1 & (^(v1)))) | 1))}
                _ = v17
            case -1:
                fmt.Println(v2, v3)
            default:
                var v18 int8 = 42
                _ = v18
            }
            _ = fi12
            _ = v13
        }
    default:
        v19 := []int{(v1 - v1)}
        _ = v19
    }
    switch "a" {
    case "hello":
        switch -600 {
        case 1:
            v20 := GenS3{F0: (v1 ^ v1), F1: 255, F2: 1}
            _ = v20
        default:
            var v21 bool = false
            _ = v21
        }
    default:
        if (!false) {
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 bool = (!false)
                v24 := GenS1{F0: 32767, F1: 1}
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: -32768, F1: 9223372036854775807}
            _ = v25
        }
    }
    v26 := GenS4{F0: 18446744073709551615}
    if (!false) {
        switch "foo bar" {
        case "foo bar":
            var v27 float32 = (1e10 - 1.5)
            _ = v27
        case "":
            v28 := GenS5{F0: 0}
            _ = v28
        default:
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS6{F0: -128}
                v31 := GenS7{F0: ((!false) || (-2147483648 > -981))}
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        v32 := []int{(v1 | v1), (v1 ^ 4294967295)}
        var v33 int8 = 1
        _ = v32
        _ = v33
    }
    v3 -= v3
    fmt.Println(v3)
    v2 = v2
    var v34 int16 = 1
    fmt.Println(v34)
    v35 := []int{(v1 | (v1 / ((v1) | 1)))}
    for fi36 := 0; fi36 < 13; fi36++ {
        v2 = v2
        _ = fi36
    }
    v37 := GenS3{F0: v1, F1: 65535, F2: uint8(v3)}
    var v38 uint64 = 0
    var v39 string = v2
    v40 := GenS2{F0: 1}
    v41 := GenS7{F0: ((!(42 >= 100)) && (986 >= 503))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}

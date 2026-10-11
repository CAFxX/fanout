package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 int = 128
    v3 := make(map[uint32]bool)
    var v4 uint64 = 1
    v5 := []int{(v2 % ((v2) | 1)), v2, int(v4)}
    var v6 int16 = 289
    v7 := GenS1{F0: (v4 | v4), F1: 2}
    if false {
        if true {
            if (255 < 127) {
                v4 = 904
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int8 = 100
                var v10 int64 = -1
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v4 = 0
            v11 := GenS1{F0: (v4 % ((v4) | 1)), F1: (v4 / ((v4) | 1))}
            _ = v11
        }
        v12 := GenS1{F0: v4, F1: v4}
        _ = v12
    } else {
        v4 = v4
        switch "hello" {
        case "a":
            delete(v1, (v2 * (-(v2))))
        case "hello":
            switch v2 {
            case 128:
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 uint16 = 1
                    var v15 int = 128
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case 127:
                switch "\n\t" {
                case "\n\t":
                    v16 := GenS1{F0: 4294967295, F1: (v4 >> ((2) & 63))}
                    _ = v16
                }
            case 772:
                var v17 uint32 = 2147483648
                _ = v17
            }
        }
        if (v2 <= 1) {
            v3[2147483648] = (765 > 65535)
        }
    }
    if (true && true) {
        var v18 string = "foo bar"
        _ = v18
    } else {
        v6 = 718
        v4 -= (^(v4))
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        if (((v2 << ((18446744073709551615) & 63)) != fi19) || (149 != 2147483648)) {
            v4 = v4
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                if (true && (1 >= 1)) {
                    var v21 int8 = 42
                    v22 := GenS2{F0: 42}
                    _ = v21
                    _ = v22
                }
                if ("foo bar" < ("" + "a")) {
                    v23 := GenS3{F0: v6, F1: "foo bar"}
                    v2 = (v2 - fi20)
                    _ = v23
                }
                _ = fi20
            }
        }
        v6 -= 0
        _ = fi19
    }
    if false {
        if ("\n\t" > "foo bar") {
            if (211 < 255) {
                for fi24 := 0; fi24 < 18; fi24++ {
                    fmt.Println(v4, v2)
                    v2 -= (fi24 | (fi24 - fi24))
                    _ = fi24
                }
                v25 := GenS3{F0: (^(v6)), F1: "\n\t"}
                v26 := GenS4{F0: 1}
                _ = v25
                _ = v26
            }
            var v27 int32 = 2147483647
            v5 = append(v5, v2)
            _ = v27
        }
        v2 += v2
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint8 = 170
        var v30 uint = 81
        var v31 string = "a"
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    v2 -= v2
    v4 = (v4 / ((v4) | 1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = (v4 | (v4 | (^(v4))))
    switch (((("foo bar" + "hello") + "") + "\n\t") + "foo bar") {
    case "a":
        var v32 float32 = -2.25
        _ = v32
    }
    v33 := GenS1{F0: 1, F1: (v4 << ((42) & 63))}
    switch "foo bar" {
    case "":
        if ((v6 + (^(v6))) < v6) {
            var v34 uint16 = 621
            _ = v34
        } else {
            fmt.Println(v2, v6, v6)
        }
    case "\n\t":
        v35 := GenS5{F0: 0, F1: -222}
        _ = v35
    default:
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 uint16 = 42
            fmt.Println(v37)
            var v38 uint16 = 32767
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

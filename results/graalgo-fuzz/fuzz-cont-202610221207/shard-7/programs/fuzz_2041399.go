package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    v3 := []int{-1, 127}
    var v4 float64 = 0.1
    for fi5 := 0; fi5 < 7; fi5++ {
        _ = len(v3)
        switch -380 {
        case -9223372036854775808:
            _ = len(v3)
        default:
            var v6 float64 = 3.14159
            _ = v6
        }
        var v7 int32 = 2147483647
        _ = fi5
        _ = v7
    }
    if (!(false && true)) {
        for fi8 := 0; fi8 < 4; fi8++ {
            v2 = ("a" + v2)
            v9 := GenS1{F0: 818, F1: -931, F2: 859}
            var v10 uint8 = 127
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(v2)
    }
    var v11 int8 = 127
    for fi12 := 0; fi12 < 1; fi12++ {
        switch fi12 {
        case 1:
            var v13 int64 = 616
            _ = v13
        case -1:
            v4 += float64(float32(-2.25))
        case 4294967295:
            switch fi12 {
            case 697:
                switch v2 {
                case "foo bar":
                    v11 = v11
                }
            case -2147483648:
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 string = (v2 + v2)
                    var v16 int64 = 9223372036854775807
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case -604:
                fmt.Println(fi12, fi12)
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            v2 = "foo bar"
            var v18 int32 = -1
            _ = fi17
            _ = v18
        }
        _ = fi12
    }
    v11 = (-(v11))
    _ = len(v3)
    v11 += (v11 % (((v11 | (v11 / ((v11) | 1)))) | 1))
    switch v2 {
    case "":
        v19 := GenS1{F0: 0, F1: -1, F2: 4294967295}
        _ = v19
    case "foo bar":
        var v20 uint = 554
        _ = v20
    default:
        v21 := GenS1{F0: 810, F1: -1, F2: 18446744073709551615}
        _ = v21
    }
    if (1e10 <= 1.5) {
        if (true || false) {
            v2 = ("\n\t" + v2)
            var v22 int16 = 1000
            for fi23 := 0; fi23 < 3; fi23++ {
                if true {
                    v24 := GenS1{F0: 42, F1: 402, F2: 494}
                    var v25 int64 = 0
                    _ = v24
                    _ = v25
                }
                _ = fi23
            }
            _ = v22
        }
        v4 = (float64(0.0) + 0.1)
        if false {
            if false {
                fmt.Println(v4, v2)
                v26 := GenS1{F0: 1, F1: 42, F2: 312}
                var v27 uint32 = 1
                _ = v26
                _ = v27
            }
        } else {
            var v28 uint64 = 402
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 12; fi29++ {
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 bool = (42 > 0)
                if v31 {
                    fmt.Println(v2, fi30, v4)
                }
                switch "\n\t" {
                case "hello":
                    v32 := GenS1{F0: 42, F1: 9223372036854775807, F2: 280}
                    _ = v32
                case "\n\t":
                    var v33 uint16 = 2
                    _ = v33
                }
                var v34 int16 = 1
                _ = fi30
                _ = v31
                _ = v34
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi29
        }
    }
    v2 = v2
    var v35 int8 = v11
    v36 := GenS1{F0: 2, F1: 1, F2: 445}
    v35 = v11
    if true {
        v35 = v35
        v1[1] = -235
        delete(v1, 42)
    }
    v35 = (v35 + (-(v11)))
    _ = len(v1)
    var v37 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}

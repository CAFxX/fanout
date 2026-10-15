package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0() uint64 {
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            switch fi7 {
            case 127:
                var v8 float64 = (1.5 - (-2.25 * float64(3.14159)))
                _ = v8
            case -944:
                var v9 int = 4294967295
                _ = v9
            case 982:
                v10 := GenS1{F0: 493, F1: -1}
                _ = v10
            }
            _ = fi7
        }
        var v11 int8 = 1
        _ = v11
    }
    switch "\n\t" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            var v12 uint8 = 139
            _ = v12
        case "\n\t":
            v13 := GenS2{F0: 224, F1: "hello", F2: 964}
            _ = v13
        default:
            var v14 bool = true
            _ = v14
        }
    case "\n\t":
        v15 := []int{}
        _ = v15
    case "foo bar":
        var v16 int8 = 100
        _ = v16
    }
    v17 := []int{-874, -1, 32768}
    var v18 float32 = 3.14159
    v18 = v18
    _ = v17
    _ = v18
    return 4294967295
}

func main() {
    var v1 int8 = 127
    var v2 uint = 120
    var v3 uint64 = 611
    var v4 int = -2147483648
    var v5 int = (-(v4))
    var v6 uint = (v2 >> ((255) & 63))
    v19 := []int{(v5 | (v5 & -91)), v4, (v4 << ((0) & 63)), v4}
    v20 := GenS3{F0: 0, F1: 0}
    v1 = v1
    v2 = 761
    if false {
        if true {
            v21 := GenS2{F0: (v2 % ((v6) | 1)), F1: ("" + "hello"), F2: (v2 % ((v6) | 1))}
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := GenS2{F0: 0, F1: "a", F2: v2}
                v3 += 42
                fmt.Println(v5)
                _ = fi22
                _ = v23
            }
            v2 = 2
            _ = v21
        } else {
            v2 = 967
            if false {
                var v24 bool = (!((v1 | v1) < v1))
                _ = v24
            }
            switch (v4 >> ((v6) & 63)) {
            case 32767:
                var v25 int64 = 1
                _ = v25
            case 9223372036854775807:
                var v26 uint32 = 215
                _ = v26
            case 872:
                var v27 int16 = -286
                _ = v27
            }
        }
    } else {
        if (false && ((v1 >> ((v2) & 63)) <= (v1 & 68))) {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            var v28 int = (-(v4))
            var v29 float32 = ((((float32(0.0) * 0.1) + 1.5) + 0.0) / -2.25)
            _ = v28
            _ = v29
        }
        if (true && (0.1 >= 1.5)) {
            v19 = append(v19, 128)
            for fi30 := 0; fi30 < 12; fi30++ {
                v19 = append(v19, fi30)
                if len(v19) > 0 {
                    v19[0] = fi30
                }
                _ = fi30
            }
        } else {
            var v31 bool = true
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            var v32 int8 = (v1 % ((38) | 1))
            _ = v31
            _ = v32
        }
    }
    v3 = (v3 << (((v2 / (((^(v6))) | 1))) & 63))
    v33 := GenS3{F0: 0, F1: 2147483647}
    var v34 uint32 = uint32(v3)
    if true {
        var v35 int64 = 9223372036854775807
        _ = v35
    }
    var v36 string = "foo bar"
    for fi37 := 0; fi37 < 15; fi37++ {
        v38 := GenS3{F0: 429, F1: 42}
        v6 += (v6 / ((v6) | 1))
        _ = fi37
        _ = v38
    }
    v4 -= int(v34)
    v39 := GenS1{F0: 65535, F1: (v1 * (v1 << (((v2 & (v6 % (((v2 + v6)) | 1)))) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
type GenS4 struct { F0 int8; F1 bool; F2 int32 }
type GenS5 struct { F0 int8 }
func genFn0() uint64 {
    var v4 string = ((("" + "foo bar") + "a") + (("foo bar" + "hello") + "foo bar"))
    for fi5 := 0; fi5 < 9; fi5++ {
        if false {
            var v6 float32 = 3.14159
            v6 += v6
            _ = v6
        } else {
            for fi7 := 0; fi7 < 6; fi7++ {
                switch ((v4 + "hello") + v4) {
                case "\n\t":
                    switch (fi7 / ((fi7) | 1)) {
                    case -863:
                        var v8 int32 = -551
                        _ = v8
                    case 2147483647:
                        v4 = "a"
                    default:
                        var v9 int32 = 353
                        _ = v9
                    }
                case "":
                    switch (v4 + (v4 + v4)) {
                    case "":
                        var v10 float64 = float64(float32(0.0))
                        _ = v10
                    }
                case "a":
                    if (true && (true && (127 > 0))) {
                        v11 := []int{611, -2147483648}
                        switch ((("" + (v4 + "foo bar")) + "\n\t") + v4) {
                        case "a":
                            var v12 float64 = ((-2.25 - 1e10) + 0.1)
                            _ = v12
                        default:
                            fmt.Println(v4, fi7)
                        }
                        fmt.Println(fi5)
                        _ = v11
                    } else {
                        v13 := []int{(fi7 ^ fi5), (fi5 * 128)}
                        switch v4 {
                        case "hello":
                            var v14 uint32 = 42
                            _ = v14
                        }
                        var v15 uint = 1
                        _ = v13
                        _ = v15
                    }
                }
                v4 = "a"
                if false {
                    var v16 uint32 = 2147483648
                    _ = v16
                }
                var v17 int32 = 582
                _ = fi7
                _ = v17
            }
            var v18 string = v4
            _ = v18
        }
        v4 = v4
        _ = fi5
    }
    fmt.Println(v4)
    v4 = v4
    switch 32767 {
    case 745:
        switch ((v4 + v4) + ("\n\t" + v4)) {
        case "a":
            v4 = v4
        case "":
            v19 := []int{84}
            _ = v19
        case "hello":
            var v20 int8 = 100
            _ = v20
        }
    case 4294967295:
        switch 128 {
        case -9223372036854775808:
            var v21 string = v4
            _ = v21
        }
    case 128:
        var v22 int64 = 749
        _ = v22
    }
    _ = v4
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 float32 = 0.1
    v3 := make(map[int]bool)
    v23 := GenS1{F0: 674, F1: 458, F2: 48}
    for fi24 := 0; fi24 < 7; fi24++ {
        v1 = (v1 - 2)
        var v25 int16 = 564
        _ = fi24
        _ = v25
    }
    v1 = 32767
    v2 = float32((((-2.25 / float64(v2)) + (0.1 / 1.0)) + 0.1))
    for fi26 := 0; fi26 < 11; fi26++ {
        if false {
            v1 = (v1 * (v1 * v1))
            var v27 float64 = 3.14159
            _ = v27
        } else {
            v28 := []int{(fi26 & -296), 9223372036854775807}
            _ = v28
        }
        _ = fi26
    }
    v29 := GenS2{F0: 18446744073709551615, F1: 100}
    switch 42 {
    case 616:
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 int8 = -1
            switch "hello" {
            case "foo bar":
                var v32 int32 = 42
                _ = v32
            case "a":
                for fi33 := 0; fi33 < 13; fi33++ {
                    v2 = 3.14159
                    fmt.Println(fi30)
                    _ = fi33
                }
            default:
                for fi34 := 0; fi34 < 2; fi34++ {
                    v2 -= 0.0
                    _ = fi34
                }
            }
            _ = fi30
            _ = v31
        }
    case 128:
        v35 := GenS3{F0: "\n\t", F1: 2, F2: -450}
        _ = v35
    }
    switch 321 {
    case 32768:
        fmt.Println(v2, v2)
    case 42:
        v36 := []int{-9223372036854775808}
        _ = v36
    case -2147483648:
        v2 = ((v2 + float32((((0.0 / float64(3.14159)) / 1.0) * float64((v2 * v2))))) + ((((((v2 * ((v2 / (v2 - float32((0.0 * 3.14159)))) / float32(3.14159))) / v2) - (v2 - v2)) * v2) - v2) - v2))
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        for fi38 := 0; fi38 < 19; fi38++ {
            v39 := GenS4{F0: 1, F1: (!(0 == 377)), F2: 0}
            var v40 float64 = -2.25
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi37
    }
    v2 -= v2
    if true {
        v2 -= 0.1
        v2 += (-2.25 + 1.5)
    } else {
        v41 := GenS5{F0: -65}
        _ = v41
    }
    var v42 float64 = 3.14159
    var v43 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v42
    _ = v43
}

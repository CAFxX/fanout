package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    var v1 int32 = -2147483648
    var v2 string = (((((((("foo bar" + "foo bar") + "\n\t") + ("hello" + ("" + ("a" + "hello")))) + "hello") + "a") + "a") + "foo bar") + ("\n\t" + ""))
    var v3 int = 42
    v4 := GenS1{F0: 127}
    var v5 float64 = -2.25
    switch (((v2 + "hello") + "hello") + ("" + ("\n\t" + (v2 + (v2 + v2))))) {
    case "hello":
        fmt.Println(v5)
    case "":
        v5 = 0.1
    default:
        fmt.Println(v1)
    }
    var v6 uint32 = 228
    if (!(false || ((v2 + v2) >= "a"))) {
        var v7 int32 = v1
        if (1 != 0) {
            v8 := GenS1{F0: 0}
            _ = v8
        } else {
            v5 -= ((3.14159 / (1e10 - ((((0.0 * (float64(3.14159) + v5)) / v5) - (-2.25 / v5)) * v5))) / v5)
            if false {
                fmt.Println(v1, v1, v1)
                if (-2.25 >= v5) {
                    v3 = (v3 * (v3 / (((v3 + 9223372036854775807)) | 1)))
                }
            } else {
                for fi9 := 0; fi9 < 16; fi9++ {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v11 := GenS1{F0: 84}
                        var v12 float32 = (0.1 / 0.1)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                var v13 bool = (382 <= 325)
                _ = v13
            }
        }
        v14 := GenS1{F0: 1}
        _ = v14
        _ = v7
    } else {
        v1 = (v1 * (v1 & (v1 & -2147483648)))
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS1{F0: 143}
            _ = fi15
            _ = v16
        }
    }
    if ((^(v6)) < v6) {
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int16 = 883
            _ = fi17
            _ = v18
        }
        v5 = (v5 * 1e10)
        for fi19 := 0; fi19 < 13; fi19++ {
            switch v2 {
            case "hello":
                v20 := []int{(fi19 >> ((65535) & 63))}
                _ = v20
            default:
                v21 := []int{}
                _ = v21
            }
            v5 += v5
            _ = fi19
        }
    }
    if false {
        var v22 uint64 = 2
        _ = v22
    }
    v23 := []int{42}
    var v24 uint16 = 776
    switch ("" + v2) {
    case "\n\t":
        var v25 uint8 = 255
        _ = v25
    default:
        var v26 int64 = 1
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        v23 = append(v23, 353)
        if true {
            if (-2147483648 >= 2147483647) {
                switch (^(v3)) {
                case 4294967295:
                    v3 = 32768
                default:
                    var v28 float32 = 1e10
                    _ = v28
                }
            }
            if len(v23) > 0 {
                v23[0] = 108
            }
            var v29 int16 = -1
            _ = v29
        } else {
            for fi30 := 0; fi30 < 1; fi30++ {
                v31 := GenS1{F0: 42}
                fmt.Println(v6)
                if (255 >= 255) {
                    v1 = -2147483648
                    v32 := GenS2{F0: 612, F1: 42}
                    v33 := GenS1{F0: 128}
                    _ = v32
                    _ = v33
                }
                v6 = v6
                var v34 string = v2
                _ = fi30
                _ = v31
                _ = v34
            }
            v6 -= 2147483647
            var v35 string = v2
            _ = v35
        }
        _ = fi27
    }
    v36 := GenS3{F0: 1, F1: 32768}
    v5 += float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    v3 := make(map[uint32]string)
    var v4 uint32 = 1
    fmt.Println(v2, v4)
    if true {
        if (v2 > ("hello" + (v2 + v2))) {
            for fi5 := 0; fi5 < 12; fi5++ {
                var v6 bool = ((!(2 == 255)) || (!true))
                _ = fi5
                _ = v6
            }
            switch v2 {
            case "foo bar":
                fmt.Println(v4, v2)
            }
        }
    }
    switch "a" {
    case "\n\t":
        v7 := GenS1{F0: 863, F1: 39}
        _ = v7
    }
    v8 := GenS2{F0: 0}
    v4 = (^(v4))
    v9 := []int{127}
    switch v2 {
    case "a":
        switch v2 {
        case "\n\t":
            switch -2147483648 {
            case 32768:
                v10 := GenS3{F0: 0}
                _ = v10
            default:
                v2 = v2
            }
        case "a":
            v2 = "a"
        }
    case "foo bar":
        if (!true) {
            v2 = v2
            v2 = v2
            var v11 string = ""
            _ = v11
        }
    case "hello":
        var v12 string = (v2 + (v2 + "a"))
        _ = v12
    default:
        v13 := GenS2{F0: 32767}
        _ = v13
    }
    var v14 float64 = 0.1
    var v15 int16 = 32767
    v2 = (v2 + ((v2 + "") + v2))
    v2 = "\n\t"
    if false {
        v16 := GenS1{F0: 488, F1: 42}
        fmt.Println(v14, v14, v2)
        _ = v16
    } else {
        fmt.Println(v4, v15)
        v2 = "a"
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS4{F0: (fi17 + fi18)}
            v20 := GenS3{F0: 1}
            var v21 float64 = v14
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        v2 = ""
        _ = fi17
    }
    if (v15 == 0) {
        switch 457 {
        case 9223372036854775807:
            v14 += 1e10
        case 32767:
            if (!false) {
                var v22 int64 = 881
                switch 4294967295 {
                case -9223372036854775808:
                    if ((v14 + (1e10 / v14)) < 3.14159) {
                        var v23 float64 = float64(-2.25)
                        if true {
                            var v24 uint64 = 180
                            var v25 int16 = (v15 % (((v15 & v15)) | 1))
                            _ = v24
                            _ = v25
                        }
                        var v26 int16 = v15
                        _ = v23
                        _ = v26
                    }
                case 4294967295:
                    v22 -= 1
                default:
                    var v27 int64 = 9223372036854775807
                    _ = v27
                }
                v28 := GenS2{F0: 322}
                _ = v22
                _ = v28
            } else {
                switch ((v2 + (v2 + (v2 + v2))) + (v2 + v2)) {
                case "\n\t":
                    var v29 float64 = (0.0 + v14)
                    _ = v29
                }
            }
        case -419:
            v14 = (float64((0.0 + 0.0)) / (float64((((0.0 / 1.5) - -2.25) * 0.1)) - v14))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

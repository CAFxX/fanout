package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int }

func main() {
    v1 := []string{"hello"}
    var v2 uint32 = 2
    var v3 float32 = -2.25
    var v4 int32 = 517
    var v5 int = 0
    v5 = v5
    var v6 uint = 18446744073709551615
    v5 = v5
    _ = len(v1)
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 float64 = (1e10 + 1e10)
            v10 := GenS1{F0: 127, F1: 80}
            var v11 int32 = v4
            v12 := GenS1{F0: 128, F1: 674}
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
            _ = v9
        }
        switch (fi7 & v5) {
        case 398:
            v1 = append(v1, "foo bar")
        case 332:
            var v13 int8 = 42
            _ = v13
        case 0:
            v1 = append(v1, "hello")
        }
        _ = fi7
    }
    if false {
        var v14 float32 = ((float32(float64(((float32(3.14159) * v3) / v3))) / -2.25) * (0.0 + float32(1e10)))
        for fi15 := 0; fi15 < 6; fi15++ {
            switch "" {
            case "":
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS1{F0: 38, F1: 4294967295}
                    fmt.Println(v6, fi15, v3)
                    v4 -= (v4 + -1)
                    var v18 uint8 = 2
                    switch (-(fi16)) {
                    case 32768:
                        fmt.Println(v14, v4, fi15)
                    case 0:
                        var v19 float32 = (1e10 + 0.0)
                        _ = v19
                    default:
                        v20 := GenS2{F0: 0}
                        fmt.Println(fi15, fi16, v2)
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            _ = len(v1)
            _ = fi15
        }
        _ = v14
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 uint16 = 255
            _ = fi21
            _ = v22
        }
        v5 += v5
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v4 = -1
        _ = fi23
    }
    v4 += 507
    var v24 uint32 = v2
    if true {
        switch (v5 - v5) {
        case 353:
            _ = len(v1)
        case 32768:
            v1 = append(v1, "a")
        case 1:
            var v25 float64 = float64(v3)
            _ = v25
        default:
            v5 -= (v5 ^ v5)
        }
    }
    v1 = append(v1, "a")
    if (!false) {
        if (!true) {
            fmt.Println(v5, v5)
        } else {
            switch 299 {
            case -1:
                v24 += v24
            default:
                switch "" {
                case "\n\t":
                    switch (("" + "\n\t") + "a") {
                    case "a":
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v5 -= fi26
                            _ = fi26
                        }
                    case "hello":
                        _ = len(v1)
                    case "":
                        fmt.Println(v2, v3)
                    default:
                        v27 := GenS3{F0: v5, F1: v5}
                        _ = v27
                    }
                case "foo bar":
                    v1 = append(v1, "")
                case "hello":
                    v2 -= v24
                }
            }
            var v28 string = ("" + ("\n\t" + ((("a" + "hello") + "\n\t") + ("a" + "foo bar"))))
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            switch ("foo bar" + "") {
            case "hello":
                _ = len(v1)
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                switch "hello" {
                case "\n\t":
                    v31 := GenS3{F0: fi29, F1: (^(v5))}
                    _ = v31
                case "a":
                    v4 += (-(v4))
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v32 uint64 = 222
                    _ = v32
                }
                var v33 uint8 = 10
                var v34 int = v5
                _ = fi30
                _ = v33
                _ = v34
            }
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi29
        }
        v35 := GenS3{F0: 4294967295, F1: 760}
        _ = v35
    } else {
        fmt.Println(v6)
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            v24 += v2
            var v38 string = "\n\t"
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    switch "hello" {
    case "":
        var v39 string = "a"
        _ = v39
    }
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

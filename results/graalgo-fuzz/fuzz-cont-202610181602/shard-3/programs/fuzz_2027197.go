package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
type GenS3 struct { F0 string; F1 string }

func main() {
    var v1 int = -526
    var v2 float64 = 1.5
    v3 := []uint32{2147483647, 2, 1, 2147483647}
    var v4 uint32 = 2
    var v5 float64 = v2
    var v6 uint32 = 2
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(v4)
        var v8 int8 = 127
        v8 += v8
        _ = fi7
        _ = v8
    }
    v5 += 0.0
    if false {
        var v9 uint8 = 183
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 float64 = (v5 + 0.1)
            var v12 int8 = -128
            var v13 int8 = (-(v12))
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!false) {
        var v14 int8 = 100
        _ = v14
    } else {
        var v15 uint8 = 0
        switch v1 {
        case 128:
            if true {
                for fi16 := 0; fi16 < 10; fi16++ {
                    switch (("\n\t" + "a") + "hello") {
                    case "":
                        var v17 int64 = 338
                        _ = v17
                    }
                    v1 += fi16
                    v6 = v4
                    _ = fi16
                }
                v5 = v2
                v4 += v6
            }
        case 42:
            v15 = (^(v15))
        }
        _ = v15
    }
    v6 -= v6
    v4 = (v6 - v6)
    var v18 float64 = (1.5 * v2)
    var v19 uint = 2
    var v20 uint32 = 65535
    for fi21 := 0; fi21 < 5; fi21++ {
        v2 = 1.5
        v1 += (fi21 * (fi21 / ((v1) | 1)))
        _ = fi21
    }
    var v22 string = "hello"
    var v23 float32 = -2.25
    if (!false) {
        var v24 uint32 = v20
        _ = v24
    }
    switch ((("foo bar" + "a") + v22) + "a") {
    case "\n\t":
        v25 := GenS1{F0: (v1 - 127), F1: 192, F2: 187}
        _ = v25
    case "foo bar":
        switch (v1 - 185) {
        case -11:
            v4 -= v20
        case 0:
            if true {
                var v26 uint = 364
                var v27 bool = true
                switch "hello" {
                case "a":
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v22 = "foo bar"
                        v26 += (v26 % (((v26 * (^(v19)))) | 1))
                        _ = fi28
                    }
                case "\n\t":
                    fmt.Println(v5)
                }
                _ = v26
                _ = v27
            }
        case 127:
            switch (v1 ^ v1) {
            case 2147483647:
                switch v22 {
                case "hello":
                    if (!(-128 == -97)) {
                        _ = len(v3)
                    }
                case "":
                    if true {
                        fmt.Println(v6, v2)
                    }
                default:
                    v29 := GenS2{F0: 142, F1: 2, F2: ("\n\t" + v22)}
                    _ = v29
                }
            case 42:
                if (((^(v19)) < v19) || true) {
                    v30 := GenS3{F0: "a", F1: v22}
                    _ = len(v3)
                    var v31 uint16 = 42
                    _ = v30
                    _ = v31
                }
            default:
                _ = len(v3)
            }
        default:
            if false {
                if len(v3) > 0 {
                    v3[0] = (v4 >> ((v19) & 63))
                }
            }
        }
    }
    var v32 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

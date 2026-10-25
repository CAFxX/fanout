package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint = 65535
    var v2 uint16 = uint16(v1)
    v3 := []int{}
    var v4 int = 0
    var v5 int8 = 42
    var v6 int64 = -11
    if (((v1 >= (v1 - 255)) && true) && ((335 > -889) && (v6 < (v6 ^ v6)))) {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch fi7 {
            case 42:
                v1 = 2
            default:
                var v8 float32 = (3.14159 / 1.5)
                _ = v8
            }
            if true {
                for fi9 := 0; fi9 < 1; fi9++ {
                    if true {
                        v10 := GenS1{F0: 127}
                        _ = v10
                    }
                    v11 := GenS1{F0: 2}
                    if false {
                        v1 = (^(v1))
                    } else {
                        fmt.Println(v1)
                        var v12 float64 = ((float64((float32(1.5) / (1.5 * 1.5))) + ((-2.25 - -2.25) - float64(-2.25))) * (1.5 / (1e10 / 3.14159)))
                        _ = v12
                    }
                    _ = fi9
                    _ = v11
                }
                v3 = append(v3, (fi7 & (fi7 << (((v1 % ((v1) | 1))) & 63))))
            }
            var v13 int = (fi7 / (((fi7 << (((v1 * v1)) & 63))) | 1))
            _ = fi7
            _ = v13
        }
        var v14 uint8 = 127
        _ = v14
    } else {
        if true {
            switch (v4 ^ (^(v4))) {
            case 9223372036854775807:
                v15 := GenS1{F0: 2}
                _ = v15
            case 42:
                v3 = append(v3, (-(v4)))
            }
            v16 := GenS1{F0: 1}
            _ = v16
        }
    }
    if true {
        fmt.Println(v1)
        if true {
            v17 := GenS2{F0: true}
            var v18 uint8 = 255
            _ = v17
            _ = v18
        } else {
            var v19 int64 = -632
            switch "hello" {
            case "\n\t":
                fmt.Println(v6, v6, v4)
            default:
                var v20 uint32 = 91
                _ = v20
            }
            _ = v19
        }
    }
    fmt.Println(v2, v1)
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int64 = -1
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v4 = (fi23 / (((^(v4))) | 1))
        _ = fi23
    }
    switch "a" {
    case "":
        var v24 int64 = int64(v1)
        _ = v24
    case "hello":
        var v25 uint32 = 42
        _ = v25
    }
    v2 -= v2
    v5 += (v5 * (v5 ^ v5))
    v26 := GenS2{F0: true}
    if (!(-106 != v5)) {
        v5 = v5
    }
    v27 := GenS2{F0: false}
    v28 := GenS3{F0: 255}
    if false {
        var v29 int8 = -128
        var v30 float32 = -2.25
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        if true {
            switch ("a" + ("foo bar" + "hello")) {
            case "":
                v32 := GenS1{F0: 128}
                _ = v32
            case "\n\t":
                v5 = 127
            default:
                if (2 == 0) {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 int16 = -458
                        var v35 uint16 = v2
                        var v36 uint16 = v35
                        _ = len(v3)
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    switch ("a" + "a") {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = int(v2)
                        }
                    case "hello":
                        var v37 uint64 = 18446744073709551615
                        _ = v37
                    }
                }
            }
        }
        var v38 int = (fi31 | -2147483648)
        var v39 uint8 = 127
        _ = fi31
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

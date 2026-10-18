package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 string = ("\n\t" + "")
    var v2 uint = 255
    var v3 float32 = float32(-2.25)
    fmt.Println(v1)
    if (!(2 >= 42)) {
        v4 := []int{127, 9223372036854775807, 2147483647}
        _ = v4
    }
    if true {
        v1 = (v1 + v1)
        v5 := []int{128, -9223372036854775808, 42}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 11; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                switch (fi7 - (fi7 | fi7)) {
                case -236:
                    var v8 uint32 = 42
                    _ = v8
                case -597:
                    var v9 uint8 = 2
                    _ = v9
                }
                v1 = (("foo bar" + "hello") + (((v1 + (((v1 + "\n\t") + "") + v1)) + ("" + ((("\n\t" + "foo bar") + ((v1 + v1) + "\n\t")) + (("foo bar" + (v1 + v1)) + "\n\t")))) + ("" + "foo bar")))
                _ = fi7
            }
            if false {
                v10 := GenS1{F0: 2, F1: (-(fi6))}
                if ((!(fi6 <= (fi6 >> (((v2 - (v2 << ((v2) & 63)))) & 63)))) || (0 >= 215)) {
                    var v11 float64 = (float64(((float32(0.0) * 0.1) / (v3 * v3))) * 0.0)
                    _ = v11
                }
                var v12 float64 = float64(v3)
                _ = v10
                _ = v12
            }
            _ = fi6
        }
        var v13 uint16 = 1
        v14 := GenS2{F0: v1}
        _ = v13
        _ = v14
    }
    var v15 int8 = -1
    v16 := []int{2147483647}
    v17 := GenS2{F0: (v1 + "hello")}
    switch (v1 + v1) {
    case "":
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(v2, v15, v1)
            _ = fi18
        }
    }
    fmt.Println(v15)
    if (v3 < v3) {
        v19 := GenS3{F0: 165, F1: 1}
        switch "" {
        case "\n\t":
            fmt.Println(v2, v15)
        case "a":
            if true {
                var v20 int32 = -515
                _ = v20
            } else {
                for fi21 := 0; fi21 < 3; fi21++ {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 uint = 42
                        v16 = append(v16, 4294967295)
                        _ = fi22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 int16 = 1000
                        _ = fi24
                        _ = v25
                    }
                    var v26 int8 = 42
                    _ = fi21
                    _ = v26
                }
            }
        default:
            v27 := GenS2{F0: ""}
            _ = v27
        }
        _ = v19
    } else {
        if ((1 > -1) || false) {
            if ((!(226 <= 128)) || true) {
                fmt.Println(v3, v1, v1)
            }
        }
    }
    if true {
        if false {
            v1 = v1
            for fi28 := 0; fi28 < 18; fi28++ {
                var v29 float64 = (-2.25 + (1.5 / (((0.0 / 1.0) + 3.14159) + (3.14159 - 0.1))))
                _ = fi28
                _ = v29
            }
            v1 = ((v1 + (("" + v1) + ((v1 + v1) + "\n\t"))) + v1)
        } else {
            v16 = append(v16, 127)
            var v30 uint = 1
            _ = v30
        }
        var v31 int = 1
        var v32 float64 = (0.0 + ((1e10 / 3.14159) + (-2.25 - -2.25)))
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 int32 = 1
        if false {
            var v35 float32 = 0.1
            _ = len(v16)
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    v36 := GenS4{F0: 2147483647}
    var v37 int8 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}

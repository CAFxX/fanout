package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int16 = 1
    var v2 int32 = 42
    var v3 uint8 = 128
    v4 := []string{"a", ("" + "a")}
    v5 := make(map[string]uint32)
    if (true && true) {
        v6 := GenS1{F0: 9223372036854775807, F1: 42}
        _ = v6
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v2, v1, v3)
    for fi7 := 0; fi7 < 9; fi7++ {
        delete(v5, "\n\t")
        var v8 float64 = (-2.25 + 1e10)
        if (790 != 9223372036854775807) {
            var v9 int64 = 9223372036854775807
            fmt.Println(v3)
            for fi10 := 0; fi10 < 5; fi10++ {
                v1 -= v1
                switch "hello" {
                case "\n\t":
                    v8 -= (v8 - (0.0 + (v8 / ((1.5 + v8) / (((-2.25 / 3.14159) / (v8 + ((v8 / (3.14159 - ((v8 * (1.5 - v8)) * float64(0.1)))) / v8))) + v8)))))
                case "a":
                    v8 += v8
                }
                _ = fi10
            }
            _ = v9
        } else {
            fmt.Println(v2)
            switch fi7 {
            case 32768:
                var v11 float64 = v8
                _ = v11
            }
        }
        _ = fi7
        _ = v8
    }
    _ = len(v4)
    delete(v5, "hello")
    var v12 uint8 = 255
    var v13 int16 = (-(v1))
    var v14 int64 = -1
    _ = len(v4)
    if ((!(true && (!(2147483647 == 1)))) || ("a" > "")) {
        if false {
            if true {
                var v15 string = ""
                if len(v4) > 0 {
                    v4[0] = (((v15 + "a") + "") + "foo bar")
                }
                _ = v15
            }
            fmt.Println(v1)
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v3, v14, v13)
            switch "" {
            case "a":
                if (!false) {
                    switch fi16 {
                    case 128:
                        fmt.Println(fi16, v14)
                    default:
                        v17 := GenS2{F0: 100}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 uint64 = 42
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        _ = fi18
                        _ = v19
                    }
                    var v20 int32 = (v2 & (v2 * v2))
                    _ = v20
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case "hello":
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            }
            _ = len(v4)
            _ = fi16
        }
        v12 = 42
    }
    v3 = (v12 << ((2) & 63))
    switch ("\n\t" + "foo bar") {
    case "a":
        v21 := GenS3{F0: v12, F1: false, F2: v1}
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v2 -= -1
        var v23 int = fi22
        v24 := GenS4{F0: 255}
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: -1}
    if len(v4) > 0 {
        v4[0] = (("foo bar" + "") + ("a" + (("" + "hello") + ("hello" + ""))))
    }
    fmt.Println(v12, v14, v3)
    v12 -= 142
    v26 := GenS4{F0: uint16(v12)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

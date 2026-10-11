package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0() string {
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        switch "\n\t" {
        case "a":
            var v4 int32 = 0
            _ = v4
        default:
            if (1 >= -1) {
                var v5 int16 = 32767
                _ = v5
            } else {
                var v6 float32 = 1.5
                fmt.Println(v6)
                _ = v6
            }
        }
    case "":
        var v7 uint16 = 592
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: 255, F1: -1}
        _ = v8
    default:
        var v9 int16 = 42
        _ = v9
    }
    var v10 uint64 = 18446744073709551615
    switch 32768 {
    case 42:
        fmt.Println(v10)
    }
    var v11 uint = 918
    _ = v10
    _ = v11
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v12 uint8 = 2
    v12 = (^(v12))
    var v13 uint16 = 1
    fmt.Println(v12, v12)
    var v14 bool = (!(true && true))
    if v14 {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint32 = 2
            _ = fi15
            _ = v16
        }
        v13 = v13
    }
    var v17 int8 = 0
    for fi18 := 0; fi18 < 12; fi18++ {
        if (-1 <= 443) {
            switch "hello" {
            case "\n\t":
                switch (fi18 * (^(fi18))) {
                case 9223372036854775807:
                    v14 = (!v14)
                case -737:
                    v17 += (-(v17))
                }
            case "foo bar":
                for fi19 := 0; fi19 < 5; fi19++ {
                    if (v14 && (!(!(float32((float64(3.14159) - 1.5)) <= float32((((((float64(float32((0.1 + (1.5 / (3.14159 / (-2.25 + 1e10)))))) * -2.25) / ((1.5 / (3.14159 / float64(((-2.25 + ((1e10 - ((((3.14159 - 1.5) + (1e10 / float32((0.0 + ((0.0 - -2.25) / 3.14159))))) - ((1e10 + 1e10) + 1e10)) - 1.5)) / float32(((3.14159 - (0.1 - (3.14159 - ((0.1 / 0.1) + 3.14159)))) * 0.1)))) / 3.14159)))) + -2.25)) * 1e10) * (float64(0.0) - (float64(1e10) * 1e10))) / 1.0)))))) {
                        if (468 != 1) {
                            v12 += v12
                        } else {
                            v14 = v14
                        }
                        v12 = uint8(v13)
                    }
                    v20 := GenS2{F0: (v17 | (v17 - (v17 + (v17 >> ((18446744073709551615) & 63))))), F1: (v13 + (^(v13)))}
                    if ((9223372036854775807 != 548) || true) {
                        switch (("\n\t" + ("" + ("\n\t" + "a"))) + "") {
                        case "a":
                            var v21 float64 = (1.5 + -2.25)
                            _ = v21
                        case "hello":
                            v22 := v1[("a" + "hello")]
                            _ = v22
                        case "":
                            if (!v14) {
                                v23 := GenS2{F0: 127, F1: (v13 & (v13 * v13))}
                                var v24 uint16 = (v13 | v13)
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                    v2 = append(v2, 339)
                    v25 := GenS3{F0: -539}
                    var v26 uint = 65535
                    _ = fi19
                    _ = v20
                    _ = v25
                    _ = v26
                }
            case "a":
                _ = len(v1)
            default:
                v27 := GenS3{F0: 42}
                _ = v27
            }
            v12 += v12
        }
        switch (("\n\t" + ("" + "a")) + "foo bar") {
        case "a":
            for fi28 := 0; fi28 < 16; fi28++ {
                if false {
                    v12 = (v12 | (^(v12)))
                    if (!("a" < "")) {
                        var v29 uint32 = 2147483648
                        _ = len(v3)
                        v17 -= v17
                        _ = v29
                    }
                } else {
                    fmt.Println(v12, v14, fi18)
                    v13 = (v13 >> ((1) & 63))
                    v17 -= 7
                }
                _ = fi28
            }
        case "foo bar":
            v17 += (^(v17))
        case "\n\t":
            v30 := GenS4{F0: true, F1: 216}
            _ = v30
        default:
            switch fi18 {
            case 0:
                _ = len(v2)
            case 272:
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            }
        }
        var v31 int32 = 2147483647
        _ = fi18
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}

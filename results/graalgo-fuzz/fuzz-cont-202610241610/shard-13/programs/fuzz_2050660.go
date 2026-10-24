package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 int32, p1 int8) uint32 {
    switch 2147483647 {
    case -8:
        v6 := GenS1{F0: 0}
        _ = v6
    case 42:
        v7 := []int{-392}
        _ = v7
    }
    var v8 uint16 = 255
    _ = p0
    _ = p1
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 bool = false
    var v4 uint64 = 2
    v5 := make(map[int]uint32)
    v3 = v3
    for fi9 := 0; fi9 < 10; fi9++ {
        v4 = (v4 << ((218) & 63))
        _ = fi9
    }
    if ((("a" + ("" + ("hello" + ((("" + "a") + "a") + "foo bar")))) + "\n\t") >= ("a" + (("" + "hello") + "\n\t"))) {
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS2{F0: 764}
                _ = fi11
                _ = v12
            }
            var v13 bool = v3
            _ = fi10
            _ = v13
        }
    }
    switch -9223372036854775808 {
    case 127:
        if (215 <= 113) {
            var v14 int64 = -1
            _ = v14
        }
    case -127:
        switch "\n\t" {
        case "\n\t":
            v2 = append(v2, 0)
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 float64 = 3.14159
                if v3 {
                    var v17 int8 = 1
                    _ = v17
                }
                var v18 uint = 1
                _ = fi15
                _ = v16
                _ = v18
            }
        }
    default:
        v4 -= v4
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        switch ("hello" + ("" + "a")) {
        case "foo bar":
            var v20 int32 = 2147483647
            _ = v20
        default:
            if false {
                v21 := GenS3{F0: v4, F1: 0}
                switch fi19 {
                case -920:
                    var v22 int32 = 728
                    _ = v22
                case 127:
                    if ((("\n\t" + "hello") != ((("\n\t" + ("foo bar" + "foo bar")) + "a") + ("foo bar" + ("a" + "a")))) && false) {
                        v3 = v3
                        if (!(1 >= 1)) {
                            v3 = ((1e10 + float32((1.5 * ((1.5 / 3.14159) - (0.0 - 0.1))))) == 0.0)
                            v23 := v1["foo bar"]
                            _ = v23
                        } else {
                            v24 := v5[(fi19 ^ (fi19 / (((fi19 + fi19)) | 1)))]
                            var v25 bool = (!v3)
                            _ = v24
                            _ = v25
                        }
                    }
                case 42:
                    fmt.Println(v4, fi19, v3)
                }
                switch "" {
                case "a":
                    v3 = (v3 && (!true))
                default:
                    v26 := GenS3{F0: v4, F1: 32767}
                    _ = v26
                }
                _ = v21
            }
        }
        v3 = (0 >= 19)
        for fi27 := 0; fi27 < 9; fi27++ {
            fmt.Println(v4, fi27, v3)
            if (32767 == 578) {
                var v28 uint64 = v4
                _ = v28
            } else {
                v29 := GenS4{F0: 2147483647, F1: (fi19 ^ 4294967295)}
                delete(v1, "a")
                v3 = (!v3)
                _ = v29
            }
            v30 := GenS5{F0: 80}
            v31 := GenS6{F0: -128, F1: (v4 % ((v4) | 1))}
            if ((v4 / (((v4 >> ((65535) & 63))) | 1)) == v4) {
                switch "a" {
                case "hello":
                    if (!v3) {
                        var v32 uint8 = 42
                        _ = v32
                    }
                case "":
                    var v33 int16 = 1000
                    _ = v33
                case "\n\t":
                    v2 = append(v2, 413)
                }
            } else {
                v4 = v4
            }
            _ = fi27
            _ = v30
            _ = v31
        }
        _ = fi19
    }
    var v34 int = 32768
    switch "" {
    case "hello":
        v4 = v4
    }
    var v35 float64 = float64((float32(0.0) / ((1e10 + (3.14159 * 1.5)) - 0.1)))
    v4 = (v4 - v4)
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := GenS2{F0: 914}
        _ = fi36
        _ = v37
    }
    if false {
        for fi38 := 0; fi38 < 19; fi38++ {
            fmt.Println(v35)
            fmt.Println(fi38, v3, v4)
            _ = fi38
        }
    } else {
        var v39 uint64 = 0
        var v40 bool = ((v34 >> ((239) & 63)) <= v34)
        v4 = (v39 % ((18446744073709551615) | 1))
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}

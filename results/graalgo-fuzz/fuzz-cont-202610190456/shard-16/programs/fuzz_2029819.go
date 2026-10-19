package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
func genFn0(p0 int64) int32 {
    if ((-1 < 605) && (0 < 32767)) {
        p0 = p0
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 uint8 = 215
                var v8 uint16 = 203
                p0 = p0
                _ = fi6
                _ = v7
                _ = v8
            }
            p0 -= 1
            v9 := GenS1{F0: 2147483648, F1: false, F2: 1}
            _ = fi5
            _ = v9
        }
    } else {
        p0 = p0
    }
    p0 = p0
    var v10 int = 2147483647
    _ = p0
    _ = v10
    return -320
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    var v3 uint16 = 786
    var v4 uint16 = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        switch fi11 {
        case 9223372036854775807:
            for fi12 := 0; fi12 < 5; fi12++ {
                v1 += (v1 - v1)
                switch ("\n\t" + "hello") {
                case "":
                    if true {
                        v4 += 859
                        var v13 float64 = float64(((float32(0.0) + 3.14159) * 3.14159))
                        _ = v13
                    }
                default:
                    v14 := []int{fi11}
                    _ = v14
                }
                v1 += (^(v1))
                fmt.Println(fi11)
                if (2 > 42) {
                    if (146 <= 18446744073709551615) {
                        v15 := GenS1{F0: uint32(v4), F1: (-1 != 408), F2: -2147483648}
                        var v16 int64 = 9223372036854775807
                        v17 := v2[fi12]
                        v17 = ("" + "\n\t")
                        v18 := GenS1{F0: 421, F1: (-128 <= 1), F2: (v1 * genFn0((v16 - (v16 >> ((2) & 63)))))}
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v19 := v2[1]
                        v19 = ("hello" + v19)
                        if (!(-29 <= 952)) {
                            v3 = v3
                            fmt.Println(v1, v4)
                            v4 -= v4
                        } else {
                            v20 := GenS1{F0: 727, F1: false, F2: v1}
                            delete(v2, 42)
                            switch fi11 {
                            case 4294967295:
                                v19 = "hello"
                            case -1:
                                var v21 int32 = (v1 | v1)
                                _ = v21
                            }
                            _ = v20
                        }
                        _ = v19
                    }
                    fmt.Println(v4, fi12)
                }
                var v22 string = "\n\t"
                _ = fi12
                _ = v22
            }
        default:
            v4 += 0
        }
        v3 -= (v3 * (v3 >> ((310) & 63)))
        _ = fi11
    }
    var v23 float64 = 1.5
    v23 -= float64((((0.0 / 1.0) / (0.1 * 1e10)) * float32(1e10)))
    v24 := GenS2{F0: 4294967295, F1: false, F2: 4294967295}
    var v25 uint = 42
    if false {
        delete(v2, 1)
        switch ("foo bar" + ("hello" + "")) {
        case "hello":
            if true {
                var v26 string = "\n\t"
                _ = v26
            }
        case "":
            var v27 string = ""
            _ = v27
        default:
            v1 -= genFn0(-543)
        }
        fmt.Println(v4, v23, v25)
    }
    v1 -= 1
    delete(v2, 4294967295)
    v28 := v2[-1]
    v29 := []int{-812, 127, 2147483647, 32768}
    fmt.Println(v25)
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := v2[128]
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}

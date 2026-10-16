package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0() uint64 {
    var v5 uint32 = 104
    v6 := []int{475}
    switch "\n\t" {
    case "foo bar":
        var v7 uint32 = 2
        _ = v7
    case "hello":
        var v8 uint64 = 4294967295
        _ = v8
    case "":
        fmt.Println(v5)
    default:
        v6 = append(v6, -1)
    }
    _ = v5
    _ = v6
    return 60
}

func main() {
    var v1 int16 = 1
    var v2 uint64 = 24
    v3 := make(map[int]bool)
    var v4 uint16 = 1
    fmt.Println(v4)
    var v9 bool = (-9223372036854775808 < 98)
    var v10 int = 32768
    v11 := GenS1{F0: 83}
    var v12 string = ""
    var v13 float32 = (((-2.25 * (1e10 / (((((3.14159 - (-2.25 * (0.1 / (1.5 - ((3.14159 - float32(3.14159)) / 1.0))))) * (((-2.25 - -2.25) + 0.0) * float32((0.0 - 0.0)))) / 0.1) - (3.14159 / -2.25)) + ((0.0 * (1e10 * 1.5)) - (1e10 - -2.25))))) * -2.25) * 3.14159)
    if v9 {
        if false {
            _ = len(v3)
        }
        if (!v9) {
            v14 := GenS2{F0: -9223372036854775808, F1: 42}
            v15 := GenS2{F0: -1, F1: 1}
            _ = v14
            _ = v15
        } else {
            if (127 < 247) {
                if (0 >= (v4 + 1)) {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            switch ("\n\t" + ("" + ((((v12 + (v12 + ("" + v12))) + v12) + v12) + "\n\t"))) {
                            case "foo bar":
                                fmt.Println(v1, v9, fi17)
                            case "hello":
                                v13 = v13
                            }
                            v13 -= (((v13 - (v13 + (v13 + ((v13 + v13) - (v13 + (v13 - 0.1)))))) / v13) - v13)
                            fmt.Println(v13)
                            _ = fi17
                        }
                        v18 := GenS2{F0: 9223372036854775807, F1: 718}
                        var v19 uint16 = v4
                        _ = fi16
                        _ = v18
                        _ = v19
                    }
                    switch "" {
                    case "foo bar":
                        if ((v13 - v13) > v13) {
                            var v20 float64 = (1.5 / (3.14159 * -2.25))
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v22 := GenS3{F0: 872, F1: 2, F2: 1}
                                v20 = 1.5
                                var v23 float32 = v13
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            fmt.Println(v20, v13, v4)
                            _ = v20
                        }
                    }
                } else {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := GenS3{F0: 42, F1: 87, F2: -9223372036854775808}
                        _ = fi24
                        _ = v25
                    }
                    fmt.Println(v12, v13, v13)
                }
            }
            if false {
                v26 := GenS2{F0: 0, F1: 0}
                _ = v26
            }
        }
        switch (v12 + (v12 + (v12 + v12))) {
        case "":
            switch -317 {
            case 128:
                var v27 uint32 = 65535
                _ = v27
            default:
                if ((!v9) && ((1e10 != (0.1 - v13)) || false)) {
                    switch v12 {
                    case "\n\t":
                        switch (v12 + v12) {
                        case "":
                            var v28 bool = v9
                            _ = v28
                        case "foo bar":
                            var v29 uint = 65535
                            _ = v29
                        }
                    case "hello":
                        switch (v10 ^ v10) {
                        case 4294967295:
                            var v30 string = ("" + v12)
                            _ = v30
                        case -2147483648:
                            for fi31 := 0; fi31 < 6; fi31++ {
                                var v32 uint8 = 252
                                v33 := GenS3{F0: 42, F1: (v32 % ((v32) | 1)), F2: 1}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    default:
                        if true {
                            v34 := []int{v10, v10, v10}
                            _ = v34
                        } else {
                            v10 = v10
                        }
                    }
                    v10 += 32767
                    fmt.Println(v1)
                } else {
                    _ = len(v3)
                }
            }
        case "hello":
            v9 = ((v2 % ((545) | 1)) != (v2 >> ((42) & 63)))
        }
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v12, v2, v2)
        _ = fi35
    }
    delete(v3, (v10 & (v10 >> ((0) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}

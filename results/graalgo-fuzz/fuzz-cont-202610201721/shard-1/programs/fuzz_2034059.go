package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    var v1 uint64 = 852
    var v2 int64 = 9223372036854775807
    var v3 int32 = 773
    v4 := make(map[string]int)
    var v5 int = 397
    var v6 bool = (!(!true))
    _ = len(v4)
    v6 = (((v1 >> ((626) & 63)) < v1) || (v6 || (!v6)))
    v5 = v5
    v5 -= (v5 ^ 2147483647)
    v3 = 2147483647
    var v7 uint64 = (v1 - 0)
    v6 = v6
    var v8 uint = 519
    switch v5 {
    case 1:
        if v6 {
            switch 42 {
            case 2147483647:
                for fi9 := 0; fi9 < 5; fi9++ {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        fmt.Println(v6, fi9, v8)
                        v7 += v1
                        fmt.Println(fi10, v2, v3)
                        _ = fi10
                    }
                    _ = fi9
                }
            case -9223372036854775808:
                fmt.Println(v3, v3, v2)
            case -1:
                switch (v5 % (((v5 % ((1) | 1))) | 1)) {
                case 32768:
                    var v11 float64 = 3.14159
                    _ = v11
                default:
                    var v12 int = (v5 & v5)
                    _ = v12
                }
            }
            v3 -= (v3 & 0)
        } else {
            fmt.Println(v8)
        }
    case 9223372036854775807:
        switch "hello" {
        case "\n\t":
            var v13 string = "hello"
            _ = v13
        case "":
            if v6 {
                if v6 {
                    var v14 uint32 = 42
                    if (v5 <= v5) {
                        v6 = v6
                        var v15 uint = 1
                        for fi16 := 0; fi16 < 14; fi16++ {
                            var v17 string = "a"
                            v18 := GenS1{F0: v3, F1: int8(v8), F2: 2}
                            if ("\n\t" <= (v17 + "foo bar")) {
                                v19 := []int{fi16, (-(fi16)), -94}
                                v19 = append(v19, v5)
                                fmt.Println(v6, v3, v5)
                                _ = v19
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = v15
                    } else {
                        var v20 uint8 = 245
                        for fi21 := 0; fi21 < 3; fi21++ {
                            v22 := GenS1{F0: v3, F1: 51, F2: (v7 % (((v7 & v1)) | 1))}
                            var v23 float32 = (((0.0 / ((3.14159 + (float32((float64(1.5) * 3.14159)) * 1e10)) * (float32(1e10) * ((float32(0.0) + 0.1) + 1e10)))) * ((0.0 * 1e10) + ((1.5 * 1e10) + 0.1))) * float32(1.5))
                            v6 = true
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        _ = v20
                    }
                    v7 += 1
                    _ = v14
                } else {
                    var v24 uint64 = v1
                    if (!((v1 * (v24 | v1)) < v7)) {
                        for fi25 := 0; fi25 < 1; fi25++ {
                            if false {
                                v1 = (v1 & (v1 * v24))
                            }
                            var v26 uint = (v8 & 0)
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        var v27 uint32 = 2147483648
                        var v28 uint16 = 1
                        _ = v27
                        _ = v28
                    }
                    v6 = ((v8 ^ (v8 % (((^(v8))) | 1))) >= 255)
                    _ = v24
                }
                var v29 uint16 = 32767
                var v30 uint8 = 127
                _ = v29
                _ = v30
            } else {
                fmt.Println(v8, v5)
                v31 := GenS2{F0: v2, F1: (v5 ^ (v5 | (^(v5))))}
                _ = v31
            }
        default:
            v32 := []int{(v5 ^ v5), (v5 / (((v5 + (v5 & v5))) | 1))}
            _ = v32
        }
    case -2147483648:
        for fi33 := 0; fi33 < 19; fi33++ {
            v6 = ("\n\t" != "hello")
            delete(v4, "a")
            _ = fi33
        }
    default:
        var v34 bool = ((-580 <= 1000) || v6)
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

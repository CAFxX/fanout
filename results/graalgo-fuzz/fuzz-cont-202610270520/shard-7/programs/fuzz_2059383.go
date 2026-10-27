package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int64 = 42
    var v2 uint64 = 42
    v3 := []int{-1, -166}
    var v4 uint64 = v2
    var v5 uint = uint(v4)
    v6 := GenS1{F0: (v2 / (((v4 - 127)) | 1))}
    var v7 int64 = v1
    var v8 int = -58
    if false {
        v5 = v5
        if (!(!true)) {
            v9 := GenS2{F0: 32767, F1: 42}
            var v10 uint16 = 801
            _ = v10
            _ = v9
        } else {
            v4 = (v4 >> ((v5) & 63))
            v3 = append(v3, v8)
        }
    }
    switch v8 {
    case 948:
        var v11 float32 = (1.5 + (1.5 + 0.0))
        _ = v11
    case 42:
        v8 -= 374
    case 9223372036854775807:
        var v12 bool = true
        _ = v12
    default:
        var v13 int8 = 42
        _ = v13
    }
    var v14 uint8 = 128
    for fi15 := 0; fi15 < 19; fi15++ {
        v7 = (-(v7))
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int8 = -128
            _ = fi16
            _ = v17
        }
        v4 = (v2 / ((1) | 1))
        _ = fi15
    }
    v18 := GenS1{F0: (v2 * v2)}
    v19 := GenS2{F0: -760, F1: -128}
    var v20 int16 = 32767
    var v21 string = "foo bar"
    var v22 uint = 0
    var v23 bool = (((42 > 65535) && (int(v4) < v8)) || ((0 >= v14) && true))
    _ = len(v3)
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(fi24, fi24, v7)
        fmt.Println(v4)
        _ = fi24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            if ((!(v23 || v23)) && (false || ((3.14159 + (1.5 * 0.0)) >= (float64((3.14159 / 3.14159)) * 0.0)))) {
                v2 += 42
                v27 := GenS3{F0: 78}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v27
            }
            v21 = v21
            if ((("\n\t" + v21) >= v21) && (int16(v22) != 2)) {
                switch "\n\t" {
                case "foo bar":
                    v2 += v2
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    v22 = v22
                }
                v14 += (^(v14))
            } else {
                switch v8 {
                case 2147483647:
                    v3 = append(v3, (fi25 ^ fi25))
                case 127:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -9223372036854775808:
                    v3 = append(v3, (-(fi26)))
                }
                var v28 uint32 = 380
                _ = v28
            }
            _ = fi26
        }
        if v23 {
            v7 = v7
        } else {
            v22 += (v22 >> (((v5 % (((v22 + 2)) | 1))) & 63))
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(v5, fi29)
            _ = fi29
        }
        _ = fi25
    }
    var v30 int16 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

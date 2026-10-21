package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint64 }

func main() {
    v1 := []int{32767, 128, -9223372036854775808}
    var v2 uint32 = 4294967295
    v3 := make(map[string]string)
    var v4 int = 9223372036854775807
    var v5 int64 = -1
    v6 := []uint32{(^(v2)), (v2 << ((255) & 63)), 2, v2}
    v3[("" + "a")] = ("foo bar" + "a")
    v7 := GenS1{F0: v5}
    switch ("hello" + "") {
    case "hello":
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(v4, fi8)
            _ = fi8
        }
    case "a":
        v5 -= v5
    case "":
        v9 := GenS1{F0: (v5 * 1)}
        _ = v9
    }
    fmt.Println(v4)
    if false {
        delete(v3, "\n\t")
        v5 -= (v5 << ((4294967295) & 63))
    }
    fmt.Println(v4)
    var v10 uint32 = v2
    var v11 uint8 = 2
    var v12 string = "\n\t"
    if false {
        var v13 float32 = (0.1 / float32((0.1 * 0.1)))
        switch v12 {
        case "a":
            if (1 != 0) {
                fmt.Println(v13)
            } else {
                var v14 bool = true
                if ((v2 / ((v2) | 1)) == (v2 << ((691) & 63))) {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        fmt.Println(fi15)
                        v3[v12] = v12
                        _ = fi15
                    }
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 bool = (!v14)
                        if ((v5 ^ int64(v11)) == (v5 | (-(v5)))) {
                            v18 := GenS1{F0: (v5 + 1)}
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                } else {
                    var v19 int16 = 1000
                    _ = v19
                }
                var v20 float32 = 0.1
                _ = v14
                _ = v20
            }
        case "\n\t":
            switch 1 {
            case 9223372036854775807:
                fmt.Println(v12, v10)
            }
        }
        _ = v13
    } else {
        var v21 uint8 = v11
        var v22 int16 = 32767
        _ = v21
        _ = v22
    }
    var v23 int32 = 1
    if false {
        if false {
            if (0.0 < -2.25) {
                var v24 int = v4
                _ = v24
            }
        } else {
            v25 := GenS2{F0: 621}
            v26 := GenS3{F0: ((!(true && false)) || (100 >= 1)), F1: 641, F2: 2}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v25
            _ = v26
        }
        fmt.Println(v11, v11, v23)
    }
    v2 = v2
    if (255 < 2) {
        v27 := GenS2{F0: int16(v4)}
        _ = v27
    }
    var v28 uint64 = uint64(v2)
    var v29 float32 = 0.0
    var v30 bool = ((v23 - 1) >= -827)
    fmt.Println(v30, v29)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

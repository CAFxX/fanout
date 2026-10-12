package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = (1.5 + (1.5 + float64((float32(0.0) / -2.25))))
    var v2 uint64 = 1
    var v3 int8 = -128
    v4 := []int{4294967295}
    var v5 int64 = 42
    var v6 float32 = 3.14159
    var v7 bool = (!(!(true && (("\n\t" + "hello") >= "hello"))))
    for fi8 := 0; fi8 < 8; fi8++ {
        v5 -= 42
        if len(v4) > 0 {
            v4[0] = -235
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v4 = append(v4, (fi8 + (fi9 / (((fi9 << ((68) & 63))) | 1))))
            v10 := GenS1{F0: v5}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if false {
        fmt.Println(v6, v2, v6)
        v2 += 42
        fmt.Println(v1, v5, v7)
    } else {
        v4 = append(v4, -2147483648)
        var v11 int = -9223372036854775808
        switch v11 {
        case 2147483647:
            if (!(65535 == 169)) {
                v6 = (((1e10 * (1.5 * 0.1)) - ((1.5 + ((float32(0.0) / v6) - (v6 + 0.1))) * 1.5)) + v6)
            } else {
                v2 -= (v2 + (^(v2)))
            }
        default:
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: (^(v5))}
                var v14 int32 = -1
                if v7 {
                    switch ("hello" + "hello") {
                    case "":
                        var v15 uint64 = (v2 - (^(v2)))
                        _ = v15
                    case "foo bar":
                        if true {
                            for fi16 := 0; fi16 < 7; fi16++ {
                                v17 := GenS2{F0: 128}
                                v7 = (-111 == (v3 | -1))
                                v5 = v5
                                _ = fi16
                                _ = v17
                            }
                            if len(v4) > 0 {
                                v4[0] = v11
                            }
                            v18 := GenS2{F0: 255}
                            _ = v18
                        } else {
                            fmt.Println(v14)
                        }
                    case "\n\t":
                        fmt.Println(v7, v1, fi12)
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = v11
    }
    v3 += (v3 / ((v3) | 1))
    var v19 uint16 = 65535
    var v20 string = "hello"
    for fi21 := 0; fi21 < 15; fi21++ {
        v20 = "a"
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 uint16 = v19
            if len(v4) > 0 {
                v4[0] = (fi21 + 0)
            }
            switch fi21 {
            case -2147483648:
                v4 = append(v4, fi21)
            case -9223372036854775808:
                var v24 uint8 = 78
                _ = v24
            }
            _ = fi22
            _ = v23
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint32 = 2147483648
            v4 = append(v4, (fi21 / ((fi25) | 1)))
            _ = fi25
            _ = v26
        }
        _ = fi21
    }
    if (((!v7) || (!v7)) && (!(!v7))) {
        v7 = v7
        var v27 int16 = -1
        _ = v27
    }
    var v28 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

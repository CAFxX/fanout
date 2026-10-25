package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 65535
    var v3 int64 = 9223372036854775807
    var v4 string = ""
    var v5 uint64 = 415
    var v6 string = v4
    v7 := GenS1{F0: 0, F1: -2147483648}
    v8 := []int{0}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v6 = (("\n\t" + "\n\t") + v4)
    v9 := GenS1{F0: 1000, F1: 638}
    v1 += float64(float32(v1))
    var v10 string = "\n\t"
    switch -9223372036854775808 {
    case -1:
        switch v4 {
        case "hello":
            for fi11 := 0; fi11 < 11; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v1 = 1e10
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi12
                }
                if (!((v3 * v3) > -1)) {
                    var v13 int16 = 1000
                    v14 := GenS1{F0: 510, F1: 126}
                    _ = v13
                    _ = v14
                } else {
                    var v15 float32 = 0.0
                    switch (fi11 - fi11) {
                    case -144:
                        v3 = int64(fi11)
                    case 9223372036854775807:
                        for fi16 := 0; fi16 < 6; fi16++ {
                            if false {
                                v1 += 1.5
                                fmt.Println(v2, fi16, v6)
                            } else {
                                var v17 uint16 = (v2 ^ (^(v2)))
                                v18 := GenS1{F0: 0, F1: -974}
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(v15, fi11, fi16)
                            v10 = (v10 + v10)
                            _ = fi16
                        }
                    case 4294967295:
                        fmt.Println(v5)
                    default:
                        if len(v8) > 0 {
                            v8[0] = (fi11 ^ fi11)
                        }
                    }
                    fmt.Println(v1, v1, fi11)
                    _ = v15
                }
                _ = fi11
            }
        }
    default:
        var v19 int64 = v3
        _ = v19
    }
    switch v4 {
    case "":
        v5 = 786
    case "\n\t":
        v8 = append(v8, 42)
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v1)
        switch (v4 + "hello") {
        case "foo bar":
            var v21 int8 = 1
            _ = v21
        case "a":
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 uint16 = (^(v2))
                _ = fi22
                _ = v23
            }
        default:
            v10 = (v4 + "a")
        }
        v8 = append(v8, (^(fi20)))
        _ = fi20
    }
    v10 = v10
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v24 int8 = 100
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

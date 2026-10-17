package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 uint; F1 uint16 }

func main() {
    var v1 float64 = 1e10
    var v2 bool = ((("\n\t" + "a") + "") == "\n\t")
    var v3 uint16 = 420
    v1 = 1.5
    if true {
        v4 := GenS1{F0: 4294967295, F1: 1, F2: 39}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 3; fi5++ {
            switch (fi5 % ((fi5) | 1)) {
            case -568:
                fmt.Println(v3, v2, v1)
            default:
                if false {
                    for fi6 := 0; fi6 < 12; fi6++ {
                        v2 = v2
                        v7 := []int{(fi6 % (((fi6 / ((fi6) | 1))) | 1)), fi6, (fi6 | fi5)}
                        _ = fi6
                        _ = v7
                    }
                } else {
                    if true {
                        v1 = v1
                        var v8 int32 = -188
                        fmt.Println(fi5, v1, v3)
                        _ = v8
                    } else {
                        fmt.Println(v3, v3)
                        switch "" {
                        case "\n\t":
                            v9 := []int{fi5, fi5, fi5}
                            _ = v9
                        }
                    }
                    v10 := GenS1{F0: (fi5 / ((fi5) | 1)), F1: 90, F2: 0}
                    v11 := []int{127, (fi5 >> ((42) & 63)), 573}
                    _ = v10
                    _ = v11
                }
            }
            var v12 int8 = 42
            var v13 int16 = 1000
            _ = fi5
            _ = v12
            _ = v13
        }
        switch "" {
        case "hello":
            fmt.Println(v1, v1)
        case "\n\t":
            var v14 int64 = 9223372036854775807
            _ = v14
        }
        v1 += v1
    }
    var v15 uint8 = 108
    var v16 uint32 = 2147483647
    v17 := []int{161, 9223372036854775807, 158}
    var v18 uint = 18446744073709551615
    var v19 int64 = 926
    v20 := GenS1{F0: 127, F1: v18, F2: 1}
    if ((false && v2) && v2) {
        for fi21 := 0; fi21 < 12; fi21++ {
            v3 += (v3 + v3)
            var v22 uint = (v18 + v18)
            v2 = v2
            _ = fi21
            _ = v22
        }
    }
    v17 = append(v17, -9223372036854775808)
    var v23 int64 = (v19 * 1)
    if v2 {
        if v2 {
            var v24 string = ("hello" + "")
            _ = v24
        } else {
            var v25 uint16 = (v3 ^ 295)
            fmt.Println(v23, v25, v18)
            v26 := GenS2{F0: (v18 % ((v18) | 1)), F1: (v3 << (((v18 - (v18 >> ((v18) & 63)))) & 63))}
            _ = v25
            _ = v26
        }
    }
    if ((v18 != (v18 % (((v18 | v18)) | 1))) && v2) {
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(v16, v19, v18)
            switch ("" + "foo bar") {
            case "":
                var v28 int32 = -179
                _ = v28
            case "\n\t":
                var v29 uint64 = 193
                _ = v29
            default:
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 float32 = (1.5 + 0.1)
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi27
        }
        var v32 float32 = -2.25
        _ = v32
    } else {
        fmt.Println(v2)
        for fi33 := 0; fi33 < 17; fi33++ {
            fmt.Println(v16, v3)
            v34 := GenS2{F0: (v18 >> (((v18 / (((v18 / (((v18 ^ v18)) | 1))) | 1))) & 63)), F1: v3}
            fmt.Println(v1)
            _ = fi33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        fmt.Println(v1, v15)
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        if len(v17) > 0 {
            v17[0] = int(v18)
        }
        _ = fi35
    }
    var v36 uint32 = uint32(v15)
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 uint8 = (v15 / (((v15 % (((v15 | (^(v15)))) | 1))) | 1))
        v39 := GenS2{F0: (^(v18)), F1: (v3 % ((2) | 1))}
        _ = fi37
        _ = v38
        _ = v39
    }
    v3 -= (v3 - (v3 ^ v3))
    var v40 uint32 = (v36 << (((^(v18))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v36
    _ = v40
}

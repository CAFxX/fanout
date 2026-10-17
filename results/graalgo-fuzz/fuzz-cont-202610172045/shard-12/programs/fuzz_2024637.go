package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 int16 }

func main() {
    var v1 int16 = 32767
    var v2 bool = true
    var v3 int64 = -1
    v4 := make(map[uint32]string)
    var v5 bool = v2
    var v6 int32 = -988
    v7 := []int{-2147483648}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch (fi8 & fi8) {
        case 137:
            v5 = v2
        case -638:
            switch (fi8 % ((int(v3)) | 1)) {
            case 127:
                switch (fi8 | fi8) {
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 3; fi9++ {
                        if (((("hello" + "hello") + "") + "\n\t") > ((("hello" + "\n\t") + "\n\t") + "a")) {
                            if (!false) {
                                v6 = v6
                                var v10 uint = 65535
                                fmt.Println(v3, v5)
                                _ = v10
                            }
                        }
                        var v11 int = 4294967295
                        delete(v4, 42)
                        v5 = (1 >= v6)
                        switch v11 {
                        case -199:
                            fmt.Println(v1, v11)
                        case 32768:
                            v7 = append(v7, v11)
                        default:
                            var v12 float32 = -2.25
                            _ = v12
                        }
                        _ = fi9
                        _ = v11
                    }
                case 127:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        fmt.Println(fi8)
                        _ = fi13
                    }
                }
            case 128:
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(v1)
                    _ = fi14
                }
            case 9223372036854775807:
                v15 := GenS1{F0: (-(v3))}
                _ = v15
            }
        }
        switch ("a" + "") {
        case "":
            v4[2147483648] = "foo bar"
        default:
            var v16 uint8 = 1
            _ = v16
        }
        _ = fi8
    }
    var v17 bool = true
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v2 = ((!false) || (false || (v6 == v6)))
            _ = fi19
        }
        _ = fi18
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS2{F0: v1, F1: 1, F2: v1}
        v22 := GenS2{F0: v1, F1: v6, F2: v1}
        delete(v4, 2)
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := GenS2{F0: 42, F1: 0, F2: (^(v1))}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v5 = v5
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 float32 = (3.14159 / 1e10)
    delete(v4, 1)
    var v27 uint32 = 42
    var v28 uint = 2
    fmt.Println(v1, v2)
    if (16 >= (v28 * uint(v27))) {
        v5 = (!(42 == 2))
    }
    switch "hello" {
    case "\n\t":
        var v29 uint8 = 2
        _ = v29
    case "":
        v2 = v17
    case "hello":
        for fi30 := 0; fi30 < 19; fi30++ {
            if v2 {
                if false {
                    v31 := GenS1{F0: (v3 / ((616) | 1))}
                    _ = v31
                }
                fmt.Println(v3, v3, fi30)
            } else {
                fmt.Println(v2, v26, v3)
                v28 += (v28 / ((uint(v27)) | 1))
                var v32 float32 = ((v26 + 3.14159) + (v26 * v26))
                _ = v32
            }
            v6 -= (v6 >> (((v28 | v28)) & 63))
            _ = fi30
        }
    }
    v3 -= v3
    v2 = (!true)
    if v17 {
        v27 -= v27
    } else {
        v5 = true
        if true {
            v7 = append(v7, int(v27))
            fmt.Println(v2, v1, v27)
            fmt.Println(v26)
        }
    }
    v3 += (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

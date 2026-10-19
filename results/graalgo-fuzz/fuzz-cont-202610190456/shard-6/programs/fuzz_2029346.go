package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint; F1 int }
func genFn0() uint64 {
    var v4 uint = 4294967295
    switch "a" {
    case "hello":
        if (65535 > 445) {
            for fi5 := 0; fi5 < 17; fi5++ {
                var v6 float64 = (1.5 * (1e10 / 1e10))
                _ = fi5
                _ = v6
            }
            fmt.Println(v4)
        }
    case "foo bar":
        fmt.Println(v4)
    default:
        for fi7 := 0; fi7 < 12; fi7++ {
            if (583 != 2147483647) {
                if false {
                    fmt.Println(fi7, fi7)
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 uint32 = 4294967295
                        v9 = v9
                        _ = fi8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 8; fi10++ {
                        if false {
                            if true {
                                v4 += (v4 | (v4 ^ 65535))
                            }
                            v11 := GenS1{F0: 127}
                            switch (fi7 ^ (fi10 * fi7)) {
                            case 0:
                                v12 := []int{(fi10 / ((fi10) | 1)), -9223372036854775808}
                                _ = v12
                            case 61:
                                v4 -= 962
                            case 32767:
                                if false {
                                    v4 += v4
                                    if true {
                                        v13 := GenS1{F0: 42}
                                        v4 -= v4
                                        v4 -= 0
                                        v14 := []int{(fi10 + 356), (fi10 << (((v4 ^ v4)) & 63))}
                                        var v15 uint32 = 65535
                                        v4 += v4
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            v16 := []int{fi7}
                                            _ = v16
                                        }
                                    }
                                }
                                v4 = 255
                            }
                            _ = v11
                        }
                        if (false || false) {
                            v4 = (v4 / ((v4) | 1))
                        } else {
                            v17 := GenS2{F0: 42}
                            fmt.Println(fi7, fi10, fi10)
                            _ = v17
                        }
                        fmt.Println(fi7, fi10, fi10)
                        _ = fi10
                    }
                }
                v18 := GenS3{F0: -126, F1: -32768}
                _ = v18
            } else {
                fmt.Println(v4)
                fmt.Println(fi7, fi7)
                var v19 uint16 = 735
                _ = v19
            }
            v4 = v4
            _ = fi7
        }
    }
    _ = v4
    return 4294967295
}

func main() {
    var v1 int32 = -1
    var v2 string = ("a" + ("\n\t" + "hello"))
    var v3 uint = 65535
    v20 := GenS4{F0: 599, F1: 32767}
    var v21 float32 = 3.14159
    v2 = v2
    var v22 int64 = -1
    for fi23 := 0; fi23 < 13; fi23++ {
        if false {
            v21 -= v21
            v24 := []int{(fi23 ^ fi23), fi23}
            if len(v24) > 0 {
                v24[0] = (fi23 >> ((v3) & 63))
            }
            _ = v24
        }
        switch (fi23 % ((int(v1)) | 1)) {
        case -1:
            v25 := []int{}
            _ = v25
        case 503:
            if (!(2 <= 32767)) {
                var v26 int64 = 1
                v22 += v22
                _ = v26
            }
        case 1:
            var v27 string = "a"
            _ = v27
        }
        v21 -= v21
        _ = fi23
    }
    var v28 int32 = v1
    v28 += v1
    for fi29 := 0; fi29 < 20; fi29++ {
        if true {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 bool = false
                _ = fi30
                _ = v31
            }
            v3 = 4294967295
            var v32 uint8 = 0
            _ = v32
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
}

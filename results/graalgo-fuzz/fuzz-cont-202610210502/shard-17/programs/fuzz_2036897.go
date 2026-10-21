package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint32; F2 int }
type GenS6 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 9223372036854775807
    var v3 uint = 2
    v4 := GenS1{F0: v3, F1: 782}
    v5 := GenS2{F0: 62, F1: (v1 <= (v1 | (v1 + v1))), F2: 1}
    var v6 int8 = 100
    v7 := GenS3{F0: v6}
    fmt.Println(v6, v6)
    switch "" {
    case "":
        var v8 uint16 = 22
        _ = v8
    case "a":
        switch 32767 {
        case -1:
            v9 := []int{1, -147}
            _ = v9
        case 0:
            switch (("foo bar" + "") + "a") {
            case "a":
                switch 9223372036854775807 {
                case 732:
                    var v10 uint16 = 1
                    _ = v10
                default:
                    v11 := GenS2{F0: 218, F1: (1.5 != -2.25), F2: -32768}
                    _ = v11
                }
            case "hello":
                v12 := []int{-1, int(v6)}
                _ = v12
            case "\n\t":
                if true {
                    v13 := []int{}
                    var v14 bool = (!false)
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(v1, v6, v3)
                    var v15 string = ("hello" + ("\n\t" + "a"))
                    _ = v15
                }
            }
        case 32767:
            var v16 uint = v3
            _ = v16
        }
    case "hello":
        var v17 uint16 = 1
        _ = v17
    }
    var v18 bool = false
    var v19 int32 = v1
    v20 := GenS4{F0: -445}
    var v21 uint8 = 127
    v22 := GenS5{F0: int64(v3), F1: 0, F2: 95}
    switch "a" {
    case "hello":
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                if false {
                    v19 = v1
                    v19 = int32(v3)
                    if (65535 >= 2147483647) {
                        v3 = uint(v21)
                        var v25 uint8 = v21
                        _ = v25
                    } else {
                        fmt.Println(v19, fi24, v2)
                        if v18 {
                            var v26 uint32 = 420
                            _ = v26
                        }
                        if true {
                            v1 += v19
                        } else {
                            switch fi23 {
                            case 860:
                                v27 := GenS3{F0: v6}
                                _ = v27
                            default:
                                var v28 int8 = v6
                                _ = v28
                            }
                            v29 := GenS3{F0: v6}
                            _ = v29
                        }
                    }
                } else {
                    fmt.Println(fi23, v1)
                    v19 += (v1 % (((^(v19))) | 1))
                    if ((!v18) && v18) {
                        fmt.Println(v21)
                        if (!(4294967295 < 257)) {
                            v3 -= v3
                            v30 := []int{(fi24 % (((fi24 ^ (fi24 / ((32768) | 1)))) | 1)), fi24, fi24, fi23}
                            v6 = v6
                            _ = v30
                        }
                    }
                }
                var v31 int32 = (^(v1))
                _ = fi24
                _ = v31
            }
            _ = fi23
        }
    case "\n\t":
        fmt.Println(v2)
    case "a":
        v32 := []int{-279}
        _ = v32
    default:
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := []int{}
            for fi35 := 0; fi35 < 1; fi35++ {
                var v36 int = fi35
                _ = fi35
                _ = v36
            }
            var v37 int64 = 0
            _ = fi33
            _ = v34
            _ = v37
        }
    }
    v38 := []int{636, 42}
    fmt.Println(v2, v19, v2)
    v18 = v18
    if v18 {
        var v39 uint32 = 42
        if v18 {
            var v40 int16 = -369
            v41 := GenS6{F0: v6, F1: 4}
            if v18 {
                v6 = 1
                v3 += (v3 & (v3 << (((v3 << ((v3) & 63))) & 63)))
            } else {
                for fi42 := 0; fi42 < 16; fi42++ {
                    fmt.Println(v2, v19, v39)
                    _ = fi42
                }
                v3 -= 583
            }
            _ = v40
            _ = v41
        } else {
            var v43 float64 = (1.5 / 1.0)
            _ = v43
        }
        switch "foo bar" {
        case "\n\t":
            v19 = v19
        case "foo bar":
            v44 := GenS3{F0: int8(v39)}
            _ = v44
        default:
            switch 127 {
            case 876:
                var v45 float32 = (1.5 / 3.14159)
                _ = v45
            case -138:
                var v46 bool = v18
                _ = v46
            case 32767:
                var v47 uint16 = 255
                _ = v47
            }
        }
        _ = v39
    } else {
        var v48 int32 = 1
        fmt.Println(v1, v3)
        if len(v38) > 0 {
            v38[0] = 42
        }
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

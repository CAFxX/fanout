package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }

func main() {
    var v1 bool = (-32768 != -370)
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 string = "a"
        var v6 float64 = (float64(0.1) + (3.14159 / 0.1))
        _ = fi4
        _ = v5
        _ = v6
    }
    v1 = (0 >= 0)
    var v7 string = "\n\t"
    v7 = (v7 + v7)
    var v8 string = v7
    fmt.Println(v8)
    var v9 int64 = 1
    fmt.Println(v9, v1, v2)
    if (v1 && (!(v8 == "foo bar"))) {
        fmt.Println(v2)
    } else {
        v9 -= -974
        v2 -= v2
    }
    switch -563 {
    case 42:
        if (18446744073709551615 == 65535) {
            fmt.Println(v9, v9)
            v10 := v3[v7]
            if v1 {
                if (42 != 623) {
                    v1 = v1
                    v11 := []int{}
                    v9 = 1
                    _ = v11
                } else {
                    var v12 float32 = 0.1
                    _ = v12
                }
                v2 = v2
            } else {
                var v13 uint64 = 2
                _ = v13
            }
            _ = v10
        } else {
            fmt.Println(v9)
        }
    case 406:
        var v14 int8 = 127
        _ = v14
    case -1:
        switch int(v2) {
        case 32767:
            for fi15 := 0; fi15 < 4; fi15++ {
                if true {
                    v16 := GenS1{F0: 0, F1: 42}
                    _ = v16
                }
                _ = fi15
            }
        case 1:
            v8 = v7
        case 32768:
            v1 = ((67 >= -128) || false)
        default:
            switch -949 {
            case 2147483647:
                for fi17 := 0; fi17 < 12; fi17++ {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v8 = v7
                        _ = fi18
                    }
                    var v19 float32 = -2.25
                    _ = fi17
                    _ = v19
                }
            }
        }
    }
    fmt.Println(v1, v8, v2)
    if (("hello" == v8) && v1) {
        switch (((v8 + (("" + v7) + v7)) + ("" + v7)) + "a") {
        case "\n\t":
            v7 = v7
        case "":
            v8 = v8
        case "hello":
            var v20 int32 = v2
            _ = v20
        }
        if v1 {
            var v21 float32 = float32((3.14159 - ((1e10 * 1.5) / 0.1)))
            if (!((2147483647 > 311) || (42 <= 60))) {
                var v22 int64 = (v9 % ((9223372036854775807) | 1))
                var v23 int32 = 1
                _ = v22
                _ = v23
            }
            v24 := []int{1, -2147483648, 32767}
            _ = v21
            _ = v24
        }
        if false {
            v25 := GenS1{F0: 42, F1: 198}
            _ = v25
        }
    } else {
        v3[v8] = 2147483648
        v26 := []int{-856, 32768, -1, -540}
        if (4294967295 < 1) {
            if len(v26) > 0 {
                v26[0] = 42
            }
            switch ("a" + v8) {
            case "hello":
                if (!v1) {
                    v27 := GenS1{F0: 32767, F1: 42}
                    var v28 uint = 42
                    fmt.Println(v28, v9, v8)
                    _ = v27
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 18; fi29++ {
                    if true {
                        var v30 float64 = (float64(0.0) + 0.0)
                        var v31 int8 = 1
                        var v32 uint = 0
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    if (!(!(1 > 247))) {
                        switch v7 {
                        case "\n\t":
                            var v33 int16 = -32768
                            _ = v33
                        case "foo bar":
                            var v34 float64 = 0.1
                            _ = v34
                        default:
                            var v35 int16 = 1
                            _ = v35
                        }
                    } else {
                        var v36 uint = 0
                        fmt.Println(fi29, fi29, v36)
                        _ = v36
                    }
                    _ = fi29
                }
            }
        }
        _ = v26
    }
    v37 := GenS2{F0: 0, F1: 2}
    if (false && (1 <= 42)) {
        for fi38 := 0; fi38 < 8; fi38++ {
            v9 = -9223372036854775808
            var v39 int32 = (-(v2))
            var v40 int = -379
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v7
    _ = v8
    _ = v9
}

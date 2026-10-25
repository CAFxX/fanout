package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 2
    var v3 int64 = -9223372036854775808
    var v4 bool = (!false)
    v5 := []string{"hello", ("a" + "a"), "\n\t"}
    v6 := make(map[int]int)
    var v7 uint16 = v2
    var v8 string = "hello"
    v4 = true
    if ((1.5 != -2.25) || v4) {
        var v9 uint8 = 128
        _ = v9
    } else {
        switch "\n\t" {
        case "hello":
            v3 += v3
        default:
            var v10 int32 = -2147483648
            _ = v10
        }
        v5 = append(v5, (v8 + v8))
        delete(v6, 769)
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v3 -= (v3 - -441)
            var v13 uint8 = 2
            v4 = true
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            v3 -= v3
            _ = fi14
        }
        _ = fi11
    }
    switch "a" {
    case "":
        if (!(v1 > (v1 - 0))) {
            if (42 > 100) {
                for fi15 := 0; fi15 < 18; fi15++ {
                    if (!(!(42 != 583))) {
                        var v16 float32 = 3.14159
                        if false {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                if len(v5) > 0 {
                                    v5[0] = "foo bar"
                                }
                                _ = fi17
                            }
                        } else {
                            v18 := GenS1{F0: 521, F1: 1}
                            _ = v18
                        }
                        _ = v16
                    } else {
                        var v19 uint16 = 41
                        _ = v19
                    }
                    var v20 float64 = -2.25
                    _ = fi15
                    _ = v20
                }
                v6[-706] = 127
            }
            var v21 float32 = 1.5
            v22 := GenS1{F0: 1, F1: 0}
            _ = v21
            _ = v22
        } else {
            if v4 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            if ((((3.14159 <= 1e10) && v4) || false) && false) {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v24 := GenS2{F0: 100, F1: -308}
                    v25 := GenS2{F0: 1, F1: 42}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
    case "\n\t":
        if ((1 == 10) || v4) {
            v26 := GenS3{F0: (^(v7)), F1: 42}
            if v4 {
                for fi27 := 0; fi27 < 16; fi27++ {
                    switch fi27 {
                    case 64:
                        v4 = (255 == 255)
                    case -159:
                        for fi28 := 0; fi28 < 8; fi28++ {
                            var v29 uint64 = (v1 ^ 18446744073709551615)
                            fmt.Println(v7)
                            _ = fi28
                            _ = v29
                        }
                    default:
                        var v30 int32 = 42
                        _ = v30
                    }
                    if (-2147483648 <= 1) {
                        var v31 int64 = v3
                        switch v8 {
                        case "\n\t":
                            fmt.Println(fi27, v31)
                        }
                        _ = v31
                    } else {
                        var v32 string = ("" + (((v8 + v8) + v8) + ""))
                        _ = v32
                    }
                    _ = fi27
                }
                for fi33 := 0; fi33 < 15; fi33++ {
                    v34 := GenS1{F0: 65535, F1: -92}
                    _ = fi33
                    _ = v34
                }
                delete(v6, -166)
            } else {
                var v35 bool = (!(!v4))
                var v36 uint64 = v1
                _ = v35
                _ = v36
            }
            _ = v26
        } else {
            for fi37 := 0; fi37 < 19; fi37++ {
                v38 := GenS1{F0: 2147483647, F1: 1000}
                switch (fi37 << ((0) & 63)) {
                case 128:
                    v2 = (v2 | v2)
                case 1:
                    var v39 int8 = 1
                    _ = v39
                }
                for fi40 := 0; fi40 < 15; fi40++ {
                    v41 := GenS4{F0: 0}
                    _ = fi40
                    _ = v41
                }
                _ = fi37
                _ = v38
            }
            var v42 uint64 = (v1 >> ((1) & 63))
            v8 = (v8 + ("hello" + "hello"))
            _ = v42
        }
    case "a":
        v4 = (v8 > v8)
    }
    v43 := GenS1{F0: 65535, F1: 0}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

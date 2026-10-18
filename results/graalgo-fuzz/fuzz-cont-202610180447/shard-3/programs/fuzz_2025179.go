package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }

func main() {
    var v1 int16 = 1000
    var v2 int16 = (v1 - (v1 + 1000))
    v3 := []uint32{2147483647, 1, 2}
    switch (("hello" + "a") + ("" + "foo bar")) {
    case "":
        v1 = v2
    case "\n\t":
        var v4 bool = ((^(v1)) >= (v1 + (v2 ^ v2)))
        _ = v4
    default:
        if (0 < 5) {
            v5 := GenS1{F0: 0, F1: -1}
            _ = v5
        } else {
            if true {
                for fi6 := 0; fi6 < 8; fi6++ {
                    switch (fi6 ^ fi6) {
                    case 2147483647:
                        v2 = v2
                    case 128:
                        for fi7 := 0; fi7 < 17; fi7++ {
                            if (18446744073709551615 <= 714) {
                                var v8 int64 = 42
                                var v9 uint16 = 65535
                                var v10 string = ("hello" + ((("a" + ("" + "\n\t")) + "hello") + "a"))
                                _ = v10
                                _ = v8
                                _ = v9
                            }
                            v1 += 42
                            if len(v3) > 0 {
                                v3[0] = 377
                            }
                            _ = fi7
                        }
                    default:
                        if (-2.25 > (1e10 + 3.14159)) {
                            var v11 int8 = 0
                            _ = v11
                        }
                    }
                    _ = fi6
                }
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(fi12)
                    _ = fi13
                }
                var v14 uint64 = 18446744073709551615
                _ = len(v3)
                _ = fi12
                _ = v14
            }
            var v15 int = 128
            _ = v15
        }
    }
    var v16 bool = ((42 < 430) || false)
    var v17 int32 = 187
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            v2 = (-(v1))
            _ = fi19
        }
        _ = fi18
    }
    var v20 uint32 = 2147483647
    var v21 uint = 4
    var v22 uint16 = 42
    v16 = ((0.0 / 1e10) < 0.0)
    _ = len(v3)
    v21 += (v21 + (v21 << (((v21 >> ((4294967295) & 63))) & 63)))
    var v23 int8 = 42
    var v24 int8 = v23
    var v25 int32 = 799
    if ((v1 << (((^(v21))) & 63)) == -1) {
        if v16 {
            v23 = v23
            var v26 int = -823
            _ = v26
        }
        v27 := GenS1{F0: 42, F1: int8(v21)}
        if v16 {
            switch ((((("a" + "") + "a") + "") + "a") + (((("" + "a") + "a") + "hello") + "foo bar")) {
            case "a":
                var v28 int = 32767
                _ = v28
            case "\n\t":
                v23 = (-(v23))
            }
            fmt.Println(v20, v2, v23)
        }
        _ = v27
    }
    fmt.Println(v20, v2, v16)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    switch "foo bar" {
    case "a":
        switch 0 {
        case -988:
            for fi29 := 0; fi29 < 20; fi29++ {
                switch ("a" + ("hello" + "hello")) {
                case "a":
                    _ = len(v3)
                default:
                    var v30 uint64 = 929
                    _ = v30
                }
                var v31 int64 = 439
                _ = fi29
                _ = v31
            }
        case -9223372036854775808:
            fmt.Println(v1, v24)
        default:
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 bool = true
                _ = fi32
                _ = v33
            }
        }
    default:
        if v16 {
            for fi34 := 0; fi34 < 13; fi34++ {
                v35 := GenS1{F0: 613, F1: -1}
                v36 := GenS2{F0: 220, F1: v21, F2: 660}
                _ = fi34
                _ = v35
                _ = v36
            }
        } else {
            v22 -= v22
            v37 := GenS2{F0: 191, F1: (^(v21)), F2: uint64(v21)}
            _ = v37
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}

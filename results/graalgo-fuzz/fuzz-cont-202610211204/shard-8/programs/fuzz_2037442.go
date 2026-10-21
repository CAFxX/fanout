package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint32, p1 int32, p2 int16) int16 {
    var v4 int64 = 362
    v5 := []int{837, 1, 9223372036854775807, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 int64) string {
    p0 += p0
    fmt.Println(p0)
    var v6 float64 = float64((0.1 + 1.5))
    _ = p0
    _ = v6
    return ("hello" + "a")
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = 2147483647
    v3 := []string{"foo bar", "hello"}
    fmt.Println(v2, v1)
    for fi7 := 0; fi7 < 7; fi7++ {
        _ = len(v3)
        _ = fi7
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v8 uint64 = 1
    if (true && (false || true)) {
        v3 = append(v3, "foo bar")
    } else {
        v8 -= 4294967295
        if false {
            switch "" {
            case "foo bar":
                var v9 uint64 = 2
                _ = v9
            case "a":
                var v10 int = (v2 - (v2 % ((v2) | 1)))
                _ = v10
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v11 uint = v1
            fmt.Println(v11, v2)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := GenS1{F0: 2147483647}
                    fmt.Println(fi13, v8)
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            v1 = (v1 / (((^(v1))) | 1))
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 string = ((("" + "\n\t") + "foo bar") + "hello")
                _ = fi15
                _ = v16
            }
        }
    }
    v1 = v1
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            _ = len(v3)
            fmt.Println(v2, v1)
            _ = fi18
        }
        if (("a" + ("hello" + "foo bar")) < "hello") {
            switch "a" {
            case "\n\t":
                if true {
                    if (!(32767 > genFn0(218, 1, 1))) {
                        if (!false) {
                            var v19 int64 = 55
                            _ = v19
                        }
                        switch ("\n\t" + ("\n\t" + "\n\t")) {
                        case "foo bar":
                            for fi20 := 0; fi20 < 8; fi20++ {
                                var v21 uint32 = 2147483647
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                    var v22 int8 = 77
                    _ = v22
                }
            }
            var v23 int = 9223372036854775807
            var v24 uint = v1
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: v8}
        _ = fi17
        _ = v25
    }
    v26 := GenS3{F0: 1, F1: 2}
    var v27 uint32 = 65535
    var v28 uint16 = 306
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v2 & (v2 & 268)) {
    case -400:
        var v29 uint32 = uint32(v28)
        _ = v29
    case 42:
        v1 = uint(v28)
    case -484:
        v8 -= uint64(v1)
    default:
        for fi30 := 0; fi30 < 15; fi30++ {
            v1 = 0
            _ = fi30
        }
    }
    switch v2 {
    case 42:
        if true {
            for fi31 := 0; fi31 < 19; fi31++ {
                v32 := GenS2{F0: uint64(v1)}
                _ = fi31
                _ = v32
            }
            v8 -= v8
        }
    }
    var v33 int64 = 42
    var v34 uint8 = 255
    var v35 int64 = 42
    v36 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v8
}

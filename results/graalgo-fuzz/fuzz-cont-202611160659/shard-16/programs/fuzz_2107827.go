package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 74
    var v3 bool = false
    v4 := []string{"\n\t"}
    if ((0 == 9223372036854775807) || v3) {
        v5 := GenS1{F0: -192, F1: 752, F2: 854}
        _ = v5
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 13; fi7++ {
                if (!(-2147483648 >= 479)) {
                    fmt.Println(v2, v2, fi6)
                    if (2147483647 >= -994) {
                        v3 = v3
                        v3 = ((1 >= 18446744073709551615) && (!(!((fi7 << ((v2) & 63)) >= (fi7 & fi6)))))
                        if v3 {
                            switch "hello" {
                            case "":
                                v2 = v2
                            default:
                                v8 := GenS2{F0: 0}
                                _ = v8
                            }
                        }
                        v9 := GenS3{F0: -829}
                        _ = v9
                    } else {
                        switch ("foo bar" + ("" + "hello")) {
                        case "a":
                            v2 = v2
                        }
                        if false {
                            v3 = (v3 && v3)
                        }
                        var v10 string = ""
                        _ = v10
                    }
                    v2 -= v2
                } else {
                    fmt.Println(fi6, fi6)
                    v2 = v2
                    delete(v1, "foo bar")
                }
                v11 := GenS2{F0: -2147483648}
                var v12 uint32 = 2147483647
                _ = fi7
                _ = v11
                _ = v12
            }
            v13 := GenS3{F0: 565}
            _ = fi6
            _ = v13
        }
        var v14 string = "hello"
        _ = v14
    }
    var v15 uint8 = 42
    var v16 uint16 = 2
    var v17 float64 = 1e10
    v17 = ((v17 + v17) - 1e10)
    if (42 < -2147483648) {
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint = 362
                var v21 uint64 = 18446744073709551615
                v16 = 428
                _ = fi19
                _ = v20
                _ = v21
            }
            v3 = (2147483647 > -104)
            v22 := GenS2{F0: 1}
            _ = fi18
            _ = v22
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v23 int64 = 0
    v23 -= (v23 % (((-(v23))) | 1))
    v16 -= v16
    v24 := v1["\n\t"]
    if v3 {
        fmt.Println(v3, v16, v24)
    } else {
        v25 := GenS2{F0: -2147483648}
        _ = v25
    }
    var v26 int8 = int8(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}

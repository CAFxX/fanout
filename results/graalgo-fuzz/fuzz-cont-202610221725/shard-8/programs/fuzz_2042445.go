package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8) int32 {
    if (!false) {
        v6 := []int{-1, 0, -1, -2147483648}
        switch -814 {
        case -214:
            if len(v6) > 0 {
                v6[0] = 0
            }
        }
        _ = v6
    } else {
        if (65535 <= 255) {
            var v7 float64 = 1e10
            _ = v7
        }
    }
    var v8 int16 = -1
    var v9 int16 = v8
    v9 = 3
    v10 := GenS1{F0: 353}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 339
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = -2.25
    var v3 int = 42
    v4 := []int{}
    v5 := []int{v3}
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            v5 = append(v5, v3)
            if false {
                if false {
                    var v12 int8 = -1
                    var v13 string = ("foo bar" + "hello")
                    _ = len(v4)
                    _ = v12
                    _ = v13
                }
            }
        }
        _ = fi11
    }
    v5 = append(v5, (-(v3)))
    v1 = 42
    switch ("" + "\n\t") {
    case "\n\t":
        switch "\n\t" {
        case "":
            var v14 bool = (!true)
            _ = v14
        case "foo bar":
            v1 = (v1 | v1)
        }
    case "hello":
        if true {
            v1 = v1
        } else {
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(fi15, v2, v2)
                if false {
                    v16 := GenS2{F0: 42}
                    for fi17 := 0; fi17 < 14; fi17++ {
                        switch 0 {
                        case 32768:
                            v5 = append(v5, 32768)
                        case -270:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v18 int16 = 32767
                            _ = v18
                        case 1:
                            v1 = v1
                        }
                        _ = len(v5)
                        var v19 uint16 = 65535
                        _ = fi17
                        _ = v19
                    }
                    if len(v5) > 0 {
                        v5[0] = (-(fi15))
                    }
                    _ = v16
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(fi15, v3, v2)
                    var v20 uint64 = 2
                    _ = v20
                }
                _ = fi15
            }
            if len(v5) > 0 {
                v5[0] = v3
            }
            fmt.Println(v3)
        }
    }
    var v21 float32 = 1e10
    var v22 int = 0
    if len(v5) > 0 {
        v5[0] = (^(v22))
    }
    if true {
        var v23 uint8 = v1
        fmt.Println(v21)
        _ = v23
    } else {
        v2 = 1e10
        for fi24 := 0; fi24 < 9; fi24++ {
            _ = len(v5)
            v3 = v3
            switch fi24 {
            case 9223372036854775807:
                if len(v5) > 0 {
                    v5[0] = fi24
                }
            case 265:
                var v25 uint16 = 2
                _ = v25
            case 421:
                v22 -= (fi24 - v3)
            }
            _ = fi24
        }
        if (genFn0((^(v1))) <= 2147483647) {
            _ = len(v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}

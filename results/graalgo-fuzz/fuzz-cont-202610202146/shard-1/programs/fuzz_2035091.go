package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0() bool {
    var v7 uint64 = 18446744073709551615
    v7 = (v7 >> ((65535) & 63))
    _ = v7
    return (!false)
}

func main() {
    v1 := []int{4294967295, 127}
    var v2 int64 = -9223372036854775808
    var v3 float64 = (0.1 - (-2.25 * 0.0))
    var v4 int64 = (v2 << ((961) & 63))
    v5 := make(map[uint32]bool)
    var v6 uint8 = 0
    var v8 int64 = v2
    switch ("hello" + ("a" + (("" + "foo bar") + ((("a" + "hello") + "foo bar") + "")))) {
    case "\n\t":
        for fi9 := 0; fi9 < 20; fi9++ {
            if true {
                switch fi9 {
                case 127:
                    v10 := GenS1{F0: -1}
                    _ = v10
                default:
                    v8 = (v2 << ((18446744073709551615) & 63))
                }
                v11 := GenS2{F0: 1, F1: 2147483648}
                fmt.Println(v3, v6)
                _ = v11
            }
            _ = fi9
        }
    }
    if (!true) {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = (fi12 % ((fi12) | 1))
            switch -82 {
            case -2147483648:
                var v14 float32 = (-2.25 - (0.1 - -2.25))
                _ = v14
            }
            fmt.Println(v6, v4)
            _ = fi12
            _ = v13
        }
        if true {
            if (4294967295 > 1) {
                if true {
                    var v15 int = -2147483648
                    _ = v15
                }
                v5[2] = true
                var v16 int16 = 1
                _ = v16
            } else {
                v17 := GenS1{F0: 32767}
                _ = v17
            }
            switch -9223372036854775808 {
            case 406:
                var v18 float64 = (v3 - v3)
                _ = v18
            case -1:
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v8)
                    v6 = (v6 / (((^(v6))) | 1))
                    switch "foo bar" {
                    case "foo bar":
                        if (-1 > 1) {
                            if true {
                                v2 = 0
                            }
                            var v20 bool = (((!(11 <= 127)) || false) || (559 == 1))
                            fmt.Println(v4, fi19)
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 11; fi21++ {
                                v8 = -1
                                fmt.Println(fi19, fi21, v8)
                                _ = fi21
                            }
                            var v22 bool = (18446744073709551615 >= 18446744073709551615)
                            _ = v22
                        }
                    case "a":
                        var v23 uint32 = 172
                        _ = v23
                    case "":
                        v3 += float64(1e10)
                    }
                    _ = fi19
                }
            default:
                v1 = append(v1, -1)
            }
            var v24 bool = (!true)
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint64 = 915
        _ = fi25
        _ = v26
    }
    if false {
        if (((true || true) && (!(!true))) || (false || true)) {
            var v27 uint16 = 65535
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v2 = (v4 << ((1) & 63))
        _ = len(v1)
        var v29 uint8 = 36
        _ = fi28
        _ = v29
    }
    var v30 float64 = 1.5
    if false {
        fmt.Println(v30, v30)
        fmt.Println(v3, v30, v2)
    } else {
        v8 = v2
    }
    var v31 bool = true
    for fi32 := 0; fi32 < 5; fi32++ {
        v8 = int64(v6)
        _ = len(v5)
        switch fi32 {
        case 9223372036854775807:
            v1 = append(v1, fi32)
        case -9223372036854775808:
            switch "\n\t" {
            case "\n\t":
                v33 := GenS2{F0: 289, F1: 227}
                _ = v33
            }
        case -2147483648:
            var v34 bool = true
            _ = v34
        default:
            var v35 int8 = 0
            _ = v35
        }
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

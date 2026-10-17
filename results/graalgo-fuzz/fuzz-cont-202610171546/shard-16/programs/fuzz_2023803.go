package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"a"}
    var v3 uint64 = 791
    var v4 int32 = -2147483648
    v5 := GenS1{F0: ("\n\t" + "\n\t")}
    fmt.Println(v4, v3)
    switch 32768 {
    case 127:
        v6 := GenS1{F0: "a"}
        _ = v6
    case -1:
        var v7 bool = false
        _ = v7
    }
    var v8 int = 785
    switch "a" {
    case "":
        var v9 int8 = 0
        _ = v9
    case "a":
        if ((!(true || (1 > -32768))) || false) {
            switch "" {
            case "foo bar":
                v3 = v3
            case "a":
                if (false && (2 != 2)) {
                    var v10 int32 = v4
                    var v11 uint16 = 670
                    _ = v10
                    _ = v11
                }
            default:
                var v12 uint16 = 65535
                _ = v12
            }
            v13 := GenS1{F0: ""}
            var v14 uint64 = (v3 % ((18446744073709551615) | 1))
            _ = v13
            _ = v14
        }
    }
    if ((!false) || false) {
        var v15 int64 = -1
        if true {
            switch "hello" {
            case "\n\t":
                var v16 int = v8
                _ = v16
            case "foo bar":
                v17 := GenS1{F0: ("" + "hello")}
                _ = v17
            case "":
                for fi18 := 0; fi18 < 8; fi18++ {
                    v19 := GenS1{F0: ("" + "hello")}
                    _ = fi18
                    _ = v19
                }
            }
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            var v20 uint64 = 18446744073709551615
            _ = v20
        } else {
            v15 += v15
        }
        v21 := v1[("" + "hello")]
        _ = v15
        _ = v21
    } else {
        if (((!true) || false) || false) {
            if false {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int8 = 107
                    _ = fi22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v2 = append(v2, "\n\t")
                    _ = fi24
                }
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v3)
                case "":
                    for fi25 := 0; fi25 < 18; fi25++ {
                        var v26 int8 = 1
                        _ = fi25
                        _ = v26
                    }
                default:
                    v3 += v3
                }
            }
            v3 = v3
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (4294967295 > (^(v8))) {
            v27 := GenS2{F0: 128, F1: (-(v4))}
            v28 := GenS2{F0: 0, F1: v4}
            for fi29 := 0; fi29 < 9; fi29++ {
                v8 -= v8
                _ = fi29
            }
            _ = v27
            _ = v28
        }
        if true {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v30 := GenS2{F0: 187, F1: v4}
            _ = v30
        }
    } else {
        delete(v1, "")
        v8 = 32768
        v4 += (v4 & -2147483648)
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        if (!(902 <= 65535)) {
            var v32 int64 = -904
            _ = v32
        }
        v33 := GenS1{F0: "hello"}
        _ = fi31
        _ = v33
    }
    if (false || true) {
        var v34 int = (v8 & (-(v8)))
        fmt.Println(v3, v34, v4)
        v35 := GenS2{F0: 1, F1: -1}
        _ = v34
        _ = v35
    }
    v3 = v3
    fmt.Println(v3)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

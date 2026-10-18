package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 uint16 = 0
    var v5 int = 4294967295
    if (446 <= 0) {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 822, F1: 42}
            _ = fi6
            _ = v7
        }
        v4 -= (v4 << ((2) & 63))
    } else {
        var v8 uint = 18446744073709551615
        _ = len(v1)
        v9 := GenS1{F0: -2147483648, F1: 36}
        _ = v8
        _ = v9
    }
    v4 = v4
    var v10 int = v5
    fmt.Println(v10, v5)
    if false {
        if len(v1) > 0 {
            v1[0] = ((("\n\t" + "\n\t") + "a") + "foo bar")
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 string = "\n\t"
            var v13 uint32 = 2147483647
            if (v4 >= (v4 << ((42) & 63))) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v13 = uint32(v4)
                    var v15 string = v12
                    _ = fi14
                    _ = v15
                }
                if (902 <= -2147483648) {
                    fmt.Println(fi11, v5, v12)
                } else {
                    if (true && (((v12 + "hello") + (v12 + (v12 + "a"))) != ("a" + v12))) {
                        var v16 bool = false
                        var v17 int32 = 935
                        _ = v16
                        _ = v17
                    }
                }
                var v18 int8 = 42
                _ = v18
            } else {
                fmt.Println(v4)
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        v4 = (v4 & (^(v4)))
    }
    v4 -= (v4 + v4)
    if false {
        switch (v5 << ((402) & 63)) {
        case 42:
            var v19 string = "\n\t"
            _ = v19
        case -515:
            var v20 uint16 = (^(v4))
            _ = v20
        default:
            switch 1 {
            case -2147483648:
                v21 := GenS1{F0: 2147483647, F1: 2}
                _ = v21
            case 128:
                if (1e10 < ((1.5 - 1.5) + 1e10)) {
                    v10 = (v10 | 0)
                } else {
                    if (255 > 65535) {
                        fmt.Println(v4, v10)
                        if ((v4 - 65535) > v4) {
                            var v22 int32 = 42
                            var v23 bool = (42 != 9223372036854775807)
                            _ = v22
                            _ = v23
                        }
                    }
                    v1 = append(v1, "\n\t")
                    var v24 uint32 = 988
                    _ = v24
                }
            default:
                v3["hello"] = (1 == 150)
            }
        }
    }
    v25 := GenS1{F0: -675, F1: 42}
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 int16 = 42
        _ = len(v2)
        v28 := GenS2{F0: (v27 % ((v27) | 1))}
        _ = fi26
        _ = v27
        _ = v28
    }
    delete(v2, (v5 + v5))
    switch ("\n\t" + "foo bar") {
    case "hello":
        v4 -= v4
    case "\n\t":
        v10 -= v5
    default:
        fmt.Println(v10)
    }
    fmt.Println(v10, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

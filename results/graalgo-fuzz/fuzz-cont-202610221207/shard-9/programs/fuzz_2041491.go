package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }

func main() {
    v1 := []string{}
    v2 := []uint32{4294967295, 2, 510}
    var v3 int = -1
    v4 := make(map[uint32]bool)
    var v5 float32 = 0.0
    var v6 uint32 = 0
    v7 := GenS1{F0: 25, F1: v6, F2: 113}
    v1 = append(v1, "foo bar")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 int64 = 761
    if (true || (546 > -1)) {
        switch ((("\n\t" + (("hello" + (("a" + ("foo bar" + (("a" + "hello") + "hello"))) + "a")) + "foo bar")) + ("hello" + "\n\t")) + "hello") {
        case "":
            var v9 int64 = v8
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 8; fi10++ {
                if true {
                    var v11 uint8 = 252
                    _ = v11
                }
                var v12 uint64 = 42
                _ = fi10
                _ = v12
            }
        case "a":
            for fi13 := 0; fi13 < 6; fi13++ {
                v5 = ((0.1 - v5) + v5)
                if false {
                    var v14 int64 = v8
                    var v15 bool = (0 == 42)
                    _ = v14
                    _ = v15
                }
                _ = fi13
            }
        }
        v5 += (float32(float64(0.0)) / 1e10)
        switch "a" {
        case "a":
            v16 := v4[(v6 / ((v6) | 1))]
            _ = v16
        case "foo bar":
            var v17 uint64 = 2
            _ = v17
        case "\n\t":
            _ = len(v2)
        }
    } else {
        if false {
            if (255 == 18446744073709551615) {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v6 = v6
                    _ = fi18
                }
                var v19 int = v3
                _ = v19
            }
            _ = len(v4)
        } else {
            if true {
                fmt.Println(v3, v3, v6)
                switch v3 {
                case -9223372036854775808:
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 uint16 = 1
                        v6 -= 787
                        _ = fi20
                        _ = v21
                    }
                case 671:
                    switch "" {
                    case "\n\t":
                        var v22 int = v3
                        _ = v22
                    }
                }
            }
        }
    }
    fmt.Println(v5, v3)
    v23 := GenS1{F0: -1, F1: 2, F2: (v8 | v8)}
    v6 -= v6
    v5 += 3.14159
    v24 := GenS1{F0: -1, F1: (v6 / ((857) | 1)), F2: (v8 << ((120) & 63))}
    var v25 int8 = 42
    v3 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

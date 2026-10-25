package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 230
    var v3 int = v2
    v4 := []uint32{505, 1, 797}
    var v5 int8 = 127
    if (629 != 2) {
        var v6 uint32 = 2147483647
        var v7 uint32 = v6
        _ = v6
        _ = v7
    }
    delete(v1, (v2 % (((v3 % (((v3 % ((32768) | 1))) | 1))) | 1)))
    v5 += v5
    v5 = v5
    v8 := GenS1{F0: 2147483648, F1: (-(v2)), F2: 42}
    if (127 < 1) {
        var v9 int64 = 0
        _ = v9
    }
    v10 := v1[v2]
    for fi11 := 0; fi11 < 7; fi11++ {
        switch ((v10 + (("foo bar" + ("hello" + ("hello" + v10))) + "\n\t")) + (v10 + "hello")) {
        case "foo bar":
            if (!true) {
                delete(v1, v2)
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 int32 = -164
                    _ = fi12
                    _ = v13
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = 42
                }
                if true {
                    v14 := GenS1{F0: 905, F1: 4294967295, F2: 427}
                    fmt.Println(v2, fi11, v2)
                    v3 = 842
                    _ = v14
                }
                _ = len(v1)
            }
        case "hello":
            _ = len(v4)
        case "\n\t":
            v2 -= (v2 / ((v3) | 1))
        }
        _ = fi11
    }
    if (!false) {
        switch v3 {
        case 32768:
            delete(v1, (v2 + (-(v2))))
        case 9223372036854775807:
            v2 -= v2
        case -678:
            v15 := GenS1{F0: 65535, F1: 2147483647, F2: 337}
            _ = v15
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint8 = 127
                switch v10 {
                case "":
                    v18 := GenS1{F0: 509, F1: (^(v2)), F2: 443}
                    _ = v18
                case "hello":
                    delete(v1, v3)
                }
                _ = fi16
                _ = v17
            }
        }
        v2 = v2
        switch (v10 + (v10 + "foo bar")) {
        case "\n\t":
            v4 = append(v4, 0)
        }
    } else {
        v19 := GenS1{F0: 600, F1: 32767, F2: -32768}
        v3 -= v2
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

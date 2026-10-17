package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    var v1 int32 = 0
    v2 := []string{}
    var v3 int32 = v1
    var v4 uint64 = 42
    v5 := []string{"\n\t"}
    fmt.Println(v1)
    switch 32768 {
    case 2147483647:
        switch "" {
        case "foo bar":
            v6 := GenS1{F0: (("" + "\n\t") + "\n\t"), F1: 256, F2: (v4 + v4)}
            _ = v6
        case "":
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint = 1
                fmt.Println(v1, v4)
                v8 -= v8
                _ = fi7
                _ = v8
            }
        }
    case 603:
        var v9 uint = 18446744073709551615
        _ = v9
    case 127:
        v1 += (v3 * v3)
    }
    v4 -= (v4 / (((v4 / (((v4 & v4)) | 1))) | 1))
    _ = len(v2)
    var v10 int64 = 0
    var v11 uint = 251
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v11)
        var v13 int = (fi12 * (-(fi12)))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        if (false || (0 != 1)) {
            var v15 uint = (v11 << (((^(v11))) & 63))
            switch (fi14 >> (((v15 * (^(v15)))) & 63)) {
            case 1:
                var v16 int8 = 96
                _ = v16
            }
            _ = v15
        } else {
            var v17 bool = (!((0.0 - 0.1) > (-2.25 - ((1e10 - (1e10 * float64(3.14159))) * (1.5 + 1.5)))))
            _ = v17
        }
        if ((866 <= 1) && true) {
            v11 = 4294967295
        } else {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            for fi18 := 0; fi18 < 9; fi18++ {
                v11 = (^(v11))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi18
            }
            var v19 uint8 = 1
            _ = v19
        }
        _ = fi14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v20 int = 128
    var v21 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

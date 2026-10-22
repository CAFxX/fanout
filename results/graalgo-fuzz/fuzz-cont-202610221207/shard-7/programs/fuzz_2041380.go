package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }

func main() {
    var v1 uint64 = 42
    v2 := []uint32{4294967295, 153, 158}
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    var v6 bool = (25 < 2)
    v1 += 42
    v6 = ((127 != 233) && false)
    fmt.Println(v6, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = v5
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v5 = uint(v1)
    v5 = (^(v5))
    switch 9223372036854775807 {
    case 128:
        delete(v3, "foo bar")
    case -775:
        var v7 uint32 = 859
        _ = v7
    case 399:
        var v8 int16 = 1000
        _ = v8
    default:
        v9 := GenS1{F0: (v5 / ((v5) | 1)), F1: "a", F2: "a"}
        _ = v9
    }
    v5 = v5
    var v10 bool = (1000 <= 32767)
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            if (!((v5 >> (((^(v5))) & 63)) > (v5 / ((1) | 1)))) {
                fmt.Println(fi12)
                var v13 float32 = (3.14159 + (3.14159 * -2.25))
                _ = v13
            } else {
                switch (fi12 - (fi12 * (fi11 % ((128) | 1)))) {
                case 2147483647:
                    delete(v3, "\n\t")
                }
            }
            if ((!(!v10)) && (((1e10 >= (0.1 * 1.5)) || true) && false)) {
                if (!v10) {
                    v2 = append(v2, 4294967295)
                    _ = len(v3)
                } else {
                    v14 := GenS1{F0: (v5 & 0), F1: "a", F2: ("\n\t" + "foo bar")}
                    var v15 float64 = -2.25
                    _ = v14
                    _ = v15
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    v6 = (v10 && (v6 || v6))
    v10 = v6
    var v16 int64 = -504
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

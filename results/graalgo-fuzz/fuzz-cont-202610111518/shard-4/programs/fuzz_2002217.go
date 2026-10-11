package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int = -1
    var v2 float32 = 3.14159
    v3 := make(map[int]bool)
    var v4 uint64 = 905
    v5 := []int{-798, v1, (v1 * 2147483647), v1}
    v1 = v1
    var v6 float64 = (((0.0 + 0.1) / 1e10) - 0.0)
    switch "\n\t" {
    case "foo bar":
        if (100 <= 127) {
            if false {
                if ((728 >= 9223372036854775807) && true) {
                    delete(v3, v1)
                } else {
                    v4 -= 4294967295
                    v2 -= 0.0
                    for fi7 := 0; fi7 < 19; fi7++ {
                        delete(v3, 32768)
                        v2 -= (v2 / float32((v6 - float64(v2))))
                        _ = fi7
                    }
                }
            } else {
                switch (v1 << ((639) & 63)) {
                case -9223372036854775808:
                    var v8 int32 = 1
                    _ = v8
                }
            }
        } else {
            if (613 == 220) {
                delete(v3, (v1 % ((-1) | 1)))
            }
            switch ("" + "\n\t") {
            case "foo bar":
                _ = len(v5)
            case "\n\t":
                delete(v3, 533)
            case "hello":
                if false {
                    v2 += 1e10
                } else {
                    var v9 bool = (true || true)
                    _ = v9
                }
            }
            var v10 bool = true
            _ = v10
        }
    case "":
        v6 = v6
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int32 = -2147483648
            var v13 uint64 = (v4 + (v4 - (v4 | 1)))
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: 1}
    v15 := GenS1{F0: 1}
    switch v1 {
    case -594:
        var v16 int = v1
        _ = v16
    case 4294967295:
        v6 = float64((float32(v6) / 1.5))
    }
    if (4294967295 != 2147483648) {
        v17 := v3[(v1 ^ v1)]
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

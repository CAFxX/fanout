package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 uint64, p2 uint32) uint16 {
    v6 := GenS1{F0: 177, F1: int32(p2)}
    v7 := GenS2{F0: 42}
    for fi8 := 0; fi8 < 12; fi8++ {
        p2 += p2
        _ = fi8
    }
    var v9 float32 = (3.14159 / (((0.1 + 0.0) + (float32(0.1) - ((0.0 / 0.1) + ((float32(float64((0.0 * -2.25))) + 3.14159) / (1.5 - ((float32(3.14159) / 3.14159) - (1e10 + 0.0))))))) / (1.5 / 1e10)))
    var v10 string = (("foo bar" + "\n\t") + ("\n\t" + ("foo bar" + ("" + ("a" + "foo bar")))))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v9
    return 255
}

func main() {
    v1 := []string{"a", "a"}
    var v2 float32 = 0.0
    var v3 bool = false
    var v4 int32 = -346
    var v5 uint16 = 32767
    for fi11 := 0; fi11 < 16; fi11++ {
        if (18446744073709551615 < 0) {
            _ = len(v1)
            fmt.Println(v5, v3)
        } else {
            v4 += v4
            v2 += v2
        }
        _ = fi11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint32 = 42
        for fi14 := 0; fi14 < 5; fi14++ {
            switch "" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                if v3 {
                    var v15 uint8 = 55
                    _ = v15
                }
            }
            v16 := GenS3{F0: "\n\t"}
            _ = fi14
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v17 := GenS3{F0: "\n\t"}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint8 = 255
        _ = fi18
        _ = v19
    }
    if ((!(((!(!(v3 || false))) && (18446744073709551615 != 274)) && (0 <= -32768))) || false) {
        v2 = float32(1e10)
    }
    fmt.Println(v5, v2, v5)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        if (1 <= 255) {
            var v20 int32 = (v4 ^ 2147483647)
            v5 += (v5 | genFn0(0, 4294967295, 1))
            _ = v20
        }
        v3 = v3
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        _ = len(v1)
        if (v3 || (true && ((42 <= 18446744073709551615) && v3))) {
            var v22 bool = (!((-(fi21)) == fi21))
            var v23 int64 = 9223372036854775807
            _ = v22
            _ = v23
        } else {
            v5 = v5
            v4 = 633
        }
        _ = fi21
    }
    fmt.Println(v3, v3)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

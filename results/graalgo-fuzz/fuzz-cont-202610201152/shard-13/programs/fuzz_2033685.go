package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0(p0 bool, p1 uint8) bool {
    v7 := []int{127, -2147483648, -9223372036854775808, 32767}
    v8 := GenS1{F0: 0, F1: 42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int = 42
    v2 := []string{("foo bar" + "\n\t"), "hello"}
    var v3 int64 = -9223372036854775808
    var v4 int16 = 153
    var v5 int64 = 9223372036854775807
    v6 := []uint32{0, 2, 65535, 42}
    switch v1 {
    case -1:
        fmt.Println(v5, v4, v5)
    default:
        var v9 int32 = 70
        _ = v9
    }
    v5 += (v5 << ((750) & 63))
    if len(v6) > 0 {
        v6[0] = 293
    }
    if true {
        if (v3 >= v3) {
            v10 := GenS1{F0: 65535, F1: 1}
            _ = v10
        }
        v5 = int64(v1)
    }
    fmt.Println(v5, v1, v5)
    var v11 float32 = 3.14159
    if (!(!true)) {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 uint16 = 395
            var v14 int8 = 110
            if (1.5 == 0.1) {
                v5 -= -9223372036854775808
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 float64 = float64(0.1)
        v2 = append(v2, "")
        _ = v15
    }
    switch (((("\n\t" + (("hello" + "hello") + "foo bar")) + "a") + ("hello" + "a")) + ("" + ("a" + (("hello" + "\n\t") + "a")))) {
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = 2
        }
    default:
        if false {
            switch "a" {
            case "hello":
                if true {
                    v1 += v1
                } else {
                    v1 = 128
                    if (v4 == (^(v4))) {
                        v16 := GenS1{F0: 0, F1: 0}
                        _ = v16
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(v11, v1, v4)
                        fmt.Println(v4, v11, v3)
                    }
                }
            case "foo bar":
                if false {
                    var v17 uint64 = 42
                    v17 += 42
                    _ = v17
                }
            case "\n\t":
                var v18 uint8 = 255
                _ = v18
            }
            var v19 uint8 = 65
            _ = v19
        }
    }
    v11 += v11
    v20 := GenS1{F0: 71, F1: 0}
    var v21 uint = 0
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
func genFn0() int {
    v5 := []int{0, 127}
    var v6 uint8 = 2
    var v7 int8 = 100
    if ((-696 >= -1) || (!true)) {
        v5 = append(v5, 127)
    }
    _ = v5
    _ = v6
    _ = v7
    return 32768
}

func main() {
    var v1 int16 = 659
    v2 := make(map[string]bool)
    var v3 int32 = -260
    var v4 uint32 = 65535
    v3 = 2147483647
    fmt.Println(v4)
    v8 := GenS1{F0: v4}
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(fi9, v3)
        _ = fi9
    }
    var v10 int8 = -10
    v11 := GenS2{F0: "foo bar", F1: v1, F2: (^(v3))}
    fmt.Println(v1, v3)
    if (-2147483648 >= 503) {
        delete(v2, (("foo bar" + ("" + "")) + ("\n\t" + ("a" + "a"))))
        var v12 int8 = 42
        if ((3.14159 != ((3.14159 / -2.25) - 1e10)) || (false && ((((!true) && ((false && true) || false)) || true) || false))) {
            var v13 int = 32767
            var v14 int32 = (v3 + 2147483647)
            var v15 int8 = v12
            _ = v13
            _ = v14
            _ = v15
        } else {
            if (true || (true && (1 <= int64(v3)))) {
                var v16 bool = true
                _ = v16
            } else {
                fmt.Println(v3, v10)
                for fi17 := 0; fi17 < 3; fi17++ {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v4 += (v4 * 42)
                        v19 := v2[("\n\t" + (("foo bar" + "foo bar") + "foo bar"))]
                        var v20 int64 = 42
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = fi17
                }
                var v21 int = 0
                _ = v21
            }
        }
        _ = v12
    } else {
        var v22 int8 = (v10 + (v10 | v10))
        _ = v22
    }
    v4 += 2147483648
    var v23 bool = true
    if (v23 && false) {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int64 = -9223372036854775808
            fmt.Println(v23, v1)
            _ = fi24
            _ = v25
        }
        v26 := v2[("foo bar" + "a")]
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v8
}

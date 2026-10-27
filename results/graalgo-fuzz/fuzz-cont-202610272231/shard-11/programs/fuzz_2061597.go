package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int64 }
func genFn0(p0 int64) uint32 {
    fmt.Println(p0)
    var v5 int64 = p0
    _ = p0
    _ = v5
    return 713
}

func main() {
    var v1 int = 128
    v2 := make(map[string]string)
    var v3 uint32 = 479
    v4 := []string{"hello", "foo bar", ("hello" + "\n\t"), "hello"}
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint = 42
            _ = fi6
            _ = v7
        }
    } else {
        fmt.Println(v1)
    }
    v8 := GenS1{F0: -978}
    v9 := GenS1{F0: -9223372036854775808}
    v10 := GenS2{F0: v1, F1: int(v3)}
    v11 := GenS1{F0: -9223372036854775808}
    if false {
        v1 = v1
    }
    if false {
        if len(v4) > 0 {
            v4[0] = (("\n\t" + "foo bar") + ((("a" + ("foo bar" + ("a" + "a"))) + ("" + "")) + "foo bar"))
        }
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    v3 += (v3 + v3)
    if (true || (!(("a" + "hello") < ""))) {
        fmt.Println(v3, v1)
        var v12 float32 = 3.14159
        fmt.Println(v1)
        _ = v12
    }
    v1 -= (-(v1))
    if false {
        v3 += v3
    }
    v1 -= v1
    for fi13 := 0; fi13 < 5; fi13++ {
        if ((576 >= 4294967295) && ((genFn0(-1) == genFn0(-9223372036854775808)) || false)) {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 string = "a"
                _ = fi14
                _ = v15
            }
        } else {
            if false {
                if (((false || false) && false) || false) {
                    var v16 int32 = -1
                    _ = v16
                } else {
                    v17 := GenS1{F0: 1}
                    _ = len(v2)
                    _ = v17
                }
                v18 := GenS2{F0: 128, F1: -826}
                _ = v18
            }
        }
        _ = len(v2)
        _ = fi13
    }
    if false {
        v2["foo bar"] = ""
        if (221 < 2147483647) {
            v19 := GenS2{F0: (v1 >> ((917) & 63)), F1: v1}
            v20 := GenS3{F0: ((((("foo bar" + "hello") + ("a" + "\n\t")) + ("hello" + "foo bar")) <= "hello") || ((42 < 18446744073709551615) && false)), F1: 618, F2: -32768}
            if true {
                var v21 uint = 0
                v22 := GenS4{F0: 144, F1: 0, F2: 347}
                _ = v21
                _ = v22
            }
            _ = v19
            _ = v20
        }
        if true {
            v23 := GenS2{F0: (v1 << ((65535) & 63)), F1: (v1 ^ v1)}
            fmt.Println(v3)
            _ = v23
        } else {
            if false {
                v24 := GenS1{F0: 9223372036854775807}
                _ = v24
            }
        }
    }
    var v25 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

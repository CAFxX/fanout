package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 uint64, p1 uint8, p2 uint32) int {
    p2 = uint32(p1)
    switch -446 {
    case -1:
        var v6 string = ("foo bar" + (((("foo bar" + "") + ("hello" + ("a" + "\n\t"))) + ("a" + "foo bar")) + "foo bar"))
        _ = v6
    case 4294967295:
        if true {
            p2 += p2
            p0 += (p0 ^ p0)
        }
    }
    fmt.Println(p0, p0, p2)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint16 = 65535
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 32767, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -841
}

func main() {
    var v1 float64 = (1e10 * (float64(3.14159) * ((0.1 * (((1e10 * 1e10) + float64(3.14159)) / float64((1.5 + 0.0)))) + (((1.5 + 0.1) / float64((1.5 * -2.25))) - float64(3.14159)))))
    var v2 int = 9223372036854775807
    var v3 uint16 = 32767
    var v4 string = "foo bar"
    v5 := make(map[int]int)
    var v10 uint64 = 4294967295
    var v11 float32 = 1.5
    v12 := GenS1{F0: (v3 / ((v3) | 1)), F1: 2147483647}
    v10 = (v10 >> ((18446744073709551615) & 63))
    v13 := GenS1{F0: (v3 + uint16(v10)), F1: -618}
    v4 = v4
    var v14 int64 = -9223372036854775808
    var v15 int = v2
    for fi16 := 0; fi16 < 16; fi16++ {
        if false {
            v14 = 42
            var v17 string = v4
            _ = v17
        } else {
            if (42 >= -994) {
                v18 := GenS1{F0: (v3 | (^(v3))), F1: int(v14)}
                v11 += 3.14159
                v3 -= 1
                _ = v18
            }
        }
        v10 = v10
        var v19 uint8 = 42
        _ = fi16
        _ = v19
    }
    if (v3 < (v3 << ((157) & 63))) {
        if ((!(42 >= -128)) || (-1 >= -2147483648)) {
            v20 := GenS1{F0: 42, F1: 2147483647}
            for fi21 := 0; fi21 < 6; fi21++ {
                v4 = "a"
                if (true && false) {
                    v22 := GenS1{F0: v3, F1: fi21}
                    _ = v22
                }
                _ = fi21
            }
            _ = v20
        }
        var v23 string = ((((v4 + ("a" + v4)) + v4) + ((v4 + ("foo bar" + ("\n\t" + v4))) + ("hello" + "foo bar"))) + v4)
        var v24 uint16 = uint16(v10)
        _ = v23
        _ = v24
    } else {
        var v25 uint32 = 2147483647
        if (true || true) {
            v1 -= v1
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := v5[(v15 ^ 4294967295)]
                var v28 bool = false
                for fi29 := 0; fi29 < 8; fi29++ {
                    v28 = v28
                    _ = fi29
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        _ = v25
    }
    var v30 uint32 = 407
    var v31 int8 = 42
    var v32 int32 = int32(v30)
    v32 = (v32 & 1)
    switch (-(v2)) {
    case 2147483647:
        delete(v5, (v15 + int(v31)))
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS2{F0: -32768, F1: "\n\t", F2: v4}
        _ = fi33
        _ = v34
    }
    var v35 string = ("\n\t" + ("foo bar" + v4))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}

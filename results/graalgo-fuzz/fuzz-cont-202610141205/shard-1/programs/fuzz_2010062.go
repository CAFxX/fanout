package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 bool, p1 int16) int16 {
    v7 := GenS1{F0: 32767, F1: "foo bar", F2: -106}
    p0 = false
    p1 = 0
    _ = p0
    _ = p1
    _ = v7
    return (p1 & 1)
}
func genFn1() bool {
    var v8 int = -260
    if (!(true || true)) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v8 += 127
            _ = fi9
        }
        v10 := []int{}
        _ = v10
    }
    switch v8 {
    case -2147483648:
        v8 = v8
    case 1:
        fmt.Println(v8)
    default:
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                fmt.Println(fi12, fi11)
                var v13 int16 = 0
                v13 = v13
                _ = fi12
                _ = v13
            }
            v8 += (v8 | (v8 + 2147483647))
            _ = fi11
        }
    }
    _ = v8
    return (1000 >= 1)
}

func main() {
    var v1 float64 = ((float64(0.0) - -2.25) / 3.14159)
    v2 := make(map[string]bool)
    var v3 int = 9223372036854775807
    var v4 int64 = 299
    var v5 bool = true
    var v6 int8 = -80
    v5 = ("" <= "")
    v4 -= -191
    v14 := GenS2{F0: 181, F1: 1, F2: "hello"}
    _ = len(v2)
    fmt.Println(v3, v6)
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint8 = 127
        v17 := GenS1{F0: fi15, F1: "hello", F2: (^(v6))}
        var v18 int64 = v4
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := []int{v3, (^(v3)), v3, 42}
    v3 -= v3
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS2{F0: 123, F1: 1, F2: "\n\t"}
            if ((v4 <= v4) || (1 >= 4294967295)) {
                var v23 int32 = 42
                _ = v23
            } else {
                v1 = 1e10
                _ = len(v19)
                if true {
                    v24 := v2["a"]
                    v3 = v3
                    _ = v24
                } else {
                    _ = len(v2)
                    v25 := GenS1{F0: (^(v3)), F1: (("a" + "a") + "hello"), F2: 100}
                    var v26 uint = 255
                    _ = v25
                    _ = v26
                }
                var v27 int64 = -254
                var v28 int64 = -9223372036854775808
                _ = v27
                _ = v28
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    fmt.Println(v4, v3, v3)
    v3 = 2147483647
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 bool = false
        _ = fi29
        _ = v30
    }
    var v31 int8 = v6
    v32 := GenS2{F0: 1, F1: 169, F2: ("a" + "foo bar")}
    v33 := GenS2{F0: 127, F1: 0, F2: (("a" + ("hello" + "a")) + ("a" + "foo bar"))}
    var v34 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

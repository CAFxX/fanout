package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int16 = -32768
    var v2 uint = 65535
    var v3 int8 = 1
    var v4 uint = 4294967295
    v5 := []int{-815, 42}
    if len(v5) > 0 {
        v5[0] = 127
    }
    fmt.Println(v2, v4, v2)
    if (18446744073709551615 != 248) {
        v2 = v2
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int32 = 146
        _ = fi6
        _ = v7
    }
    var v8 bool = (true || true)
    switch -740 {
    case 128:
        v2 -= (v2 + (v4 & 0))
    }
    v9 := GenS1{F0: v3, F1: 1}
    var v10 uint32 = 689
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: (v3 * (v3 % ((v3) | 1))), F1: fi11}
        if true {
            v13 := GenS2{F0: v1, F1: ""}
            v3 -= (-(v3))
            fmt.Println(v10, v4)
            _ = v13
        } else {
            if ((!(v8 || v8)) && (!(v2 > (^(v4))))) {
                var v14 uint32 = v10
                _ = v14
            } else {
                v10 = (v10 ^ (v10 >> ((v2) & 63)))
                v15 := GenS3{F0: 877, F1: 664}
                _ = v15
            }
            v16 := GenS4{F0: 561}
            if ((("foo bar" + "") > ("hello" + "foo bar")) && (int(v1) == 32768)) {
                v10 = v10
                if v8 {
                    _ = len(v5)
                    v2 -= v2
                    v3 = v3
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v17 string = "hello"
                var v18 int16 = v1
                _ = v17
                _ = v18
            }
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    v3 -= -85
    v19 := GenS3{F0: 2, F1: 2}
    var v20 bool = (1 < -404)
    var v21 int = 10
    v20 = v20
    v22 := GenS1{F0: v3, F1: v21}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

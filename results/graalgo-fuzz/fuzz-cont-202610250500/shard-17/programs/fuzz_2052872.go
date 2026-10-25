package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    var v1 bool = false
    var v2 int16 = 130
    var v3 int32 = 1
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: 42, F1: (!v1), F2: -2147483648}
        _ = fi4
        _ = v5
    }
    v6 := []int{32767, -214, 42, -934}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch -9223372036854775808 {
    case -1:
        v7 := GenS1{F0: 0, F1: v1, F2: v3}
        _ = v7
    case 801:
        if len(v6) > 0 {
            v6[0] = 32768
        }
    case -266:
        if ((!false) && true) {
            v1 = (0 >= 106)
        } else {
            var v8 uint8 = 128
            if (9223372036854775807 < 1) {
                var v9 float64 = (3.14159 + 0.0)
                _ = v9
            }
            if ((!v1) || false) {
                v10 := GenS1{F0: 42, F1: false, F2: (v3 * v3)}
                _ = v10
            }
            _ = v8
        }
    }
    v2 = v2
    fmt.Println(v1, v3)
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS2{F0: -1}
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    var v13 int8 = -1
    v6 = append(v6, 32767)
    var v14 int8 = (v13 | v13)
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 string = ("\n\t" + "foo bar")
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS3{F0: 18446744073709551615, F1: 255, F2: 615}
            v3 += 42
            var v19 int8 = v14
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    var v20 int16 = v2
    var v21 int8 = (-(v14))
    v22 := GenS2{F0: -9223372036854775808}
    for fi23 := 0; fi23 < 12; fi23++ {
        v1 = v1
        v20 += v20
        v20 = v2
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8; F1 int32; F2 int }

func main() {
    v1 := []int{-2147483648}
    var v2 uint = 1
    var v3 string = (((("" + ("foo bar" + "foo bar")) + (((("foo bar" + "hello") + "a") + "foo bar") + "foo bar")) + "hello") + "hello")
    var v4 int8 = -77
    var v5 bool = false
    var v6 string = ""
    var v7 uint8 = 1
    if false {
        for fi8 := 0; fi8 < 13; fi8++ {
            v2 = 619
            if ((52 < -288) && (fi8 > (fi8 / ((fi8) | 1)))) {
                v4 += v4
                v9 := GenS1{F0: 0, F1: 2147483647, F2: 1}
                var v10 int32 = int32(v2)
                _ = v10
                _ = v9
            } else {
                v11 := GenS1{F0: -2147483648, F1: -1, F2: -983}
                var v12 uint = (v2 - v2)
                _ = v11
                _ = v12
            }
            _ = fi8
        }
    } else {
        v13 := GenS1{F0: 868, F1: 1, F2: 9223372036854775807}
        _ = v13
    }
    var v14 int16 = 629
    for fi15 := 0; fi15 < 9; fi15++ {
        v3 = "foo bar"
        _ = fi15
    }
    if true {
        var v16 int = int(v4)
        for fi17 := 0; fi17 < 1; fi17++ {
            v6 = v6
            fmt.Println(v4)
            var v18 int16 = int16(v7)
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        var v19 int16 = (-(v14))
        var v20 string = v3
        _ = v19
        _ = v20
    }
    fmt.Println(v4, v3)
    var v21 int8 = v4
    v22 := GenS1{F0: -2147483648, F1: -2147483648, F2: -1}
    switch 42 {
    case 32767:
        v23 := GenS1{F0: -937, F1: -677, F2: 42}
        _ = v23
    case -660:
        if true {
            var v24 int = -644
            var v25 uint64 = 531
            v6 = "\n\t"
            _ = v24
            _ = v25
        }
    case -9223372036854775808:
        v26 := GenS1{F0: 0, F1: 2147483647, F2: -1}
        _ = v26
    }
    fmt.Println(v5, v21, v6)
    v27 := GenS1{F0: -294, F1: -70, F2: -1}
    v2 = (v2 ^ 65535)
    var v28 uint64 = 905
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 float32 = 1e10
        if (false || (!v5)) {
            var v31 string = v3
            switch ((((v6 + v31) + "\n\t") + ("foo bar" + "a")) + ((v31 + ("a" + "\n\t")) + v31)) {
            case "foo bar":
                for fi32 := 0; fi32 < 10; fi32++ {
                    v4 += v4
                    _ = fi32
                }
            case "\n\t":
                v3 = "\n\t"
            }
            var v33 float32 = float32(1.5)
            _ = v31
            _ = v33
        } else {
            v21 += 0
            var v34 int = fi29
            v35 := GenS2{F0: 1, F1: 2147483647, F2: (fi29 << (((v2 * (^(v2)))) & 63))}
            _ = v34
            _ = v35
        }
        var v36 float64 = (float64(0.0) * (0.1 / 1.5))
        _ = fi29
        _ = v30
        _ = v36
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        v5 = (v5 || ((v7 == (v7 << (((^(v2))) & 63))) || v5))
        _ = fi37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

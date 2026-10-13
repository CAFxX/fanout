package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 int16 }

func main() {
    var v1 int = 32767
    v2 := make(map[string]string)
    var v3 int64 = 1
    var v4 int8 = 42
    var v5 uint64 = 4294967295
    var v6 int = (v1 % (((v1 & int(v5))) | 1))
    v7 := GenS1{F0: true, F1: "\n\t", F2: 52}
    var v8 uint = 18446744073709551615
    v9 := []int{(v1 << ((v8) & 63)), v1}
    v8 += (v8 / ((v8) | 1))
    var v10 uint = v8
    switch (v1 % ((v6) | 1)) {
    case 127:
        for fi11 := 0; fi11 < 14; fi11++ {
            if true {
                var v12 uint64 = v5
                _ = v12
            }
            _ = fi11
        }
    case -366:
        switch 440 {
        case -9223372036854775808:
            for fi13 := 0; fi13 < 7; fi13++ {
                switch v6 {
                case 32768:
                    var v14 string = ""
                    _ = v14
                case 1:
                    var v15 uint32 = 2147483648
                    _ = v15
                }
                fmt.Println(v5, v10)
                _ = fi13
            }
        default:
            _ = len(v9)
        }
    default:
        v8 += (v10 / (((v8 >> (((^(v8))) & 63))) | 1))
    }
    v16 := GenS2{F0: 360, F1: v3, F2: 1000}
    if false {
        switch v6 {
        case 1:
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint8 = 206
                v6 = (-(v6))
                v18 += v18
                _ = fi17
                _ = v18
            }
        default:
            var v19 uint16 = 1
            _ = v19
        }
        v4 = -1
        var v20 int32 = 0
        _ = v20
    }
    v21 := GenS2{F0: 350, F1: (v3 % ((-521) | 1)), F2: -12}
    var v22 int8 = v4
    var v23 int64 = (-(v3))
    v4 = 42
    var v24 string = "\n\t"
    var v25 int32 = 1
    v8 += v8
    var v26 uint = v10
    if ((420 == 4294967295) || (0.1 < 3.14159)) {
        v25 -= (v25 - int32(v10))
        if false {
            if false {
                v6 = (v1 - (v6 / ((int(v23)) | 1)))
                v4 += 42
            }
            v25 -= v25
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                fmt.Println(v25, v5, v3)
                _ = fi27
            }
            for fi28 := 0; fi28 < 1; fi28++ {
                fmt.Println(v24, v22, fi28)
                v29 := GenS1{F0: (-824 > 1000), F1: v24, F2: 128}
                _ = fi28
                _ = v29
            }
            v30 := GenS2{F0: v1, F1: v23, F2: 257}
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 6; fi31++ {
            switch v24 {
            case "foo bar":
                var v32 int = (v1 * (fi31 & v1))
                _ = v32
            default:
                _ = len(v9)
            }
            var v33 float64 = 1e10
            _ = fi31
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

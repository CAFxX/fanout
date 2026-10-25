package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    var v4 int = -672
    v5 := make(map[uint32]bool)
    var v6 uint8 = 255
    v6 -= (^(v6))
    var v7 int8 = 0
    v8 := []int{}
    v9 := GenS1{F0: 2}
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            switch (v4 << ((368) & 63)) {
            case 1:
                if v1 {
                    v6 += v6
                    v12 := GenS2{F0: "a"}
                    _ = v12
                } else {
                    v6 -= v6
                }
            case -2147483648:
                if (32767 < 0) {
                    v13 := GenS2{F0: (("\n\t" + "") + "\n\t")}
                    var v14 uint16 = 255
                    _ = v13
                    _ = v14
                } else {
                    var v15 int64 = -9223372036854775808
                    fmt.Println(v4, fi11)
                    v16 := v5[749]
                    _ = v15
                    _ = v16
                }
            default:
                var v17 uint64 = 998
                _ = v17
            }
            v18 := GenS1{F0: 1}
            _ = fi11
            _ = v18
        }
        var v19 uint8 = 251
        var v20 bool = ((0 <= 42) && (v1 && v1))
        _ = fi10
        _ = v19
        _ = v20
    }
    fmt.Println(v7, v7)
    var v21 bool = (float32(1e10) != (1.5 + 1.5))
    var v22 int8 = (v7 & (v7 | 127))
    for fi23 := 0; fi23 < 5; fi23++ {
        switch v4 {
        case 2147483647:
            v24 := GenS2{F0: "foo bar"}
            _ = v24
        case 9223372036854775807:
            v25 := GenS1{F0: 18446744073709551615}
            _ = v25
        default:
            var v26 uint8 = v6
            _ = v26
        }
        v5[4294967295] = ((v1 || false) || false)
        for fi27 := 0; fi27 < 19; fi27++ {
            switch ("hello" + "hello") {
            case "\n\t":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch v4 {
                case 4294967295:
                    v4 += (fi23 >> ((0) & 63))
                case 128:
                    v5[42] = false
                }
            default:
                v21 = true
            }
            _ = fi27
        }
        _ = fi23
    }
    switch "a" {
    case "\n\t":
        if ((3.14159 * 3.14159) < -2.25) {
            var v28 int = v4
            _ = v28
        }
    case "hello":
        v21 = (!(v6 <= v6))
    case "":
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 bool = (v1 && (0.0 < 1.5))
            _ = fi29
            _ = v30
        }
    default:
        v1 = false
    }
    v22 += (v7 << ((65535) & 63))
    var v31 int64 = -9223372036854775808
    if false {
        var v32 float32 = (float32((float64(float32(float64(0.1))) - ((0.1 * (1.5 - 3.14159)) - (0.1 / 0.1)))) + ((1.5 * 3.14159) * (float32(((((1e10 - 0.1) / 1.5) / 1e10) + 0.1)) - 1.5)))
        v22 -= v22
        v31 -= v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

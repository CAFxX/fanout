package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }

func main() {
    var v1 int8 = 0
    var v2 int = -9223372036854775808
    var v3 uint64 = 4294967295
    if (405 >= 528) {
        v3 = (v3 & v3)
        var v4 int8 = (v1 ^ -13)
        _ = v4
    }
    fmt.Println(v2)
    fmt.Println(v1)
    var v5 uint8 = 127
    var v6 float32 = (0.1 / (-2.25 + ((1e10 - 0.1) / 0.1)))
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint32 = 312
        var v9 uint8 = 128
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint16 = 1
    if (32767 == -1) {
        fmt.Println(v5)
        switch (("" + "foo bar") + ("" + ((("" + "a") + "\n\t") + "foo bar"))) {
        case "a":
            var v11 uint32 = uint32(v5)
            _ = v11
        case "hello":
            switch "hello" {
            case "\n\t":
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := GenS1{F0: (v1 & (^(v1))), F1: 2147483647, F2: "hello"}
                    v6 += 0.0
                    _ = fi12
                    _ = v13
                }
            default:
                switch (v2 % (((v2 - (-(v2)))) | 1)) {
                case 866:
                    v14 := GenS1{F0: v1, F1: 0, F2: (("" + "a") + "\n\t")}
                    _ = v14
                default:
                    v3 = (v3 ^ (v3 + v3))
                }
            }
        case "\n\t":
            var v15 uint32 = 359
            _ = v15
        default:
            switch (v2 % (((v2 ^ v2)) | 1)) {
            case 32767:
                v16 := GenS1{F0: 25, F1: 4294967295, F2: ("hello" + "")}
                _ = v16
            case -235:
                if (42 <= -347) {
                    fmt.Println(v10, v6, v6)
                    fmt.Println(v1)
                } else {
                    v17 := []int{v2, v2, (v2 & v2)}
                    v3 = v3
                    _ = v17
                }
            case 42:
                for fi18 := 0; fi18 < 15; fi18++ {
                    if true {
                        v10 -= (v10 << ((338) & 63))
                        v3 -= (v3 & v3)
                        v19 := []int{(v2 & fi18), -209, fi18}
                        _ = v19
                    } else {
                        fmt.Println(v5)
                    }
                    fmt.Println(v3, v1)
                    _ = fi18
                }
            }
        }
    }
    v1 += (v1 % (((v1 << ((488) & 63))) | 1))
    var v20 uint8 = (v5 & (v5 - (^(v5))))
    if (0 <= v10) {
        v6 = v6
        var v21 float64 = ((float64(v6) * 0.0) * float64(v6))
        switch ("\n\t" + (("hello" + ("foo bar" + (("hello" + "hello") + ""))) + "\n\t")) {
        case "\n\t":
            var v22 int8 = v1
            _ = v22
        case "a":
            var v23 int16 = 32767
            _ = v23
        case "foo bar":
            v3 = uint64(v10)
        }
        _ = v21
    }
    if ((true || false) || (255 <= 127)) {
        for fi24 := 0; fi24 < 14; fi24++ {
            v3 = (v3 % ((v3) | 1))
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(v2, fi25, v2)
                _ = fi25
            }
            _ = fi24
        }
        v5 -= (v5 & 172)
    }
    fmt.Println(v5, v20, v6)
    v26 := []int{(^(v2)), 9223372036854775807, (v2 ^ 9223372036854775807), v2}
    for fi27 := 0; fi27 < 17; fi27++ {
        v26 = append(v26, 4294967295)
        switch (fi27 >> ((255) & 63)) {
        case -9223372036854775808:
            var v28 bool = (!(((^(v10)) <= (v10 | v10)) || (0.0 > (1.5 - 3.14159))))
            _ = v28
        }
        _ = len(v26)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v5
    _ = v6
}

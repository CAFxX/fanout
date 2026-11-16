package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 4294967295
    var v2 int64 = 1
    var v3 uint16 = 65535
    for fi4 := 0; fi4 < 19; fi4++ {
        if ((1e10 * 0.0) == 0.0) {
            if true {
                var v5 int8 = 0
                v3 -= v3
                if (((1 != 1) && true) && ((("\n\t" + "foo bar") + "\n\t") == "hello")) {
                    v6 := GenS1{F0: (fi4 + (fi4 / ((32767) | 1))), F1: (v5 >> ((v1) & 63)), F2: v2}
                    v3 += 1
                    _ = v6
                } else {
                    v2 = v2
                }
                _ = v5
            } else {
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 uint8 = 127
                    switch fi7 {
                    case 495:
                        var v9 int = 1
                        _ = v9
                    default:
                        var v10 int = int(v8)
                        _ = v10
                    }
                    v2 = (v2 / ((9223372036854775807) | 1))
                    _ = fi7
                    _ = v8
                }
                v11 := GenS1{F0: (^(fi4)), F1: -82, F2: v2}
                _ = v11
            }
        } else {
            v2 -= (v2 + v2)
            var v12 uint16 = (^(v3))
            _ = v12
        }
        v13 := GenS2{F0: 2}
        fmt.Println(v1)
        _ = fi4
        _ = v13
    }
    var v14 bool = (42 > 0)
    v1 += v1
    var v15 bool = (false && (v3 >= 2))
    v15 = ((("" + ("" + "foo bar")) + "\n\t") != "")
    var v16 uint64 = 18446744073709551615
    if (!true) {
        fmt.Println(v16)
        switch -2147483648 {
        case 128:
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := []int{}
                v14 = ((fi17 * (fi17 >> (((v1 >> ((v1) & 63))) & 63))) >= 979)
                _ = len(v18)
                _ = fi17
                _ = v18
            }
        }
    }
    v3 = (v3 / (((v3 ^ (v3 | (v3 + v3)))) | 1))
    if (v16 < 1) {
        if v14 {
            switch "" {
            case "hello":
                fmt.Println(v2, v14, v1)
            case "":
                switch ("foo bar" + (("a" + ((("a" + "a") + (("" + ("foo bar" + ("\n\t" + "\n\t"))) + "")) + (("foo bar" + "foo bar") + "foo bar"))) + "hello")) {
                case "a":
                    if v15 {
                        fmt.Println(v2, v1)
                    }
                }
            default:
                v3 += v3
            }
        }
        v1 -= v1
    } else {
        v1 = (v1 | v1)
        var v19 int16 = -1
        v20 := []int{32768, 32767}
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}

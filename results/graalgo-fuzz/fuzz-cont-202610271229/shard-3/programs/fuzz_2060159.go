package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]string)
    v3 := []int{1, -807, 2147483647}
    var v4 int8 = 56
    var v5 int64 = 724
    _ = len(v2)
    if true {
        if len(v3) > 0 {
            v3[0] = 128
        }
        v4 -= v4
    }
    v5 = 1
    if (!true) {
        var v6 float64 = float64(0.1)
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 1, F1: ("\n\t" + "")}
            for fi9 := 0; fi9 < 13; fi9++ {
                switch (fi9 | int(v1)) {
                case 32768:
                    var v10 float32 = 1.5
                    _ = v10
                case 42:
                    v4 -= -95
                case -75:
                    fmt.Println(fi9, v4, v5)
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        var v11 uint16 = 32767
        _ = v11
        _ = v6
    } else {
        if true {
            var v12 string = "\n\t"
            v4 += v4
            _ = v12
        } else {
            switch (("a" + "\n\t") + "\n\t") {
            case "foo bar":
                fmt.Println(v5)
            default:
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            }
        }
    }
    if true {
        switch 127 {
        case 32768:
            v13 := GenS1{F0: 65535, F1: (("\n\t" + "") + "")}
            _ = v13
        default:
            var v14 bool = (0.0 > (1e10 - -2.25))
            _ = v14
        }
        v15 := GenS1{F0: 2, F1: ((("" + "") + ("\n\t" + "")) + "")}
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 1; fi17++ {
                fmt.Println(v4, v1)
                _ = fi17
            }
            _ = fi16
        }
        _ = v15
    }
    var v18 int16 = 32767
    if ((1 > 127) && (1 > 0)) {
        var v19 uint64 = 18446744073709551615
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

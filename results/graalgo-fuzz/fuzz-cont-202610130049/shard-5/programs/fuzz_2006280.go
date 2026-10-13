package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 uint64, p1 uint16) uint16 {
    switch (("hello" + ("foo bar" + ("foo bar" + "\n\t"))) + "foo bar") {
    case "":
        var v6 string = "\n\t"
        _ = v6
    default:
        var v7 string = "a"
        _ = v7
    }
    var v8 uint16 = (^(p1))
    var v9 float64 = 3.14159
    var v10 int = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return uint16(p0)
}

func main() {
    var v1 bool = ((0 <= 163) && true)
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.0
    v4 := make(map[int]int)
    v5 := []int{433, 32768}
    if true {
        v11 := GenS1{F0: 42, F1: 1, F2: 65535}
        delete(v4, 32768)
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        switch (fi12 / ((fi12) | 1)) {
        case -150:
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(fi12)
                _ = fi13
            }
        case 1:
            v3 += (v3 * v3)
        case -1:
            switch "" {
            case "\n\t":
                _ = len(v5)
            default:
                if (true && v1) {
                    var v14 uint32 = 1
                    var v15 int = fi12
                    var v16 uint8 = 2
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
        }
        v1 = (!v1)
        _ = fi12
    }
    v1 = (0 <= 127)
    v3 -= float64(float32(v3))
    v1 = false
    fmt.Println(v1, v1)
    v1 = (true || v1)
    for fi17 := 0; fi17 < 18; fi17++ {
        if (798 <= 65535) {
            var v18 int64 = -1
            if false {
                var v19 bool = true
                _ = v19
            } else {
                v20 := GenS1{F0: 4294967295, F1: 1, F2: 2147483648}
                if ((fi17 << ((0) & 63)) > (-(fi17))) {
                    v21 := v2[65535]
                    _ = v21
                }
                _ = v20
            }
            _ = v18
        } else {
            var v22 uint64 = 676
            if ((v22 * 2) >= (v22 * 0)) {
                var v23 uint32 = 2
                switch fi17 {
                case -2147483648:
                    delete(v2, 42)
                case -986:
                    var v24 float32 = 1e10
                    _ = v24
                }
                switch fi17 {
                case 102:
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                    case "hello":
                        v23 = 2147483647
                    case "":
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v23 = (v23 & v23)
                            _ = fi25
                        }
                    }
                case -199:
                    v1 = (-1 == -9223372036854775808)
                case 9223372036854775807:
                    v26 := GenS1{F0: (v22 * 18446744073709551615), F1: 65535, F2: v23}
                    _ = v26
                }
                _ = v23
            } else {
                if (!(fi17 <= fi17)) {
                    v27 := GenS1{F0: 1, F1: 585, F2: 2147483647}
                    v2[582] = 17
                    _ = v27
                }
                var v28 uint64 = v22
                _ = v28
            }
            _ = v22
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
func genFn0() int16 {
    v6 := GenS1{F0: 42}
    var v7 int16 = 38
    _ = v6
    _ = v7
    return v7
}
func genFn1() uint {
    var v8 uint8 = 238
    for fi9 := 0; fi9 < 1; fi9++ {
        switch -1 {
        case -295:
            v10 := []int{(fi9 % ((fi9) | 1)), fi9, 128, (fi9 - (fi9 % (((fi9 / ((42) | 1))) | 1)))}
            _ = v10
        default:
            v8 -= 26
        }
        _ = fi9
    }
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (1 >= 480)
    v3 := []int{-1}
    v4 := make(map[string]string)
    var v5 int8 = 55
    if (v2 && (-244 >= 0)) {
        fmt.Println(v5)
    }
    _ = len(v1)
    if (!v2) {
        if v2 {
            v11 := GenS2{F0: 42}
            _ = v11
        } else {
            fmt.Println(v2, v2)
            fmt.Println(v5, v2)
            switch 9223372036854775807 {
            case -276:
                switch "a" {
                case "foo bar":
                    v2 = v2
                default:
                    for fi12 := 0; fi12 < 5; fi12++ {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v2 = v2
                            _ = fi13
                        }
                        v5 = 1
                        _ = fi12
                    }
                }
            case 127:
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 uint32 = 2147483648
                    _ = fi14
                    _ = v15
                }
            }
        }
        var v16 string = "a"
        _ = v16
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                if v2 {
                    var v19 float64 = (-2.25 + (-2.25 * -2.25))
                    var v20 int = (fi18 << ((18446744073709551615) & 63))
                    v2 = v2
                    _ = len(v3)
                    v20 = (v20 << ((1) & 63))
                    v5 = 1
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint16 = 402
                    switch "" {
                    case "a":
                        fmt.Println(fi17, fi18, fi18)
                    case "hello":
                        fmt.Println(v2)
                    }
                    _ = v21
                }
                _ = fi18
            }
            _ = fi17
        }
        v22 := v1[("a" + "foo bar")]
        _ = v22
    }
    fmt.Println(v2, v5)
    v5 -= -67
    v23 := GenS2{F0: 42}
    fmt.Println(v2, v2)
    var v24 float32 = 1.5
    fmt.Println(v5, v2)
    var v25 uint64 = 4294967295
    var v26 uint16 = 1
    var v27 uint64 = (v25 * (v25 % (((v25 * v25)) | 1)))
    fmt.Println(v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

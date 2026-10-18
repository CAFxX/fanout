package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 uint = 65535
    var v3 uint = (v2 % ((v2) | 1))
    var v4 int64 = -1
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    v1 = 0.0
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            switch fi7 {
            case 32768:
                v8 := v6[127]
                _ = v8
            case -793:
                if (42 == 65535) {
                    if true {
                        var v9 int32 = -1
                        v2 = (v3 - (v3 * (v2 - 255)))
                        var v10 int = (fi7 << (((v3 & (v3 ^ 255))) & 63))
                        _ = v10
                        _ = v9
                    }
                }
            case 708:
                var v11 uint16 = 255
                _ = v11
            }
            v4 += v4
            _ = fi7
        }
    }
    switch 4294967295 {
    case 68:
        switch ("foo bar" + (("" + ("hello" + "")) + "hello")) {
        case "hello":
            for fi12 := 0; fi12 < 20; fi12++ {
                v2 = (v2 / (((^(v2))) | 1))
                _ = fi12
            }
        case "a":
            var v13 uint64 = 2
            _ = v13
        }
    default:
        var v14 float32 = (0.1 * 0.1)
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v1 = 1.5
        switch (fi15 - (fi15 - fi15)) {
        case 864:
            v16 := []int{fi15, (fi15 | (fi15 << ((v2) & 63))), 127}
            _ = v16
        case 1:
            v1 -= (float64((3.14159 - 1.5)) * v1)
        }
        fmt.Println(fi15, v2)
        _ = fi15
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        if (!true) {
            fmt.Println(v2)
        } else {
            var v18 string = ("\n\t" + ("hello" + "foo bar"))
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint16 = 1
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        v1 -= v1
        for fi21 := 0; fi21 < 18; fi21++ {
            if true {
                if (false || true) {
                    if (!true) {
                        v22 := GenS1{F0: 18446744073709551615}
                        if ((4294967295 == 4294967295) && (7 > 858)) {
                            if false {
                                _ = len(v6)
                                var v23 int = fi17
                                var v24 int16 = 42
                                _ = v23
                                _ = v24
                            } else {
                                var v25 int8 = -11
                                v2 = (v2 ^ v2)
                                _ = v25
                            }
                            v4 = v4
                            if (681 < (v2 | (v2 % ((v2) | 1)))) {
                                fmt.Println(v3, v2, fi21)
                                var v26 uint16 = 563
                                _ = v26
                            } else {
                                v4 += (-(v4))
                            }
                            v27 := GenS1{F0: 164}
                            _ = v27
                        }
                        _ = v22
                    }
                    if (fi21 != -2147483648) {
                        var v28 int16 = 42
                        delete(v5, 12)
                        var v29 string = "\n\t"
                        v30 := []int{(fi17 & 127), (fi17 << (((^(v2))) & 63))}
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    v1 = v1
                }
            }
            _ = fi21
        }
        _ = fi17
    }
    var v31 int64 = (v4 | v4)
    var v32 int16 = -1
    fmt.Println(v32, v31, v32)
    for fi33 := 0; fi33 < 3; fi33++ {
        fmt.Println(v2, v32, v3)
        if ((!(fi33 >= 127)) && true) {
            _ = len(v6)
            v34 := GenS2{F0: "foo bar", F1: -1}
            v1 = v1
            _ = v34
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

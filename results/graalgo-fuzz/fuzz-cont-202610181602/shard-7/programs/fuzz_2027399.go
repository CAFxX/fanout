package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
func genFn0(p0 int, p1 int8) uint32 {
    switch ("" + ("foo bar" + (((("" + "foo bar") + ("" + "a")) + "foo bar") + ("" + "hello")))) {
    case "":
        p0 -= (p0 / ((p0) | 1))
    case "a":
        fmt.Println(p0, p1)
    case "foo bar":
        if true {
            var v5 uint64 = 1
            v6 := GenS1{F0: p1, F1: 255, F2: 18446744073709551615}
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 float32 = (((float32(3.14159) * -2.25) * (0.0 + 3.14159)) / (-2.25 / float32((((3.14159 - 1e10) + 0.1) / ((1e10 + 1.5) - float64(3.14159))))))
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        } else {
            var v9 int = 128
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS2{F0: 108, F1: false, F2: 2147483648}
        v12 := GenS1{F0: 0, F1: 87, F2: 0}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint = 18446744073709551615
    switch "hello" {
    case "hello":
        var v14 uint32 = 1
        _ = v14
    case "a":
        if true {
            for fi15 := 0; fi15 < 11; fi15++ {
                fmt.Println(p1, v13, fi15)
                _ = fi15
            }
            if true {
                for fi16 := 0; fi16 < 14; fi16++ {
                    if true {
                        p0 = (p0 / ((fi16) | 1))
                        v13 = v13
                        for fi17 := 0; fi17 < 7; fi17++ {
                            switch "hello" {
                            case "a":
                                var v18 uint32 = 2
                                _ = v18
                            case "foo bar":
                                var v19 uint16 = 65535
                                _ = v19
                            }
                            p1 -= (p1 & p1)
                            if true {
                                v20 := []int{}
                                if len(v20) > 0 {
                                    v20[0] = 9223372036854775807
                                }
                                p1 = -1
                                _ = v20
                            }
                            var v21 float32 = ((3.14159 - (0.1 + float32(3.14159))) + float32((float64((0.0 - (3.14159 - -2.25))) / (1e10 - 0.0))))
                            _ = fi17
                            _ = v21
                        }
                    }
                    p1 = (p1 << (((^(v13))) & 63))
                    _ = fi16
                }
                v13 += (v13 + (^(v13)))
            }
        }
    default:
        fmt.Println(v13)
    }
    _ = p0
    _ = p1
    _ = v13
    return 620
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 int8 = 100
    var v4 bool = (!false)
    switch ((((("\n\t" + "hello") + "\n\t") + "") + "foo bar") + "a") {
    case "":
        v4 = ((true || true) || v4)
    case "hello":
        switch "hello" {
        case "hello":
            fmt.Println(v3)
        }
    case "a":
        var v22 string = (("" + (("hello" + "foo bar") + "foo bar")) + ("foo bar" + ("" + "")))
        _ = v22
    }
    var v23 uint32 = 651
    for fi24 := 0; fi24 < 15; fi24++ {
        v3 = v3
        var v25 uint16 = 2
        v26 := []int{(fi24 - (fi24 ^ 2147483647)), fi24, (fi24 % (((fi24 | fi24)) | 1)), (^(fi24))}
        _ = fi24
        _ = v25
        _ = v26
    }
    v2["foo bar"] = 128
    var v27 int = -2147483648
    var v28 uint8 = 255
    for fi29 := 0; fi29 < 14; fi29++ {
        if (42 < -2147483648) {
            switch (fi29 ^ fi29) {
            case 127:
                for fi30 := 0; fi30 < 10; fi30++ {
                    _ = len(v1)
                    if v4 {
                        var v31 int8 = v3
                        _ = v31
                    }
                    v32 := GenS1{F0: v3, F1: v28, F2: 65535}
                    switch (fi30 ^ (v27 + (fi30 | -666))) {
                    case -2147483648:
                        v23 += genFn0(fi29, (v3 * (-(v3))))
                    default:
                        var v33 bool = (1 < 65535)
                        _ = v33
                    }
                    var v34 int32 = -2147483648
                    _ = fi30
                    _ = v32
                    _ = v34
                }
            case 1:
                switch "" {
                case "foo bar":
                    var v35 uint64 = 1
                    _ = v35
                }
            }
            if (v23 != (v23 ^ (v23 % ((v23) | 1)))) {
                v28 = 150
                v4 = (2147483647 == 42)
                v36 := GenS2{F0: 363, F1: (4294967295 != 597), F2: 0}
                _ = v36
            }
        } else {
            v37 := v2["hello"]
            _ = v37
        }
        v38 := []int{}
        if ((fi29 > fi29) && v4) {
            v27 = (fi29 % (((v27 / ((-9223372036854775808) | 1))) | 1))
        } else {
            var v39 string = "foo bar"
            _ = v39
        }
        _ = fi29
        _ = v38
    }
    fmt.Println(v28, v4, v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}

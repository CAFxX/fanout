package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint64; F1 uint; F2 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 float32 = (((1e10 + -2.25) - (3.14159 - 1e10)) + -2.25)
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint64 = 1
        v4 = float32(float64(float32(1e10)))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if (!(577 == 745)) {
            var v8 bool = (-1 <= 0)
            if (1 >= 2147483648) {
                var v9 int16 = 1
                _ = v9
            }
            v8 = (-9223372036854775808 <= fi7)
            _ = v8
        } else {
            v10 := GenS1{F0: 65535, F1: ((("\n\t" + ("a" + ((("\n\t" + "") + ((("a" + "") + "a") + "a")) + "a"))) + "\n\t") + "a")}
            v11 := GenS1{F0: 255, F1: "foo bar"}
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    var v12 string = "a"
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int = 1
        v16 := GenS1{F0: 2, F1: "foo bar"}
        for fi17 := 0; fi17 < 6; fi17++ {
            v15 = fi17
            _ = fi17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v18 := GenS1{F0: 319, F1: v12}
    fmt.Println(v4)
    switch "a" {
    case "\n\t":
        var v19 uint32 = 4294967295
        _ = v19
    case "":
        v4 -= 1e10
    case "foo bar":
        _ = len(v2)
    }
    switch 32768 {
    case 430:
        if false {
            v20 := GenS2{F0: 91, F1: -1, F2: 1}
            var v21 float64 = 3.14159
            if (-1 > 4294967295) {
                var v22 int16 = -32768
                _ = len(v2)
                _ = v22
            }
            _ = v20
            _ = v21
        }
    case -2147483648:
        v23 := GenS3{F0: 4294967295, F1: 32768}
        _ = v23
    case 128:
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS4{F0: 18446744073709551615, F1: 42, F2: (4294967295 == 229)}
            if true {
                v26 := GenS2{F0: int8(fi24), F1: -306, F2: 2}
                if (int16(v13) != -1) {
                    delete(v3, 2147483648)
                    v27 := v2[(fi24 & (fi24 - (fi24 | -1)))]
                    _ = v27
                } else {
                    var v28 int32 = (v13 & (v13 / ((-1) | 1)))
                    if true {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            if ((255 != 42) || (0 <= 764)) {
                                var v30 int64 = int64(v28)
                                v12 = (((v12 + v12) + (v12 + (v12 + (v12 + ("foo bar" + "foo bar"))))) + "foo bar")
                                delete(v2, (^(fi29)))
                                _ = v30
                            }
                            _ = fi29
                        }
                    } else {
                        switch (v12 + (v12 + "foo bar")) {
                        case "foo bar":
                            v31 := GenS2{F0: -128, F1: v13, F2: 1}
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 8; fi32++ {
                            v33 := v3[645]
                            _ = fi32
                            _ = v33
                        }
                    }
                    v34 := GenS2{F0: -128, F1: v28, F2: 65535}
                    _ = v28
                    _ = v34
                }
                fmt.Println(fi24, v12)
                _ = v26
            } else {
                var v35 uint8 = 214
                for fi36 := 0; fi36 < 5; fi36++ {
                    v35 = (v35 + 0)
                    _ = fi36
                }
                _ = v35
            }
            _ = fi24
            _ = v25
        }
    }
    v13 = (v13 >> ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

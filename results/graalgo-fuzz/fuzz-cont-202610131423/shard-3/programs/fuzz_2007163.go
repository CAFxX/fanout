package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint8 }
func genFn0() int32 {
    var v6 string = ((("" + ("foo bar" + "\n\t")) + "hello") + "hello")
    v7 := GenS1{F0: -623, F1: 0, F2: v6}
    v8 := []int{-571, 32768, -268, 9223372036854775807}
    _ = v6
    _ = v7
    _ = v8
    return -462
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.0
    v3 := []int{1, 9223372036854775807, 32768, -2147483648}
    v4 := make(map[string]bool)
    var v5 int64 = -9223372036854775808
    var v9 uint = 42
    switch ("\n\t" + (("a" + "a") + "\n\t")) {
    case "foo bar":
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10, v5)
            switch (^(fi10)) {
            case -1:
                if ((200 < 65535) || false) {
                    if true {
                        v5 -= (v5 ^ 9223372036854775807)
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v2 += 1e10
                            v12 := GenS1{F0: (fi11 << (((v9 * v9)) & 63)), F1: 65535, F2: "\n\t"}
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        v9 = 42
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v14 := GenS1{F0: 1, F1: 65535, F2: "hello"}
                            _ = fi13
                            _ = v14
                        }
                        if true {
                            var v15 uint = v9
                            _ = v15
                        } else {
                            v9 = (v9 - (^(v9)))
                            var v16 uint8 = 42
                            _ = v16
                        }
                    }
                }
            default:
                if false {
                    v2 += 1e10
                    _ = len(v4)
                    var v17 float64 = ((1.5 * (1.5 / 3.14159)) / 0.1)
                    _ = v17
                } else {
                    v18 := GenS2{F0: 220}
                    var v19 int16 = -366
                    _ = v18
                    _ = v19
                }
            }
            v5 = -1
            _ = fi10
        }
    default:
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS2{F0: 128}
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    v5 += (v5 - 1)
    fmt.Println(v9)
    switch "foo bar" {
    case "foo bar":
        switch 1 {
        case 32767:
            if (2 == 42) {
                if (-2147483648 < 0) {
                    v2 += v2
                    var v24 string = "\n\t"
                    _ = v24
                }
                for fi25 := 0; fi25 < 16; fi25++ {
                    switch "hello" {
                    case "":
                        var v26 int = 465
                        _ = v26
                    }
                    _ = fi25
                }
            }
        }
    case "":
        var v27 uint16 = 546
        _ = v27
    case "a":
        fmt.Println(v9)
    default:
        v9 += (v9 >> ((v9) & 63))
    }
    var v28 uint32 = 2147483647
    if false {
        if (true && (!(1 > (^(v28))))) {
            for fi29 := 0; fi29 < 20; fi29++ {
                for fi30 := 0; fi30 < 18; fi30++ {
                    switch fi29 {
                    case 2147483647:
                        v28 -= 65535
                    case 1:
                        v5 = 42
                        var v31 float32 = (float32(((1.5 + (0.1 / (((((((3.14159 - (float64(-2.25) * (-2.25 / (float64((float32(float64(v2)) + v2)) * ((-2.25 * 3.14159) - (1.5 * (float64(v2) - -2.25))))))) + 3.14159) - -2.25) / ((float64(v2) / float64(((float32(3.14159) * (v2 + (v2 * v2))) / (v2 + v2)))) * 3.14159)) * -2.25) / 0.1) + 3.14159))) + 3.14159)) + (1e10 * (0.0 - 1e10)))
                        var v32 uint8 = 23
                        _ = v31
                        _ = v32
                    }
                    _ = fi30
                }
                v33 := GenS3{F0: (v28 >> (((^(v9))) & 63)), F1: 183}
                _ = fi29
                _ = v33
            }
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        }
        v34 := GenS2{F0: 1}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 18; fi35++ {
            if len(v3) > 0 {
                v3[0] = (fi35 - (fi35 >> ((v9) & 63)))
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 uint32 = (v28 / ((v28) | 1))
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        v38 := GenS4{F0: 128}
        _ = v38
    }
    v39 := GenS1{F0: 128, F1: v28, F2: "foo bar"}
    v28 += (v28 % ((v28) | 1))
    switch ("" + "\n\t") {
    case "":
        var v40 uint32 = v28
        _ = v40
    }
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}

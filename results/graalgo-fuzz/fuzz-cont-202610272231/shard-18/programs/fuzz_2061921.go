package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 int; F1 int16; F2 uint16 }
func genFn0() uint8 {
    switch -9223372036854775808 {
    case 969:
        v4 := []int{}
        _ = v4
    case -2147483648:
        if (-43 >= 1) {
            if (-6 == 127) {
                if (18446744073709551615 < 18446744073709551615) {
                    for fi5 := 0; fi5 < 7; fi5++ {
                        for fi6 := 0; fi6 < 12; fi6++ {
                            if (929 != 822) {
                                fmt.Println(fi6, fi6)
                                fmt.Println(fi6, fi5)
                                if (!(683 <= 1)) {
                                    var v7 int64 = -914
                                    _ = v7
                                }
                            } else {
                                v8 := GenS1{F0: (("" + "") + "foo bar"), F1: 224, F2: 0}
                                var v9 string = ((("" + "hello") + ("foo bar" + "\n\t")) + "\n\t")
                                var v10 float64 = float64((-2.25 + (-2.25 - ((0.0 * float32(1.5)) + (float32(3.14159) - 1.5)))))
                                if true {
                                    fmt.Println(v9, v10, fi6)
                                    v11 := GenS1{F0: "hello", F1: 4294967295, F2: 42}
                                    v12 := GenS1{F0: v9, F1: 1, F2: 9223372036854775807}
                                    _ = v11
                                    _ = v12
                                }
                                _ = v10
                                _ = v8
                                _ = v9
                            }
                            _ = fi6
                        }
                        fmt.Println(fi5)
                        v13 := []int{}
                        _ = fi5
                        _ = v13
                    }
                } else {
                    var v14 uint64 = 909
                    _ = v14
                }
                var v15 bool = (false || true)
                var v16 int16 = 0
                _ = v15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := []int{32768, fi17, (^(fi17)), fi17}
                _ = fi17
                _ = v18
            }
        }
    default:
        var v19 int = 1
        _ = v19
    }
    v20 := []int{}
    _ = v20
    return 237
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 0
    v3 := make(map[int]uint32)
    var v21 uint32 = 742
    v2 = v2
    v22 := GenS1{F0: ("\n\t" + "foo bar"), F1: v2, F2: -396}
    v23 := GenS2{F0: 1, F1: -53, F2: -2147483648}
    v24 := GenS3{F0: -2147483648, F1: 32767, F2: 65535}
    switch (("\n\t" + ("foo bar" + (("hello" + "hello") + "\n\t"))) + "hello") {
    case "\n\t":
        if true {
            var v25 int8 = 110
            v25 = (v25 * v25)
            _ = v25
        }
    }
    v2 = (v2 % (((v2 & (v2 + 18446744073709551615))) | 1))
    var v26 uint32 = (v21 % ((v21) | 1))
    var v27 int = 1
    var v28 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}

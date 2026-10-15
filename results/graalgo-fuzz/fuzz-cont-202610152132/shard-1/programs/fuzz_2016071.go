package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 uint8; F1 string; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 int16 = 1
    v4 := make(map[int]uint32)
    var v5 uint = 471
    v6 := make(map[string]int)
    fmt.Println(v3, v3)
    if false {
        if (!(!(0 <= 2147483648))) {
            var v7 int64 = -9223372036854775808
            _ = v7
        }
        var v8 int8 = 100
        _ = v8
    } else {
        v2[4294967295] = "\n\t"
    }
    if ((v3 > (v3 / (((v3 << ((373) & 63))) | 1))) || true) {
        if true {
            v5 += (v5 | (v5 % (((^(v5))) | 1)))
            if (!(741 < -2147483648)) {
                v9 := GenS1{F0: 2147483648, F1: -9223372036854775808, F2: 941}
                v10 := GenS1{F0: 1, F1: -9223372036854775808, F2: (v5 | (v5 ^ v5))}
                var v11 float32 = 3.14159
                _ = v10
                _ = v11
                _ = v9
            } else {
                switch (("foo bar" + ("\n\t" + "")) + (("hello" + "\n\t") + "a")) {
                case "a":
                    fmt.Println(v3)
                case "\n\t":
                    v12 := []int{32767, 42, 2147483647, 32767}
                    _ = v12
                case "foo bar":
                    var v13 int32 = 291
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        if true {
                            var v15 uint8 = 42
                            _ = v15
                        }
                        v16 := GenS1{F0: 4294967295, F1: -390, F2: v5}
                        _ = fi14
                        _ = v16
                    }
                }
                if true {
                    v3 += v3
                }
            }
            fmt.Println(v5, v5)
        } else {
            fmt.Println(v5)
            for fi17 := 0; fi17 < 17; fi17++ {
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 float64 = (3.14159 + (0.1 - 1e10))
                    v3 = v3
                    v20 := []int{fi18, (fi18 << ((v5) & 63)), -512, int(v3)}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch ("\n\t" + "") {
                case "hello":
                    v21 := GenS2{F0: 1, F1: (("foo bar" + "") + ""), F2: 1}
                    _ = v21
                case "":
                    var v22 uint32 = 65535
                    _ = v22
                case "foo bar":
                    v23 := GenS3{F0: 1, F1: 2, F2: 42}
                    _ = v23
                default:
                    v5 -= 0
                }
                v24 := []int{}
                _ = fi17
                _ = v24
            }
            fmt.Println(v5, v3)
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            if ((3.14159 < (float64(3.14159) * ((0.0 + (0.1 - 0.0)) + (1e10 - float64(0.0))))) || (!(-9223372036854775808 <= (fi25 ^ 2147483647)))) {
                var v26 int32 = -2147483648
                _ = v26
            }
            _ = fi25
        }
    }
    v3 = v3
    fmt.Println(v3, v3)
    v27 := GenS3{F0: 818, F1: (v5 >> ((4294967295) & 63)), F2: 18446744073709551615}
    var v28 int32 = 42
    v3 = 1000
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int16 = v3
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

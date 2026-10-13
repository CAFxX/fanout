package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0(p0 int8) int {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: -161, F1: 42}
        if (1 > 18446744073709551615) {
            if true {
                if ((0.1 / (3.14159 - 0.1)) == (0.1 / ((1.5 / (1e10 + 0.1)) * -2.25))) {
                    v6 := []int{fi4, (fi4 ^ (fi4 << ((18446744073709551615) & 63))), (^(fi4))}
                    _ = v6
                } else {
                    v7 := GenS1{F0: fi4, F1: 127}
                    var v8 int = (fi4 << ((42) & 63))
                    var v9 int8 = p0
                    _ = v7
                    _ = v8
                    _ = v9
                }
                var v10 uint8 = 255
                p0 = (^(p0))
                _ = v10
            } else {
                p0 -= p0
                p0 = 127
                switch ("hello" + ("hello" + "a")) {
                case "":
                    if true {
                        p0 += (p0 - (^(p0)))
                    }
                default:
                    p0 = -1
                }
            }
            p0 = (p0 & 91)
            var v11 int64 = 0
            _ = v11
        }
        _ = fi4
        _ = v5
    }
    var v12 uint = 255
    fmt.Println(p0, p0)
    _ = p0
    _ = v12
    return -9223372036854775808
}
func genFn1() uint {
    if true {
        v13 := []int{1, 101, 616, 2147483647}
        var v14 string = ""
        if (0 >= 0) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if true {
                var v15 int8 = 0
                if len(v13) > 0 {
                    v13[0] = 1
                }
                _ = v15
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 uint64 = 2
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
        _ = v13
        _ = v14
    } else {
        if false {
            if true {
                var v19 string = ""
                _ = v19
            }
            var v20 uint = 4294967295
            v20 += (v20 - (v20 | (v20 << ((980) & 63))))
            _ = v20
        }
    }
    var v21 int16 = -1
    _ = v21
    return 65535
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{}
    v3 := make(map[int]bool)
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(v1, v1)
        _ = fi22
    }
    var v23 float32 = (1.5 * 0.0)
    v1 += (v1 % ((4294967295) | 1))
    fmt.Println(v1)
    v2 = append(v2, (("\n\t" + (((("" + ((("" + (("\n\t" + "hello") + "")) + "\n\t") + "")) + "hello") + "foo bar") + "")) + ("foo bar" + "foo bar")))
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 uint16 = 65535
        v3[127] = (0.0 <= float32(3.14159))
        switch "foo bar" {
        case "":
            v25 = (v25 / (((v25 + (^(v25)))) | 1))
        }
        _ = fi24
        _ = v25
    }
    v1 = 2147483648
    v26 := GenS2{F0: 2, F1: (65535 == 42)}
    v27 := GenS3{F0: 42, F1: 0, F2: 2147483647}
    delete(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
}

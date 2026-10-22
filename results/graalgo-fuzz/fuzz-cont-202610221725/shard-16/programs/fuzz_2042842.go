package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 string, p1 uint64) int {
    var v6 int16 = 450
    v7 := GenS1{F0: 32767, F1: 255}
    v8 := []int{128, -595, 2147483647, 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]string)
    var v3 uint64 = 0
    var v4 uint64 = v3
    var v5 float64 = 3.14159
    v9 := GenS1{F0: 312, F1: 954}
    v10 := GenS1{F0: 629, F1: 316}
    var v11 int64 = 0
    fmt.Println(v11, v1, v11)
    fmt.Println(v5, v1, v11)
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 32767, F1: 335}
        v14 := GenS1{F0: 637, F1: 65535}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: 32767, F2: 977}
    var v16 uint32 = 475
    switch "a" {
    case "hello":
        v17 := []int{1, int(v4), v1, (-(v1))}
        _ = v17
    case "\n\t":
        v18 := GenS1{F0: 65535, F1: 42}
        _ = v18
    case "a":
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS3{F0: 18446744073709551615, F1: 42}
            for fi21 := 0; fi21 < 6; fi21++ {
                v22 := []int{(v1 | -479)}
                switch ("a" + "\n\t") {
                case "foo bar":
                    if false {
                        if true {
                            var v23 int16 = 42
                            _ = v23
                        }
                    }
                case "\n\t":
                    v3 = v4
                case "":
                    var v24 float32 = (float32(v5) / (1e10 / 1.5))
                    _ = v24
                }
                v25 := GenS4{F0: v16, F1: 33, F2: ("hello" + ("\n\t" + "foo bar"))}
                _ = fi21
                _ = v22
                _ = v25
            }
            _ = fi19
            _ = v20
        }
    default:
        var v26 uint16 = 0
        _ = v26
    }
    v27 := GenS1{F0: 65535, F1: 0}
    v2["a"] = "foo bar"
    v28 := GenS5{F0: (v16 * (v16 << ((18446744073709551615) & 63))), F1: -2147483648}
    var v29 float32 = 1e10
    var v30 float32 = v29
    if ((!((v16 >> ((368) & 63)) != 2)) && true) {
        v31 := v2["\n\t"]
        v29 = (float32(1e10) * v29)
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS6{F0: v16, F1: 1, F2: 1000}
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    v29 = (3.14159 / (((1.5 - -2.25) * (0.1 * v30)) + v30))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}

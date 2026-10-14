package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int8; F2 int64 }
type GenS6 struct { F0 bool; F1 int32 }
func genFn0() int16 {
    var v5 uint16 = 1000
    v5 += (v5 * (v5 >> ((255) & 63)))
    _ = v5
    return -124
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]string)
    var v3 float32 = float32((0.0 + float64(1.5)))
    var v4 float64 = (0.1 / 3.14159)
    switch ("" + ("\n\t" + "")) {
    case "":
        if ((-9223372036854775808 >= v1) && false) {
            var v6 int8 = 89
            _ = v6
        } else {
            fmt.Println(v3)
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: (("hello" + ("foo bar" + "foo bar")) + "")}
                v9 := GenS1{F0: "foo bar"}
                v3 -= 1e10
                _ = fi7
                _ = v8
                _ = v9
            }
            switch "foo bar" {
            case "hello":
                v1 = v1
            }
        }
    }
    var v10 int = 128
    var v11 uint32 = 65535
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 string = ("hello" + ("a" + ("foo bar" + "a")))
            if true {
                var v15 int32 = 0
                v4 = float64((v3 / 1.5))
                if (v11 >= v11) {
                    v11 = v11
                }
                _ = v15
            }
            v16 := GenS2{F0: 573, F1: 1000}
            _ = fi13
            _ = v14
            _ = v16
        }
        v3 += (0.0 / 1.5)
        _ = fi12
    }
    if (615 < 42) {
        v17 := GenS1{F0: "a"}
        v18 := []int{v10, v10, 9223372036854775807, (v10 / (((^(v10))) | 1))}
        _ = v17
        _ = v18
    } else {
        v11 -= (v11 | (v11 / (((v11 & v11)) | 1)))
        delete(v2, v10)
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v10 -= v10
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS3{F0: fi20, F1: 770}
            var v22 uint = 153
            _ = fi20
            _ = v21
            _ = v22
        }
        v23 := GenS4{F0: 662, F1: 1, F2: 1}
        _ = fi19
        _ = v23
    }
    v10 = -2147483648
    v24 := []int{968, (-(v10))}
    var v25 uint = 65535
    switch ("" + "") {
    case "a":
        switch -2147483648 {
        case 297:
            if ((!false) || (((float32(v4) / 1.5) - 0.1) != (3.14159 - v3))) {
                if false {
                    switch "hello" {
                    case "foo bar":
                        if ("foo bar" < "hello") {
                            fmt.Println(v4, v25)
                            var v26 int64 = v1
                            v27 := GenS5{F0: 255, F1: 6, F2: v26}
                            _ = v26
                            _ = v27
                        }
                    }
                    switch "hello" {
                    case "hello":
                        v28 := v2[v10]
                        _ = v28
                    case "\n\t":
                        fmt.Println(v25, v1, v4)
                    case "a":
                        v29 := GenS5{F0: 69, F1: 100, F2: (v1 >> ((v25) & 63))}
                        _ = v29
                    }
                }
            } else {
                v1 = (^(v1))
                if ((v11 != (v11 * (v11 << (((v25 - v25)) & 63)))) || true) {
                    v30 := GenS6{F0: false, F1: 931}
                    v11 = uint32(v25)
                    var v31 uint16 = 0
                    _ = v30
                    _ = v31
                }
            }
        case 4294967295:
            fmt.Println(v25, v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}

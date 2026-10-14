package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0() int {
    var v7 float64 = float64(3.14159)
    var v8 float64 = v7
    switch "" {
    case "a":
        v9 := []int{127, 1, 767, -34}
        _ = v9
    default:
        switch -1 {
        case -9223372036854775808:
            for fi10 := 0; fi10 < 15; fi10++ {
                if false {
                    v11 := GenS1{F0: 720, F1: 100, F2: ("\n\t" + "\n\t")}
                    if false {
                        switch ("hello" + ("" + "a")) {
                        case "":
                            var v12 uint64 = 2
                            _ = v12
                        case "foo bar":
                            var v13 uint64 = 4294967295
                            _ = v13
                        case "a":
                            v14 := []int{9223372036854775807, fi10, 999}
                            _ = v14
                        default:
                            var v15 uint = 229
                            _ = v15
                        }
                        var v16 uint8 = 42
                        _ = v16
                    } else {
                        var v17 float64 = 1.5
                        _ = v17
                    }
                    _ = v11
                } else {
                    v18 := GenS2{F0: 1, F1: 1}
                    _ = v18
                }
                var v19 int8 = 42
                _ = fi10
                _ = v19
            }
        case -1:
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint16 = 255
                _ = fi20
                _ = v21
            }
        case 32768:
            v22 := []int{2147483647, 2147483647, -2147483648}
            _ = v22
        }
    }
    v23 := []int{-2147483648}
    if ((!false) || (0 == -1)) {
        v24 := GenS3{F0: -21, F1: -3, F2: ("a" + ("foo bar" + ""))}
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 int8 = 127
            _ = fi25
            _ = v26
        }
        fmt.Println(v8, v7)
        _ = v24
    }
    _ = v23
    _ = v7
    _ = v8
    return -660
}
func genFn1() uint8 {
    v27 := GenS3{F0: 0, F1: -32768, F2: ("hello" + "foo bar")}
    v28 := GenS2{F0: 0, F1: 1}
    _ = v27
    _ = v28
    return 20
}

func main() {
    var v1 uint = 2
    var v2 int8 = 127
    var v3 float64 = 3.14159
    var v4 bool = false
    var v5 float32 = (1.5 + 0.1)
    v6 := make(map[string]bool)
    v29 := GenS4{F0: "a"}
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 15; fi31++ {
            fmt.Println(v2)
            v6["hello"] = ((fi31 <= (fi31 / (((fi31 % (((fi31 & fi31)) | 1))) | 1))) || v4)
            v32 := GenS5{F0: 2, F1: v2, F2: 266}
            _ = fi31
            _ = v32
        }
        v33 := GenS6{F0: 1, F1: 717, F2: 255}
        _ = fi30
        _ = v33
    }
    v1 = (v1 & (v1 + (v1 >> (((v1 << (((v1 << (((v1 >> (((^(v1))) & 63))) & 63))) & 63))) & 63))))
    v34 := v6["a"]
    var v35 int16 = 0
    fmt.Println(v5, v3, v34)
    v3 = (1.5 + v3)
    fmt.Println(v4, v35, v34)
    v34 = ("foo bar" != "foo bar")
    var v36 int64 = -769
    if v34 {
        var v37 int64 = v36
        v2 += -103
        _ = v37
    }
    if v4 {
        for fi38 := 0; fi38 < 15; fi38++ {
            switch "a" {
            case "":
                if (v34 && false) {
                    var v39 bool = (v34 || ((fi38 << ((18446744073709551615) & 63)) >= genFn0()))
                    _ = v39
                }
            default:
                for fi40 := 0; fi40 < 7; fi40++ {
                    fmt.Println(v5)
                    v35 -= (^(v35))
                    _ = fi40
                }
            }
            fmt.Println(v5, v34)
            v41 := GenS6{F0: 18446744073709551615, F1: 18446744073709551615, F2: (v1 / (((v1 & v1)) | 1))}
            _ = fi38
            _ = v41
        }
    }
    var v42 uint16 = 1
    v3 = (v3 - v3)
    v43 := GenS3{F0: (v35 & 1000), F1: 32767, F2: ("hello" + ("hello" + ""))}
    if v34 {
        if (!v34) {
            var v44 int16 = v35
            for fi45 := 0; fi45 < 5; fi45++ {
                v46 := GenS4{F0: "hello"}
                _ = fi45
                _ = v46
            }
            _ = v44
        }
    }
    for fi47 := 0; fi47 < 9; fi47++ {
        if (!(((true || v4) && v34) || (-2.25 >= 3.14159))) {
            v3 = v3
        } else {
            var v48 string = ""
            _ = v48
        }
        _ = fi47
    }
    var v49 float32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v49
    _ = v5
    _ = v6
}

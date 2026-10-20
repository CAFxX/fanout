package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint32; F1 uint64; F2 int }
func genFn0() bool {
    var v5 int = 32767
    v6 := GenS1{F0: v5, F1: ("" + "hello")}
    switch ("\n\t" + "") {
    case "":
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{-2147483648, 1, (-(fi7))}
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v5)
    _ = v5
    _ = v6
    return (false && (!(!false)))
}
func genFn1(p0 uint) uint16 {
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int = (-(fi9))
        if false {
            switch fi9 {
            case -709:
                var v11 float64 = 3.14159
                _ = v11
            }
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(p0)
    switch "" {
    case "a":
        p0 = p0
    case "hello":
        v12 := GenS1{F0: -1, F1: "\n\t"}
        _ = v12
    default:
        if false {
            switch "" {
            case "\n\t":
                var v13 uint16 = 0
                _ = v13
            case "foo bar":
                if true {
                    v14 := GenS1{F0: 4294967295, F1: ""}
                    _ = v14
                }
            default:
                p0 -= (p0 % ((p0) | 1))
            }
            v15 := GenS1{F0: 9223372036854775807, F1: "hello"}
            _ = v15
        }
    }
    v16 := []int{0, 460, 4294967295}
    _ = p0
    _ = v16
    return 32767
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 string = ("a" + "hello")
    var v4 int8 = 1
    if (!true) {
        v3 = v3
    } else {
        var v17 float64 = (0.0 * float64(0.0))
        v18 := GenS2{F0: ((v4 & v4) > (v4 % ((v4) | 1))), F1: 886, F2: ((("" + v3) + (((v3 + (v3 + "hello")) + v3) + v3)) + "")}
        if (397 >= 18446744073709551615) {
            v19 := GenS1{F0: 458, F1: "a"}
            _ = v19
        } else {
            v3 = v3
            _ = len(v1)
            var v20 uint32 = 2147483648
            _ = v20
        }
        _ = v17
        _ = v18
    }
    v21 := GenS2{F0: true, F1: 65535, F2: "a"}
    switch 4294967295 {
    case 127:
        for fi22 := 0; fi22 < 1; fi22++ {
            v3 = v3
            switch (v3 + v3) {
            case "hello":
                v3 = v3
            case "":
                v4 += 83
            case "\n\t":
                v23 := GenS1{F0: -71, F1: v3}
                _ = v23
            }
            _ = fi22
        }
    case 920:
        v24 := GenS1{F0: -829, F1: (v3 + v3)}
        _ = v24
    case -452:
        v3 = v3
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            v4 = (v4 / (((v4 / (((v4 & -1)) | 1))) | 1))
            var v27 int8 = -1
            var v28 int32 = 2147483647
            _ = fi26
            _ = v27
            _ = v28
        }
        v3 = (v3 + ("\n\t" + v3))
        v3 = "hello"
        _ = fi25
    }
    v4 += v4
    v3 = v3
    v3 = ((("foo bar" + ("hello" + (v3 + ("foo bar" + v3)))) + "a") + v3)
    if (!true) {
        if true {
            _ = len(v2)
            switch 127 {
            case 42:
                var v29 uint8 = 2
                _ = v29
            case 613:
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint32 = 753
                    _ = fi30
                    _ = v31
                }
            case -94:
                v32 := GenS1{F0: -1, F1: v3}
                _ = v32
            default:
                if (100 != 1) {
                    var v33 float32 = 0.1
                    _ = v33
                } else {
                    v34 := GenS1{F0: 42, F1: v3}
                    v3 = v3
                    if (!(0 > 1)) {
                        var v35 uint16 = 719
                        for fi36 := 0; fi36 < 12; fi36++ {
                            v37 := GenS3{F0: 4294967295, F1: 42, F2: -9223372036854775808}
                            var v38 uint16 = (v35 & genFn1(55))
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    } else {
                        v39 := GenS1{F0: -956, F1: v3}
                        _ = v39
                    }
                    _ = v34
                }
            }
            v3 = v3
        }
        v40 := []int{4294967295, 2147483647}
        _ = v40
    }
    v4 = 32
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}

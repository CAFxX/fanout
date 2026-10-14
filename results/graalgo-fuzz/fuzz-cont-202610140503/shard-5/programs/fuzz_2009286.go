package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 bool; F1 int8; F2 int }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint32, p2 uint) uint32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            fmt.Println(fi5, p1, p2)
            _ = fi5
        }
        _ = fi4
    }
    v6 := GenS1{F0: -89, F1: "hello"}
    if true {
        var v7 int16 = 1000
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int32 = 2147483647
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float32 = (0.1 / 1e10)
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v9, fi8, fi10)
            if (0 <= 869) {
                v11 := []int{fi8, fi10, fi8}
                v9 = (v9 * (0.0 * (1.5 - float32((1.5 + 0.1)))))
                _ = v11
            }
            v12 := GenS2{F0: 824}
            _ = fi10
            _ = v12
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            if ((0.1 == float32((0.1 + ((1e10 * -2.25) * -2.25)))) && true) {
                delete(v2, "foo bar")
                v14 := GenS3{F0: 985, F1: 1}
                _ = v14
            }
            switch fi8 {
            case 286:
                v15 := GenS4{F0: (true || true), F1: 42, F2: 4294967295}
                _ = v15
            }
            v16 := GenS1{F0: v3, F1: ("foo bar" + "a")}
            _ = fi13
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := []int{(fi18 & (fi17 >> ((18446744073709551615) & 63)))}
            if (2 != 0) {
                switch "" {
                case "\n\t":
                    fmt.Println(v3, fi17)
                    v3 = v3
                    fmt.Println(v3)
                default:
                    if len(v19) > 0 {
                        v19[0] = fi17
                    }
                    var v20 uint = 42
                    var v21 int32 = 1
                    _ = v20
                    _ = v21
                }
                v2["hello"] = ("" + "")
                delete(v1, (fi17 ^ -1))
                v22 := GenS4{F0: (!false), F1: -104, F2: (fi18 ^ (fi18 - fi18))}
                var v23 string = ("a" + "a")
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v24 int32 = v3
    v25 := GenS5{F0: 235, F1: 1}
    v24 = 42
    v24 = (v24 & (v3 | -2147483648))
    v26 := GenS4{F0: (3.14159 == 0.1), F1: -128, F2: 128}
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := v2["a"]
        _ = fi27
        _ = v28
    }
    v24 = -1
    if (!true) {
        var v29 uint32 = genFn0((v3 >> ((18446744073709551615) & 63)), 674, 90)
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        if true {
            switch "" {
            case "":
                v24 = (v3 % ((v24) | 1))
            default:
                v24 = v24
            }
        }
        _ = fi30
    }
    v31 := []int{}
    v24 = 42
    v32 := GenS6{F0: 42}
    var v33 uint16 = 0
    if true {
        switch -9223372036854775808 {
        case 32767:
            fmt.Println(v24)
        case 9223372036854775807:
            var v34 int8 = 0
            _ = v34
        default:
            for fi35 := 0; fi35 < 7; fi35++ {
                v36 := GenS6{F0: 1}
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 int8 = 0
                    v33 -= v33
                    _ = fi37
                    _ = v38
                }
                v33 += (v33 - 2)
                _ = fi35
                _ = v36
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}

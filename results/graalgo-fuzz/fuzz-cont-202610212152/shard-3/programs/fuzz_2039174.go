package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }

func main() {
    var v1 int16 = 0
    var v2 int8 = 127
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 18446744073709551615, F1: 42, F2: 243}
        var v6 uint8 = 42
        _ = fi4
        _ = v5
        _ = v6
    }
    v2 += (^(v2))
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint32 = 42
            _ = fi8
            _ = v9
        }
        switch (fi7 ^ 127) {
        case -6:
            v2 = v2
        case 1:
            for fi10 := 0; fi10 < 16; fi10++ {
                if (((("hello" + ((("a" + "") + "foo bar") + "\n\t")) + ("" + ("a" + "\n\t"))) + "a") >= "a") {
                    v11 := []int{fi10}
                    _ = v11
                } else {
                    v2 = (v2 & 0)
                    var v12 string = "\n\t"
                    var v13 string = "hello"
                    _ = v12
                    _ = v13
                }
                var v14 int64 = -781
                if (fi7 == (-(fi10))) {
                    v2 -= -128
                }
                v1 = (v1 << ((754) & 63))
                switch "foo bar" {
                case "hello":
                    v15 := GenS1{F0: 898, F1: 65535, F2: 138}
                    _ = v15
                default:
                    var v16 int8 = v2
                    _ = v16
                }
                _ = fi10
                _ = v14
            }
        case 0:
            v17 := GenS1{F0: 42, F1: 4294967295, F2: 0}
            _ = v17
        default:
            var v18 float64 = 0.0
            _ = v18
        }
        v19 := GenS1{F0: 198, F1: 2147483648, F2: 39}
        _ = fi7
        _ = v19
    }
    switch "" {
    case "\n\t":
        var v20 uint16 = 1
        _ = v20
    case "":
        if true {
            v21 := []int{2147483647, 32768, 1}
            _ = v21
        }
    default:
        var v22 bool = false
        _ = v22
    }
    v2 = (v2 | (v2 >> ((4294967295) & 63)))
    if false {
        v1 = v1
    }
    fmt.Println(v2, v2)
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := []int{(fi23 + fi23)}
        var v25 int32 = -873
        if false {
            v26 := GenS1{F0: 708, F1: 2147483647, F2: 1}
            for fi27 := 0; fi27 < 12; fi27++ {
                v28 := GenS2{F0: 1, F1: (v2 ^ -128), F2: 2}
                v29 := GenS1{F0: 0, F1: 937, F2: 42}
                _ = fi27
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                var v31 uint16 = 255
                var v32 float64 = ((((1e10 + (0.1 / 1.5)) + float64(1e10)) - -2.25) * 1e10)
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v33 uint8 = 0
    v2 -= 95
    var v34 uint32 = 983
    v3[uint32(v33)] = 2147483648
    if false {
        v35 := []int{317, 42, 237}
        _ = v35
    } else {
        var v36 float32 = 1.5
        for fi37 := 0; fi37 < 3; fi37++ {
            v36 += float32((1e10 * float64(v36)))
            var v38 float64 = (1.5 * (0.0 + 1.5))
            v39 := GenS2{F0: 42, F1: v2, F2: 65535}
            _ = fi37
            _ = v38
            _ = v39
        }
        var v40 bool = true
        _ = v36
        _ = v40
    }
    fmt.Println(v2, v33)
    var v41 uint = 18446744073709551615
    switch 128 {
    case 4294967295:
        fmt.Println(v1)
    case -1:
        switch "foo bar" {
        case "a":
            var v42 int = 4294967295
            _ = v42
        case "hello":
            for fi43 := 0; fi43 < 13; fi43++ {
                v44 := []int{}
                v41 = (v41 << ((v41) & 63))
                _ = fi43
                _ = v44
            }
        }
    case 2147483647:
        var v45 int16 = (v1 + (v1 - (v1 - (-(v1)))))
        _ = v45
    default:
        var v46 int = 2147483647
        _ = v46
    }
    if (!(1.5 != 1e10)) {
        _ = len(v3)
    } else {
        delete(v3, v34)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v41
}

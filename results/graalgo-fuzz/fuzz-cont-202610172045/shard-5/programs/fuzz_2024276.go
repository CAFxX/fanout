package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint; F2 string }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0() uint {
    var v5 float64 = 0.1
    switch ("a" + "\n\t") {
    case "a":
        for fi6 := 0; fi6 < 8; fi6++ {
            switch "" {
            case "\n\t":
                var v7 uint = 501
                _ = v7
            }
            var v8 int32 = -463
            var v9 bool = true
            _ = fi6
            _ = v8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint64 = 2
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := []int{}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            switch ("a" + "") {
            case "a":
                v15 := GenS1{F0: (227 == 939), F1: (!(((0.0 * -2.25) * float32(0.0)) != float32(3.14159)))}
                _ = v15
            default:
                if true {
                    v5 = float64(1e10)
                    v16 := GenS2{F0: 59, F1: 0}
                    v17 := GenS1{F0: true, F1: ((350 != 1000) && false)}
                    _ = v16
                    _ = v17
                }
            }
            _ = fi13
            _ = v14
        }
        if false {
            var v18 uint64 = 42
            if (1 == 4294967295) {
                v19 := GenS1{F0: true, F1: ((false && true) && ((false || false) || (true && true)))}
                v5 = (v5 / 3.14159)
                _ = v19
            }
            _ = v18
        } else {
            var v20 uint8 = 0
            var v21 uint = 42
            v22 := []int{fi12, 42, (fi12 | fi12), fi12}
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi12
    }
    _ = v5
    return 112
}
func genFn1(p0 uint64) int64 {
    p0 += (^(p0))
    var v23 uint = 575
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(fi24)
        if ((((270 != 42) && true) && false) && (((!false) || false) || true)) {
            v23 += (v23 % (((v23 - 65535)) | 1))
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 float64 = ((((3.14159 - ((1.5 + 3.14159) * ((3.14159 / 1.5) + ((1e10 - (float64(float32(1e10)) - 0.1)) / 1e10)))) / (((0.1 + 1.5) + -2.25) * ((1.5 - (float64(0.0) / 1.0)) - -2.25))) * 0.0) * -2.25)
                var v27 uint64 = p0
                _ = fi25
                _ = v26
                _ = v27
            }
        } else {
            v28 := []int{fi24, 9223372036854775807}
            v28 = append(v28, (fi24 >> (((v23 << ((18446744073709551615) & 63))) & 63)))
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 bool = (!(0.1 >= (0.0 - 3.14159)))
                var v31 int32 = -952
                v32 := GenS3{F0: 595, F1: genFn0(), F2: ("a" + ("a" + ""))}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = v28
        }
        _ = fi24
    }
    p0 = p0
    _ = p0
    _ = v23
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (((0.0 - 3.14159) / 1.5) - float64(0.0))
    var v3 bool = true
    var v4 int32 = 0
    var v33 float64 = v2
    var v34 int8 = 62
    for fi35 := 0; fi35 < 2; fi35++ {
        v36 := GenS4{F0: -32768, F1: "\n\t"}
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        for fi38 := 0; fi38 < 8; fi38++ {
            v2 = v33
            var v39 int32 = 699
            v1["hello"] = ((v34 <= v34) || false)
            _ = fi38
            _ = v39
        }
        v40 := GenS4{F0: 32767, F1: "foo bar"}
        _ = fi37
        _ = v40
    }
    v2 -= (v33 + 1.5)
    v41 := []int{1, -492, -9223372036854775808}
    var v42 string = "\n\t"
    for fi43 := 0; fi43 < 11; fi43++ {
        switch (v42 + "") {
        case "a":
            delete(v1, v42)
        case "foo bar":
            var v44 uint64 = 216
            _ = v44
        default:
            var v45 int8 = (v34 / (((v34 / (((v34 % (((^(v34))) | 1))) | 1))) | 1))
            _ = v45
        }
        delete(v1, "hello")
        for fi46 := 0; fi46 < 5; fi46++ {
            switch (fi46 % (((fi43 + fi43)) | 1)) {
            case 2147483647:
                var v47 int16 = -161
                _ = v47
            case 9223372036854775807:
                v48 := GenS5{F0: v3}
                _ = v48
            case 32767:
                _ = len(v1)
            }
            if (!(!v3)) {
                v49 := GenS6{F0: v4, F1: 77, F2: 255}
                _ = v49
            }
            _ = fi46
        }
        _ = fi43
    }
    var v50 uint8 = 77
    var v51 uint16 = 820
    var v52 int64 = 420
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v50
    _ = v51
    _ = v52
}

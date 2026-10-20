package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint32; F1 uint32; F2 string }
func genFn0() int {
    for fi4 := 0; fi4 < 2; fi4++ {
        if false {
            var v5 uint32 = 191
            var v6 int16 = -310
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: (v6 ^ v6), F1: 127}
                var v9 uint = 824
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v5
            _ = v6
        }
        if false {
            var v10 uint = 443
            v11 := []int{}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                if (!false) {
                    var v13 uint32 = 2147483647
                    var v14 int = (fi4 - fi4)
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(fi12)
                    var v15 int = fi12
                    _ = v15
                }
                v16 := GenS2{F0: fi12, F1: 4294967295, F2: -264}
                _ = fi12
                _ = v16
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(fi17, fi17)
                _ = fi17
            }
            fmt.Println(fi4)
        }
        v18 := GenS3{F0: 1, F1: 2147483647, F2: 329}
        _ = fi4
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint64 = 18446744073709551615
            var v22 float32 = -2.25
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 float64 = 1.5
        _ = fi23
        _ = v24
    }
    switch -277 {
    case 127:
        v25 := []int{9223372036854775807, -734, -1, -251}
        _ = v25
    default:
        var v26 uint8 = 2
        _ = v26
    }
    v27 := GenS3{F0: 640, F1: 532, F2: 0}
    _ = v27
    return 128
}
func genFn1() uint {
    v28 := []int{128, 32767, 881, 32767}
    var v29 float32 = ((-2.25 - 0.1) * -2.25)
    v28 = append(v28, genFn0())
    _ = v28
    _ = v29
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 1
    var v3 uint32 = 4294967295
    switch 127 {
    case 2147483647:
        v30 := GenS4{F0: v3, F1: -582}
        _ = v30
    case 4294967295:
        switch ("a" + "foo bar") {
        case "foo bar":
            v3 -= v3
        case "":
            if ((v3 ^ 0) == (v3 % ((4294967295) | 1))) {
                delete(v1, v3)
            }
        case "a":
            var v31 uint64 = 42
            _ = v31
        default:
            switch ((("\n\t" + "") + "a") + "a") {
            case "\n\t":
                v2 += (^(v2))
            default:
                fmt.Println(v3)
            }
        }
    }
    v2 = (v2 * (v2 - (^(v2))))
    var v32 bool = (!(1000 == 42))
    fmt.Println(v32)
    v33 := GenS1{F0: -32768, F1: 1}
    var v34 int32 = 2147483647
    v35 := []int{128}
    if (!((!(false && v32)) && v32)) {
        v36 := GenS5{F0: (v3 ^ 410), F1: 2147483648, F2: "\n\t"}
        _ = v36
    } else {
        v37 := GenS5{F0: (v3 / ((v3) | 1)), F1: 504, F2: "foo bar"}
        _ = v37
    }
    var v38 int16 = 1
    var v39 int16 = v38
    var v40 int16 = (^(v38))
    switch "foo bar" {
    case "hello":
        var v41 int32 = (^(v34))
        _ = v41
    case "a":
        var v42 int8 = -58
        _ = v42
    case "foo bar":
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v40
}

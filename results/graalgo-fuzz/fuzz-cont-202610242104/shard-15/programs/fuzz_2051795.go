package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint16 }
func genFn0(p0 uint8, p1 uint, p2 uint16) string {
    v4 := []int{148, 127}
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 float64 = 1e10
        var v7 int32 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    p1 = 613
    var v8 float32 = (float32((((((0.0 * (-2.25 * -2.25)) - (1.5 * float64((0.1 / -2.25)))) - 3.14159) * (-2.25 / 3.14159)) * 0.1)) * 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return "\n\t"
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    v3 := []uint32{42, 2147483648}
    switch -607 {
    case -433:
        for fi9 := 0; fi9 < 16; fi9++ {
            if (!((160 > 538) && true)) {
                v1 = (v1 - (v1 * v1))
                v3 = append(v3, 324)
                var v10 float32 = (v1 * v1)
                _ = v10
            }
            if true {
                v11 := v2[333]
                fmt.Println(fi9, v11, v11)
                v3 = append(v3, 718)
                _ = v11
            }
            switch (-(fi9)) {
            case 128:
                switch fi9 {
                case -564:
                    var v12 int32 = -2147483648
                    _ = v12
                case 15:
                    fmt.Println(fi9, fi9)
                }
            case 127:
                for fi13 := 0; fi13 < 12; fi13++ {
                    v2[2147483648] = "foo bar"
                    v14 := v2[982]
                    v15 := GenS1{F0: ((!true) && (!true)), F1: false}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case -181:
                if (!(176 == -761)) {
                    var v16 int32 = 0
                    fmt.Println(v1, v1, fi9)
                    _ = v16
                } else {
                    v1 += v1
                    var v17 uint = 2
                    _ = v17
                }
            }
            _ = fi9
        }
    case 421:
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint64 = 18446744073709551615
                if true {
                    var v21 int16 = -32768
                    fmt.Println(fi18)
                    _ = v21
                } else {
                    _ = len(v2)
                }
                var v22 int = (^(fi18))
                fmt.Println(v20)
                _ = fi19
                _ = v20
                _ = v22
            }
            v23 := GenS2{F0: 4294967295}
            var v24 int32 = 42
            _ = fi18
            _ = v23
            _ = v24
        }
    case -9223372036854775808:
        v25 := GenS3{F0: 1}
        _ = v25
    default:
        fmt.Println(v1)
    }
    var v26 uint = 421
    if len(v3) > 0 {
        v3[0] = 338
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        _ = len(v2)
        _ = len(v2)
        switch "foo bar" {
        case "hello":
            v28 := v2[2147483647]
            _ = v28
        }
        _ = fi27
    }
    fmt.Println(v1, v26)
    v3 = append(v3, 4294967295)
    v29 := GenS4{F0: -9223372036854775808, F1: 1}
    v30 := GenS5{F0: 0}
    if true {
        var v31 string = (("" + ("hello" + "")) + (((("\n\t" + ("foo bar" + "a")) + "a") + ("\n\t" + "foo bar")) + "hello"))
        _ = v31
    }
    v26 = 477
    v32 := GenS6{F0: -2147483648, F1: 42, F2: 2}
    v33 := v2[1]
    var v34 int64 = -1
    var v35 string = "\n\t"
    var v36 uint32 = 2147483648
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 int16 = 278
        v39 := v2[v36]
        v40 := GenS7{F0: 127, F1: true, F2: 32767}
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}

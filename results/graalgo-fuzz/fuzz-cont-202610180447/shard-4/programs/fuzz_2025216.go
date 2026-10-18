package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0() int32 {
    var v5 uint8 = 31
    if (!false) {
        v5 = (v5 ^ v5)
    } else {
        switch 9223372036854775807 {
        case 32768:
            var v6 string = (("" + "a") + "")
            _ = v6
        case -975:
            v5 += v5
        case 2147483647:
            v5 += (v5 / (((v5 >> ((933) & 63))) | 1))
        default:
            switch 4294967295 {
            case 2147483647:
                if false {
                    for fi7 := 0; fi7 < 3; fi7++ {
                        v5 = v5
                        v8 := []int{fi7}
                        _ = fi7
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 20; fi9++ {
                        for fi10 := 0; fi10 < 4; fi10++ {
                            v5 += (^(v5))
                            _ = fi10
                        }
                        for fi11 := 0; fi11 < 16; fi11++ {
                            var v12 int32 = 1
                            v13 := GenS1{F0: 1, F1: 718}
                            v14 := GenS1{F0: 929, F1: 430}
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        v15 := []int{4294967295, fi9}
                        _ = fi9
                        _ = v15
                    }
                } else {
                    v5 -= (v5 | (v5 | (v5 % (((v5 | v5)) | 1))))
                    v5 += v5
                }
            }
        }
    }
    _ = v5
    return 0
}
func genFn1(p0 bool, p1 int, p2 int8) int32 {
    var v16 uint16 = 255
    v17 := []int{(-(p1))}
    v18 := GenS2{F0: (p2 + -58), F1: 42}
    var v19 int8 = int8(p1)
    fmt.Println(v19)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    v1 := []int{128, -686}
    v2 := []string{"hello"}
    var v3 float32 = 1.5
    var v4 float32 = v3
    _ = len(v1)
    switch (("hello" + ("foo bar" + ("\n\t" + ("" + "")))) + ("foo bar" + ("foo bar" + ""))) {
    case "hello":
        fmt.Println(v4, v4)
    default:
        var v20 uint = 0
        _ = v20
    }
    var v21 float64 = ((float64(float32(0.1)) / ((1e10 + (3.14159 - float64((((v4 + (-2.25 / float32((float64(v3) / (0.1 / -2.25))))) / (-2.25 + v3)) * v4)))) - -2.25)) - float64(3.14159))
    v4 -= v3
    fmt.Println(v3)
    var v22 uint16 = 160
    v4 += (v4 + v3)
    var v23 string = "\n\t"
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(v3, v3)
        var v25 float64 = (v21 / v21)
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 int16 = -644
            v25 -= v25
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((!true) && (42 <= 32767)) {
        var v28 uint32 = 2147483648
        v21 = ((v21 * (v21 - float64((v3 / (v4 + v3))))) + ((1.5 / (v21 - ((0.1 / v21) + v21))) / v21))
        switch (v23 + "a") {
        case "\n\t":
            if true {
                v3 = v3
            } else {
                var v29 bool = true
                _ = v29
            }
        }
        _ = v28
    } else {
        var v30 uint64 = 214
        if len(v1) > 0 {
            v1[0] = -984
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            if (-9223372036854775808 == (^(fi31))) {
                var v32 int64 = -9223372036854775808
                var v33 string = (v23 + v23)
                _ = v32
                _ = v33
            } else {
                if (42 != 2147483647) {
                    var v34 bool = false
                    fmt.Println(v4, v23)
                    _ = v34
                }
            }
            v35 := GenS1{F0: 1, F1: 42}
            v36 := GenS1{F0: 172, F1: 2}
            _ = fi31
            _ = v35
            _ = v36
        }
        _ = v30
    }
    var v37 bool = true
    var v38 uint64 = uint64(v22)
    fmt.Println(v4)
    var v39 uint8 = 42
    v1 = append(v1, -1)
    fmt.Println(v37)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}

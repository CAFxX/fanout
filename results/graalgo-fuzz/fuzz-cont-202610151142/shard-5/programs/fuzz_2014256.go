package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
func genFn0() uint64 {
    v7 := []int{9223372036854775807, -9223372036854775808, 853, -576}
    if false {
        v8 := GenS1{F0: 18446744073709551615, F1: 42}
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int32 = 85
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint64 = 586
                _ = fi11
                _ = v12
            }
            v10 = -1
            _ = fi9
            _ = v10
        }
        var v13 uint = 65535
        _ = v13
        _ = v8
    } else {
        var v14 int8 = 127
        switch 42 {
        case 127:
            v14 += (v14 * v14)
        case 118:
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 uint32 = 42
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int16 = 1000
                    v18 = (v18 >> ((0) & 63))
                    v19 := GenS1{F0: 2, F1: 127}
                    if ((v18 / (((v18 / (((v18 / ((v18) | 1))) | 1))) | 1)) == -364) {
                        var v20 uint64 = 0
                        _ = v20
                    }
                    fmt.Println(v14, v16)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v21 int8 = v14
                _ = fi15
                _ = v16
                _ = v21
            }
        case 42:
            var v22 float32 = 0.0
            _ = v22
        default:
            for fi23 := 0; fi23 < 12; fi23++ {
                v14 = (v14 ^ (v14 + (^(v14))))
                if false {
                    if false {
                        var v24 uint16 = 0
                        if (true || false) {
                            v14 -= -128
                        }
                        _ = v24
                    }
                    v14 -= (v14 - v14)
                }
                for fi25 := 0; fi25 < 16; fi25++ {
                    v14 += 1
                    _ = fi25
                }
                _ = fi23
            }
        }
        _ = v14
    }
    var v26 bool = ((false || true) && true)
    v26 = ((244 != 440) && false)
    fmt.Println(v26)
    _ = v26
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 int64, p2 int8) int32 {
    var v27 string = ((("hello" + ((("" + ("hello" + "foo bar")) + ("\n\t" + ("foo bar" + "\n\t"))) + "")) + ("a" + "a")) + "a")
    v28 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return 560
}

func main() {
    var v1 uint32 = 706
    var v2 int32 = -2147483648
    var v3 int16 = -817
    var v4 uint = 1
    var v5 int = 415
    var v6 int32 = v2
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 int32 = v6
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint64 = 980
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v2 = v6
    var v33 float64 = 1e10
    fmt.Println(v5, v1, v5)
    var v34 int32 = (v6 ^ v2)
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 uint64 = 42
        _ = fi35
        _ = v36
    }
    if false {
        fmt.Println(v3, v1)
        v37 := []int{0, 2147483647, (v5 ^ 4294967295)}
        _ = v37
    } else {
        v38 := GenS3{F0: (v3 ^ (^(v3)))}
        _ = v38
    }
    v3 += v3
    var v39 int64 = 9223372036854775807
    v40 := []int{4294967295, (v5 - 1), 0, v5}
    if true {
        v41 := GenS4{F0: 357}
        if len(v40) > 0 {
            v40[0] = v5
        }
        fmt.Println(v39, v4)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

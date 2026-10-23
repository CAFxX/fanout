package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint }

func main() {
    var v1 string = "hello"
    var v2 int32 = 0
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 uint64 = 18446744073709551615
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint8 = 255
        _ = fi7
        _ = v8
    }
    var v9 uint = 42
    if false {
        var v10 uint8 = uint8(v9)
        v2 += v2
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int32 = (v2 | v2)
            fmt.Println(v6, v2)
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        if (-879 >= 0) {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := []int{(fi13 / (((fi13 / ((1) | 1))) | 1))}
                v15 := GenS1{F0: v5}
                if (v5 >= v5) {
                    fmt.Println(v2, v9, v1)
                    switch 1 {
                    case 32768:
                        var v16 int16 = -748
                        _ = v16
                    case -2147483648:
                        v14 = append(v14, fi13)
                    case 2147483647:
                        v1 = (v1 + (v1 + ""))
                    default:
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v6 = (v6 * (v6 | (-(v6))))
                            v6 = v6
                            var v18 uint64 = (^(v5))
                            var v19 int8 = -1
                            v18 += (v5 / (((v18 / ((42) | 1))) | 1))
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 string = ((("foo bar" + (("" + (((v1 + ("hello" + ((v1 + v1) + ("hello" + ((v1 + v1) + ((((v1 + (v1 + v1)) + v1) + v1) + v1)))))) + (v1 + ("a" + (v1 + v1)))) + v1)) + "\n\t")) + "a") + "\n\t")
                    v2 -= v2
                    v23 := GenS1{F0: 100}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v6 = v6
                var v24 float64 = 3.14159
                _ = fi20
                _ = v24
            }
        }
    }
    v6 = v6
    var v25 bool = (-2.25 < ((-2.25 * (1e10 * (3.14159 / 1.5))) + 1e10))
    var v26 int32 = v2
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := GenS1{F0: v5}
        var v29 uint = 221
        switch v1 {
        case "":
            var v30 int32 = v2
            _ = v30
        default:
            v9 += (v9 >> ((875) & 63))
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            switch (v1 + v1) {
            case "":
                var v33 int8 = 1
                _ = v33
            case "foo bar":
                fmt.Println(v26, v1)
            case "\n\t":
                v6 = v6
            }
            v2 -= v26
            _ = fi32
        }
        var v34 uint32 = 65535
        for fi35 := 0; fi35 < 7; fi35++ {
            fmt.Println(fi31)
            _ = fi35
        }
        _ = fi31
        _ = v34
    }
    if (v25 || (true || (v26 == 42))) {
        if v25 {
            for fi36 := 0; fi36 < 20; fi36++ {
                if ((v9 <= v9) || (v25 || v25)) {
                    v37 := GenS1{F0: v5}
                    fmt.Println(v9, v6)
                    _ = v37
                } else {
                    v38 := GenS1{F0: v5}
                    fmt.Println(v9)
                    v26 = v26
                    _ = v38
                }
                _ = fi36
            }
            v39 := []int{int(v6), int(v2)}
            v25 = (!(!(61 < 1)))
            _ = v39
        }
        var v40 int8 = 101
        _ = v40
    }
    v41 := GenS2{F0: 352, F1: (v6 + (v6 - (v6 ^ (v6 * (^(v6)))))), F2: (v9 & (v9 * 42))}
    v42 := GenS2{F0: (v5 ^ v5), F1: v6, F2: uint(v5)}
    for fi43 := 0; fi43 < 10; fi43++ {
        for fi44 := 0; fi44 < 17; fi44++ {
            v45 := v3[2147483648]
            v5 = v5
            v45 = 65535
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}

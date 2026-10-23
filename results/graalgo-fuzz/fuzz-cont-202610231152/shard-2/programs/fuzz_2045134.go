package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
func genFn0() bool {
    v6 := GenS1{F0: 65535}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int32 = 0
        var v9 int32 = (v8 ^ v8)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 0
    fmt.Println(v10)
    v10 += (v10 % ((v10) | 1))
    _ = v10
    _ = v6
    return false
}
func genFn1(p0 int64, p1 string, p2 string) int {
    v11 := GenS1{F0: 0}
    var v12 uint = 383
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int16 = 1
        fmt.Println(fi13, p2, v12)
        p0 = 586
        _ = fi13
        _ = v14
    }
    v12 -= (^(v12))
    v12 -= 83
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -823
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = true
    v3 := make(map[uint32]bool)
    var v4 uint8 = 243
    v5 := make(map[uint32]bool)
    var v15 string = (("hello" + (("a" + ("\n\t" + "hello")) + ("\n\t" + "foo bar"))) + "foo bar")
    v4 = (^(v4))
    v4 = v4
    for fi16 := 0; fi16 < 8; fi16++ {
        v2 = (false && true)
        var v17 uint32 = 809
        var v18 float64 = -2.25
        _ = fi16
        _ = v17
        _ = v18
    }
    delete(v1, 95)
    _ = len(v3)
    delete(v5, 574)
    switch -510 {
    case 127:
        if (42 > -84) {
            if ((!(!(42 != 1))) && (true || (0.1 != 0.0))) {
                v15 = v15
                for fi19 := 0; fi19 < 18; fi19++ {
                    v15 = v15
                    fmt.Println(v4, fi19, v15)
                    _ = fi19
                }
            } else {
                v20 := []int{6, genFn1(42, v15, v15), 870}
                _ = v20
            }
        }
    case -9223372036854775808:
        if v2 {
            var v21 int8 = 0
            switch v15 {
            case "foo bar":
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 int64 = 42
                    _ = fi22
                    _ = v23
                }
            case "hello":
                switch -9223372036854775808 {
                case 127:
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v25 := GenS1{F0: 255}
                        v26 := v1[2147483647]
                        v27 := []int{fi24}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case 1:
                    v28 := v5[784]
                    _ = v28
                case 0:
                    if v2 {
                        v21 = -58
                        var v29 float32 = (3.14159 * 1e10)
                        switch 4294967295 {
                        case 9223372036854775807:
                            var v30 float64 = 3.14159
                            _ = v30
                        }
                        _ = v29
                    }
                }
            }
            var v31 int8 = (v21 % ((1) | 1))
            _ = v21
            _ = v31
        } else {
            if v2 {
                v2 = (0.1 >= 0.0)
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                if ((18446744073709551615 >= 795) || ("foo bar" != v15)) {
                    v33 := GenS1{F0: 1}
                    var v34 bool = ((1 >= 65535) && ((false && (((2147483647 <= 136) || (4294967295 < 2147483647)) && (127 <= 127))) && false))
                    for fi35 := 0; fi35 < 7; fi35++ {
                        v34 = (((42 < -1) && false) && ((!(!v2)) || (-2147483648 > 42)))
                        var v36 bool = v34
                        v4 = 127
                        v37 := GenS2{F0: true, F1: genFn1(42, ("" + (v15 + "\n\t")), ((("hello" + "") + (v15 + v15)) + v15))}
                        v36 = ((!v36) || (v36 || true))
                        fmt.Println(v4, v34, fi32)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = v33
                    _ = v34
                }
                _ = fi32
            }
            for fi38 := 0; fi38 < 2; fi38++ {
                var v39 string = ("foo bar" + "hello")
                v40 := v5[0]
                var v41 bool = (97 == 255)
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        }
    case 42:
        var v42 float32 = 3.14159
        _ = v42
    default:
        var v43 uint16 = 255
        _ = v43
    }
    if (false || v2) {
        v44 := []int{4294967295, -1, 0}
        _ = v44
    } else {
        v4 += v4
    }
    v45 := GenS2{F0: v2, F1: -28}
    v46 := v3[65535]
    fmt.Println(v15)
    v1[2] = ((v15 + "hello") + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}

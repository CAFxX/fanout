package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int8, p1 uint8) int64 {
    p1 = 58
    var v7 bool = true
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 string, p1 uint32) int64 {
    switch ("" + p0) {
    case "foo bar":
        v8 := GenS1{F0: 2}
        _ = v8
    default:
        if false {
            var v9 uint16 = 65535
            if (-8 == 1) {
                if (((1e10 >= 1e10) && false) && (true && false)) {
                    if (42 == 2147483647) {
                        var v10 bool = (false && ((-1 != -2147483648) || ((!(!(!true))) && (false && true))))
                        v11 := GenS1{F0: 255}
                        v12 := GenS1{F0: 18446744073709551615}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        v9 -= 1
                    }
                    v9 = 308
                    var v13 int8 = 1
                    _ = v13
                }
                v14 := []int{1, -2147483648, -87, 127}
                switch p0 {
                case "\n\t":
                    var v15 float64 = -2.25
                    _ = v15
                case "foo bar":
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if (false || (618 > 9223372036854775807)) {
                            if false {
                                p0 = ("a" + "")
                                var v17 int = fi16
                                _ = v17
                            }
                        }
                        switch fi16 {
                        case 566:
                            v14 = append(v14, 128)
                        }
                        _ = fi16
                    }
                default:
                    fmt.Println(v9)
                }
                _ = v14
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    fmt.Println(p0)
                    v19 := GenS1{F0: 0}
                    _ = fi18
                    _ = v19
                }
                if (!(0 <= 450)) {
                    fmt.Println(p1)
                    v20 := []int{32768, 128}
                    _ = v20
                }
                fmt.Println(p0, p1, p0)
            }
            v21 := GenS1{F0: 420}
            _ = v21
            _ = v9
        } else {
            for fi22 := 0; fi22 < 15; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    if ((0.1 <= 1.5) || (false && (p0 != ((("\n\t" + ("hello" + p0)) + "\n\t") + p0)))) {
                        var v24 int8 = 1
                        if true {
                            v25 := []int{0}
                            var v26 int16 = -1
                            _ = v25
                            _ = v26
                        } else {
                            p1 += p1
                            p0 = ((("foo bar" + p0) + "hello") + ("a" + "\n\t"))
                        }
                        _ = v24
                    } else {
                        p1 += (^(p1))
                        switch -9223372036854775808 {
                        case 128:
                            p0 = (p0 + ("hello" + "a"))
                        }
                    }
                    v27 := []int{}
                    _ = len(v27)
                    var v28 uint16 = 65535
                    _ = fi23
                    _ = v27
                    _ = v28
                }
                if (128 < 42) {
                    var v29 int16 = 32767
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 15; fi30++ {
                        p0 = ""
                        var v31 uint16 = 32767
                        _ = fi30
                        _ = v31
                    }
                }
                _ = fi22
            }
        }
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := []int{fi33, fi33, (fi33 / ((fi33) | 1))}
            p0 = p0
            _ = fi33
            _ = v34
        }
        v35 := GenS1{F0: 0}
        _ = fi32
        _ = v35
    }
    if (-2147483648 > 210) {
        fmt.Println(p1, p1)
        if (false || false) {
            var v36 bool = true
            var v37 int64 = -9223372036854775808
            _ = v36
            _ = v37
        } else {
            v38 := []int{127, -574}
            var v39 float64 = (float64(1e10) / (float64(1.5) * (3.14159 / 1.0)))
            var v40 float64 = ((v39 / float64(float32((-2.25 / 1.0)))) - 0.0)
            _ = v38
            _ = v39
            _ = v40
        }
    }
    p1 = p1
    p0 = ("hello" + "hello")
    _ = p0
    _ = p1
    return -35
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 float32 = (0.1 - 1.5)
    var v4 int64 = 42
    var v5 bool = (641 < 539)
    v6 := []string{("" + "a"), "hello", ""}
    var v41 bool = v5
    if (-1 > 100) {
        v5 = (1 >= 255)
    }
    v42 := GenS1{F0: 2}
    v43 := GenS1{F0: 300}
    v44 := v2[568]
    v44 = false
    var v45 string = "\n\t"
    v4 += v4
    v44 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}

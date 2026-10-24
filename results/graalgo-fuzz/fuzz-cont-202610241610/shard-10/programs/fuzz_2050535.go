package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 string; F1 uint; F2 int32 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 string, p1 bool, p2 int16) string {
    v7 := GenS1{F0: 18446744073709551615, F1: -83}
    p0 = "foo bar"
    v8 := GenS1{F0: 472, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 bool) uint32 {
    p0 = true
    v9 := []int{}
    fmt.Println(p0)
    _ = len(v9)
    _ = p0
    _ = v9
    return 4294967295
}

func main() {
    var v1 int64 = 0
    var v2 uint16 = 42
    var v3 uint8 = uint8(v2)
    var v4 uint = 1
    v5 := make(map[uint32]int)
    var v6 int32 = -2147483648
    v1 += (v1 - 9223372036854775807)
    v10 := GenS1{F0: 0, F1: -1}
    var v11 float32 = (float32((1e10 / 1.0)) + (1e10 + 3.14159))
    var v12 int = int(v4)
    v12 = v12
    v13 := GenS1{F0: 18446744073709551615, F1: 127}
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: 0, F1: (v1 | (v1 + v1))}
        for fi16 := 0; fi16 < 2; fi16++ {
            v12 = fi14
            v17 := v5[892]
            v18 := GenS3{F0: false, F1: 1}
            if true {
                _ = len(v5)
                v12 -= v17
            } else {
                fmt.Println(v4, v2)
                fmt.Println(v3, v17, v11)
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    var v19 uint16 = v2
    var v20 uint16 = v2
    v20 = (^(v20))
    if (false && ((((v1 >> (((^(v4))) & 63)) >= (-(v1))) && false) || true)) {
        var v21 uint = v4
        for fi22 := 0; fi22 < 15; fi22++ {
            if false {
                v23 := GenS2{F0: 4294967295, F1: v1}
                for fi24 := 0; fi24 < 4; fi24++ {
                    switch (fi24 + 32768) {
                    case 703:
                        fmt.Println(fi22, v1, v11)
                    }
                    if false {
                        if (uint8(v20) == v3) {
                            var v25 string = (("hello" + "foo bar") + ("hello" + "foo bar"))
                            var v26 int8 = -128
                            _ = v25
                            _ = v26
                        }
                        var v27 uint32 = genFn1((42 >= (v6 & v6)))
                        _ = v27
                    }
                    _ = fi24
                }
                _ = v23
            }
            _ = fi22
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                switch fi29 {
                case -2147483648:
                    var v30 uint8 = 66
                    _ = v30
                case -1:
                    v6 = v6
                default:
                    v21 += v21
                    v31 := GenS4{F0: (("a" + ("foo bar" + ("" + "foo bar"))) + ""), F1: (v21 - (v4 >> ((v21) & 63))), F2: v6}
                    _ = v31
                }
                v21 = 2
                _ = fi29
            }
            if ("\n\t" >= "\n\t") {
                if (2 == 255) {
                    v32 := []int{(fi28 & (v12 / ((int(v3)) | 1))), v12, 127}
                    var v33 string = "hello"
                    var v34 int16 = 0
                    _ = v32
                    _ = v33
                    _ = v34
                }
                v11 -= 1.5
                var v35 uint32 = 211
                _ = v35
            } else {
                var v36 uint64 = 42
                switch ("hello" + "\n\t") {
                case "\n\t":
                    fmt.Println(v2, v2, v19)
                case "a":
                    v6 += -1
                }
                if true {
                    for fi37 := 0; fi37 < 16; fi37++ {
                        v11 -= 0.1
                        _ = fi37
                    }
                    v38 := []int{fi28}
                    _ = v38
                }
                _ = v36
            }
            _ = fi28
        }
        _ = v21
    } else {
        for fi39 := 0; fi39 < 19; fi39++ {
            v40 := GenS5{F0: int64(v4), F1: 0}
            _ = fi39
            _ = v40
        }
        if ((^(v2)) == v19) {
            v12 = (v12 - v12)
            v20 += uint16(v3)
            if (uint16(v4) == (v19 / ((2) | 1))) {
                var v41 string = ""
                if false {
                    switch (v12 & 773) {
                    case -699:
                        v42 := []int{}
                        _ = v42
                    case 1:
                        for fi43 := 0; fi43 < 12; fi43++ {
                            var v44 bool = (1e10 >= -2.25)
                            fmt.Println(v1, v44)
                            v11 = v11
                            _ = fi43
                            _ = v44
                        }
                    }
                    v3 -= (v3 ^ 42)
                    var v45 int32 = (v6 % ((v6) | 1))
                    _ = v45
                }
                v46 := v5[1]
                _ = v41
                _ = v46
            }
        } else {
            fmt.Println(v19, v4, v19)
            if (v3 != (v3 & 61)) {
                v47 := []int{v12, (v12 >> (((v4 + (^(v4)))) & 63)), v12}
                _ = v47
            } else {
                fmt.Println(v3, v1)
                v48 := []int{v12, int(v6)}
                _ = v48
            }
        }
    }
    v49 := v5[65535]
    v50 := GenS1{F0: 1, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}

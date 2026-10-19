package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32) int8 {
    var v4 uint = 960
    if ((false && (100 == 100)) || true) {
        p0 = (^(p0))
    }
    if true {
        if false {
            p0 += (p0 & (p0 % (((^(p0))) | 1)))
            for fi5 := 0; fi5 < 16; fi5++ {
                for fi6 := 0; fi6 < 8; fi6++ {
                    var v7 float64 = 1.5
                    _ = fi6
                    _ = v7
                }
                _ = fi5
            }
        }
        v8 := GenS1{F0: 143, F1: 42}
        _ = v8
    } else {
        v9 := []int{47, -2147483648, -651, 2147483647}
        _ = v9
    }
    switch "hello" {
    case "\n\t":
        var v10 float32 = 1.5
        _ = v10
    default:
        v11 := GenS1{F0: 225, F1: 255}
        _ = v11
    }
    _ = p0
    _ = v4
    return 100
}
func genFn1(p0 string) int {
    v12 := GenS1{F0: 1, F1: 42}
    if ((!(9223372036854775807 == 754)) && true) {
        var v13 uint64 = 1
        var v14 int16 = -1
        if true {
            v13 = (^(v13))
        }
        _ = v13
        _ = v14
    } else {
        switch -327 {
        case 127:
            var v15 string = p0
            _ = v15
        default:
            if false {
                v16 := []int{}
                _ = v16
            } else {
                var v17 float32 = ((3.14159 - 1e10) * (0.1 * float32(float64((float32(-2.25) - 3.14159)))))
                var v18 uint8 = 128
                fmt.Println(v17)
                _ = v17
                _ = v18
            }
        }
    }
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 uint8 = 0
    var v2 string = "foo bar"
    var v3 int64 = 42
    if false {
        for fi19 := 0; fi19 < 8; fi19++ {
            if true {
                v2 = v2
                v20 := []int{(^(fi19)), genFn1(v2)}
                _ = v20
            }
            _ = fi19
        }
    } else {
        var v21 int32 = 0
        v2 = ""
        v22 := []int{2147483647, 42}
        _ = v21
        _ = v22
    }
    v2 = (("a" + "") + v2)
    fmt.Println(v2, v2)
    var v23 int64 = v3
    for fi24 := 0; fi24 < 18; fi24++ {
        v1 += 128
        v25 := GenS2{F0: -1}
        _ = fi24
        _ = v25
    }
    if (!true) {
        var v26 int16 = 1000
        _ = v26
    } else {
        switch (((v2 + (v2 + "foo bar")) + v2) + "") {
        case "a":
            if (0.1 >= 3.14159) {
                if (true || false) {
                    fmt.Println(v2, v23, v2)
                    fmt.Println(v2)
                } else {
                    var v27 int32 = 202
                    if false {
                        if false {
                            var v28 bool = true
                            var v29 bool = true
                            _ = v28
                            _ = v29
                        }
                        v30 := []int{466, genFn1(v2)}
                        _ = v30
                    }
                    _ = v27
                }
            } else {
                var v31 int8 = 0
                if (!false) {
                    v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
                    var v32 uint = 2
                    _ = v32
                }
                _ = v31
            }
        case "\n\t":
            if (!(false || true)) {
                var v33 uint64 = 4294967295
                v34 := []int{genFn1(v2), 1, genFn1("foo bar"), 4294967295}
                var v35 uint16 = 739
                _ = v33
                _ = v34
                _ = v35
            }
        }
        v36 := []int{genFn1((v2 + (v2 + v2))), 127, -1}
        v37 := GenS1{F0: v1, F1: uint16(v1)}
        _ = v36
        _ = v37
    }
    v2 = (("\n\t" + ("\n\t" + "")) + "hello")
    var v38 int = 32767
    v23 += v23
    var v39 float64 = (0.0 * (-2.25 + 0.1))
    for fi40 := 0; fi40 < 16; fi40++ {
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 uint8 = v1
            v39 += v39
            _ = fi41
            _ = v42
        }
        for fi43 := 0; fi43 < 19; fi43++ {
            v44 := GenS1{F0: 127, F1: 0}
            v23 = v3
            var v45 uint8 = (v1 + 127)
            _ = fi43
            _ = v44
            _ = v45
        }
        v3 += 42
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
}

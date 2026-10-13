package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0(p0 uint) uint32 {
    if false {
        v4 := GenS1{F0: 42}
        _ = v4
    } else {
        v5 := []int{}
        var v6 uint16 = 1
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int8 = 100
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    if (-949 >= -1) {
        p0 = (p0 ^ p0)
        v9 := []int{2147483647, 9223372036854775807, 9223372036854775807, -208}
        p0 = p0
        _ = v9
    }
    v10 := GenS1{F0: 1}
    _ = p0
    _ = v10
    return 42
}

func main() {
    v1 := []string{(("" + "foo bar") + "hello"), "", (("\n\t" + "") + "foo bar"), ("a" + "hello")}
    var v2 float64 = 1.5
    var v3 float32 = 0.1
    var v11 int16 = -32768
    var v12 uint16 = 255
    if (v12 == (v12 >> ((480) & 63))) {
        v13 := GenS1{F0: (v11 + v11)}
        var v14 uint8 = 2
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int16 = (v11 / (((v11 / ((-1) | 1))) | 1))
            v18 := GenS1{F0: (v11 | v17)}
            v19 := GenS1{F0: v11}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        if true {
            v3 = 3.14159
            v3 = (v3 - 0.0)
        } else {
            var v20 int16 = v11
            v20 -= v11
            _ = v20
        }
        if (!((("" + "") + "hello") <= "foo bar")) {
            switch (fi15 * (fi15 + fi15)) {
            case 0:
                v12 = v12
            case 9223372036854775807:
                var v21 float64 = v2
                _ = v21
            default:
                var v22 int32 = 1
                _ = v22
            }
        } else {
            var v23 float64 = float64(1e10)
            _ = v23
        }
        _ = fi15
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v2 = (-2.25 - ((v2 - v2) - -2.25))
        _ = fi24
    }
    var v25 int32 = -1
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v25 += -953
    v26 := GenS1{F0: (v11 % ((0) | 1))}
    if ((((("foo bar" + ("\n\t" + "hello")) + "") + "\n\t") + "") < "foo bar") {
        _ = len(v1)
        var v27 uint8 = 52
        _ = v27
    }
    if (!true) {
        for fi28 := 0; fi28 < 10; fi28++ {
            if (v3 < v3) {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                for fi29 := 0; fi29 < 5; fi29++ {
                    v2 -= 1e10
                    var v30 int8 = -128
                    _ = fi29
                    _ = v30
                }
                var v31 int32 = -1
                _ = v31
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v25 = (v25 / ((v25) | 1))
            _ = fi28
        }
        _ = len(v1)
    }
    var v32 int8 = 42
    var v33 int64 = 295
    if false {
        if (!false) {
            var v34 int16 = v11
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v34
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v35 float64 = 0.0
            _ = v35
        }
        v11 += 0
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 int16 = v11
            if true {
                var v38 int = fi36
                _ = v38
            } else {
                v39 := GenS2{F0: false, F1: 283, F2: 1}
                var v40 uint32 = 2147483648
                _ = v39
                _ = v40
            }
            if (!(true && (false && (!(!false))))) {
                if (float64(v3) < v2) {
                    _ = len(v1)
                } else {
                    _ = len(v1)
                    if false {
                        fmt.Println(v12)
                    } else {
                        fmt.Println(v2, v12)
                        var v41 int64 = v33
                        var v42 int = fi36
                        _ = v41
                        _ = v42
                    }
                    v33 = v33
                }
                _ = len(v1)
            }
            _ = fi36
            _ = v37
        }
    }
    var v43 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v43
}

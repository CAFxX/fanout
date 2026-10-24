package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64, p1 int8) uint16 {
    var v5 int = 9223372036854775807
    if ((((0.1 - (0.1 * (float64(((1e10 * 0.0) / (((((((0.1 - 1.5) + (-2.25 * 1.5)) * 1e10) + 3.14159) + 0.1) - 3.14159) - (((0.1 + (float32(1.5) * (float32(1.5) - float32(0.1)))) * 1.5) * -2.25)))) / (float64(-2.25) + 0.0)))) / 0.1) + 1.5) < 0.1) {
        v5 += 127
    }
    v6 := GenS1{F0: ((false || true) || (42 == 81))}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = ((1.5 - float32(1e10)) - (3.14159 * float32(float64(float32(1.5)))))
        p1 = p1
        _ = fi7
        _ = v8
    }
    var v9 int32 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return 255
}

func main() {
    var v1 uint8 = 62
    var v2 int32 = 789
    var v3 string = "hello"
    var v4 bool = (!(false && (!false)))
    if (-32768 > -32768) {
        v10 := []int{9223372036854775807, 2147483647}
        var v11 uint32 = 1
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v10
        _ = v11
    }
    v2 = -508
    v12 := []int{4294967295}
    if len(v12) > 0 {
        v12[0] = 1
    }
    var v13 int8 = 1
    if len(v12) > 0 {
        v12[0] = int(v2)
    }
    fmt.Println(v2)
    v14 := GenS1{F0: true}
    v1 -= 0
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            switch v3 {
            case "a":
                var v17 bool = true
                _ = v17
            case "foo bar":
                fmt.Println(fi15)
            case "hello":
                if false {
                    var v18 int8 = v13
                    if (v4 || (!((v18 ^ v18) >= 1))) {
                        if v4 {
                            v12 = append(v12, fi15)
                            v1 -= v1
                        } else {
                            var v19 int16 = 32767
                            var v20 uint32 = 1
                            if v4 {
                                if false {
                                    fmt.Println(v2, v13)
                                    if (true || (!true)) {
                                        var v21 int32 = 0
                                        if (-1 >= 167) {
                                            if v4 {
                                                if v4 {
                                                    if true {
                                                        v22 := GenS1{F0: (!true)}
                                                        _ = v22
                                                    }
                                                }
                                                v4 = true
                                            }
                                            var v23 int32 = -1
                                            var v24 int = (fi16 - fi15)
                                            _ = v23
                                            _ = v24
                                        } else {
                                            fmt.Println(v1, v3, v13)
                                        }
                                        var v25 int8 = v13
                                        var v26 float64 = (0.0 * (0.1 / 0.1))
                                        var v27 bool = (32767 < 0)
                                        _ = v21
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                } else {
                                    v28 := GenS1{F0: (0 <= 42)}
                                    _ = v28
                                }
                                var v29 uint16 = 42
                                v2 = (^(v2))
                                _ = v29
                            }
                            _ = v19
                            _ = v20
                        }
                    } else {
                        var v30 uint = 981
                        _ = v30
                    }
                    _ = v18
                }
            }
            _ = fi16
        }
        switch v3 {
        case "":
            switch (fi15 << ((32) & 63)) {
            case -874:
                v31 := GenS1{F0: false}
                _ = v31
            case 1:
                v1 = (v1 / ((v1) | 1))
            }
        case "\n\t":
            if ((-(v13)) > v13) {
                if v4 {
                    v1 = (v1 % ((v1) | 1))
                } else {
                    v4 = v4
                    v3 = v3
                }
            }
        }
        _ = fi15
    }
    v32 := GenS1{F0: (42 <= -232)}
    var v33 int32 = (^(v2))
    var v34 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 bool) uint {
    p0 = false
    var v4 float32 = 3.14159
    _ = p0
    _ = v4
    return 42
}
func genFn1() int8 {
    if false {
        if (!false) {
            var v5 float32 = -2.25
            var v6 bool = false
            _ = v5
            _ = v6
        } else {
            var v7 float64 = ((0.0 / ((float64(-2.25) + -2.25) / (0.0 + (0.0 - -2.25)))) * -2.25)
            if true {
                for fi8 := 0; fi8 < 18; fi8++ {
                    fmt.Println(fi8, v7)
                    _ = fi8
                }
                v9 := []int{-2147483648, 115, -343, -2147483648}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = v9
            } else {
                v10 := []int{-737, 32767, 32767}
                v11 := GenS1{F0: -1}
                switch -1 {
                case 42:
                    if (0 != 2147483647) {
                        fmt.Println(v7)
                    } else {
                        var v12 bool = ((("" + ("foo bar" + "a")) != "foo bar") || false)
                        _ = v12
                    }
                default:
                    if (true || (127 < 0)) {
                        v13 := GenS1{F0: 127}
                        v14 := GenS1{F0: -46}
                        _ = v13
                        _ = v14
                    }
                }
                _ = v10
                _ = v11
            }
            _ = v7
        }
    }
    if true {
        var v15 uint32 = 922
        _ = v15
    }
    if false {
        if false {
            var v16 int32 = -451
            v17 := []int{-9223372036854775808, -9223372036854775808, -4, 9223372036854775807}
            fmt.Println(v16)
            _ = v16
            _ = v17
        } else {
            var v18 uint8 = 177
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint16 = 182
            fmt.Println(v21)
            v22 := GenS2{F0: "\n\t"}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    return 114
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -1
    var v3 int = 293
    var v23 uint = genFn0(true)
    v24 := GenS3{F0: 65535}
    v3 += v3
    v25 := GenS4{F0: 231}
    v26 := GenS5{F0: 0, F1: 1, F2: 2147483648}
    for fi27 := 0; fi27 < 16; fi27++ {
        v28 := GenS4{F0: 0}
        var v29 int8 = 127
        _ = fi27
        _ = v28
        _ = v29
    }
    switch "hello" {
    case "":
        for fi30 := 0; fi30 < 3; fi30++ {
            if false {
                v31 := []int{(v3 + (v3 << (((^(v23))) & 63))), (-(v3))}
                v32 := GenS3{F0: 39}
                _ = v31
                _ = v32
            }
            if false {
                switch (("\n\t" + "a") + ("a" + "foo bar")) {
                case "hello":
                    var v33 int8 = -1
                    _ = v33
                }
                for fi34 := 0; fi34 < 20; fi34++ {
                    var v35 uint = 0
                    _ = fi34
                    _ = v35
                }
            } else {
                if true {
                    _ = len(v1)
                } else {
                    var v36 string = "a"
                    for fi37 := 0; fi37 < 9; fi37++ {
                        v38 := []int{fi37}
                        if ((false && (!(!(false || ((true && true) || true))))) || (!(754 > uint64(v23)))) {
                            v39 := GenS6{F0: 0}
                            _ = v39
                        } else {
                            v36 = v36
                        }
                        _ = fi37
                        _ = v38
                    }
                    fmt.Println(fi30, v23)
                    _ = v36
                }
                var v40 int64 = 0
                v41 := GenS1{F0: -1}
                _ = v40
                _ = v41
            }
            if (false || false) {
                v42 := []int{v3, (fi30 % ((32768) | 1)), v3}
                for fi43 := 0; fi43 < 1; fi43++ {
                    v2 = v2
                    v2 = -591
                    if len(v42) >= 2 {
                        v42 = v42[1:len(v42)]
                    }
                    _ = fi43
                }
                _ = v42
            } else {
                fmt.Println(v23)
                var v44 int = v3
                _ = v44
            }
            _ = fi30
        }
    case "\n\t":
        var v45 uint32 = 664
        _ = v45
    case "foo bar":
        var v46 int32 = 2147483647
        _ = v46
    }
    var v47 uint32 = 2
    var v48 uint32 = uint32(v23)
    if (-9223372036854775808 >= 134) {
        v49 := v1[(v47 << ((genFn0(("a" > (("\n\t" + "\n\t") + "a")))) & 63))]
        _ = v49
    }
    var v50 float64 = (1.5 - (((((-2.25 - (1.5 / (1e10 - float64((0.0 + -2.25))))) - (1e10 - ((3.14159 / ((0.1 - 0.1) + 1e10)) + 0.1))) + -2.25) - 1e10) - ((((0.0 * 0.1) * -2.25) * -2.25) / (3.14159 / -2.25))))
    v47 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v47
    _ = v48
    _ = v50
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 bool, p1 int16) int32 {
    var v4 int64 = 42
    switch 522 {
    case 32768:
        fmt.Println(v4, p1)
    default:
        v5 := []int{32767, 32767, 243, -1}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return -996
}

func main() {
    var v1 int8 = -64
    var v2 uint8 = 128
    var v3 int32 = 0
    v1 += 127
    var v6 int32 = (-(v3))
    v1 = v1
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v2, fi7, v3)
        v8 := []int{(-(fi7)), fi7, fi7, 815}
        v2 = v2
        _ = fi7
        _ = v8
    }
    v1 -= v1
    fmt.Println(v2, v1, v6)
    if false {
        var v9 int32 = v6
        var v10 float64 = 1.5
        _ = v10
        _ = v9
    } else {
        if (((0.0 / 1.5) * 1.5) != 0.0) {
            fmt.Println(v2, v1, v2)
            v11 := GenS1{F0: 42, F1: v6, F2: 1}
            if false {
                var v12 int64 = 109
                v13 := GenS1{F0: -755, F1: int32(v12), F2: 2}
                v14 := []int{32767}
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v11
        } else {
            v15 := []int{42}
            switch 127 {
            case -2147483648:
                v2 -= (v2 % (((v2 << ((772) & 63))) | 1))
            default:
                if false {
                    if (!(4294967295 >= 42)) {
                        _ = len(v15)
                        for fi16 := 0; fi16 < 8; fi16++ {
                            v3 = (v6 | -1)
                            for fi17 := 0; fi17 < 4; fi17++ {
                                v18 := GenS1{F0: int(v3), F1: v3, F2: 2}
                                v19 := GenS1{F0: fi16, F1: (v3 ^ genFn0((("a" < ("hello" + ("foo bar" + "foo bar"))) && ((593 >= 4294967295) || (false || ((v2 >> ((18446744073709551615) & 63)) <= v2)))), 42)), F2: 4294967295}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = fi16
                        }
                        v6 = 2147483647
                    }
                    var v20 int64 = 9223372036854775807
                    var v21 bool = false
                    _ = v20
                    _ = v21
                }
            }
            switch 0 {
            case -172:
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            case 9223372036854775807:
                fmt.Println(v3, v1)
            }
            _ = v15
        }
    }
    if true {
        if false {
            for fi22 := 0; fi22 < 5; fi22++ {
                v3 += (v6 - -1)
                switch (("foo bar" + "hello") + "") {
                case "\n\t":
                    var v23 string = ("\n\t" + "hello")
                    _ = v23
                default:
                    v24 := GenS1{F0: (^(fi22)), F1: (v6 / (((v6 & (v3 / ((v6) | 1)))) | 1)), F2: 0}
                    _ = v24
                }
                _ = fi22
            }
            v1 = 1
            var v25 float32 = 1e10
            _ = v25
        } else {
            if true {
                var v26 int16 = -32768
                for fi27 := 0; fi27 < 1; fi27++ {
                    if ((v3 % ((v6) | 1)) >= v6) {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            var v29 string = (("hello" + "a") + "\n\t")
                            v6 = (v6 + 0)
                            v3 += v6
                            v26 = -32768
                            v26 = v26
                            _ = fi28
                            _ = v29
                        }
                        var v30 int32 = 2147483647
                        for fi31 := 0; fi31 < 13; fi31++ {
                            v6 = (v3 ^ v30)
                            _ = fi31
                        }
                        _ = v30
                    }
                    v6 = int32(v26)
                    var v32 int32 = v6
                    _ = fi27
                    _ = v32
                }
                _ = v26
            }
            v2 -= 127
        }
    } else {
        var v33 float32 = 3.14159
        v1 = (v1 / (((^(v1))) | 1))
        var v34 bool = true
        _ = v33
        _ = v34
    }
    var v35 int = 9223372036854775807
    fmt.Println(v35, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v6
}

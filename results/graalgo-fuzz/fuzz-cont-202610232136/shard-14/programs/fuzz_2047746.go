package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }

func main() {
    var v1 int16 = -700
    v2 := []string{"hello"}
    var v3 int64 = -561
    var v4 int64 = 1
    _ = len(v2)
    var v5 int = -1
    v6 := GenS1{F0: (!false), F1: 42, F2: -2147483648}
    switch -1 {
    case 118:
        _ = len(v2)
    case -257:
        v2 = append(v2, "")
    case 42:
        switch 42 {
        case -614:
            if false {
                v4 = v4
                v5 = -9223372036854775808
            }
        case 0:
            fmt.Println(v4, v1, v4)
        case -864:
            if false {
                var v7 uint = 255
                _ = v7
            }
        default:
            var v8 uint = 613
            _ = v8
        }
    default:
        v9 := GenS1{F0: (((!false) || true) && ((!true) && false)), F1: 1, F2: 2147483647}
        _ = v9
    }
    v4 = v4
    if len(v2) > 0 {
        v2[0] = ("\n\t" + (("" + (("hello" + "\n\t") + "hello")) + "\n\t"))
    }
    switch 4294967295 {
    case -205:
        if false {
            if true {
                fmt.Println(v3, v1)
                v3 = (v3 | (v3 >> ((65535) & 63)))
                if ((v5 / ((v5) | 1)) < (^(v5))) {
                    v10 := GenS1{F0: false, F1: 4294967295, F2: 875}
                    fmt.Println(v4, v3)
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                _ = len(v2)
                _ = fi11
            }
            v4 -= v3
        }
    default:
        for fi12 := 0; fi12 < 11; fi12++ {
            if ((-(v5)) <= fi12) {
                if true {
                    v4 -= -693
                    if true {
                        if (18446744073709551615 <= 2) {
                            v13 := GenS1{F0: (42 != 1), F1: 2, F2: -602}
                            for fi14 := 0; fi14 < 11; fi14++ {
                                var v15 int16 = -157
                                v4 = (v4 + (-(v4)))
                                _ = len(v2)
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        }
                        if false {
                            var v16 uint = 42
                            var v17 uint = 4294967295
                            var v18 uint16 = 42
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            var v19 uint16 = 42
                            var v20 uint = 255
                            for fi21 := 0; fi21 < 5; fi21++ {
                                if len(v2) > 0 {
                                    v2[0] = ("a" + "\n\t")
                                }
                                _ = fi21
                            }
                            _ = v19
                            _ = v20
                        }
                    }
                } else {
                    if (!(true && true)) {
                        v22 := GenS1{F0: (!(true && (false && true))), F1: 255, F2: -1}
                        v1 = (v1 | 42)
                        v2 = append(v2, "foo bar")
                        _ = v22
                    } else {
                        var v23 bool = ("" != (("a" + ("foo bar" + "hello")) + "foo bar"))
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v25 := GenS2{F0: 18446744073709551615, F1: -1, F2: (^(v1))}
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    switch 749 {
                    case 32768:
                        _ = len(v2)
                    case 4294967295:
                        if ((^(v4)) <= v3) {
                            v26 := GenS2{F0: 42, F1: -499, F2: v1}
                            _ = v26
                        } else {
                            var v27 uint16 = 2
                            _ = v27
                        }
                    default:
                        v5 = (v5 * (fi12 * v5))
                    }
                }
            } else {
                v4 = (v4 - (v4 ^ (v3 >> ((201) & 63))))
                var v28 int8 = 127
                _ = v28
            }
            _ = fi12
        }
    }
    fmt.Println(v1, v1)
    v3 = (v4 & (v4 + v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

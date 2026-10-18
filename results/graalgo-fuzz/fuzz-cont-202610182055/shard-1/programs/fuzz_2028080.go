package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint64 = 555
    var v2 int = 32767
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 int32 = 0
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 4294967295}
            if true {
                fmt.Println(fi7, v2)
            } else {
                v9 := GenS1{F0: 0}
                v5 = v5
                _ = v9
            }
            v10 := v4["\n\t"]
            _ = fi7
            _ = v10
            _ = v8
        }
        _ = fi6
    }
    switch "foo bar" {
    case "a":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint16 = 577
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v2, v3)
    var v13 int = v2
    v3 = false
    if (("\n\t" <= "foo bar") || v3) {
        v14 := GenS1{F0: 0}
        if (!((v5 * (v5 & (v5 * 42))) > (-(v5)))) {
            switch (v13 >> ((18446744073709551615) & 63)) {
            case 42:
                v15 := GenS2{F0: 739, F1: 844}
                _ = v15
            case 1:
                if v3 {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v17 := v4["\n\t"]
                        if v17 {
                            v18 := GenS1{F0: 0}
                            _ = v18
                        } else {
                            v19 := []int{(v13 << ((560) & 63))}
                            v20 := GenS2{F0: v1, F1: 65535}
                            v13 += (fi16 - (v13 - v13))
                            _ = v19
                            _ = v20
                        }
                        _ = fi16
                        _ = v17
                    }
                    var v21 bool = ((v13 * 42) <= (-(v13)))
                    fmt.Println(v5, v5, v21)
                    _ = v21
                }
            case -2147483648:
                var v22 float64 = float64(0.0)
                _ = v22
            default:
                var v23 uint64 = v1
                _ = v23
            }
        } else {
            if true {
                v24 := []int{(-(v13)), 4294967295, (v2 - v2)}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 5; fi25++ {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 uint = 255
                        v5 += 633
                        switch 42 {
                        case 42:
                            if ((9 > 0) && v3) {
                                fmt.Println(v1)
                                v1 -= v1
                            }
                        case -499:
                            v28 := []int{}
                            _ = v28
                        case 128:
                            var v29 int64 = -1
                            fmt.Println(v1, v2, fi25)
                            fmt.Println(v1)
                            _ = v29
                        default:
                            if false {
                                v30 := []int{(-(fi26)), fi25, (fi25 >> ((v27) & 63)), 128}
                                _ = v30
                            } else {
                                _ = len(v4)
                                fmt.Println(v27)
                            }
                        }
                        v31 := GenS1{F0: 2}
                        _ = fi26
                        _ = v27
                        _ = v31
                    }
                    _ = fi25
                }
            }
            v1 += (v1 | (^(v1)))
        }
        _ = v14
    }
    var v32 uint8 = 2
    v3 = (false && (!(!(65535 > 18446744073709551615))))
    v1 -= v1
    var v33 int = (v2 - 239)
    v32 = v32
    var v34 uint16 = 42
    var v35 uint = 332
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}

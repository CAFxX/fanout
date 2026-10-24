package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) int32 {
    var v6 int16 = 32767
    var v7 int8 = -1
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 1
    var v2 bool = false
    var v3 float32 = 1e10
    v4 := []int{}
    var v5 float32 = float32(0.0)
    v8 := GenS1{F0: 2147483647}
    for fi9 := 0; fi9 < 9; fi9++ {
        if false {
            if v2 {
                v3 = v3
                switch "foo bar" {
                case "a":
                    switch ("" + "\n\t") {
                    case "foo bar":
                        if true {
                            if v2 {
                                v10 := GenS1{F0: 592}
                                if ("" > ("" + (("foo bar" + "foo bar") + "\n\t"))) {
                                    var v11 uint8 = 2
                                    _ = v11
                                } else {
                                    for fi12 := 0; fi12 < 15; fi12++ {
                                        v4 = append(v4, (fi9 / (((fi9 >> ((751) & 63))) | 1)))
                                        var v13 int64 = -681
                                        switch fi9 {
                                        case 32768:
                                            v14 := GenS2{F0: 4294967295}
                                            _ = v14
                                        }
                                        _ = fi12
                                        _ = v13
                                    }
                                    v5 += v5
                                }
                                v3 = (((-2.25 * 1.5) + 0.1) / v5)
                                _ = v10
                            } else {
                                if v2 {
                                    var v15 int8 = 16
                                    for fi16 := 0; fi16 < 11; fi16++ {
                                        v17 := GenS2{F0: 42}
                                        fmt.Println(fi9, v1)
                                        v5 = (0.0 - (((v3 * v5) + (v5 + float32((-2.25 + -2.25)))) + float32(1e10)))
                                        _ = fi16
                                        _ = v17
                                    }
                                    _ = v15
                                }
                            }
                            v1 -= 87
                            v3 = v3
                        }
                    }
                case "foo bar":
                    v2 = ((1 != 42) || (false || v2))
                case "\n\t":
                    v3 = float32((3.14159 + 0.0))
                }
            } else {
                var v18 uint = 4294967295
                _ = v18
            }
            v5 = v3
            v5 -= 3.14159
        } else {
            var v19 string = "a"
            for fi20 := 0; fi20 < 12; fi20++ {
                if len(v4) > 0 {
                    v4[0] = (fi9 - (fi20 - fi9))
                }
                _ = fi20
            }
            v1 = v1
            _ = v19
        }
        _ = fi9
    }
    v3 -= 1e10
    if v2 {
        if (!(914 >= 549)) {
            v5 += v5
            fmt.Println(v3)
        }
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 bool = v2
        switch fi21 {
        case -883:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 32768:
            if len(v4) > 0 {
                v4[0] = (fi21 + fi21)
            }
        case -439:
            v4 = append(v4, fi21)
        }
        _ = fi21
        _ = v22
    }
    if v2 {
        v23 := GenS2{F0: 0}
        v5 = ((3.14159 - (0.1 / 0.1)) - v3)
        _ = v23
    }
    switch 2147483647 {
    case -1:
        v24 := GenS2{F0: 2147483648}
        _ = v24
    }
    var v25 uint = 378
    v4 = append(v4, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

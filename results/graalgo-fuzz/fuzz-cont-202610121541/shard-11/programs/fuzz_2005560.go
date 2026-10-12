package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint64) int8 {
    var v4 int = 101
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 42}
            switch "a" {
            case "foo bar":
                if (((799 >= (p1 % ((p1) | 1))) || true) && (1 == 1)) {
                    var v8 uint = 255
                    _ = v8
                }
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    fmt.Println(p0, p1)
    var v9 int = (^(v4))
    v9 -= v9
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return int8(p0)
}

func main() {
    var v1 uint8 = 0
    var v2 string = "foo bar"
    var v3 float64 = 1.5
    if (-366 == 435) {
        v10 := GenS1{F0: 32767}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 uint64 = 624
                _ = fi12
                _ = v13
            }
            v14 := []int{fi11, (fi11 & fi11), (-(fi11))}
            if true {
                if false {
                    v14 = append(v14, (-(fi11)))
                    if (1 <= -394) {
                        v1 -= (^(v1))
                        if len(v14) > 0 {
                            v14[0] = 506
                        }
                        v15 := GenS2{F0: 42, F1: (fi11 - (fi11 / ((32767) | 1)))}
                        _ = v15
                    } else {
                        if (false || false) {
                            for fi16 := 0; fi16 < 15; fi16++ {
                                v1 = v1
                                if (v2 <= v2) {
                                    var v17 int8 = 100
                                    v14 = append(v14, fi11)
                                    _ = v17
                                } else {
                                    v18 := GenS3{F0: 32767}
                                    v14 = append(v14, -700)
                                    _ = v18
                                }
                                _ = len(v14)
                                _ = fi16
                            }
                        }
                    }
                } else {
                    v3 -= v3
                    for fi19 := 0; fi19 < 16; fi19++ {
                        fmt.Println(v3)
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = fi19
                    }
                }
                var v20 int16 = -182
                var v21 int = (fi11 % (((fi11 % (((fi11 / ((-1) | 1))) | 1))) | 1))
                _ = v20
                _ = v21
            }
            _ = fi11
            _ = v14
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v2, v2, v1)
            var v23 int64 = 1
            var v24 int8 = 1
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 int32 = -1
        _ = v25
    }
    if ("" <= v2) {
        v2 = v2
    }
    v2 = v2
    v26 := GenS2{F0: 42, F1: -9223372036854775808}
    var v27 int8 = 10
    v28 := GenS4{F0: -754}
    v29 := []int{-574}
    var v30 int = 2147483647
    for fi31 := 0; fi31 < 12; fi31++ {
        v32 := GenS1{F0: -1}
        for fi33 := 0; fi33 < 10; fi33++ {
            v27 += 42
            var v34 int16 = -115
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}

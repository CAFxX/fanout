package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0() bool {
    switch 1 {
    case -423:
        var v4 float32 = 1e10
        _ = v4
    default:
        for fi5 := 0; fi5 < 3; fi5++ {
            if true {
                var v6 uint8 = 135
                v6 -= 2
                var v7 uint64 = 0
                _ = v6
                _ = v7
            } else {
                var v8 float32 = 3.14159
                v8 = (float32(1.5) + v8)
                _ = v8
            }
            fmt.Println(fi5)
            _ = fi5
        }
    }
    var v9 string = ("\n\t" + "hello")
    v10 := GenS1{F0: v9, F1: 1, F2: 1}
    if (false || false) {
        fmt.Println(v9)
    }
    v9 = "hello"
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 int8 = 58
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v1 -= (v1 << ((18446744073709551615) & 63))
    var v11 int16 = 1
    v11 = v11
    if (!false) {
        v1 -= 120
        v12 := []int{2147483647, -893, -49}
        v13 := GenS1{F0: (("\n\t" + "hello") + ("foo bar" + (("\n\t" + "a") + ("a" + ("hello" + ""))))), F1: 1, F2: 18446744073709551615}
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: ("a" + ""), F1: 1, F2: 1}
        switch "foo bar" {
        case "a":
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v11 = v11
                    v17 := GenS1{F0: "a", F1: 4294967295, F2: 42}
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 1; fi18++ {
                    v11 += 697
                    fmt.Println(v11)
                    _ = fi18
                }
                _ = fi15
            }
        default:
            var v19 float64 = (1.5 + (1.5 / 1.0))
            _ = v19
        }
        _ = v14
    }
    if true {
        var v20 bool = (1 <= 306)
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 13; fi21++ {
            if ((fi21 - fi21) != -590) {
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := []int{fi21}
                    v11 += (v11 / (((^(v11))) | 1))
                    v1 = -128
                    if false {
                        _ = len(v3)
                        var v24 int16 = v11
                        _ = v24
                    } else {
                        v25 := v2[("hello" + (("hello" + ("\n\t" + "")) + ("foo bar" + ("hello" + "foo bar"))))]
                        fmt.Println(v1, fi21, v1)
                        if ((!(791 != 169)) || true) {
                            _ = len(v23)
                            v26 := v2["foo bar"]
                            _ = v26
                        }
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
            } else {
                v27 := GenS1{F0: "foo bar", F1: 0, F2: 0}
                var v28 int16 = -63
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 int8 = (v1 / (((v1 % (((v1 % (((v1 + (v1 << ((65535) & 63)))) | 1))) | 1))) | 1))
                    v31 := GenS2{F0: 536, F1: 1}
                    var v32 float64 = 1.5
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = v27
                _ = v28
            }
            _ = fi21
        }
    }
    v11 += 42
    for fi33 := 0; fi33 < 20; fi33++ {
        v34 := []int{(fi33 / ((fi33) | 1)), fi33}
        var v35 uint = 2
        v36 := GenS1{F0: "a", F1: 0, F2: 18446744073709551615}
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}

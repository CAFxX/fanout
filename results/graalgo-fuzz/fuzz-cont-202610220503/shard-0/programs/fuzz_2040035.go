package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint8, p1 string, p2 uint32) uint32 {
    v4 := GenS1{F0: -40, F1: 255, F2: 4294967295}
    if true {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 float64 = 3.14159
            var v7 int = (^(fi5))
            var v8 uint = 18446744073709551615
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            if true {
                v10 := GenS2{F0: 785}
                v11 := GenS2{F0: 1}
                switch ((p1 + "hello") + "a") {
                case "":
                    v12 := []int{-174}
                    _ = v12
                case "hello":
                    if (-125 != -99) {
                        v13 := GenS3{F0: 255, F1: 1, F2: 65535}
                        v14 := GenS3{F0: 2, F1: 0, F2: 255}
                        _ = v13
                        _ = v14
                    }
                }
                _ = v10
                _ = v11
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), ("a" + ("hello" + "")), ""}
    v2 := make(map[uint32]bool)
    var v3 int16 = 0
    v3 -= (^(v3))
    var v15 uint = 42
    v16 := GenS1{F0: 1, F1: (^(v15)), F2: (v15 * (v15 + 2))}
    var v17 uint32 = 2147483647
    if (!true) {
        fmt.Println(v3)
        fmt.Println(v3, v3)
    }
    v3 = v3
    delete(v2, 4294967295)
    v15 += v15
    for fi18 := 0; fi18 < 4; fi18++ {
        if false {
            v17 = v17
            if (0 < 9223372036854775807) {
                _ = len(v2)
                if (65535 == genFn0(42, "a", v17)) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if (false && true) {
                            if ((v3 > 1000) && false) {
                                fmt.Println(fi18, v15)
                                if (v17 <= v17) {
                                    var v20 int32 = 1
                                    v1 = append(v1, "foo bar")
                                    _ = len(v2)
                                    _ = v20
                                }
                            }
                            v17 -= (v17 | (v17 % ((2) | 1)))
                            var v21 int32 = 0
                            _ = v21
                        } else {
                            v22 := GenS1{F0: 1, F1: (v15 * (v15 / ((570) | 1))), F2: (^(v15))}
                            v15 = (^(v15))
                            var v23 int64 = 9223372036854775807
                            _ = v22
                            _ = v23
                        }
                        var v24 uint16 = 338
                        _ = fi19
                        _ = v24
                    }
                }
            }
            v25 := GenS4{F0: "a", F1: (^(v15))}
            _ = v25
        }
        var v26 int16 = v3
        _ = fi18
        _ = v26
    }
    var v27 int32 = -1
    var v28 uint32 = (v17 | genFn0(208, "hello", genFn0(127, ("foo bar" + ("\n\t" + "\n\t")), (v17 & (v17 & v17)))))
    v28 -= (v17 & 833)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}

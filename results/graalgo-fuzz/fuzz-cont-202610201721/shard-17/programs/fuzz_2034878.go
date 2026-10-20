package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8, p1 uint32, p2 int16) int32 {
    var v5 int8 = p0
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 16; fi6++ {
            p2 += (p2 % (((p2 % ((p2) | 1))) | 1))
            _ = fi6
        }
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(p0)
            _ = fi8
        }
        var v9 uint64 = 18446744073709551615
        _ = fi7
        _ = v9
    }
    var v10 int16 = (p2 >> ((18446744073709551615) & 63))
    v11 := GenS1{F0: 428, F1: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v5
    return 698
}
func genFn1(p0 int, p1 string, p2 int32) string {
    if (p0 != (p0 * p0)) {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := []int{(fi12 - fi12), (p0 | (fi12 % ((-751) | 1))), (^(p0))}
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        p1 = "hello"
        var v15 uint32 = 2147483647
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint = 757
    v2 := make(map[string]uint32)
    var v3 uint8 = 2
    var v4 int8 = int8(v1)
    var v16 int64 = 0
    fmt.Println(v3)
    v17 := GenS1{F0: 32767, F1: v3}
    fmt.Println(v16, v4)
    var v18 int8 = v4
    v1 -= v1
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int32 = int32(v3)
        _ = fi19
        _ = v20
    }
    switch ("hello" + (("hello" + ("hello" + ("\n\t" + "foo bar"))) + (("a" + "foo bar") + "foo bar"))) {
    case "":
        for fi21 := 0; fi21 < 4; fi21++ {
            if (0.0 >= (0.0 + 0.0)) {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 float64 = 0.1
                    _ = fi22
                    _ = v23
                }
                var v24 int16 = 1
                if true {
                    v4 += (v18 & (v18 - (v4 - (v18 % (((-(v18))) | 1)))))
                    var v25 float64 = ((1e10 + -2.25) * (1.5 * -2.25))
                    v26 := GenS1{F0: 195, F1: (v3 % ((73) | 1))}
                    _ = v25
                    _ = v26
                } else {
                    if (v1 != (v1 << ((v1) & 63))) {
                        for fi27 := 0; fi27 < 15; fi27++ {
                            if ((3.14159 + 0.0) < (0.0 - (0.1 - -2.25))) {
                                var v28 int16 = (v24 | v24)
                                v1 += (^(v1))
                                v3 += (v3 + (v3 | v3))
                                _ = v28
                            } else {
                                v29 := GenS2{F0: 547, F1: v24}
                                _ = v29
                            }
                            var v30 int32 = -793
                            var v31 uint8 = (v3 | 254)
                            _ = fi27
                            _ = v30
                            _ = v31
                        }
                    }
                }
                _ = v24
            }
            if true {
                v32 := GenS3{F0: false, F1: 76}
                v33 := []int{}
                if (!(!true)) {
                    v16 = int64(v3)
                    if len(v33) > 0 {
                        v33[0] = 2147483647
                    }
                } else {
                    var v34 int16 = -32768
                    _ = v34
                }
                _ = v32
                _ = v33
            }
            _ = fi21
        }
    }
    v35 := GenS2{F0: 4294967295, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}

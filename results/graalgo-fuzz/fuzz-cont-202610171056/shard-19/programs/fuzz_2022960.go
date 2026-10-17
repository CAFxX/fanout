package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 bool, p2 uint32) int32 {
    switch ("\n\t" + "hello") {
    case "a":
        p2 -= (p2 >> (((^(p0))) & 63))
    default:
        v7 := []int{4294967295, 402, -9223372036854775808}
        _ = v7
    }
    p0 -= (^(p0))
    var v8 bool = p1
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 130
}

func main() {
    v1 := []uint32{65535, 65535}
    var v2 int32 = 1
    var v3 uint = 4294967295
    var v4 int32 = -2147483648
    var v5 uint32 = 437
    v6 := make(map[uint32]string)
    v10 := GenS1{F0: 932}
    v4 += v4
    var v11 int32 = (v2 * (v2 + v4))
    fmt.Println(v4, v4)
    if false {
        v6[v5] = "\n\t"
    }
    v3 = v3
    v4 = genFn0((v3 % ((v3) | 1)), (!false), (^(v5)))
    v5 += v5
    if len(v1) > 0 {
        v1[0] = v5
    }
    v4 = (v2 * (v11 % ((int32(v5)) | 1)))
    if ((368 >= (^(v5))) && false) {
        var v12 int16 = -607
        if (660 == 42) {
            if (-2.25 > 3.14159) {
                fmt.Println(v3, v3, v4)
                if (("foo bar" + "foo bar") < ("a" + (("\n\t" + "foo bar") + "\n\t"))) {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 bool = true
                        var v15 uint32 = 732
                        v2 = (v11 | genFn0(v3, (2 <= 2), 129))
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    v3 += v3
                    for fi16 := 0; fi16 < 7; fi16++ {
                        var v17 int64 = 1
                        v18 := GenS2{F0: 255}
                        for fi19 := 0; fi19 < 19; fi19++ {
                            fmt.Println(v11, v17)
                            v12 -= v12
                            var v20 uint64 = 42
                            switch fi16 {
                            case -2147483648:
                                v21 := GenS1{F0: v17}
                                _ = v21
                            case -41:
                                var v22 bool = true
                                _ = v22
                            }
                            if len(v1) > 0 {
                                v1[0] = v5
                            }
                            _ = fi19
                            _ = v20
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v6[v5] = "hello"
                }
                fmt.Println(v11)
            } else {
                v12 += 1000
                var v23 int64 = -1
                var v24 float64 = (0.0 - 0.0)
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 10; fi25++ {
                if true {
                    v5 += (v5 * v5)
                    var v26 int64 = -76
                    _ = v26
                }
                _ = fi25
            }
        } else {
            var v27 uint16 = 608
            var v28 int = 127
            _ = v27
            _ = v28
        }
        _ = v12
    }
    fmt.Println(v3)
    fmt.Println(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

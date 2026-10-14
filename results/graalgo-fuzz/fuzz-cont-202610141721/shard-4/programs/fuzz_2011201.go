package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0() uint64 {
    v7 := []int{947, 32768, 127}
    var v8 int = 4294967295
    var v9 string = ("\n\t" + "a")
    fmt.Println(v9, v8)
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    v3 := []string{"\n\t"}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    v6 := make(map[string]int)
    v10 := v2[""]
    var v11 uint16 = 255
    v11 += (v11 | 0)
    var v12 int64 = -371
    fmt.Println(v11)
    if (42 < 42) {
        for fi13 := 0; fi13 < 3; fi13++ {
            if (!(2 != 381)) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 bool = (0 < fi14)
                    if len(v3) > 0 {
                        v3[0] = (v10 + v10)
                    }
                    if (false || v15) {
                        var v16 int8 = 100
                        v17 := GenS1{F0: v12}
                        v15 = (!(!false))
                        v12 = 557
                        if (false && v15) {
                            var v18 bool = (v15 && v15)
                            _ = len(v5)
                            var v19 int64 = (v12 << ((51) & 63))
                            v10 = (("foo bar" + ("hello" + ((((v10 + "hello") + v10) + (((v10 + (("foo bar" + v10) + "")) + "foo bar") + v10)) + (("" + (v10 + v10)) + v10)))) + "hello")
                            _ = v18
                            _ = v19
                        }
                        if (!(!v15)) {
                            var v20 float64 = float64(0.1)
                            var v21 uint8 = 128
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v20
                            _ = v21
                        }
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                if true {
                    var v22 bool = (!true)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v24 := GenS2{F0: 814}
                        var v25 uint32 = 921
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                } else {
                    v26 := GenS1{F0: (v12 & (v12 ^ (v12 % (((^(v12))) | 1))))}
                    fmt.Println(v10)
                    v27 := GenS1{F0: 9223372036854775807}
                    _ = v26
                    _ = v27
                }
                var v28 int64 = v12
                _ = v28
            } else {
                switch fi13 {
                case -739:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32768:
                    v1 -= v1
                default:
                    v29 := GenS3{F0: 65535, F1: 1}
                    _ = v29
                }
            }
            _ = fi13
        }
    } else {
        fmt.Println(v12, v11)
        for fi30 := 0; fi30 < 2; fi30++ {
            if ((float64(-2.25) / 1.0) == float64((0.0 * 1.5))) {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v11 += (v11 & 255)
                    var v32 int64 = 9223372036854775807
                    v3 = append(v3, v10)
                    _ = fi31
                    _ = v32
                }
                fmt.Println(v10)
            }
            _ = fi30
        }
    }
    v1 += v1
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

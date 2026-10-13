package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64) int {
    v7 := []int{479}
    var v8 uint16 = 0
    v9 := GenS1{F0: 1000, F1: 127}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []uint32{0, 1, 2147483648, 510}
    v2 := make(map[int]bool)
    var v3 int64 = -372
    var v4 float32 = 1.5
    var v5 float64 = 0.1
    var v6 int8 = 100
    v4 -= v4
    v10 := GenS1{F0: 1000, F1: 128}
    for fi11 := 0; fi11 < 16; fi11++ {
        if false {
            v2[0] = (!false)
            var v12 int64 = 490
            _ = v12
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                if (v5 <= float64(v4)) {
                    if len(v1) > 0 {
                        v1[0] = 8
                    }
                    var v14 uint16 = 455
                    var v15 int = genFn0(18446744073709551615)
                    switch 127 {
                    case 128:
                        v15 = fi13
                    }
                    _ = v14
                    _ = v15
                } else {
                    var v16 int = genFn0(0)
                    v17 := v2[(v16 % (((v16 + v16)) | 1))]
                    _ = v16
                    _ = v17
                }
                var v18 uint64 = 745
                _ = fi13
                _ = v18
            }
            if (127 >= 1) {
                v4 = v4
                _ = len(v2)
            }
            v4 = 0.0
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            v3 = (v3 - (v3 & (v3 | (v3 - (v3 >> ((0) & 63))))))
            var v20 float64 = v5
            var v21 uint64 = 18446744073709551615
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi11
    }
    var v22 uint8 = 155
    v23 := GenS2{F0: false}
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v24 int8 = v6
    var v25 float32 = -2.25
    v26 := GenS1{F0: 172, F1: genFn0(2)}
    var v27 uint32 = 65535
    if true {
        delete(v2, 0)
        var v28 bool = (42 > -32768)
        for fi29 := 0; fi29 < 2; fi29++ {
            fmt.Println(v22)
            v22 = 67
            if v28 {
                switch (("a" + "\n\t") + "a") {
                case "\n\t":
                    fmt.Println(v6, fi29, v22)
                case "a":
                    var v30 int32 = 1
                    _ = v30
                case "hello":
                    if (true && ((v4 + 1.5) < -2.25)) {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            if len(v1) > 0 {
                                v1[0] = (v27 & (v27 % ((v27) | 1)))
                            }
                            v22 = v22
                            v22 = 128
                            _ = fi31
                        }
                        var v32 string = "foo bar"
                        v33 := GenS1{F0: -1, F1: (fi29 ^ 32767)}
                        _ = v32
                        _ = v33
                    }
                default:
                    var v34 int = (fi29 / ((genFn0(42)) | 1))
                    _ = v34
                }
                v6 = -1
                var v35 uint64 = 4294967295
                _ = v35
            } else {
                v36 := GenS3{F0: 4294967295}
                _ = v36
            }
            _ = fi29
        }
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

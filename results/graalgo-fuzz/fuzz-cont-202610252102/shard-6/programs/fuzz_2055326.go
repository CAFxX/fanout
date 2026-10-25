package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    if (!false) {
        for fi4 := 0; fi4 < 16; fi4++ {
            v5 := GenS1{F0: 42}
            if true {
                var v6 uint8 = 14
                var v7 int = fi4
                _ = v6
                _ = v7
            } else {
                v8 := []int{157, (^(fi4)), (fi4 * fi4), 628}
                _ = len(v8)
                _ = v8
            }
            _ = fi4
            _ = v5
        }
    }
    p0 += (p0 | p0)
    _ = p0
    return 663
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v9 bool = true
    v1 = (v1 + (v1 % (((v1 ^ -1)) | 1)))
    v1 += (-(v1))
    var v10 int32 = (v1 - (v1 - (v1 / ((2147483647) | 1))))
    for fi11 := 0; fi11 < 14; fi11++ {
        if (true && (128 != 127)) {
            var v12 int8 = int8(fi11)
            fmt.Println(v9, v12, fi11)
            switch (fi11 * (-(fi11))) {
            case 4294967295:
                if (!(0.0 != 3.14159)) {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v14 := GenS1{F0: v12}
                        switch ("\n\t" + (("\n\t" + "a") + "\n\t")) {
                        case "hello":
                            var v15 int16 = 280
                            _ = v15
                        default:
                            if (-32768 < 1) {
                                if v9 {
                                    var v16 int = (fi13 << ((42) & 63))
                                    v2[v16] = ("foo bar" + "\n\t")
                                    switch "\n\t" {
                                    case "a":
                                        v1 -= v1
                                    }
                                    _ = v16
                                } else {
                                    fmt.Println(v1)
                                    if v9 {
                                        v17 := []int{fi13, (fi13 - (-(fi13))), 1}
                                        _ = v17
                                    } else {
                                        v9 = (3.14159 == (float64(((3.14159 - (float32(-2.25) - -2.25)) - 1e10)) / -2.25))
                                        v10 -= v10
                                        v1 = int32(fi13)
                                    }
                                }
                                v3[909] = 4294967295
                                var v18 int32 = (v10 << ((0) & 63))
                                _ = v18
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                    var v19 uint64 = 1
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if v9 {
                            var v21 int = fi11
                            v10 = v1
                            v1 = -1
                            _ = v21
                        }
                        _ = fi20
                    }
                }
            case -244:
                for fi22 := 0; fi22 < 4; fi22++ {
                    fmt.Println(v12, v10, fi22)
                    v1 -= 42
                    if v9 {
                        v12 -= v12
                    } else {
                        v9 = v9
                        v1 = (v1 + 2147483647)
                    }
                    _ = fi22
                }
            default:
                var v23 uint8 = 1
                _ = v23
            }
            _ = v12
        } else {
            fmt.Println(v10, v10)
        }
        fmt.Println(v9, v1)
        for fi24 := 0; fi24 < 13; fi24++ {
            if (276 <= 18446744073709551615) {
                var v25 int64 = 632
                fmt.Println(v10, v10, v25)
                _ = v25
            }
            if false {
                var v26 uint = 18446744073709551615
                v1 = (v10 / (((v1 - (v1 * (v1 >> (((v26 >> (((v26 << ((v26) & 63))) & 63))) & 63))))) | 1))
                _ = v26
            }
            v9 = (true && (326 != 65535))
            var v27 bool = true
            _ = fi24
            _ = v27
        }
        _ = fi11
    }
    var v28 uint8 = 128
    v28 = v28
    var v29 uint16 = 204
    if (100 != 100) {
        v9 = ((v29 ^ 262) >= 32767)
        switch 4294967295 {
        case 0:
            v10 -= v1
        case 2147483647:
            fmt.Println(v29, v29, v28)
        case -2147483648:
            for fi30 := 0; fi30 < 19; fi30++ {
                v9 = v9
                v29 -= 51
                var v31 float32 = -2.25
                _ = fi30
                _ = v31
            }
        }
    } else {
        v32 := v3[4294967295]
        _ = v32
    }
    fmt.Println(v28, v1)
    v2[1] = ""
    v33 := GenS1{F0: 127}
    for fi34 := 0; fi34 < 19; fi34++ {
        v9 = v9
        _ = fi34
    }
    v35 := []int{2147483647, 814, -686, 691}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v35
    _ = v9
}

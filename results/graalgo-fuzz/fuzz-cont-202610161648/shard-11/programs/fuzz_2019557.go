package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    v1 = (v1 & 404)
    v4 := []int{128, 623, 127}
    var v5 string = ("\n\t" + "\n\t")
    if true {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 1}
            var v8 uint = 910
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            if true {
                var v10 uint = 4294967295
                v5 = v5
                _ = v10
            } else {
                switch (fi9 >> ((2) & 63)) {
                case 913:
                    v1 -= -32768
                default:
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 uint64 = 4294967295
                        v1 = (v1 - v1)
                        _ = fi11
                        _ = v12
                    }
                }
                v13 := GenS2{F0: (v1 % ((v1) | 1)), F1: false}
                switch ((v5 + (v5 + (v5 + ""))) + "hello") {
                case "a":
                    var v14 uint64 = 42
                    _ = v14
                default:
                    v15 := GenS1{F0: 9223372036854775807}
                    _ = v15
                }
                _ = v13
            }
            _ = fi9
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            v1 = v1
            _ = len(v2)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
        }
    }
    var v17 int64 = -1
    fmt.Println(v1)
    v18 := GenS1{F0: (v17 % (((^(v17))) | 1))}
    v19 := GenS2{F0: -32768, F1: (((v5 + "\n\t") + v5) <= ((v5 + "\n\t") + (v5 + (("foo bar" + (v5 + "foo bar")) + "hello"))))}
    v1 -= (v1 << ((131) & 63))
    v20 := GenS2{F0: (v1 * (v1 | v1)), F1: true}
    var v21 float32 = 1e10
    fmt.Println(v3)
    v5 = v5
    v22 := GenS2{F0: -1, F1: (float32((0.1 / 1.0)) == v3)}
    for fi23 := 0; fi23 < 17; fi23++ {
        v3 = ((0.1 * 0.1) / float32((((3.14159 - (3.14159 - (0.1 * -2.25))) / 1.0) + 1e10)))
        v5 = v5
        _ = fi23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v2[fi24] = fi24
        var v25 string = (v5 + v5)
        switch fi24 {
        case -484:
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 uint8 = 186
                v4 = append(v4, 32768)
                if (255 == 786) {
                    delete(v2, 19)
                    if ((-(fi24)) >= fi26) {
                        v17 += 0
                        var v28 uint8 = (^(v27))
                        _ = v28
                    }
                    v29 := GenS1{F0: v17}
                    if false {
                        var v30 int = fi24
                        _ = v30
                    } else {
                        v31 := GenS1{F0: 42}
                        v2[-83] = 1
                        _ = v31
                    }
                    _ = v29
                } else {
                    if true {
                        var v32 int16 = -207
                        if false {
                            if (255 == 1) {
                                switch v25 {
                                case "\n\t":
                                    var v33 bool = (829 <= 2)
                                    _ = v33
                                case "foo bar":
                                    var v34 uint = 1
                                    _ = v34
                                }
                            } else {
                                var v35 int32 = 42
                                _ = v35
                            }
                        } else {
                            v36 := GenS1{F0: v17}
                            if (42 < 568) {
                                switch fi26 {
                                case 4294967295:
                                    fmt.Println(fi26, v25)
                                    var v37 int16 = 1
                                    _ = v37
                                case 0:
                                    fmt.Println(v21)
                                case 9223372036854775807:
                                    v27 = 198
                                }
                                if len(v4) > 0 {
                                    v4[0] = (fi24 % ((fi26) | 1))
                                }
                                v5 = v5
                            }
                            delete(v2, (-(fi26)))
                            _ = v36
                        }
                        var v38 bool = false
                        if ((^(v27)) < (v27 & (v27 / (((v27 & v27)) | 1)))) {
                            if true {
                                if false {
                                    v3 = v21
                                    if false {
                                        v1 += -32768
                                        if v38 {
                                            var v39 float32 = ((v21 / 0.1) + float32(0.0))
                                            v40 := GenS3{F0: 42, F1: -1}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v32 = 846
                                        }
                                        v17 += 9223372036854775807
                                    }
                                } else {
                                    switch (v5 + v25) {
                                    case "\n\t":
                                        v21 = 3.14159
                                    }
                                }
                                if v38 {
                                    v41 := GenS2{F0: (v32 % ((v1) | 1)), F1: v38}
                                    var v42 bool = (100 >= -1)
                                    _ = v41
                                    _ = v42
                                } else {
                                    _ = len(v4)
                                    v17 = v17
                                }
                                var v43 int8 = 127
                                v43 = (v43 + 127)
                                delete(v2, (fi26 * fi24))
                                _ = v43
                            }
                        }
                        _ = v32
                        _ = v38
                    } else {
                        var v44 int8 = 100
                        var v45 float64 = -2.25
                        var v46 uint8 = (^(v27))
                        v47 := v2[9223372036854775807]
                        _ = v44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    _ = len(v4)
                }
                _ = fi26
                _ = v27
            }
        default:
            var v48 uint8 = 55
            _ = v48
        }
        _ = fi24
        _ = v25
    }
    if true {
        var v49 string = v5
        v21 = (float32(((-2.25 * -2.25) * float64(v21))) / v3)
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}

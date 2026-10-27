package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 int8; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
func genFn0() uint64 {
    v6 := GenS1{F0: 18446744073709551615, F1: -566, F2: ("foo bar" + ("hello" + ""))}
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            v8 := []int{(fi7 / (((fi7 | fi7)) | 1)), fi7, fi7, (fi7 << ((255) & 63))}
            _ = v8
        } else {
            v9 := []int{}
            _ = v9
        }
        var v10 int16 = int16(fi7)
        if false {
            var v11 uint64 = 18446744073709551615
            _ = v11
        }
        _ = fi7
        _ = v10
    }
    var v12 bool = false
    var v13 uint = 487
    var v14 uint16 = 2
    _ = v12
    _ = v13
    _ = v14
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 uint16 = 255
    v4 := make(map[int]uint32)
    var v5 int = 618
    v15 := GenS2{F0: 255, F1: -1}
    var v16 float64 = float64(1e10)
    v4[v5] = 2
    v17 := GenS3{F0: 42, F1: 1, F2: 1}
    if (127 < 0) {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 int64 = -986
            _ = fi18
            _ = v19
        }
    }
    var v20 bool = ((!true) || (730 >= 38))
    switch (v5 << ((371) & 63)) {
    case -910:
        var v21 uint32 = 84
        _ = v21
    default:
        var v22 float32 = (0.0 - 1e10)
        _ = v22
    }
    if ((1 >= 4294967295) && (-128 == -9)) {
        v23 := GenS4{F0: "hello", F1: -1, F2: -9223372036854775808}
        if v20 {
            var v24 int64 = 9223372036854775807
            var v25 int = (v5 | v5)
            if (v20 || true) {
                if v20 {
                    if false {
                        var v26 uint8 = 255
                        if ((v20 || (-1 == -1)) && (1 >= 42)) {
                            if false {
                                v5 += (v5 + v25)
                            }
                            v5 += v25
                        } else {
                            if (!(v20 || v20)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    fmt.Println(v24)
                                    var v27 bool = v20
                                    var v28 float64 = (v16 * ((0.1 - float64(3.14159)) / (0.1 / 1.5)))
                                    _ = v27
                                    _ = v28
                                } else {
                                    for fi29 := 0; fi29 < 11; fi29++ {
                                        fmt.Println(v5, v24)
                                        var v30 int8 = int8(v24)
                                        v31 := GenS1{F0: 2, F1: 0, F2: "foo bar"}
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                    v20 = v20
                                    var v32 uint32 = 42
                                    _ = v32
                                }
                                switch (^(v25)) {
                                case 883:
                                    _ = len(v2)
                                }
                            }
                            v4[(v5 - (v5 - 4294967295))] = 0
                        }
                        _ = v26
                    } else {
                        v33 := GenS5{F0: (v3 << ((538) & 63)), F1: 42}
                        var v34 int32 = 599
                        _ = v33
                        _ = v34
                    }
                    v16 = 0.0
                    fmt.Println(v5)
                } else {
                    v35 := GenS6{F0: 125}
                    v3 -= v3
                    for fi36 := 0; fi36 < 4; fi36++ {
                        v37 := GenS7{F0: v5}
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
                for fi38 := 0; fi38 < 5; fi38++ {
                    fmt.Println(v20, v16)
                    _ = fi38
                }
                var v39 uint64 = 0
                _ = v39
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v24
            _ = v25
        } else {
            var v40 float32 = 1.5
            _ = v40
        }
        _ = v23
    } else {
        delete(v1, "")
        v16 = 1.5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

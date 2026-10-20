package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 int32 }
func genFn0() int {
    if true {
        var v6 bool = true
        _ = v6
    }
    var v7 string = ""
    v8 := []int{-274, 223, 9223372036854775807, -225}
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]uint32)
    var v3 uint32 = 4294967295
    var v4 int32 = 42
    var v5 uint = 807
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: (true || true), F1: -320}
            if true {
                var v12 uint64 = 42
                if true {
                    v13 := GenS1{F0: false, F1: -9223372036854775808}
                    if ((fi9 % ((genFn0()) | 1)) >= 4294967295) {
                        if true {
                            v12 += 18446744073709551615
                            if false {
                                var v14 int8 = -125
                                v1 = v1
                                _ = v14
                            } else {
                                v15 := v2[(fi9 + -321)]
                                _ = v15
                            }
                            if (false || true) {
                                v16 := GenS1{F0: true, F1: 1}
                                _ = v16
                            }
                        } else {
                            v17 := []int{genFn0(), 2147483647, (fi10 - (fi10 - -371)), fi9}
                            v18 := GenS1{F0: true, F1: 42}
                            fmt.Println(v1, fi9)
                            _ = v17
                            _ = v18
                        }
                        if (!true) {
                            v2[fi10] = 572
                            fmt.Println(fi10, v5)
                        } else {
                            var v19 float64 = 0.0
                            v12 += 273
                            _ = v19
                        }
                        if false {
                            v4 = (-(v4))
                            switch (fi10 / ((-986) | 1)) {
                            case -2147483648:
                                var v20 uint = v5
                                _ = v20
                            case 0:
                                if ((!((v3 < v3) || false)) && true) {
                                    delete(v2, 51)
                                    var v21 uint8 = 50
                                    v22 := GenS2{F0: (v4 + v4), F1: true, F2: 255}
                                    _ = v21
                                    _ = v22
                                }
                            }
                            var v23 int8 = 42
                            _ = v23
                        }
                    } else {
                        v3 = 1
                        v4 = (^(v4))
                        var v24 uint = (v5 * 42)
                        _ = v24
                    }
                    v25 := GenS1{F0: (406 < 255), F1: -1}
                    v26 := GenS1{F0: false, F1: -818}
                    var v27 uint16 = 672
                    _ = v13
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    v28 := []int{0, fi9, genFn0(), genFn0()}
                    var v29 uint32 = (v3 ^ v3)
                    v4 = -1
                    _ = v28
                    _ = v29
                }
                var v30 int32 = -1
                var v31 float64 = (0.1 + 1.5)
                _ = v12
                _ = v30
                _ = v31
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS1{F0: true, F1: 42}
        _ = fi32
        _ = v33
    }
    v34 := v2[32767]
    v35 := []int{42}
    v36 := GenS2{F0: (v4 / ((2147483647) | 1)), F1: (!true), F2: v5}
    var v37 float32 = (-2.25 - -2.25)
    for fi38 := 0; fi38 < 19; fi38++ {
        v4 += v4
        _ = fi38
    }
    var v39 int16 = 42
    switch genFn0() {
    case 32767:
        var v40 string = "foo bar"
        _ = v40
    case 845:
        var v41 int16 = (v39 % ((v39) | 1))
        _ = v41
    }
    v2[-483] = v3
    var v42 string = "hello"
    if len(v35) > 0 {
        v35[0] = 641
    }
    v43 := GenS3{F0: -1, F1: (v4 % ((321) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}

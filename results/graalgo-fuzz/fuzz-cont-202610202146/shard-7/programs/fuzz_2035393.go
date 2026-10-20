package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int32; F1 string }
func genFn0() int {
    var v7 uint64 = 18446744073709551615
    if (true || true) {
        fmt.Println(v7)
        v7 -= 929
    } else {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int32 = -1
            _ = fi8
            _ = v9
        }
    }
    v10 := []int{127, -1, 42}
    _ = v10
    _ = v7
    return 824
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    v4 := []string{"", "a"}
    var v5 int8 = 0
    var v6 int64 = int64(v5)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch (fi12 >> ((42) & 63)) {
            case 968:
                var v13 int64 = v6
                _ = v13
            case 9223372036854775807:
                v1 -= (v1 * 4294967295)
            case 435:
                v6 = v6
            }
            if true {
                v14 := GenS1{F0: v1}
                v1 -= (^(v1))
                if false {
                    if true {
                        var v15 uint8 = 128
                        v6 += v6
                        var v16 float32 = 1e10
                        _ = v15
                        _ = v16
                    }
                    var v17 int32 = -244
                    v18 := GenS1{F0: (v1 * (v1 ^ v1))}
                    _ = v17
                    _ = v18
                } else {
                    switch "" {
                    case "":
                        if true {
                            delete(v3, "foo bar")
                        } else {
                            switch "" {
                            case "\n\t":
                                var v19 int64 = (v6 - (v6 + 1))
                                _ = v19
                            case "hello":
                                v20 := GenS1{F0: v1}
                                _ = v20
                            case "":
                                var v21 uint8 = 0
                                _ = v21
                            default:
                                v22 := GenS1{F0: 243}
                                _ = v22
                            }
                            v5 -= 42
                        }
                    }
                    v5 += v5
                }
                _ = v14
            } else {
                if true {
                    v23 := GenS1{F0: v1}
                    delete(v3, "foo bar")
                    fmt.Println(v5, v5, fi12)
                    fmt.Println(fi12, fi12, fi12)
                    _ = v23
                } else {
                    fmt.Println(v5, v5)
                }
                var v24 float32 = 0.0
                v5 = -30
                _ = v24
            }
            var v25 float32 = (float32(1.5) - ((3.14159 + 0.1) * 0.1))
            if ((v5 - 42) == (v5 & (v5 / (((v5 >> ((255) & 63))) | 1)))) {
                v26 := GenS2{F0: 1, F1: 0, F2: -2147483648}
                var v27 int8 = (^(v5))
                if false {
                    v6 += (^(v6))
                    var v28 int16 = int16(v27)
                    if (((0 == 1) || (!(false || false))) && true) {
                        v1 -= v1
                    }
                    _ = v28
                } else {
                    var v29 int64 = -9223372036854775808
                    _ = v29
                }
                _ = v26
                _ = v27
            }
            _ = fi12
            _ = v25
        }
        _ = fi11
    }
    switch (("foo bar" + "\n\t") + "a") {
    case "":
        fmt.Println(v1, v6)
    case "foo bar":
        if (v5 <= v5) {
            delete(v2, 2147483648)
            v30 := GenS2{F0: 1, F1: 32767, F2: 0}
            v6 = v6
            _ = v30
        }
    }
    v31 := v3["\n\t"]
    fmt.Println(v5)
    _ = len(v4)
    var v32 int16 = 42
    var v33 uint = 2
    for fi34 := 0; fi34 < 2; fi34++ {
        if (v31 == v31) {
            v5 -= v5
        } else {
            for fi35 := 0; fi35 < 10; fi35++ {
                if false {
                    var v36 uint32 = v1
                    var v37 float32 = 3.14159
                    _ = v36
                    _ = v37
                }
                v5 += -1
                v6 = int64(v33)
                _ = fi35
            }
            for fi38 := 0; fi38 < 19; fi38++ {
                v39 := GenS3{F0: false, F1: (v6 >> (((v33 ^ (v33 >> ((v33) & 63)))) & 63))}
                var v40 uint = (v33 & v33)
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        _ = fi34
    }
    var v41 int32 = 42
    var v42 uint32 = (v1 ^ (^(v1)))
    if (1 != (v32 ^ v32)) {
        for fi43 := 0; fi43 < 14; fi43++ {
            v44 := GenS1{F0: 42}
            if len(v4) > 0 {
                v4[0] = v31
            }
            _ = fi43
            _ = v44
        }
    } else {
        for fi45 := 0; fi45 < 19; fi45++ {
            switch (("\n\t" + (("\n\t" + ((v31 + v31) + ((((v31 + v31) + "foo bar") + v31) + v31))) + v31)) + v31) {
            case "hello":
                v6 += int64(v41)
            }
            v31 = v31
            _ = fi45
        }
    }
    switch 1 {
    case -1:
        var v46 int64 = (v6 / ((v6) | 1))
        _ = v46
    case 4294967295:
        v1 = (v42 & v42)
    }
    var v47 int16 = (v32 + (v32 & (v32 | int16(v5))))
    v48 := GenS4{F0: v41, F1: "hello"}
    var v49 int16 = v32
    var v50 int16 = v47
    v33 += (v33 + v33)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}

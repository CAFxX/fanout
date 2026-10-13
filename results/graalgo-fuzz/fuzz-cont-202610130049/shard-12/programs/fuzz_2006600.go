package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int8; F1 int64; F2 int32 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int16) string {
    v7 := GenS1{F0: 2147483647, F1: 127}
    v8 := []int{-1, 0}
    _ = p0
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint = 829
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    var v6 string = "hello"
    var v9 uint64 = 980
    v10 := GenS1{F0: 28, F1: 1}
    v11 := GenS1{F0: 979, F1: 2}
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int8 = 127
        v14 := GenS1{F0: 405, F1: 6}
        v15 := []int{fi12, (fi12 + fi12), (fi12 + fi12), (fi12 ^ (fi12 << ((v1) & 63)))}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: -61, F1: 100, F2: -102}
    if false {
        if ((true || (("" + v6) > (("" + "") + v6))) || true) {
            v17 := GenS2{F0: 0, F1: 100, F2: 32768}
            _ = v17
        }
        v18 := GenS3{F0: false, F1: 1, F2: 42}
        _ = v18
    }
    if false {
        var v19 float64 = 0.0
        var v20 float32 = (3.14159 + 1.5)
        var v21 int16 = -32768
        _ = v19
        _ = v20
        _ = v21
    } else {
        if true {
            var v22 int16 = 587
            _ = v22
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                if (1 > 2147483647) {
                    var v24 float64 = (0.1 * float64((3.14159 + 0.0)))
                    _ = v24
                } else {
                    var v25 int32 = int32(fi23)
                    if (((3.14159 / (float32(-2.25) + 3.14159)) < (0.0 + 0.0)) || (v6 >= "a")) {
                        v26 := GenS4{F0: 0, F1: 42, F2: (^(v25))}
                        _ = v26
                    }
                    _ = v25
                }
                _ = fi23
            }
            v27 := GenS4{F0: 0, F1: -9223372036854775808, F2: 716}
            _ = v27
        }
        switch (v6 + "hello") {
        case "\n\t":
            if false {
                v28 := v3[158]
                v29 := []int{(v28 * v28), (v28 & (^(v28))), (v28 + -773)}
                if ((0 == 276) && false) {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v28 += (^(v28))
                        switch (v6 + (("" + (v6 + "\n\t")) + ((("\n\t" + "\n\t") + v6) + v6))) {
                        case "hello":
                            var v31 int8 = 93
                            _ = v31
                        case "foo bar":
                            switch ((((v6 + v6) + "") + "foo bar") + v6) {
                            case "foo bar":
                                v1 = v1
                            }
                        case "\n\t":
                            v28 += (fi30 * (v28 << ((v1) & 63)))
                        default:
                            v32 := GenS2{F0: 127, F1: 1, F2: fi30}
                            _ = v32
                        }
                        v33 := GenS5{F0: 240}
                        _ = fi30
                        _ = v33
                    }
                    fmt.Println(v28)
                } else {
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                _ = v28
                _ = v29
            }
        default:
            v9 = (v9 >> (((v1 % ((31) | 1))) & 63))
        }
    }
    fmt.Println(v6)
    v6 = ""
    v6 = v6
    v1 -= (^(v1))
    switch 128 {
    case 32767:
        v34 := GenS5{F0: 37}
        _ = v34
    case 4294967295:
        for fi35 := 0; fi35 < 11; fi35++ {
            v9 = (v9 * 185)
            if true {
                var v36 float32 = 1e10
                if ("foo bar" > (v6 + v6)) {
                    if true {
                        var v37 int64 = 759
                        var v38 int64 = (v37 & 9223372036854775807)
                        _ = v37
                        _ = v38
                    } else {
                        var v39 string = v6
                        switch fi35 {
                        case 526:
                            v39 = v39
                        default:
                            var v40 int = fi35
                            _ = v40
                        }
                        switch v39 {
                        case "foo bar":
                            v41 := GenS5{F0: 6}
                            _ = v41
                        case "hello":
                            v42 := GenS5{F0: 104}
                            _ = v42
                        }
                        _ = v39
                    }
                    v9 = (v9 / ((v9) | 1))
                }
                _ = v36
            } else {
                var v43 int = fi35
                var v44 int64 = 42
                _ = v43
                _ = v44
            }
            v3[1] = 127
            _ = fi35
        }
    case 706:
        var v45 uint64 = 18446744073709551615
        _ = v45
    default:
        var v46 string = (v6 + v6)
        _ = v46
    }
    switch 32768 {
    case 128:
        fmt.Println(v6)
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        fmt.Println(v6, fi47)
        fmt.Println(v9)
        _ = fi47
    }
    if (v9 > v9) {
        var v48 string = "\n\t"
        var v49 float64 = 1.5
        v1 = 255
        _ = v48
        _ = v49
    } else {
        for fi50 := 0; fi50 < 1; fi50++ {
            v6 = ((((("a" + (v6 + v6)) + "") + ((((v6 + (v6 + ((v6 + v6) + v6))) + "foo bar") + "foo bar") + (v6 + "a"))) + ("\n\t" + v6)) + (v6 + v6))
            _ = fi50
        }
    }
    var v51 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v9
}

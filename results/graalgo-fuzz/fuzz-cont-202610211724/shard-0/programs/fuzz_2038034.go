package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 bool = (761 >= 255)
    v3 := []int{127}
    var v4 float32 = ((float32(1e10) - 3.14159) * 3.14159)
    if v2 {
        v5 := v1["a"]
        v6 := GenS1{F0: 0, F1: 1}
        _ = v5
        _ = v6
    }
    v7 := v1[("foo bar" + (("" + "a") + "a"))]
    if false {
        v8 := v1["foo bar"]
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int16 = -510
            for fi11 := 0; fi11 < 16; fi11++ {
                if (!v2) {
                    v4 += 0.1
                    switch ("a" + "") {
                    case "":
                        v12 := v1[(("foo bar" + ("\n\t" + "hello")) + "foo bar")]
                        _ = v12
                    case "foo bar":
                        v4 = 3.14159
                    }
                    var v13 int32 = 589
                    _ = v13
                }
                switch fi11 {
                case -1:
                    if (v2 || false) {
                        switch v7 {
                        case 106:
                            v2 = (!true)
                        case 2147483647:
                            var v14 int16 = v10
                            var v15 bool = (1 == 4294967295)
                            v10 = 1000
                            _ = v14
                            _ = v15
                        default:
                            switch (-(fi11)) {
                            case -1:
                                v2 = v2
                            default:
                                v16 := GenS1{F0: 1, F1: 1}
                                _ = v16
                            }
                            v17 := GenS2{F0: v2}
                            v10 = v10
                            v2 = true
                            _ = v17
                        }
                        if v2 {
                            fmt.Println(v8)
                            v2 = (73 > -79)
                            fmt.Println(v2, fi9)
                            switch fi9 {
                            case -410:
                                _ = len(v1)
                            default:
                                v18 := GenS3{F0: 65535, F1: fi9, F2: "hello"}
                                v19 := v1["foo bar"]
                                _ = v18
                                _ = v19
                            }
                            if (!true) {
                                v20 := v1[""]
                                _ = v20
                            } else {
                                v21 := v1[""]
                                if (!v2) {
                                    v1["\n\t"] = -2147483648
                                    _ = len(v3)
                                } else {
                                    v10 += (v10 ^ v10)
                                    var v22 float64 = 0.1
                                    _ = v22
                                }
                                var v23 int8 = 77
                                _ = v21
                                _ = v23
                            }
                            var v24 int32 = 0
                            _ = v24
                        } else {
                            v25 := GenS1{F0: 1, F1: 2}
                            fmt.Println(v8)
                            if ((v7 - 42) == v7) {
                                var v26 int16 = v10
                                v10 = -353
                                var v27 bool = (!v2)
                                _ = v26
                                _ = v27
                            } else {
                                if true {
                                    v7 = v7
                                    v2 = v2
                                }
                                v28 := GenS2{F0: false}
                                var v29 float32 = ((v4 / ((float32(3.14159) / 3.14159) / (float32(float64((((v4 - 3.14159) - float32(0.1)) * (v4 * v4)))) - v4))) * -2.25)
                                _ = v28
                                _ = v29
                            }
                            _ = v25
                        }
                    } else {
                        if true {
                            fmt.Println(fi11)
                            var v30 int64 = 927
                            switch (fi9 ^ (v8 >> ((677) & 63))) {
                            case -414:
                                var v31 string = "a"
                                _ = v31
                            default:
                                fmt.Println(v4, v30, fi9)
                            }
                            _ = v30
                        }
                        var v32 float64 = (1e10 - 1e10)
                        v4 += (1.5 * (v4 / (((v4 + v4) - v4) - v4)))
                        v8 = (^(v7))
                        v3 = append(v3, (fi11 & fi9))
                        _ = v32
                    }
                case 127:
                    v33 := GenS1{F0: 42, F1: 1}
                    _ = v33
                case 990:
                    v2 = true
                default:
                    v4 = float32((float64(-2.25) * (1e10 + 3.14159)))
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v2 = (731 > 302)
    var v34 int = 2147483647
    _ = len(v1)
    v35 := GenS2{F0: v2}
    v7 = -2147483648
    var v36 uint64 = 2
    var v37 int32 = 1
    var v38 uint16 = 42
    v3 = append(v3, v7)
    v2 = false
    v39 := GenS2{F0: (!v2)}
    var v40 int32 = v37
    v38 = v38
    v41 := GenS2{F0: ((!v2) && true)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v7
}

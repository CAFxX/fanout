package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 string; F2 int32 }
func genFn0(p0 int, p1 uint32) uint {
    var v5 bool = (42 < 0)
    var v6 int16 = -842
    v7 := GenS1{F0: "", F1: 4294967295, F2: "a"}
    if false {
        v8 := []int{}
        v5 = ((-(v6)) == (v6 - -32768))
        if v5 {
            if false {
                if v5 {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v8 = append(v8, (-(fi9)))
                        for fi10 := 0; fi10 < 1; fi10++ {
                            var v11 uint32 = p1
                            var v12 uint32 = (p1 * 0)
                            switch p0 {
                            case 170:
                                v13 := GenS2{F0: -160, F1: -422}
                                _ = v13
                            case 303:
                                var v14 float64 = 0.0
                                _ = v14
                            case 9223372036854775807:
                                var v15 uint = 65535
                                _ = v15
                            default:
                                var v16 uint64 = 0
                                _ = v16
                            }
                            var v17 int = -9223372036854775808
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v17
                        }
                        if len(v8) > 0 {
                            v8[0] = p0
                        }
                        _ = fi9
                    }
                    for fi18 := 0; fi18 < 20; fi18++ {
                        switch (p0 >> ((408) & 63)) {
                        case 0:
                            v19 := GenS2{F0: int64(v6), F1: -9223372036854775808}
                            _ = v19
                        case 9223372036854775807:
                            v20 := GenS3{F0: 98}
                            _ = v20
                        }
                        var v21 int = (-(fi18))
                        _ = fi18
                        _ = v21
                    }
                    var v22 int32 = 2147483647
                    _ = v22
                }
                if v5 {
                    var v23 bool = (!((1.5 / float32(float64(((1e10 - ((1e10 + 3.14159) - (1e10 / ((float32(3.14159) * -2.25) / ((((1e10 / -2.25) * (0.1 + 1.5)) / 3.14159) / (3.14159 * 0.1)))))) - (0.0 * ((1e10 * -2.25) * (0.1 / ((-2.25 + 1e10) / (float32(3.14159) * 1.5))))))))) > 1e10))
                    _ = v23
                } else {
                    var v24 uint64 = 0
                    _ = v24
                }
            } else {
                p0 -= -766
                for fi25 := 0; fi25 < 20; fi25++ {
                    p1 = p1
                    _ = fi25
                }
            }
        } else {
            v26 := GenS1{F0: "a", F1: 1, F2: "foo bar"}
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 uint64 = 18446744073709551615
                p1 += 4294967295
                v28 -= v28
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -1
    var v3 float32 = (0.1 / float32(3.14159))
    var v4 int16 = v2
    var v29 int16 = v4
    v3 += (v3 + (v3 * ((1e10 / 1.0) * 1.5)))
    v30 := []int{-2147483648, -1, -906}
    v29 += -1
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 uint64 = 18446744073709551615
        v32 = v32
        switch "" {
        case "foo bar":
            v4 = v29
        }
        _ = fi31
        _ = v32
    }
    _ = len(v30)
    var v33 int32 = 847
    v34 := GenS2{F0: -1, F1: 0}
    for fi35 := 0; fi35 < 7; fi35++ {
        switch (^(fi35)) {
        case 127:
            for fi36 := 0; fi36 < 20; fi36++ {
                fmt.Println(v2, v2, fi36)
                v37 := GenS1{F0: "\n\t", F1: 988, F2: ""}
                _ = fi36
                _ = v37
            }
        case -2147483648:
            if (v33 > 42) {
                if true {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        var v39 uint64 = 945
                        v40 := GenS1{F0: "", F1: 65535, F2: (("foo bar" + "a") + "")}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                } else {
                    switch "a" {
                    case "a":
                        v41 := GenS1{F0: ("" + "\n\t"), F1: 255, F2: "\n\t"}
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 12; fi42++ {
                        if true {
                            v2 = -1
                            v2 = v2
                            var v43 string = "hello"
                            _ = v43
                        }
                        var v44 int8 = -1
                        _ = fi42
                        _ = v44
                    }
                }
            }
        }
        v45 := GenS4{F0: (-1 != 1), F1: ("a" + "foo bar"), F2: (^(v33))}
        _ = fi35
        _ = v45
    }
    v29 -= 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
}

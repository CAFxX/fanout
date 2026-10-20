package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 int8 }
type GenS4 struct { F0 string; F1 string; F2 uint }
func genFn0(p0 int16) uint32 {
    p0 = (p0 >> ((0) & 63))
    v7 := GenS1{F0: 18446744073709551615}
    var v8 int64 = -1
    v9 := []int{0, -218, -531, -347}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 0
    v3 := make(map[uint32]uint32)
    v4 := []string{""}
    v5 := make(map[int]string)
    v6 := []string{("\n\t" + "hello"), (("foo bar" + "\n\t") + ("" + ((("foo bar" + ("hello" + ("a" + "foo bar"))) + "\n\t") + (("" + "hello") + "a")))), "hello", "hello"}
    var v10 bool = true
    var v11 float64 = ((((3.14159 * 1e10) + 0.1) / 1.0) * 3.14159)
    v12 := GenS1{F0: 1}
    var v13 uint16 = 32767
    if v10 {
        if v10 {
            fmt.Println(v10)
            if v10 {
                if true {
                    v10 = false
                    v13 += v13
                }
                switch ("a" + "foo bar") {
                case "\n\t":
                    v2 -= (v2 | v2)
                case "hello":
                    var v14 string = "\n\t"
                    _ = v14
                default:
                    if false {
                        v11 += float64(3.14159)
                    } else {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            var v16 int16 = 821
                            var v17 int8 = 100
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 14; fi18++ {
                            for fi19 := 0; fi19 < 8; fi19++ {
                                delete(v1, "a")
                                v20 := GenS1{F0: 65535}
                                fmt.Println(v2, fi19)
                                _ = fi19
                                _ = v20
                            }
                            v6 = append(v6, "\n\t")
                            _ = fi18
                        }
                        var v21 float64 = (v11 * (0.1 + v11))
                        _ = v21
                    }
                }
            } else {
                var v22 uint8 = 2
                var v23 int32 = v2
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 bool = v10
                v26 := GenS2{F0: -18, F1: 0}
                _ = len(v1)
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 12; fi27++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v28 := GenS3{F0: 2, F1: -1, F2: -128}
                v29 := GenS3{F0: 153, F1: 1, F2: 127}
                _ = fi27
                _ = v28
                _ = v29
            }
            if (42 != -82) {
                if v10 {
                    if (v2 >= (v2 ^ (-(v2)))) {
                        var v30 int = -48
                        var v31 uint64 = uint64(v13)
                        v32 := GenS4{F0: "", F1: "a", F2: 1}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                }
                fmt.Println(v2, v11)
                var v33 bool = v10
                _ = v33
            }
            var v34 int = -869
            _ = v34
        }
        var v35 uint8 = 116
        _ = v35
    }
    var v36 uint64 = 42
    v37 := GenS4{F0: "foo bar", F1: "hello", F2: 747}
    v36 = (^(v36))
    var v38 int16 = 0
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 uint64 = 975
        _ = fi39
        _ = v40
    }
    fmt.Println(v11)
    var v41 float32 = ((0.1 / 1e10) + float32(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}

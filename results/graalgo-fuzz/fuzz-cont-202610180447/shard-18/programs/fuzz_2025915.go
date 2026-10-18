package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 uint32 = 2147483647
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 1, F1: "hello", F2: ""}
    var v6 float32 = 0.0
    v7 := GenS2{F0: 42, F1: v3, F2: -9223372036854775808}
    v8 := GenS3{F0: (v3 ^ (^(v3)))}
    v3 = (v3 & (v3 | v3))
    fmt.Println(v3)
    if v2 {
        if (v2 && (-1 < 58)) {
            fmt.Println(v6)
            var v9 int16 = 42
            switch "foo bar" {
            case "hello":
                if v2 {
                    for fi10 := 0; fi10 < 17; fi10++ {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            var v12 uint = 65535
                            switch (fi10 - (fi10 * 128)) {
                            case -2147483648:
                                if (223 >= 32767) {
                                    switch fi10 {
                                    case 9223372036854775807:
                                        var v13 bool = (((false || (42 <= 0)) && (32767 != 255)) || v2)
                                        _ = v13
                                    }
                                    if (!false) {
                                        if v2 {
                                            fmt.Println(v2)
                                            fmt.Println(v9, v12)
                                            var v14 bool = false
                                            _ = v14
                                        } else {
                                            fmt.Println(v2)
                                            var v15 float32 = v6
                                            switch fi10 {
                                            case -132:
                                                var v16 bool = v2
                                                _ = v16
                                            case 4294967295:
                                                v17 := GenS1{F0: v12, F1: ("foo bar" + "\n\t"), F2: ""}
                                                _ = v17
                                            case -9223372036854775808:
                                                switch fi11 {
                                                case -2147483648:
                                                    var v18 uint64 = 0
                                                    _ = v18
                                                case 890:
                                                    v19 := []int{(fi11 << ((4294967295) & 63)), -378, fi11}
                                                    _ = v19
                                                default:
                                                    v20 := []int{(fi11 ^ (-(fi11))), -352}
                                                    _ = v20
                                                }
                                            default:
                                                var v21 uint64 = 18446744073709551615
                                                _ = v21
                                            }
                                            _ = v15
                                        }
                                    } else {
                                        if ((v2 && false) && v2) {
                                            v22 := GenS1{F0: 18446744073709551615, F1: ("hello" + "hello"), F2: "foo bar"}
                                            _ = v22
                                        }
                                        var v23 float64 = 0.1
                                        _ = v23
                                    }
                                    v24 := []int{}
                                    _ = v24
                                }
                            default:
                                v6 += v6
                            }
                            _ = fi11
                            _ = v12
                        }
                        _ = fi10
                    }
                }
            }
            _ = v9
        } else {
            var v25 int64 = 1
            if v2 {
                v2 = (-950 == 32768)
                var v26 int64 = (-(v25))
                _ = v26
            }
            switch 127 {
            case 127:
                if true {
                    v6 -= v6
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 int = 2147483647
                        for fi29 := 0; fi29 < 12; fi29++ {
                            if ((!(4294967295 <= 42)) || v2) {
                                v2 = v2
                                v30 := GenS2{F0: 1, F1: 2, F2: (v25 / ((1) | 1))}
                                _ = v30
                            }
                            if v2 {
                                delete(v1, (v3 | v3))
                                v3 = v3
                                v2 = (("hello" > "hello") && (false && ("\n\t" == ("a" + ("\n\t" + (("hello" + "foo bar") + "\n\t"))))))
                                var v31 uint8 = 42
                                _ = v31
                            }
                            var v32 float32 = v6
                            _ = fi29
                            _ = v32
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            }
            _ = v25
        }
    }
    v6 = ((0.1 + v6) - (v6 * ((float32(0.1) + -2.25) * 0.0)))
    if false {
        v33 := []int{1}
        v6 -= v6
        _ = v33
    }
    if (((-2.25 + (0.0 - (float64(v6) - -2.25))) + (3.14159 * float64(v6))) != ((float64(float32(((float64(-2.25) + 1.5) - ((-2.25 / -2.25) / float64(float32(0.1)))))) / -2.25) + 3.14159)) {
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16, p1 uint32, p2 uint16) int16 {
    v6 := []int{2147483647}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 bool = false
        var v9 uint64 = uint64(p1)
        p2 = 32767
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}

func main() {
    v1 := []uint32{}
    var v2 int = 0
    v3 := make(map[string]int)
    var v4 float64 = 0.0
    var v5 float32 = (((float32(0.0) / (1.5 / ((3.14159 - (-2.25 * 1.5)) + 1e10))) / 1.0) * float32((v4 / ((v4 - 1.5) / v4))))
    v4 = ((0.0 / v4) + (float64(v5) - -2.25))
    for fi10 := 0; fi10 < 11; fi10++ {
        v2 -= (fi10 ^ 206)
        _ = fi10
    }
    if len(v1) > 0 {
        v1[0] = 946
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: fi12, F1: 627}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 uint16 = 2
            fmt.Println(fi11)
            var v16 string = (("foo bar" + "") + ("foo bar" + (("" + "a") + "a")))
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    if (!(false || true)) {
        if len(v1) > 0 {
            v1[0] = 2
        }
    }
    var v17 float32 = v5
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(v5)
        _ = fi18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint64 = 0
        if (42 > 184) {
            for fi21 := 0; fi21 < 14; fi21++ {
                v22 := v3["\n\t"]
                if true {
                    var v23 bool = false
                    v5 = (3.14159 / v5)
                    _ = v23
                } else {
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    if true {
                        if ((!true) || false) {
                            v1 = append(v1, 4294967295)
                        } else {
                            v17 += (3.14159 - v5)
                        }
                        v24 := GenS2{F0: 494}
                        if false {
                            v20 += v20
                            var v25 float64 = v4
                            _ = v25
                        } else {
                            v26 := GenS2{F0: 18446744073709551615}
                            v27 := GenS3{F0: 18446744073709551615, F1: 32767}
                            var v28 int64 = 0
                            fmt.Println(fi21, v2, v22)
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = v24
                    }
                    fmt.Println(fi19, fi19)
                    if (true || false) {
                        if len(v1) > 0 {
                            v1[0] = 618
                        }
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    } else {
                        var v29 bool = false
                        var v30 int16 = -582
                        fmt.Println(v20)
                        _ = v29
                        _ = v30
                    }
                    var v31 uint8 = 2
                    _ = v31
                }
                switch "a" {
                case "hello":
                    _ = len(v3)
                case "foo bar":
                    v17 = v5
                    delete(v3, (("\n\t" + "foo bar") + ((("foo bar" + "hello") + "") + "foo bar")))
                    var v32 uint16 = 519
                    _ = v32
                case "a":
                    if false {
                        var v33 float64 = v4
                        v4 -= 3.14159
                        v3["foo bar"] = fi21
                        _ = len(v3)
                        _ = v33
                    } else {
                        v4 -= (v4 + float64(v17))
                        var v34 uint32 = 439
                        _ = v34
                    }
                default:
                    v4 = v4
                }
                _ = fi21
                _ = v22
            }
            var v35 uint = 883
            if true {
                var v36 uint8 = 222
                switch -1 {
                case -367:
                    switch fi19 {
                    case 42:
                        if true {
                            var v37 uint32 = 65535
                            var v38 uint64 = 18446744073709551615
                            _ = v37
                            _ = v38
                        }
                    case -2147483648:
                        for fi39 := 0; fi39 < 16; fi39++ {
                            v36 += v36
                            _ = fi39
                        }
                    }
                case 0:
                    for fi40 := 0; fi40 < 6; fi40++ {
                        _ = len(v1)
                        if true {
                            var v41 bool = (1 > -2147483648)
                            v20 = 0
                            _ = v41
                        }
                        _ = fi40
                    }
                default:
                    var v42 float64 = (((((v4 + float64(v17)) * v4) / (v4 / 1.5)) / 1.0) + ((float64(v5) / 0.1) * ((((v4 - v4) * 3.14159) / 1.5) / (v4 - v4))))
                    _ = v42
                }
                _ = len(v1)
                _ = v36
            } else {
                v43 := v3[("hello" + ("foo bar" + "\n\t"))]
                var v44 string = "\n\t"
                _ = v43
                _ = v44
            }
            _ = v35
        }
        _ = fi19
        _ = v20
    }
    v45 := GenS4{F0: 42}
    switch "hello" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        v17 -= v5
    default:
        switch 0 {
        case -621:
            if true {
                _ = len(v1)
                fmt.Println(v17, v17, v2)
                var v46 int64 = 9223372036854775807
                _ = v46
            } else {
                var v47 uint8 = 155
                var v48 uint64 = 18446744073709551615
                _ = v47
                _ = v48
            }
        }
    }
    fmt.Println(v4)
    v2 += 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}

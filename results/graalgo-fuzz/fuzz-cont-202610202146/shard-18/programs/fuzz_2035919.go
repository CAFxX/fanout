package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 string; F2 string }
func genFn0() string {
    if true {
        var v4 string = "foo bar"
        fmt.Println(v4)
        var v5 int16 = -1
        _ = v4
        _ = v5
    }
    var v6 uint16 = 65535
    v6 = v6
    _ = v6
    return "hello"
}
func genFn1(p0 uint8, p1 string) bool {
    v7 := GenS1{F0: 42}
    if ((((!true) || ((!true) && false)) || true) || false) {
        v8 := []int{127, 42, -9223372036854775808}
        v9 := GenS1{F0: 440}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    var v3 int = -1
    var v10 uint64 = 4294967295
    v11 := GenS1{F0: 10}
    switch "a" {
    case "":
        switch 127 {
        case 0:
            var v12 bool = true
            _ = v12
        }
    case "foo bar":
        if (("" + "hello") > "") {
            var v13 float64 = 1.5
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                switch "a" {
                case "foo bar":
                    if true {
                        v3 -= (v3 % ((v3) | 1))
                    } else {
                        fmt.Println(v3, v13)
                        v14 := []int{}
                        _ = v14
                    }
                case "":
                    v15 := []int{}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 12; fi16++ {
                        if true {
                            for fi17 := 0; fi17 < 8; fi17++ {
                                fmt.Println(fi17, v2, fi16)
                                _ = fi17
                            }
                            if false {
                                v10 = 2
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    fmt.Println(fi16, v10, v3)
                                    _ = fi18
                                }
                            } else {
                                for fi19 := 0; fi19 < 9; fi19++ {
                                    v2 -= v2
                                    v20 := GenS1{F0: -127}
                                    if true {
                                        switch "" {
                                        case "a":
                                            v21 := GenS2{F0: -693, F1: -2147483648, F2: "hello"}
                                            _ = v21
                                        default:
                                            v10 = (v10 & 1)
                                        }
                                    } else {
                                        v10 -= v2
                                        var v22 uint32 = 1
                                        _ = v22
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                if (512 >= 42) {
                                    fmt.Println(v10)
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        for fi23 := 0; fi23 < 10; fi23++ {
                                            switch "\n\t" {
                                            case "":
                                                switch fi16 {
                                                case -634:
                                                    fmt.Println(v3, fi23, v2)
                                                }
                                            case "\n\t":
                                                var v24 string = "foo bar"
                                                _ = v24
                                            }
                                            fmt.Println(fi23, v10)
                                            _ = fi23
                                        }
                                    case "":
                                        var v25 float32 = 1e10
                                        _ = v25
                                    }
                                    v3 += (v3 * 9223372036854775807)
                                }
                                if (2147483648 > 42) {
                                    v10 -= v10
                                }
                            }
                            var v26 bool = (true || (!(!true)))
                            _ = v26
                        }
                        _ = fi16
                    }
                }
            case "":
                switch v3 {
                case -386:
                    var v27 float32 = ((-2.25 - float32(3.14159)) - float32(v13))
                    _ = v27
                case 9223372036854775807:
                    v2 -= 635
                case 2147483647:
                    v28 := []int{(v3 ^ (v3 | v3)), (v3 >> ((255) & 63)), (v3 ^ (v3 & (-(v3))))}
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 15; fi29++ {
                        v2 += v2
                        _ = fi29
                    }
                }
            case "\n\t":
                fmt.Println(v2, v3)
            default:
                delete(v1, "a")
            }
            _ = v13
        }
    case "\n\t":
        var v30 uint = 65535
        _ = v30
    default:
        if (float64(0.0) != (1e10 - 0.0)) {
            var v31 int16 = 880
            var v32 uint8 = 103
            v33 := GenS3{F0: 42}
            _ = v31
            _ = v32
            _ = v33
        } else {
            fmt.Println(v10, v2, v2)
        }
    }
    fmt.Println(v10, v2)
    if (127 < 127) {
        for fi34 := 0; fi34 < 19; fi34++ {
            v3 = 2147483647
            if true {
                v35 := []int{fi34}
                fmt.Println(v10, fi34)
                _ = v35
            }
            _ = fi34
        }
    }
    v36 := GenS1{F0: 32767}
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := []int{(-(fi37)), v3}
        for fi39 := 0; fi39 < 5; fi39++ {
            if (42 != (v10 >> ((4294967295) & 63))) {
                fmt.Println(v10, fi39, fi39)
            } else {
                v40 := GenS1{F0: 1000}
                _ = v40
            }
            v1["\n\t"] = fi39
            _ = fi39
        }
        v41 := GenS4{F0: 0, F1: "a", F2: (("\n\t" + "foo bar") + "a")}
        _ = fi37
        _ = v38
        _ = v41
    }
    var v42 float32 = 1e10
    v43 := GenS1{F0: int16(v3)}
    var v44 uint8 = 133
    var v45 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}

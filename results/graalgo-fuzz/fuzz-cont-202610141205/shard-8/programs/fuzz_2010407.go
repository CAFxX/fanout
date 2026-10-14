package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int) uint64 {
    v4 := []int{(p1 << ((42) & 63)), (p1 / ((p1) | 1)), (p1 * (-(p1)))}
    if len(v4) > 0 {
        v4[0] = p1
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1(p0 uint32, p1 int8, p2 uint8) int8 {
    v5 := GenS1{F0: -9223372036854775808, F1: 127}
    if false {
        var v6 uint = 104
        _ = v6
    } else {
        p1 = (p1 * -26)
        v7 := []int{-9223372036854775808}
        switch -988 {
        case 2147483647:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int = fi8
                switch (("\n\t" + "a") + ("" + "\n\t")) {
                case "hello":
                    p1 += p1
                case "\n\t":
                    var v10 int32 = 1
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
        case 0:
            v11 := GenS1{F0: 9223372036854775807, F1: -462}
            _ = v11
        case 428:
            v12 := GenS1{F0: -1, F1: -2147483648}
            _ = v12
        }
        _ = v7
    }
    p0 -= 65535
    p2 = 127
    var v13 int32 = 508
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return (p1 | 1)
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (32767 < -32768)
    var v3 uint16 = 777
    v2 = (3.14159 != 0.1)
    var v14 float64 = (1e10 * 1e10)
    if false {
        v15 := []int{-2147483648, -2147483648, 42, 1}
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS1{F0: -887, F1: 9223372036854775807}
                var v19 uint = 2
                var v20 int16 = 42
                if v2 {
                    fmt.Println(v20, v19, v2)
                }
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            v21 := GenS1{F0: 834, F1: fi16}
            _ = fi16
            _ = v21
        }
        _ = v15
    } else {
        v2 = false
    }
    if (v3 <= 255) {
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 int16 = 857
            if (true && v2) {
                switch fi22 {
                case 0:
                    v24 := v1["\n\t"]
                    _ = v24
                }
                v2 = true
                v2 = ((!(false || false)) && (v2 && true))
            } else {
                v25 := []int{fi22, fi22, 42, -25}
                v14 -= 1e10
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        v26 := GenS2{F0: -2147483648, F1: 2147483647}
        _ = v26
    } else {
        _ = len(v1)
        var v27 float64 = v14
        _ = v27
    }
    v2 = ((false || false) || v2)
    v3 = 255
    if (false && (0 < -1)) {
        v2 = (!(2 > 849))
        if v2 {
            for fi28 := 0; fi28 < 8; fi28++ {
                if ("a" < "hello") {
                    switch ("hello" + "\n\t") {
                    case "a":
                        for fi29 := 0; fi29 < 10; fi29++ {
                            if false {
                                v30 := GenS2{F0: -134, F1: 42}
                                v31 := GenS2{F0: 1, F1: 2147483648}
                                v32 := []int{(fi29 * 287), -2147483648, fi28, fi28}
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                var v33 uint16 = (^(v3))
                                v34 := GenS3{F0: -9223372036854775808}
                                fmt.Println(v14)
                                _ = v33
                                _ = v34
                            }
                            var v35 int32 = -2147483648
                            _ = fi29
                            _ = v35
                        }
                    default:
                        var v36 float32 = (((float32(-2.25) + (3.14159 - (float32((v14 / (v14 + float64(0.0)))) - 3.14159))) - float32(((v14 / 1e10) / (3.14159 / v14)))) + (((0.1 * (0.0 * float32((v14 * float64(0.0))))) - (0.0 / (0.1 - (((0.1 * 3.14159) + (-2.25 + ((-2.25 - float32(float64(0.0))) * 0.0))) - 3.14159)))) * 0.1))
                        _ = v36
                    }
                    fmt.Println(v3)
                }
                switch ("\n\t" + "a") {
                case "a":
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v38 := GenS3{F0: 42}
                        v39 := GenS3{F0: 0}
                        var v40 string = ((((("foo bar" + "hello") + ("" + "")) + ("foo bar" + "hello")) + "hello") + "hello")
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                case "\n\t":
                    if (v2 && v2) {
                        v41 := GenS2{F0: 208, F1: 42}
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 1; fi42++ {
                            v43 := GenS1{F0: -1, F1: (fi28 ^ -9223372036854775808)}
                            var v44 int = fi28
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        v1[("foo bar" + "")] = fi28
                    }
                }
                _ = fi28
            }
            v3 = v3
        }
    } else {
        var v45 string = "hello"
        var v46 uint64 = 1
        switch 127 {
        case 0:
            var v47 int64 = 42
            _ = v47
        default:
            var v48 bool = (-9223372036854775808 <= -1)
            _ = v48
        }
        _ = v45
        _ = v46
    }
    var v49 uint16 = (^(v3))
    v14 = v14
    v2 = v2
    v2 = true
    v50 := GenS4{F0: 65535}
    fmt.Println(v49, v14, v49)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v49
    _ = v50
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 uint; F1 string }
func genFn0() int64 {
    var v4 uint8 = 67
    var v5 int16 = 1000
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 int8, p2 uint32) string {
    fmt.Println(p2, p1, p2)
    v6 := GenS1{F0: 4294967295, F1: 0, F2: (p1 * (p1 + p1))}
    var v7 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    var v8 float32 = (0.0 - 0.0)
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(v8)
        _ = fi9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    v3 += v3
    fmt.Println(v3)
    v3 = v3
    if (200 == 127) {
        v8 = (v8 * v8)
    }
    v3 -= v3
    if (false && (!false)) {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint8 = 2
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 32767, F1: 9223372036854775807, F2: 127}
        _ = v14
    } else {
        fmt.Println(v8)
        switch ("\n\t" + (("" + (("a" + "hello") + "a")) + ("hello" + "a"))) {
        case "hello":
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint32 = 2
                _ = fi15
                _ = v16
            }
        case "\n\t":
            var v17 uint8 = 0
            _ = v17
        case "":
            v3 += (^(v3))
        }
        var v18 bool = true
        _ = v18
    }
    if false {
        if (true && true) {
            for fi19 := 0; fi19 < 18; fi19++ {
                if ((97 >= 0) || false) {
                    if (!(v3 >= v3)) {
                        var v20 string = ("" + ("\n\t" + "a"))
                        if (0.0 > 1e10) {
                            for fi21 := 0; fi21 < 10; fi21++ {
                                _ = len(v2)
                                v20 = v20
                                _ = fi21
                            }
                        } else {
                            v22 := GenS1{F0: fi19, F1: 9223372036854775807, F2: 0}
                            _ = v22
                        }
                        _ = v20
                    }
                    switch 4294967295 {
                    case 636:
                        v3 -= 155
                    default:
                        var v23 float64 = 1e10
                        _ = v23
                    }
                    if (0 == 1) {
                        switch -9223372036854775808 {
                        case 32767:
                            fmt.Println(v3, v3)
                        case 42:
                            var v24 uint64 = 408
                            _ = v24
                        case 127:
                            v25 := GenS2{F0: 65535, F1: "a"}
                            _ = v25
                        default:
                            if true {
                                var v26 uint = uint(v3)
                                switch fi19 {
                                case -9223372036854775808:
                                    for fi27 := 0; fi27 < 19; fi27++ {
                                        v3 += v3
                                        var v28 uint = 42
                                        v26 = v28
                                        _ = fi27
                                        _ = v28
                                    }
                                case 406:
                                    if ((((v26 - v26) == v26) && ((v26 & (v26 + 1)) >= (^(v26)))) && false) {
                                        v3 = (^(v3))
                                    }
                                }
                                _ = v26
                            } else {
                                v29 := []int{fi19, -912, 32767, (fi19 / ((fi19) | 1))}
                                var v30 float32 = 1.5
                                _ = v29
                                _ = v30
                            }
                        }
                        var v31 string = ("\n\t" + "a")
                        _ = v31
                    }
                }
                _ = fi19
            }
        } else {
            if false {
                fmt.Println(v3, v8)
                var v32 int32 = 94
                v33 := []int{9223372036854775807, 32767}
                _ = v32
                _ = v33
            }
        }
    }
    switch -2147483648 {
    case -459:
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 float64 = 1.5
            _ = fi34
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 uint32, p1 uint16) uint64 {
    v7 := []int{127, 128, 0, 1}
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(p0, fi8)
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 bool = (3.14159 > 1e10)
            fmt.Println(p0, v10, fi9)
            _ = len(v7)
            _ = fi9
            _ = v10
        }
        if (1e10 == 0.1) {
            var v11 int16 = 1
            _ = v11
        } else {
            v12 := GenS1{F0: "hello", F1: (p0 & p0)}
            var v13 uint16 = (p1 ^ (p1 >> ((65535) & 63)))
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    p0 = uint32(p1)
    var v14 int64 = 0
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: ("foo bar" + ((("" + "hello") + "a") + "\n\t")), F1: (p0 & p0)}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 552
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 int32 = 542
    var v4 int = 1
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    var v17 int16 = -748
    v18 := GenS1{F0: "\n\t", F1: 2147483648}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int32 = -2147483648
        _ = fi19
        _ = v20
    }
    v17 += (v17 % (((v17 ^ v17)) | 1))
    v21 := GenS1{F0: "hello", F1: 65535}
    v3 -= (v3 - 42)
    if false {
        for fi22 := 0; fi22 < 17; fi22++ {
            fmt.Println(v17)
            switch ("\n\t" + (("a" + ("" + "a")) + (("\n\t" + ("hello" + "hello")) + "a"))) {
            case "a":
                v23 := GenS1{F0: ("\n\t" + ("a" + "a")), F1: 0}
                _ = v23
            case "":
                v17 = 0
            default:
                if (4294967295 < 969) {
                    var v24 int32 = (-(v3))
                    var v25 int16 = (^(v17))
                    v26 := v2[(-(fi22))]
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    var v27 int32 = (-(v3))
                    if ((false && false) && (!true)) {
                        if false {
                            var v28 uint32 = 4294967295
                            _ = v28
                        } else {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                var v30 uint = 80
                                var v31 bool = (v30 <= 65535)
                                v3 += (v3 - 466)
                                var v32 uint32 = 1
                                var v33 int32 = 2147483647
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        }
                        fmt.Println(v3, fi22, fi22)
                        if (0 <= v4) {
                            v17 -= (-(v17))
                        } else {
                            v3 -= v3
                            fmt.Println(v27, v17, fi22)
                        }
                    }
                    _ = v27
                }
            }
            _ = fi22
        }
    }
    var v34 int = v4
    var v35 int16 = -32768
    var v36 int64 = 1
    if (32767 == 2) {
        v17 -= v35
    }
    if false {
        v36 = (v36 | v36)
        var v37 float32 = 1.5
        _ = v37
    }
    v36 += v36
    var v38 uint32 = 70
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

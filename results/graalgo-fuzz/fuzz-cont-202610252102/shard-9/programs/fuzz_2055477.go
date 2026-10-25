package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0(p0 int8, p1 int16, p2 uint8) uint16 {
    p2 = (^(p2))
    if true {
        var v4 uint32 = 2147483648
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 bool = ((((1.5 - 0.1) * (float32(1e10) - (float32(0.1) + 1.5))) / 3.14159) < 3.14159)
    var v3 uint32 = 4294967295
    if (!v2) {
        var v5 float64 = 3.14159
        v6 := GenS1{F0: 90}
        _ = v5
        _ = v6
    } else {
        switch ("foo bar" + "foo bar") {
        case "a":
            v1 -= v3
        }
    }
    switch 2147483647 {
    case -9223372036854775808:
        fmt.Println(v1)
    case 32767:
        var v7 uint32 = (v1 % (((^(v1))) | 1))
        _ = v7
    case 32768:
        var v8 int16 = -1
        _ = v8
    }
    if v2 {
        var v9 uint8 = 2
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                v2 = (1 == -9223372036854775808)
                v1 -= (^(v1))
                var v12 uint16 = 454
                _ = fi11
                _ = v12
            }
            v3 = (v1 / (((v1 % ((1) | 1))) | 1))
            _ = fi10
        }
        var v13 int64 = -221
        _ = v13
        _ = v9
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        switch fi14 {
        case 943:
            if true {
                var v15 uint = 106
                var v16 uint16 = 724
                var v17 string = ""
                _ = v15
                _ = v16
                _ = v17
            } else {
                if false {
                    fmt.Println(v1)
                    for fi18 := 0; fi18 < 11; fi18++ {
                        switch (fi14 % (((fi18 % ((fi14) | 1))) | 1)) {
                        case 1:
                            v19 := []int{(fi18 / ((42) | 1))}
                            _ = v19
                        }
                        if v2 {
                            fmt.Println(fi18, v1)
                            v20 := GenS1{F0: 1}
                            _ = v20
                        }
                        v3 -= v3
                        _ = fi18
                    }
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v22 := []int{}
                        _ = fi21
                        _ = v22
                    }
                    v23 := GenS2{F0: 956, F1: 1}
                    v1 = (v1 * v3)
                    _ = v23
                }
            }
        default:
            v24 := GenS2{F0: 255, F1: -578}
            _ = v24
        }
        switch "foo bar" {
        case "a":
            v1 = (v3 ^ (^(v1)))
        case "\n\t":
            if (v2 || (v1 >= 498)) {
                var v25 string = "hello"
                fmt.Println(v2, v3, v25)
                var v26 uint32 = v1
                _ = v25
                _ = v26
            }
        }
        _ = fi14
    }
    var v27 string = "a"
    for fi28 := 0; fi28 < 20; fi28++ {
        v1 = v3
        for fi29 := 0; fi29 < 3; fi29++ {
            if (v2 || v2) {
                var v30 float32 = ((3.14159 / 1.0) / 0.1)
                v31 := []int{fi29, fi28, (fi28 | (-(fi28)))}
                _ = v30
                _ = v31
            }
            var v32 string = "foo bar"
            var v33 uint16 = 32767
            v34 := []int{(fi29 / ((fi28) | 1)), fi29, 42, (fi28 % ((fi29) | 1))}
            _ = fi29
            _ = v32
            _ = v33
            _ = v34
        }
        v27 = "\n\t"
        _ = fi28
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 uint32 = (v1 % ((v3) | 1))
        _ = fi35
        _ = v36
    }
    var v37 uint8 = 2
    v38 := GenS3{F0: -128, F1: (v1 - (v1 + 2147483648)), F2: 65535}
    fmt.Println(v3, v1, v27)
    var v39 uint8 = (v37 & (v37 / ((v37) | 1)))
    v39 += (v39 + (v39 & v37))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}

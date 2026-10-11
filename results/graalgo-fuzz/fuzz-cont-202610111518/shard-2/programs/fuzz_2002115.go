package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 int32; F2 uint8 }
func genFn0() int64 {
    switch (("\n\t" + ("foo bar" + "")) + "") {
    case "hello":
        if false {
            var v5 int = -1
            v5 = v5
            v6 := GenS1{F0: 9223372036854775807, F1: (true && true)}
            _ = v5
            _ = v6
        }
    case "":
        switch -640 {
        case 32767:
            var v7 string = "foo bar"
            _ = v7
        case 1:
            var v8 uint64 = 820
            _ = v8
        case 9223372036854775807:
            var v9 float64 = float64(((3.14159 / 0.1) - (0.1 + -2.25)))
            _ = v9
        default:
            if true {
                v10 := GenS2{F0: 561}
                _ = v10
            } else {
                v11 := GenS2{F0: 18446744073709551615}
                var v12 int16 = 0
                var v13 bool = false
                _ = v11
                _ = v12
                _ = v13
            }
        }
    case "\n\t":
        var v14 bool = false
        _ = v14
    }
    var v15 int16 = 32767
    var v16 float64 = 3.14159
    v17 := GenS3{F0: v15, F1: "\n\t"}
    v15 -= (v15 % ((v15) | 1))
    _ = v15
    _ = v16
    _ = v17
    return 42
}
func genFn1(p0 int64) uint32 {
    p0 -= (p0 & p0)
    for fi18 := 0; fi18 < 4; fi18++ {
        p0 += genFn0()
        var v19 string = (("" + "hello") + "\n\t")
        _ = fi18
        _ = v19
    }
    var v20 int16 = 1000
    _ = p0
    _ = v20
    return 592
}

func main() {
    var v1 bool = false
    var v2 uint16 = 2
    var v3 float64 = (0.1 * 3.14159)
    var v4 float64 = float64(0.0)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int32 = 1
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v2, v2)
            switch "" {
            case "hello":
                if (!true) {
                    switch fi23 {
                    case 9223372036854775807:
                        switch 128 {
                        case -1:
                            v25 := GenS4{F0: 127}
                            _ = v25
                        }
                    default:
                        switch (fi24 * int(v2)) {
                        case 0:
                            v3 += v4
                        }
                    }
                    if v1 {
                        fmt.Println(fi23, v2)
                    }
                    if (-267 <= 2147483647) {
                        v3 = 1e10
                        v4 -= v3
                    } else {
                        if (-74 <= 1000) {
                            v26 := GenS4{F0: 1}
                            v3 -= (3.14159 * (float64(-2.25) + 1e10))
                            _ = v26
                        } else {
                            v4 = v3
                            v27 := GenS5{F0: 0, F1: ((((("\n\t" + "hello") + ("a" + "")) + (("a" + "hello") + "\n\t")) + "a") + ("hello" + ("" + "")))}
                            _ = v27
                        }
                        if (42 == 1) {
                            v28 := GenS1{F0: 9223372036854775807, F1: false}
                            switch (fi23 % (((fi23 ^ (fi24 << ((18446744073709551615) & 63)))) | 1)) {
                            case 1:
                                v29 := []int{(fi24 + fi24), (fi23 + fi23), -15}
                                _ = v29
                            case -1:
                                v1 = v1
                            case -604:
                                var v30 int16 = -32768
                                _ = v30
                            default:
                                var v31 uint16 = (v2 >> ((0) & 63))
                                _ = v31
                            }
                            var v32 int64 = -901
                            _ = v28
                            _ = v32
                        }
                        var v33 float64 = (((v3 / v4) + v4) - v4)
                        _ = v33
                    }
                } else {
                    fmt.Println(v1, fi24)
                    v34 := []int{}
                    v2 += (v2 >> ((436) & 63))
                    _ = v34
                }
            case "foo bar":
                var v35 int = fi24
                _ = v35
            }
            _ = fi24
        }
        _ = fi23
    }
    v36 := GenS6{F0: ""}
    switch 42 {
    case 42:
        for fi37 := 0; fi37 < 5; fi37++ {
            fmt.Println(v3, v3, v2)
            v38 := []int{fi37, (fi37 * (fi37 * fi37)), 2147483647, fi37}
            _ = fi37
            _ = v38
        }
    }
    fmt.Println(v2, v3)
    var v39 uint8 = 255
    var v40 int32 = -1
    v41 := GenS7{F0: uint32(v39), F1: (-(v40)), F2: 1}
    var v42 int64 = -1
    for fi43 := 0; fi43 < 1; fi43++ {
        for fi44 := 0; fi44 < 12; fi44++ {
            v3 -= v4
            v45 := []int{-1, (^(fi44))}
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    var v46 uint16 = 42
    if true {
        v47 := []int{int(v42), 548}
        _ = v47
    } else {
        v1 = v1
    }
    if v1 {
        if v1 {
            fmt.Println(v46)
            for fi48 := 0; fi48 < 16; fi48++ {
                v3 = 3.14159
                _ = fi48
            }
            v49 := GenS7{F0: 65535, F1: v40, F2: (v39 - (v39 - (v39 << ((0) & 63))))}
            _ = v49
        }
        for fi50 := 0; fi50 < 4; fi50++ {
            var v51 uint32 = 65535
            if v1 {
                for fi52 := 0; fi52 < 2; fi52++ {
                    var v53 uint = 65535
                    _ = fi52
                    _ = v53
                }
                var v54 int64 = (v42 / ((v42) | 1))
                v42 -= 670
                _ = v54
            } else {
                var v55 int = fi50
                _ = v55
            }
            _ = fi50
            _ = v51
        }
        v56 := GenS5{F0: -1, F1: "\n\t"}
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
}

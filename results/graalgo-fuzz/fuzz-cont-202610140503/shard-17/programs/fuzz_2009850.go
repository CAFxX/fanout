package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0() int64 {
    if false {
        v5 := GenS1{F0: 1, F1: 0, F2: -32768}
        var v6 int16 = 32767
        _ = v5
        _ = v6
    } else {
        v7 := GenS2{F0: -1}
        v8 := GenS3{F0: 255}
        if true {
            switch "\n\t" {
            case "\n\t":
                switch -223 {
                case 364:
                    for fi9 := 0; fi9 < 18; fi9++ {
                        var v10 uint = 42
                        v11 := GenS1{F0: -762, F1: uint32(v10), F2: 32767}
                        if (622 != 2) {
                            v12 := GenS3{F0: 0}
                            var v13 uint = (v10 / ((v10) | 1))
                            _ = v12
                            _ = v13
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case 160:
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 int64 = 42
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v17 := GenS1{F0: 32767, F1: 909, F2: -32768}
                            v15 += (v15 | v15)
                            v15 = v15
                            v15 -= -1
                            v15 = (v15 / (((v15 & v15)) | 1))
                            _ = fi16
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                case 32767:
                    v18 := []int{-9223372036854775808, 127}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 20; fi19++ {
                        if (-113 < int64(fi19)) {
                            var v20 float32 = 0.0
                            _ = v20
                        }
                        _ = fi19
                    }
                }
            default:
                v21 := []int{}
                _ = v21
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 11; fi23++ {
                    v24 := []int{fi23, 32768, (fi23 | 424)}
                    _ = fi23
                    _ = v24
                }
                if false {
                    if (!true) {
                        var v25 string = "a"
                        _ = v25
                    }
                }
                var v26 uint32 = 895
                _ = fi22
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 7; fi27++ {
                v28 := GenS2{F0: (fi27 - (-(fi27)))}
                fmt.Println(fi27)
                _ = fi27
                _ = v28
            }
            var v29 uint64 = 0
            var v30 uint16 = 255
            _ = v29
            _ = v30
        }
        _ = v7
        _ = v8
    }
    switch 127 {
    case 860:
        var v31 int8 = -128
        _ = v31
    }
    return -589
}

func main() {
    var v1 uint = 4294967295
    var v2 float32 = (0.0 - -2.25)
    var v3 int16 = 32767
    v4 := make(map[uint32]uint32)
    v1 += (v1 - 299)
    if false {
        var v32 bool = false
        _ = v32
    } else {
        fmt.Println(v1, v2)
        var v33 int32 = 0
        _ = v33
    }
    var v34 uint32 = 42
    var v35 int = 308
    if ("a" != ("" + "")) {
        for fi36 := 0; fi36 < 7; fi36++ {
            switch "a" {
            case "\n\t":
                v3 += (v3 & v3)
            case "a":
                v1 -= v1
            }
            var v37 string = ("\n\t" + (("hello" + "") + (("hello" + "hello") + (("\n\t" + (("a" + "hello") + ("foo bar" + "a"))) + ("" + "foo bar")))))
            _ = fi36
            _ = v37
        }
    }
    v38 := v4[(v34 * (v34 >> ((v1) & 63)))]
    if ((^(v1)) != 574) {
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 int16 = v3
            if true {
                v41 := []int{9223372036854775807}
                _ = v41
            } else {
                v42 := GenS2{F0: -1}
                fmt.Println(v35)
                _ = v42
            }
            v43 := GenS4{F0: true, F1: v1}
            _ = fi39
            _ = v40
            _ = v43
        }
    } else {
        _ = len(v4)
    }
    v4[364] = v34
    v44 := GenS4{F0: true, F1: 715}
    v45 := GenS2{F0: 0}
    v4[v38] = (^(v34))
    var v46 int16 = 348
    switch ("a" + "foo bar") {
    case "hello":
        v1 = 0
    case "a":
        var v47 uint16 = 852
        _ = v47
    case "":
        v48 := GenS4{F0: false, F1: (v1 << ((v1) & 63))}
        _ = v48
    }
    var v49 int = 0
    fmt.Println(v34, v46)
    v50 := v4[2147483647]
    v51 := []int{v49}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v49
    _ = v50
    _ = v51
}

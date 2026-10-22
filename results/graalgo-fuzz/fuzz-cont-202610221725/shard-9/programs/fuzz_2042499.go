package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 uint }
type GenS4 struct { F0 uint; F1 bool; F2 uint }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (1.5 - 3.14159)
    var v3 bool = ((false && true) && true)
    v4 := GenS1{F0: 38, F1: 0, F2: 1000}
    if (1000 != 0) {
        v5 := GenS1{F0: -117, F1: -2147483648, F2: 147}
        var v6 string = "\n\t"
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 42, F1: -1, F2: 404}
        if (0.0 <= 1e10) {
            if (!true) {
                v2 += (v2 / v2)
            } else {
                var v8 uint16 = 65535
                switch ("" + "foo bar") {
                case "":
                    for fi9 := 0; fi9 < 3; fi9++ {
                        for fi10 := 0; fi10 < 7; fi10++ {
                            if ((!true) && true) {
                                v11 := GenS1{F0: -13, F1: 981, F2: -116}
                                v12 := v1[4294967295]
                                var v13 float64 = -2.25
                                _ = v11
                                _ = v12
                                _ = v13
                            } else {
                                v2 = v2
                                switch fi9 {
                                case 4294967295:
                                    var v14 int64 = 42
                                    _ = v14
                                case 1:
                                    var v15 uint = 18446744073709551615
                                    _ = v15
                                default:
                                    fmt.Println(fi10, v3, fi9)
                                }
                            }
                            _ = fi10
                        }
                        switch ("\n\t" + (("a" + "\n\t") + "foo bar")) {
                        case "foo bar":
                            v1[fi9] = fi9
                        case "\n\t":
                            v16 := GenS1{F0: 100, F1: -1, F2: 0}
                            _ = v16
                        case "a":
                            var v17 string = "\n\t"
                            _ = v17
                        }
                        v18 := GenS2{F0: v8}
                        _ = fi9
                        _ = v18
                    }
                default:
                    v3 = v3
                }
                _ = v8
            }
        } else {
            v2 = (1e10 + v2)
        }
        _ = v7
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 18; fi19++ {
        v3 = true
        switch "" {
        case "":
            var v20 uint32 = 266
            _ = v20
        }
        _ = fi19
    }
    switch "foo bar" {
    case "hello":
        v2 = v2
    case "foo bar":
        v3 = v3
    default:
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                if (!v3) {
                    var v23 uint = 387
                    fmt.Println(fi21)
                    v24 := v1[(fi22 << ((v23) & 63))]
                    v25 := []int{v24, -1, (-(fi22))}
                    v26 := v1[-2147483648]
                    if v3 {
                        if ((v23 | v23) != 4294967295) {
                            var v27 float64 = 1e10
                            _ = v27
                        } else {
                            if v3 {
                                v24 += fi21
                            } else {
                                fmt.Println(fi22)
                                var v28 float64 = 0.0
                                _ = v28
                            }
                            fmt.Println(v3, v3)
                            v29 := GenS3{F0: -32768, F1: 42, F2: v23}
                            _ = v29
                        }
                    }
                    v26 = fi21
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    var v30 bool = v3
                    _ = v30
                }
                v31 := GenS4{F0: 0, F1: (v3 || v3), F2: 696}
                _ = fi22
                _ = v31
            }
            v32 := GenS1{F0: -1, F1: 2147483647, F2: -1}
            _ = fi21
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        switch (fi33 * fi33) {
        case 1:
            v34 := []int{}
            _ = v34
        case 9223372036854775807:
            if false {
                v1[32767] = (-(fi33))
                v35 := []int{}
                _ = v35
            } else {
                v36 := []int{fi33, fi33, -691, fi33}
                var v37 int32 = 42
                fmt.Println(v2)
                _ = v36
                _ = v37
            }
        }
        _ = fi33
    }
    if v3 {
        v3 = false
        v38 := []int{128}
        _ = v38
    } else {
        v1[-934] = -2147483648
    }
    v3 = ((!false) || ("foo bar" < ("\n\t" + "foo bar")))
    switch (("" + ("\n\t" + ("\n\t" + "\n\t"))) + "a") {
    case "foo bar":
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}

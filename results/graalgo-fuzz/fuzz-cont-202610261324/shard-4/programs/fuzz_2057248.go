package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int64) int64 {
    switch "a" {
    case "hello":
        if true {
            v5 := []int{2147483647, -9223372036854775808, 42, 89}
            v6 := GenS1{F0: -1, F1: (false || (1e10 < 3.14159)), F2: "foo bar"}
            _ = v5
            _ = v6
        }
    case "\n\t":
        fmt.Println(p0)
    case "a":
        p0 -= p0
    }
    v7 := []int{-600}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return (p0 / (((p0 + p0)) | 1))
}
func genFn1(p0 int32, p1 int8) int64 {
    v8 := []int{-598, 1}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    var v1 int8 = -96
    var v2 uint32 = 743
    var v3 int32 = 1
    var v4 uint32 = 65535
    v1 -= v1
    v9 := []int{32767, 4294967295, 9223372036854775807, -1}
    fmt.Println(v3, v1)
    v1 += v1
    var v10 uint16 = 2
    switch 128 {
    case 4294967295:
        if (0 == 0) {
            v1 = v1
            var v11 uint64 = 4294967295
            var v12 bool = ((0.0 - (0.1 / 1.5)) >= -2.25)
            _ = v11
            _ = v12
        }
    case -9223372036854775808:
        if (true || true) {
            v13 := GenS2{F0: -912}
            v14 := GenS1{F0: 128, F1: (!true), F2: "hello"}
            _ = v13
            _ = v14
        } else {
            v3 -= v3
            if (true || false) {
                v2 = (^(v2))
                for fi15 := 0; fi15 < 1; fi15++ {
                    fmt.Println(v4)
                    _ = fi15
                }
                fmt.Println(v10, v4, v10)
            }
        }
    default:
        if false {
            v16 := GenS2{F0: 32767}
            v17 := GenS1{F0: 32767, F1: (true && ((v3 | (-(v3))) >= (-(v3)))), F2: ""}
            if true {
                if true {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 int32 = (v3 | 0)
                        _ = fi18
                        _ = v19
                    }
                    var v20 uint64 = 875
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        switch (^(fi21)) {
                        case -999:
                            v22 := GenS3{F0: int64(v4)}
                            _ = v22
                        case 703:
                            if true {
                                var v23 int32 = v3
                                _ = v23
                            } else {
                                var v24 string = ""
                                _ = v24
                            }
                        }
                        var v25 int8 = (v1 >> ((2) & 63))
                        var v26 uint16 = v10
                        _ = fi21
                        _ = v25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 3; fi27++ {
                        var v28 uint = 18446744073709551615
                        if ("\n\t" <= "") {
                            if (!true) {
                                var v29 uint32 = 2
                                v10 = (^(v10))
                                _ = v29
                            } else {
                                var v30 float32 = ((-2.25 / 1e10) - (1e10 + ((0.1 / ((-2.25 + 1.5) / 3.14159)) + 1.5)))
                                v30 = v30
                                _ = v30
                            }
                            switch 2147483647 {
                            case 9223372036854775807:
                                if false {
                                    var v31 uint16 = (v10 * (v10 | (v10 | (v10 << (((v28 | (^(v28)))) & 63)))))
                                    _ = v31
                                }
                            case 433:
                                var v32 int = fi27
                                _ = v32
                            }
                            var v33 uint8 = 141
                            _ = v33
                        }
                        var v34 int = int(v3)
                        _ = fi27
                        _ = v28
                        _ = v34
                    }
                }
            } else {
                v35 := GenS4{F0: (v10 / ((42) | 1))}
                _ = v35
            }
            _ = v16
            _ = v17
        }
    }
    var v36 bool = true
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := GenS4{F0: v10}
        v36 = false
        var v39 uint32 = v2
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v9
}

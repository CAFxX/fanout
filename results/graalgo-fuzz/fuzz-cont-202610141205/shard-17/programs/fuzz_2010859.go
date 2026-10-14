package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint8 }
func genFn0() int8 {
    var v5 int = -57
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{(v5 * fi6)}
        _ = fi6
        _ = v7
    }
    var v8 int32 = 1
    v5 -= v5
    v9 := GenS1{F0: 2, F1: 42}
    _ = v5
    _ = v8
    _ = v9
    return -106
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -2147483648
    var v3 int32 = (v2 << ((2) & 63))
    var v4 int8 = 97
    var v10 int = int(v2)
    v2 = v3
    if true {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int64 = 9223372036854775807
            v13 := GenS2{F0: 32767}
            _ = fi11
            _ = v12
            _ = v13
        }
        v4 -= v4
        v1[("hello" + "foo bar")] = (v10 ^ 9223372036854775807)
    }
    delete(v1, "foo bar")
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int64 = 0
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 42
    if ((692 > 0) || true) {
        if (0 != 32767) {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint64 = 1
                v4 -= v4
                _ = fi17
                _ = v18
            }
            v19 := GenS1{F0: 42, F1: genFn0()}
            v20 := []int{}
            _ = v19
            _ = v20
        } else {
            fmt.Println(v4)
            for fi21 := 0; fi21 < 2; fi21++ {
                if ((-(v4)) >= (^(v4))) {
                    var v22 int16 = 1
                    var v23 int32 = v2
                    _ = v22
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v10 -= (fi21 + v10)
                        v4 += v4
                        if (!(!false)) {
                            var v25 int16 = -6
                            if (!false) {
                                v26 := v1[("\n\t" + "")]
                                if (false || ((((float64(-2.25) * (3.14159 * 1e10)) - 1e10) != 1e10) && true)) {
                                    fmt.Println(v3)
                                }
                                _ = v26
                            }
                            switch "\n\t" {
                            case "a":
                                v27 := GenS3{F0: "foo bar", F1: -9223372036854775808, F2: v16}
                                _ = v27
                            default:
                                v28 := []int{fi24, v10, (v10 & (v10 ^ 4294967295))}
                                _ = v28
                            }
                            var v29 int32 = (-(v3))
                            var v30 int64 = int64(fi21)
                            switch ("hello" + ("hello" + ((("\n\t" + "\n\t") + ("a" + "hello")) + ("hello" + "\n\t")))) {
                            case "\n\t":
                                if false {
                                    v4 += 42
                                    var v31 string = "a"
                                    v32 := GenS4{F0: true, F1: "\n\t"}
                                    _ = v31
                                    _ = v32
                                }
                            case "":
                                _ = len(v1)
                            case "a":
                                if (false || true) {
                                    var v33 int16 = v25
                                    switch 778 {
                                    case 2147483647:
                                        var v34 bool = true
                                        _ = v34
                                    default:
                                        var v35 float32 = -2.25
                                        _ = v35
                                    }
                                    var v36 string = ((("\n\t" + "a") + ("foo bar" + ("a" + "\n\t"))) + "")
                                    _ = v33
                                    _ = v36
                                } else {
                                    v37 := []int{fi24, (v10 % ((int(v3)) | 1)), (^(fi21))}
                                    fmt.Println(fi24)
                                    if len(v37) > 0 {
                                        v37[0] = 539
                                    }
                                    fmt.Println(fi24)
                                    var v38 float64 = 3.14159
                                    delete(v1, ("hello" + ("hello" + "\n\t")))
                                    _ = v37
                                    _ = v38
                                }
                            }
                            _ = v25
                            _ = v29
                            _ = v30
                        }
                        _ = fi24
                    }
                    var v39 int16 = 1
                    var v40 int8 = (v4 - (v4 >> ((1) & 63)))
                    _ = v39
                    _ = v40
                }
                var v41 int64 = -391
                var v42 int8 = v4
                _ = fi21
                _ = v41
                _ = v42
            }
        }
        var v43 int16 = 0
        v3 = (v2 / ((-1) | 1))
        _ = v43
    }
    v44 := GenS5{F0: 128}
    v10 -= v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0() int {
    var v7 uint32 = 0
    switch ("a" + ("\n\t" + ("hello" + "\n\t"))) {
    case "hello":
        v7 = v7
    }
    var v8 uint8 = 255
    var v9 bool = false
    if v9 {
        var v10 int16 = -855
        _ = v10
    } else {
        v9 = ((v7 & 1) > v7)
        var v11 float64 = (float64(((0.1 + float32(0.0)) * ((-2.25 * 3.14159) / 1.0))) + (-2.25 + 1.5))
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return -853
}

func main() {
    v1 := []uint32{642, 1}
    var v2 bool = ((1.5 * 3.14159) > 3.14159)
    var v3 float32 = (1e10 * (3.14159 / float32((0.0 + -2.25))))
    var v4 float32 = (1.5 * 0.0)
    v5 := make(map[string]bool)
    var v6 int16 = -235
    var v12 uint = 18446744073709551615
    if v2 {
        v2 = ((42 >= 2) || (v12 < v12))
        if ((2 < 39) || ("a" != "")) {
            if true {
                var v13 int64 = 42
                fmt.Println(v3)
                _ = v13
            } else {
                var v14 uint16 = 32767
                v4 += v4
                var v15 uint = v12
                _ = v14
                _ = v15
            }
        }
    }
    switch 42 {
    case 2147483647:
        v16 := GenS1{F0: 32768, F1: 0, F2: -608}
        _ = v16
    case 985:
        var v17 float32 = (1.5 / -2.25)
        _ = v17
    }
    v3 -= v3
    v18 := v5[(("foo bar" + "") + "hello")]
    delete(v5, "")
    var v19 int = -1
    v19 = (^(v19))
    var v20 bool = (827 > 1)
    if v18 {
        for fi21 := 0; fi21 < 15; fi21++ {
            v18 = true
            switch "hello" {
            case "":
                switch 767 {
                case 127:
                    v22 := GenS2{F0: (v19 * v19)}
                    _ = v22
                case 128:
                    var v23 int8 = 116
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := GenS3{F0: (!(true && false)), F1: 2}
                    if ((32767 == 9223372036854775807) && (fi21 != fi24)) {
                        v12 = v12
                        var v26 int32 = -1
                        if true {
                            v12 = 255
                            var v27 int8 = -70
                            if v2 {
                                v1 = append(v1, 2)
                                var v28 int16 = (v6 / ((v6) | 1))
                                _ = v28
                            }
                            _ = v27
                        } else {
                            var v29 float64 = (3.14159 + ((1e10 - (1e10 / ((float64(-2.25) + (float64(v3) - (float64(0.0) / (0.1 - float64(v3))))) + 0.0))) + (3.14159 + (float64(float32((0.0 + 0.1))) * 0.0))))
                            _ = v29
                        }
                        delete(v5, "")
                        _ = v26
                    }
                    fmt.Println(v6, v2)
                    var v30 float64 = 1e10
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (!v2) {
                        var v31 int32 = 813
                        _ = v31
                    } else {
                        v5["a"] = v20
                        _ = len(v5)
                    }
                    if v2 {
                        _ = len(v5)
                        if (!(1 == 128)) {
                            if (1 > 0) {
                                switch "hello" {
                                case "hello":
                                    if (42 != 18446744073709551615) {
                                        v3 += ((((-2.25 - v3) * (3.14159 - (v4 / v3))) + float32(3.14159)) - v3)
                                    }
                                case "a":
                                    v20 = false
                                case "":
                                    var v32 int = int(v12)
                                    _ = v32
                                default:
                                    var v33 float64 = v30
                                    _ = v33
                                }
                                fmt.Println(fi21, v3)
                                var v34 float64 = (((v30 + 1e10) + 1.5) - v30)
                                _ = v34
                            } else {
                                var v35 uint64 = 1
                                var v36 int16 = v6
                                var v37 int8 = 127
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        } else {
                            var v38 float64 = v30
                            _ = len(v5)
                            _ = v38
                        }
                        var v39 bool = (228 >= 2)
                        _ = v39
                    }
                    _ = fi24
                    _ = v25
                    _ = v30
                }
            }
            _ = fi21
        }
        var v40 int32 = 2147483647
        var v41 int64 = int64(v40)
        _ = v40
        _ = v41
    }
    if (65535 != 4294967295) {
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 uint = (v12 & v12)
            _ = fi42
            _ = v43
        }
        var v44 int = (v19 >> ((880) & 63))
        _ = v44
    } else {
        v2 = (4294967295 >= 2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

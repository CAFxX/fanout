package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64; F1 int }
func genFn0() uint64 {
    var v7 string = "a"
    var v8 string = v7
    var v9 uint = 18446744073709551615
    v10 := []int{322}
    v11 := GenS1{F0: 128, F1: 271, F2: v7}
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 705
    var v3 float64 = 0.1
    var v4 uint64 = 0
    v5 := make(map[string]int)
    var v6 int = 4294967295
    v12 := []int{}
    var v13 int64 = -1
    _ = len(v1)
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v2, v6)
        v12 = append(v12, fi14)
        _ = fi14
    }
    v6 = 0
    switch int(v13) {
    case -1:
        var v15 float32 = (float32(v3) / 1e10)
        _ = v15
    }
    var v16 int16 = 0
    fmt.Println(v6, v2)
    for fi17 := 0; fi17 < 20; fi17++ {
        delete(v5, "hello")
        fmt.Println(v13, fi17)
        _ = fi17
    }
    if false {
        fmt.Println(v13, v2)
    }
    v18 := GenS2{F0: -48, F1: (v16 - 631)}
    for fi19 := 0; fi19 < 2; fi19++ {
        if (2147483648 < (v2 & uint32(v4))) {
            if false {
                v13 -= int64(v4)
            } else {
                fmt.Println(fi19, v13)
                if (2 == 65535) {
                    if false {
                        v12 = append(v12, (^(v6)))
                        v20 := GenS3{F0: genFn0(), F1: fi19}
                        fmt.Println(v6)
                        _ = v20
                    } else {
                        var v21 string = ("foo bar" + "")
                        for fi22 := 0; fi22 < 18; fi22++ {
                            v23 := GenS1{F0: 1, F1: v13, F2: "foo bar"}
                            switch (("" + ((v21 + "foo bar") + v21)) + (v21 + (v21 + v21))) {
                            case "a":
                                var v24 uint32 = 2
                                _ = v24
                            case "":
                                if false {
                                    var v25 uint = 2
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: -2147483648, F1: (-(v16))}
                                    v4 -= 1
                                    var v27 int8 = 127
                                    _ = v26
                                    _ = v27
                                }
                            }
                            if true {
                                var v28 uint16 = 782
                                var v29 bool = false
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int = v6
                                _ = v30
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                    var v31 uint32 = v2
                    var v32 float64 = v3
                    _ = v31
                    _ = v32
                }
            }
            var v33 bool = (!false)
            _ = v33
        }
        v34 := GenS3{F0: genFn0(), F1: -9223372036854775808}
        _ = fi19
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

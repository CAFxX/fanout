package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int32; F1 int64; F2 uint64 }

func main() {
    var v1 int = 127
    v2 := make(map[string]int)
    var v3 int32 = 42
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 1}
    var v6 uint8 = 127
    v1 = 32768
    v7 := GenS2{F0: (^(v6)), F1: 42}
    v1 -= 4294967295
    if true {
        v6 += 127
    }
    fmt.Println(v3)
    var v8 uint8 = (^(v6))
    var v9 int8 = -128
    if true {
        switch 127 {
        case 765:
            var v10 uint16 = 55
            _ = v10
        default:
            if false {
                for fi11 := 0; fi11 < 19; fi11++ {
                    switch (fi11 * (fi11 & int(v6))) {
                    case -213:
                        v12 := GenS2{F0: (v8 / ((v8) | 1)), F1: 510}
                        _ = v12
                    case 4294967295:
                        v13 := v4[v1]
                        _ = v13
                    case 32767:
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v15 := GenS1{F0: (v9 % ((v9) | 1))}
                            fmt.Println(v9, v1)
                            v16 := GenS3{F0: true, F1: "a", F2: v9}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    v3 = v3
                    _ = fi11
                }
                if true {
                    if false {
                        v2[("\n\t" + "\n\t")] = v1
                        if (1e10 < 0.0) {
                            v8 = (^(v6))
                        } else {
                            fmt.Println(v3, v1)
                            fmt.Println(v1, v8)
                        }
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            delete(v4, v1)
                        case "a":
                            var v17 int8 = (v9 << ((0) & 63))
                            _ = v17
                        default:
                            v18 := GenS4{F0: 1, F1: v9, F2: 128}
                            _ = v18
                        }
                    } else {
                        _ = len(v4)
                        v19 := []int{v1, (-(v1))}
                        _ = len(v4)
                        _ = v19
                    }
                    v20 := GenS1{F0: (v9 ^ (v9 + v9))}
                    _ = v20
                }
                if (false || true) {
                    if (0 < 1) {
                        for fi21 := 0; fi21 < 3; fi21++ {
                            v1 -= fi21
                            v2["a"] = (fi21 + (fi21 - (v1 | (v1 ^ -648))))
                            v2[("a" + "hello")] = fi21
                            _ = fi21
                        }
                        var v22 float32 = 3.14159
                        _ = v22
                    } else {
                        v3 = v3
                    }
                }
            }
        }
        v23 := GenS4{F0: 986, F1: 1, F2: v6}
        v24 := []int{v1, 4294967295}
        _ = v23
        _ = v24
    } else {
        var v25 uint64 = 31
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := v2["a"]
            _ = len(v2)
            v9 -= v9
            _ = fi26
            _ = v27
        }
        v28 := GenS5{F0: 0, F1: (v1 % (((v1 >> ((255) & 63))) | 1))}
        _ = v25
        _ = v28
    }
    v29 := GenS3{F0: true, F1: ("foo bar" + ("" + "\n\t")), F2: v9}
    var v30 uint16 = 32767
    v31 := GenS4{F0: 797, F1: (-(v9)), F2: v8}
    switch "foo bar" {
    case "a":
        var v32 float64 = 0.0
        _ = v32
    }
    v33 := []int{(v1 << ((255) & 63))}
    v34 := GenS6{F0: v3, F1: -1, F2: 582}
    if (!(uint64(v30) != 1)) {
        var v35 int64 = -9223372036854775808
        var v36 string = "\n\t"
        _ = v35
        _ = v36
    } else {
        var v37 float32 = (0.0 / 0.1)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

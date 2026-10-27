package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 uint32, p1 int8) bool {
    v6 := []int{}
    v7 := GenS1{F0: 1}
    v8 := GenS1{F0: 2}
    if false {
        var v9 uint16 = 982
        var v10 float32 = 0.0
        _ = v10
        _ = v9
    } else {
        _ = len(v6)
        _ = len(v6)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 float32 = (3.14159 - 0.1)
    v2 := make(map[string]int)
    var v3 int64 = 1
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    switch 32768 {
    case 81:
        var v11 int16 = 195
        _ = v11
    case -395:
        v12 := []int{32767}
        _ = v12
    default:
        var v13 uint8 = 91
        _ = v13
    }
    if true {
        var v14 int = -1
        _ = v14
    } else {
        fmt.Println(v3, v1)
    }
    var v15 string = ""
    var v16 uint64 = 0
    for fi17 := 0; fi17 < 20; fi17++ {
        v1 = v1
        _ = fi17
    }
    v1 -= (v1 * ((v1 / 1.0) / 0.1))
    var v18 uint8 = 1
    switch 128 {
    case 1:
        if true {
            if true {
                _ = len(v2)
                for fi19 := 0; fi19 < 3; fi19++ {
                    v20 := GenS1{F0: 0}
                    var v21 float32 = float32(1e10)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                v22 := []int{1}
                _ = v22
            } else {
                if (852 <= 430) {
                    if (("" + "") == "\n\t") {
                        fmt.Println(v16, v3, v16)
                    } else {
                        var v23 int = 4294967295
                        _ = v23
                    }
                } else {
                    switch 128 {
                    case -1:
                        if (0 != uint16(v18)) {
                            for fi24 := 0; fi24 < 7; fi24++ {
                                v25 := GenS2{F0: "", F1: v16}
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    var v27 bool = true
                                    fmt.Println(fi26)
                                    _ = fi26
                                    _ = v27
                                }
                                _ = fi24
                                _ = v25
                            }
                        } else {
                            var v28 int64 = 9223372036854775807
                            switch 128 {
                            case 1:
                                if true {
                                    var v29 uint = 65535
                                    v30 := GenS1{F0: v16}
                                    _ = v29
                                    _ = v30
                                }
                            case -9223372036854775808:
                                v16 = v16
                            case 770:
                                var v31 int64 = (v28 % ((v3) | 1))
                                _ = v31
                            }
                            _ = v28
                        }
                    case 42:
                        switch (("foo bar" + (v15 + "foo bar")) + ("a" + v15)) {
                        case "a":
                            v16 = (v16 / (((v16 | (v16 >> ((42) & 63)))) | 1))
                        case "hello":
                            var v32 int32 = -706
                            _ = v32
                        }
                    default:
                        v1 = 1.5
                    }
                    v15 = ("" + v15)
                    if (v16 <= v16) {
                        var v33 uint64 = 18446744073709551615
                        _ = v33
                    }
                }
            }
        } else {
            for fi34 := 0; fi34 < 3; fi34++ {
                switch (v15 + ("\n\t" + v15)) {
                case "hello":
                    if true {
                        var v35 int8 = 1
                        for fi36 := 0; fi36 < 6; fi36++ {
                            var v37 int8 = v35
                            _ = fi36
                            _ = v37
                        }
                        v38 := []int{127, fi34, fi34}
                        _ = v35
                        _ = v38
                    }
                }
                v39 := GenS2{F0: "", F1: 2}
                v1 = ((float32((0.0 * ((((0.0 + 1.5) * 0.0) / float64(v1)) - (float64(v1) * -2.25)))) / (v1 - -2.25)) * v1)
                _ = fi34
                _ = v39
            }
        }
    case 127:
        var v40 uint64 = v16
        _ = v40
    case 4294967295:
        fmt.Println(v1, v16)
    default:
        v3 -= v3
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

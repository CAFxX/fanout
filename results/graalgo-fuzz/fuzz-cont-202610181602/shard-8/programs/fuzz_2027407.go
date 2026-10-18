package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint32) uint16 {
    v5 := []int{403, -2147483648, 32767, 149}
    v6 := GenS1{F0: 4294967295, F1: 867, F2: ("" + ("foo bar" + ""))}
    v7 := GenS2{F0: "", F1: -105}
    p0 += (p0 | -27)
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1() string {
    switch 0 {
    case 0:
        var v9 string = "foo bar"
        _ = v9
    case 4294967295:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int = fi10
            v12 := GenS3{F0: 0}
            _ = fi10
            _ = v11
            _ = v12
        }
    case 42:
        var v13 uint16 = 726
        _ = v13
    default:
        var v14 int8 = -1
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        if true {
            if true {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 int8 = -100
                    if ((604 != 986) || (128 >= 255)) {
                        if false {
                            var v18 int8 = (^(v17))
                            var v19 float32 = 3.14159
                            _ = v18
                            _ = v19
                        } else {
                            v20 := GenS1{F0: 1, F1: -1, F2: "a"}
                            _ = v20
                        }
                        v21 := []int{128}
                        _ = v21
                    } else {
                        v17 = (v17 << ((274) & 63))
                        var v22 float64 = 3.14159
                        _ = v22
                    }
                    _ = fi16
                    _ = v17
                }
                var v23 uint16 = 615
                for fi24 := 0; fi24 < 12; fi24++ {
                    if false {
                        v23 = v23
                        var v25 int32 = -1
                        if ((-89 <= (v25 >> ((65535) & 63))) || false) {
                            v26 := []int{fi24, int(v23)}
                            v25 += 42
                            _ = v26
                        } else {
                            v27 := []int{}
                            var v28 string = ""
                            fmt.Println(v25, v23, v25)
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    } else {
                        v29 := []int{(fi15 & (fi15 ^ fi15)), fi15}
                        _ = v29
                    }
                    v30 := GenS4{F0: 42, F1: 1}
                    fmt.Println(fi15)
                    _ = fi24
                    _ = v30
                }
                _ = v23
            }
            var v31 int = (fi15 & -1)
            _ = v31
        }
        _ = fi15
    }
    v32 := []int{}
    _ = v32
    return ""
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{2147483648}
    v3 := make(map[int]int)
    v4 := []string{"\n\t", "foo bar", "", "hello"}
    if ((65535 <= 784) || (18446744073709551615 == 278)) {
        var v33 int32 = 1
        var v34 float32 = (float32(0.0) + (0.1 - (-2.25 * 0.0)))
        _ = v33
        _ = v34
    }
    var v35 float64 = (float64(-2.25) / float64(((0.0 / 1.0) - (-2.25 + 3.14159))))
    switch 832 {
    case -9223372036854775808:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    case 2147483647:
        if false {
            v35 = v35
            v36 := GenS5{F0: (^(v1))}
            _ = v36
        }
    case 128:
        var v37 uint64 = 33
        _ = v37
    }
    switch ("" + "") {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 1
        }
    case "":
        var v38 uint16 = 65535
        _ = v38
    default:
        var v39 bool = true
        _ = v39
    }
    v1 = (v1 ^ v1)
    v35 -= (float64(3.14159) * ((float64(((0.1 / 1e10) + 0.0)) - (v35 / v35)) / (v35 * 0.1)))
    v40 := GenS4{F0: -19, F1: 0}
    _ = len(v3)
    fmt.Println(v35)
    var v41 bool = ((((!(128 != 24)) && true) && true) && false)
    var v42 int = 9223372036854775807
    v35 += (1e10 + ((v35 + v35) - 0.1))
    v43 := GenS6{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}

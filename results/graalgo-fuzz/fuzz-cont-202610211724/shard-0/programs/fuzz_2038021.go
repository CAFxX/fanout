package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 int8, p1 uint8, p2 int32) string {
    v4 := GenS1{F0: ""}
    for fi5 := 0; fi5 < 9; fi5++ {
        if (-32768 <= -991) {
            for fi6 := 0; fi6 < 16; fi6++ {
                fmt.Println(p2)
                _ = fi6
            }
        } else {
            switch "" {
            case "foo bar":
                if (65535 >= 2) {
                    fmt.Println(p1, fi5, p1)
                    p0 -= p0
                } else {
                    p2 = (-(p2))
                    p0 -= p0
                }
            case "hello":
                fmt.Println(p2, p0, p1)
            case "":
                var v7 bool = (!false)
                _ = v7
            }
            for fi8 := 0; fi8 < 1; fi8++ {
                if false {
                    var v9 bool = false
                    _ = v9
                }
                _ = fi8
            }
        }
        p1 = p1
        v10 := []int{42}
        _ = fi5
        _ = v10
    }
    p0 -= p0
    v11 := []int{}
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return ""
}

func main() {
    var v1 int8 = 26
    v2 := make(map[uint32]string)
    v3 := []string{("" + "\n\t"), ""}
    if ((0.0 * (1e10 * -2.25)) <= 0.0) {
        var v12 int = 127
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int = (^(v12))
            v1 = (v1 / ((v1) | 1))
            _ = fi13
            _ = v14
        }
        if true {
            v15 := GenS1{F0: "a"}
            _ = v15
        }
        _ = v12
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if (-9223372036854775808 >= (-(fi16))) {
            var v17 float32 = -2.25
            if true {
                if false {
                    v18 := GenS1{F0: ("foo bar" + "")}
                    if (!true) {
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 int8 = -128
                            v3 = append(v3, ("\n\t" + "a"))
                            v21 := GenS2{F0: 42}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        v3 = append(v3, (("\n\t" + "foo bar") + "hello"))
                    }
                    v1 -= v1
                    _ = v18
                } else {
                    if ((-9223372036854775808 >= -9223372036854775808) && (65535 > 894)) {
                        v1 = (v1 << ((0) & 63))
                        var v22 float64 = 0.1
                        v17 = v17
                        _ = v22
                    }
                    fmt.Println(v1, v1, v1)
                    v23 := GenS2{F0: 42}
                    _ = v23
                }
                v1 = v1
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                v1 = v1
                switch "" {
                case "a":
                    var v25 float32 = 1.5
                    _ = v25
                }
                if false {
                    fmt.Println(v17)
                }
                _ = fi24
            }
            _ = v17
        }
        _ = fi16
    }
    v1 -= 0
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println(v1)
    v26 := GenS3{F0: 2147483647, F1: 127, F2: -2147483648}
    fmt.Println(v1)
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 uint8 = 2
        _ = fi27
        _ = v28
    }
    v1 = v1
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    _ = len(v3)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
}

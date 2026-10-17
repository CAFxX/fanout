package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
func genFn0(p0 uint64) int32 {
    p0 += 2
    p0 = 0
    _ = p0
    return 700
}
func genFn1() uint16 {
    var v5 bool = false
    v6 := []int{32767, -2147483648, -170, 32767}
    for fi7 := 0; fi7 < 1; fi7++ {
        v5 = true
        _ = fi7
    }
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := []uint32{1, 42, 4294967295}
    v2 := make(map[uint32]int)
    v3 := []int{-9223372036854775808, -625}
    var v4 float32 = 3.14159
    v8 := GenS1{F0: "\n\t"}
    v4 = v4
    v4 = -2.25
    if false {
        v9 := GenS1{F0: ("foo bar" + "hello")}
        if true {
            v4 = -2.25
            var v10 int32 = 2147483647
            _ = v10
        } else {
            v4 = v4
            v4 -= (v4 + v4)
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    fmt.Println(fi12)
                    v13 := GenS1{F0: "\n\t"}
                    v4 = ((((float32((1.5 / 1.5)) * (v4 - 1e10)) / v4) - v4) + (float32(-2.25) + v4))
                    _ = fi12
                    _ = v13
                }
                switch "hello" {
                case "foo bar":
                    switch (-(fi11)) {
                    case -2147483648:
                        v4 = ((0.0 - v4) - v4)
                    case 2147483647:
                        v14 := v2[448]
                        _ = v14
                    case 614:
                        switch "foo bar" {
                        case "foo bar":
                            v15 := GenS1{F0: "hello"}
                            _ = v15
                        default:
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v4 += v4
                                _ = fi16
                            }
                        }
                    default:
                        v4 += float32(float64(v4))
                    }
                default:
                    v4 -= v4
                }
                v2[1] = fi11
                _ = fi11
            }
        }
        _ = v9
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v4 = v4
        var v18 string = "hello"
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            v4 += v4
            _ = fi20
        }
        fmt.Println(fi19)
        _ = fi19
    }
    v21 := GenS1{F0: ("hello" + "")}
    switch "hello" {
    case "a":
        v4 += ((v4 - v4) / v4)
    case "foo bar":
        for fi22 := 0; fi22 < 3; fi22++ {
            if false {
                var v23 int32 = 2147483647
                fmt.Println(v4, v23)
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 bool = (!true)
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                v4 = v4
                if (-1 > genFn0(2)) {
                    v2[4294967295] = fi22
                    var v26 uint64 = 18446744073709551615
                    _ = v26
                }
            }
            v4 = -2.25
            _ = fi22
        }
    case "":
        var v27 uint8 = 44
        _ = v27
    }
    fmt.Println(v4)
    if (2 > genFn1()) {
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 int64 = -850
            var v30 uint64 = 0
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        v4 += 1.5
    }
    if (0 != -1) {
        v4 = v4
        for fi31 := 0; fi31 < 20; fi31++ {
            if false {
                v4 -= 3.14159
                v32 := GenS2{F0: ""}
                _ = v32
            }
            if (!true) {
                v4 = (v4 / float32(float64(v4)))
            } else {
                v1 = append(v1, 1)
                for fi33 := 0; fi33 < 20; fi33++ {
                    fmt.Println(v4, fi33, fi33)
                    v4 -= (((1.5 - v4) / (0.1 + v4)) / v4)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi33
                }
                v2[0] = 4294967295
            }
            for fi34 := 0; fi34 < 13; fi34++ {
                if true {
                    var v35 int64 = 0
                    _ = v35
                }
                _ = fi34
            }
            _ = fi31
        }
    }
    if (!true) {
        if len(v1) > 0 {
            v1[0] = 123
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            v4 += v4
            _ = fi36
        }
    }
    v4 = (v4 + (((v4 * -2.25) * float32(float64(float32(3.14159)))) + ((0.0 * (v4 / 1.0)) - float32(0.1))))
    if ((("" + ("" + "\n\t")) + "") > "\n\t") {
        var v37 float32 = (0.1 + (3.14159 + float32(1.5)))
        fmt.Println(v4, v37)
        _ = v37
    }
    var v38 int8 = -1
    v4 = (v4 / v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v8
}

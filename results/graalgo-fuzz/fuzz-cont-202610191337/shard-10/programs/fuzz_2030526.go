package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint }
func genFn0() bool {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{(fi5 * fi5)}
        _ = fi5
        _ = v6
    }
    if true {
        switch 32767 {
        case 0:
            switch 968 {
            case -2147483648:
                var v7 int32 = -2147483648
                _ = v7
            }
        }
        var v8 int64 = 54
        v8 = (v8 * (v8 ^ v8))
        _ = v8
    }
    return true
}
func genFn1(p0 uint, p1 uint64) uint32 {
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            var v10 float32 = (((0.0 / 1e10) * -2.25) - 0.0)
            v11 := GenS1{F0: 0, F1: (("a" + "foo bar") > "hello")}
            _ = v10
            _ = v11
        }
        var v12 bool = true
        _ = fi9
        _ = v12
    }
    p1 -= (p1 | 4294967295)
    v13 := GenS1{F0: -1, F1: (255 == 776)}
    _ = p0
    _ = p1
    _ = v13
    return uint32(p0)
}

func main() {
    var v1 int32 = 2147483647
    var v2 int = -510
    var v3 uint = 255
    v4 := []uint32{1, 0, 310}
    var v14 uint = (v3 << ((v3) & 63))
    for fi15 := 0; fi15 < 3; fi15++ {
        v1 -= v1
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 string = "hello"
        _ = fi16
        _ = v17
    }
    var v18 uint = 1
    var v19 int64 = 1
    switch v2 {
    case -9223372036854775808:
        if len(v4) > 0 {
            v4[0] = 2
        }
    case 9223372036854775807:
        var v20 float64 = -2.25
        _ = v20
    default:
        if false {
            v3 = (v14 * (v18 ^ v3))
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 string = ("" + "")
                _ = fi21
                _ = v22
            }
        } else {
            v18 -= v14
            switch ("\n\t" + ("\n\t" + "")) {
            case "hello":
                for fi23 := 0; fi23 < 10; fi23++ {
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    switch ((("foo bar" + "") + (((("a" + "\n\t") + "foo bar") + "") + "")) + ("" + "foo bar")) {
                    case "a":
                        fmt.Println(v1)
                    default:
                        for fi24 := 0; fi24 < 4; fi24++ {
                            var v25 int16 = -807
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = fi23
                }
            default:
                if false {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v3 = v18
                        for fi27 := 0; fi27 < 20; fi27++ {
                            if true {
                                _ = len(v4)
                                v28 := GenS2{F0: 283, F1: ""}
                                _ = v28
                            } else {
                                v29 := GenS3{F0: "hello", F1: (v3 | v18), F2: 190}
                                var v30 uint16 = 45
                                v31 := GenS3{F0: ((((("\n\t" + "hello") + "hello") + "a") + ("foo bar" + "")) + ""), F1: (v18 >> ((42) & 63)), F2: 127}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            v1 = (v1 | 2147483647)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi27
                        }
                        _ = fi26
                    }
                } else {
                    v32 := GenS2{F0: 255, F1: "a"}
                    v3 = v14
                    var v33 int16 = 32767
                    _ = v32
                    _ = v33
                }
            }
            fmt.Println(v14)
        }
    }
    switch (("foo bar" + "") + "a") {
    case "":
        v34 := GenS4{F0: v18}
        _ = v34
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch "\n\t" {
    case "":
        v19 += -1
    }
    v4 = append(v4, genFn1(v18, 42))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

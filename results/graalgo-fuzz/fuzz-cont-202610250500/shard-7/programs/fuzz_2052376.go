package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32, p1 string) bool {
    if (0 > 1) {
        fmt.Println(p1, p0)
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := GenS1{F0: 0}
            _ = fi5
            _ = v6
        }
        fmt.Println(p0, p0)
    } else {
        switch 1 {
        case 0:
            for fi7 := 0; fi7 < 6; fi7++ {
                if (false && false) {
                    switch (p1 + ("foo bar" + "a")) {
                    case "hello":
                        if false {
                            v8 := GenS2{F0: true}
                            var v9 uint64 = 0
                            _ = v8
                            _ = v9
                        } else {
                            if true {
                                var v10 string = "a"
                                v10 = "hello"
                                fmt.Println(p0, v10)
                                _ = v10
                            } else {
                                var v11 float64 = 3.14159
                                var v12 uint8 = 13
                                _ = v11
                                _ = v12
                            }
                        }
                    case "\n\t":
                        fmt.Println(fi7, p0, p0)
                    default:
                        var v13 int32 = 1
                        _ = v13
                    }
                    var v14 int64 = 1
                    fmt.Println(p1)
                    _ = v14
                }
                fmt.Println(p0)
                _ = fi7
            }
        }
        var v15 uint64 = 0
        fmt.Println(p0, p0, p1)
        _ = v15
    }
    var v16 float32 = float32(0.0)
    if false {
        switch p1 {
        case "hello":
            var v17 uint8 = 5
            _ = v17
        }
        v18 := []int{4294967295, 375, -582}
        if (2 <= 866) {
            if false {
                v16 -= v16
                if len(v18) > 0 {
                    v18[0] = -83
                }
                var v19 int8 = -128
                _ = v19
            }
            p0 = 81
        } else {
            p1 = ("" + (p1 + "hello"))
            var v20 uint16 = 255
            _ = v20
        }
        _ = v18
    }
    switch "hello" {
    case "hello":
        var v21 int16 = 0
        _ = v21
    default:
        v22 := []int{128}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v16
    return true
}

func main() {
    v1 := []int{}
    var v2 int8 = 42
    var v3 float64 = 1e10
    var v4 float64 = 1e10
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v4 -= (0.0 + float64((1e10 + ((-2.25 + -2.25) - float32(v4)))))
            if ((!true) && false) {
                if len(v1) > 0 {
                    v1[0] = -1
                }
                var v25 bool = true
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v26 int32 = -873
                _ = v25
                _ = v26
            } else {
                v1 = append(v1, fi24)
                v27 := GenS3{F0: 1}
                v28 := GenS4{F0: 837}
                fmt.Println(v3, fi23)
                v29 := GenS2{F0: (900 > 0)}
                if true {
                    var v30 int8 = v2
                    _ = v30
                }
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi24
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v4 += ((v3 + 1.5) * (v3 * v3))
            var v32 int16 = 970
            switch "foo bar" {
            case "hello":
                var v33 float32 = 1.5
                _ = v33
            default:
                if false {
                    v2 = v2
                    fmt.Println(fi23, v3, v2)
                }
            }
            _ = fi31
            _ = v32
        }
        _ = fi23
    }
    _ = len(v1)
    switch 32767 {
    case 128:
        v2 -= v2
    case 127:
        v3 -= 0.0
    case -610:
        v4 = 3.14159
    default:
        for fi34 := 0; fi34 < 6; fi34++ {
            if false {
                v2 -= (v2 | v2)
                _ = len(v1)
                _ = len(v1)
            } else {
                v3 = 3.14159
            }
            _ = len(v1)
            _ = fi34
        }
    }
    v4 += (v4 * v3)
    fmt.Println(v3)
    var v35 uint8 = 233
    for fi36 := 0; fi36 < 20; fi36++ {
        var v37 uint8 = v35
        var v38 int16 = 538
        fmt.Println(v38, v38)
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v4, v35)
    v2 = 14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}

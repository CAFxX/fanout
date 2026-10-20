package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 bool) int {
    fmt.Println(p0, p1)
    var v7 int8 = 1
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    v3 := []string{}
    v4 := []string{"", ("a" + "foo bar"), ""}
    var v5 int16 = 42
    v6 := make(map[string]string)
    var v8 float32 = 0.1
    if (true || (true && false)) {
        if true {
            v5 -= (-(v5))
            var v9 int64 = 986
            _ = v9
        }
        v5 += v5
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v3 = append(v3, "\n\t")
        fmt.Println(v2)
        v11 := GenS1{F0: ((!false) && false)}
        _ = fi10
        _ = v11
    }
    if false {
        var v12 int8 = (^(v2))
        _ = v12
    }
    var v13 uint8 = 239
    fmt.Println(v8, v2, v2)
    v14 := GenS2{F0: v13, F1: -2147483648}
    if (true || (!false)) {
        var v15 string = "hello"
        v13 = v13
        _ = v15
    }
    v5 = (v5 & 0)
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int64 = 698
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v19 uint16 = 42
            _ = fi17
            _ = v18
            _ = v19
        }
        v2 += v2
        _ = fi16
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "hello" {
        case "hello":
            _ = len(v1)
        case "foo bar":
            if ((v5 % ((v5) | 1)) > 42) {
                var v21 uint16 = 905
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            switch "" {
            case "\n\t":
                v23 := GenS1{F0: (1.5 <= 3.14159)}
                var v24 uint = 507
                _ = v23
                _ = v24
            }
            v13 = 1
            if (!false) {
                var v25 int16 = v5
                var v26 uint = 738
                _ = v25
                _ = v26
            } else {
                if false {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(fi20, v2)
                    }
                    if true {
                        v8 -= float32((1.5 / -2.25))
                    } else {
                        var v27 int8 = (^(v2))
                        switch (fi22 / ((fi20) | 1)) {
                        case -2147483648:
                            v13 = (v13 + v13)
                        case -9223372036854775808:
                            var v28 uint64 = 18446744073709551615
                            _ = v28
                        }
                        _ = v27
                    }
                    switch (("a" + "\n\t") + "") {
                    case "hello":
                        v29 := v6["\n\t"]
                        _ = v29
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            }
            _ = fi22
        }
        for fi30 := 0; fi30 < 1; fi30++ {
            if (0 != 2147483648) {
                var v31 int8 = (v2 << ((2) & 63))
                _ = v31
            } else {
                v8 = (1e10 - v8)
                switch "a" {
                case "\n\t":
                    if false {
                        var v32 int = (fi30 * -599)
                        _ = v32
                    } else {
                        if true {
                            var v33 uint16 = 0
                            _ = v33
                        }
                    }
                default:
                    fmt.Println(v2, v13)
                }
                switch fi30 {
                case -34:
                    switch (fi30 + (fi30 / ((fi20) | 1))) {
                    case 42:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        _ = len(v1)
                    }
                default:
                    if (!true) {
                        v5 -= v5
                        var v34 uint32 = 606
                        _ = v34
                    } else {
                        if false {
                            var v35 uint32 = 2147483648
                            _ = v35
                        } else {
                            v4 = append(v4, ("\n\t" + ("\n\t" + ("foo bar" + ("\n\t" + "\n\t")))))
                        }
                        var v36 uint16 = 101
                        delete(v1, 907)
                        _ = v36
                    }
                }
            }
            fmt.Println(v13, v13)
            var v37 int64 = -1
            _ = fi30
            _ = v37
        }
        _ = fi20
    }
    v13 -= v13
    v38 := GenS1{F0: false}
    v8 += (v8 * 0.1)
    fmt.Println(v8)
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := GenS1{F0: false}
        var v41 int64 = -1
        if true {
            v42 := GenS3{F0: (v5 / ((32767) | 1))}
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            _ = v42
        } else {
            v2 += (v2 | 0)
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := v1[2147483648]
        v1[4294967295] = v44
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

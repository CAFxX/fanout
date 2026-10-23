package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8) int8 {
    if true {
        var v4 uint = 0
        _ = v4
    } else {
        v5 := []int{1, 1}
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 bool = (p1 > p1)
            var v9 bool = v8
            v10 := []int{9223372036854775807}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        p1 = p1
        _ = fi6
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int8 = 33
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    return -6
}
func genFn1(p0 int16, p1 bool, p2 uint8) string {
    p0 = p0
    v13 := []int{-9223372036854775808, -1, 0, 127}
    var v14 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return "\n\t"
}

func main() {
    v1 := []string{"\n\t", ("a" + "a"), "\n\t", ""}
    v2 := make(map[string]uint32)
    v3 := []int{-2147483648, -396, -2147483648, 9223372036854775807}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 float64 = 1e10
        var v17 uint8 = 19
        v17 -= (v17 + (v17 ^ v17))
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS1{F0: ((("" + ((("" + ("" + "foo bar")) + (("a" + ((("" + "hello") + "hello") + "foo bar")) + ("foo bar" + ""))) + "hello")) + (("hello" + "\n\t") + "hello")) + "\n\t")}
        if (57 <= genFn0(65535, 67)) {
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int32 = 682
                if (!((int64(v21) > -407) || true)) {
                    if false {
                        v22 := GenS1{F0: ""}
                        fmt.Println(fi18, v21, fi18)
                        _ = v22
                    }
                    v23 := GenS1{F0: "hello"}
                    _ = v23
                } else {
                    v24 := GenS1{F0: "foo bar"}
                    fmt.Println(fi20)
                    _ = v24
                }
                if false {
                    var v25 uint = 0
                    switch (fi18 % ((382) | 1)) {
                    case -695:
                        if true {
                            var v26 uint16 = uint16(v25)
                            var v27 uint8 = 4
                            v25 += v25
                            _ = v26
                            _ = v27
                        }
                    case 4294967295:
                        v1 = append(v1, "hello")
                    case -1:
                        v28 := GenS2{F0: 128}
                        _ = v28
                    }
                    var v29 int8 = 1
                    _ = v25
                    _ = v29
                }
                _ = fi20
                _ = v21
            }
        }
        _ = fi18
        _ = v19
    }
    var v30 int64 = -1
    if (4294967295 < 369) {
        v31 := GenS2{F0: 127}
        v30 = v30
        _ = v31
    }
    v32 := GenS1{F0: ("a" + "")}
    if (-91 <= 746) {
        var v33 int64 = v30
        switch 1 {
        case -771:
            fmt.Println(v30, v30)
        case -1:
            v33 = (v30 & (-(v33)))
        }
        _ = v33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        v2[(("\n\t" + "foo bar") + "\n\t")] = 611
        _ = fi34
    }
    v35 := GenS1{F0: "a"}
    for fi36 := 0; fi36 < 3; fi36++ {
        v30 = (v30 / ((v30) | 1))
        for fi37 := 0; fi37 < 14; fi37++ {
            fmt.Println(fi36, fi36, v30)
            v30 = (v30 * -70)
            fmt.Println(fi37)
            v30 += (-(v30))
            v38 := GenS3{F0: 128, F1: "\n\t", F2: 127}
            _ = fi37
            _ = v38
        }
        v39 := GenS4{F0: 63}
        _ = fi36
        _ = v39
    }
    var v40 int8 = -1
    if false {
        for fi41 := 0; fi41 < 1; fi41++ {
            v30 = v30
            v30 -= -1
            _ = fi41
        }
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        fmt.Println(fi42)
        var v43 int = (^(fi42))
        _ = fi42
        _ = v43
    }
    v2["a"] = 65535
    switch "" {
    case "foo bar":
        if (-1 > (v30 ^ -360)) {
            var v44 uint = 18446744073709551615
            if false {
                var v45 uint32 = 70
                v46 := GenS5{F0: true}
                v45 = 740
                _ = v45
                _ = v46
            }
            _ = v44
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v35
    _ = v40
}

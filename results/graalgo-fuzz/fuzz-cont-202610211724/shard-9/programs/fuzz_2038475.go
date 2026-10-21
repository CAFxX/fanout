package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 string; F1 int; F2 int16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 uint8; F2 int32 }
func genFn0(p0 bool, p1 string, p2 bool) string {
    p2 = p2
    v7 := []int{1, 597, 595, 127}
    for fi8 := 0; fi8 < 20; fi8++ {
        p0 = p2
        var v9 uint = 4294967295
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 bool = (!(128 <= (fi10 & (fi10 % (((fi10 / ((fi10) | 1))) | 1)))))
        var v12 bool = true
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1() uint {
    v13 := []int{2147483647, -9223372036854775808}
    switch ("hello" + ("foo bar" + "foo bar")) {
    case "":
        var v14 uint = 401
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 uint = 497
            _ = fi15
            _ = v16
        }
    case "hello":
        var v17 string = "a"
        _ = v17
    }
    var v18 uint = 65535
    v18 += v18
    _ = v13
    _ = v18
    return v18
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v5 := []uint32{484, 421, 294}
    var v6 int64 = 42
    fmt.Println(v6)
    fmt.Println(v6)
    var v19 float64 = 3.14159
    v20 := GenS1{F0: 4294967295, F1: 42, F2: 2147483647}
    fmt.Println(v19, v6)
    if ("foo bar" >= "") {
        var v21 int8 = 0
        fmt.Println(v19)
        if false {
            v6 -= (v6 * (-(v6)))
        } else {
            fmt.Println(v19, v19)
        }
        _ = v21
    } else {
        if ("foo bar" > ("hello" + ("hello" + ((((("\n\t" + "\n\t") + ("hello" + ((("foo bar" + (("" + ("foo bar" + "hello")) + "")) + "") + "a"))) + "a") + "") + "hello")))) {
            var v22 uint32 = 413
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(v22)
                if true {
                    v24 := GenS1{F0: -304, F1: genFn1(), F2: v22}
                    v25 := GenS2{F0: 0, F1: v6}
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 int8 = -128
                    _ = fi26
                    _ = v27
                }
                _ = fi23
            }
            var v28 uint32 = 129
            _ = v22
            _ = v28
        }
    }
    v5 = append(v5, 231)
    if (-2147483648 < -669) {
        if len(v5) > 0 {
            v5[0] = 1
        }
        if true {
            for fi29 := 0; fi29 < 3; fi29++ {
                v19 += v19
                _ = fi29
            }
            switch ("a" + "\n\t") {
            case "":
                fmt.Println(v6, v19)
            case "foo bar":
                v30 := GenS3{F0: "foo bar", F1: 128, F2: 172}
                _ = v30
            default:
                v6 += -9223372036854775808
            }
        }
        if false {
            v19 -= ((v19 + float64(((((3.14159 + 1e10) / -2.25) * ((0.1 - ((((0.1 * 3.14159) + 1e10) * 3.14159) + 1e10)) / -2.25)) - (1.5 * (1e10 / float32(v19)))))) * v19)
        } else {
            switch 32768 {
            case 32768:
                for fi31 := 0; fi31 < 19; fi31++ {
                    v5 = append(v5, 0)
                    fmt.Println(fi31, v6, fi31)
                    v6 -= v6
                    _ = fi31
                }
            case 127:
                v19 = 0.0
            case -658:
                v19 = (3.14159 * v19)
            }
            switch 43 {
            case -729:
                for fi32 := 0; fi32 < 10; fi32++ {
                    v19 -= v19
                    v33 := GenS4{F0: 1, F1: (^(v6))}
                    v34 := GenS5{F0: "\n\t"}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            case 1:
                fmt.Println(v6, v19)
            default:
                delete(v2, 1)
            }
            var v35 int16 = 923
            _ = v35
        }
    }
    var v36 int32 = 42
    switch "foo bar" {
    case "\n\t":
        for fi37 := 0; fi37 < 16; fi37++ {
            v5 = append(v5, 1)
            switch (fi37 << ((334) & 63)) {
            case 9223372036854775807:
                v36 += 42
            case 127:
                switch "hello" {
                case "a":
                    var v38 uint = 254
                    _ = v38
                default:
                    switch "" {
                    case "a":
                        var v39 uint16 = 2
                        _ = v39
                    default:
                        if (157 > 973) {
                            v1[0] = "foo bar"
                        } else {
                            fmt.Println(v19, v6, fi37)
                            v19 += 0.0
                        }
                    }
                }
            case -470:
                var v40 int8 = -1
                _ = v40
            }
            if (false && (true || (345 == 255))) {
                fmt.Println(fi37, v19)
                var v41 int16 = -1
                var v42 float32 = -2.25
                _ = v41
                _ = v42
            }
            _ = fi37
        }
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        switch ("a" + "foo bar") {
        case "\n\t":
            var v44 int = -144
            _ = v44
        }
        _ = fi43
    }
    v45 := v3[32768]
    var v46 int64 = (v6 - (v6 + (v6 & v6)))
    fmt.Println(v45)
    if true {
        v47 := GenS6{F0: 43, F1: 42, F2: v36}
        var v48 uint32 = (v45 << ((286) & 63))
        var v49 string = "\n\t"
        _ = v47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        _ = len(v5)
        _ = fi50
    }
    fmt.Println(v19, v45)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}

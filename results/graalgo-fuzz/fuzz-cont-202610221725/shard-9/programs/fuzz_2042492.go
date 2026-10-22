package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }

func main() {
    v1 := []uint32{}
    v2 := []int{1, 0, -9223372036854775808, 433}
    var v3 bool = true
    v4 := []uint32{1}
    var v5 uint32 = 65535
    var v6 int8 = 127
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 float32 = 0.0
                _ = fi8
                _ = v9
            }
            if (1e10 < float64(3.14159)) {
                switch ((("\n\t" + ((("foo bar" + "hello") + "foo bar") + "")) + "foo bar") + "") {
                case "foo bar":
                    v5 -= v5
                case "":
                    if len(v2) > 0 {
                        v2[0] = (^(fi7))
                    }
                }
            } else {
                v6 = v6
            }
            fmt.Println(v6)
            _ = fi7
        }
    }
    v3 = ((((-2.25 * (1e10 / 3.14159)) / 1.5) < (3.14159 / 3.14159)) && v3)
    v10 := GenS1{F0: 42, F1: true, F2: 2}
    v6 = v6
    v5 -= 2
    if v3 {
        var v11 uint16 = 1
        if v3 {
            var v12 uint64 = 4294967295
            v1 = append(v1, (v5 | (v5 ^ uint32(v12))))
            v5 = v5
            _ = v12
        } else {
            var v13 int = 128
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 string = ""
                _ = fi14
                _ = v15
            }
            var v16 float64 = 1.5
            _ = v13
            _ = v16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            if false {
                v11 = 255
            }
            _ = fi17
        }
        _ = v11
    }
    v5 += 854
    if v3 {
        for fi18 := 0; fi18 < 18; fi18++ {
            if (v6 >= 42) {
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 uint = 42
                    _ = fi19
                    _ = v20
                }
            } else {
                v21 := GenS2{F0: 44, F1: 93, F2: 200}
                _ = v21
            }
            _ = fi18
        }
        v22 := GenS3{F0: (v6 - 1)}
        v5 += v5
        _ = v22
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v1) > 0 {
            v1[0] = v5
        }
    }
    _ = len(v4)
    switch 781 {
    case -59:
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(fi23, v5, fi23)
            _ = fi23
        }
    case -458:
        _ = len(v1)
    case -9223372036854775808:
        v4 = append(v4, 978)
    }
    switch "hello" {
    case "foo bar":
        for fi24 := 0; fi24 < 19; fi24++ {
            if ("a" < ("\n\t" + ("" + ("hello" + "a")))) {
                v25 := GenS2{F0: 255, F1: 0, F2: -822}
                _ = len(v2)
                var v26 int8 = v6
                _ = v25
                _ = v26
            } else {
                var v27 int16 = 42
                fmt.Println(v27, v3, v6)
                _ = v27
            }
            var v28 uint16 = 255
            _ = fi24
            _ = v28
        }
    default:
        fmt.Println(v5, v3)
    }
    var v29 string = "\n\t"
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 uint = 18446744073709551615
            if len(v4) > 0 {
                v4[0] = (v5 + v5)
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    var v33 int64 = 42
    switch v29 {
    case "hello":
        fmt.Println(v33)
    case "\n\t":
        var v34 int = -2147483648
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v33 += (-(v33))
        _ = fi35
    }
    v36 := GenS4{F0: -9223372036854775808}
    switch "foo bar" {
    case "hello":
        for fi37 := 0; fi37 < 13; fi37++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi37
        }
    case "":
        if v3 {
            var v38 float64 = 0.1
            _ = v38
        }
    case "\n\t":
        fmt.Println(v3, v5, v29)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

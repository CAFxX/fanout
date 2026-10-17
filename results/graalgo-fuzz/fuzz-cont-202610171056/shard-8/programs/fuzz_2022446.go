package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int16; F2 uint64 }
func genFn0(p0 uint64) int32 {
    fmt.Println(p0)
    var v7 int16 = -94
    _ = p0
    _ = v7
    return 494
}

func main() {
    v1 := []string{(("\n\t" + (((((("foo bar" + (("" + "hello") + ("a" + (("foo bar" + "a") + "\n\t")))) + ("hello" + (("" + "foo bar") + ("\n\t" + "\n\t")))) + ("a" + "hello")) + ("a" + "hello")) + ("foo bar" + "a")) + ("hello" + ""))) + ("foo bar" + "foo bar")), "foo bar", "\n\t"}
    v2 := make(map[uint32]string)
    var v3 float32 = 3.14159
    var v4 int32 = 42
    var v5 uint64 = 127
    var v6 int16 = 0
    if (!(false && true)) {
        v4 = v4
    }
    v6 = 541
    if (("\n\t" + "\n\t") != ("" + "\n\t")) {
        if (42 >= (v6 + v6)) {
            v8 := GenS1{F0: v5}
            switch "foo bar" {
            case "":
                switch 42 {
                case 32768:
                    v9 := GenS1{F0: 184}
                    _ = v9
                case -2147483648:
                    var v10 float64 = (1.5 + ((float64(v3) + ((-2.25 + 3.14159) / 3.14159)) - -2.25))
                    _ = v10
                }
            case "a":
                fmt.Println(v4)
            }
            var v11 float64 = 3.14159
            _ = v11
            _ = v8
        } else {
            v12 := GenS1{F0: (^(v5))}
            var v13 int64 = -1
            var v14 int32 = (v4 / ((v4) | 1))
            _ = v12
            _ = v13
            _ = v14
        }
        v6 = (v6 | v6)
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            if false {
                if true {
                    _ = len(v1)
                    var v17 uint32 = 65535
                    v18 := GenS2{F0: ""}
                    _ = v17
                    _ = v18
                } else {
                    _ = len(v2)
                    switch (^(fi16)) {
                    case 9223372036854775807:
                        v2[2] = "a"
                    case -504:
                        fmt.Println(v3, v6, v3)
                    case 127:
                        v4 += v4
                    }
                    if false {
                        var v19 uint64 = (v5 >> ((484) & 63))
                        v3 = (0.0 + 0.1)
                        v20 := GenS3{F0: 0}
                        _ = v19
                        _ = v20
                    } else {
                        var v21 float64 = float64(v3)
                        v22 := GenS4{F0: -91, F1: 0, F2: v5}
                        _ = v21
                        _ = v22
                    }
                }
                switch ("hello" + (("\n\t" + "\n\t") + (("" + ("hello" + "\n\t")) + ("\n\t" + "\n\t")))) {
                case "":
                    v4 = (v4 & (^(v4)))
                }
            } else {
                v4 -= (v4 - (v4 & v4))
            }
            v5 -= 4294967295
            v6 += v6
            _ = fi16
        }
        if (0 > 42) {
            v23 := GenS3{F0: 464}
            v1 = append(v1, "hello")
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            v4 = v4
            _ = fi24
        }
        _ = fi15
    }
    var v25 float32 = 0.1
    v25 = -2.25
    if false {
        var v26 int16 = (v6 >> ((42) & 63))
        var v27 uint16 = 1
        fmt.Println(v5)
        _ = v26
        _ = v27
    } else {
        delete(v2, 42)
        for fi28 := 0; fi28 < 12; fi28++ {
            v2[65535] = "foo bar"
            v5 = (v5 & v5)
            _ = fi28
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint8 = 1
            v6 += v6
            switch (-(fi29)) {
            case 1:
                var v31 uint64 = 4294967295
                _ = v31
            }
            _ = fi29
            _ = v30
        }
    }
    switch "hello" {
    case "a":
        v25 = (3.14159 + (0.0 / 0.1))
    case "hello":
        v4 = v4
    }
    var v32 uint = 2
    var v33 uint64 = (v5 & v5)
    if ((true && true) || false) {
        if ((false || (true && false)) && (!false)) {
            var v34 float32 = 1.5
            _ = v34
        }
        if (42 <= -1) {
            v6 = v6
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            v32 -= v32
        }
    } else {
        v2[2147483648] = "hello"
    }
    var v35 int16 = (v6 ^ (v6 - v6))
    var v36 float32 = 1.5
    var v37 string = "foo bar"
    v36 = v36
    v38 := v2[723]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

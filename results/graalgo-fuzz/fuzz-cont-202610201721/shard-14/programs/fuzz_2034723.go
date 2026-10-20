package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int16, p1 uint8) int64 {
    fmt.Println(p1, p1)
    var v5 int = -2147483648
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 6; fi7++ {
        v5 -= 298
        _ = fi7
    }
    var v8 int16 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v8
    return -662
}

func main() {
    var v1 uint8 = 0
    var v2 int = 128
    var v3 int8 = -128
    v4 := make(map[string]string)
    var v9 uint = 1
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            v1 = (v1 | 153)
            _ = fi11
        }
        v3 = (-(v3))
        _ = fi10
    }
    switch (v2 * v2) {
    case 108:
        v2 -= v2
    case 565:
        delete(v4, "foo bar")
    case -2147483648:
        fmt.Println(v3, v2)
    }
    v9 += (v9 | (v9 & (v9 + (^(v9)))))
    v2 = 2147483647
    v12 := GenS1{F0: ("a" + "foo bar")}
    if true {
        if true {
            fmt.Println(v1)
            var v13 uint = v9
            _ = v13
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    _ = len(v4)
                    var v16 uint32 = 740
                    _ = fi15
                    _ = v16
                }
                if (1 > 32767) {
                    v3 = (^(v3))
                    if (1 >= (v9 | v9)) {
                        switch 9223372036854775807 {
                        case 127:
                            v2 = (^(v2))
                        }
                        v1 = v1
                    }
                }
                if (!false) {
                    v17 := GenS2{F0: 1, F1: 511}
                    fmt.Println(v3)
                    _ = v17
                } else {
                    if false {
                        if false {
                            fmt.Println(v1)
                        }
                        if true {
                            var v18 uint = (v9 >> ((v9) & 63))
                            _ = v18
                        }
                    } else {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            fmt.Println(v1)
                            _ = fi19
                        }
                        v20 := GenS3{F0: 0, F1: "foo bar"}
                        _ = v20
                    }
                    if (!(("foo bar" + (((("hello" + "\n\t") + "a") + ("\n\t" + ("\n\t" + "foo bar"))) + "hello")) != ("hello" + "foo bar"))) {
                        v21 := GenS2{F0: 2147483647, F1: 1}
                        fmt.Println(fi14, v2, v1)
                        v22 := GenS1{F0: "\n\t"}
                        _ = v21
                        _ = v22
                    }
                    var v23 float32 = (0.1 / 1e10)
                    _ = v23
                }
                _ = fi14
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    if (!false) {
                        var v26 uint16 = 42
                        _ = v26
                    }
                    var v27 float32 = 1e10
                    v4["\n\t"] = ("a" + (("foo bar" + ("hello" + "a")) + ""))
                    _ = fi25
                    _ = v27
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        v29 := []int{9223372036854775807}
                        _ = v29
                    case "":
                        v3 = v3
                        fmt.Println(fi28, fi28)
                    default:
                        var v30 float64 = (0.0 * 1.5)
                        _ = v30
                    }
                    var v31 string = "a"
                    _ = fi28
                    _ = v31
                }
                fmt.Println(v9, v9, fi24)
                _ = fi24
            }
            v32 := v4[""]
            _ = v32
        }
        v9 = (v9 >> ((v9) & 63))
        var v33 string = "foo bar"
        _ = v33
    } else {
        v1 += v1
        fmt.Println(v1, v1, v1)
    }
    var v34 uint = v9
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint32 = 2147483648
        for fi37 := 0; fi37 < 14; fi37++ {
            v9 += (v34 + v9)
            if (-32768 == 42) {
                _ = len(v4)
            }
            v1 = v1
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    var v38 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v9
}

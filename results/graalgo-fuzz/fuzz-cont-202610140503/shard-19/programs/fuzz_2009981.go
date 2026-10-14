package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
func genFn0(p0 uint16) int64 {
    fmt.Println(p0)
    if false {
        p0 = (p0 / ((65535) | 1))
    }
    v4 := []int{-9223372036854775808, -531}
    if (1 != 502) {
        switch -799 {
        case -6:
            if (-306 >= 544) {
                p0 += (^(p0))
                for fi5 := 0; fi5 < 19; fi5++ {
                    v6 := GenS1{F0: ("hello" + "\n\t"), F1: ("hello" + "\n\t")}
                    _ = fi5
                    _ = v6
                }
            } else {
                for fi7 := 0; fi7 < 10; fi7++ {
                    fmt.Println(p0, p0)
                    p0 -= (p0 + (p0 & 814))
                    if false {
                        if true {
                            v8 := GenS1{F0: ("\n\t" + ("foo bar" + (("" + "") + "\n\t"))), F1: "hello"}
                            _ = v8
                        }
                    }
                    _ = fi7
                }
                for fi9 := 0; fi9 < 13; fi9++ {
                    p0 = 255
                    switch "a" {
                    case "hello":
                        p0 -= (p0 - 451)
                    case "foo bar":
                        for fi10 := 0; fi10 < 8; fi10++ {
                            p0 += (^(p0))
                            _ = fi10
                        }
                    }
                    p0 += (p0 | (^(p0)))
                    _ = fi9
                }
            }
        case -857:
            for fi11 := 0; fi11 < 10; fi11++ {
                switch (("hello" + ("foo bar" + "hello")) + "hello") {
                case "hello":
                    var v12 int8 = 1
                    _ = v12
                case "":
                    if len(v4) > 0 {
                        v4[0] = -2147483648
                    }
                }
                _ = fi11
            }
        case -649:
            var v13 int16 = -437
            _ = v13
        default:
            if true {
                p0 -= p0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        switch "" {
        case "hello":
            v14 := GenS1{F0: "a", F1: "\n\t"}
            _ = v14
        case "":
            var v15 int = 974
            _ = v15
        default:
            _ = len(v4)
        }
        switch "foo bar" {
        case "foo bar":
            var v16 int64 = 0
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            if len(v4) > 0 {
                v4[0] = (fi17 * (fi17 % (((fi17 * fi17)) | 1)))
            }
            _ = fi17
        }
        var v18 uint = 1
        _ = v18
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1() string {
    var v19 int16 = 1000
    var v20 uint8 = 1
    for fi21 := 0; fi21 < 15; fi21++ {
        if ((v20 >> ((65535) & 63)) < v20) {
            fmt.Println(fi21, fi21, v20)
            v22 := []int{}
            _ = v22
        } else {
            var v23 uint8 = 0
            v20 = v23
            switch -478 {
            case -1:
                if false {
                    v19 += v19
                    var v24 string = ("\n\t" + "")
                    _ = v24
                }
            case 761:
                fmt.Println(v19, v23, fi21)
            }
            _ = v23
        }
        _ = fi21
    }
    _ = v19
    _ = v20
    return (((("\n\t" + (("foo bar" + "a") + ("" + ""))) + "a") + "\n\t") + "")
}

func main() {
    var v1 int16 = 1
    var v2 string = ("" + "hello")
    v3 := make(map[string]int)
    for fi25 := 0; fi25 < 20; fi25++ {
        if false {
            var v26 string = "\n\t"
            _ = v26
        } else {
            v27 := GenS1{F0: v2, F1: v2}
            _ = v27
        }
        fmt.Println(fi25, v1, v1)
        fmt.Println(v1, v2)
        _ = fi25
    }
    var v28 int = 9223372036854775807
    v28 = (v28 | (-(v28)))
    v29 := GenS1{F0: v2, F1: "\n\t"}
    var v30 int = (v28 ^ 0)
    var v31 float64 = -2.25
    if (255 > 127) {
        fmt.Println(v2, v28)
        fmt.Println(v28, v1)
        var v32 int = (v30 & v28)
        _ = v32
    }
    v1 = (-(v1))
    _ = len(v3)
    v2 = v2
    var v33 uint8 = 42
    var v34 uint32 = 272
    var v35 uint8 = (v33 << ((227) & 63))
    var v36 bool = (v28 == (-(v28)))
    v37 := GenS2{F0: 32767, F1: -9223372036854775808, F2: 42}
    fmt.Println(v1, v35, v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}

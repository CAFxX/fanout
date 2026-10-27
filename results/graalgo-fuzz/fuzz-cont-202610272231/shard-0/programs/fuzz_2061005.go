package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
type GenS3 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int = 2147483647
    v2 := []uint32{511}
    var v3 uint8 = 127
    var v4 int = v1
    v5 := make(map[string]string)
    var v6 float32 = float32(-2.25)
    if (!(42 > -47)) {
        v7 := GenS1{F0: -1}
        var v8 uint = 255
        v6 = 0.1
        _ = v7
        _ = v8
    } else {
        switch "foo bar" {
        case "hello":
            var v9 int16 = 32767
            _ = v9
        case "":
            var v10 uint32 = 300
            _ = v10
        case "a":
            if false {
                v4 -= v4
                v6 = v6
                if false {
                    fmt.Println(v3, v3, v3)
                } else {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            var v13 string = ((("hello" + "\n\t") + "foo bar") + "\n\t")
                            v14 := GenS1{F0: 1}
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                    }
                    fmt.Println(v4, v1, v1)
                    v3 = (v3 * (v3 % ((v3) | 1)))
                }
            }
        default:
            var v15 string = "a"
            _ = v15
        }
    }
    fmt.Println(v3, v4)
    var v16 float32 = v6
    var v17 int16 = 42
    fmt.Println(v6, v17)
    for fi18 := 0; fi18 < 11; fi18++ {
        if true {
            if (!false) {
                var v19 int8 = -82
                _ = len(v2)
                v20 := GenS1{F0: (v17 * v17)}
                _ = v19
                _ = v20
            } else {
                var v21 float64 = 0.0
                v22 := GenS2{F0: (v3 + (v3 | v3)), F1: 50, F2: ("foo bar" + "\n\t")}
                _ = v21
                _ = v22
            }
            v17 += (v17 & v17)
            v4 += 1
        }
        switch "\n\t" {
        case "a":
            var v23 string = ""
            _ = v23
        case "":
            v17 = 0
        }
        _ = fi18
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            switch (^(fi25)) {
            case 1:
                var v26 uint32 = 2147483647
                _ = v26
            }
            v4 = v1
            v6 -= 1.5
            _ = fi25
        }
        v1 -= v4
        _ = fi24
    }
    v5["a"] = ("" + (("" + ("foo bar" + "hello")) + ("\n\t" + "hello")))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    for fi27 := 0; fi27 < 8; fi27++ {
        if false {
            v17 = 32767
        } else {
            var v28 float32 = (1.5 * v16)
            for fi29 := 0; fi29 < 8; fi29++ {
                v1 = v1
                var v30 uint16 = 42
                if false {
                    v4 += (v4 % (((^(fi29))) | 1))
                    v30 = v30
                }
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        v17 += -239
        v4 = v1
        _ = fi27
    }
    v3 -= 255
    for fi31 := 0; fi31 < 20; fi31++ {
        if true {
            v32 := GenS3{F0: 2147483647, F1: 42}
            delete(v5, "\n\t")
            _ = v32
        } else {
            fmt.Println(fi31)
            v6 = ((1.5 - v16) * 3.14159)
            v33 := v5["\n\t"]
            _ = v33
        }
        _ = fi31
    }
    v34 := GenS3{F0: 1, F1: 42}
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0(p0 int64) uint64 {
    p0 = (p0 ^ 9223372036854775807)
    switch "foo bar" {
    case "a":
        if true {
            var v4 int32 = 0
            if true {
                v5 := []int{}
                _ = v5
            }
            var v6 uint32 = 2147483647
            _ = v4
            _ = v6
        }
    case "":
        for fi7 := 0; fi7 < 18; fi7++ {
            p0 = (p0 % (((-(p0))) | 1))
            _ = fi7
        }
    case "hello":
        v8 := []int{-9223372036854775808, 622}
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 666
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[string]int)
    var v3 float32 = (0.1 / (((float32((3.14159 + -2.25)) - 1e10) - 0.1) * (0.1 - (1.5 * 3.14159))))
    if false {
        var v9 float64 = 1e10
        var v10 int = 180
        v3 -= (v3 * v3)
        _ = v10
        _ = v9
    }
    if false {
        var v11 float32 = v3
        switch 32767 {
        case 42:
            var v12 string = ("\n\t" + "\n\t")
            _ = v12
        case 941:
            v3 -= v3
        case -218:
            fmt.Println(v11)
        }
        if (100 < -1) {
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int16 = 990
                _ = len(v1)
                var v15 int8 = -1
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            fmt.Println(v3, v11)
            v3 -= ((v3 + 0.0) / ((v3 - v3) + v11))
            var v16 uint = 1
            _ = v16
        }
        _ = v11
    }
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v17 := GenS1{F0: 42}
    delete(v2, "foo bar")
    v3 -= float32(0.1)
    switch "\n\t" {
    case "a":
        fmt.Println(v3)
    case "\n\t":
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := v2[(((("\n\t" + "hello") + "a") + ("\n\t" + "foo bar")) + "a")]
            v20 := GenS1{F0: 32767}
            var v21 uint16 = 32767
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    case "foo bar":
        var v22 int16 = 1
        _ = v22
    default:
        delete(v2, (("hello" + (("\n\t" + "foo bar") + ("hello" + "hello"))) + "\n\t"))
    }
    v3 = 0.1
    for fi23 := 0; fi23 < 19; fi23++ {
        switch (((("" + "foo bar") + "") + ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + "\n\t"))) + "a") {
        case "":
            if (fi23 < (fi23 >> ((585) & 63))) {
                fmt.Println(fi23, v3)
                if true {
                    if false {
                        fmt.Println(v3, fi23)
                        v24 := GenS1{F0: -1}
                        for fi25 := 0; fi25 < 15; fi25++ {
                            var v26 int8 = 100
                            var v27 int16 = 1
                            switch "" {
                            case "\n\t":
                                v28 := GenS2{F0: (-(v27)), F1: 0}
                                v26 -= 0
                                var v29 int16 = (v27 + v27)
                                _ = v28
                                _ = v29
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                    } else {
                        for fi30 := 0; fi30 < 18; fi30++ {
                            var v31 uint64 = 4294967295
                            fmt.Println(v3, v31, fi30)
                            _ = fi30
                            _ = v31
                        }
                        if true {
                            fmt.Println(v3, v3)
                            if len(v1) > 0 {
                                v1[0] = ("hello" + ("" + ""))
                            }
                        } else {
                            if false {
                                var v32 string = "a"
                                var v33 int8 = 125
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                } else {
                    v3 = v3
                }
                var v34 int = (fi23 << ((2) & 63))
                _ = v34
            }
        case "foo bar":
            delete(v2, "")
        default:
            v35 := GenS3{F0: 0, F1: 128, F2: 42}
            _ = v35
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 float32 = (v3 * -2.25)
            _ = fi36
            _ = v37
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            v3 = (v3 / 0.1)
            switch fi38 {
            case 32767:
                var v39 int8 = 116
                _ = v39
            case 739:
                fmt.Println(fi23)
            case 9223372036854775807:
                var v40 uint8 = 2
                _ = v40
            }
            _ = fi38
        }
        _ = fi23
    }
    var v41 uint8 = 140
    v41 += v41
    fmt.Println(v41)
    for fi42 := 0; fi42 < 10; fi42++ {
        fmt.Println(v3)
        if false {
            var v43 int = fi42
            _ = v43
        }
        _ = fi42
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    delete(v2, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v41
}

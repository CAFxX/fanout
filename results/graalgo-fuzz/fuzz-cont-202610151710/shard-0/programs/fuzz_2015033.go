package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint16, p1 int8) int64 {
    var v7 bool = (54 <= 127)
    v8 := []int{990}
    v9 := GenS1{F0: 0, F1: 45, F2: 128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 610
}
func genFn1(p0 uint16, p1 uint, p2 uint32) int64 {
    v10 := GenS2{F0: 1, F1: 0}
    fmt.Println(p1, p0)
    p2 += p2
    switch "" {
    case "\n\t":
        v11 := GenS2{F0: 0, F1: 0}
        _ = v11
    case "a":
        var v12 uint32 = (^(p2))
        _ = v12
    }
    var v13 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 576
    var v3 string = "hello"
    var v4 int = 127
    v5 := make(map[uint32]string)
    var v6 string = ""
    v6 = "a"
    for fi14 := 0; fi14 < 14; fi14++ {
        v5[2] = ((("\n\t" + ("\n\t" + "hello")) + v3) + ("" + "foo bar"))
        for fi15 := 0; fi15 < 5; fi15++ {
            v2 = (v2 * (v2 << ((65535) & 63)))
            _ = fi15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v6 = v6
        _ = fi16
    }
    v17 := GenS2{F0: 42, F1: 125}
    if (!true) {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := []int{(fi18 * (fi18 | v4)), v4}
            v3 = ""
            switch ((v6 + (v3 + ("\n\t" + v3))) + ((v3 + "foo bar") + (v3 + "\n\t"))) {
            case "":
                _ = len(v19)
            case "a":
                var v20 uint8 = 2
                _ = v20
            case "hello":
                if ("hello" != (("hello" + "\n\t") + ((v6 + (v3 + v6)) + (v3 + ((("" + v3) + v6) + v3))))) {
                    fmt.Println(v4, v6)
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                }
            }
            _ = fi18
            _ = v19
        }
        v21 := []int{}
        fmt.Println(v3, v2, v4)
        _ = v21
    } else {
        if (-32768 < -32768) {
            var v22 float64 = (0.0 / (1e10 * 1e10))
            var v23 uint = 255
            _ = v22
            _ = v23
        }
        var v24 bool = (304 > 472)
        if v24 {
            var v25 uint64 = 42
            var v26 string = v3
            var v27 string = (("foo bar" + "hello") + "")
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v24
    }
    v2 = (v2 + (-(v2)))
    v2 += v2
    v4 = v4
    var v28 float32 = (1e10 * float32((0.1 * (-2.25 + float64(0.1)))))
    v4 -= (v4 / (((v4 >> ((1) & 63))) | 1))
    switch 720 {
    case 0:
        _ = len(v1)
    case 1:
        v29 := GenS1{F0: 1, F1: 718, F2: 42}
        _ = v29
    case -953:
        v30 := v1[(v6 + v6)]
        _ = v30
    }
    switch v4 {
    case 9223372036854775807:
        v31 := []int{}
        _ = v31
    default:
        var v32 int8 = 127
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        switch ("\n\t" + v3) {
        case "\n\t":
            fmt.Println(v2, v28)
        case "hello":
            v34 := []int{int(v2), 2147483647}
            _ = v34
        case "":
            v3 = "foo bar"
        }
        var v35 int8 = 100
        _ = fi33
        _ = v35
    }
    var v36 int = 526
    var v37 int64 = -1
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 bool = true
        v40 := GenS1{F0: 4294967295, F1: 122, F2: 128}
        switch (fi38 / ((128) | 1)) {
        case 127:
            switch v3 {
            case "hello":
                var v41 int8 = 127
                _ = v41
            case "foo bar":
                for fi42 := 0; fi42 < 15; fi42++ {
                    v28 -= v28
                    _ = fi42
                }
            case "":
                v3 = "\n\t"
            default:
                v43 := v5[847]
                _ = v43
            }
        case -2147483648:
            fmt.Println(v39)
        case 308:
            v6 = ("a" + "")
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    var v44 int = (-(v36))
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}

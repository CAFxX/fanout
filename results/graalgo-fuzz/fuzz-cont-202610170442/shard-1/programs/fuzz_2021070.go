package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 string) int16 {
    v6 := []int{2147483647, 331, 32768}
    var v7 uint32 = 4294967295
    _ = len(v6)
    var v8 int = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int8, p1 uint64, p2 uint16) uint {
    p0 = -10
    var v9 int = 426
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 943
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v4 := []string{"\n\t", "hello"}
    var v5 float32 = (1.5 - 0.0)
    v4 = append(v4, "hello")
    v5 += (v5 / v5)
    var v10 int = -824
    v1 = v1
    if true {
        switch 4294967295 {
        case 42:
            if ((-202 != 1) || true) {
                v11 := GenS1{F0: 352}
                if (!(v1 > v1)) {
                    _ = len(v4)
                    switch (v10 ^ v10) {
                    case 0:
                        v4 = append(v4, (("hello" + "hello") + (("foo bar" + "\n\t") + "")))
                    }
                }
                _ = v11
            } else {
                v12 := GenS1{F0: 65535}
                v1 = v1
                _ = v12
            }
        }
        switch ("a" + "a") {
        case "hello":
            if (65535 > genFn1(1, (v1 - v1), 74)) {
                v10 -= v10
                fmt.Println(v5, v5, v5)
            }
        case "foo bar":
            v5 = ((float32(1.5) - float32(1e10)) * 1.5)
        case "\n\t":
            switch ((("hello" + "hello") + "hello") + "\n\t") {
            case "foo bar":
                v10 -= 158
            case "hello":
                var v13 int64 = 9223372036854775807
                _ = v13
            case "\n\t":
                var v14 int8 = 115
                _ = v14
            }
        default:
            switch 128 {
            case -73:
                if (true && (!false)) {
                    v2["a"] = ("hello" + "\n\t")
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 uint64 = 4294967295
                        v17 := GenS1{F0: genFn1(1, (v1 & v1), 255)}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
            default:
                fmt.Println(v1)
            }
        }
    } else {
        var v18 uint8 = 2
        _ = v18
    }
    if true {
        var v19 int16 = -32768
        for fi20 := 0; fi20 < 20; fi20++ {
            if (!false) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := GenS1{F0: genFn1(-1, (v1 * 0), 147)}
                    _ = fi21
                    _ = v22
                }
            }
            v19 = 1
            _ = fi20
        }
        _ = v19
    } else {
        fmt.Println(v5, v1)
        v1 = v1
        for fi23 := 0; fi23 < 16; fi23++ {
            _ = len(v4)
            _ = fi23
        }
    }
    var v24 float32 = 0.0
    fmt.Println(v10)
    for fi25 := 0; fi25 < 1; fi25++ {
        v24 += v5
        var v26 uint32 = 4294967295
        var v27 int32 = -2147483648
        _ = fi25
        _ = v26
        _ = v27
    }
    if (((false && (!true)) && ((("hello" + "a") + "") <= "foo bar")) || (!false)) {
        var v28 uint64 = (v1 * 2)
        if len(v4) > 0 {
            v4[0] = ("foo bar" + "")
        }
        _ = v28
    }
    fmt.Println(v24, v24)
    if (339 != 1) {
        var v29 string = ("a" + "")
        _ = v29
    } else {
        var v30 int64 = -429
        _ = v30
    }
    v24 += -2.25
    v1 = (v1 >> ((2) & 63))
    v31 := GenS1{F0: 1}
    if (42 >= v1) {
        fmt.Println(v5, v1, v1)
        v1 += (v1 - v1)
        v24 -= (-2.25 / v24)
    } else {
        var v32 string = "hello"
        v5 = float32(float64(v5))
        var v33 uint64 = (v1 * v1)
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        v1 = 4294967295
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}

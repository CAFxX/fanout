package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int64 }
func genFn0() uint8 {
    var v6 uint32 = 42
    v7 := []int{-689, 32768}
    if len(v7) > 0 {
        v7[0] = 32767
    }
    if (("a" + "") > ("hello" + "")) {
        v6 = (v6 ^ (v6 & (v6 | 4294967295)))
        switch "hello" {
        case "a":
            if (-2147483648 < 2147483647) {
                var v8 int8 = 127
                v9 := GenS1{F0: v8, F1: 358, F2: (v8 & (v8 & (v8 % (((v8 | (v8 % (((v8 << ((42) & 63))) | 1)))) | 1))))}
                _ = v8
                _ = v9
            }
        case "":
            v10 := GenS1{F0: 0, F1: 42, F2: -128}
            _ = v10
        }
    } else {
        var v11 float32 = 1e10
        switch "foo bar" {
        case "\n\t":
            var v12 uint16 = 0
            _ = v12
        case "foo bar":
            switch (("" + ("hello" + (("\n\t" + (("" + ("" + ("hello" + "a"))) + ("\n\t" + "\n\t"))) + ""))) + "foo bar") {
            case "\n\t":
                v13 := GenS1{F0: -110, F1: 65535, F2: 100}
                _ = v13
            case "":
                var v14 uint8 = 99
                _ = v14
            case "foo bar":
                v15 := GenS1{F0: 42, F1: 0, F2: 1}
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 3; fi16++ {
                v11 += (float32(0.0) / float32(((3.14159 - (0.0 * 1.5)) + 1.5)))
                var v17 int32 = 42
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 uint64 = 1
                    if len(v7) > 0 {
                        v7[0] = (fi18 | (fi16 % ((185) | 1)))
                    }
                    var v20 int64 = 9223372036854775807
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
        }
        _ = v11
    }
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint16, p1 uint16) int8 {
    p1 -= (p0 - p1)
    p0 = 2
    var v21 uint = 65535
    if true {
        fmt.Println(v21, p1)
        var v22 uint8 = 128
        _ = v22
    } else {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int = (fi23 - fi23)
            _ = fi23
            _ = v24
        }
        p0 -= (p0 % ((641) | 1))
        v25 := GenS1{F0: 49, F1: 4294967295, F2: -128}
        _ = v25
    }
    var v26 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v21
    _ = v26
    return -36
}

func main() {
    v1 := []uint32{868, 42}
    var v2 uint16 = 32767
    v3 := make(map[string]uint32)
    var v4 int32 = 42
    v5 := []int{548}
    var v27 bool = false
    v28 := GenS1{F0: 42, F1: 139, F2: genFn1(v2, 416)}
    var v29 int64 = 9223372036854775807
    v2 += v2
    switch 32767 {
    case -9223372036854775808:
        _ = len(v3)
    case 32767:
        var v30 uint64 = 135
        _ = v30
    case 2147483647:
        delete(v3, ("" + ("\n\t" + ("hello" + "\n\t"))))
    default:
        fmt.Println(v27, v29, v2)
    }
    v31 := GenS1{F0: -49, F1: 1, F2: genFn1(374, v2)}
    var v32 int8 = genFn1(v2, v2)
    v33 := GenS1{F0: (v32 | v32), F1: 0, F2: v32}
    switch ("a" + "foo bar") {
    case "foo bar":
        v4 -= (v4 / ((v4) | 1))
    case "":
        var v34 uint8 = 179
        _ = v34
    }
    fmt.Println(v29, v4)
    switch "hello" {
    case "hello":
        var v35 float64 = 3.14159
        _ = v35
    }
    v29 += (v29 / ((v29) | 1))
    v27 = ((v32 << ((18446744073709551615) & 63)) == (v32 - v32))
    v2 += v2
    if true {
        var v36 float32 = (0.0 / float32(1.5))
        v5 = append(v5, -988)
        _ = v36
    }
    if true {
        fmt.Println(v2, v32)
        v2 = (v2 % ((v2) | 1))
        v37 := GenS2{F0: (v2 >> ((65535) & 63)), F1: (v29 * 42), F2: (v29 << ((463) & 63))}
        _ = v37
    } else {
        v5 = append(v5, 127)
        v38 := GenS3{F0: 796}
        if false {
            if ((v2 % (((v2 + (v2 | v2))) | 1)) > 255) {
                if ((!(!false)) || true) {
                    v4 += 589
                } else {
                    var v39 int8 = genFn1(1, v2)
                    if (!false) {
                        fmt.Println(v27)
                    } else {
                        v40 := GenS3{F0: v29}
                        var v41 int16 = 1000
                        _ = v40
                        _ = v41
                    }
                    _ = v39
                }
                for fi42 := 0; fi42 < 7; fi42++ {
                    var v43 uint8 = genFn0()
                    _ = fi42
                    _ = v43
                }
            }
            v1 = append(v1, 686)
        }
        _ = v38
    }
    if true {
        delete(v3, "foo bar")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}

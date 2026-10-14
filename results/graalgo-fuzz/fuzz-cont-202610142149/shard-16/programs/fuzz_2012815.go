package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
func genFn0() uint8 {
    var v4 int = -356
    v5 := []int{(v4 / ((v4) | 1)), (v4 - -1)}
    v6 := GenS1{F0: 2}
    var v7 uint16 = 42
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int8) uint8 {
    v8 := []int{4294967295, -929, -963, 128}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return genFn0()
}

func main() {
    var v1 int32 = 1
    var v2 int8 = -13
    var v3 float32 = 1.5
    if false {
        var v9 int16 = -622
        if (!(840 >= 917)) {
            switch "" {
            case "hello":
                fmt.Println(v3, v3)
            case "a":
                switch "" {
                case "hello":
                    v1 = (-(v1))
                case "foo bar":
                    var v10 float32 = ((v3 - v3) / (((v3 - 1e10) + v3) * v3))
                    _ = v10
                }
            case "\n\t":
                var v11 uint = 42
                _ = v11
            default:
                var v12 int16 = 1000
                _ = v12
            }
            var v13 uint32 = 4294967295
            var v14 uint32 = 42
            _ = v13
            _ = v14
        }
        _ = v9
    }
    if (true && ((!false) && false)) {
        v2 -= v2
        if true {
            if (v2 != (-(v2))) {
                var v15 int8 = 0
                switch -368 {
                case 42:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v17 := []int{1, (fi16 | fi16)}
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v15
            }
        } else {
            fmt.Println(v3, v1)
            if (!false) {
                var v18 int64 = 779
                v19 := []int{1, 128, -807}
                fmt.Println(v3, v2)
                _ = v18
                _ = v19
            } else {
                v20 := []int{4294967295, -9223372036854775808, -9223372036854775808, 9223372036854775807}
                switch 1 {
                case 745:
                    switch 128 {
                    case 128:
                        v1 += (-(v1))
                    default:
                        for fi21 := 0; fi21 < 16; fi21++ {
                            fmt.Println(v1, v3, v2)
                            fmt.Println(fi21, v3)
                            _ = fi21
                        }
                    }
                case 32767:
                    v2 = v2
                case 1:
                    var v22 uint8 = 128
                    _ = v22
                }
                for fi23 := 0; fi23 < 1; fi23++ {
                    fmt.Println(v1, v3)
                    v24 := GenS2{F0: -32768}
                    v20 = append(v20, fi23)
                    _ = fi23
                    _ = v24
                }
                _ = v20
            }
        }
        v2 = v2
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS3{F0: 133, F1: "a", F2: -540}
        v27 := []int{fi25}
        v3 = v3
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint8 = 1
    v2 += -69
    if ((("a" + "") + ("" + ("a" + "a"))) < "a") {
        v28 = (v28 | 42)
        if false {
            v28 = genFn0()
            v1 -= v1
        }
        v29 := []int{-9223372036854775808, 9223372036854775807, 42}
        _ = v29
    }
    v3 -= float32(0.0)
    var v30 uint64 = 2
    if (!false) {
        v31 := []int{}
        switch 1 {
        case 2147483647:
            var v32 uint64 = v30
            _ = v32
        case 1:
            switch (("" + "a") + (("foo bar" + "\n\t") + "\n\t")) {
            case "":
                var v33 float64 = -2.25
                _ = v33
            case "hello":
                var v34 float64 = 0.1
                _ = v34
            }
        }
        _ = v31
    }
    var v35 uint8 = (v28 >> ((18446744073709551615) & 63))
    if (v35 == (v28 >> ((342) & 63))) {
        v28 -= genFn1((v2 * -121))
    } else {
        if true {
            if false {
                fmt.Println(v1)
                var v36 int16 = -32768
                _ = v36
            } else {
                var v37 int16 = -738
                v38 := GenS2{F0: v37}
                switch (((("" + "foo bar") + "a") + ("\n\t" + ("hello" + (("a" + "\n\t") + "a")))) + ("\n\t" + "")) {
                case "hello":
                    v2 -= (v2 & 1)
                }
                _ = v37
                _ = v38
            }
            switch "hello" {
            case "":
                var v39 int64 = -327
                _ = v39
            case "hello":
                for fi40 := 0; fi40 < 2; fi40++ {
                    v41 := []int{}
                    v30 = (v30 - v30)
                    fmt.Println(v30, v28, v1)
                    _ = fi40
                    _ = v41
                }
            case "foo bar":
                v42 := []int{1, 42, 9223372036854775807}
                _ = v42
            }
            v43 := GenS1{F0: 42}
            _ = v43
        }
    }
    v30 += (v30 * 1)
    v44 := []int{32767}
    for fi45 := 0; fi45 < 2; fi45++ {
        v46 := GenS1{F0: 666}
        _ = fi45
        _ = v46
    }
    if len(v44) > 0 {
        v44[0] = -733
    }
    fmt.Println(v2)
    if len(v44) > 0 {
        v44[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v35
    _ = v44
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint32; F1 string; F2 uint32 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []int{127, -439, 1, 0}
    v2 := make(map[string]uint32)
    v3 := []string{("\n\t" + "foo bar"), "", "hello"}
    v4 := []int{-2147483648}
    if true {
        var v5 float64 = ((0.0 / 3.14159) - (1e10 - 0.1))
        var v6 float32 = (((1.5 - 1e10) - 1e10) + 1e10)
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                switch ("foo bar" + ("foo bar" + "a")) {
                case "a":
                    var v9 string = "\n\t"
                    _ = v9
                case "foo bar":
                    _ = len(v2)
                case "":
                    var v10 int32 = 0
                    _ = v10
                }
                if (0 == 538) {
                    v6 += v6
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v5 = -2.25
                }
                _ = fi8
            }
            if true {
                _ = len(v2)
                v2["a"] = 1
            } else {
                var v11 bool = false
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 int16 = 0
                    fmt.Println(v11)
                    v13 += (v13 + (v13 & v13))
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v11, v6)
                    _ = fi14
                }
                _ = v11
            }
            _ = fi7
        }
        _ = v5
        _ = v6
    } else {
        var v15 float64 = 3.14159
        v15 += float64(1e10)
        var v16 int16 = 42
        _ = v15
        _ = v16
    }
    if false {
        var v17 int32 = 42
        _ = v17
    } else {
        switch 0 {
        case 9223372036854775807:
            var v18 bool = (("hello" >= "foo bar") && false)
            _ = v18
        default:
            var v19 uint = 456
            _ = v19
        }
        if ((4294967295 < 4294967295) || (("\n\t" + (("" + (("\n\t" + ((("\n\t" + "a") + (("\n\t" + "a") + ((("\n\t" + (("foo bar" + ("hello" + "foo bar")) + "\n\t")) + "\n\t") + (("foo bar" + "foo bar") + "")))) + "\n\t")) + "a")) + "foo bar")) != "hello")) {
            v20 := GenS1{F0: 32768, F1: 32767}
            _ = v20
        }
        var v21 uint32 = 1
        _ = v21
    }
    switch 1 {
    case 4294967295:
        var v22 uint16 = 2
        _ = v22
    default:
        v23 := GenS2{F0: 0, F1: 0, F2: 319}
        _ = v23
    }
    if len(v4) > 0 {
        v4[0] = -1
    }
    v24 := GenS3{F0: 1, F1: ("\n\t" + "\n\t"), F2: 2147483648}
    v25 := GenS3{F0: 0, F1: "", F2: 638}
    var v26 int = 9223372036854775807
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 int32 = -1
        _ = fi27
        _ = v28
    }
    switch (v26 | (v26 + v26)) {
    case 32768:
        var v29 int = (-(v26))
        _ = v29
    case 9223372036854775807:
        v26 += (v26 + v26)
    case -370:
        v26 += (v26 >> ((65535) & 63))
    }
    v26 = v26
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v26)
    fmt.Println(v26)
    v30 := GenS1{F0: v26, F1: 32767}
    v31 := GenS4{F0: 1}
    switch ("\n\t" + ("a" + ("foo bar" + ("\n\t" + "\n\t")))) {
    case "":
        var v32 uint16 = 65535
        _ = v32
    case "hello":
        v26 = (-(v26))
    case "a":
        _ = len(v2)
    }
    fmt.Println(v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}

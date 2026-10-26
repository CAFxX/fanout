package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 bool) uint {
    if (804 >= -508) {
        var v7 int16 = -32768
        v8 := []int{}
        var v9 uint8 = 2
        _ = v7
        _ = v8
        _ = v9
    }
    p0 = (422 == 0)
    var v10 uint = 65535
    var v11 uint = v10
    _ = p0
    _ = v10
    _ = v11
    return (v10 % (((v10 - v10)) | 1))
}
func genFn1(p0 int8, p1 bool) uint64 {
    v12 := []int{-827, 32767, 0, 4294967295}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    _ = p0
    _ = p1
    _ = v12
    return 4294967295
}

func main() {
    v1 := []uint32{2, 2147483648}
    v2 := []int{1, 786, 607, 622}
    var v3 int = 1
    var v4 float32 = -2.25
    v5 := make(map[string]string)
    var v6 float64 = 1.5
    if (genFn1(-42, false) == 4294967295) {
        var v13 int8 = 100
        var v14 uint16 = 59
        if (65535 <= 2) {
            for fi15 := 0; fi15 < 16; fi15++ {
                v4 -= v4
                _ = fi15
            }
        }
        _ = v13
        _ = v14
    }
    var v16 uint = 2
    if true {
        var v17 float32 = v4
        for fi18 := 0; fi18 < 7; fi18++ {
            switch 0 {
            case 797:
                switch v3 {
                case 127:
                    fmt.Println(v4, v17)
                case 32767:
                    v17 = v17
                default:
                    v6 = (v6 / 1.0)
                }
            case -371:
                fmt.Println(fi18, fi18)
            case -238:
                var v19 float32 = (((0.1 + v17) + ((0.0 * -2.25) + 3.14159)) - float32((((3.14159 * 3.14159) / 0.1) + float64(v17))))
                _ = v19
            }
            v20 := v5["a"]
            v21 := GenS1{F0: -765, F1: 18446744073709551615}
            _ = fi18
            _ = v20
            _ = v21
        }
        _ = v17
    } else {
        fmt.Println(v3)
        v22 := GenS1{F0: 442, F1: 2}
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 float64 = 0.1
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v1 = append(v1, 1)
    var v25 float64 = 1.5
    var v26 float64 = -2.25
    v1 = append(v1, 65535)
    switch ("" + ("foo bar" + (("hello" + "foo bar") + (("foo bar" + ("" + "\n\t")) + (("a" + "foo bar") + (("hello" + "foo bar") + "hello")))))) {
    case "foo bar":
        delete(v5, "a")
    case "a":
        switch (v3 & v3) {
        case 111:
            v25 = (float64((float32(v25) * v4)) + v6)
        case 831:
            for fi27 := 0; fi27 < 17; fi27++ {
                switch ("\n\t" + ((("foo bar" + "") + "hello") + ("a" + "a"))) {
                case "\n\t":
                    if true {
                        v28 := GenS1{F0: -308, F1: 4294967295}
                        if (!(9223372036854775807 < -9223372036854775808)) {
                            v1 = append(v1, 4294967295)
                        }
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v5["\n\t"] = (("\n\t" + "hello") + (("a" + "foo bar") + "hello"))
                            var v30 int = (v3 >> (((^(v16))) & 63))
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                case "a":
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v32 := GenS2{F0: 255, F1: 42}
                        _ = len(v5)
                        fmt.Println(v6, v26)
                        _ = fi31
                        _ = v32
                    }
                case "hello":
                    var v33 uint16 = 65535
                    _ = v33
                }
                var v34 uint32 = 1
                _ = fi27
                _ = v34
            }
        case -838:
            var v35 int = v3
            _ = v35
        }
    case "":
        var v36 uint32 = 27
        _ = v36
    default:
        v5["hello"] = "a"
    }
    fmt.Println(v25, v26)
    v3 = v3
    v37 := GenS3{F0: "a", F1: 76, F2: 40}
    if (1 > 128) {
        var v38 int64 = 0
        var v39 int16 = 1
        v40 := GenS2{F0: 972, F1: 339}
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 string; F2 uint8 }

func main() {
    v1 := []uint32{2, 612, 555, 151}
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: 547, F1: 991}
            var v7 int16 = 32767
            switch (("foo bar" + (("foo bar" + "") + "\n\t")) + (((((("a" + "hello") + (("foo bar" + "") + (("\n\t" + "\n\t") + ("\n\t" + ("foo bar" + ""))))) + ((("foo bar" + ("hello" + "hello")) + "\n\t") + (("foo bar" + "") + ("" + ("" + (("" + ("" + "")) + "\n\t")))))) + "\n\t") + ("\n\t" + "hello")) + "foo bar")) {
            case "foo bar":
                var v8 uint32 = 4294967295
                _ = v8
            case "hello":
                if false {
                    switch fi4 {
                    case -1:
                        v9 := GenS2{F0: 65535}
                        _ = v9
                    case 32767:
                        v10 := GenS1{F0: 1, F1: -1}
                        _ = v10
                    case 128:
                        v7 -= v7
                    }
                    fmt.Println(v7, v7, fi5)
                    var v11 float64 = (float64(0.1) / (3.14159 + ((0.0 + ((3.14159 * (3.14159 * float64(0.0))) - 1e10)) - (0.1 * float64((1.5 + (((-2.25 / (1e10 - ((float32(1.5) / (0.1 + 0.1)) - ((-2.25 + ((0.0 * 0.1) * 0.0)) + 3.14159)))) * (0.0 - 3.14159)) / float32((1.5 / float64(((float32(3.14159) / 1e10) - -2.25)))))))))))
                    _ = v11
                }
                v12 := GenS2{F0: 0}
                _ = v12
            case "\n\t":
                v13 := GenS3{F0: 255, F1: 42}
                v7 = -59
                v3[2] = "foo bar"
                _ = v13
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        var v14 int = -60
        _ = fi4
        _ = v14
    }
    switch -66 {
    case -941:
        if true {
            if true {
                var v15 int = 4294967295
                v16 := GenS1{F0: 2, F1: -4}
                for fi17 := 0; fi17 < 20; fi17++ {
                    switch "" {
                    case "a":
                        v15 -= fi17
                    case "foo bar":
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if (-128 == 0) {
                                v15 += -1
                            }
                            _ = fi18
                        }
                    }
                    var v19 uint64 = 42
                    _ = len(v3)
                    _ = fi17
                    _ = v19
                }
                _ = v15
                _ = v16
            }
            var v20 int64 = 42
            _ = v20
        }
    case -511:
        v21 := GenS4{F0: 2}
        _ = v21
    }
    var v22 int16 = 1
    v22 = 42
    v2["hello"] = "\n\t"
    var v23 int64 = -500
    if (42 >= 4294967295) {
        var v24 float64 = (1e10 + 1e10)
        for fi25 := 0; fi25 < 5; fi25++ {
            if ((fi25 % (((^(fi25))) | 1)) <= (fi25 - (fi25 % (((fi25 ^ fi25)) | 1)))) {
                _ = len(v3)
                v22 = (v22 - v22)
            }
            fmt.Println(fi25, v22, fi25)
            _ = fi25
        }
        _ = v24
    }
    v22 += 42
    var v26 string = ""
    v26 = (v26 + v26)
    if true {
        if true {
            if true {
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 uint16 = 42
                    _ = fi27
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v23 = (v23 >> ((1) & 63))
                fmt.Println(v26, v23, v26)
                var v30 float64 = 3.14159
                _ = fi29
                _ = v30
            }
            v22 += -1
        } else {
            v31 := GenS5{F0: 4294967295, F1: v26, F2: 128}
            fmt.Println(v22, v23)
            _ = v31
        }
    } else {
        var v32 int64 = 208
        v32 += (^(v32))
        var v33 uint16 = 42
        _ = v32
        _ = v33
    }
    var v34 uint32 = 691
    v34 = (v34 & (v34 % (((^(v34))) | 1)))
    var v35 int8 = -118
    v3[65535] = "foo bar"
    var v36 uint16 = 1
    v37 := GenS1{F0: 1, F1: 9223372036854775807}
    v22 = v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}

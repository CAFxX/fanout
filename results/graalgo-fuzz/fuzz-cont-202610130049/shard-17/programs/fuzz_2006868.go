package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 bool }
func genFn0() uint32 {
    v6 := GenS1{F0: 867, F1: -313, F2: -1}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint64 = 42
        v9 := GenS2{F0: 18446744073709551615, F1: (-2.25 <= 1e10)}
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        var v10 bool = ((!(float64(float32(3.14159)) <= 1.5)) && (128 > 144))
        var v11 int64 = -1
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 221, F1: 4294967295, F2: -101}
        v13 := []int{4294967295, 128}
        switch ("foo bar" + "a") {
        case "foo bar":
            var v14 bool = false
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 bool = (255 < 247)
                v17 := GenS2{F0: 65535, F1: v16}
                _ = fi15
                _ = v16
                _ = v17
            }
        case "":
            switch ((("\n\t" + "") + ("hello" + "a")) + "foo bar") {
            case "hello":
                var v18 int8 = 127
                _ = v18
            case "\n\t":
                if true {
                    if (127 == 108) {
                        if len(v13) > 0 {
                            v13[0] = -928
                        }
                        if true {
                            var v19 float32 = ((0.0 * ((float32((((1e10 - 0.0) - 1e10) * (-2.25 - float64((((0.0 / ((0.1 / float32((((-2.25 / 1.5) - float64((3.14159 + -2.25))) - -2.25))) - -2.25)) - (0.1 + (-2.25 + 3.14159))) - -2.25))))) / 3.14159) / (1e10 / 3.14159))) * (1.5 * float32(-2.25)))
                            var v20 uint64 = 2
                            _ = len(v13)
                            _ = v19
                            _ = v20
                        }
                        v13 = append(v13, -426)
                    } else {
                        v13 = append(v13, -446)
                        var v21 int16 = 42
                        if (!false) {
                            var v22 uint = 65535
                            _ = v22
                        } else {
                            var v23 float64 = -2.25
                            var v24 int64 = 975
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    }
                    for fi25 := 0; fi25 < 10; fi25++ {
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v13 = append(v13, fi26)
                            if len(v13) > 0 {
                                v13[0] = fi25
                            }
                            var v27 int16 = -1
                            switch -9223372036854775808 {
                            case -1:
                                fmt.Println(fi25, fi26)
                            }
                            _ = fi26
                            _ = v27
                        }
                        var v28 float32 = 1.5
                        _ = fi25
                        _ = v28
                    }
                }
            }
        default:
            var v29 uint = 469
            _ = v29
        }
        _ = v12
        _ = v13
    }
    if (!true) {
        if false {
            var v30 uint32 = 9
            _ = v30
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 float32 = 0.0
            _ = fi31
            _ = v32
        }
    } else {
        var v33 uint = 255
        _ = v33
    }
    var v34 int32 = 1
    _ = v34
    _ = v6
    return 4294967295
}

func main() {
    v1 := []string{"hello", "\n\t"}
    var v2 int16 = -1
    v3 := make(map[string]uint32)
    var v4 uint = 764
    var v5 float64 = float64(-2.25)
    v35 := GenS2{F0: 42, F1: true}
    v2 = (v2 | (v2 << (((v4 - v4)) & 63)))
    var v36 int64 = 1
    for fi37 := 0; fi37 < 16; fi37++ {
        _ = len(v3)
        _ = fi37
    }
    switch (("foo bar" + "a") + "foo bar") {
    case "foo bar":
        var v38 uint = (^(v4))
        _ = v38
    }
    v2 = v2
    v4 = (v4 ^ v4)
    var v39 uint16 = 211
    v36 = v36
    fmt.Println(v5, v5)
    v5 = (v5 * (v5 - (float64(-2.25) + (0.1 + float64(0.1)))))
    if false {
        var v40 uint64 = 0
        _ = v40
    } else {
        delete(v3, "\n\t")
    }
    v36 += (^(v36))
    v41 := GenS1{F0: -2147483648, F1: -184, F2: 42}
    _ = len(v1)
    v42 := GenS2{F0: v4, F1: true}
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}

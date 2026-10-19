package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 bool = (!true)
    var v2 bool = false
    v3 := make(map[uint32]string)
    var v4 uint16 = 514
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: -1, F1: uint32(v4), F2: 274}
        if (348 <= 0) {
            fmt.Println(v4)
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 int8 = -73
                var v9 bool = v1
                var v10 int16 = -1
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 float32 = 1e10
                var v13 float32 = 1e10
                v14 := GenS2{F0: ("hello" >= "foo bar"), F1: 0}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        } else {
            v4 = 65535
            for fi15 := 0; fi15 < 12; fi15++ {
                fmt.Println(fi5)
                _ = fi15
            }
            v2 = v2
        }
        if (1.5 != ((1e10 + (((float64(0.0) / 1e10) / 3.14159) * 0.1)) * 1.5)) {
            switch "\n\t" {
            case "":
                switch fi5 {
                case 32767:
                    v2 = (float64((-2.25 - 0.0)) != 1.5)
                }
            case "foo bar":
                v1 = ((((("" + "\n\t") + "\n\t") + "hello") >= "a") || (!(60 >= 147)))
            }
            v1 = v2
        }
        _ = fi5
        _ = v6
    }
    v2 = v1
    v16 := GenS2{F0: (27 <= 42), F1: -849}
    v17 := GenS2{F0: true, F1: 42}
    fmt.Println(v2)
    var v18 int8 = 100
    v19 := []int{384, -2147483648}
    _ = len(v3)
    var v20 float32 = ((0.0 / 1e10) - 1.5)
    if len(v19) > 0 {
        v19[0] = 4294967295
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int = -2147483648
        switch "\n\t" {
        case "hello":
            v23 := GenS1{F0: v18, F1: 4294967295, F2: 2}
            _ = v23
        case "":
            var v24 bool = (0 >= (v4 >> ((640) & 63)))
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    var v25 uint64 = 4294967295
    if (!false) {
        var v26 int = 4294967295
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int64 = 42
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                fmt.Println(v1)
            }
            var v29 float32 = 3.14159
            _ = fi27
            _ = v28
            _ = v29
        }
        if true {
            if (2147483647 > (v26 - (-(v26)))) {
                v20 = ((3.14159 / (float32(-2.25) / 3.14159)) + (3.14159 * 1e10))
            }
        }
        _ = v26
    }
    v4 -= 32767
    v30 := GenS2{F0: v2, F1: 1}
    for fi31 := 0; fi31 < 11; fi31++ {
        switch fi31 {
        case 9223372036854775807:
            var v32 int32 = 725
            _ = v32
        case -623:
            if ((uint(v25) > 4294967295) || false) {
                if v2 {
                    if v2 {
                        v33 := GenS2{F0: ((((0 > 65) || v2) || v2) && v1), F1: int64(v4)}
                        _ = v33
                    } else {
                        var v34 int64 = int64(v4)
                        for fi35 := 0; fi35 < 1; fi35++ {
                            v20 = (0.0 * ((((-2.25 - 3.14159) - ((float32(3.14159) - ((((v20 * float32(float64(-2.25))) * (v20 * (1e10 + v20))) - v20) * 3.14159)) + 0.1)) - v20) - (0.1 + float32(0.0))))
                            v19 = append(v19, (fi35 - fi31))
                            _ = fi35
                        }
                        var v36 int16 = 375
                        _ = v34
                        _ = v36
                    }
                    v18 -= -1
                }
            } else {
                if v1 {
                    v37 := GenS2{F0: (true && (false || v2)), F1: 9223372036854775807}
                    var v38 int8 = v18
                    _ = v37
                    _ = v38
                }
            }
        case 2147483647:
            var v39 uint64 = (v25 - v25)
            _ = v39
        default:
            if ((287 < 65535) && v1) {
                for fi40 := 0; fi40 < 2; fi40++ {
                    var v41 float64 = 0.1
                    var v42 bool = (!(((((0.0 + v20) * v20) - v20) >= (((((v20 - (((v20 - 1e10) - 0.0) + ((v20 / ((((((v20 - ((v20 - (v20 / v20)) - v20)) + ((v20 * v20) - 0.1)) - ((((float32((v41 * float64(v20))) / v20) * float32(v41)) / (v20 - v20)) / 1e10)) * v20) / 1.0) * 1e10)) * 1e10))) / (v20 - v20)) * (float32(((v41 * v41) - (0.1 / 1e10))) / v20)) + (v20 / 1e10)) - v20)) && (!v2)))
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}

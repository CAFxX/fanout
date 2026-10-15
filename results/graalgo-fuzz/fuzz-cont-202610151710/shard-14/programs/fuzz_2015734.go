package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
func genFn0() uint {
    if false {
        if true {
            v4 := []int{-9223372036854775808}
            _ = v4
        }
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := []int{(fi5 + fi5)}
            for fi7 := 0; fi7 < 17; fi7++ {
                switch (fi5 - -9223372036854775808) {
                case -2147483648:
                    var v8 uint = 873
                    _ = v8
                default:
                    var v9 uint32 = 42
                    _ = v9
                }
                _ = fi7
            }
            _ = fi5
            _ = v6
        }
    } else {
        switch 0 {
        case -9223372036854775808:
            var v10 uint32 = 42
            _ = v10
        }
        v11 := []int{9223372036854775807, 127, -511}
        var v12 int32 = -2147483648
        _ = v11
        _ = v12
    }
    var v13 int32 = 884
    var v14 uint8 = 42
    for fi15 := 0; fi15 < 5; fi15++ {
        switch "" {
        case "a":
            var v16 int = fi15
            _ = v16
        case "\n\t":
            switch fi15 {
            case 797:
                v17 := []int{}
                _ = v17
            case 4294967295:
                var v18 uint16 = uint16(v14)
                _ = v18
            }
        case "":
            v14 = v14
        default:
            var v19 uint32 = 1
            _ = v19
        }
        if true {
            v20 := GenS1{F0: -128}
            _ = v20
        }
        _ = fi15
    }
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    var v1 float64 = float64(-2.25)
    var v2 string = ""
    var v3 int8 = 100
    v1 += float64(float32(v1))
    if (false && (1 > 86)) {
        v21 := GenS1{F0: 1}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint32 = 65535
                switch (-(fi23)) {
                case -9223372036854775808:
                    fmt.Println(fi22, v1, v1)
                }
                var v25 float64 = v1
                if (0 > -1) {
                    var v26 int16 = -744
                    fmt.Println(v24)
                    if true {
                        v27 := GenS2{F0: "hello"}
                        _ = v27
                    }
                    _ = v26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        v28 := []int{-2147483648, 32767, 837}
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := []int{42}
        var v31 int = fi29
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 float64 = v1
    v33 := []int{515, -303, -444, 715}
    v34 := GenS1{F0: (v3 | v3)}
    var v35 float64 = (float64(1e10) / ((0.1 * -2.25) + (0.1 * v32)))
    var v36 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}

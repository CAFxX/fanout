package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
func genFn0(p0 int8, p1 uint8) uint64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 65535
        var v9 string = ""
        _ = fi7
        _ = v8
        _ = v9
    }
    if (!(2147483647 != -2147483648)) {
        v10 := GenS1{F0: 730, F1: -697, F2: 42}
        v11 := []int{}
        _ = v10
        _ = v11
    }
    v12 := []int{596, 2147483647}
    if ((p0 >> ((0) & 63)) == p0) {
        var v13 uint = 675
        fmt.Println(p0)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 float64 = 1.5
    var v5 uint8 = 1
    v6 := make(map[string]string)
    v5 -= (v5 | (^(v5)))
    v14 := []int{-775, -2147483648}
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS1{F0: 9223372036854775807, F1: 1, F2: 954}
        var v17 string = v2
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 += (v5 >> ((100) & 63))
    var v18 float64 = 1.5
    if (0 <= 2147483647) {
        fmt.Println(v18, v4)
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint8 = v5
        for fi21 := 0; fi21 < 7; fi21++ {
            switch (("" + (v2 + v2)) + v2) {
            case "hello":
                var v22 string = (v2 + v2)
                _ = v22
            case "a":
                fmt.Println(v18, v2)
            case "\n\t":
                v2 = v2
            default:
                switch fi21 {
                case 42:
                    var v23 int16 = 1
                    _ = v23
                case 2147483647:
                    var v24 bool = (false || false)
                    _ = v24
                }
            }
            switch (fi19 << ((42) & 63)) {
            case -417:
                if true {
                    v25 := GenS1{F0: 9223372036854775807, F1: 42, F2: -1}
                    fmt.Println(v4, fi21)
                    var v26 uint32 = 0
                    _ = v25
                    _ = v26
                } else {
                    var v27 float32 = float32(v4)
                    delete(v6, "foo bar")
                    _ = v27
                }
            case -9223372036854775808:
                var v28 int64 = 0
                _ = v28
            case -661:
                v29 := GenS1{F0: 1, F1: 0, F2: -587}
                _ = v29
            }
            v20 = (v20 / (((v20 ^ 128)) | 1))
            v18 = 0.0
            v30 := v3[(^(fi21))]
            _ = fi21
            _ = v30
        }
        switch (fi19 % (((fi19 | (fi19 % (((fi19 & fi19)) | 1)))) | 1)) {
        case -1:
            for fi31 := 0; fi31 < 3; fi31++ {
                v32 := GenS1{F0: 219, F1: 1000, F2: int64(fi31)}
                var v33 uint8 = (v5 | v20)
                v34 := GenS1{F0: -337, F1: -903, F2: 42}
                v35 := GenS1{F0: 1, F1: 1, F2: 42}
                v1[766] = 830
                var v36 float64 = 0.0
                switch 32768 {
                case 174:
                    v37 := GenS1{F0: 1, F1: 1, F2: 734}
                    _ = v37
                case 9223372036854775807:
                    v33 = v5
                default:
                    if ((fi19 * 9223372036854775807) >= fi31) {
                        var v38 int = (fi19 ^ fi19)
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        if false {
                            if (v2 == "") {
                                if true {
                                    if len(v14) > 0 {
                                        v14[0] = v38
                                    }
                                } else {
                                    fmt.Println(v18, v18, v18)
                                }
                                v39 := GenS1{F0: 1, F1: -32768, F2: -9223372036854775808}
                                v2 = v2
                                _ = v39
                            }
                        }
                        _ = v38
                    }
                    var v40 uint = 850
                    _ = v40
                }
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
                _ = v35
                _ = v36
            }
        case -818:
            var v41 string = (("hello" + ((v2 + (v2 + (v2 + ("\n\t" + (("" + v2) + "a"))))) + (v2 + (v2 + (v2 + (v2 + (v2 + "a"))))))) + "a")
            _ = v41
        default:
            var v42 int64 = -1
            _ = v42
        }
        _ = fi19
        _ = v20
    }
    v43 := GenS1{F0: -1, F1: 32767, F2: -9223372036854775808}
    var v44 int16 = 1
    var v45 bool = false
    v46 := GenS2{F0: 42, F1: -9223372036854775808, F2: 7}
    v47 := GenS2{F0: -1, F1: 0, F2: -25}
    v44 -= v44
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

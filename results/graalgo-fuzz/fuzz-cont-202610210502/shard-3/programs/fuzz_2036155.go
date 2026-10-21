package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 string }

func main() {
    v1 := []int{42}
    var v2 int64 = -1
    var v3 float32 = -2.25
    v4 := make(map[uint32]bool)
    var v5 bool = false
    _ = len(v4)
    if (32767 < 32767) {
        var v6 float64 = ((1.5 / 1.5) * (0.0 / 1.0))
        v3 -= float32(v6)
        var v7 bool = true
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 158, F1: (v2 + -224), F2: 768}
        if (!(2 < 185)) {
            v9 := GenS2{F0: 100, F1: -73}
            var v10 float32 = ((((v3 * v3) - 1e10) + float32(-2.25)) + 0.1)
            _ = v10
            _ = v9
        } else {
            if false {
                switch "\n\t" {
                case "hello":
                    var v11 uint8 = 180
                    _ = v11
                }
                if len(v1) > 0 {
                    v1[0] = 9223372036854775807
                }
                v2 = 0
            }
            fmt.Println(v2)
            v12 := GenS2{F0: 71, F1: -128}
            _ = v12
        }
        _ = v8
    }
    switch -2147483648 {
    case 2147483647:
        for fi13 := 0; fi13 < 2; fi13++ {
            v5 = (!true)
            var v14 uint16 = 363
            v14 += (v14 + (v14 % ((32767) | 1)))
            _ = fi13
            _ = v14
        }
    }
    v15 := v4[80]
    for fi16 := 0; fi16 < 19; fi16++ {
        if ((733 > 65535) || v15) {
            v2 = v2
            var v17 int32 = 682
            v3 = float32(-2.25)
            _ = v17
        } else {
            if v5 {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v3 += ((v3 - (v3 + 1.5)) + v3)
                    v3 = v3
                    _ = fi18
                }
                var v19 bool = ((42 > fi16) || true)
                if true {
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        var v20 uint32 = 0
                        _ = v20
                    case "hello":
                        v19 = (v15 || v19)
                    case "\n\t":
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 float64 = 0.0
                            if (v15 || v15) {
                                v19 = v5
                                var v23 string = "foo bar"
                                _ = v23
                            }
                            v3 -= -2.25
                            _ = fi21
                            _ = v22
                        }
                    }
                    for fi24 := 0; fi24 < 6; fi24++ {
                        if (!(((((float64(v3) + (0.0 / 1.5)) - -2.25) * 0.1) - 1.5) == 1e10)) {
                            v5 = (-43 < -1)
                            v19 = (1 == 255)
                            switch (fi24 * 0) {
                            case 128:
                                v15 = (1 >= 2147483647)
                            case 760:
                                v19 = ((true && true) || false)
                            }
                        } else {
                            if v19 {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v25 uint64 = 453
                                var v26 int16 = -824
                                switch 32767 {
                                case -173:
                                    var v27 uint8 = 2
                                    _ = v27
                                case -863:
                                    v19 = false
                                case 4294967295:
                                    var v28 int = int(v25)
                                    _ = v28
                                }
                                if v5 {
                                    v29 := v4[2147483648]
                                    var v30 uint32 = 65535
                                    if ((v26 * int16(v25)) < (v26 + v26)) {
                                        _ = len(v1)
                                        v31 := GenS3{F0: v25}
                                        _ = v31
                                    }
                                    _ = v29
                                    _ = v30
                                }
                                var v32 int8 = 1
                                _ = v25
                                _ = v26
                                _ = v32
                            }
                            var v33 int = 4294967295
                            var v34 uint64 = 0
                            _ = v33
                            _ = v34
                        }
                        _ = fi24
                    }
                    fmt.Println(v15)
                }
                _ = v19
            }
            var v35 float64 = 1.5
            _ = v35
        }
        fmt.Println(v15, fi16, v3)
        _ = fi16
    }
    switch -150 {
    case 127:
        var v36 int32 = 0
        _ = v36
    case -2147483648:
        var v37 float64 = (0.1 - -2.25)
        _ = v37
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        fmt.Println(v15, fi38, v15)
        fmt.Println(v15, v3)
        _ = fi38
    }
    fmt.Println(v15)
    var v39 uint = 756
    if false {
        if ((2 <= 211) || v15) {
            var v40 string = "\n\t"
            v39 += v39
            _ = v40
        }
    }
    v41 := GenS4{F0: "\n\t", F1: "\n\t"}
    var v42 bool = v5
    v43 := GenS2{F0: 57, F1: 1}
    var v44 int16 = 1
    var v45 float32 = ((v3 - v3) / v3)
    v46 := GenS4{F0: ("\n\t" + "hello"), F1: "a"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
